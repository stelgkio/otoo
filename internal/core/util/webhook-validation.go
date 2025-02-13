package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
)

func ExtractWebhookHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract headers
		deliveryId := c.Request().Header.Get("x-wc-webhook-delivery-id")
		webhookId := c.Request().Header.Get("x-wc-webhook-id")
		signature := c.Request().Header.Get("x-wc-webhook-signature")
		event := c.Request().Header.Get("x-wc-webhook-event")
		resource := c.Request().Header.Get("x-wc-webhook-resource")
		topic := c.Request().Header.Get("x-wc-webhook-topic")
		source := c.Request().Header.Get("x-wc-webhook-source")
		trimmedDomain := strings.TrimRight(source, "/")

		if deliveryId == "" || webhookId == "" || signature == "" || event == "" || resource == "" || topic == "" || source == "" {
			return c.String(http.StatusOK, "Missing required headers")
		}
		// Pass headers to the handler via context
		c.Set("webhookDeliveryId", deliveryId)
		c.Set("webhookId", webhookId)
		c.Set("webhookSignature", signature)
		c.Set("webhookEvent", event)
		c.Set("webhookResource", resource)
		c.Set("webhookTopic", topic)
		c.Set("webhookSource", trimmedDomain)

		return next(c)
	}
}

func ValidateWebhookSignature2(c echo.Context, secretKey string) error {
	// Extract the signature from the headers
	signature := c.Get("webhookSignature").(string)
	if signature == "" {
		return errors.New("missing signature")
	}

	// Read the body of the request
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	defer c.Request().Body.Close()

	// Compute the HMAC SHA256 hash
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write(body)
	expectedMAC := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// Compare the computed hash with the signature
	if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
		return errors.New("invalid signature")
	}

	return nil
}
func ValidateWebhookSignature(ctx echo.Context, secret string, body []byte) error {

	signature := ctx.Get("webhookSignature").(string)
	if signature == "" {
		return errors.New("missing signature")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	if signature != expectedSignature {
		return errors.New("invalid signature")
	}

	return nil
}

// AllErrorsEmpty checks if all elements in the webhooks slice have an empty Error field
func AllErrorsEmpty(webhooks []w.WebhookRecord) bool {
	var hasError = false
	for _, webhook := range webhooks {
		if webhook.Error != "" {
			hasError = true
		}
	}
	return hasError
}
