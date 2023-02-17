package main

import (
	"fmt"
	"net/url"
	"os/exec"
)

func main() {
	message := "Hola! Te quiero mucho y no puedo esperar a verte en México!"
	phone := "+521XXXXXXXXXX" // Replace XXXXXXXXXX with your nephew's phone number including the country code

	// Set up the API endpoint and parameters
	apiEndpoint := "https://api.whatsapp.com/send"
	params := url.Values{}
	params.Add("phone", phone)
	params.Add("text", message)

	// Create the full URL with the API endpoint and parameters
	fullURL := fmt.Sprintf("%s?%s", apiEndpoint, params.Encode())

	// Open the URL in the user's default web browser
	err := openURL(fullURL)
	if err != nil {
		fmt.Println("Error opening URL:", err)
	}
}

// Helper function to open a URL in the user's default web browser
func openURL(url string) error {
	return exec.Command("xdg-open", url).Start()
}
