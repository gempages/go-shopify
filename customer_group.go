package goshopify

import (
	"context"
	"fmt"
	"time"
)

const customerSavedSearchBasePath = "customer_saved_searches"
const customerSavedSearchName = "customerSavedSearch"

// CustomerSavedSearchService is an interface for interfacing with the customer group endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customersavedsearch
type CustomerSavedSearchService interface {
	List(context.Context, interface{}) ([]CustomerSavedSearch, error)
	GetBySinceId(ctx context.Context, sinceId int64, limit int, options interface{}) ([]CustomerSavedSearch, error)
	MetafieldsService
}

// CustomerSavedSearchServiceOp handles communication with the product related methods of
// the Shopify API.
type CustomerSavedSearchServiceOp struct {
	client *Client
}

func (s *CustomerSavedSearchServiceOp) ListMetafields(ctx context.Context, i int64, i2 interface{}) ([]Metafield, error) {
	panic("implement me")
}

func (s *CustomerSavedSearchServiceOp) CountMetafields(ctx context.Context, i int64, i2 interface{}) (int, error) {
	panic("implement me")
}

func (s *CustomerSavedSearchServiceOp) GetMetafield(ctx context.Context, i int64, i2 int64, i3 interface{}) (*Metafield, error) {
	panic("implement me")
}

func (s *CustomerSavedSearchServiceOp) CreateMetafield(ctx context.Context, i int64, metafield Metafield) (*Metafield, error) {
	panic("implement me")
}

func (s *CustomerSavedSearchServiceOp) UpdateMetafield(ctx context.Context, i int64, metafield Metafield) (*Metafield, error) {
	panic("implement me")
}

func (s *CustomerSavedSearchServiceOp) DeleteMetafield(ctx context.Context, i int64, i2 int64) error {
	panic("implement me")
}

// CustomerSavedSearch represents a Shopify CustomerSavedSearch
type CustomerSavedSearch struct {
	ID                int64       `json:"id,omitempty"`
	Name              string      `json:"name,omitempty"`
	CreatedAt         *time.Time  `json:"created_at,omitempty"`
	UpdatedAt         *time.Time  `json:"updated_at,omitempty"`
	Metafields        []Metafield `json:"metafields,omitempty"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Query             string      `json:"query,omitempty"`
}

// CustomerSavedSearchResource Represents the result from the customers/X.json endpoint
type CustomerSavedSearchResource struct {
	CustomerSavedSearch *CustomerSavedSearch `json:"customer_saved_search"`
}

// CustomerSavedSearchesResource Represents the result from the customers.json endpoint
type CustomerSavedSearchesResource struct {
	CustomerSavedSearches []CustomerSavedSearch `json:"customer_saved_searches"`
}

// List customers
func (s *CustomerSavedSearchServiceOp) List(ctx context.Context, options interface{}) ([]CustomerSavedSearch, error) {
	path := fmt.Sprintf("%s.json", customerSavedSearchBasePath)
	resource := new(CustomerSavedSearchesResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.CustomerSavedSearches, err
}

func (s *CustomerSavedSearchServiceOp) GetBySinceId(ctx context.Context, sinceId int64, limit int, options interface{}) ([]CustomerSavedSearch, error) {
	path := fmt.Sprintf("%s.json?since_id=%v&limit=%v", customerSavedSearchBasePath, sinceId, limit)
	resource := new(CustomerSavedSearchesResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.CustomerSavedSearches, err
}
