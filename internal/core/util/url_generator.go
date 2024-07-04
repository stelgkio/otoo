package util

import (
	"fmt"
	"os"
)

func ResetPasswordLinkGenerator(textplan string) (string, error) {
	domain := os.Getenv("SITE_URL")
	hash, err := Encrypt(textplan)
	if err != nil {
		return "", err
	}
	if domain == "" && hash == "" {
		return "", nil
	}
	return fmt.Sprintf("%s/resetpassword/%s", domain, hash), nil
}
