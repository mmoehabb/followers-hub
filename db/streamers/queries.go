package streamers

import (
	"errors"
	"goweb/ancillaries"
	"goweb/db"
)

func Add(d *DataModel) error {
  res, err := db.SeqQuery("SELECT * FROM streamers WHERE id=$1", d.Id)
  if len(res) != 0 {
    db.Disconnect()
    return errors.New("user already found.")
  }
  _, err = db.Query(
    "INSERT INTO streamers VALUES ($1, $2, $3, $4, $5)", 
    d.Id, d.DisplayName, d.ImgUrl, d.AccessToken, d.RefreshToken,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(id string) (DataModel, error) {
  res, err := db.Query("SELECT * FROM streamers WHERE id=$1", id)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, nil
  }
  row := res[0].([]any)
  streamer := parseRow(row)
  return streamer, nil
}

func Exists(id string) (bool, error) {
  res, err := db.Query("SELECT * FROM streamers WHERE id=$1", id)
  if err != nil {
    return false, err
  }
  return len(res) > 0, nil
}

func Update(data *DataModel) error {
  if data.Id == "" {
    return errors.New("Id must be specified in order to update data.")
  }
  res, err := db.SeqQuery("SELECT * FROM streamers WHERE id=$1", data.Id)
  if err != nil {
    db.Disconnect()
    return err
  }
  if len(res) == 0 {
    db.Disconnect()
    return errors.New("entity not found.")
  }

  q, values := ancillaries.GenUpdateQuery("streamers", parseModel(data), "id")
  _, err = db.Query(q, values...)
  if err != nil {
    return err
  }

  return nil
}

func Remove(id string) error {
  _, err := db.Query("DELETE FROM streamers WHERE id=$1", id)
  return err
}
