package channels

import "strconv"

type DataModel struct {
	Id             int
	StreamerId     string
	Name           string
	PrimaryColor   string
	SecondaryColor string
	AccentColor    string
	TextColor      string
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:             int(row[0].(int32)),
		StreamerId:     row[1].(string),
		Name:           row[2].(string),
		PrimaryColor:   row[3].(string),
		SecondaryColor: row[4].(string),
		AccentColor:    row[5].(string),
		TextColor:      row[6].(string),
	}
}

func parseModel(m *DataModel) map[string]string {
	var modelmap = make(map[string]string)
	if m.Id != 0 {
		modelmap["id"] = strconv.Itoa(m.Id)
	}
	if m.StreamerId != "" {
		modelmap["streamer_id"] = m.StreamerId
	}
	if m.Name != "" {
		modelmap["name"] = m.Name
	}
	if m.PrimaryColor != "" {
		modelmap["primary_color"] = m.PrimaryColor
	}
	if m.SecondaryColor != "" {
		modelmap["secondary_color"] = m.SecondaryColor
	}
	if m.AccentColor != "" {
		modelmap["accent_color"] = m.AccentColor
	}
	if m.TextColor != "" {
		modelmap["text_color"] = m.TextColor
	}
	return modelmap
}
