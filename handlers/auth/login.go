package auth

import (
	"encoding/json"
	"goweb/db/streamers"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// login hanlder for fiber endpoint /login
// it expects a POST request
func LoginWithTwitch(c *fiber.Ctx) error {
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
    err := streamers.UpdateTokens(claims.Id, body.AccessToken, body.RefreshToken)
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

func LoginWithEmail(c *fiber.Ctx) error {
  return c.SendString("Nothing to be done.")
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

