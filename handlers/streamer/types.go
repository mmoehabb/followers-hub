package streamer

type LoginBody struct {
	Email string `json:"email" xml:"email" form:"email"`
}

type UpdateBody struct {
  DisplayName string `json:"display_name" xml:"display_name" form:"display_name"`
  ImgUrl string `json:"img_url" xml:"img_url" form:"img_url"`
}

var HTML_INPUTS_NAMES = struct{
  Email string
  Id string
  DisplayName string
  ImgUrl string
} {
    Email: "email",
    DisplayName: "display_name",
    ImgUrl: "img_url",
  }
