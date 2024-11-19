package sections

import "strconv"

type DataModel struct {
	Id        int
	ChannelId int
	Name      string
}

func parseRow(row []any) DataModel {
	return DataModel{
		Id:        int(row[0].(int32)),
		ChannelId: int(row[1].(int32)),
		Name:      row[2].(string),
	}
}

func parseModel(m *DataModel) map[string]string {
	var modelmap = make(map[string]string)
	if m.Id != 0 {
		modelmap["id"] = strconv.Itoa(m.Id)
	}
	if m.ChannelId != 0 {
		modelmap["channel_id"] = strconv.Itoa(m.ChannelId)
	}
	if m.Name != "" {
		modelmap["name"] = m.Name
	}
	return modelmap
}
