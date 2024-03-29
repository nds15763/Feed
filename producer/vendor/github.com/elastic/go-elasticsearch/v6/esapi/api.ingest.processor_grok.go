// Code generated from specification version 6.7.2: DO NOT EDIT

package esapi

import (
	"context"
	"strings"
)

func newIngestProcessorGrokFunc(t Transport) IngestProcessorGrok {
	return func(o ...func(*IngestProcessorGrokRequest)) (*Response, error) {
		var r = IngestProcessorGrokRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IngestProcessorGrok returns a list of the built-in patterns.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/grok-processor.html#grok-processor-rest-get.
//
type IngestProcessorGrok func(o ...func(*IngestProcessorGrokRequest)) (*Response, error)

// IngestProcessorGrokRequest configures the Ingest  Processor Grok API request.
//
type IngestProcessorGrokRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IngestProcessorGrokRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_ingest/processor/grok"))
	path.WriteString("/_ingest/processor/grok")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f IngestProcessorGrok) WithContext(v context.Context) func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IngestProcessorGrok) WithPretty() func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IngestProcessorGrok) WithHuman() func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IngestProcessorGrok) WithErrorTrace() func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IngestProcessorGrok) WithFilterPath(v ...string) func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.FilterPath = v
	}
}
