package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"goweb/db/streamers"
)

func Auth(c *fiber.Ctx) error {
	if c.Cookies("streamer_id") == "" {
		return c.Redirect("/login")
	}
	data, err := streamers.Get(c.Cookies("streamer_id"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if c.Cookies("token") != data.AccessToken {
		return c.Redirect("/login")
	}
	return c.Next()
}
