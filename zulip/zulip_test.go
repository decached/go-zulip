// zulip_test.go
// Copyright (C) 2016 akash <akash@decached.com>
//
// Distributed under terms of the MIT license.

package zulip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(nil, server.URL)
}

func teardown() {
	server.Close()
}
