package streamers

import (
	"errors"
	"goweb/db"
)

type DataModel struct{
  Id string
  DisplayName string
  ImgUrl string
  AccessToken string
  RefreshToken string
}

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
  streamer := DataModel{
    Id: row[0].(string),
    DisplayName: row[1].(string),
    ImgUrl: row[2].(string),
    AccessToken: row[3].(string),
    RefreshToken: row[4].(string),
  }
  return streamer, nil
}

func Exists(id string) (bool, error) {
  res, err := db.Query("SELECT * FROM streamers WHERE id=$1", id)
  if err != nil {
    return false, err
  }
  return len(res) > 0, nil
}

func UpdateTokens(id string, at string, rt string) error {
  res, err := db.SeqQuery("SELECT * FROM streamers WHERE id=$1", id)
  if len(res) == 0 {
    db.Disconnect()
    return errors.New("entity not found.")
  }
  _, err = db.Query("UPDATE streamers SET access_token=$1, refresh_token=$2 WHERE id=$3", at, rt, id)
  if err != nil {
    return err
  }
  return nil
}

func Remove(id string) error {
  _, err := db.Query("DELETE FROM streamers WHERE id=$1", id)
  return err
}
