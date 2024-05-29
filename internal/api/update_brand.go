package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/brand"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) UpdateBrand(ctx context.Context, req *desc.UpdateBrandRequest) (*desc.UpdateBrandResponse, error) {
	id := req.GetId()
	err := i.brandService.Update(ctx, id, converter.ToBrandInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated brand with id: %d", id)

	return &desc.UpdateBrandResponse{
		Message: "Success",
	}, nil
}
