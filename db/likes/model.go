package likes

type DataModel struct{
  VideoId int
  FollowerEmail string
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    VideoId: int(row[0].(int32)),
    FollowerEmail: row[1].(string),
  }
}

