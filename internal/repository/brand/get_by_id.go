package brand

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
	converter "product-catalog-service/internal/converter/brand"
	"product-catalog-service/internal/repository/brand/model"
)

func (r *repo) GetById(ctx context.Context, id uint32) (*model.Brand, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, descriptionColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "brand_repository.GetById",
		QueryRaw: query,
	}

	var brand model.Brand
	err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&brand.ID, &brand.Info.Name, &brand.Info.Slug, &brand.Info.Description, &brand.CreatedAt, &brand.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToBrandFromRepo(&brand), nil
}
