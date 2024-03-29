package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type URL struct {
	Scheme string
	Path   string
}

func parseURL(textURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(textURL)
	if err != nil {
		return nil, err
	}

	return parsedURL, nil
}

func (u *URL) UnmarshalJSON(input []byte) error {
	var textURL string

	err := json.Unmarshal(input, &textURL)
	if err != nil {
		return err
	}

	parsedURL, err := parseURL(textURL)
	if err != nil {
		return err
	}

	u.Scheme = parsedURL.Scheme
	u.Path = parsedURL.Path

	return nil
}

func main() {
	// Example JSON data containing a URL
	jsonData := []byte(`"https://example.com/path"`)

	var u URL
	err := u.UnmarshalJSON(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parsed URL:")
	fmt.Println("Scheme:", u)
	fmt.Println("Path:", u.Path)
}
