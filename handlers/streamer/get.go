package streamer

import (
	"context"
	"log"

	"goweb/db/channels"
	"goweb/ui/collections"

	"github.com/gofiber/fiber/v2"
)

func GetChannels(admin bool) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		streamer_id := c.Params("streamer_id")
		if admin {
			streamer_id = c.Cookies("streamer_id")
		}

		channels, err := channels.GetChannelsOf(streamer_id)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		collections.Channels(channels, admin).Render(context.Background(), c.Response().BodyWriter())
		return c.SendStatus(fiber.StatusOK)
	}
}
