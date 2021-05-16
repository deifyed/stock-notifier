package gotify

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/deifyed/stock-notifier/pkg/notification"
)

type client struct {
	serverURL url.URL
}

func (c client) Notify(message notification.Message) error {
	_, err := http.PostForm(c.serverURL.String(), message.AsValues())
	if err != nil {
		return fmt.Errorf("posting notification: %w", err)
	}

	return nil
}

func NewGotifyClient(serverURL url.URL) notification.Client {
	return &client{serverURL: serverURL}
}
