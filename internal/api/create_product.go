package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/product"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	id, err := i.productService.Create(ctx, converter.ToProductInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted product with id: %s", id)

	return &desc.CreateProductResponse{
		Id: id,
	}, nil
}
