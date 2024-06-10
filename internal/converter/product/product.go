package converter

import (
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/product/model"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProductInfoFromDesc(info *desc.CreateProductInfo) *model.CreateProduct {
	product := &model.CreateProduct{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
		Gender:      info.Gender,
		Categories:  info.Categories,
		BrandId:     info.BrandId,
	}

	return product
}

func ToProductFromService(product *model.GetProduct) *desc.Product {
	var updatedAt *timestamppb.Timestamp
	if product.UpdatedAt.Valid {
		updatedAt = timestamppb.New(product.UpdatedAt.Time)
	}

	return &desc.Product{
		Id:        product.ID,
		Info:      ToProductInfoFromService(product.Info),
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToProductArrayFromService(products *[]model.GetProduct) []*desc.Product {
	var resProducts []*desc.Product

	for _, product := range *products {
		var updatedAt *timestamppb.Timestamp
		if product.UpdatedAt.Valid {
			updatedAt = timestamppb.New(product.UpdatedAt.Time)
		}

		resProducts = append(resProducts, &desc.Product{
			Id:        product.ID,
			Info:      ToProductInfoFromService(product.Info),
			CreatedAt: timestamppb.New(product.CreatedAt),
			UpdatedAt: updatedAt,
		})
	}

	return resProducts
}

func ToProductInfoFromService(info model.GetProductInfo) *desc.ProductInfo {
	product := &desc.ProductInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
		Gender:      info.Gender,
		Brand: &desc.Brand{
			Id: info.Brand.ID,
			Info: &desc.BrandInfo{
				Name:        info.Brand.Info.Name,
				Slug:        info.Brand.Info.Slug,
				Description: info.Brand.Info.Description,
			},
		},
	}

	for _, category := range info.Categories {
		product.Categories = append(product.Categories, &desc.Category{
			Id: category.ID,
			Info: &desc.CategoryInfo{
				Name: category.Info.Name,
				Slug: category.Info.Slug,
			},
		})
	}

	return product
}

func ToProductFromRepo(product *model.GetProduct) *model.GetProduct {
	return &model.GetProduct{
		ID:        product.ID,
		Info:      ToProductInfoFromRepo(product.Info),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToProductInfoFromRepo(info model.GetProductInfo) model.GetProductInfo {
	return model.GetProductInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
		Gender:      info.Gender,
		Categories:  info.Categories,
		Brand:       info.Brand,
	}
}

func ToProductFiltersFromDesc(filters *desc.GetAllProductsRequest) *model.FilterProduct {
	return &model.FilterProduct{
		BrandIds:    filters.BrandIds,
		CategoryIds: filters.CategoryIds,
		MinPrice:    filters.MinPrice,
		MaxPrice:    filters.MaxPrice,
		Gender:      filters.Gender,
	}
}
