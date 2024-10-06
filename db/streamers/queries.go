package streamers

import (
	"errors"
	"goweb/db"
)

type DataModel struct{
  Id string
  Username string
  DisplayName string
  ImgUrl string
  AccessToken string
  RefreshToken string
}

func Add(d *DataModel) error {
  res, err := db.SeqQuery("SELECT * FROM streamers WHERE username=$1", d.Username)
  if len(res) != 0 {
    db.Disconnect()
    return errors.New("user already found.")
  }
  _, err = db.Query(
    "INSERT INTO streamers VALUES ($1, $2, $3, $4, $5, $6)", 
    d.Id, d.Username, d.DisplayName, d.ImgUrl, d.AccessToken, d.RefreshToken,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(username string) (DataModel, error) {
  res, err := db.Query("SELECT * FROM streamers WHERE username=$1", username)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, errors.New("couldn't find username.")
  }
  row := res[0].([]any)
  streamer := DataModel{ 
    Username: row[0].(string), 
    Email: row[1].(string),
    DisplayName: row[2].(string),
    ImgUrl: row[3].(string),
    Token: row[4].(string),
  }
  return streamer, nil
}

func Remove(username string) error {
  _, err := db.Query("DELETE FROM streamers WHERE username=$1", username)
  return err
}
