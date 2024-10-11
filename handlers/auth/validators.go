package auth

func ValidateTwitchAuthBody(body *TwitchAuthBody) (bool, map[string]string) {
  ok := true
  errs := make(map[string]string)
  if body.AccessToken == "" {
    errs["access_token"] = "access_token required."
    ok = false
  }
  return ok, errs
}

func ValidateAccountAuthBody(body *AccountAuthBody) (bool, map[string]string) {
  ok := true
  errs := make(map[string]string)
  if body.AccessToken == "" {
    errs["access_token"] = "access_token required."
    ok = false
  }
  if body.Id == "" {
    errs["id"] = "user id required."
    ok = false
  }
  return ok, errs
}

