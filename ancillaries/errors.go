package ancillaries

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ancillary function used to clean the code in handlers;
// by wrapping db methods and return the result if no error
// is found, or panic otherwise.
func Must(res any, err error) any {
	if err != nil {
		panic(err)
	}
	return res
}

// ancillary function that shall be used (deferred) in the
// beginning of functions where anc.Must is used within.
func Recover(c *fiber.Ctx) error {
	if r := recover(); r != nil {
		log.Println(r)
		Notify(c, "Something went wrong!", "bg-error")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return nil
}
