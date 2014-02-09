package twilio

import (
	"fmt"
	"log"
	"net/url"
)

// Calls is an intermediate class that for now is here just to emulate the
// Python API behaviour
type Calls struct {
	client *TwilioRestClient
}

// Create a new call. You will need an external XML endpoint with the TwiML XML
func (c *Calls) Create(from string, to string, URL string) {
	log.Printf("Call %s -> %s: %s", from, to, URL)

	apiURL := fmt.Sprintf("%s/Calls.json", c.client.getAPIBaseURL())
	values := url.Values{"Url": {URL}, "From": {from}, "To": {to}}

	if _, err := c.client.post(apiURL, values); err != nil {
		log.Panic(err)
	}
}
