package domain

import (
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

type Project struct {
	Base
	Name        string      `json:"name" pg:"name,notnull"`
	Description string      `json:"description" pg:"description,notnull"`
	ProjectType ProjectType `json:"project_type" pg:"project_type,notnull"`
	WoocommerceProject
	ShopifyProject
	UserId      uuid.UUID `pg:"fk:user_id,type:uuid"`
	User        *User     `pg:"rel:has-one"`
	ValidatedAt types.NullTime
	IsActive    bool `pg:"is_active"`
}

// NewProject creates a instance of user with hashed password
func NewProject(p *ProjectRequest) (*Project, error) {

	pj := new(Project)
	pj.Name = p.Name
	pj.Description = p.Description
	pj.ProjectType = p.ProjectType
	return pj, nil

}

type WoocommerceProject struct {
	Name           string `json:"woocommerce_name" pg:"woocommerce_name"`
	Description    string `json:"woocommerce_description" pg:"woocommerce_description"`
	Domain         string `json:"woocommerce_domain" pg:"woocommerce_domain"`
	ConsumerKey    string `json:"woocommerce_consumer_key" pg:"woocommerce_consumer_key"`
	ConsumerSecret string `json:"woocommerce_consumer_secret" pg:"woocommerce_consumer_secret"`
}

func NewWoocommerceProject(p *ProjectRequest) (WoocommerceProject, error) {
	wc := WoocommerceProject{}
	wc.Name = p.Name
	wc.Description = p.Description
	wc.Domain = p.Domain
	wc.ConsumerKey = p.ConsumerKey
	wc.ConsumerSecret = p.ConsumerSecret
	return wc, nil
}

type ShopifyProject struct {
	Name           string `json:"shopify_consumer_name" pg:"shopify_consumer_name"`
	Description    string `json:"shopify_description" pg:"shopify_description"`
	Domain         string `json:"shopify_domain" pg:"shopify_domain"`
	ConsumerKey    string `json:"shopify_consumer_key" pg:"shopify_consumer_key"`
	ConsumerSecret string `json:"shopify_consumer_secret" pg:"shopify_consumer_secret"`
}

func NewShopifyProject(p *ProjectRequest) (ShopifyProject, error) {
	sp := ShopifyProject{}
	sp.Name = p.Name
	sp.Description = p.Description
	sp.Domain = p.Domain
	sp.ConsumerKey = p.ConsumerKey
	sp.ConsumerSecret = p.ConsumerSecret
	return sp, nil
}

// projectRequest represents the request body for creating a user
type ProjectRequest struct {
	Name           string      `json:"name" form:"name"`
	Description    string      `json:"description" form:"description"`
	ProjectType    ProjectType `json:"project_type" form:"project_type"`
	Domain         string      `json:"domain" form:"domain"`
	ConsumerKey    string      `json:"consumer_key" form:"consumer_key"`
	ConsumerSecret string      `json:"consumer_secret" form:"consumer_secret"`
	AccessToken    string      `json:"accesstoken" form:"accesstoken"`
}

func (p *ProjectRequest) Validate() map[string](string) {

	errors := make(map[string]string)

	if p.Name == "" {
		errors["name"] = "Name is required"
	}
	if p.Domain == "" {
		errors["domain"] = "Domain is required"
	}
	if p.ConsumerKey == "" {
		errors["consumer_key"] = "Consumer key is required"
	}
	if p.ConsumerSecret == "" {
		errors["consumer_secret"] = "Consumer secret is required"
	}
	// if p.AccessToken == "" {
	// 	errors["accesstoken"] = "Access token is required"
	// }

	return errors
}

type FindProjectRequest struct {
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	ProjectType ProjectType `json:"project_type" form:"project_type"`
	Domain      string      `json:"domain" form:"domain"`
}
