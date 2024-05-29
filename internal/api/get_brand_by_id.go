package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/brand"
	desc "product-catalog-service/pkg/product-catalog-service"
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
