package streamer

import (
	"net/mail"
)

func ValidateLoginBody(body *LoginBody) (bool, map[string]string) {
  ok := true
  errs := make(map[string]string)
  if body.Email == "" {
    errs[HTML_INPUTS_NAMES.Email] = "email address required."
    ok = false
  }
  if _, err := mail.ParseAddress(body.Email); err != nil {
    errs[HTML_INPUTS_NAMES.Email] = "Invalid email address."
    ok = false
  }
  return ok, errs
}

