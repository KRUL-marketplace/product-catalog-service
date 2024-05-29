package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/brand"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) GetBrandBySlug(ctx context.Context, req *desc.GetBrandBySlugRequest) (*desc.GetBrandBySlugResponse, error) {
	brand, err := i.brandService.GetBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, slug: %s, description: %s, created_at: %v, updated_at: %v",
		brand.ID, brand.Info.Name, brand.Info.Slug, brand.Info.Description, brand.CreatedAt, brand.UpdatedAt)

	return &desc.GetBrandBySlugResponse{
		Brand: converter.ToBrandFromService(brand),
	}, nil
}
