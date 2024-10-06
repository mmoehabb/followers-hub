package followers

import (
	"errors"
	"goweb/db"
)

type DataModel struct{
  Email string
  DisplayName string
  Token string
}

func Add(d *DataModel) error {
  res, err := db.SeqQuery("SELECT * FROM followers WHERE email=$1", d.Email)
  if len(res) != 0 {
    db.Disconnect()
    return errors.New("email already found.")
  }
  _, err = db.Query(
    "INSERT INTO followers VALUES ($1, $2, $3)", 
    d.Email, d.DisplayName, d.Token,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(email string) (DataModel, error) {
  res, err := db.Query("SELECT * FROM followers WHERE email=$1", email)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, errors.New("couldn't find email.")
  }
  row := res[0].([]any)
  obj := DataModel{ 
    Email: row[0].(string),
    DisplayName: row[1].(string),
    Token: row[2].(string),
  }
  return obj, nil
}

func Remove(email string) error {
  _, err := db.Query("DELETE FROM followers WHERE email=$1", email)
  return err
}
