package model

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        uint32       `db:"id"`
	Info      CategoryInfo `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type CategoryInfo struct {
	Name string `db:"name"`
	Slug string `db:"slug"`
}
