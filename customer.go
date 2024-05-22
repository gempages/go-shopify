package goshopify

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const customersBasePath = "customers"
const customersResourceName = "customers"

// CustomerService is an interface for interfacing with the customers endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customer
type CustomerService interface {
	List(context.Context, interface{}) ([]Customer, error)
	GetBySinceId(ctx context.Context, sinceId int64, limit int, options interface{}) ([]Customer, error)
	Count(context.Context, interface{}) (int, error)
	Get(context.Context, int64, interface{}) (*Customer, error)
	Search(context.Context, interface{}) ([]Customer, error)
	Create(context.Context, Customer) (*Customer, error)
	Update(context.Context, Customer) (*Customer, error)
	Delete(context.Context, int64) error
	ListOrders(context.Context, int64, interface{}) ([]Order, error)
	ListTags(context.Context, interface{}) ([]string, error)

	// MetafieldsService used for Customer resource to communicate with Metafields resource
	MetafieldsService
}

// CustomerServiceOp handles communication with the product related methods of
// the Shopify API.
type CustomerServiceOp struct {
	client *Client
}

// Customer represents a Shopify customer
type Customer struct {
	ID                  int64              `json:"id,omitempty"`
	Email               string             `json:"email,omitempty"`
	FirstName           string             `json:"first_name,omitempty"`
	LastName            string             `json:"last_name,omitempty"`
	State               string             `json:"state,omitempty"`
	Note                string             `json:"note,omitempty"`
	VerifiedEmail       bool               `json:"verified_email,omitempty"`
	MultipassIdentifier string             `json:"multipass_identifier,omitempty"`
	OrdersCount         int                `json:"orders_count,omitempty"`
	TaxExempt           bool               `json:"tax_exempt,omitempty"`
	TotalSpent          *decimal.Decimal   `json:"total_spent,omitempty"`
	Phone               string             `json:"phone,omitempty"`
	Tags                string             `json:"tags,omitempty"`
	LastOrderId         int64              `json:"last_order_id,omitempty"`
	LastOrderName       string             `json:"last_order_name,omitempty"`
	AcceptsMarketing    bool               `json:"accepts_marketing,omitempty"`
	DefaultAddress      *CustomerAddress   `json:"default_address,omitempty"`
	Addresses           []*CustomerAddress `json:"addresses,omitempty"`
	CreatedAt           *time.Time         `json:"created_at,omitempty"`
	UpdatedAt           *time.Time         `json:"updated_at,omitempty"`
	Metafields          []Metafield        `json:"metafields,omitempty"`
	AdminGraphqlAPIID   string             `json:"admin_graphql_api_id"`
	Currency            string             `json:"currency"`
}

// Represents the result from the customers/X.json endpoint
type CustomerResource struct {
	Customer *Customer `json:"customer"`
}

// Represents the result from the customers.json endpoint
type CustomersResource struct {
	Customers []Customer `json:"customers"`
}

// Represents the result from the customers/tags.json endpoint
type CustomerTagsResource struct {
	Tags []string `json:"tags"`
}

// Represents the options available when searching for a customer
type CustomerSearchOptions struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Fields string `url:"fields,omitempty"`
	Order  string `url:"order,omitempty"`
	Query  string `url:"query,omitempty"`
}

// List customers
func (s *CustomerServiceOp) List(ctx context.Context, options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Customers, err
}

func (s *CustomerServiceOp) GetBySinceId(ctx context.Context, sinceId int64, limit int, options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s.json?since_id=%v&limit=%v", customersBasePath, sinceId, limit)
	resource := new(CustomersResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Customers, err
}

// Count customers
func (s *CustomerServiceOp) Count(ctx context.Context, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customersBasePath)
	return s.client.Count(ctx, path, options)
}

// Get customer
func (s *CustomerServiceOp) Get(ctx context.Context, customerID int64, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%v.json", customersBasePath, customerID)
	resource := new(CustomerResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Customer, err
}

// Create a new customer
func (s *CustomerServiceOp) Create(ctx context.Context, customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	wrappedData := CustomerResource{Customer: &customer}
	resource := new(CustomerResource)
	err := s.client.Post(ctx, path, wrappedData, resource)
	return resource.Customer, err
}

// Update an existing customer
func (s *CustomerServiceOp) Update(ctx context.Context, customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customer.ID)
	wrappedData := CustomerResource{Customer: &customer}
	resource := new(CustomerResource)
	err := s.client.Put(ctx, path, wrappedData, resource)
	return resource.Customer, err
}

// Delete an existing customer
func (s *CustomerServiceOp) Delete(ctx context.Context, customerID int64) error {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customerID)
	return s.client.Delete(ctx, path)
}

// Search customers
func (s *CustomerServiceOp) Search(ctx context.Context, options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s/search.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Customers, err
}

// ListOrders retrieves all orders from a customer
func (s *CustomerServiceOp) ListOrders(ctx context.Context, customerID int64, options interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s/%d/orders.json", customersBasePath, customerID)
	resource := new(OrdersResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Orders, err
}

// ListTags retrieves all unique tags across all customers
func (s *CustomerServiceOp) ListTags(ctx context.Context, options interface{}) ([]string, error) {
	path := fmt.Sprintf("%s/tags.json", customersBasePath)
	resource := new(CustomerTagsResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Tags, err
}

// List metafields for a customer
func (s *CustomerServiceOp) ListMetafields(ctx context.Context, customerID int64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.List(ctx, options)
}

// Count metafields for a customer
func (s *CustomerServiceOp) CountMetafields(ctx context.Context, customerID int64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Count(ctx, options)
}

// Get individual metafield for a customer
func (s *CustomerServiceOp) GetMetafield(ctx context.Context, customerID int64, metafieldID int64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Get(ctx, metafieldID, options)
}

// Create a new metafield for a customer
func (s *CustomerServiceOp) CreateMetafield(ctx context.Context, customerID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Create(ctx, metafield)
}

// Update an existing metafield for a customer
func (s *CustomerServiceOp) UpdateMetafield(ctx context.Context, customerID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Update(ctx, metafield)
}

// // Delete an existing metafield for a customer
func (s *CustomerServiceOp) DeleteMetafield(ctx context.Context, customerID int64, metafieldID int64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Delete(ctx, metafieldID)
}
