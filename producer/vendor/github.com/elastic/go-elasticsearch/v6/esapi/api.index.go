// Code generated from specification version 6.7.2: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"strconv"
	"strings"
	"time"
)

func newIndexFunc(t Transport) Index {
	return func(index string, body io.Reader, o ...func(*IndexRequest)) (*Response, error) {
		var r = IndexRequest{Index: index, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// Index creates or updates a document in an index.
//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html.
//
type Index func(index string, body io.Reader, o ...func(*IndexRequest)) (*Response, error)

// IndexRequest configures the Index API request.
//
type IndexRequest struct {
	Index        string
	DocumentType string
	DocumentID   string

	Body io.Reader

	IfPrimaryTerm       *int
	IfSeqNo             *int
	OpType              string
	Parent              string
	Pipeline            string
	Refresh             string
	Routing             string
	Timeout             time.Duration
	Version             *int
	VersionType         string
	WaitForActiveShards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IndexRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	if r.DocumentID != "" {
		method = "PUT"
	} else {
		method = "POST"
	}

	if r.DocumentType == "" {
		r.DocumentType = "_doc"
	}

	path.Grow(1 + len(r.Index) + 1 + len(r.DocumentType) + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString(r.Index)
	path.WriteString("/")
	path.WriteString(r.DocumentType)
	if r.DocumentID != "" {
		path.WriteString("/")
		path.WriteString(r.DocumentID)
	}

	params = make(map[string]string)

	if r.IfPrimaryTerm != nil {
		params["if_primary_term"] = strconv.FormatInt(int64(*r.IfPrimaryTerm), 10)
	}

	if r.IfSeqNo != nil {
		params["if_seq_no"] = strconv.FormatInt(int64(*r.IfSeqNo), 10)
	}

	if r.OpType != "" {
		params["op_type"] = r.OpType
	}

	if r.Parent != "" {
		params["parent"] = r.Parent
	}

	if r.Pipeline != "" {
		params["pipeline"] = r.Pipeline
	}

	if r.Refresh != "" {
		params["refresh"] = r.Refresh
	}

	if r.Routing != "" {
		params["routing"] = r.Routing
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Version != nil {
		params["version"] = strconv.FormatInt(int64(*r.Version), 10)
	}

	if r.VersionType != "" {
		params["version_type"] = r.VersionType
	}

	if r.WaitForActiveShards != "" {
		params["wait_for_active_shards"] = r.WaitForActiveShards
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
func (f Index) WithContext(v context.Context) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.ctx = v
	}
}

// WithDocumentID - document ID.
//
func (f Index) WithDocumentID(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.DocumentID = v
	}
}

// WithDocumentType - the type of the document.
//
func (f Index) WithDocumentType(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.DocumentType = v
	}
}

// WithIfPrimaryTerm - only perform the index operation if the last operation that has changed the document has the specified primary term.
//
func (f Index) WithIfPrimaryTerm(v int) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.IfPrimaryTerm = &v
	}
}

// WithIfSeqNo - only perform the index operation if the last operation that has changed the document has the specified sequence number.
//
func (f Index) WithIfSeqNo(v int) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.IfSeqNo = &v
	}
}

// WithOpType - explicit operation type.
//
func (f Index) WithOpType(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.OpType = v
	}
}

// WithParent - ID of the parent document.
//
func (f Index) WithParent(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Parent = v
	}
}

// WithPipeline - the pipeline ID to preprocess incoming documents with.
//
func (f Index) WithPipeline(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Pipeline = v
	}
}

// WithRefresh - if `true` then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` (the default) then do nothing with refreshes..
//
func (f Index) WithRefresh(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Refresh = v
	}
}

// WithRouting - specific routing value.
//
func (f Index) WithRouting(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Routing = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f Index) WithTimeout(v time.Duration) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Timeout = v
	}
}

// WithVersion - explicit version number for concurrency control.
//
func (f Index) WithVersion(v int) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Version = &v
	}
}

// WithVersionType - specific version type.
//
func (f Index) WithVersionType(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.VersionType = v
	}
}

// WithWaitForActiveShards - sets the number of shard copies that must be active before proceeding with the index operation. defaults to 1, meaning the primary shard only. set to `all` for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
//
func (f Index) WithWaitForActiveShards(v string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f Index) WithPretty() func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f Index) WithHuman() func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f Index) WithErrorTrace() func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f Index) WithFilterPath(v ...string) func(*IndexRequest) {
	return func(r *IndexRequest) {
		r.FilterPath = v
	}
}
