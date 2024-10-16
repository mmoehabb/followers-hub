package main

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"goweb/db"
	"goweb/db/streamers"

	"goweb/handlers/auth"
	"goweb/handlers/streamer"

	"goweb/pages"
	"goweb/ui/components"
	"goweb/ui/forms"
)

func main() {
  // initialize a context to share data between different templ components
  ctx := context.WithValue(context.Background(), "version", "v0.0.1")

  app := fiber.New()
  app.Static("/public", "./public/")

  // seed endpoint: it shall be used once and commented afterwards,
  // and maybe completelly removed in production.
  app.Get("/seed", func(c *fiber.Ctx) error {
    err := db.Seed()
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).SendString("internal error: " + err.Error())
    }
    return c.SendString("Database has been seeded.")
  })

  app.Get("/login", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    pages.Login().Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Get("/forms/login", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    forms.LoginForm(nil).Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Post("/login/email", streamer.Login)
  app.Post("/auth/twitch", auth.Twitch)
  app.Post("/auth/account", auth.Account)

  app.Use(func(c *fiber.Ctx) error {
    if c.Cookies("streamer_id") == "" {
      return c.Redirect("/login")
    }
    data, err := streamers.Get(c.Cookies("streamer_id"))
    if err != nil {
      return c.SendStatus(fiber.StatusInternalServerError)
    }
    if c.Cookies("token") != data.AccessToken {
      return c.Redirect("/login")
    }
    return c.Next()
  })
  
  app.Get("/", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    pages.Index().Render(ctx, c.Response().BodyWriter())
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

  app.Get("/component/comments", func(c *fiber.Ctx) error {
    c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    components.CommentsDrawer().Render(ctx, c.Response().BodyWriter())
    return c.SendStatus(200)
  })

  app.Listen(":3000")
}
