package streamers

import (
	"errors"
	"goweb/db"
)

type StreamerData struct{
  Username string
  Email string
  DisplayName string
  ImgUrl string
  Token string
}

func Add(d *StreamerData) error {
  res, err := db.SeqQuery("SELECT * FROM streamers WHERE username=$1", d.Username)
  if len(res) != 0 {
    db.Disconnect()
    return errors.New("username already found.")
  }
  _, err = db.Query(
    "INSERT INTO streamers VALUES ($1, $2, $3, $4, $5)", 
    d.Username, d.Email, d.DisplayName, d.ImgUrl, d.Token,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(username string) (StreamerData, error) {
  res, err := db.Query("SELECT * FROM users WHERE username=$1", username)
  if err != nil {
    return StreamerData{}, err
  }
  if len(res) == 0 {
    return StreamerData{}, errors.New("couldn't find username.")
  }
  streamer := StreamerData{ 
    Username: res[0].(string), 
    Email: res[1].(string),
    DisplayName: res[2].(string),
    ImgUrl: res[3].(string),
    Token: res[4].(string),
  }
  return streamer, nil
}
