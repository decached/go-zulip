// zulip.go
// Copyright (C) 2016 akash <akash@decached.com>
//
// Distributed under terms of the MIT license.

package zulip

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.zulip.com/v1/"
)

type Client struct {
	BaseURL *url.URL
	Email   string
	APIKey  string
	client  *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	return c
}
