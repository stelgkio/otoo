package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AcsCourier represents the structure for storing ACS Courier credentials in MongoDB
type AcsCourier struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`                            // MongoDB ObjectID
	ProjectID       string             `json:"project_id" bson:"project_id,omitempty"`   // Associated project ID
	CompanyID       string             `json:"company_id" bson:"company_id"`             // Company ID provided by ACS
	CompanyPassword string             `json:"company_password" bson:"company_password"` // Company password provided by ACS
	UserID          string             `json:"user_id" bson:"user_id"`                   // User ID provided by ACS
	UserPassword    string             `json:"user_password" bson:"user_password"`       // User password provided by ACS
	AcsAPIKey       string             `json:"acs_api_key" bson:"acs_api_key"`           // ACS API key
}

// CreateVoucherResponse represents the structure of the response from the Create Voucher API
type CreateVoucherResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSOutputResponse        struct {
		ACSValueOutput []struct {
			VoucherNo string `json:"Voucher_No"`
		} `json:"ACSValueOutput"`
	} `json:"ACSOutputResponce"`
}

// ACSRequest defines the structure of the ACS request
type ACSRequest struct {
	ACSAlias           string          `json:"ACSAlias"`
	ACSInputParameters *VoucherRequest `json:"ACSInputParameters"`
}

// VoucherRequest defines the structure of the Create Voucher request
type VoucherRequest struct {
	CompanyID              string   `json:"Company_ID" validate:"required"`
	CompanyPassword        string   `json:"Company_Password" validate:"required"`
	UserID                 string   `json:"User_ID" validate:"required"`
	UserPassword           string   `json:"User_Password" validate:"required"`
	PickupDate             string   `json:"Pickup_Date" validate:"required,datetime=2006-01-02"`
	Sender                 string   `json:"Sender" validate:"required"`
	RecipientName          string   `json:"Recipient_Name" validate:"required"`
	RecipientAddress       string   `json:"Recipient_Address" validate:"required"`
	RecipientAddressNumber int      `json:"Recipient_Address_Number" validate:"required,min=1"`
	RecipientZipcode       int      `json:"Recipient_Zipcode" validate:"required,min=1000,max=99999"`
	RecipientRegion        string   `json:"Recipient_Region" validate:"required"`
	RecipientPhone         int64    `json:"Recipient_Phone,omitempty" validate:"required_without=RecipientCellPhone"`
	RecipientCellPhone     int64    `json:"Recipient_Cell_Phone,omitempty" validate:"required_without=RecipientPhone"`
	RecipientFloor         *string  `json:"Recipient_Floor,omitempty"`
	RecipientCompanyName   *string  `json:"Recipient_Company_Name,omitempty"`
	RecipientCountry       string   `json:"Recipient_Country" validate:"required,oneof=GR CY"`
	AcsStationDestination  *string  `json:"Acs_Station_Destination,omitempty"`
	AcsStationBranchDest   int      `json:"Acs_Station_Branch_Destination" validate:"required,oneof=0 1"`
	BillingCode            string   `json:"Billing_Code" validate:"required"`
	ChargeType             int      `json:"Charge_Type" validate:"required,oneof=2 4"`
	CostCenterCode         *string  `json:"Cost_Center_Code,omitempty"`
	ItemQuantity           int      `json:"Item_Quantity" validate:"required,min=1"`
	Weight                 float64  `json:"Weight" validate:"required,min=0.5"`
	DimensionXInCm         *int     `json:"Dimension_X_In_Cm,omitempty"`
	DimensionYInCm         *int     `json:"Dimension_Y_in_Cm,omitempty"`
	DimensionZInCm         *int     `json:"Dimension_Z_in_Cm,omitempty"`
	CodAmount              *float64 `json:"Cod_Ammount,omitempty" validate:"omitempty,min=0"`
	CodPaymentWay          *int     `json:"Cod_Payment_Way,omitempty" validate:"omitempty,oneof=0 1"`
	AcsDeliveryProducts    string   `json:"Acs_Delivery_Products" validate:"required"`
	InsuranceAmount        *float64 `json:"Insurance_Ammount,omitempty" validate:"omitempty,min=0,max=3000"`
	DeliveryNotes          *string  `json:"Delivery_Notes,omitempty"`
	AppointmentUntilTime   *string  `json:"Appointment_Until_Time,omitempty" validate:"omitempty,datetime=15:04"`
	RecipientEmail         *string  `json:"Recipient_Email,omitempty" validate:"omitempty,email"`
	ReferenceKey1          *string  `json:"Reference_Key1,omitempty"`
	ReferenceKey2          *string  `json:"Reference_Key2,omitempty"`
	WithReturnVoucher      *int     `json:"With_Return_Voucher,omitempty" validate:"omitempty,oneof=0 1"`
	ContentTypeID          *int     `json:"Content_Type_ID,omitempty"`
	Language               string   `json:"Language" validate:"required,oneof=GR EN"`
}

// Validate ensures the VoucherRequest fields are valid
func (v *VoucherRequest) Validate() error {
	// Check if both RecipientPhone and RecipientCellPhone are empty
	if v.RecipientPhone == 0 && v.RecipientCellPhone == 0 {
		return errors.New("either RecipientPhone or RecipientCellPhone must be provided")
	}

	// Ensure the pickup date is not in the past
	pickupDate, err := time.Parse("2006-01-02", v.PickupDate)
	if err != nil {
		return errors.New("invalid Pickup_Date format, must be YYYY-MM-DD")
	}

	if pickupDate.Before(time.Now()) {
		return errors.New("Pickup_Date cannot be in the past")
	}

	// Validate RecipientCountry
	if v.RecipientCountry != "GR" && v.RecipientCountry != "CY" {
		return errors.New("Recipient_Country must be GR or CY")
	}

	// Validate ChargeType
	if v.ChargeType != 2 && v.ChargeType != 4 {
		return errors.New("Charge_Type must be 2 (sender) or 4 (recipient)")
	}

	// Validate CodPaymentWay
	if v.CodPaymentWay != nil && *v.CodPaymentWay != 0 && *v.CodPaymentWay != 1 {
		return errors.New("Cod_Payment_Way must be 0 (cash) or 1 (check)")
	}

	// Validate Weight
	if v.Weight < 0.5 {
		return errors.New("Weight must be at least 0.5 kg")
	}

	return nil
}

// ErrorMessages maps error codes to user-friendly descriptions
var ErrorMessages = map[string]string{
	"Invalid pick-up date": "Error: The pick-up date is invalid.",
	"Pickup date is not allowed on Sunday or national holiday":                         "Error: The pick-up date falls on a Sunday or national holiday.",
	"The recipient’s name cannot be empty":                                             "Error: The recipient's name is empty.",
	"The recipient's address cannot be empty":                                          "Error: The recipient's address is empty.",
	"Unacceptable zip code or country":                                                 "Error: The zip code or country is incorrect.",
	"Unacceptable ACS Destination Value":                                               "Error: The ACS destination value is invalid.",
	"Not more than 99 pieces per shipment are supported":                               "Error: The number of pieces exceeds the maximum limit.",
	"Unacceptable weight value (0.5-999)":                                              "Error: The weight value is out of acceptable range.",
	"Unacceptable shipping charge value":                                               "Error: The shipping charge value is incorrect.",
	"Unacceptable payment method for COD":                                              "Error: The payment method for COD is incorrect.",
	"Cannot find the COD":                                                              "Error: COD service is not available.",
	"No valid credit code":                                                             "Error: The billing code is invalid.",
	"The Acs-SmartPoint destination must have 1 mobile phone":                          "Error: A mobile phone is required for Acs-SmartPoint destination.",
	"The product \"RVO\" is only combined with return voucher (with_return = 1)":       "Error: The RVO product requires a return voucher.",
	"You cannot create Cyprus Economy (EC) shipments with the billing code: 2XXXXXXXX": "Error: The billing code is not valid for Cyprus Economy shipments.",
	"You cannot combine these products together":                                       "Error: The selected products cannot be combined.",
	"The product 5Σ is not supported for this destination":                             "Error: The selected product is not supported for the destination.",
	"If time is entered in the field \"Appointment_Until_Time\" then in the shipment will be automatically added the product of the Morning delivery": "Error: Time entered in Appointment_Until_Time automatically adds Morning delivery.",
}

// HandleAcsError processes and prints user-friendly error messages
func HandleAcsError(errMsg string) {
	if message, exists := ErrorMessages[errMsg]; exists {
		fmt.Println(message)
	} else {
		fmt.Printf("Error: %s\n", errMsg)
	}
}

// ACSMultipleRequest defines the structure of the ACS request
type ACSMultipleRequest struct {
	ACSAlias           string                   `json:"ACSAlias"`
	ACSInputParameters *MultipartVoucherRequest `json:"ACSInputParameters"`
}

// MultipartVoucherRequest represents the structure of the request for multipart vouchers
type MultipartVoucherRequest struct {
	CompanyID       string `json:"Company_ID"`
	CompanyPassword string `json:"Company_Password"`
	UserID          string `json:"User_ID"`
	UserPassword    string `json:"User_Password"`
	Language        string `json:"Language"`
	MainVoucherNo   string `json:"Main_Voucher_No"`
}

// MultipartVouchersResponse represents the structure of the response from the Multipart Vouchers API
type MultipartVouchersResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSOutputResponse        struct {
		ACSTableOutput struct {
			TableData []struct {
				MultiPartVoucherNo string `json:"MultiPart_Voucher_No"`
			} `json:"Table_Data"`
		} `json:"ACSTableOutput"`
	} `json:"ACSOutputResponce"`
}

// ACSMultipleRequest defines the structure of the ACS request
type ACSPrintRequest struct {
	ACSAlias           string               `json:"ACSAlias"`
	ACSInputParameters *PrintVoucherRequest `json:"ACSInputParameters"`
}

// PrintVoucherRequest represents the structure of the request for printing vouchers
type PrintVoucherRequest struct {
	CompanyID         string   `json:"Company_ID"`
	CompanyPassword   string   `json:"Company_Password"`
	UserID            string   `json:"User_ID"`
	UserPassword      string   `json:"User_Password"`
	VoucherNo         []string `json:"Voucher_No"`
	PrintType         int      `json:"Print_Type"`
	StartPosition     int      `json:"Start_Position"`
	WithReturnVoucher int      `json:"With_Return_Voucher"`
	ItemQuantity      int      `json:"Item_Quantity"`
}

// Validate validates the PrintVoucherRequest
func (req *PrintVoucherRequest) Validate() error {
	// Validate PrintType
	if req.PrintType != 1 && req.PrintType != 2 {
		return errors.New("invalid PrintType. Must be 1 for Thermal or 2 for Laser")
	}

	// Validate StartPosition
	if req.PrintType == 2 && (req.StartPosition < 1 || req.StartPosition > 3) {
		return errors.New("invalid StartPosition. Must be 1, 2, or 3 for Laser printing")
	}

	// Validate VoucherNo
	if len(req.VoucherNo) == 0 || len(req.VoucherNo) > 10 {
		return errors.New("VoucherNo should contain between 1 and 10 vouchers")
	}
	// Check if VoucherNo contains separators and process accordingly
	voucherNoStr := strings.Join(req.VoucherNo, "|")
	voucherCount := len(strings.Split(voucherNoStr, "|"))
	if voucherCount > 10 {
		return errors.New("VoucherNo should not contain more than 10 vouchers")
	}

	// Special handling for With_Return_Voucher and Item_Quantity
	if req.WithReturnVoucher == 1 && req.ItemQuantity > 1 {
		// Additional checks or modifications if needed
		// This is where you can add any specific handling logic
	}

	// No additional validations specified for reprinting
	return nil

}

// ACSDeleteRequest defines the structure of the ACS request
type ACSDeleteRequest struct {
	ACSAlias           string                `json:"ACSAlias"`
	ACSInputParameters *DeleteVoucherRequest `json:"ACSInputParameters"`
}

// DeleteVoucherRequest represents the structure of the request for deleting vouchers
type DeleteVoucherRequest struct {
	CompanyID       string   `json:"Company_ID"`
	CompanyPassword string   `json:"Company_Password"`
	UserID          string   `json:"User_ID"`
	UserPassword    string   `json:"User_Password"`
	VoucherNo       []string `json:"Voucher_No"`
}

// Validate validates the DeleteVoucherRequest
func (req *DeleteVoucherRequest) Validate() error {
	// Validate VoucherNo
	if len(req.VoucherNo) == 0 || len(req.VoucherNo) > 20 {
		return errors.New("VoucherNo should contain between 1 and 20 vouchers")
	}

	// Check if VoucherNo contains separators and process accordingly
	voucherNoStr := strings.Join(req.VoucherNo, ",")
	voucherCount := len(strings.Split(voucherNoStr, ","))
	if voucherCount > 20 {
		return errors.New("VoucherNo should not contain more than 20 vouchers")
	}

	// No additional validations specified for Language or other fields
	return nil
}

// DeleteVoucherResponse represents the structure of the response from the Delete Voucher API
type DeleteVoucherResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSOutputResponse        struct {
		ACSValueOutput []struct {
			ErrorMessage string `json:"Error_Message"`
		} `json:"ACSValueOutput"`
		ACSTableOutput struct{} `json:"ACSTableOutput"`
	} `json:"ACSOutputResponce"`
}

// ACSIssuePickupListRequest defines the structure of the ACS request
type ACSIssuePickupListRequest struct {
	ACSAlias           string                  `json:"ACSAlias"`
	ACSInputParameters *IssuePickupListRequest `json:"ACSInputParameters"`
}

// IssuePickupListRequest represents the structure of the request for issuing a pickup list
type IssuePickupListRequest struct {
	CompanyID       string  `json:"Company_ID"`
	CompanyPassword string  `json:"Company_Password"`
	UserID          string  `json:"User_ID"`
	UserPassword    string  `json:"User_Password"`
	Language        *string `json:"Language"`
	PickupDate      string  `json:"Pickup_Date"`
	MyData          *int    `json:"MyData"` // Use a pointer to allow null values
}

// Validate validates the IssuePickupListRequest
func (req *IssuePickupListRequest) Validate() error {
	if req.PickupDate == "" {
		return errors.New("Pickup_Date is required")
	}
	if req.MyData != nil && (*req.MyData != 0 && *req.MyData != 1) {
		return errors.New("MyData must be either 0 or 1")
	}
	return nil
}

// IssuePickupListResponse represents the structure of the response for issuing a pickup list
type IssuePickupListResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSValueOutput           []struct {
		PickupListNo   string `json:"PickupList_No"`
		UnprintedFound int    `json:"Unprinted_Found"`
		ErrorMessage   string `json:"Error_Message"`
	} `json:"ACSValueOutput"`
	ACSTableOutput struct {
		TableData []struct {
			UnprintedVouchers string `json:"Unprinted_Vouchers"`
		} `json:"Table_Data"`
	} `json:"ACSTableOutput"`
}

// ACSPrintPickupListRequest defines the structure of the ACS request
type ACSPrintPickupListRequest struct {
	ACSAlias           string                  `json:"ACSAlias"`
	ACSInputParameters *PrintPickupListRequest `json:"ACSInputParameters"`
}

// PrintPickupListRequest represents the structure of the request for printing a pickup list
type PrintPickupListRequest struct {
	CompanyID       string  `json:"Company_ID"`
	CompanyPassword string  `json:"Company_Password"`
	UserID          string  `json:"User_ID"`
	UserPassword    string  `json:"User_Password"`
	Language        *string `json:"Language"`
	MassNumber      string  `json:"Mass_Number"`
	PickupDate      string  `json:"Pickup_Date"`
}

// Validate validates the PrintPickupListRequest
func (req *PrintPickupListRequest) Validate() error {
	if req.MassNumber == "" {
		return errors.New("Mass_Number is required")
	}
	if req.PickupDate == "" {
		return errors.New("Pickup_Date is required")
	}
	return nil
}

// PrintPickupListResponse represents the response from printing a pickup list
type PrintPickupListResponse struct {
	// Assuming similar structure as other responses, but you may need to adjust based on actual API
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	// Response details will vary depending on how the printed list is handled
}

// ACSDisplayPickupListVouchersRequest defines the structure of the ACS request
type ACSDisplayPickupListVouchersRequest struct {
	ACSAlias           string                            `json:"ACSAlias"`
	ACSInputParameters *DisplayPickupListVouchersRequest `json:"ACSInputParameters"`
}

// DisplayPickupListVouchersRequest represents the structure of the request for displaying vouchers in a pickup list
type DisplayPickupListVouchersRequest struct {
	CompanyID       string  `json:"Company_ID"`
	CompanyPassword string  `json:"Company_Password"`
	UserID          string  `json:"User_ID"`
	UserPassword    string  `json:"User_Password"`
	Language        *string `json:"Language"`
	PickupListNo    string  `json:"PickupList_No"`
	PickupDate      string  `json:"Pickup_Date"`
}

// Validate validates the DisplayPickupListVouchersRequest
func (req *DisplayPickupListVouchersRequest) Validate() error {
	if req.PickupListNo == "" {
		return errors.New("PickupList_No is required")
	}
	if req.PickupDate == "" {
		return errors.New("Pickup_Date is required")
	}
	return nil
}

// DisplayPickupListVouchersResponse represents the structure of the response for displaying vouchers in a pickup list
type DisplayPickupListVouchersResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSValueOutput           []struct {
		ListVouchersCount int    `json:"List_Vouchers_Count"`
		ErrorMessage      string `json:"Error_Message"`
	} `json:"ACSValueOutput"`
	ACSTableOutput struct {
		TableData []struct {
			VoucherNo     string  `json:"Voucher_no"`
			ReferenceKey1 *string `json:"Reference_Key1"`
			ReferenceKey2 *string `json:"Reference_Key2"`
		} `json:"Table_Data"`
	} `json:"ACSTableOutput"`
}

// ACSGetPickupListsRequest defines the structure of the ACS request
type ACSGetPickupListsRequest struct {
	ACSAlias           string                 `json:"ACSAlias"`
	ACSInputParameters *GetPickupListsRequest `json:"ACSInputParameters"`
}

// GetPickupListsRequest represents the structure of the request for getting pickup lists
type GetPickupListsRequest struct {
	CompanyID       string  `json:"Company_ID"`
	CompanyPassword string  `json:"Company_Password"`
	UserID          string  `json:"User_ID"`
	UserPassword    string  `json:"User_Password"`
	Language        *string `json:"Language"`
	PickupDate      string  `json:"Pickup_Date"`
}

// Validate validates the GetPickupListsRequest
func (req *GetPickupListsRequest) Validate() error {
	if req.PickupDate == "" {
		return errors.New("Pickup_Date is required")
	}
	return nil
}

// GetPickupListsResponse represents the structure of the response for getting pickup lists
type GetPickupListsResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSValueOutput           []struct {
		ErrorMessage string `json:"Error_Message"`
	} `json:"ACSValueOutput"`
	ACSTableOutput struct {
		TableData []struct {
			PickupDate         string `json:"Pickup_date"`
			PickupListDateTime string `json:"Pickup_List_DateTime"`
			UserID             string `json:"User_ID"`
			PickupListNo       string `json:"PickupList_No"`
			ListVouchersCount  int    `json:"List_Vouchers_Count"`
		} `json:"Table_Data"`
	} `json:"ACSTableOutput"`
}

// ACSGetPickupListVouchersRequest defines the structure of the ACS request
type ACSTrackingSummaryRequest struct {
	ACSAlias           string                  `json:"ACSAlias"`
	ACSInputParameters *TrackingSummaryRequest `json:"ACSInputParameters"`
}

// TrackingSummaryRequest represents the request body for tracking summary
type TrackingSummaryRequest struct {
	CompanyID       string `json:"Company_ID"`
	CompanyPassword string `json:"Company_Password"`
	UserID          string `json:"User_ID"`
	UserPassword    string `json:"User_Password"`
	VoucherNo       string `json:"Voucher_No"`
}

// TrackingSummaryResponse represents the response from tracking summary
type TrackingSummaryResponse struct {
	ACSExecutionHasError     bool   `json:"ACSExecution_HasError"`
	ACSExecutionErrorMessage string `json:"ACSExecutionErrorMessage"`
	ACSOutputResponse        struct {
		TrackingDetails struct {
			// Add fields as necessary to match tracking details
			Status      string `json:"Status"`
			LastUpdated string `json:"Last_Updated"`
			// Add other fields as necessary
		} `json:"Tracking_Details"`
	} `json:"ACSOutputResponce"`
}
