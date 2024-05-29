package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/category"
	desc "product-catalog-service/pkg/product-catalog-service"
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
