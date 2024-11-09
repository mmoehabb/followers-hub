package streamers

type DataModel struct{
  Id string
  DisplayName string
  ImgUrl string
  AccessToken string
  RefreshToken string
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    Id: row[0].(string),
    DisplayName: row[1].(string),
    ImgUrl: row[2].(string),
    AccessToken: row[3].(string),
    RefreshToken: row[4].(string),
  }
}
