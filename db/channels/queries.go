package channels

import (
	"goweb/db"
)

func Add(d *DataModel) (bool, error) {
  res, err := db.SeqQuery("SELECT * FROM channels WHERE streamer_id=$1 AND name=$2", d.StreamerId, d.Name)
  if len(res) != 0 {
    db.Disconnect()
    return false, nil
  }
  _, err = db.Query(
    `INSERT INTO channels (
      streamer_id, 
      name, 
      primary_color, 
      secondary_color, 
      accent_color, 
      text_color
    ) VALUES ($1, $2, $3, $4, $5, $6);`,
    d.StreamerId, d.Name, d.PrimaryColor, d.SecondaryColor, d.AccentColor, d.TextColor,
  )
  if err != nil {
    return false, err
  }
  return true, nil
}

func Get(id int) (DataModel, error) {
  res, err := db.Query("SELECT * FROM channels WHERE id=$1", id)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, nil
  }
  row := res[0].([]any)
  return parseRow(row), nil
}

func GetChannelsOf(username string) ([]DataModel, error) {
  res, err := db.Query("SELECT * FROM channels WHERE streamer_id=$1", username)
  list := make([]DataModel, len(res))
  for i, row := range res {
    list[i] = parseRow(row.([]any))
  }
  return list, err
}

func Remove(id int) error {
  _, err := db.Query("DELETE FROM channels WHERE id=$1", id)
  return err
}

