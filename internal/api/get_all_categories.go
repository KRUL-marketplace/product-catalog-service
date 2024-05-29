package api

import (
	"context"
	"log"
	converter "product-catalog-service/internal/converter/category"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func (i *Implementation) GetAllCategories(ctx context.Context, _ *desc.GetAllCategoriesRequest) (*desc.GetAllCategoriesResponse, error) {
	categories, err := i.categoryService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var categoriesMessages []*desc.Category

	for _, category := range categories {
		log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at:%v\n",
			category.ID, category.Info.Name, category.Info.Slug, category.CreatedAt, category.UpdatedAt)

		categoryMessage := converter.ToCategoryFromService(category)
		categoriesMessages = append(categoriesMessages, categoryMessage)
	}

	response := &desc.GetAllCategoriesResponse{
		Categories: categoriesMessages,
	}

	return response, nil
}
