# go-zulip
Go library to access the [Zulip API][zulip-api]

## Usage

```go
import "github.com/decached/go-zulip/zulip"
```

Construct a new Zulip client, then use the various services on the client to
access different parts of the GitHub API. For example:

```go
c := zulip.New(nil, "https://api.zulip.com/v1/")
```

#### Messages

```go
// Send a message

type_ := "stream"
to := "Denmark"
subject := "Castle"
content := "Something is rotten in the state of Denmark."

c.Messages.Send(type_, to, subject, content)
```

#### Events

```go
// Register a queue to receive new messages

eventTypes := []string{"message", "subscriptions"}
applyMarkdown := true

c.Events.Register(eventTypes, applyMarkdown)
```

[zulip-api]: https://zulip.com/api/endpoints/
