package video

import (
	anc "goweb/ancillaries"
	"goweb/db/videos"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  body := new(VideoBody)
  if err := c.BodyParser(body); err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
  }

  ok, errs := ValidateBody(body)
  if ok == false {
    anc.Notify(c, "Ensure all values have been set.", "bg-error")
    return c.Status(fiber.StatusBadRequest).JSON(errs)
  }

  ok = anc.Must(videos.Add(&videos.DataModel{
    SectionId: body.SectionId,
    Title: body.Title,
    Url: body.Url,
  })).(bool)

  if ok == false {
    anc.Notify(c, "Video already found!", "bg-error")
    return c.SendStatus(fiber.StatusConflict)
  }
  anc.Notify(c, "Video has been successfully created.", "bg-success")
  return c.SendStatus(fiber.StatusCreated)
}
