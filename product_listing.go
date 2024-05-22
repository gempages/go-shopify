package goshopify

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const productListingBasePath = "product_listings"
const productsListingResourceName = "product_listings"

// ProductListingService is an interface for interfacing with the product listing endpoints
// of the Shopify API.
// See: https://shopify.dev/docs/admin-api/rest/reference/sales-channels/productlisting
type ProductListingService interface {
	List(context.Context, interface{}) ([]ProductListing, error)
	ListWithPagination(context.Context, interface{}) ([]ProductListing, *Pagination, error)
	Count(context.Context, interface{}) (int, error)
	Get(context.Context, int64, interface{}) (*ProductListing, error)
	GetProductIDs(context.Context, interface{}) ([]int64, error)
	Publish(context.Context, int64) (*ProductListing, error)
	Delete(context.Context, int64) error
}

// ProductListingServiceOp handles communication with the product related methods of
// the Shopify API.
type ProductListingServiceOp struct {
	client *Client
}

// ProductListing represents a Shopify product published to your sales channel app
type ProductListing struct {
	ID          int64           `json:"product_id,omitempty"`
	Title       string          `json:"title,omitempty"`
	BodyHTML    string          `json:"body_html,omitempty"`
	Vendor      string          `json:"vendor,omitempty"`
	ProductType string          `json:"product_type,omitempty"`
	Handle      string          `json:"handle,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	PublishedAt *time.Time      `json:"published_at,omitempty"`
	Tags        string          `json:"tags,omitempty"`
	Options     []ProductOption `json:"options,omitempty"`
	Variants    []Variant       `json:"variants,omitempty"`
	Images      []Image         `json:"images,omitempty"`
}

// Represents the result from the product_listings/X.json endpoint
type ProductListingResource struct {
	ProductListing *ProductListing `json:"product_listing"`
}

// Represents the result from the product_listings.json endpoint
type ProductsListingsResource struct {
	ProductListings []ProductListing `json:"product_listings"`
}

// Represents the result from the product_listings/product_ids.json endpoint
type ProductListingIDsResource struct {
	ProductIDs []int64 `json:"product_ids"`
}

// Resource which create product_listing endpoint expects in request body
// e.g.
// PUT /admin/api/2020-07/product_listings/921728736.json
//
//	{
//	  "product_listing": {
//	    "product_id": 921728736
//	  }
//	}
type ProductListingPublishResource struct {
	ProductListing struct {
		ProductID int64 `json:"product_id"`
	} `json:"product_listing"`
}

// List products
func (s *ProductListingServiceOp) List(ctx context.Context, options interface{}) ([]ProductListing, error) {
	products, _, err := s.ListWithPagination(ctx, options)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ListWithPagination lists products and return pagination to retrieve next/previous results.
func (s *ProductListingServiceOp) ListWithPagination(ctx context.Context, options interface{}) ([]ProductListing, *Pagination, error) {
	path := fmt.Sprintf("%s.json", productListingBasePath)
	resource := new(ProductsListingsResource)
	headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders(ctx, "GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")

	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.ProductListings, pagination, nil
}

// Count products listings published to your sales channel app
func (s *ProductListingServiceOp) Count(ctx context.Context, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", productListingBasePath)
	return s.client.Count(ctx, path, options)
}

// Get individual product_listing by product ID
func (s *ProductListingServiceOp) Get(ctx context.Context, productID int64, options interface{}) (*ProductListing, error) {
	path := fmt.Sprintf("%s/%d.json", productListingBasePath, productID)
	resource := new(ProductListingResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.ProductListing, err
}

// GetProductIDs lists all product IDs that are published to your sales channel
func (s *ProductListingServiceOp) GetProductIDs(ctx context.Context, options interface{}) ([]int64, error) {
	path := fmt.Sprintf("%s/product_ids.json", productListingBasePath)
	resource := new(ProductListingIDsResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.ProductIDs, err
}

// Publish an existing product listing to your sales channel app
func (s *ProductListingServiceOp) Publish(ctx context.Context, productID int64) (*ProductListing, error) {
	path := fmt.Sprintf("%s/%v.json", productListingBasePath, productID)
	wrappedData := new(ProductListingPublishResource)
	wrappedData.ProductListing.ProductID = productID
	resource := new(ProductListingResource)
	err := s.client.Put(ctx, path, wrappedData, resource)
	return resource.ProductListing, err
}

// Delete unpublishes an existing product from your sales channel app.
func (s *ProductListingServiceOp) Delete(ctx context.Context, productID int64) error {
	return s.client.Delete(ctx, fmt.Sprintf("%s/%d.json", productListingBasePath, productID))
}
