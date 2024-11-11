package video

type VideoBody struct {
  SectionId int `json:"section_id" xml:"section_id" form:"section_id"`
  Title string `json:"Video Title" xml:"Video Title" form:"Video Title"`
  Url string `json:"Embedded Url" xml:"Embedded Url" form:"Embedded Url"`
}
