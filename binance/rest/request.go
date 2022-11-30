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

// Request define an API Request
type Request struct {
	method   string
	endpoint string
	query    url.Values
	form     url.Values
	secType  secType
	header   http.Header
	fullURL  string
}

// setParam set param with key/value to query string
func (r *Request) setParam(key string, value interface{}) *Request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set params with key/values to query string
func (r *Request) setParams(m params) *Request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

func (r *Request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return nil
}

// RequestOption define option type for request
type RequestOption func(*Request)

// WithHeader set or add a header value to the request
func WithHeader(key, value string, replace bool) RequestOption {
	return func(r *Request) {
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
