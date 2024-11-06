package streamer

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"goweb/ancillaries"
	"goweb/db/streamers"
	"goweb/ui/components"
)

// used by local check and recoverf function
type pvalue struct{
  Msg string
  Code int
}

// an ancillary function to clean Login func handler a little
func check(err error, statuscode int) {
  if err != nil {
    log.Println(err.Error())
    panic(&pvalue{err.Error(), statuscode})
  }
}

// an ancillary function to clean Login func handler a little
func recoverf(c *fiber.Ctx) error {
  if r := recover(); r != nil {
    components.Notification(r.(*pvalue).Msg, "bg-error").
      Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(r.(*pvalue).Code)
  }
  return nil
}

func Login(c *fiber.Ctx) error {
  defer recoverf(c)

  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  body := new(LoginBody)
  check(c.BodyParser(body), fiber.StatusBadRequest)

  ok, _ := ValidateLoginBody(body)
  if ok == false {
    ancillaries.Notify(c, "Ensure the email is valid.", "bg-success")
    return c.SendStatus(fiber.StatusBadRequest)
  }

  // Check if the user id already exists;
  found, err := streamers.Exists(body.Email)
  check(err, fiber.StatusInternalServerError)

  // if it does exist update RefreshToken and send a mail with /auth/* link accordingly
  newtoken := generateToken()
  if found {
    err = streamers.Update(&streamers.DataModel{
      Id: body.Email,
      RefreshToken: newtoken,
    })
    check(err, fiber.StatusInternalServerError)

    err := sendAuthMail(body.Email, newtoken)
    check(err, fiber.StatusInternalServerError)

    okmsg := "Checkout your email inbox; a mail with access link shall be sent to you."
    ancillaries.Notify(c, okmsg, "bg-success")

    return c.SendStatus(fiber.StatusOK)
  }
  // otherwise generate an AccessToken, register the user, and send a mail with /auth/* link.
  err = streamers.Add(&streamers.DataModel{
    Id: body.Email, 
    DisplayName: "Ross Geller",
    ImgUrl: "/public/images/user.jpg",
    AccessToken: newtoken,
    RefreshToken: newtoken,
  })
  check(err, fiber.StatusInternalServerError)

  err = sendAuthMail(body.Email, newtoken)
  check(err, fiber.StatusInternalServerError)

  okmsg := "Checkout your email inbox; a mail with access link shall be sent to you."
  ancillaries.Notify(c, okmsg, "bg-success")

  return c.SendStatus(fiber.StatusOK)
}

func sendAuthMail(address string, token string) error {
  msgbody := fmt.Sprintf(`
    <html>
      <body>
        <p>Use the below hyper-link to get access to your FollowersHub account.</p>
        <p style="color: red;">Never share it with anyone.</p>
        <a src='http://localhost:3000/auth/account?id=%s&refresh_token=%s'>Click here to login.</a>
      </body>
    </html>`, address, token)
  // @TODO: use SMTP server to send the message
  fmt.Println(msgbody)
  return nil
}

func generateToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return fmt.Sprintf("%x", b)
}
