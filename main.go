package main

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"goweb/db"
	"goweb/pages"
	"goweb/ui/components"
)

func main() {
  // initialize a context to share data between different templ components
  ctx := context.WithValue(context.Background(), "version", "v0.0.1")

  app := fiber.New()
  app.Static("/public", "./public/")

  // shall be used once and commented afterwards,
  // and maybe completed removed in production.
  app.Get("/seed", func(c *fiber.Ctx) error {
    err := db.Seed()
    if err != nil {
      c.Status(fiber.StatusInternalServerError).SendString("internal error.")
    }
    return c.SendString("Database has been seeded.")
  })

  app.Get("/", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    pages.Index().Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Get("/login", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    pages.Login().Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Get("/component/video", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    if c.Query("url") == "" {
      return c.SendStatus(fiber.ErrBadRequest.Code)
    }
    components.VideoOverlay(c.Query("url")).Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Get("/component/chat", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    components.ChatDrawer().Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Listen(":3000")
}
