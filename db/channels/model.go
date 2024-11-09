package channels

type DataModel struct{
  Id int
  StreamerId string
  Name string
  PrimaryColor string
  SecondaryColor string
  AccentColor string
  TextColor string
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    Id: int(row[0].(int32)),
    StreamerId: row[1].(string),
    Name: row[2].(string),
    PrimaryColor: row[3].(string),
    SecondaryColor: row[4].(string),
    AccentColor: row[5].(string),
    TextColor: row[6].(string),
  }
}

