package sections

import (
	"errors"
	"goweb/db"
)

func Add(d *DataModel) (bool, error) {
  res, err := db.SeqQuery("SELECT * FROM sections WHERE channel_id=$1 AND name=$2", d.ChannelId, d.Name)
  if len(res) != 0 {
    db.Disconnect()
    return false, nil;
  }
  _, err = db.Query(
    "INSERT INTO sections VALUES ($1, $2, $3)", 
    d.Id, d.ChannelId, d.Name,
  )
  if err != nil {
    return false, err
  }
  return true, nil
}

func Get(id int) (DataModel, error) {
  res, err := db.Query("SELECT * FROM sections WHERE id=$1", id)
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

func GetSectionsOf(channel_id int) ([]DataModel, error) {
  res, err := db.Query("SELECT * FROM sections WHERE channel_id=$1", channel_id)
  list := make([]DataModel, len(res))
  for i, row := range res {
    list[i] = parseRow(row.([]any))
  }
  return list, err
}

func Remove(id int) error {
  _, err := db.Query("DELETE FROM sections WHERE id=$1", id)
  return err
}
