package domain

import (
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

type Contact struct {
	MongoBase
	FullName string `json:"full_name" bson:"full_name,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Phone    string `json:"phone" bson:"phone,omitempty"`
	Message  string `json:"message" bson:"message,omitempty"`
}

// NewProject creates a instance of user with hashed password
func NewContact(p *ContactRequest) (*Contact, error) {

	pj := new(Contact)
	pj.FullName = p.FullName
	pj.Email = p.Email
	pj.Phone = p.Phone
	pj.Message = p.Message
	pj.MongoBase = NewMongoBase()

	return pj, nil

}

type ContactRequest struct {
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Message  string `json:"message" form:"message" validate:"required"`
}

func (c *ContactRequest) Validate() map[string](string) {
	errors := make(map[string]string)

	// // Validate the contact struct
	// validate := validator.New()
	// if err := validate.Struct(c); err != nil {
	// 	errors["error"] = "Invalid struct"
	// 	return errors
	// }

	// Sanitize the input to prevent XSS
	p := bluemonday.UGCPolicy()
	c.FullName = p.Sanitize(c.FullName)
	c.Email = p.Sanitize(c.Email)
	c.Phone = p.Sanitize(c.Phone)
	c.Message = p.Sanitize(c.Message)

	if c.Email == "" {
		errors["email"] = "Email is required"
	} else {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(c.Email) {
			errors["email"] = "Invalid email format"
		}
	}
	if c.Phone == "" {
		errors["phone"] = "Phone is required"
	} else {
		phoneRegex := regexp.MustCompile(`^\d{10}$`) // Change regex pattern as per your phone number format
		if !phoneRegex.MatchString(c.Phone) {
			errors["phone"] = "Invalid phone format (should be 10 digits)"
		}
	}
	if c.Message == "" {
		errors["message"] = "Message is required"
	}
	if c.FullName == "" {
		errors["full_name"] = "FullName is required"
	}

	return errors
}
