package model

import (
	"database/sql"
	brandModel "product-catalog-service/internal/repository/brand/model"
	categoryModel "product-catalog-service/internal/repository/category/model"

	"time"
)

type GetProduct struct {
	ID        string         `db:"id"`
	Info      GetProductInfo `db:""`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

type GetProductInfo struct {
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	Description string `db:"description"`
	Price       uint32 `db:"price"`
	Categories  []categoryModel.Category
	Brand       brandModel.Brand `db:""`
}

type CreateProduct struct {
	Name        string
	Slug        string
	Description string
	Price       uint32
	Categories  []uint32
	BrandId     uint32
}
