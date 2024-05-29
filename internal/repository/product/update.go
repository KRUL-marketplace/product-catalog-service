package product

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/product/model"
	"time"
)

func (r *repo) Update(ctx context.Context, id string, info *model.CreateProduct) (string, error) {
	productUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	// Обновление информации о продукте
	builder := sq.Update(tableName).
		Where(sq.Eq{idColumn: productUUID}).
		Set(nameColumn, info.Name).
		Set(descriptionColumn, info.Description).
		Set(slugColumn, info.Slug).
		Set(priceColumn, info.Price).
		Set(brandIdColumn, info.BrandId).
		Set(updatedAtColumn, time.Now())

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "category_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return "", err
	}

	// Удаление существующих записей о категориях
	deleteBuilder := sq.Delete("product_categories").
		Where(sq.Eq{"product_id": productUUID})

	deleteQuery, deleteArgs, err := deleteBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return "", err
	}

	q = db.Query{
		Name:     "category_repository.DeleteCategories",
		QueryRaw: deleteQuery,
	}

	_, err = r.db.DB().ExecContext(ctx, q, deleteArgs...)
	if err != nil {
		return "", err
	}

	// Добавление новых записей о категориях
	if len(info.Categories) > 0 {
		insertBuilder := sq.Insert("product_categories").
			Columns("product_id", "category_id").
			Values(productUUID, sq.Expr("unnest(?::int[])", info.Categories))

		insertQuery, insertArgs, err := insertBuilder.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			return "", err
		}

		q = db.Query{
			Name:     "category_repository.UpdateCategories",
			QueryRaw: insertQuery,
		}

		_, err = r.db.DB().ExecContext(ctx, q, insertArgs...)
		if err != nil {
			return "", err
		}
	}

	return "Success", nil
}
