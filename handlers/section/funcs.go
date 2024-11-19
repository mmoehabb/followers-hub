package section

import (
	"context"
	"strconv"

	anc "goweb/ancillaries"
	"goweb/db/sections"
	"goweb/db/videos"
	"goweb/ui/collections"

	"github.com/gofiber/fiber/v2"
)

func GetVideos(admin bool) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		defer anc.Recover(c)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		section_id := anc.Must(strconv.Atoi(c.Params("section_id"))).(int)

		found := anc.Must(sections.Get(section_id)).(sections.DataModel)
		if found.Name == "" {
			return c.SendStatus(fiber.StatusNotFound)
		}

		list := anc.Must(videos.GetVideosOf(section_id)).([]videos.DataModel)
		collections.Videos(list, section_id, admin).Render(context.Background(), c.Response().BodyWriter())
		return c.SendStatus(fiber.StatusOK)
	}
}

func Create(c *fiber.Ctx) error {
	defer anc.Recover(c)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	body := new(SectionBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ok, errs := ValidateBody(body)
	if ok == false {
		anc.Notify(c, "Ensure all values have been set.", "bg-error")
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	ok = anc.Must(sections.Add(&sections.DataModel{
		ChannelId: body.ChannelId,
		Name:      body.Name,
	})).(bool)

	if ok == false {
		anc.Notify(c, "Section already found!", "bg-error")
		return c.SendStatus(fiber.StatusConflict)
	}
	anc.Notify(c, "Section has been successfully created.", "bg-success")
	return c.SendStatus(fiber.StatusCreated)
}

func Update(c *fiber.Ctx) error {
	defer anc.Recover(c)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	body := new(SectionBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ok, errs := ValidateBody(body)
	if ok == false {
		anc.Notify(c, "Ensure all values have been set.", "bg-error")
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	section_id := anc.Must(c.ParamsInt("section_id")).(int)
	// check if the section already exists and it the client
	// permitted to update it or not
	owner_id := anc.Must(sections.GetSteamerId(section_id)).(string)
	if owner_id == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if owner_id != c.Cookies("streamer_id") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	data := &sections.DataModel{
		ChannelId: body.ChannelId,
		Name:      body.Name,
	}
	anc.Must(nil, sections.Update(section_id, data))

	return c.SendStatus(fiber.StatusOK)
}

func Remove(c *fiber.Ctx) error {
	defer anc.Recover(c)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	sections_id := anc.Must(c.ParamsInt("sections_id")).(int)
	// check if the section already exists and it the client
	// permitted to update it or not
	owner_id := anc.Must(sections.GetSteamerId(sections_id)).(string)
	if owner_id == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if owner_id != c.Cookies("streamer_id") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	anc.Must(nil, sections.Remove(sections_id))
	return c.SendStatus(fiber.StatusOK)
}
