package likes

import (
	"errors"
	"goweb/db"
)

type DataModel struct{
  VideoId int
  FollowerEmail string
}

func Add(d *DataModel) error {
  _, err := db.Query(
    "INSERT INTO likes VALUES ($1, $2)", 
    d.VideoId, d.FollowerEmail,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(id int) (DataModel, error) {
  res, err := db.Query("SELECT * FROM likes WHERE id=$1", id)
  if err != nil {
    return DataModel{}, err
  }
  if len(res) == 0 {
    return DataModel{}, errors.New("data not found.")
  }
  row := res[0].([]any)
  obj := DataModel{ 
    VideoId: row[0].(int),
    FollowerEmail: row[1].(string),
  }
  return obj, nil
}

func Remove(id int) error {
  _, err := db.Query("DELETE FROM likes WHERE id=$1", id)
  return err
}
