package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// encrypt encrypts plain text string using AES algorithm and returns the encrypted text in hexadecimal format.
func encrypt(plainText, key string) (string, error) {
	// Create a new AES cipher using the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Make the plain text into a byte slice
	plainTextBytes := []byte(plainText)

	// Generate a new AES-GCM cipher based on the AES block cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce of the appropriate length
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the plain text using AES-GCM
	cipherText := aesGCM.Seal(nonce, nonce, plainTextBytes, nil)

	// Return the encrypted text in hexadecimal format
	return hex.EncodeToString(cipherText), nil
}

// decrypt decrypts the encrypted text in hexadecimal format using AES algorithm and returns the decrypted plain text.
func decrypt(cipherText, key string) (string, error) {
	// Convert the encrypted text from hexadecimal format to a byte slice
	cipherTextBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher using the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Generate a new AES-GCM cipher based on the AES block cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the encrypted text
	nonceSize := aesGCM.NonceSize()
	nonce, cipherTextBytes := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]

	// Decrypt the encrypted text using AES-GCM
	plainTextBytes, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	// Return the decrypted plain text
	return string(plainTextBytes), nil
}
