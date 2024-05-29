package api

import (
	"context"
	"log"
	desc "product-catalog-service/pkg/product-catalog-service"
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
