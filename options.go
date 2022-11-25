package goshopify

import (
	"context"
	"fmt"
	"net/http"
)

// Option is used to configure client with options
type Option func(c *Client)

// WithContext optionally sets the custom context
func WithContext(ctx context.Context) Option {
	return func(c *Client) {
		c.ctx = ctx
	}
}

// WithVersion optionally sets the api-version if the passed string is valid
func WithVersion(apiVersion string) Option {
	return func(c *Client) {
		pathPrefix := defaultApiPathPrefix
		if len(apiVersion) > 0 && (apiVersionRegex.MatchString(apiVersion) || apiVersion == UnstableApiVersion) {
			pathPrefix = fmt.Sprintf("admin/api/%s", apiVersion)
		}
		c.apiVersion = apiVersion
		c.pathPrefix = pathPrefix
	}
}

func WithRetry(retries int) Option {
	return func(c *Client) {
		c.retries = retries
	}
}

func WithLogger(logger LeveledLoggerInterface) Option {
	return func(c *Client) {
		c.log = logger
	}
}

// WithHTTPClient is used to set a custom http client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.Client = client
	}
}
