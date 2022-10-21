package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPResponse200 return http response with status code 200
func HTTPResponse200(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}

// HTTPResponse201 return http response with status code 201
func HTTPResponse201(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "201 Created",
		StatusCode:    201,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}

// HTTPResponse400 return http response with status code 400
func HTTPResponse400(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "400 Bad Request",
		StatusCode:    400,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}

// HTTPResponse402 return http response with status code 402
func HTTPResponse402(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "402 Payment Required",
		StatusCode:    402,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}

// HTTPResponse403 return http response with status code 403
func HTTPResponse403(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "403 Forbidden",
		StatusCode:    403,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}

// HTTPResponse422 return http response with status code 422
func HTTPResponse422(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "422 Unprocessable Entity",
		StatusCode:    422,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}

// HTTPResponse500 return http response with status code 500
func HTTPResponse500(body string, req *http.Request) *http.Response {
	return &http.Response{
		Status:        "500 Internal server error",
		StatusCode:    500,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
}
