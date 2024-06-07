package category

import (
	"context"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/category"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/category/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) GetById(ctx context.Context, id uint32) (*model.Category, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "category_repository.GetById",
		QueryRaw: query,
	}

	var category model.Category
	err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&category.ID, &category.Info.Name, &category.Info.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToCategoryFromRepo(&category), nil
}
