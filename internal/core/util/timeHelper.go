package util

import (
	"fmt"
	"time"
)

// CalculateTimeDifference calculates the time difference between two dates
func CalculateTimeDifference(startDateStr, endDateStr string) string {
	// Define the layout of the input date strings
	layout := "2006-01-02 15:04:05" // Layout for parsing date strings

	// Parse the start date string
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return ""
	}

	// Parse the end date string
	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return ""
	}

	// Calculate the duration between the two dates
	duration := endDate.Sub(startDate)

	// Extract days, hours, and minutes from the duration
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	if days == 0 && hours == 0 {
		return fmt.Sprintf(" %d minutes ago", minutes)
	} else if days == 0 {
		return fmt.Sprintf("%d hours, and %d minutes ago", hours, minutes)
	} else if days > 0 {
		return fmt.Sprintf("%d days, %d hours, and %d minutes ago", days, hours, minutes)
	} else {
		return "just now"
	}

}

// ConvertDateString  Function to convert and format the date
func ConvertDateString(dateStr string) (string, error) {
	// Define the layout for parsing the original string
	layoutInput := "2006-01-02T15:04:05"

	// Parse the input date string into a time.Time object
	parsedTime, err := time.Parse(layoutInput, dateStr)
	if err != nil {
		return "", err
	}

	// Change the date to 2024-09-25 while keeping the same time
	newDate := time.Date(2024, 9, 25, parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, parsedTime.Location())

	// Define the layout for formatting the output string
	layoutOutput := "2006-01-02 15:04:05"

	// Format the new date into the desired string format
	formattedTime := newDate.Format(layoutOutput)

	return formattedTime, nil
}
