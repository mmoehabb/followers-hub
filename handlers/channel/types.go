package channel

type ChannelBody struct {
  Name string `json:"Channel Name" xml:"Channel Name" form:"Channel Name"`
  PrimaryColor string `json:"primary-color" xml:"primary-color" form:"primary-color"`
  SecondaryColor string `json:"secondary-color" xml:"secondary-color" form:"secondary-color"`
  AccentColor string `json:"accent-color" xml:"accent-color" form:"accent-color"`
  TextColor string `json:"text-color" xml:"text-color" form:"text-color"`
}

