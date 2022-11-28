package rest

import (
	"fmt"
	"net/http"
	"net/url"
)

type secType int32

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned // if the 'timestamp' parameter is required
)

type params map[string]interface{}

// request define an API request
type request struct {
	method   string
	endpoint string
	query    url.Values
	form     url.Values
	secType  secType
	header   http.Header
	fullURL  string
}

// setParam set param with key/value to query string
func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set params with key/values to query string
func (r *request) setParams(m params) *request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return nil
}

// RequestOption define option type for request
type RequestOption func(*request)

// WithHeader set or add a header value to the request
func WithHeader(key, value string, replace bool) RequestOption {
	return func(r *request) {
		if r.header == nil {
			r.header = http.Header{}
		}
		if replace {
			r.header.Set(key, value)
		} else {
			r.header.Add(key, value)
		}
	}
}
