package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/category"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) UpdateCategory(ctx context.Context, req *desc.UpdateCategoryRequest) (*desc.UpdateCategoryResponse, error) {
	id := req.GetId()
	err := i.categoryService.Update(ctx, id, converter.ToCategoryInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated product with id: %d", id)

	return &desc.UpdateCategoryResponse{
		Message: "Success",
	}, nil
}
