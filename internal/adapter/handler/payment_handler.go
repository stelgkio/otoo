package handler

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/auth"
	woo "github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/sub"
	"github.com/stripe/stripe-go/v72/webhook"
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
		Discounts:           discounts,
		AllowPromotionCodes: stripe.Bool(true),
		Mode:                stripe.String(string(stripe.CheckoutSessionModeSubscription)),
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

func getEvent(c echo.Context) (*stripe.Event, error) {
	const MaxBodyBytes = int64(65536)
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return nil, err
	}

	// Get the Stripe-Signature header value
	sigHeader := c.Request().Header.Get("Stripe-Signature")
	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	// Verify the webhook signature
	event, err := webhook.ConstructEvent(payload, sigHeader, webhookSecret)
	if err != nil {
		return nil, fmt.Errorf("Failed to verify webhook signature: %s", err)
	}

	return &event, nil
}

// PaymentEvent handles the Stripe webhook event
func (dh *DashboardHandler) PaymentEvent(c echo.Context) error {
	// Retrieve the Stripe event
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	event, err := getEvent(c)
	if err != nil {
		slog.Error("Failed to get event:", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid event payload"})
	}

	slog.Info("Received event", "event: ", event.Type)

	// Handle the event based on its type
	if event.Type == "customer.subscription.created" {
		slog.Info("Unhandled event type", "eventType", event.Type)
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

		slog.Info("Subscription created metadata - ProjectId", "projectID", projectID)
		slog.Info("Subscription created metadata - ExtensionID", "extensionID", extensionID)

	}
	if event.Type == "customer.subscription.deleted" {
		slog.Info("Unhandled event type", "eventType", event.Type)
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

		slog.Info("Subscription deleted metadata - ProjectId", "projectID", projectID)
		slog.Info("Subscription deleted metadata - ExtensionID", "extensionID", extensionID)

		err = dh.SubscriptionDeleted(c, subscriptionID, projectID, extensionID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete subscription"})
		}
	}
	if event.Type == "invoice.payment_succeeded" {
		// This event is triggered when a subscription payment succeeds
		slog.Info("Unhandled event type", "eventType", event.Type)
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
		slog.Info("Payment succeeded", "subscriptionID", subscriptionID)

		// You can also access metadata from the invoice or subscription object if needed
		projectID := subscription.Metadata["ProjectId"]
		extensionID := subscription.Metadata["ExtensionID"]
		slog.Info("Payment succeeded for project", "projectID", projectID)
		slog.Info("Payment succeeded  - ExtensionID", "extensionID", extensionID)
		err = dh.SubscriptionCreated(c, subscriptionID, projectID, extensionID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create subscription"})
		}
		projectExtensions, err := dh.extensionSvc.GetProjectExtensionByID(nil, extensionID, projectID)
		if err != nil {
			return err
		}

		now := time.Now().UTC()

		nextpay := now.Add(time.Duration(projectExtensions.SubscriptionPeriod) * 24 * time.Hour)
		extensions, err := dh.extensionSvc.GetExtensionByID(nil, extensionID)
		if err != nil {
			return err
		}
		payment := woo.NewPaymentSuccess(projectID, projectExtensions.ID.Hex(), projectExtensions.Code, int64(extensions.Price), nextpay)
		dh.paymentSvc.CreatePayment(nil, payment)

	}
	if event.Type == "invoice.payment_failed" {
		// This event is triggered when a subscription payment succeeds
		slog.Info("Unhandled event type", "eventType", event.Type)
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
		slog.Info("Payment failed", "subscriptionID", subscriptionID)

		// You can also access metadata from the invoice or subscription object if needed
		projectID := subscription.Metadata["ProjectId"]
		extensionID := subscription.Metadata["ExtensionID"]
		slog.Info("Payment failed for project", "projectID", projectID)
		slog.Info("Payment failed - ExtensionID", "extensionID", extensionID)

		projectExtensions, err := dh.extensionSvc.GetProjectExtensionByID(nil, extensionID, projectID)
		if projectExtensions != nil {

			if err != nil {
				return err
			}
			extensions, err := dh.extensionSvc.GetExtensionByID(nil, extensionID)
			if err != nil {
				return err
			}
			now := time.Now().UTC()

			payment := woo.NewPaymentFail(projectID, projectExtensions.ID.Hex(), projectExtensions.Code, int64(extensions.Price), now)
			dh.paymentSvc.CreatePayment(nil, payment)
			err = dh.SubscriptionDeleted(c, subscriptionID, projectID, extensionID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed payment delete subscription"})
			}
		}
	}

	return c.NoContent(http.StatusOK)
}

// SubscriptionCreated asd
func (dh *DashboardHandler) SubscriptionCreated(ctx echo.Context, subscriptionID, projectID, extensionID string) error {
	extension, err := dh.extensionSvc.GetExtensionByID(ctx, extensionID)
	if err != nil {
		return err
	}

	if extension.Code == "asc-courier" {
		if err := dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, subscriptionID); err != nil {
			return err
		}
	}
	if extension.Code == "courier4u" {
		if err := dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, subscriptionID); err != nil {
			return err
		}
	}
	if extension.Code == "redcourier" {
		if err := dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, subscriptionID); err != nil {
			return err
		}
	}
	if extension.Code == "wallet-expences" {
		if err := dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, subscriptionID); err != nil {
			return err
		}
	}
	if extension.Code == "team-member" {
		if err := dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, subscriptionID); err != nil {
			return err
		}
	}
	if extension.Code == "data-synchronizer" {
		if err := dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 365, subscriptionID); err != nil {
			return err
		}
	}

	return nil
}

// SubscriptionDeleted delete
func (dh *DashboardHandler) SubscriptionDeleted(ctx echo.Context, subscriptionID, projectID, extensionID string) error {
	err := dh.extensionSvc.DeleteProjectExtension(ctx, extensionID, projectID)
	if err != nil {
		return nil
	}

	return nil
}

// PaymentTable return payment
func (dh *DashboardHandler) PaymentTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	pageNum, err := strconv.Atoi(page)
	sort := ctx.QueryParam("sort")
	direction := ctx.QueryParam("direction")
	itemsPerPage := 12
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid page number: %v", err))
	}

	if sort == "" {
		sort = "extension_name"
	}
	if direction == "" {
		direction = "asc"
	}
	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	paymentCountChan := make(chan int64, 1)
	paymentListChan := make(chan []*woo.Payment, 1)
	errChan := make(chan error, 1)
	errListChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		dh.paymentSvc.GetPaymentCountAsync(ctx, projectID, paymentCountChan, errChan)
	}()

	// Fetch  10 orders
	go func() {
		defer wg.Done()
		dh.paymentSvc.FindPaymentByProjectIDAsync(ctx, projectID, itemsPerPage, pageNum, sort, direction, paymentListChan, errListChan)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(paymentCountChan)
		close(paymentListChan)
		close(errChan)
		close(errListChan)
	}()

	var totalItems int64
	var paymentRecord []*woo.Payment

	for err := range errChan {
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch order count": err.Error()})
		}
	}
	for errList := range errListChan {
		if errList != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error to fetch order": errList.Error()})
		}
	}

	for item := range paymentCountChan {
		totalItems = item
	}
	for item := range paymentListChan {
		paymentRecord = item
	}

	// Convert orderRecords to OrderTableList for the response
	var payments []woo.PaymentTableList
	if paymentRecord != nil {
		for _, record := range paymentRecord {
			payments = append(payments, woo.PaymentTableList{
				ID:            record.ID,
				ProjectID:     record.ProjectID,
				ExtensionName: record.ExtensionName,
				Amount:        float64(record.Amount),
				IsPaid:        record.IsPaid,
				IsFail:        record.IsFail,
				CreatedAt:     record.CreatedAt,
			})
		}
	}

	// Prepare metadata

	totalPages := int(totalItems) / itemsPerPage
	if int(totalItems)%itemsPerPage > 0 {
		totalPages++
	}

	// Create response object
	response := woo.PaymentTableResponde{
		Data: payments,
		Meta: woo.Meta{
			TotalItems:   int(totalItems),
			CurrentPage:  pageNum,
			ItemsPerPage: itemsPerPage,
			TotalPages:   totalPages,
		},
	}

	return ctx.JSON(http.StatusOK, response)

}
