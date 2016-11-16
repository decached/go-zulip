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

func TestEventsService_Register(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"msg": "", "result": "success", "last_event_id": -1, "queue_id": "1375801870:2942"}`)
	})

	gotQID, gotEID, err := client.Events.Register([]string{"message", "subscriptions"}, false)
	if err != nil {
		t.Errorf("Events.Register returned error: %v", err)
	}

	wantQID, wantEID := "1375801870:2942", int64(-1)
	if (gotQID != wantQID) || (gotEID != wantEID) {
		t.Errorf("Events.Register returned %+v, %+v, want %+v, %+v", gotQID, gotEID, wantQID, wantEID)
	}
}
