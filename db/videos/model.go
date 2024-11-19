package videos

import "strconv"

type DataModel struct {
	Id        int
	SectionId int
	Title     string
	Url       string
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:        int(row[0].(int32)),
		SectionId: int(row[1].(int32)),
		Title:     row[2].(string),
		Url:       row[3].(string),
	}
}

func parseModel(m *DataModel) map[string]string {
	var modelmap = make(map[string]string)
	if m.Id != 0 {
		modelmap["id"] = strconv.Itoa(m.Id)
	}
	if m.SectionId != 0 {
		modelmap["section_id"] = strconv.Itoa(m.SectionId)
	}
	if m.Id != 0 {
		modelmap["title"] = m.Title
	}
	if m.Id != 0 {
		modelmap["url"] = m.Url
	}
	return modelmap
}
