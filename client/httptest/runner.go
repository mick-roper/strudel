package httptest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Request that describes a test
type Request struct {
	url     string
	method  string
	payload string
	Options *RequestOptions
}

// RequestOptions that help to configue a request
type RequestOptions struct {
	Headers               map[string]string
	ReturnResponseHeaders bool
}

// Response from an invoked test
type Response struct {
	StatusCode    int
	StatusMessage string
	Data          string
	ElapsedMs     int64
	Headers       map[string]string
}

// NewRequest creates a new request
func NewRequest(url, method, payload string, headers map[string]string) (*Request, error) {
	if url == "" {
		return nil, errors.New("url is a required parameter")
	}

	if method == "" {
		return nil, errors.New("method is a required parameter")
	}

	opts := &RequestOptions{Headers: headers}

	req := &Request{url, method, payload, opts}

	return req, nil
}

// CreateHTTPRequest from a test request
func (req *Request) CreateHTTPRequest() (*http.Request, error) {
	buffer := bytes.NewBufferString(req.payload)
	return http.NewRequest(req.method, req.url, buffer)
}

// TestRunner that runs tests
type TestRunner struct {
	client *http.Client
}

// NewTestRunner creates a new test runner
func NewTestRunner() *TestRunner {
	return &TestRunner{client: &http.Client{}}
}

// ExecuteRequest executes the test request and returns a test response
func (runner *TestRunner) ExecuteRequest(req *Request) (*Response, error) {
	if req == nil {
		return nil, errors.New("request must be provided")
	}

	httpRequest, err := req.CreateHTTPRequest()

	if err != nil {
		return nil, err
	}

	start := time.Now().UnixNano() * 1000000

	res, err := runner.client.Do(httpRequest)

	end := time.Now().UnixNano() * 1000000

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	respBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	testResponse := &Response{
		StatusCode:    res.StatusCode,
		StatusMessage: res.Status,
		ElapsedMs:     end - start,
		Data:          string(respBytes),
	}

	return testResponse, nil
}
