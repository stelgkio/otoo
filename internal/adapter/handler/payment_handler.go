package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/sub"
)

// Payment get extention
func (dh *DashboardHandler) Payment(c echo.Context) error {

	req := new(CheckoutInput)
	if err := c.Bind(req); err != nil {
		log.Println("Failed to bind input:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	userID, err := auth.GetUserID(c)
	if err != nil {
		return err
	}
	user, err := dh.userSvc.GetUserById(c, userID)
	if err != nil {
		return err
	}
	extem, err := dh.extensionSvc.GetExtensionByID(c, req.ExtensionID)
	if err != nil {
		return err
	}

	priceID := extem.PriceID
	stripeSession, err := checkout(user.Email, req.ProjectID, req.ExtensionID, priceID)
	if err != nil {
		log.Println("Failed to create Stripe session:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create checkout session"})
	}

	return c.JSON(http.StatusOK, &SessionOutput{ID: stripeSession.ID})

}

func checkout(email, projectID, extensionID, priceID string) (*stripe.CheckoutSession, error) {
	var discounts []*stripe.CheckoutSessionDiscountParams
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Initialize variable to store the customer ID
	var customerID string

	// Attempt to find an existing customer with the provided email
	params := &stripe.CustomerListParams{
		Email: stripe.String(email),
	}
	i := customer.List(params)

	// Iterate through the list to find the customer (if it exists)
	for i.Next() {
		c := i.Customer()
		if c.Email == email {
			customerID = c.ID
			log.Printf("Found existing customer: %s", customerID)
			break
		}
	}

	if customerID == "" {
		// If no customer exists with this email, create a new one
		customerParams := &stripe.CustomerParams{
			Email: stripe.String(email),
		}
		customerParams.AddMetadata("FinalEmail", email)
		customerParams.AddMetadata("ProjectId", projectID)
		newCustomer, err := customer.New(customerParams)
		if err != nil {
			return nil, fmt.Errorf("failed to create new customer: %w", err)
		}
		customerID = newCustomer.ID
		log.Printf("Created new customer: %s", customerID)
	}

	// Metadata to be added to the subscription
	meta := map[string]string{
		"FinalEmail":  email,
		"ProjectId":   projectID,
		"ExtensionID": extensionID,
	}

	log.Println("Creating meta for user: ", meta)

	// Create the Stripe Checkout session
	paramsCheckout := &stripe.CheckoutSessionParams{
		Customer:   stripe.String(customerID),
		SuccessURL: stripe.String(fmt.Sprintf("%s/extension/%s/%s/success", os.Getenv("DELIVERY_URL"), projectID, extensionID)),
		CancelURL:  stripe.String(fmt.Sprintf("%s/extension/%s/%s/fail", os.Getenv("DELIVERY_URL"), projectID, extensionID)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Discounts: discounts,
		Mode:      stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Metadata: meta,
		},
	}
	return session.New(paramsCheckout)
}

// CheckoutInput is the input for the email
type CheckoutInput struct {
	ProjectID   string `json:"projectId"`
	ExtensionID string `json:"extensionId"`
}

// SessionOutput is the input for the email
type SessionOutput struct {
	ID string `json:"id"`
}

// MaxBodyBytes is the maximum size of the request body
const MaxBodyBytes = int64(65536)

func getEvent(c echo.Context) (*stripe.Event, error) {
	// Limit the size of the request body
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, MaxBodyBytes)

	// Read the request body
	payload, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the payload into a Stripe Event
	var event stripe.Event
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}

	return &event, nil
}

// PaymentEvent handles the Stripe webhook event
func (dh *DashboardHandler) PaymentEvent(c echo.Context) error {
	// Retrieve the Stripe event
	event, err := getEvent(c)
	if err != nil {
		slog.Error("Failed to get event:", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid event payload"})
	}

	slog.Info("Received event", "event: ", event.Type)

	// Handle the event based on its type
	if event.Type == "customer.subscription.created" {
		// Access the customer ID and subscription ID from the event data
		customerID := event.Data.Object["customer"].(string)
		subscriptionID := event.Data.Object["id"].(string)

		// Retrieve the customer details from Stripe
		cust, err := customer.Get(customerID, nil)
		if err != nil {
			slog.Error("Failed to retrieve customer", "error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve customer"})
		}

		// Retrieve the subscription details from Stripe
		subscription, err := sub.Get(subscriptionID, nil)
		if err != nil {
			slog.Error("Failed to retrieve subscription", "error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve subscription"})
		}

		// Access metadata from the customer object
		email := cust.Metadata["FinalEmail"]
		slog.Info("Customer metadata - Subscription created by", "email", email)

		// Access metadata from the subscription object
		projectID := subscription.Metadata["ProjectId"]
		extensionID := subscription.Metadata["ExtensionID"]

		slog.Info("Subscription metadata - ProjectId", "projectID", projectID)
		slog.Info("Subscription metadata - ExtensionID", "extensionID", extensionID)
	}

	if event.Type == "invoice.payment_succeeded" {
		// This event is triggered when a subscription payment succeeds

		// Access the customer ID and invoice ID from the event data
		customerID := event.Data.Object["customer"].(string)
		invoiceID := event.Data.Object["id"].(string)

		// Retrieve the customer details from Stripe
		cust, err := customer.Get(customerID, nil)

		if err != nil {
			slog.Error("Failed to retrieve customer", "error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve customer"})
		}
		// Access metadata from the customer object
		email := cust.Metadata["FinalEmail"]
		slog.Info("Customer metadata - Subscription payment by", "email", email)

		// Retrieve the invoice details from Stripe
		invoice, err := invoice.Get(invoiceID, nil)
		if err != nil {
			slog.Error("Failed to retrieve invoice", "error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve invoice"})
		}

		// Access the subscription ID related to the payment
		subscriptionID := invoice.Subscription.ID

		// Retrieve subscription details from Stripe
		subscription, err := sub.Get(subscriptionID, nil)
		if err != nil {
			slog.Error("Failed to retrieve subscription", "error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve subscription"})
		}

		// Log the payment success for the subscription
		slog.Info("Subscription payment succeeded", "subscriptionID", subscriptionID)

		// You can also access metadata from the invoice or subscription object if needed
		projectID := subscription.Metadata["ProjectId"]
		extensionID := subscription.Metadata["ExtensionID"]
		slog.Info("Subscription payment for project", "projectID", projectID)
		slog.Info("Subscription metadata - ExtensionID", "extensionID", extensionID)

		// Additional logic to update your system for successful payment
		// (e.g., mark invoice as paid in your database)

	}

	if event.Type == "invoice.payment_failed" {
		slog.Warn("Unhandled event type", "eventType", event.Type)
	}

	return c.NoContent(http.StatusOK)
}
