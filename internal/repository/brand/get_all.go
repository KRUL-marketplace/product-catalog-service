package brand

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
	converter "product-catalog-service/internal/converter/brand"
	"product-catalog-service/internal/repository/brand/model"
)

func (r *repo) GetAll(ctx context.Context) ([]*model.Brand, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, descriptionColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "brand_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []*model.Brand
	for rows.Next() {
		var brand model.Brand
		err = rows.Scan(&brand.ID, &brand.Info.Name, &brand.Info.Slug, &brand.Info.Description, &brand.CreatedAt, &brand.UpdatedAt)

		if err != nil {
			return nil, err
		}

		brands = append(brands, converter.ToBrandFromRepo(&brand))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}
