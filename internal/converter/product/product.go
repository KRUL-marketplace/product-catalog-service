package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"product-catalog-service/internal/repository/product/model"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func ToProductInfoFromDesc(info *desc.CreateProductInfo) *model.CreateProduct {
	product := &model.CreateProduct{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
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

func ToProductInfoFromService(info model.GetProductInfo) *desc.ProductInfo {
	product := &desc.ProductInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
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
		Categories:  info.Categories,
		Brand:       info.Brand,
	}
}
