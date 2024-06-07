package category

import (
	"context"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/category"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/category/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) GetAll(ctx context.Context) ([]*model.Category, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "category_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Info.Name, &category.Info.Slug, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			return nil, err
		}

		categories = append(categories, converter.ToCategoryFromRepo(&category))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
