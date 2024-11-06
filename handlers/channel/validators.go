package channel

func ValidateBody(body *ChannelBody) (bool, map[string]string) {
  var ok = true;
  var errs = make(map[string]string)
  if (body.Name == "") {
    errs["Channel Name"] = "Channel name is required"
    ok = false
  }
  if (body.PrimaryColor == "") {
    errs["primary-color"] = "primary-color is required"
    ok = false
  }
  if (body.SecondaryColor == "") {
    errs["secondary-color"] = "secondary-color is required"
    ok = false
  }
  if (body.AccentColor == "") {
    errs["accent-color"] = "accent-color is required"
    ok = false
  }
  if (body.TextColor == "") {
    errs["text-color"] = "text-color is required"
    ok = false
  }
  return ok, errs
}
