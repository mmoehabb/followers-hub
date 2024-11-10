package section

func ValidateBody(body *SectionBody) (bool, map[string]string) {
  var ok bool = true
  var errs = make(map[string]string)
  if body.ChannelId == 0 {
    ok = false
    errs["channel_id"] = "channel_id must be specified."
  }
  if body.Name == "" {
    ok = false
    errs["name"] = "name field is required."
  }
  return ok, errs
}
