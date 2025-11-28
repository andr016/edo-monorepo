package models

import "database/sql"

type User struct {
	ID         int64         `db:"id"`
	Login      string        `db:"login"`
	PositionID sql.NullInt64 `db:"position_id"`
}
