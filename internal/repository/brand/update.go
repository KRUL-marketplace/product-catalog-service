package brand

import (
	"context"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/brand/model"
	sq "github.com/Masterminds/squirrel"
	"time"
)

func (r *repo) Update(ctx context.Context, id uint32, info *model.BrandInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, info.Name).
		Set(slugColumn, info.Slug).
		Set(descriptionColumn, info.Description).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "brand_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
