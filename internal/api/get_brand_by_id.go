package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/brand"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) GetBrandById(ctx context.Context, req *desc.GetBrandByIdRequest) (*desc.GetBrandByIdResponse, error) {
	brand, err := i.brandService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at: %v",
		brand.ID, brand.Info.Name, brand.Info.Slug, brand.CreatedAt, brand.UpdatedAt)

	return &desc.GetBrandByIdResponse{
		Brand: converter.ToBrandFromService(brand),
	}, nil
}
