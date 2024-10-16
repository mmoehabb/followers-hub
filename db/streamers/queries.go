package streamers

import (
	"errors"
	"fmt"
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

  var i int = 1
  var values []any
  q := "UPDATE streamers SET "
  if data.DisplayName != "" {
    q = fmt.Sprintf("%s%s=$%d ", q, "display_name", i)
    values = append(values, data.DisplayName)
    i = i + 1
  }
  if data.ImgUrl != "" {
    q = fmt.Sprintf("%s%s=$%d ", q, "img_url", i)
    values = append(values, data.ImgUrl)
    i = i + 1
  }
  if data.AccessToken != "" {
    q = fmt.Sprintf("%s%s=$%d ", q, "access_token", i)
    values = append(values, data.AccessToken)
    i = i + 1
  }
  if data.RefreshToken != "" {
    q = fmt.Sprintf("%s%s=$%d ", q, "refresh_token", i)
    values = append(values, data.RefreshToken)
    i = i + 1
  }
  q = fmt.Sprintf("%sWHERE id='%s'", q, data.Id)

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
