package product

import (
	"context"
	"database/sql"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/category/model"
	productModel "github.com/KRUL-marketplace/product-catalog-service/internal/repository/product/model"
	sq "github.com/Masterminds/squirrel"
	"strconv"
	"strings"
)

func (r *repo) GetAll(ctx context.Context) ([]*productModel.GetProduct, error) {
	builder := sq.Select("p.id AS product_id", "p.name AS product_name", "p.slug AS product_slug",
		"p.description AS product_description", "p.price AS product_price",
		"p.gender AS product_gender", "p.created_at AS product_created_at",
		"p.updated_at AS product_updated_at",
		"b.id AS brand_id", "b.name AS brand_name", "b.slug AS brand_slug", "b.description AS brand_description",
		"b.created_at AS brand_created_at", "b.updated_at AS brand_updated_at",
		"STRING_AGG(c.id::character varying, ',') AS category_ids", "STRING_AGG(c.name, ',') AS categories",
		"STRING_AGG(c.slug, ',') AS category_slugs").
		From("products p").
		LeftJoin("product_categories pc ON p.id = pc.product_id").
		LeftJoin("categories c ON pc.category_id = c.id").
		LeftJoin("brands b ON p.brand_id = b.id").
		GroupBy("p.id, p.name, p.slug, p.description, p.price, p.gender, p.created_at, p.updated_at, " +
			"b.id, b.name, b.slug, b.description").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "product_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*productModel.GetProduct
	for rows.Next() {
		var categoriesIds sql.NullString
		var categories sql.NullString
		var categorySlugs sql.NullString

		var product productModel.GetProduct

		err := rows.Scan(&product.ID, &product.Info.Name, &product.Info.Slug, &product.Info.Description, &product.Info.Price,
			&product.Info.Gender, &product.CreatedAt, &product.UpdatedAt, &product.Info.Brand.ID,
			&product.Info.Brand.Info.Name, &product.Info.Brand.Info.Slug, &product.Info.Brand.Info.Description,
			&product.Info.Brand.CreatedAt, &product.Info.Brand.UpdatedAt, &categoriesIds, &categories, &categorySlugs)

		if err != nil {
			return nil, err
		}

		if categoriesIds.Valid {
			categoryIdsSlice := strings.Split(categoriesIds.String, ",")
			categoryNamesSlice := strings.Split(categories.String, ",")
			categorySlugsSlice := strings.Split(categorySlugs.String, ",")

			var categoryList []model.Category
			for i, name := range categoryNamesSlice {
				id, _ := strconv.ParseUint(categoryIdsSlice[i], 10, 32)
				category := model.Category{
					ID: uint32(id),
					Info: model.CategoryInfo{
						Name: name,
						Slug: categorySlugsSlice[i],
					},
				}
				categoryList = append(categoryList, category)
			}

			product.Info.Categories = categoryList
		}

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
