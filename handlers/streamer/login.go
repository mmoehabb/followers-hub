package streamer

import (
	"crypto/rand"
	"fmt"

	"github.com/gofiber/fiber/v2"

	anc "goweb/ancillaries"
	"goweb/db/streamers"
)

func Login(c *fiber.Ctx) error {
	defer anc.Recover(c)

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	body := new(LoginBody)
	if err := c.BodyParser(body); err != nil {
		anc.Notify(c, "Ensure that you have supplied valid data.", "bg-error")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	ok, _ := ValidateLoginBody(body)
	if ok == false {
		anc.Notify(c, "Ensure the email is valid.", "bg-success")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if the user id already exists;
	found := anc.Must(streamers.Exists(body.Email)).(bool)

	// if it does exist update RefreshToken and send a mail with /auth/* link accordingly
	newtoken := generateToken()
	if found {
		anc.Must(nil, streamers.Update(&streamers.DataModel{
			Id:           body.Email,
			RefreshToken: newtoken,
		}))
		anc.Must(nil, sendAuthMail(body.Email, newtoken))

		okmsg := "Checkout your email inbox; a mail with access link shall be sent to you."
		anc.Notify(c, okmsg, "bg-success")
		return c.SendStatus(fiber.StatusOK)
	}
	// otherwise generate an AccessToken, register the user, and send a mail with /auth/* link.
	anc.Must(nil, streamers.Add(&streamers.DataModel{
		Id:           body.Email,
		DisplayName:  "Ross Geller",
		ImgUrl:       "/public/images/user.jpg",
		AccessToken:  newtoken,
		RefreshToken: newtoken,
	}))
	anc.Must(nil, sendAuthMail(body.Email, newtoken))

	okmsg := "Checkout your email inbox; a mail with access link shall be sent to you."
	anc.Notify(c, okmsg, "bg-success")
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
