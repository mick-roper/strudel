package httptest

import "errors"

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
	ElapsedMs     float64
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

// RunTest executes the tet request and returns a test response
func RunTest(req *Request) (*Response, error) {
	return nil, errors.New("not implemented")
}
