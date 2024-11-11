package section

type SectionBody struct {
  ChannelId int `json:"channel_id" xml:"channel_id" form:"channel_id"`
  Name string `json:"Section Name" xml:"Section Name" form:"Section Name"`
}
