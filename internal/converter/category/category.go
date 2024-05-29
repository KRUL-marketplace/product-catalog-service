package category

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"product-catalog-service/internal/repository/category/model"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func ToCategoryInfoFromDesc(info *desc.CategoryInfo) *model.CategoryInfo {
	return &model.CategoryInfo{
		Name: info.Name,
		Slug: info.Slug,
	}
}

func ToCategoryFromService(category *model.Category) *desc.Category {
	var updatedAt *timestamppb.Timestamp
	if category.UpdatedAt.Valid {
		updatedAt = timestamppb.New(category.UpdatedAt.Time)
	}

	return &desc.Category{
		Id:        category.ID,
		Info:      ToCategoryInfoFromService(category.Info),
		CreatedAt: timestamppb.New(category.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToCategoryInfoFromService(info model.CategoryInfo) *desc.CategoryInfo {
	return &desc.CategoryInfo{
		Name: info.Name,
		Slug: info.Slug,
	}
}

func ToCategoryFromRepo(category *model.Category) *model.Category {
	return &model.Category{
		ID:        category.ID,
		Info:      ToCategoryInfoFromRepo(category.Info),
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToCategoryInfoFromRepo(info model.CategoryInfo) model.CategoryInfo {
	return model.CategoryInfo{
		Name: info.Name,
		Slug: info.Slug,
	}
}
