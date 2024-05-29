package category

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/category/model"
	"time"
)

func (r *repo) Update(ctx context.Context, id uint32, info *model.CategoryInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, info.Name).
		Set(slugColumn, info.Slug).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "category_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
