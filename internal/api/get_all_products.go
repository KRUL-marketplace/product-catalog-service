package api

import (
	"context"
	converter "github.com/KRUL-marketplace/product-catalog-service/internal/converter/product"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"log"
)

func (i *Implementation) GetAllProducts(ctx context.Context, filters *desc.GetAllProductsRequest) (*desc.GetAllProductsResponse, error) {
	products, err := i.productService.GetAll(ctx, converter.ToProductFiltersFromDesc(filters))
	if err != nil {
		return nil, err
	}

	var productMessages []*desc.Product

	for _, product := range products {
		log.Printf("id: %s, name: %s, slug: %s, description: %s, price: %d, created_at: %v, updated_at:%v\n",
			product.ID, product.Info.Name, product.Info.Slug, product.Info.Description, product.Info.Description)

		productMessage := converter.ToProductFromService(product)
		productMessages = append(productMessages, productMessage)
	}

	response := &desc.GetAllProductsResponse{
		Product: productMessages,
	}

	return response, nil
}
