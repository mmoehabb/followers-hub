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

func parseModel(m *DataModel) map[string]string {
  var modelMap = make(map[string]string)
  if m.Id != "" {
    modelMap["id"] = m.Id
  }
  if m.DisplayName != "" {
    modelMap["display_name"] = m.DisplayName
  }
  if m.ImgUrl != "" {
    modelMap["img_url"] = m.ImgUrl
  }
  if m.AccessToken != "" {
    modelMap["access_token"] = m.AccessToken
  }
  if m.RefreshToken != "" {
    modelMap["refresh_token"] = m.RefreshToken
  }
  return modelMap
}
