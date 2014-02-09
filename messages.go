package twilio

import (
	"fmt"
	"log"
	"net/url"
)

// Messages is an intermediate class that will allow the creation of lists of
// recipients/messages
type Messages struct {
	client *TwilioRestClient
}

// Create a new SMS
func (m *Messages) Create(from string, to string, body string) {
	log.Printf("SMS %s -> %s: %s", from, to, body)

	apiURL := fmt.Sprintf("%s/Messages.json", m.client.getAPIBaseURL())
	values := url.Values{"Body": {body}, "From": {from}, "To": {to}}

	if _, err := m.client.post(apiURL, values); err != nil {
		log.Panic(err)
	}
}
