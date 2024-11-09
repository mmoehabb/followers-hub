package subscriptions

import (
	"errors"
	"goweb/db"
)

func Add(d *DataModel) error {
  res, err := db.SeqQuery("SELECT * FROM subscriptions WHERE streamer_id=$1 AND follower_email=$2", d.StreamerId, d.FollowerEmail)
  if len(res) != 0 {
    db.Disconnect()
    return errors.New("email already found.")
  }
  _, err = db.Query(
    "INSERT INTO subscriptions VALUES ($1, $2, $3, $4)", 
    d.StreamerId, d.FollowerEmail, d.SubscribedAt, d.Bending,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(email string) (DataModel, error) {
  res, err := db.Query("SELECT * FROM subscriptions WHERE email=$1", email)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, errors.New("couldn't find email.")
  }
  row := res[0].([]any)
  obj := parseRow(row)
  return obj, nil
}

func GetSubsOf(username string) ([]DataModel, error) {
  res, err := db.Query("SELECT * FROM subscriptions WHERE strearmer_id=$1 AND bending=FALSE", username)
  list := make([]DataModel, len(res))
  for i, row := range res {
    list[i] = parseRow(row.([]any))
  }
  return list, err
}

func GetBendingSubs(username string) ([]DataModel, error) {
  res, err := db.Query("SELECT * FROM subscriptions WHERE strearmer_id=$1 AND bending=TRUE", username)
  list := make([]DataModel, len(res))
  for i, row := range res {
    list[i] = parseRow(row.([]any))
  }
  return list, err
}

func Remove(StreamerUsername, FollowerEmail string) error {
  _, err := db.Query(
    "DELETE FROM subscriptions WHERE strearmer_id=$1 AND follower_email=$2", 
    StreamerUsername, FollowerEmail,
  )
  return err
}
