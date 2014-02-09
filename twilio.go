// Package twilio is s small wrapper to be used with the Twilio API. Twilio
// lets you use standard web languages to build voice, VoIP and SMS
// applications via a web API.
//
// To use it, all that you need is to create a client:
//
//     client := NewTwilioRestClient("yourAccount", "yourToken")
//
// And to send a message just write:
//
//     client.Messages.Create("fromNumber", "toNumber", "body")
package twilio

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// TwilioRestClient is an struct to store your Twilio credentials and available
// methods
type TwilioRestClient struct {
	account  string
	token    string
	Calls    *Calls
	Messages *Messages
}

// NewTwilioRestClient is going to create a new Twilio client providing the
// account and token details
func NewTwilioRestClient(account string, token string) *TwilioRestClient {
	client := &TwilioRestClient{account, token, &Calls{}, &Messages{}}
	client.Messages.client = client // Keep track of the parent
	client.Calls.client = client
	return client
}

// Return the base URL for the Twilio account
func (c TwilioRestClient) getAPIBaseURL() string {
	version := "2010-04-01"
	return fmt.Sprintf("https://api.twilio.com/%s/Accounts/%s", version, c.account)
}

// Do a POST request to Twilio setting the needed headers.
func (c TwilioRestClient) post(methodURL string, values url.Values) (*http.Response, error) {
	httpClient := &http.Client{}

	req, _ := http.NewRequest("POST", methodURL, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.account, c.token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBodyContent, _ := ioutil.ReadAll(resp.Body)
	log.Printf("POST response body: %s", responseBodyContent)

	return resp, nil
}
