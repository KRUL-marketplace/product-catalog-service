package api

import (
	brandService "product-catalog-service/internal/service/brand"
	categoryService "product-catalog-service/internal/service/category"
	productService "product-catalog-service/internal/service/product"
	desc "product-catalog-service/pkg/product-catalog-service"
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
