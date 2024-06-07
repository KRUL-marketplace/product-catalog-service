package product

import (
	"context"
	"database/sql"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	categoryModel "github.com/KRUL-marketplace/product-catalog-service/internal/repository/category/model"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/product/model"
	sq "github.com/Masterminds/squirrel"
	"strconv"
	"strings"
)

func (r *repo) GetById(ctx context.Context, ids []string) (*[]model.GetProduct, error) {
	builder := sq.Select("p.id AS product_id",
		"p.name AS product_name",
		"p.slug AS product_slug",
		"p.description AS product_description",
		"p.price AS product_price",
		"p.created_at AS product_created_at",
		"p.updated_at AS product_updated_at",
		"b.id AS brand_id", "b.name AS brand_name", "b.slug AS brand_slug", "b.description AS brand_description",
		"b.created_at AS brand_created_at", "b.updated_at AS brand_updated_at",
		"STRING_AGG(c.id::character varying, ',') AS category_ids",
		"STRING_AGG(c.name, ',') AS category_names",
		"STRING_AGG(c.slug, ',') AS category_slugs").
		From("products p").
		LeftJoin("product_categories pc ON p.id = pc.product_id").
		LeftJoin("categories c ON pc.category_id = c.id").
		LeftJoin("brands b ON p.brand_id = b.id").
		Where(sq.Eq{"p.id": ids}).
		GroupBy("p.id, p.name, p.description, p.price, b.id, b.name, b.slug, b.description")

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "product_repository.GetByIds",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.GetProduct
	for rows.Next() {
		var product model.GetProduct
		var categoriesIds, categoriesNames, categoriesSlugs sql.NullString

		err := rows.Scan(&product.ID, &product.Info.Name, &product.Info.Slug, &product.Info.Description, &product.Info.Price,
			&product.CreatedAt, &product.UpdatedAt,
			&product.Info.Brand.ID, &product.Info.Brand.Info.Name, &product.Info.Brand.Info.Slug,
			&product.Info.Brand.Info.Description, &product.Info.Brand.CreatedAt, &product.Info.Brand.UpdatedAt,
			&categoriesIds, &categoriesNames, &categoriesSlugs)
		if err != nil {
			return nil, err
		}

		if categoriesIds.Valid {
			categoryIdsSlice := strings.Split(categoriesIds.String, ",")
			categoryNamesSlice := strings.Split(categoriesNames.String, ",")
			categorySlugsSlice := strings.Split(categoriesSlugs.String, ",")

			var categoryList []categoryModel.Category
			for i, name := range categoryNamesSlice {
				categoryId, _ := strconv.ParseUint(categoryIdsSlice[i], 10, 32)

				category := categoryModel.Category{
					ID: uint32(categoryId),
					Info: categoryModel.CategoryInfo{
						Name: name,
						Slug: categorySlugsSlice[i],
					},
				}
				categoryList = append(categoryList, category)
			}

			product.Info.Categories = categoryList
		}

		products = append(products, product)
	}

	return &products, nil
}
