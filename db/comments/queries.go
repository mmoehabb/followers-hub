package comments

import (
	"errors"
	"goweb/db"

	"github.com/jackc/pgx/v5/pgtype"
)

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
	obj := parseRow(row)
	return obj, nil
}

func GetVideosOf(video_id int) ([]DataModel, error) {
	res, err := db.Query("SELECT * FROM comments WHERE video_id=$1", video_id)
	list := make([]DataModel, len(res))
	for i, row := range res {
		list[i] = parseRow(row.([]any))
	}
	return list, err
}

func Remove(id int) error {
	_, err := db.Query("DELETE FROM comments WHERE id=$1", id)
	return err
}
