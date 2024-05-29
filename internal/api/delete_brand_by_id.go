package api

import (
	"context"
	"log"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) DeleteBrandById(ctx context.Context, req *desc.DeleteBrandByIdRequest) (*desc.DeleteBrandByIdResponse, error) {
	err := i.brandService.DeleteById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("brand deleted by id %d", req.GetId())

	return &desc.DeleteBrandByIdResponse{
		Message: "Success",
	}, nil
}
