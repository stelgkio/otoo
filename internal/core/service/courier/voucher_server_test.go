package service_test

import (
	"testing"

	service "github.com/stelgkio/otoo/internal/core/service/courier"

	"github.com/stretchr/testify/assert"
)

func TestSplitFullName(t *testing.T) {
	tests := []struct {
		name              string
		fullName          string
		expectedFirstName string
		expectedLastName  string
	}{
		{
			name:              "Single part name",
			fullName:          "John",
			expectedFirstName: "John",
			expectedLastName:  "",
		},
		{
			name:              "Two part name",
			fullName:          "John Doe",
			expectedFirstName: "John",
			expectedLastName:  "Doe",
		},
		{
			name:              "Three part name",
			fullName:          "John Michael Doe",
			expectedFirstName: "John",
			expectedLastName:  "Michael Doe",
		},
		{
			name:              "Four part name",
			fullName:          "John Michael Smith Doe",
			expectedFirstName: "John",
			expectedLastName:  "Michael Smith Doe",
		},
		{
			name:              "Empty name",
			fullName:          "",
			expectedFirstName: "",
			expectedLastName:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstName, lastName := service.SplitFullName(tt.fullName)
			assert.Equal(t, tt.expectedFirstName, firstName)
			assert.Equal(t, tt.expectedLastName, lastName)
		})
	}
}
