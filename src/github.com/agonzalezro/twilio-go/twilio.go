package twilio

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TwilioRestClient struct {
	account  string
	token    string
	Messages *Messages
}

func NewTwilioRestClient(account string, token string) *TwilioRestClient {
	client := &TwilioRestClient{account, token, &Messages{}}
	client.Messages.client = client // Keep track of the parent
	return client
}

// Return the base URL for the Twilio account
func (c TwilioRestClient) getAPIBaseURL() string {
	version := "2010-04-01"
	return fmt.Sprintf("https://api.twilio.com/%s/Accounts/%s", version, c.account)
}

// Do a POST request to Twilio setting the needed headers.
func (c TwilioRestClient) post(methodUrl string, values url.Values) (*http.Response, error) {
	httpClient := &http.Client{}

	req, _ := http.NewRequest("POST", methodUrl, strings.NewReader(values.Encode()))
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
