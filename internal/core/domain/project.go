package domain

import (
	"github.com/go-pg/pg/types"
	"github.com/google/uuid"
)

// consumer_key ck_8c344061bcfda558d2f114efb8d1b892b4330a73

// Consumer secret  cs_562132bdb7e4e4c9e37b53a0fb703718e2dad5f7

// ProjectType is an enum for user's role
type ProjectType string

// UserRole enum values
const (
	Woocommerce ProjectType = "Woocommerce"
	Shopify     ProjectType = "Shopify"
)

type Project struct {
	Base
	Name        string
	Description string
	ProjectType ProjectType

	WoocommerceProjectId uuid.UUID           `pg:"woocommerce_id,type:uuid"`
	WoocommerceProject   *WoocommerceProject `pg:"rel:has-one"`

	ShopifyProjectId uuid.UUID       `pg:"shopify_id,type:uuid"`
	ShopifyProject   *ShopifyProject `pg:"rel:has-one"`

	UserId uuid.UUID `pg:"fk:user_id,type:uuid"`
	User   *User     `pg:"rel:has-one"`

	ValidatedAt types.NullTime
	IsActive    bool
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
	Base
	Name           string
	Description    string
	Domain         string
	ConsumerKey    string
	ConsumerSecret string
}

func NewWoocommerceProject(p *ProjectRequest) (*WoocommerceProject, error) {
	wc := new(WoocommerceProject)
	wc.Name = p.Name
	wc.Description = p.Description
	wc.Domain = p.Domain
	wc.ConsumerKey = p.ConsumerKey
	wc.ConsumerSecret = p.ConsumerSecret
	return wc, nil
}

type ShopifyProject struct {
	Base
	Name           string
	Description    string
	Domain         string
	ConsumerKey    string
	ConsumerSecret string
}

func NewShopifyProject(p *ProjectRequest) (*ShopifyProject, error) {
	sp := new(ShopifyProject)
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
}
