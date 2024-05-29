package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/category"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) CreateCategory(ctx context.Context, req *desc.CreateCategoryRequest) (*desc.CreateCategoryResponse, error) {
	id, err := i.categoryService.Create(ctx, converter.ToCategoryInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted product with id: %d", id)

	return &desc.CreateCategoryResponse{
		Id: id,
	}, nil
}
