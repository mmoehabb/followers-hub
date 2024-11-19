package channel

import (
	"context"
	anc "goweb/ancillaries"
	"goweb/db/channels"
	"goweb/db/sections"
	"goweb/ui/collections"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetSections(admin bool) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		defer anc.Recover(c)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		channel_id := anc.Must(strconv.Atoi(c.Params("channel_id"))).(int)

		found := anc.Must(channels.Get(channel_id)).(channels.DataModel)
		if found.Name == "" {
			return c.SendStatus(fiber.StatusNotFound)
		}

		list := anc.Must(sections.GetSectionsOf(channel_id)).([]sections.DataModel)
		collections.Sections(list, channel_id, admin).Render(context.Background(), c.Response().BodyWriter())
		return c.SendStatus(fiber.StatusOK)
	}
}

func Create(c *fiber.Ctx) error {
	defer anc.Recover(c)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	body := new(ChannelBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ok, errs := ValidateBody(body)
	if ok == false {
		anc.Notify(c, "Ensure all values have been set.", "bg-error")
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	ok = anc.Must(channels.Add(&channels.DataModel{
		StreamerId:     c.Cookies("streamer_id"),
		Name:           body.Name,
		PrimaryColor:   body.PrimaryColor,
		SecondaryColor: body.SecondaryColor,
		AccentColor:    body.AccentColor,
		TextColor:      body.TextColor,
	})).(bool)

	if ok == false {
		anc.Notify(c, "Channel already found!", "bg-error")
		return c.SendStatus(fiber.StatusConflict)
	}
	anc.Notify(c, "Channel has been successfully created.", "bg-success")
	return c.SendStatus(fiber.StatusCreated)
}

func Update(c *fiber.Ctx) error {
	defer anc.Recover(c)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	body := new(ChannelBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ok, errs := ValidateBody(body)
	if ok == false {
		anc.Notify(c, "Ensure all values have been set.", "bg-error")
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	channel_id := anc.Must(c.ParamsInt("channel_id")).(int)
	// check if the channel already exists and it the client
	// permitted to update it or not
	found_channel := anc.Must(channels.Get(channel_id)).(*channels.DataModel)
	if found_channel.StreamerId == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if found_channel.StreamerId != c.Cookies("streamer_id") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	data := &channels.DataModel{
		StreamerId:     c.Cookies("streamer_id"),
		Name:           body.Name,
		PrimaryColor:   body.PrimaryColor,
		SecondaryColor: body.SecondaryColor,
		AccentColor:    body.AccentColor,
		TextColor:      body.TextColor,
	}
	anc.Must(nil, channels.Update(channel_id, data))

	return c.SendStatus(fiber.StatusOK)
}

func Remove(c *fiber.Ctx) error {
	defer anc.Recover(c)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	channel_id := anc.Must(c.ParamsInt("channel_id")).(int)
	// check if the channel already exists and it the client
	// permitted to update it or not
	found_channel := anc.Must(channels.Get(channel_id)).(*channels.DataModel)
	if found_channel.StreamerId == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if found_channel.StreamerId != c.Cookies("streamer_id") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	anc.Must(nil, channels.Remove(channel_id))
	return c.SendStatus(fiber.StatusOK)
}
