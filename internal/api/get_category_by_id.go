package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/category"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) GetCategoryById(ctx context.Context, req *desc.GetCategoryByIdRequest) (*desc.GetCategoryByIdResponse, error) {
	category, err := i.categoryService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at: %v",
		category.ID, category.Info.Name, category.Info.Slug, category.CreatedAt, category.UpdatedAt)

	return &desc.GetCategoryByIdResponse{
		Category: converter.ToCategoryFromService(category),
	}, nil
}
