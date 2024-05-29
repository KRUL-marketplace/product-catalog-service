package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/category"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
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
