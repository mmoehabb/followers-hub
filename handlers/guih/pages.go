package guih

import (
	"context"
	"goweb/db/streamers"
	"goweb/pages"

	"github.com/gofiber/fiber/v2"
)

func IndexPage(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  data, err := streamers.Get(c.Cookies("streamer_id"))
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if (data.Id == "") {
    return c.SendStatus(fiber.StatusNotFound)
  }
  pages.Index(&data).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}

func LoginPage(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  pages.Login().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}
