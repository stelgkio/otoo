package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	courier_domain "github.com/stelgkio/otoo/internal/core/domain/courier"
	"github.com/stelgkio/otoo/internal/core/port"
)

// Define the ACS API URL for Multipart Vouchers and Printing
const (
	redCourierURL = "https://web.redcourier.gr"
	courier4uURL  = "https://web.courier4u.gr"
)

// HermesService defines the methods for interacting with the Voucher service
type HermesService struct {
	repo port.VoucherRepository
}

// NewHermesService creates a new voucher service instance
func NewHermesService(repo port.VoucherRepository) *HermesService {
	return &HermesService{
		repo,
	}
}

// CreateVoucher inserts a new Voucher into the database
func (vs *HermesService) CreateVoucher(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, hermesVoucerRequest *courier_domain.HermesVoucerRequest, projectID string) (*courier_domain.VoucherResponse, error) {

	url := ""
	token := ""
	if courier4u == nil && redcourier == nil {
		return nil, fmt.Errorf("internal server error")
	}
	if courier4u != nil {
		url = courier4uURL + "/api/v5.0/CreateVoucher"
		token = courier4u.CourierAPIKey
	}
	if redcourier != nil {
		url = redCourierURL + "/api/v5.0/CreateVoucher"
		token = redcourier.CourierAPIKey
	}

	// Encode the struct to JSON
	jsonBody, err := json.Marshal(hermesVoucerRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Initialize HTTP client
	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	// Print the response body

	// Decode the response JSON into the VoucherResponse struct
	var voucherResponse *courier_domain.VoucherResponse
	err = json.Unmarshal(body, &voucherResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}
	// Print the structured response

	return voucherResponse, nil
}

// PrintVoucher inserts a new Voucher into the database
func (vs *HermesService) PrintVoucher(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, voucherId int64, projectID, printType string) ([]byte, error) {

	url := ""
	token := ""
	if courier4u == nil && redcourier == nil {
		return nil, fmt.Errorf("internal server error")
	}
	if courier4u != nil {
		url = courier4uURL + "/api/v5.0/PrintVouchers"
		token = courier4u.CourierAPIKey
	}
	if redcourier != nil {
		url = redCourierURL + "/api/v5.0/PrintVouchers"
		token = redcourier.CourierAPIKey
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	print := "a4"
	if printType == "thermal" {
		print = "a6"
	} else {
		print = "a4"
	}
	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Set query parameters
	q := req.URL.Query()
	q.Add("type", print)
	q.Add("vouchers", fmt.Sprintf("%d", voucherId))
	req.URL.RawQuery = q.Encode()

	// Initialize HTTP client
	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// Check if the response is a success
	if res.StatusCode == http.StatusOK {
		// Read the PDF from the response body
		pdfData, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		return pdfData, nil
	}

	// Handle potential error response
	var errorResponse *courier_domain.VoucherPrintResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// You can parse the error response if needed
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return nil, fmt.Errorf("failed to print vouchers: %s", errorResponse.Message)
}

// PrintMultipleVoucher inserts a new Voucher into the database
func (vs *HermesService) PrintMultipleVoucher(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, voucherId []int64, projectID, printType string) ([]byte, error) {

	url := ""
	token := ""
	if courier4u == nil && redcourier == nil {
		return nil, fmt.Errorf("internal server error")
	}
	if courier4u != nil {
		url = courier4uURL + "/api/v5.0/PrintVouchers"
		token = courier4u.CourierAPIKey
	}
	if redcourier != nil {
		url = redCourierURL + "/api/v5.0/PrintVouchers"
		token = redcourier.CourierAPIKey
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	print := "a4"
	if printType == "thermal" {
		print = "a6"
	} else {
		print = "a4"
	}
	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Set query parameters
	q := req.URL.Query()
	q.Add("type", print)
	// Convert voucherId slice to a comma-separated string
	voucherIDStr := make([]string, len(voucherId))
	for i, id := range voucherId {
		voucherIDStr[i] = fmt.Sprintf("%d", id)
	}
	q.Add("vouchers", strings.Join(voucherIDStr, ","))
	req.URL.RawQuery = q.Encode()

	// Initialize HTTP client
	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// Check if the response is a success
	if res.StatusCode == http.StatusOK {
		// Read the PDF from the response body
		pdfData, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		return pdfData, nil
	}

	// Handle potential error response
	var errorResponse *courier_domain.VoucherPrintResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// You can parse the error response if needed
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return nil, fmt.Errorf("failed to print vouchers: %s", errorResponse.Message)
}

// UpdateVoucher update a Voucher
func (vs *HermesService) UpdateVoucher(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, hermesVoucerRequest *courier_domain.HermesVoucerUpdateRequest, projectID string) (*courier_domain.VoucherResponse, error) {

	url := ""
	token := ""
	if courier4u == nil && redcourier == nil {
		return nil, fmt.Errorf("internal server error")
	}
	if courier4u != nil {
		url = courier4uURL + "/api/v5.0/EditVoucher"
		token = courier4u.CourierAPIKey
	}
	if redcourier != nil {
		url = redCourierURL + "/api/v5.0/EditVoucher"
		token = redcourier.CourierAPIKey
	}

	// Encode the struct to JSON
	jsonBody, err := json.Marshal(hermesVoucerRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Initialize HTTP client
	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	// Print the response body

	// Decode the response JSON into the VoucherResponse struct
	var voucherResponse *courier_domain.VoucherResponse
	err = json.Unmarshal(body, &voucherResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}
	// Print the structured response

	return voucherResponse, nil
}

// TrackingHermerVoucherStatus get the status of a voucher
func (vs *HermesService) TrackingHermerVoucherStatus(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, voucherId int64) (*courier_domain.TrackingResponse, error) {

	url := ""
	token := ""
	if courier4u == nil && redcourier == nil {
		return nil, fmt.Errorf("internal server error")
	}
	if courier4u != nil {
		url = courier4uURL + "/api/v5.0/GetVoucherFullTrackingStatus"
		token = courier4u.CourierAPIKey
	}
	if redcourier != nil {
		url = redCourierURL + "/api/v5.0/GetVoucherFullTrackingStatus"
		token = redcourier.CourierAPIKey
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Set query parameters
	q := req.URL.Query()

	q.Add("voucher", fmt.Sprintf("%d", voucherId))
	req.URL.RawQuery = q.Encode()

	// Initialize HTTP client
	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	// Print the response body
	//fmt.Println("Response:", string(body))
	// Decode the response JSON into the VoucherResponse struct
	var trackingResponse *courier_domain.TrackingResponse
	err = json.Unmarshal(body, &trackingResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}
	// Print the structured response
	//fmt.Printf("Response Struct: %+v\n", trackingResponse)
	return trackingResponse, nil
}
