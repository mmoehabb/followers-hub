package auth

func ValidateTwitchAuthBody(body *TwitchAuthBody) (bool, map[string]string) {
  ok := true
  errs := make(map[string]string)
  if body.AccessToken == "" {
    errs["id"] = "no access_token found."
    ok = false
  }
  return ok, errs
}

