package twilio

import (
	"fmt"
	"log"
	"net/url"
)

type Messages struct {
	client *TwilioRestClient
}

// Send a new SMS
func (m *Messages) Create(from string, to string, body string) {
	log.Printf("SMS %s -> %s: %s", from, to, body)

	apiUrl := fmt.Sprintf("%s/Messages.json", m.client.getAPIBaseURL())
	values := url.Values{"Body": {body}, "From": {from}, "To": {to}}

	if _, err := m.client.post(apiUrl, values); err != nil {
		log.Panic(err)
	}
}
