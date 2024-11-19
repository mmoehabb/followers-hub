package subscriptions

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type DataModel struct {
	StreamerId    string
	FollowerEmail string
	SubscribedAt  pgtype.Timestamp
	Bending       bool
}

func parseRow(row []any) DataModel {
	return DataModel{
		StreamerId:    row[0].(string),
		FollowerEmail: row[1].(string),
		SubscribedAt:  row[2].(pgtype.Timestamp),
		Bending:       row[3].(bool),
	}
}

func parseModel(m *DataModel) map[string]string {
	var modelmap = make(map[string]string)
	if m.StreamerId != "" {
		modelmap["streamer_id"] = m.StreamerId
	}
	if m.FollowerEmail != "" {
		modelmap["follower_email"] = m.FollowerEmail
	}
	if m.Bending != false {
		modelmap["bending"] = strconv.FormatBool(m.Bending)
	}
	return modelmap
}
