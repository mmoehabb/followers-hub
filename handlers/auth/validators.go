package auth

func ValidateTwitchAuthBody(body *TwitchAuthBody) (bool, map[string]string) {
	ok := true
	errs := make(map[string]string)
	if body.AccessToken == "" {
		errs[HTML_INPUTS_NAMES.AccessToken] = "access_token required."
		ok = false
	}
	return ok, errs
}

func ValidateAccountAuthBody(body *AccountAuthBody) (bool, map[string]string) {
	ok := true
	errs := make(map[string]string)
	if body.RefreshToken == "" {
		errs[HTML_INPUTS_NAMES.RefreshToken] = "refresh_token required."
		ok = false
	}
	if body.Id == "" {
		errs[HTML_INPUTS_NAMES.Id] = "user id required."
		ok = false
	}
	return ok, errs
}
