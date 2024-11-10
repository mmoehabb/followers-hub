package video

type VideoBody struct {
  SectionId int `json:"section_id" xml:"section_id" form:"section_id"`
  Title string `json:"title" xml:"title" form:"title"`
  Url string `json:"url" xml:"url" form:"url"`
}
