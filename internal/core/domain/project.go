package domain

import (
	"regexp"
	"strings"

	"github.com/go-pg/pg/types"
	"github.com/google/uuid"
)

// consumer_key ck_8c344061bcfda558d2f114efb8d1b892b4330a73

// Consumer secret  cs_562132bdb7e4e4c9e37b53a0fb703718e2dad5f7

// ProjectType is an enum for project type of woo or shopify
type ProjectType string

// UserRole enum values
const (
	Woocommerce ProjectType = "Woocommerce"
	Shopify     ProjectType = "Shopify"
)

// Project represents a project
type Project struct {
	tableName struct{} `pg:"project,alias:project"`
	Base
	Name        string      `json:"name" pg:"name,notnull"`
	Description string      `json:"description" pg:"description,notnull"`
	ProjectType ProjectType `json:"project_type" pg:"project_type,notnull"`
	WoocommerceProject
	ShopifyProject
	ValidatedAt types.NullTime
	Users       []*User `pg:"many2many:user_projects"`
}

// NewProject creates a instance of user with hashed password
func NewProject(p *ProjectRequest) (*Project, error) {

	pj := new(Project)
	pj.Name = p.Name
	pj.Description = p.Description
	pj.ProjectType = p.ProjectType
	return pj, nil

}

// WoocommerceProject represents a shopify project
type WoocommerceProject struct {
	Name           string `json:"woocommerce_name" pg:"woocommerce_name"`
	Description    string `json:"woocommerce_description" pg:"woocommerce_description"`
	Domain         string `json:"woocommerce_domain" pg:"woocommerce_domain"`
	ConsumerKey    string `json:"woocommerce_consumer_key" pg:"woocommerce_consumer_key"`
	ConsumerSecret string `json:"woocommerce_consumer_secret" pg:"woocommerce_consumer_secret"`
}

// NewWoocommerceProject creates a instance of user with hashed password
func NewWoocommerceProject(p *ProjectRequest) (WoocommerceProject, error) {
	wc := WoocommerceProject{}
	wc.Name = p.Name
	wc.Description = p.Description
	wc.Domain = strings.TrimRight(p.Domain, "/")
	wc.ConsumerKey = p.ConsumerKey
	wc.ConsumerSecret = p.ConsumerSecret
	return wc, nil
}

// ShopifyProject represents a shopify project
type ShopifyProject struct {
	Name           string `json:"shopify_consumer_name" pg:"shopify_consumer_name"`
	Description    string `json:"shopify_description" pg:"shopify_description"`
	Domain         string `json:"shopify_domain" pg:"shopify_domain"`
	ConsumerKey    string `json:"shopify_consumer_key" pg:"shopify_consumer_key"`
	ConsumerSecret string `json:"shopify_consumer_secret" pg:"shopify_consumer_secret"`
}

// NewShopifyProject creates a instance of user with hashed password
func NewShopifyProject(p *ProjectRequest) (ShopifyProject, error) {
	sp := ShopifyProject{}
	sp.Name = p.Name
	sp.Description = p.Description
	sp.Domain = strings.TrimRight(p.Domain, "/")
	sp.ConsumerKey = p.ConsumerKey
	sp.ConsumerSecret = p.ConsumerSecret
	return sp, nil
}

// ProjectRequest represents the request body for creating a user
type ProjectRequest struct {
	Name           string      `json:"name" form:"name"`
	Description    string      `json:"description" form:"description"`
	ProjectType    ProjectType `json:"project_type" form:"project_type"`
	Domain         string      `json:"domain" form:"domain"`
	ConsumerKey    string      `json:"consumer_key" form:"consumer_key"`
	ConsumerSecret string      `json:"consumer_secret" form:"consumer_secret"`
	AccessToken    string      `json:"accesstoken" form:"accesstoken"`
}

// Validate validates the request body
func (p *ProjectRequest) Validate() map[string](string) {

	errors := make(map[string]string)

	if p.Name == "" {
		errors["name"] = "Name is required"
	}
	if p.Description == "" {
		errors["description"] = "Description is required"
	}
	if p.Domain == "" {
		errors["domain"] = "Domain is required"
	} else if !isValidHttpsURL(p.Domain) {
		errors["domain"] = "Domain must be a valid HTTPS URL"
	}
	if p.ConsumerKey == "" {
		errors["consumer_key"] = "Consumer key is required"
	}
	if p.ConsumerSecret == "" {
		errors["consumer_secret"] = "Consumer secret is required"
	}
	if p.ConsumerKey == p.ConsumerSecret {
		errors["consumer_key"] = "Consumer Key is the same as Consumer Secret"
		errors["consumer_secret"] = "Consumer Secret is the same as Consumer Key"
	}

	return errors
}

// FindProjectRequest represents the request body for creating a user
type FindProjectRequest struct {
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	ProjectType ProjectType `json:"project_type" form:"project_type"`
	Domain      string      `json:"domain" form:"domain"`
}

// isValidHttpsURL checks if a given URL is a valid HTTPS URL
func isValidHttpsURL(url string) bool {
	re := regexp.MustCompile(`^https:\/\/[a-zA-Z0-9\-._~:\/?#@!$&'()*+,;=%]+$`)
	return re.MatchString(url)
}

// AddUser to user
func (u *Project) AddUser(userID uuid.UUID) {

}
