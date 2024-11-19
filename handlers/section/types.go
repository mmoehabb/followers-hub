package section

type SectionBody struct {
	ChannelId int    `json:"channel_id" xml:"channel_id" form:"channel_id"`
	Name      string `json:"Section Name" xml:"Section Name" form:"Section Name"`
}

var HTML_INPUTS_NAMES = struct {
	ChannelId string
	Name      string
}{
	ChannelId: "channel_id",
	Name:      "Section Name",
}
