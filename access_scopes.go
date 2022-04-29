package goshopify

import (
	"fmt"
)

const AccessScopeBasePath = "access_scopes"

// AccessScopeService is an interface for interfacing with the storefront access
// token endpoints of the Shopify API.
// See: https://help.shopify.com/api/reference/access/storefrontaccesstoken
type AccessScopeService interface {
	List(interface{}) ([]AccessScope, error)
}

// AccessScopeServiceOp handles communication with the storefront access token
// related methods of the Shopify API.
type AccessScopeServiceOp struct {
	client *Client
}

// AccessScope represents a Shopify storefront access token
type AccessScope struct {
	Handle             string     `json:"handle,omitempty"`
}


// AccessScopesResource is the root object for a storefront access tokens get request.
type AccessScopesResource struct {
	AccessScopes []AccessScope `json:"access_scopes"`
}

// List storefront access tokens
func (s *AccessScopeServiceOp) List(options interface{}) ([]AccessScope, error) {
	path := fmt.Sprintf("%s.json", AccessScopeBasePath)
	resource := new(AccessScopesResource)
	err := s.client.Get(path, resource, options)
	return resource.AccessScopes, err
}