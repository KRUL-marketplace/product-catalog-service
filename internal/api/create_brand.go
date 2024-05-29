package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/brand"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
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
