package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/product"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) UpdateProduct(ctx context.Context, req *desc.UpdateProductRequest) (*desc.CreateProductResponse, error) {
	id := req.GetId()
	_, err := i.productService.Update(ctx, id, converter.ToProductInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated product with id: %s", id)

	return &desc.CreateProductResponse{
		Id: id,
	}, nil
}
