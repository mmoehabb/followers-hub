package video

type VideoBody struct {
	SectionId int    `json:"section_id" xml:"section_id" form:"section_id"`
	Title     string `json:"Video Title" xml:"Video Title" form:"Video Title"`
	Url       string `json:"Embedded Url" xml:"Embedded Url" form:"Embedded Url"`
}

var HTML_INPUTS_NAMES = struct {
	SectionId string
	Title     string
	Url       string
}{
	SectionId: "section_id",
	Title:     "Video Title",
	Url:       "Embedded Url",
}
