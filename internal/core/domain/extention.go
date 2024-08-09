package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Extention represents available extention
type Extention struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Type        string             `json:"type" bson:"type,omitempty"`
	Price       float64            `json:"price" bson:"price,omitempty"`
	MongoBase
}

// ProjectExtention represents active project extention
type ProjectExtention struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Type        string             `json:"type" bson:"type,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string             `json:"project_id" bson:"project_id,omitempty"`
	ExtentionID string             `json:"extention_id" bson:"extention_id,omitempty"`
	MongoBase
}

// AcsCourierExtention represents active acs details for acs courier
type AcsCourierExtention struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Type        string             `json:"type" bson:"type,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string             `json:"project_id" bson:"project_id,omitempty"`
	ExtentionID string             `json:"extention_id" bson:"extention_id,omitempty"`
	MongoBase
}
