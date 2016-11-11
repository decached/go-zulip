// messages_test.go
// Copyright (C) 2016 akash <akash@decached.com>
//
// Distributed under terms of the MIT license.

package zulip

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMessageService_Send(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"msg": "", "result": "success", "id": 123456789}`)
	})

	got, err := client.Messages.Send("type_", "to", "subject", "content")
	if err != nil {
		t.Errorf("Messages.Send returned error: %v", err)
	}

	want := int64(123456789)
	if got != want {
		t.Errorf("Message.Send returned %+v, want %+v", got, want)
	}
}
