package engine

import (
	"bytes"
	"net/http"
)

// TestDefinition that the test runner will execute
type TestDefinition struct {
	url     string
	method  string
	payload string
}

// CreateHTTPRequest creates a HTTP request from a test definition
func (def *TestDefinition) CreateHTTPRequest() (*http.Request, error) {
	data := bytes.NewBufferString(def.payload)
	return http.NewRequest(def.method, def.url, data)
}

// TestResult of a running TestDefinition
type TestResult struct {
	StatusCode    int
	StatusMessage string
	Payload       string
	ElapsedMs     int
}
