// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type Dbmate struct {
	ID       sql.NullInt32
	Name     sql.NullString
	Email    string
	Age      sql.NullInt32
	IsActive sql.NullBool
}
