package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/category"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) GetCategoryBySlug(ctx context.Context, req *desc.GetCategoryBySlugRequest) (*desc.GetCategoryBySlugResponse, error) {
	category, err := i.categoryService.GetBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at: %v",
		category.ID, category.Info.Name, category.Info.Slug, category.CreatedAt, category.UpdatedAt)

	return &desc.GetCategoryBySlugResponse{
		Category: converter.ToCategoryFromService(category),
	}, nil
}
