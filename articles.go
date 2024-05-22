package goshopify

import (
	"context"
	"fmt"
	"time"
)

// ArticleService is an interface for interfacing with the article endpoints
// of the Shopify API.
// See: https://shopify.dev/api/admin-rest/2021-10/resources/article#top
type ArticleService interface {
	GetByBlogID(ctx context.Context, blogID int64, limit int, sinceId int64, options interface{}) (*[]Article, error)
	GetByBlogIDAndArticleID(context.Context, int64, int64, interface{}) (*Article, error)
	GetCountByBlogID(context.Context, int64, interface{}) (int, error)
	Create(context.Context, int64, *Article) (*Article, error)
	Update(context.Context, int64, int64, *Article) (*Article, error)
	Delete(context.Context, int64, int64) error
}

// ArticleServiceOp handles communication with the blog related methods of
// the Shopify API.
type ArticleServiceOp struct {
	client *Client
}

type ImageArticle struct {
	CreatedAt *time.Time `json:"created_at"`
	Alt       string     `json:"alt"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Src       string     `json:"src"`
}

// Article represents a Shopify Article
type Article struct {
	ID                int64         `json:"id"`
	Title             string        `json:"title"`
	BodyHtml          string        `json:"body_html"`
	BlogID            int64         `json:"blog_id"`
	Author            string        `json:"author"`
	UserID            int64         `json:"user_id"`
	SummaryHtml       string        `json:"summary_html"`
	TemplateSuffix    string        `json:"template_suffix"`
	Handle            string        `json:"handle"`
	Tags              string        `json:"tags"`
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	Image             *ImageArticle `json:"image"`
	CreatedAt         *time.Time    `json:"created_at"`
	UpdatedAt         *time.Time    `json:"updated_at"`
	PublishedAt       *time.Time    `json:"published_at"`
}

// ArticlesResource is the result from the article.json endpoint
type ArticlesResource struct {
	Articles []Article `json:"articles"`
}

// ArticleResource Represents the result from the articles.json endpoint
type ArticleResource struct {
	Article *Article `json:"article"`
}

// GetByBlogID - Retrieves a list of all articles from a blog
func (s *ArticleServiceOp) GetByBlogID(ctx context.Context, blogID int64, limit int, sinceId int64, options interface{}) (*[]Article, error) {
	path := fmt.Sprintf("blogs/%v/articles.json?limit=%v&since_id=%v", blogID, limit, sinceId)
	resource := new(ArticlesResource)
	err := s.client.Get(ctx, path, resource, options)
	return &resource.Articles, err
}

// GetByBlogIDAndArticleID Receive a single Article
func (s *ArticleServiceOp) GetByBlogIDAndArticleID(ctx context.Context, blogID int64, articleID int64, options interface{}) (resource *Article, err error) {
	path := fmt.Sprintf("blogs/%v/articles/%v.json", blogID, articleID)
	err = s.client.Get(ctx, path, resource, options)
	return resource, err
}

// GetCountByBlogID - Retrieves a count of all articles from a blog
func (s *ArticleServiceOp) GetCountByBlogID(ctx context.Context, blogID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("blogs/%v/articles/count.json", blogID)
	return s.client.Count(ctx, path, options)
}

// Create - Create a new article
func (s *ArticleServiceOp) Create(ctx context.Context, blogID int64, article *Article) (*Article, error) {
	path := fmt.Sprintf("blogs/%v/articles.json", blogID)
	wrappedData := ArticleResource{Article: article}
	resource := ArticleResource{}
	err := s.client.Post(ctx, path, wrappedData, &resource)
	return resource.Article, err

}

// Update - Update a new article
func (s *ArticleServiceOp) Update(ctx context.Context, blogID int64, articleID int64, article *Article) (*Article, error) {
	path := fmt.Sprintf("blogs/%v/articles/%v.json", blogID, articleID)
	wrappedData := ArticleResource{Article: article}
	resource := new(ArticleResource)
	err := s.client.Put(ctx, path, wrappedData, resource)
	return resource.Article, err
}

// Delete an existing product
func (s *ArticleServiceOp) Delete(ctx context.Context, blogID int64, articleID int64) error {
	return s.client.Delete(ctx, fmt.Sprintf("blogs/%v/articles/%v.json", blogID, articleID))
}
