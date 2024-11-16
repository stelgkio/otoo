package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// HermesVoucerRequest represents a voucher record
type HermesVoucerRequest struct {
	ReceiverName      string   `json:"ReceiverName" bson:"ReceiverName" validate:"required,max=64"`                                 // Receiver's name (max 64 chars)
	ReceiverAddress   string   `json:"ReceiverAddress" bson:"ReceiverAddress" validate:"required_unless=ServiceReception 1,max=64"` // Receiver's address (max 64 chars) *Required if ServiceReception is not selected
	ReceiverCity      string   `json:"ReceiverCity" bson:"ReceiverCity" validate:"required,max=64"`                                 // Receiver's city (max 64 chars)
	ReceiverPostal    int      `json:"ReceiverPostal" bson:"ReceiverPostal" validate:"required,min=10000,max=99999"`                // Receiver's postal code (5 digits)
	ReceiverTelephone string   `json:"ReceiverTelephone" bson:"ReceiverTelephone" validate:"required,max=32,telephone"`
	ReceiverEmail     string   `json:"ReceiverEmail" bson:"ReceiverEmail" validate:"required,email"`  // Receiver's telephone (custom telephone format validation)
	Notes             *string  `json:"Notes" bson:"Notes" validate:"omitempty,max=128"`               // Delivery instructions (max 128 chars)
	OrderID           string   `json:"OrderID" bson:"OrderID" validate:"required,max=36"`             // Internal order note (max 36 chars)
	Cod               float64  `json:"Cod" bson:"Cod" validate:"required,max=499.99"`                 // Cash on Delivery amount (max 499.99)
	ServiceSavvato    *int     `json:"ServiceSavvato" bson:"ServiceSavvato" validate:"oneof=0 1"`     // Service option for Saturday delivery (0 or 1)
	ServiceEpigon     *int     `json:"ServiceEpigon" bson:"ServiceEpigon" validate:"oneof=0 1"`       // Service option for next-day delivery (0 or 1)
	ServiceEpistrofi  *int     `json:"ServiceEpistrofi" bson:"ServiceEpistrofi" validate:"oneof=0 1"` // Service option for return (0 or 1)
	ServiceSameday    *int     `json:"ServiceSameday" bson:"ServiceSameday" validate:"oneof=0 1"`     // Service option for same-day delivery (0 or 1)
	ServiceProtocol   *int     `json:"ServiceProtocol" bson:"ServiceProtocol" validate:"oneof=0 1"`   // Service option for protocol number (0 or 1)
	ServiceReception  *int     `json:"ServiceReception" bson:"ServiceReception" validate:"oneof=0 1"` // Service option for pickup from store (0 or 1)
	ParcelWeight      int      `json:"ParcelWeight" bson:"ParcelWeight" validate:"required,min=1"`    // Parcel weight (minimum 1)
	ParcelDepth       *float64 `json:"ParcelDepth" bson:"ParcelDepth" validate:"required,min=1.0"`    // Parcel depth (minimum 1.0)
	ParcelWidth       *float64 `json:"ParcelWidth" bson:"ParcelWidth" validate:"required,min=1.0"`    // Parcel width (minimum 1.0)
	ParcelHeight      *float64 `json:"ParcelHeight" bson:"ParcelHeight" validate:"required,min=1.0"`
	CustomOrderID     bool     `json:"CustomOrderID"` // Parcel height (minimum 1.0)

}

// Validate checks the fields of HermesVoucerRequest for validation errors.
func (r *HermesVoucerRequest) Validate() error {
	// Validate required fields
	if r.ReceiverName == "" || len(r.ReceiverName) > 64 {
		return errors.New("invalid ReceiverName: required, max 64 characters")
	}
	if r.ReceiverCity == "" || len(r.ReceiverCity) > 64 {
		return errors.New("invalid ReceiverCity: required, max 64 characters")
	}
	if r.ReceiverPostal < 10000 || r.ReceiverPostal > 99999 {
		return errors.New("invalid ReceiverPostal: must be a 5-digit postal code between 10000 and 99999")
	}
	if r.ReceiverTelephone == "" || len(r.ReceiverTelephone) > 32 {
		return errors.New("invalid ReceiverTelephone: required, max 32 characters")
	}
	if r.OrderID == "" || len(r.OrderID) > 36 {
		return errors.New("invalid OrderID: required, max 36 characters")
	}
	if r.Cod < 0 || r.Cod > 499.99 {
		return errors.New("invalid Cod: must be between 0 and 499.99")
	}
	if r.ParcelWeight < 1 {
		return errors.New("invalid ParcelWeight: minimum value is 1")
	}

	// Conditionally required fields
	if r.ServiceReception == nil || *r.ServiceReception != 1 {
		if r.ReceiverAddress == "" || len(r.ReceiverAddress) > 64 {
			return errors.New("invalid ReceiverAddress: required unless ServiceReception is 1, max 64 characters")
		}
	}

	// Optional Notes
	if r.Notes != nil && len(*r.Notes) > 128 {
		return errors.New("invalid Notes: max 128 characters")
	}

	// Validate service options (must be 0 or 1)
	services := map[string]*int{
		"ServiceSavvato":   r.ServiceSavvato,
		"ServiceEpigon":    r.ServiceEpigon,
		"ServiceEpistrofi": r.ServiceEpistrofi,
		"ServiceSameday":   r.ServiceSameday,
		"ServiceProtocol":  r.ServiceProtocol,
		"ServiceReception": r.ServiceReception,
	}
	for name, service := range services {
		if service != nil && *service != 0 && *service != 1 {
			return fmt.Errorf("invalid %s: must be 0 or 1 if set", name)
		}
	}

	// Parcel dimensions (all required if any one is provided)
	if r.ParcelDepth != nil || r.ParcelWidth != nil || r.ParcelHeight != nil {
		if r.ParcelDepth == nil || *r.ParcelDepth < 1.0 {
			return errors.New("invalid ParcelDepth: required if any parcel dimension is set, minimum 1.0")
		}
		if r.ParcelWidth == nil || *r.ParcelWidth < 1.0 {
			return errors.New("invalid ParcelWidth: required if any parcel dimension is set, minimum 1.0")
		}
		if r.ParcelHeight == nil || *r.ParcelHeight < 1.0 {
			return errors.New("invalid ParcelHeight: required if any parcel dimension is set, minimum 1.0")
		}
	}

	return nil
}

// WithNotes Optional field setter for Notes
func WithNotes(notes string) func(*HermesVoucerRequest) error {
	return func(r *HermesVoucerRequest) error {
		if len(notes) > 128 {
			return errors.New("notes exceed max length of 128")
		}
		r.Notes = &notes
		return nil
	}
}

// WithServiceSavvato  Optional field setter for ServiceSavvato
func WithServiceSavvato(enabled int) func(*HermesVoucerRequest) error {
	return func(r *HermesVoucerRequest) error {
		if enabled != 0 && enabled != 1 {
			return errors.New("invalid value for ServiceSavvato; must be 0 or 1")
		}
		r.ServiceSavvato = &enabled
		return nil
	}
}

// WithParcelDepth Optional field setter for ParcelDepth
func WithParcelDepth(depth float64) func(*HermesVoucerRequest) error {
	return func(r *HermesVoucerRequest) error {
		if depth < 1.0 {
			return errors.New("parcel depth must be at least 1.0")
		}
		r.ParcelDepth = &depth
		return nil
	}
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

// VoucherResponse struct represents the response from the voucher serviceß
type VoucherResponse struct {
	Success bool   `json:"success"`           // Indicates if the operation was successful
	Error   bool   `json:"error"`             // Indicates if there was an error
	Message string `json:"message"`           // Contains the response message
	Voucher int64  `json:"voucher,omitempty"` // Voucher ID (optional, might be empty if not applicable)
}

// VoucherPrintResponse defines the expected structure of the response for the PrintVouchers endpoint
type VoucherPrintResponse struct {
	Success             bool             `json:"success"`               // Indicates if the operation was successful
	Error               bool             `json:"error"`                 // Indicates if there was an error
	SingleVoucherFailed bool             `json:"single_voucher_failed"` // Indicates if a single voucher failed
	Message             string           `json:"message"`               // Message containing any issues
	Data                []VoucherFailure `json:"data,omitempty"`        // Vouchers that failed to print
}

// VoucherFailure represents a voucher that failed to print
type VoucherFailure struct {
	Voucher int `json:"voucher"` // Voucher ID
}

// HermesVoucerUpdateRequest defines the expected structure of the request for the UpdateVoucher endpoint
type HermesVoucerUpdateRequest struct {
	VoucherID int64 `json:"voucher"`
	HermesVoucerRequest
}

// NewHermesVoucerUpdateRequest creates a new instance of HermesVoucerUpdateRequest
func NewHermesVoucerUpdateRequest(voucherID int64, request *HermesVoucerRequest) *HermesVoucerUpdateRequest {
	return &HermesVoucerUpdateRequest{
		VoucherID:           voucherID,
		HermesVoucerRequest: *request,
	}
}

// TrackingStatus defines the expected structure of the response for the UpdateVoucher endpoint
type TrackingStatus struct {
	StatusID int    `json:"status_id"`
	Status   string `json:"status"`
	Date     string `json:"date"`
}

// TrackingResponse defines the expected structure of the response for the Tracking endpoint
type TrackingResponse struct {
	Success bool             `json:"success"`
	Error   bool             `json:"error"`
	Message string           `json:"message"`
	Data    []TrackingStatus `json:"data"`
}
