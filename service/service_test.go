package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

/*
Mock RoundTripper to intercept http.Get calls
*/
type mockRoundTripper struct {
	response string
	status   int
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: m.status,
		Body:       ioutil.NopCloser(bytes.NewBufferString(m.response)),
		Header:     make(http.Header),
	}, nil
}

func TestGetCountry_Success(t *testing.T) {

	mockJSON := `
	[
		{
			"name": {
				"common": "India"
			},
			"capital": ["New Delhi"],
			"currencies": {
				"INR": {
					"name": "Indian rupee"
				}
			},
			"population": 1400000000
		}
	]
	`

	// Save original transport
	originalTransport := http.DefaultTransport

	// Mock transport
	http.DefaultTransport = &mockRoundTripper{
		response: mockJSON,
		status:   http.StatusOK,
	}

	// Restore transport after test
	defer func() {
		http.DefaultTransport = originalTransport
	}()

	country, err := GetCountry("India")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if country.Name != "India" {
		t.Errorf("expected Name India, got %s", country.Name)
	}

	if country.Capital != "New Delhi" {
		t.Errorf("expected Capital New Delhi, got %s", country.Capital)
	}

	if country.Currency != "Indian rupee" {
		t.Errorf("expected Currency Indian rupee, got %s", country.Currency)
	}

	if country.Population != 1400000000 {
		t.Errorf("expected Population 1400000000, got %d", country.Population)
	}
}


func TestGetCountry_Non200Status(t *testing.T) {

	originalTransport := http.DefaultTransport

	http.DefaultTransport = &mockRoundTripper{
		response: `{}`,
		status:   http.StatusNotFound,
	}

	defer func() {
		http.DefaultTransport = originalTransport
	}()

	_, err := GetCountry("InvalidCountry")

	if err == nil {
		t.Fatalf("expected error for non-200 status code")
	}
}

func TestGetCountry_InvalidJSON(t *testing.T) {

	originalTransport := http.DefaultTransport

	http.DefaultTransport = &mockRoundTripper{
		response: `invalid-json`,
		status:   http.StatusOK,
	}

	defer func() {
		http.DefaultTransport = originalTransport
	}()

	_, err := GetCountry("India")

	if err == nil {
		t.Fatalf("expected JSON unmarshal error")
	}
}
