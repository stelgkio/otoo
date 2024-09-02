package port

import domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"

// BestSellers defines the methods for interacting with the BestSellers repository
type BestSellers interface {
	FindBestSellers(projectID string, size, page int) ([]*domain.ProductBestSellerRecord, error)
	CreateBestSellers(projectID string, data []*domain.ProductBestSellerRecord) error
	DeleteBestSellers(projectID string) error
}

// ProductBestSellers defines the methods for interacting with the ProductBestSellers repository
type ProductBestSellers interface {
	RunAProductBestSellerInitializerJob(projectID string) error
}
