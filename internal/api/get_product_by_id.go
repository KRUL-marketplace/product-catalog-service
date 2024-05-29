package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/product"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) GetProductById(ctx context.Context, req *desc.GetProductByIdRequest) (*desc.GetProductByIdResponse, error) {
	productObj, err := i.productService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, name: %s, slug: %s, description: %s, price: %d, created_at: %v, updated_at:%v\n",
		productObj.ID, productObj.Info.Name, productObj.Info.Slug, productObj.Info.Description, productObj.Info.Description)

	return &desc.GetProductByIdResponse{
		Product: converter.ToProductFromService(productObj),
	}, nil
}
