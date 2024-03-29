// Code generated from specification version 6.7.2: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"strconv"
	"strings"
	"time"
)

func newSnapshotCreateRepositoryFunc(t Transport) SnapshotCreateRepository {
	return func(repository string, body io.Reader, o ...func(*SnapshotCreateRepositoryRequest)) (*Response, error) {
		var r = SnapshotCreateRepositoryRequest{Repository: repository, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SnapshotCreateRepository creates a repository.
//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html.
//
type SnapshotCreateRepository func(repository string, body io.Reader, o ...func(*SnapshotCreateRepositoryRequest)) (*Response, error)

// SnapshotCreateRepositoryRequest configures the Snapshot  Create Repository API request.
//
type SnapshotCreateRepositoryRequest struct {
	Body io.Reader

	Repository string

	MasterTimeout time.Duration
	Timeout       time.Duration
	Verify        *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SnapshotCreateRepositoryRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_snapshot") + 1 + len(r.Repository))
	path.WriteString("/")
	path.WriteString("_snapshot")
	path.WriteString("/")
	path.WriteString(r.Repository)

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Verify != nil {
		params["verify"] = strconv.FormatBool(*r.Verify)
	}

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

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f SnapshotCreateRepository) WithContext(v context.Context) func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
//
func (f SnapshotCreateRepository) WithMasterTimeout(v time.Duration) func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f SnapshotCreateRepository) WithTimeout(v time.Duration) func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.Timeout = v
	}
}

// WithVerify - whether to verify the repository after creation.
//
func (f SnapshotCreateRepository) WithVerify(v bool) func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.Verify = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SnapshotCreateRepository) WithPretty() func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SnapshotCreateRepository) WithHuman() func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SnapshotCreateRepository) WithErrorTrace() func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SnapshotCreateRepository) WithFilterPath(v ...string) func(*SnapshotCreateRepositoryRequest) {
	return func(r *SnapshotCreateRepositoryRequest) {
		r.FilterPath = v
	}
}
