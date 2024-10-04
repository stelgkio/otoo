package domain

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// HermesVoucerRequest represents a voucher record
type HermesVoucerRequest struct {
	ReceiverName      string  `json:"receiver_name" bson:"receiver_name" validate:"required,max=64"`                                 // Receiver's name (max 64 chars)
	ReceiverAddress   string  `json:"receiver_address" bson:"receiver_address" validate:"required_unless=ServiceReception 1,max=64"` // Receiver's address (max 64 chars) *Required if ServiceReception is not selected
	ReceiverCity      string  `json:"receiver_city" bson:"receiver_city" validate:"required,max=64"`                                 // Receiver's city (max 64 chars)
	ReceiverPostal    int     `json:"receiver_postal" bson:"receiver_postal" validate:"required,min=10000,max=99999"`                // Receiver's postal code (5 digits)
	ReceiverTelephone string  `json:"receiver_telephone" bson:"receiver_telephone" validate:"required,max=32,telephone"`             // Receiver's telephone (custom telephone format validation)
	Notes             string  `json:"notes" bson:"notes" validate:"omitempty,max=128"`                                               // Delivery instructions (max 128 chars)
	OrderID           string  `json:"order_id" bson:"order_id" validate:"required,max=36"`                                           // Internal order note (max 36 chars)
	Cod               float64 `json:"cod" bson:"cod" validate:"required,max=499.99"`                                                 // Cash on Delivery amount (max 499.99)
	ServiceSavvato    int     `json:"service_savvato" bson:"service_savvato" validate:"oneof=0 1"`                                   // Service option for Saturday delivery (0 or 1)
	ServiceEpigon     int     `json:"service_epigon" bson:"service_epigon" validate:"oneof=0 1"`                                     // Service option for next-day delivery (0 or 1)
	ServiceEpistrofi  int     `json:"service_epistrofi" bson:"service_epistrofi" validate:"oneof=0 1"`                               // Service option for return (0 or 1)
	ServiceSameday    int     `json:"service_sameday" bson:"service_sameday" validate:"oneof=0 1"`                                   // Service option for same-day delivery (0 or 1)
	ServiceProtocol   int     `json:"service_protocol" bson:"service_protocol" validate:"oneof=0 1"`                                 // Service option for protocol number (0 or 1)
	ServiceReception  int     `json:"service_reception" bson:"service_reception" validate:"oneof=0 1"`                               // Service option for pickup from store (0 or 1)
	ParcelWeight      int     `json:"parcel_weight" bson:"parcel_weight" validate:"required,min=1"`                                  // Parcel weight (minimum 1)
	ParcelDepth       float64 `json:"parcel_depth" bson:"parcel_depth" validate:"required,min=1.0"`                                  // Parcel depth (minimum 1.0)
	ParcelWidth       float64 `json:"parcel_width" bson:"parcel_width" validate:"required,min=1.0"`                                  // Parcel width (minimum 1.0)
	ParcelHeight      float64 `json:"parcel_height" bson:"parcel_height" validate:"required,min=1.0"`                                // Parcel height (minimum 1.0)

}

// Custom validation for telephone number format
// telephoneValidator checks if the receiver's telephone numbers are valid
// Custom telephone validator function
func telephoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	// Split by allowed separators (space, comma, or dash)
	separators := []string{" ", ",", "-"}
	var numbers []string

	for _, sep := range separators {
		if strings.Contains(phone, sep) {
			numbers = strings.Split(phone, sep)
			break
		}
	}

	// If no separator is found or not exactly two parts, it's invalid
	if len(numbers) != 2 {
		return false
	}

	// Regular expression to match exactly 10 digits
	re := regexp.MustCompile(`^\d{10}$`)

	// Check if both parts are valid 10-digit numbers
	for _, number := range numbers {
		number = strings.TrimSpace(number)
		if !re.MatchString(number) {
			return false
		}
	}

	// If both numbers are valid, return true
	return true
}

// Helper function to check if a string is numeric (all digits)
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// ValidateParcel function to validate a Parcel struct and return any validation errors
func ValidateParcel(parcel HermesVoucerRequest) error {
	// Create a new validator instance
	validate := validator.New()

	// Register custom validation for telephone numbers
	validate.RegisterValidation("telephone", telephoneValidator)

	// Validate the Parcel struct
	err := validate.Struct(parcel)

	// Return any validation errors
	return err
}
