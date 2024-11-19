package streamer

type LoginBody struct {
	Email string `json:"email" xml:"email" form:"email"`
}

var HTML_INPUTS_NAMES = LoginBody{
	Email: "email",
}
