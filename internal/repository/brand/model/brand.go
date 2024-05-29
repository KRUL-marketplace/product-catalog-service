package model

import (
	"database/sql"
	"time"
)

type Brand struct {
	ID        uint32       `db:"id"`
	Info      BrandInfo    `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type BrandInfo struct {
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	Description string `db:"description"`
}
