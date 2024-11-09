package subscriptions

import (
  "github.com/jackc/pgx/v5/pgtype"
)

type DataModel struct{
  StreamerId string
  FollowerEmail string
  SubscribedAt pgtype.Timestamp
  Bending bool
}

func parseRow(row []any) DataModel {
  return DataModel{ 
    StreamerId: row[0].(string),
    FollowerEmail: row[1].(string),
    SubscribedAt: row[2].(pgtype.Timestamp),
    Bending: row[3].(bool),
  }
}
