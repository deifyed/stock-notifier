package config

import "net/url"

type StringGetter func(string) string

type Config struct {
	Symbols []string

	PushServerURL url.URL
}
