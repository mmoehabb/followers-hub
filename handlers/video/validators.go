package video

func ValidateBody(body *VideoBody) (bool, map[string]string) {
  var ok bool = true
  var errs = make(map[string]string)
  if body.SectionId == 0 {
    ok = false
    errs["section_id"] = "section_id must be specified."
  }
  if body.Title == "" {
    ok = false
    errs["title"] = "title field is required."
  }
  if body.Url == "" {
    ok = false
    errs["url"] = "url field is required."
  }
  return ok, errs
}
