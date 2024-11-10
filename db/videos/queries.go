package videos

import (
	"errors"
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
    return DataModel{}, errors.New("data not found.")
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

func Remove(id int) error {
  _, err := db.Query("DELETE FROM videos WHERE id=$1", id)
  return err
}
