package streamer

import (
	"net/mail"
)

func ValidateLoginBody(body *LoginBody) (bool, map[string]string) {
  ok := true
  errs := make(map[string]string)
  if body.Email == "" {
    errs["email"] = "email address required."
    ok = false
  }
  if _, err := mail.ParseAddress(body.Email); err != nil {
    errs["email"] = "Invalid email address."
    ok = false
  }
  return ok, errs
}

