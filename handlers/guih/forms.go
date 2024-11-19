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
	channel_id, err := c.ParamsInt("channel_id")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	forms.SectionForm(channel_id).Render(context.Background(), c.Response().BodyWriter())
	return c.SendStatus(fiber.StatusOK)
}

func VideoForm(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	section_id, err := c.ParamsInt("section_id")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	forms.VideoForm(section_id).Render(context.Background(), c.Response().BodyWriter())
	return c.SendStatus(200)
}
