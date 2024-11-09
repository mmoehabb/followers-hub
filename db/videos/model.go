package videos

type DataModel struct{
  Id int
  SectionId int
  Title string
  Url string
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    Id: int(row[0].(int32)),
    SectionId: int(row[1].(int32)),
    Title: row[2].(string),
    Url: row[3].(string),
  }
}
