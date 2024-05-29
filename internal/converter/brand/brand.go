package brand

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"product-catalog-service/internal/repository/brand/model"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func ToBrandInfoFromDesc(info *desc.BrandInfo) *model.BrandInfo {
	return &model.BrandInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
	}
}

func ToBrandFromService(brand *model.Brand) *desc.Brand {
	var updatedAt *timestamppb.Timestamp
	if brand.UpdatedAt.Valid {
		updatedAt = timestamppb.New(brand.UpdatedAt.Time)
	}

	return &desc.Brand{
		Id:        brand.ID,
		Info:      ToBrandInfoFromService(brand.Info),
		CreatedAt: timestamppb.New(brand.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToBrandInfoFromService(info model.BrandInfo) *desc.BrandInfo {
	return &desc.BrandInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
	}
}

func ToBrandFromRepo(brand *model.Brand) *model.Brand {
	return &model.Brand{
		ID:        brand.ID,
		Info:      ToBrandInfoFromRepo(brand.Info),
		CreatedAt: brand.CreatedAt,
		UpdatedAt: brand.UpdatedAt,
	}
}

func ToBrandInfoFromRepo(info model.BrandInfo) model.BrandInfo {
	return model.BrandInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
	}
}
