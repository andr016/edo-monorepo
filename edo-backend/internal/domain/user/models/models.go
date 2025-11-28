package models

import "database/sql"

type User struct {
	ID          int64
	Login       string
	PositionID  sql.NullInt64
	Permissions []Permission
}

type Permission struct {
	ID          int64
	Name        string
	Code        string
	Description string
}
