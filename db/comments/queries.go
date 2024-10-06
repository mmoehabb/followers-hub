package comments

import (
	"errors"
	"goweb/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type DataModel struct{
  VideoId int
  FollowerEmail string
  Content string
  CommentedAt pgtype.Timestamp
}

func Add(d *DataModel) error {
  _, err := db.Query(
    "INSERT INTO comments VALUES ($1, $2, $3)", 
    d.VideoId, d.FollowerEmail, d.Content,
  )
  if err != nil {
    return err
  }
  return nil
}

func Get(id int) (DataModel, error) {
  res, err := db.Query("SELECT * FROM comments WHERE id=$1", id)
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
    Content: row[2].(string),
    CommentedAt: row[3].(pgtype.Timestamp),
  }
  return obj, nil
}

func Remove(id int) error {
  _, err := db.Query("DELETE FROM comments WHERE id=$1", id)
  return err
}
