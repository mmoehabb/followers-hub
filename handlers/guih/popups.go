package guih

import (
	"context"
	"goweb/ui/components"

	"github.com/gofiber/fiber/v2"
)

func Video(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  if c.Query("url") == "" {
    return c.SendStatus(fiber.ErrBadRequest.Code)
  }
  components.VideoOverlay(c.Query("url")).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}

func Comments(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  components.CommentsDrawer().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}
