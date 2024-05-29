package api

import (
	"context"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
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
