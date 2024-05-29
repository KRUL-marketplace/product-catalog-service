package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/brand"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) GetAllBrands(ctx context.Context, _ *desc.GetAllBrandsRequest) (*desc.GetAllBrandsResponse, error) {
	brands, err := i.brandService.GetAll(ctx)

	var brandsMessages []*desc.Brand

	for _, brand := range brands {
		log.Printf("id: %d, name: %s, slug: %s, description: %s, created_at: %v, updated_at:%v\n",
			brand.ID, brand.Info.Name, brand.Info.Slug, brand.Info.Description, brand.CreatedAt, brand.UpdatedAt)

		brandMessage := converter.ToBrandFromService(brand)
		brandsMessages = append(brandsMessages, brandMessage)
	}

	response := &desc.GetAllBrandsResponse{
		Brands: brandsMessages,
	}

	return response, err
}
