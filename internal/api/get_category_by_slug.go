package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/category"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
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
