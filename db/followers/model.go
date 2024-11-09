package followers

type DataModel struct{
  Email string
  DisplayName string
  Token string
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    Email: row[0].(string),
    DisplayName: row[1].(string),
    Token: row[2].(string),
  }
}

