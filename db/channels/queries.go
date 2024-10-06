package channels

import (
	"errors"
	"goweb/db"
)

type DataModel struct{
  Id int
  StreamerUsername string
  Name string
}

func Add(d *DataModel) error {
  res, err := db.SeqQuery("SELECT * FROM channels WHERE streamer_username=$1 AND name=$2", d.StreamerUsername, d.Name)
  if len(res) != 0 {
    db.Disconnect()
    return errors.New("channel already found.")
  }
  _, err = db.Query(
    "INSERT INTO channels VALUES ($1, $2, $3)", 
    d.Id, d.StreamerUsername, d.Name,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(id int) (DataModel, error) {
  res, err := db.Query("SELECT * FROM channels WHERE id=$1", id)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, errors.New("data not found.")
  }
  row := res[0].([]any)
  obj := DataModel{ 
    Id: row[0].(int),
    StreamerUsername: row[1].(string),
    Name: row[2].(string),
  }
  return obj, nil
}

func GetChannelsOf(username string) ([]DataModel, error) {
  res, err := db.Query("SELECT * FROM channels WHERE streamer_username=$1", username)
  list := make([]DataModel, len(res))
  for i, row := range res {
    list[i] = row.(DataModel)
  }
  return list, err
}

func Remove(id int) error {
  _, err := db.Query("DELETE FROM channels WHERE id=$1", id)
  return err
}
