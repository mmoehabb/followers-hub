package main

import (
	"github.com/gofiber/fiber/v2"

	"goweb/db"
	"goweb/middlewares"

	"goweb/handlers/auth"
	"goweb/handlers/channel"
	"goweb/handlers/section"
	"goweb/handlers/guih"
	"goweb/handlers/streamer"
)

func main() {
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

  // Streamer page endpoints
  app.Get("/hub/:streamer_id", guih.IndexPage(false))
  app.Get("/hub/:streamer_id/channels", streamer.GetChannels)
  app.Get("/hub/:channel_id<int>/sections", channel.GetSections)
  app.Get("/hub/:section_id<int/>/videos", section.GetVideos)

  // forms HTMX endpoints
  app.Get("/forms/login", guih.LoginForm)
  app.Get("/forms/channel", guih.ChannelForm)
  app.Get("/forms/section", guih.SectionForm)
  app.Get("/forms/video", guih.VideoForm)

  // Login and Authentication endpoints
  app.Get("/login", guih.LoginPage)

  app.Get("/auth/account", auth.Account) // GET method as it's dispatched by an access link
  app.Post("/auth/twitch", auth.Twitch)
  app.Post("/login/email", streamer.Login)

  // token authentication middleware
  app.Use(middlewares.Auth)
  
  // the index page; only works for logged in users
  // it used basically for every thing: adding/editing 
  // contents and more generally maintaining the "Hub"
  app.Get("/", guih.IndexPage(true))

  // POST create handlers
  app.Post("/create/channel", channel.Create)

  // components endpoints; these are mostly used for pop-up layouts
  app.Get("/component/video", guih.Video)
  app.Get("/component/comments", guih.Comments)

  app.Listen(":3000")
}
