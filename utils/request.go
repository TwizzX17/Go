package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

func GetCsvFromRemote(url string) (*csv.Reader, error) {
	// Define the request
	req, err := createRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Execute the request
	body, err := executeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	// Read the csv body
	reader, err := readCsvBody(body)
	if err != nil {
		return nil, fmt.Errorf("error reading csv body: %v", err)
	}

	return reader, nil
}

func createRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	return req, nil
}

// Execute the request and return body
func executeRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body before we close the resp body
	body, err := readBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func readBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return body, nil
}

func readCsvBody(body []byte) (*csv.Reader, error) {
	reader := csv.NewReader(io.NopCloser(bytes.NewReader(body)))

	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading csv header: %v", err)
	}

	return reader, nil
}
