// Code generated from specification version 6.7.2: DO NOT EDIT

package esapi

import (
	"context"
	"strings"
)

func newClusterRemoteInfoFunc(t Transport) ClusterRemoteInfo {
	return func(o ...func(*ClusterRemoteInfoRequest)) (*Response, error) {
		var r = ClusterRemoteInfoRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// ClusterRemoteInfo returns the information about configured remote clusters.
//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-remote-info.html.
//
type ClusterRemoteInfo func(o ...func(*ClusterRemoteInfoRequest)) (*Response, error)

// ClusterRemoteInfoRequest configures the Cluster  Remote Info API request.
//
type ClusterRemoteInfoRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ClusterRemoteInfoRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_remote/info"))
	path.WriteString("/_remote/info")

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
func (f ClusterRemoteInfo) WithContext(v context.Context) func(*ClusterRemoteInfoRequest) {
	return func(r *ClusterRemoteInfoRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ClusterRemoteInfo) WithPretty() func(*ClusterRemoteInfoRequest) {
	return func(r *ClusterRemoteInfoRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ClusterRemoteInfo) WithHuman() func(*ClusterRemoteInfoRequest) {
	return func(r *ClusterRemoteInfoRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ClusterRemoteInfo) WithErrorTrace() func(*ClusterRemoteInfoRequest) {
	return func(r *ClusterRemoteInfoRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ClusterRemoteInfo) WithFilterPath(v ...string) func(*ClusterRemoteInfoRequest) {
	return func(r *ClusterRemoteInfoRequest) {
		r.FilterPath = v
	}
}
