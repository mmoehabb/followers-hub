package ancillaries

import (
	"context"
	"goweb/ui/components"

	"github.com/gofiber/fiber/v2"
)

// an ancillary function whose only purpose it to follow the DRY principle!
// it shall be used in handler function wherever a notification should be
// rendered (htmx) to the user.
func Notify(c *fiber.Ctx, msg, color string) {
  components.Notification(msg, color).Render(context.Background(), c.Response().BodyWriter())
}
