package api

import (
	"context"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) DeleteCategoryById(ctx context.Context, req *desc.DeleteCategoryByIdRequest) (*desc.DeleteCategoryByIdResponse, error) {
	err := i.categoryService.DeleteById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("product deleted by id %d", req.GetId())

	return &desc.DeleteCategoryByIdResponse{
		Message: "Success",
	}, nil
}
