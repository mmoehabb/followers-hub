package streamer

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"goweb/db/streamers"
	"goweb/ui/forms"
)

func Login(c *fiber.Ctx) error {
  body := new(LoginBody)
  if err := c.BodyParser(body); err != nil {
    return err
  }
  ok, errs := ValidateLoginBody(body)
  if ok == false {
    forms.LoginForm(errs).Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(fiber.StatusBadRequest)
  }
  // Check if the user id already exists;
  found, err := streamers.Exists(body.Email)
  if err != nil {
    log.Println(err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  // if it does exist update RefreshToken and send a mail with /auth/* link accordingly
  if found {
    err = streamers.Update(&streamers.DataModel{
      Id: body.Email,
      RefreshToken: generateToken(),
    })
    if err != nil {
      log.Println(err)
      return c.SendStatus(fiber.StatusInternalServerError)
    }
    return c.SendStatus(fiber.StatusOK)
  }
  // otherwise generate an AccessToken, register the user, and send a mail with /auth/* link.
  token := generateToken()
  err = streamers.Add(&streamers.DataModel{
    Id: body.Email, 
    DisplayName: "MyNameIsJeff",
    ImgUrl: "/images/user.jpg",
    AccessToken: token,
    RefreshToken: token,
  })
  if err != nil {
    log.Println(err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  return c.SendStatus(fiber.StatusOK)
}

func generateToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return fmt.Sprintf("%x", b)
}
