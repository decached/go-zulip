// messages.go
// Copyright (C) 2016 akash <akash@decached.com>
//
// Distributed under terms of the MIT license.

package zulip

type MessagesService service

type MessageResponse struct {
	msg    string
	result string
	ID     int64
}

type MessageRequest struct {
	Type    string `json:"type"`
	To      string `json:"to"`
	Subject string `json:"subject,omitempty"`
	Content string `json:"content"`
}

func (s *MessagesService) Send(type_ string, to string, subject string, content string) (ID int64, err error) {
	body := &MessageRequest{
		Type:    type_,
		To:      to,
		Subject: subject,
		Content: content,
	}
	mResponse := new(MessageResponse)

	_, err = s.client.sling.New().Post("messages").BodyJSON(body).Receive(mResponse, nil)
	if err != nil {
		return 0, err
	}

	return mResponse.ID, err
}
