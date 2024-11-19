package followers

type DataModel struct {
	Email       string
	DisplayName string
	Token       string
}

func parseRow(row []any) DataModel {
	return DataModel{
		Email:       row[0].(string),
		DisplayName: row[1].(string),
		Token:       row[2].(string),
	}
}

func parseModel(m *DataModel) map[string]string {
	var modelmap = make(map[string]string)
	if m.Email != "" {
		modelmap["email"] = m.Email
	}
	if m.DisplayName != "" {
		modelmap["display_name"] = m.DisplayName
	}
	if m.Token != "" {
		modelmap["token"] = m.Token
	}
	return modelmap
}
