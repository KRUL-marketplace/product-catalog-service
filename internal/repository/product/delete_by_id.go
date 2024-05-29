package product

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
)

func (r *repo) DeleteById(ctx context.Context, id string) error {
	builder := sq.Delete("product_categories").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"product_id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "product_repository.DeleteById",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	builder = sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err = builder.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		Name:     "product_repository.DeleteById",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
