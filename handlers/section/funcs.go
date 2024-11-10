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

func Create(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  body := new(SectionBody)
  if err := c.BodyParser(body); err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
  }

  ok, errs := ValidateBody(body)
  if ok == false {
    anc.Notify(c, "Ensure all values have been set.", "bg-error")
    return c.Status(fiber.StatusBadRequest).JSON(errs)
  }

  ok = anc.Must(sections.Add(&sections.DataModel{
    ChannelId: body.ChannelId,
    Name: body.Name,
  })).(bool)

  if ok == false {
    anc.Notify(c, "Section already found!", "bg-error")
    return c.SendStatus(fiber.StatusConflict)
  }
  anc.Notify(c, "Section has been successfully created.", "bg-success")
  return c.SendStatus(fiber.StatusCreated)
}

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
