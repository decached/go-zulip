// zulip.go
// Copyright (C) 2016 akash <akash@decached.com>
//
// Distributed under terms of the MIT license.

package zulip

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Client struct {
	client   *http.Client
	sling    *sling.Sling
	Email    string
	APIKey   string
	common   service
	Messages *MessagesService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	s := sling.New().Client(httpClient).Base(baseURL)
	c := &Client{sling: s}

	c.common.client = c
	c.Messages = (*MessagesService)(&c.common)

	return c
}
