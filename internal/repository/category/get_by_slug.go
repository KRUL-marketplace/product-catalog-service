package category

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
	converter "product-catalog-service/internal/converter/category"
	"product-catalog-service/internal/repository/category/model"
)

func (r *repo) GetBySlug(ctx context.Context, slug string) (*model.Category, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{slugColumn: slug}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "category_repository.GetBySlug",
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
