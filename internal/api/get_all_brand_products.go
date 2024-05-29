package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/product"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) GetAllBrandProducts(ctx context.Context, req *desc.GetAllBrandProductsRequest) (*desc.GetAllBrandProductsResponse, error) {
	products, err := i.productService.GetByBrand(ctx, req.GetId())
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

	response := &desc.GetAllBrandProductsResponse{
		Products: productMessages,
	}

	return response, nil
}
