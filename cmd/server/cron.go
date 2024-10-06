package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

// InitCronScheduler initializes and starts the cron scheduler
func InitCronScheduler() *cron.Cron {

	// Create a new cron instance
	c := cron.New()

	// Start the cron scheduler
	c.Start()
	// Add a cron job that runs every 10 seconds
	// _, err := c.AddFunc("@every 10s", RunAnalyticsJob)
	// _, err = c.AddFunc("@every 10s", RunAProductBestSellerJob)
	_, err := c.AddFunc("0 4 * * 1", RunOrderWeeklyBalanceJob)
	// Add a cron job that runs every day at 4 AM
	_, err = c.AddFunc("0 4 * * *", RunAProductBestSellerDailyJob)

	if err != nil {
		fmt.Println("Error reading the response body:", err)
	}
	fmt.Println("Cron scheduler initialized")

	return c
}

// RunAnalyticsJob is the function to be executed by the cron job
func RunAnalyticsJob() {
	domain := os.Getenv("SITE_URL")
	resp, err := http.Get(fmt.Sprintf("%s/RunAnalyticsJob", domain))
	if err != nil {
		fmt.Println("Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

}

// RunAProductBestSellerDailyJob is the function to be executed by the cron job
func RunAProductBestSellerDailyJob() {
	domain := os.Getenv("SITE_URL")
	resp, err := http.Get(fmt.Sprintf("%s/RunAProductBestSellerDailyJob", domain))
	if err != nil {
		fmt.Println("Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

}

// RunAProductBestSellerInitializerJob is the function to be executed by the cron job
func RunAProductBestSellerInitializerJob() {
	domain := os.Getenv("SITE_URL")
	resp, err := http.Get(fmt.Sprintf("%s/RunAProductBestSellerInitializerJob", domain))
	if err != nil {
		fmt.Println("Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

}

// RunCustomerBestBuyerJob is the function to be executed by the cron job
func RunCustomerBestBuyerJob() {
	domain := os.Getenv("SITE_URL")
	resp, err := http.Get(fmt.Sprintf("%s/RunCustomerBestBuyerJob", domain))
	if err != nil {
		fmt.Println("Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	//fmt.Println(string(body))
}

// RunOrderWeeklyBalanceJob is the function to be executed by the cron job
func RunOrderWeeklyBalanceJob() {
	domain := os.Getenv("SITE_URL")
	resp, err := http.Get(fmt.Sprintf("%s/RunOrderWeeklyBalanceJob", domain))
	if err != nil {
		fmt.Println("Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}
}
