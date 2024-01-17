package apirequest

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/v2"
	"golang.org/x/net/context/ctxhttp"
)

type RequestParams struct {
	URL      string
	Method   string
	Body     []byte
	Headers  map[string]string
	SpanName string
}

// MakeHTTPRequest is a generic function to make an HTTP request with APM instrumentation
func APIRequest(ctx context.Context, params RequestParams) (*http.Response, error) {

	span, _ := apm.StartSpan(ctx, params.SpanName, "custom")
	defer span.End()
	// Wrap http.Client with apmhttp.WrapClient
	client := apmhttp.WrapClient(&http.Client{})

	// Create an HTTP request
	req, err := http.NewRequest(params.Method, params.URL, bytes.NewBuffer(params.Body))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	// Make the HTTP request with APM instrumentation
	resp, err := ctxhttp.Do(ctx, client, req)
	if err != nil || resp.StatusCode != http.StatusOK {
		// Capture error and send to APM
		apm.CaptureError(ctx, fmt.Errorf("error in HTTP request. response: %s %v", resp.Status, err)).Send()
		return nil, fmt.Errorf("error in HTTP request. response: %s %v", resp.Status, err)
	}

	return resp, nil
}
