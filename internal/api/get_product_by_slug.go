package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/product"
	desc "product-catalog-service/pkg/product-catalog-service"
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
