package brand

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/brand/model"
)

func (r *repo) Create(ctx context.Context, info *model.BrandInfo) (uint32, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, slugColumn, descriptionColumn).
		Values(info.Name, info.Slug, info.Description).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "brand_repository.Create",
		QueryRaw: query,
	}

	var id uint32
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
