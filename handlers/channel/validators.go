package channel

func ValidateBody(body *ChannelBody) (bool, map[string]string) {
	var ok = true
	var errs = make(map[string]string)
	if body.Name == "" {
		errs[HTML_INPUTS_NAMES.Name] = "Channel name is required"
		ok = false
	}
	if body.PrimaryColor == "" {
		errs[HTML_INPUTS_NAMES.PrimaryColor] = "primary-color is required"
		ok = false
	}
	if body.SecondaryColor == "" {
		errs[HTML_INPUTS_NAMES.SecondaryColor] = "secondary-color is required"
		ok = false
	}
	if body.AccentColor == "" {
		errs[HTML_INPUTS_NAMES.AccentColor] = "accent-color is required"
		ok = false
	}
	if HTML_INPUTS_NAMES.TextColor == "" {
		errs["text-color"] = "text-color is required"
		ok = false
	}
	return ok, errs
}
