package port

import domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"

// BestSellers defines the methods for interacting with the BestSellers repository
type BestSellers interface {
	FindBestSellers(projectID string) ([]*domain.ProductBestSellerRecord, error)
	CreateBestSellers(projectID string, data []*domain.ProductBestSellerRecord) error
}
