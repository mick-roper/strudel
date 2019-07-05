package engine

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// TestRunner is responsible for running tests
type TestRunner struct {
	client *http.Client
}

// NewTestRunner creates a new test runner
func NewTestRunner() *TestRunner {
	client := &http.Client{}
	return &TestRunner{client}
}

// Execute a test definition
func (runner *TestRunner) Execute(def *TestDefinition) (*TestResult, error) {
	if def != nil {
		return nil, errors.New("no definition")
	}

	req, err := def.CreateHTTPRequest()

	if err != nil {
		return nil, err
	}

	start := time.Now().UnixNano() * 1000000

	resp, err := runner.client.Do(req)

	end := time.Now().UnixNano() * 1000000
	elapsedMs := end - start

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body := ioutil.ReadAll(resp.Body)
	payload := string(body)
	result := &TestResult{
		StatusCode:    resp.StatusCode,
		StatusMessage: resp.Status,
		Payload:       payload,
		ElapsedMs:     elapsedMs,
	}

	return result, nil
}
