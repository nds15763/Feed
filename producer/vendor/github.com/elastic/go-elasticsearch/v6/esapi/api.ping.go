// Code generated from specification version 6.7.2: DO NOT EDIT

package esapi

import (
	"context"
	"strings"
)

func newPingFunc(t Transport) Ping {
	return func(o ...func(*PingRequest)) (*Response, error) {
		var r = PingRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// Ping returns whether the cluster is running.
//
// See full documentation at http://www.elastic.co/guide/.
//
type Ping func(o ...func(*PingRequest)) (*Response, error)

// PingRequest configures the Ping API request.
//
type PingRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r PingRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "HEAD"

	path.Grow(len("/"))
	path.WriteString("/")

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
func (f Ping) WithContext(v context.Context) func(*PingRequest) {
	return func(r *PingRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f Ping) WithPretty() func(*PingRequest) {
	return func(r *PingRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f Ping) WithHuman() func(*PingRequest) {
	return func(r *PingRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f Ping) WithErrorTrace() func(*PingRequest) {
	return func(r *PingRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f Ping) WithFilterPath(v ...string) func(*PingRequest) {
	return func(r *PingRequest) {
		r.FilterPath = v
	}
}
