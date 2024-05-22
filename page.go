package goshopify

import (
	"context"
	"fmt"
	"time"
)

const pagesBasePath = "pages"
const pagesResourceName = "pages"

// PagesPageService is an interface for interacting with the pages
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/online_store/page
type PageService interface {
	List(context.Context, interface{}) ([]Page, error)
	GetBySinceId(context.Context, int64, int64, interface{}) ([]Page, error)
	Count(context.Context, interface{}) (int, error)
	Get(context.Context, int64, interface{}) (*Page, error)
	Create(context.Context, Page) (*Page, error)
	Update(context.Context, Page) (*Page, error)
	Delete(context.Context, int64) error

	// MetafieldsService used for Pages resource to communicate with Metafields
	// resource
	MetafieldsService
}

// PageServiceOp handles communication with the page related methods of the
// Shopify API.
type PageServiceOp struct {
	client *Client
}

// Page represents a Shopify page.
type Page struct {
	ID                int64       `json:"id,omitempty"`
	Author            string      `json:"author,omitempty"`
	Handle            string      `json:"handle,omitempty"`
	Title             string      `json:"title,omitempty"`
	CreatedAt         *time.Time  `json:"created_at,omitempty"`
	UpdatedAt         *time.Time  `json:"updated_at,omitempty"`
	BodyHTML          string      `json:"body_html,omitempty"`
	TemplateSuffix    string      `json:"template_suffix,omitempty"`
	PublishedAt       *time.Time  `json:"published_at,omitempty"`
	ShopID            int64       `json:"shop_id,omitempty"`
	Metafields        []Metafield `json:"metafields,omitempty"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
}

// PageResource represents the result from the pages/X.json endpoint
type PageResource struct {
	Page *Page `json:"page"`
}

// PagesResource represents the result from the pages.json endpoint
type PagesResource struct {
	Pages []Page `json:"pages"`
}

// List pages
func (s *PageServiceOp) List(ctx context.Context, options interface{}) ([]Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	resource := new(PagesResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Pages, err
}

func (s *PageServiceOp) GetBySinceId(ctx context.Context, sinceId int64, limit int64, options interface{}) ([]Page, error) {
	path := fmt.Sprintf("%s.json?since_id=%v&limit=%v", pagesBasePath, sinceId, limit)
	resource := new(PagesResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Pages, err
}

// Count pages
func (s *PageServiceOp) Count(ctx context.Context, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", pagesBasePath)
	return s.client.Count(ctx, path, options)
}

// Get individual page
func (s *PageServiceOp) Get(ctx context.Context, pageID int64, options interface{}) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, pageID)
	resource := new(PageResource)
	err := s.client.Get(ctx, path, resource, options)
	return resource.Page, err
}

// Create a new page
func (s *PageServiceOp) Create(ctx context.Context, page Page) (*Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := s.client.Post(ctx, path, wrappedData, resource)
	return resource.Page, err
}

// Update an existing page
func (s *PageServiceOp) Update(ctx context.Context, page Page) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, page.ID)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := s.client.Put(ctx, path, wrappedData, resource)
	return resource.Page, err
}

// Delete an existing page.
func (s *PageServiceOp) Delete(ctx context.Context, pageID int64) error {
	return s.client.Delete(ctx, fmt.Sprintf("%s/%d.json", pagesBasePath, pageID))
}

// List metafields for a page
func (s *PageServiceOp) ListMetafields(ctx context.Context, pageID int64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldService.List(ctx, options)
}

// Count metafields for a page
func (s *PageServiceOp) CountMetafields(ctx context.Context, pageID int64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldService.Count(ctx, options)
}

// Get individual metafield for a page
func (s *PageServiceOp) GetMetafield(ctx context.Context, pageID int64, metafieldID int64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldService.Get(ctx, metafieldID, options)
}

// Create a new metafield for a page
func (s *PageServiceOp) CreateMetafield(ctx context.Context, pageID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldService.Create(ctx, metafield)
}

// Update an existing metafield for a page
func (s *PageServiceOp) UpdateMetafield(ctx context.Context, pageID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldService.Update(ctx, metafield)
}

// Delete an existing metafield for a page
func (s *PageServiceOp) DeleteMetafield(ctx context.Context, pageID int64, metafieldID int64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldService.Delete(ctx, metafieldID)
}
