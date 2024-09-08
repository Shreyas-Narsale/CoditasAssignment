package models

import "database/sql"

type IsUserNameExists struct {
	Count  int            `db:"count"`
	UserID sql.NullString `db:"userid"`
}
