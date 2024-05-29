package app

import (
	"context"
	"log"
	"product-catalog-service/client/db"
	"product-catalog-service/client/db/pg"
	"product-catalog-service/client/db/transaction"
	"product-catalog-service/internal/api"
	"product-catalog-service/internal/config"
	brandRepository "product-catalog-service/internal/repository/brand"
	categoryRepository "product-catalog-service/internal/repository/category"
	productRepository "product-catalog-service/internal/repository/product"
	brandService "product-catalog-service/internal/service/brand"
	categoryService "product-catalog-service/internal/service/category"
	productService "product-catalog-service/internal/service/product"
)

type serviceProvider struct {
	productRepository productRepository.Repository
	productService    productService.ProductService

	categoryRepository categoryRepository.Repository
	categoryService    categoryService.CategoryService

	brandRepository brandRepository.Repository
	brandService    brandService.BrandService

	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	pgConfig      config.PGConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager

	productCatalogImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) ProductRepository(ctx context.Context) productRepository.Repository {
	if s.productRepository == nil {
		s.productRepository = productRepository.NewRepository(s.DBClient(ctx))
	}

	return s.productRepository
}

func (s *serviceProvider) ProductService(ctx context.Context) productService.ProductService {
	if s.productService == nil {
		s.productService = productService.NewService(
			s.ProductRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.productService
}

func (s *serviceProvider) CategoryRepository(ctx context.Context) categoryRepository.Repository {
	if s.categoryRepository == nil {
		s.categoryRepository = categoryRepository.NewRepository(s.DBClient(ctx))
	}

	return s.categoryRepository
}

func (s *serviceProvider) BrandRepository(ctx context.Context) brandRepository.Repository {
	if s.brandRepository == nil {
		s.brandRepository = brandRepository.NewRepository(s.DBClient(ctx))
	}

	return s.brandRepository
}

func (s *serviceProvider) CategoryService(ctx context.Context) categoryService.CategoryService {
	if s.categoryService == nil {
		s.categoryService = categoryService.NewService(
			s.CategoryRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.categoryService
}

func (s *serviceProvider) BrandService(ctx context.Context) brandService.BrandService {
	if s.brandService == nil {
		s.brandService = brandService.NewService(
			s.BrandRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.brandService
}

func (s *serviceProvider) ProductCatalogImpl(ctx context.Context) *api.Implementation {
	if s.productCatalogImpl == nil {
		s.productCatalogImpl = api.NewImplementation(s.ProductService(ctx), s.CategoryService(ctx), s.BrandService(ctx))
	}

	return s.productCatalogImpl
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}
