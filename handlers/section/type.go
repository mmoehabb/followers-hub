package section

type SectionBody struct {
  ChannelId int `json:"channel_id" xml:"channel_id" form:"channel_id"`
  Name string `json:"name" xml:"name" form:"name"`
}
