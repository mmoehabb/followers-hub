package comments

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type DataModel struct {
	VideoId       int
	FollowerEmail string
	Content       string
	CommentedAt   pgtype.Timestamp
}

func parseRow(row []any) DataModel {
	return DataModel{
		VideoId:       int(row[0].(int32)),
		FollowerEmail: row[1].(string),
		Content:       row[2].(string),
		CommentedAt:   row[3].(pgtype.Timestamp),
	}
}
