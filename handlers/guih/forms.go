package guih

import (
	"context"
	"goweb/ui/forms"

	"github.com/gofiber/fiber/v2"
)

func LoginForm(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  forms.LoginForm(nil).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}

func ChannelForm(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  forms.ChannelForm().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}

func SectionForm(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  forms.SectionForm().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}
