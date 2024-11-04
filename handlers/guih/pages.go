package guih

import (
	"context"
	"goweb/db/streamers"
	"goweb/pages"

	"github.com/gofiber/fiber/v2"
)

func IndexPage(isadmin bool) func(*fiber.Ctx) error {
  return func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    data, err := streamers.Get(c.Params("streamer_id"))
    if err != nil {
      return c.SendStatus(fiber.StatusInternalServerError)
    }
    if (data.Id == "") {
      return c.SendStatus(fiber.StatusNotFound)
    }
    pages.Index(&data, isadmin).Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(200)
  }
}

func LoginPage(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  pages.Login().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}
