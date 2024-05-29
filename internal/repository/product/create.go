package product

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/product/model"
)

func (r *repo) Create(ctx context.Context, product *model.CreateProduct) (string, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, slugColumn, descriptionColumn, priceColumn, brandIdColumn).
		Values(product.Name, product.Slug, product.Description, product.Price, product.BrandId).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "product_repository.Created",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	for _, categoryID := range product.Categories {
		err := r.createProductCategory(ctx, id, categoryID)
		if err != nil {
			return "Incorrect category", err
		}
	}

	return id, nil
}

func (r *repo) createProductCategory(ctx context.Context, productID string, categoryID uint32) error {
	builder := sq.Insert("product_categories").
		PlaceholderFormat(sq.Dollar).
		Columns("product_id", "category_id").
		Values(productID, categoryID)

	log.Printf("PRDOUCT ID %s CATEGORY ID %d", productID, categoryID)

	// Build the SQL query
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "product_repository.CategoriesAssociation",
		QueryRaw: query,
	}

	// Execute the SQL query
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
