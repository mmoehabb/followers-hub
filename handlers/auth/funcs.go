package auth

import (
	"context"
	"encoding/json"
	"goweb/db/streamers"
	"goweb/ui/forms"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// login hanlder for fiber endpoint /login
// it expects a POST request
func Twitch(c *fiber.Ctx) error {
  body := new(TwitchAuthBody)
  if err := c.BodyParser(body); err != nil {
    return err
  }
  // verify that body.AccessToken is passed
  ok, errs := ValidateTwitchAuthBody(body)
  if ok == false {
    return c.Status(fiber.StatusBadRequest).JSON(errs)
  }
  // otherwise, get user (streamer) info
  claims, httpStatusInt := getTwitchClaims(body.AccessToken)
  if httpStatusInt != 0 {
    return c.SendStatus(httpStatusInt)
  }
  // if user id already exists in the database update the AccessToken
  // otherwise add new entity and generate it's channel
  found, err := streamers.Exists(claims.Id)
  if err != nil {
    log.Println("internal error: ", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if found {
    err := streamers.Update(&streamers.DataModel{
      Id: claims.Id,
      AccessToken: body.AccessToken,
      RefreshToken: body.RefreshToken,
    })
    if err != nil {
      log.Println("internal error: ", err)
      return c.SendStatus(fiber.StatusInternalServerError)
    }
  } else {
    newStreamerData := streamers.DataModel{
      Id: claims.Id,
      DisplayName: claims.DisplayName,
      ImgUrl: claims.ImgUrl,
      AccessToken: body.AccessToken,
      RefreshToken: body.RefreshToken,
    }
    err := streamers.Add(&newStreamerData)
    if err != nil {
      log.Println("internal error: ", err)
      return c.SendStatus(fiber.StatusInternalServerError)
    }
  }
  c.Cookie(&fiber.Cookie{
    Name: "token",
    Value: body.AccessToken,
  })
  c.Cookie(&fiber.Cookie{
    Name: "streamer_id",
    Value: claims.Id,
  })
  return c.Redirect("/")
}

func Account(c *fiber.Ctx) error {
  body := new(AccountAuthBody)
  if err := c.ParamsParser(body); err != nil {
    log.Println(err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  ok, errs := ValidateAccountAuthBody(body)
  if ok == false {
    return c.Status(fiber.StatusBadRequest).JSON(errs)
  }
  // if the the user id exists in the db and the access token is valid
  // set cookies and redirect the client to the root.
  data, err := streamers.Get(body.Id)
  if err != nil {
    log.Println(err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if data.AccessToken == body.AccessToken {
    c.Cookie(&fiber.Cookie{
      Name: "token",
      Value: body.AccessToken,
    })
    c.Cookie(&fiber.Cookie{
      Name: "streamer_id",
      Value: body.Id,
    })
    return c.Redirect("/")
  }
  // otherwise render the login/register form
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  forms.LoginForm(nil).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}

// returns TwitchClaims struct and an http status error integer
func getTwitchClaims(accessToken string) (TwitchClaims, int) {
  req, err := http.NewRequest("GET", `https://id.twitch.tv/oauth2/userinfo?claims={"userinfo":{"preferred_username":null,"picture":null}}`, nil)
  if err != nil {
    return TwitchClaims{}, fiber.StatusInternalServerError
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer " + accessToken)
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Println("internal error: ", err)
    return TwitchClaims{}, fiber.StatusInternalServerError
  }
  defer resp.Body.Close()
  buf, err := io.ReadAll(resp.Body)
  info := TwitchClaims{}
  json.Unmarshal(buf, &info)
  return info, 0
}
