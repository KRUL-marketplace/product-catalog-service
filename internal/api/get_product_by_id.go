package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/product"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) GetProductById(ctx context.Context, req *desc.GetProductByIdRequest) (*desc.GetProductByIdResponse, error) {
	productObj, err := i.productService.GetById(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}

	return &desc.GetProductByIdResponse{
		Product: converter.ToProductArrayFromService(productObj),
	}, nil
}
