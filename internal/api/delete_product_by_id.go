package api

import (
	"context"
	"log"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) DeleteProductById(ctx context.Context, req *desc.DeleteProductByIdRequest) (*desc.DeleteProductByIdResponse, error) {
	err := i.productService.DeleteById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("product deleted by id")

	return &desc.DeleteProductByIdResponse{
		Message: "Success",
	}, nil
}
