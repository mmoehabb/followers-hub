package videos

import (
  anc "goweb/ancillaries"
	"goweb/db"
)

func Add(d *DataModel) (bool, error) {
  res, err := db.SeqQuery("SELECT * FROM videos WHERE section_id=$1 AND title=$2", d.SectionId, d.Title)
  if len(res) != 0 {
    db.Disconnect()
    return false, nil
  }
  _, err = db.Query(
    "INSERT INTO videos VALUES ($1, $2, $3, $4)", 
    d.Id, d.SectionId, d.Title, d.Url,
  )
  if err != nil {
    return false, err
  }
  return true, nil
}

func Get(id int) (DataModel, error) {
  res, err := db.Query("SELECT * FROM videos WHERE id=$1", id)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, nil
  }
  row := res[0].([]any)
  obj := parseRow(row)
  return obj, nil
}

func GetVideosOf(section_id int) ([]DataModel, error) {
  res, err := db.Query("SELECT * FROM videos WHERE section_id=$1", section_id)
  list := make([]DataModel, len(res))
  for i, row := range res {
    list[i] = parseRow(row.([]any))
  }
  return list, err
}

func GetSteamerId(video_id int) (string, error) {
  res, err := db.SeqQuery("SELECT streamer_id FROM videos v JOIN sections s ON v.section_id = s.id JOIN channels c ON s.channel_id = c.id WHERE s.id=$1", video_id)
  if err != nil {
    return "", err
  }
  if len(res) == 0 {
    return "", nil
  }
  row := res[0].([]any)
  return row[0].(string), nil
}

func Update(id int, data *DataModel) error {
  data.Id = id
  query, values := anc.GenUpdateQuery("videos", parseModel(data), "id")
  _, err := db.Query(query, values...)
  return err
}

func Remove(id int) error {
  _, err := db.Query("DELETE FROM videos WHERE id=$1", id)
  return err
}
