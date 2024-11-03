package channel

import (
	"context"
	"log"

	"goweb/db/channels"
	"goweb/ui/collections"

	"github.com/gofiber/fiber/v2"
)

func GetChannels(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  channels, err := channels.GetChannelsOf(c.Params("streamer_id"))
  if err != nil {
    log.Println(err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  collections.Channels(channels).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(fiber.StatusOK)
}