package sections

type DataModel struct{
  Id int
  ChannelId int
  Name string
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    Id: int(row[0].(int32)),
    ChannelId: int(row[1].(int32)),
    Name: row[2].(string),
  }
}
