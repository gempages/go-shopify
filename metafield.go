package goshopify

import (
	"context"
	"fmt"
	"time"
)

// MetafieldService is an interface for interfacing with the metafield endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldService interface {
	List(context.Context, interface{}) ([]Metafield, error)
	Count(context.Context, interface{}) (int, error)
	Get(context.Context, int64, interface{}) (*Metafield, error)
	Create(context.Context, Metafield) (*Metafield, error)
	Update(context.Context, Metafield) (*Metafield, error)
	Delete(context.Context, int64) error
}

// MetafieldsService is an interface for other Shopify resources
// to interface with the metafield endpoints of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldsService interface {
	ListMetafields(context.Context, int64, interface{}) ([]Metafield, error)
	CountMetafields(context.Context, int64, interface{}) (int, error)
	GetMetafield(context.Context, int64, int64, interface{}) (*Metafield, error)
	CreateMetafield(context.Context, int64, Metafield) (*Metafield, error)
	UpdateMetafield(context.Context, int64, Metafield) (*Metafield, error)
	DeleteMetafield(context.Context, int64, int64) error
}

// MetafieldServiceOp handles communication with the metafield
// related methods of the Shopify API.
type MetafieldServiceOp struct {
	client     *Client
	resource   string
	resourceID int64
}

// Metafield represents a Shopify metafield.
type Metafield struct {
	ID                int64       `json:"id,omitempty"`
	Key               string      `json:"key,omitempty"`
	Value             interface{} `json:"value,omitempty"`
	ValueType         string      `json:"value_type,omitempty"`
	Namespace         string      `json:"namespace,omitempty"`
	Description       string      `json:"description,omitempty"`
	OwnerId           int64       `json:"owner_id,omitempty"`
	CreatedAt         *time.Time  `json:"created_at,omitempty"`
	UpdatedAt         *time.Time  `json:"updated_at,omitempty"`
	OwnerResource     string      `json:"owner_resource,omitempty"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id,omitempty"`
}

// MetafieldResource represents the result from the metafields/X.json endpoint
type MetafieldResource struct {
	Metafield *Metafield `json:"metafield"`
}

// MetafieldsResource represents the result from the metafields.json endpoint
type MetafieldsResource struct {
	Metafields []Metafield `json:"metafields"`
}

// List metafields
func (s *MetafieldServiceOp) List(ctx context.Context, options interface{}) ([]Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	resource := new(MetafieldsResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Metafields, err
}

// Count metafields
func (s *MetafieldServiceOp) Count(ctx context.Context, options interface{}) (int, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return s.client.Count(ctx, path, options)
}

// Get individual metafield
func (s *MetafieldServiceOp) Get(ctx context.Context, metafieldID int64, options interface{}) (*Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafieldID)
	resource := new(MetafieldResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Metafield, err
}

// Create a new metafield
func (s *MetafieldServiceOp) Create(ctx context.Context, metafield Metafield) (*Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Post(ctx, path, wrappedData, resource)
	return resource.Metafield, err
}

// Update an existing metafield
func (s *MetafieldServiceOp) Update(ctx context.Context, metafield Metafield) (*Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafield.ID)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Put(ctx, path, wrappedData, resource)
	return resource.Metafield, err
}

// Delete an existing metafield
func (s *MetafieldServiceOp) Delete(ctx context.Context, metafieldID int64) error {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	return s.client.Delete(ctx, fmt.Sprintf("%s/%d.json", prefix, metafieldID))
}
