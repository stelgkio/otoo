package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	domain "github.com/stelgkio/otoo/internal/core/domain"
)

// Define the ACS API URL for Multipart Vouchers and Printing
const (
	acsAPIURL               = "https://webservices.acscourier.net/ACSRestServices/api/ACSAutoRest"
	acsMultipartVouchersURL = "https://webservices.acscourier.net/ACSRestServices/api/ACSAutoRest"
	acsPrintVoucherURL      = "https://webservices.acscourier.net/ACSRestServices/api/ACSAutoRest"
	acsDeleteVoucherURL     = "https://webservices.acscourier.net/ACSRestServices/api/ACSAutoRest"
)

// ACSService defines the service structure for interacting with the ACS API
type ACSService struct {
	APIKey string
}

// NewACSService creates a new instance of ACSService
func NewACSService(apiKey string) *ACSService {
	return &ACSService{APIKey: apiKey}
}

// CreateVoucher sends the VoucherRequest to the ACS API and returns the voucher number
func (service *ACSService) CreateVoucher(request *domain.VoucherRequest) (string, error) {
	// Validate the request
	if err := request.Validate(); err != nil {
		return "", err
	}

	// Prepare the request body
	body := domain.ACSRequest{
		ACSAlias:           "ACS_Create_Voucher",
		ACSInputParameters: request,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.CreateVoucherResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ACSExecutionHasError {
		return "", errors.New(response.ACSExecutionErrorMessage)
	}
	// Check if there was an error in the execution
	if response.ACSExecutionHasError {
		domain.HandleAcsError(response.ACSExecutionErrorMessage)
		return "", errors.New(response.ACSExecutionErrorMessage)
	}

	// Return the voucher number if available
	if len(response.ACSOutputResponse.ACSValueOutput) > 0 {
		return response.ACSOutputResponse.ACSValueOutput[0].VoucherNo, nil
	}

	return "", errors.New("no voucher number returned from ACS")
}

// CreateMultipleVouchers handles creating and printing multiple vouchers
func (service *ACSService) CreateMultipleVouchers(request *domain.VoucherRequest) ([]string, error) {
	// Create the main voucher
	mainVoucherNo, err := service.CreateVoucher(request)
	if err != nil {
		return nil, fmt.Errorf("failed to create main voucher: %w", err)
	}

	// If Item_Quantity > 1, retrieve multipart vouchers
	if request.ItemQuantity > 1 {
		multiVoucherNos, err := service.GetMultipartVouchers(mainVoucherNo)
		if err != nil {
			return nil, fmt.Errorf("failed to get multipart vouchers: %w", err)
		}

		// Combine main voucher number with multipart voucher numbers
		voucherNos := append([]string{mainVoucherNo}, multiVoucherNos...)
		return voucherNos, nil
	}

	return []string{mainVoucherNo}, nil
}

// GetMultipartVouchers retrieves the associated sub-vouchers for a given main voucher number
func (service *ACSService) GetMultipartVouchers(mainVoucherNo string) ([]string, error) {
	// Prepare the request body
	body := domain.ACSMultipleRequest{
		ACSAlias: "ACS_Get_Multipart_Vouchers",
		ACSInputParameters: &domain.MultipartVoucherRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			Language:        "EN",
			MainVoucherNo:   mainVoucherNo,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsMultipartVouchersURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.MultipartVouchersResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if there was an error in the execution
	if response.ACSExecutionHasError {
		return nil, fmt.Errorf("error from ACS API: %s", response.ACSExecutionErrorMessage)
	}

	// Return the voucher numbers
	var voucherNos []string
	for _, entry := range response.ACSOutputResponse.ACSTableOutput.TableData {
		voucherNos = append(voucherNos, entry.MultiPartVoucherNo)
	}

	return voucherNos, nil
}

// PrintVouchers prints the vouchers using the Print Voucher API
func (service *ACSService) PrintVouchers(voucherNos []string, printType int, startPosition int, ItemQuantity int, WithReturnVoucher int) error {
	// Prepare the request body
	body := domain.ACSPrintRequest{
		ACSAlias: "ACS_Print_Voucher",
		ACSInputParameters: &domain.PrintVoucherRequest{
			CompanyID:         "demo",
			CompanyPassword:   "demo",
			UserID:            "demo",
			UserPassword:      "demo",
			VoucherNo:         voucherNos,
			PrintType:         printType,
			StartPosition:     startPosition,
			ItemQuantity:      ItemQuantity,
			WithReturnVoucher: WithReturnVoucher,
		},
	}
	if err := body.ACSInputParameters.Validate(); err != nil {
		return err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsPrintVoucherURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Process the response body (e.g., save PDFs or handle byte arrays)
	fmt.Printf("Print response: %s\n", bodyBytes)

	return nil
}

// DeleteVouchers handles the deletion of one or multiple vouchers
func (service *ACSService) DeleteVouchers(voucherNos []string) error {
	// Prepare the request body
	body := domain.ACSDeleteRequest{
		ACSAlias: "ACS_Delete_Voucher",
		ACSInputParameters: &domain.DeleteVoucherRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			VoucherNo:       voucherNos,
		},
	}
	if err := body.ACSInputParameters.Validate(); err != nil {
		return err
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsDeleteVoucherURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.DeleteVoucherResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if there was an error in the execution
	if response.ACSExecutionHasError {
		return fmt.Errorf("error from ACS API: %s", response.ACSExecutionErrorMessage)
	}

	fmt.Println("Delete response:", response)
	return nil
}

// IssuePickupList issues a new pickup list and returns the pickup list number
func (service *ACSService) IssuePickupList(pickupDate string) (string, error) {
	// Prepare the request body
	body := domain.ACSIssuePickupListRequest{
		ACSAlias: "ACS_Issue_Pickup_List",
		ACSInputParameters: &domain.IssuePickupListRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			PickupDate:      pickupDate,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.IssuePickupListResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ACSExecutionHasError {
		return "", errors.New(response.ACSExecutionErrorMessage)
	}

	if len(response.ACSValueOutput) > 0 {
		return response.ACSValueOutput[0].PickupListNo, nil
	}

	return "", errors.New("no pickup list number returned from ACS")
}

// PrintPickupList prints a pickup list given the pickup list number and date
func (service *ACSService) PrintPickupList(pickupListNo string, pickupDate string) error {
	// Prepare the request body
	body := domain.ACSPrintPickupListRequest{
		ACSAlias: "ACS_Print_Pickup_List",
		ACSInputParameters: &domain.PrintPickupListRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			MassNumber:      pickupListNo,
			PickupDate:      pickupDate,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Process the response body (e.g., save PDFs or handle byte arrays)
	fmt.Printf("Print pickup list response: %s\n", bodyBytes)

	return nil
}

// DisplayPickupListVouchers displays vouchers in a given pickup list
func (service *ACSService) DisplayPickupListVouchers(pickupListNo string, pickupDate string) (*domain.DisplayPickupListVouchersResponse, error) {
	// Prepare the request body
	body := domain.ACSDisplayPickupListVouchersRequest{
		ACSAlias: "ACS_Pickup_List_Display_Voucher",
		ACSInputParameters: &domain.DisplayPickupListVouchersRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			PickupListNo:    pickupListNo,
			PickupDate:      pickupDate,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.DisplayPickupListVouchersResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ACSExecutionHasError {
		return nil, errors.New(response.ACSExecutionErrorMessage)
	}

	return &response, nil
}

// GetPickupLists retrieves pickup lists for a given date
func (service *ACSService) GetPickupLists(pickupDate string) (*domain.GetPickupListsResponse, error) {
	// Prepare the request body
	body := domain.ACSGetPickupListsRequest{
		ACSAlias: "ACS_Get_Pickup_Lists",
		ACSInputParameters: &domain.GetPickupListsRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			PickupDate:      pickupDate,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.GetPickupListsResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ACSExecutionHasError {
		return nil, errors.New(response.ACSExecutionErrorMessage)
	}

	return &response, nil
}

// TrackingSummary retrieves tracking information for a given voucher number
func (service *ACSService) TrackingSummary(voucherNo string) (*domain.TrackingSummaryResponse, error) {
	// Prepare the request body
	body := domain.ACSTrackingSummaryRequest{
		ACSAlias: "ACS_Trackingsummary",
		ACSInputParameters: &domain.TrackingSummaryRequest{
			CompanyID:       "demo",
			CompanyPassword: "demo",
			UserID:          "demo",
			UserPassword:    "demo",
			VoucherNo:       voucherNo,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", acsAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACSApiKey", service.APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to ACS API: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var response domain.TrackingSummaryResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ACSExecutionHasError {
		return nil, errors.New(response.ACSExecutionErrorMessage)
	}

	return &response, nil
}
