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
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
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

func Update(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  body := new(VideoBody)
  if err := c.BodyParser(body); err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
  }

  ok, errs := ValidateBody(body)
  if ok == false {
    anc.Notify(c, "Ensure all values have been set.", "bg-error")
    return c.Status(fiber.StatusBadRequest).JSON(errs)
  }

  video_id := anc.Must(c.ParamsInt("video_id")).(int)
  // check if the video already exists and it the client
  // permitted to update it or not
  owner_id := anc.Must(videos.GetSteamerId(video_id)).(string)
  if owner_id == "" {
    return c.SendStatus(fiber.StatusNotFound)
  } 
  if owner_id != c.Cookies("streamer_id") {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

  data := &videos.DataModel{
    SectionId: body.SectionId,
    Title: body.Title,
    Url: body.Url,
  }
  anc.Must(nil, videos.Update(video_id, data))

  return c.SendStatus(fiber.StatusOK)
}

func Remove(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

  video_id := anc.Must(c.ParamsInt("video_id")).(int)
  // check if the section already exists and it the client
  // permitted to update it or not
  owner_id := anc.Must(videos.GetSteamerId(video_id)).(string)
  if owner_id == "" {
    return c.SendStatus(fiber.StatusNotFound)
  } 
  if owner_id != c.Cookies("streamer_id") {
    return c.SendStatus(fiber.StatusUnauthorized)
  }

  anc.Must(nil, videos.Remove(video_id))
  return c.SendStatus(fiber.StatusOK)
}

