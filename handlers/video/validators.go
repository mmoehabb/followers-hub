package video

func ValidateBody(body *VideoBody) (bool, map[string]string) {
	var ok bool = true
	var errs = make(map[string]string)
	if body.SectionId == 0 {
		ok = false
		errs[HTML_INPUTS_NAMES.SectionId] = "section_id must be specified."
	}
	if body.Title == "" {
		ok = false
		errs[HTML_INPUTS_NAMES.Title] = "title field is required."
	}
	if body.Url == "" {
		ok = false
		errs[HTML_INPUTS_NAMES.Url] = "url field is required."
	}
	return ok, errs
}
