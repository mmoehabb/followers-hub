package streamer

import (
	"context"
	"log"

	"goweb/db/channels"
	"goweb/db/streamers"
	"goweb/ui/collections"
  anc "goweb/ancillaries"

	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  streamer_id := c.Cookies("streamer_id")
  
  var body UpdateBody
  anc.Must(nil, c.BodyParser(&body))

  anc.Must(nil, streamers.Update(&streamers.DataModel{
    Id: streamer_id,
    DisplayName: body.DisplayName,
    ImgUrl: body.ImgUrl,
  }))

  anc.Notify(c, "Your info has been successfully updated.", "bg-success")
  return c.SendStatus(fiber.StatusOK)
}

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
