package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/brand"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) CreateBrand(ctx context.Context, req *desc.CreateBrandRequest) (*desc.CreateBrandResponse, error) {
	id, err := i.brandService.Create(ctx, converter.ToBrandInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted product with id: %d", id)

	return &desc.CreateBrandResponse{
		Id: id,
	}, nil
}
