package notification

type Message struct {
	Title    string
	Message  string
	Priority int
}

type Client interface {
	Notify(Message) error
}
