package notification

import (
	"net/url"
	"strconv"
)

func (m Message) AsValues() url.Values {
	return url.Values{
		"title":    {m.Title},
		"message":  {m.Message},
		"priority": {strconv.Itoa(m.Priority)},
	}
}
