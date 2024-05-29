package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/product"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) GetProductBySlug(ctx context.Context, req *desc.GetProductBySlugRequest) (*desc.GetProductBySlugResponse, error) {
	productObj, err := i.productService.GetBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, name: %s, slug: %s, description: %s, price: %d, created_at: %v, updated_at:%v\n",
		productObj.ID, productObj.Info.Name, productObj.Info.Slug, productObj.Info.Description, productObj.Info.Description)

	return &desc.GetProductBySlugResponse{
		Product: converter.ToProductFromService(productObj),
	}, nil
}
