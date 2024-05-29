package api

import (
	brandService "github.com/KRUL-marketplace/product-catalog-service/internal/service/brand"
	categoryService "github.com/KRUL-marketplace/product-catalog-service/internal/service/category"
	productService "github.com/KRUL-marketplace/product-catalog-service/internal/service/product"
	desc "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
)

type Implementation struct {
	desc.UnimplementedProductCatalogServiceServer
	productService  productService.ProductService
	categoryService categoryService.CategoryService
	brandService    brandService.BrandService
}

func NewImplementation(
	productService productService.ProductService,
	categoryService categoryService.CategoryService,
	brandService brandService.BrandService,
) *Implementation {
	return &Implementation{
		productService:  productService,
		categoryService: categoryService,
		brandService:    brandService,
	}
}
