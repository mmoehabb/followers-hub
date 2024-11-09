package section

import (
	"context"
	"strconv"

	"goweb/db/sections"
	"goweb/db/videos"
	"goweb/ui/collections"
	anc "goweb/ancillaries"

	"github.com/gofiber/fiber/v2"
)

func GetVideos(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  section_id := anc.Must(strconv.Atoi(c.Params("section_id"))).(int)

  found := anc.Must(sections.Get(section_id)).(sections.DataModel)
  if found.Name == "" {
    return c.SendStatus(fiber.StatusNotFound)
  }

  list := anc.Must(videos.GetVideosOf(section_id)).([]videos.DataModel)
  collections.Videos(list).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(fiber.StatusOK)
}
