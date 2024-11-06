package channel

import (
	"goweb/ancillaries"
	"goweb/db/channels"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
  body := new(ChannelBody)
  if err := c.BodyParser(body); err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
  }
  ok, errs := ValidateBody(body)
  if ok == false {
    ancillaries.Notify(c, "Ensure all values have been set.", "bg-error")
    return c.Status(fiber.StatusBadRequest).JSON(errs)
  }
  ok, err := channels.Add(&channels.DataModel{
    StreamerId: c.Cookies("streamer_id"),
    Name: body.Name,
    PrimaryColor: body.PrimaryColor,
    SecondaryColor: body.SecondaryColor,
    AccentColor: body.AccentColor,
    TextColor: body.TextColor,
  })
  if err != nil {
    ancillaries.Notify(c, "Something went wrong!", "bg-error")
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if ok == false {
    ancillaries.Notify(c, "Channel already found!", "bg-error")
    return c.SendStatus(fiber.StatusConflict)
  }
  ancillaries.Notify(c, "Channel has been successfully created.", "bg-success")
  return c.SendStatus(fiber.StatusCreated)
}
