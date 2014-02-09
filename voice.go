package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

// Calls is an intermediate class that for now is here just to emulate the
// Python API behaviour
type Calls struct {
	client *TwilioRestClient
}

// CallsResponse is the struct where we will Unmarshal the API response
type CallsResponse struct {
	Sid             string
	DateCreated     string `json:"date_created"`
	DateUpdated     string `json:"date_updated"`
	ParentCallSid   string `json:"parent_call_sid"`
	AccountSid      string `json:"account_sid"`
	To              string
	ToFormated      string `json:"to_formatted"`
	From            string
	FromFormatted   string `json:"from_formatted"`
	PhoneNumberSid  string `json:"phone_number_sid"`
	Status          string
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	Duration        string
	Price           string
	PriceUnit       string `json:"price_unit"`
	Direction       string
	AnsweredBy      string `json:"answered_by"`
	APIVersion      string `json:"api_version"`
	Annotation      string
	ForwardedFrom   string `json:"forwaded_from"`
	GroupSid        string `json:"group_sid"`
	CallerName      string `json:"caller_name"`
	Uri             string
	SubresourceUris struct {
		Notifications string
		Recordings    string
	} `json:"subresource_uris"`
}

func (c *Calls) getResponseStruct(body io.Reader) (*CallsResponse, error) {
	callsResponse := CallsResponse{}
	bodyContent, _ := ioutil.ReadAll(body)
	if err := json.Unmarshal(bodyContent, &callsResponse); err != nil {
		return nil, err
	}
	return &callsResponse, nil
}

// Create a new call. You will need an external XML endpoint with the TwiML XML
func (c *Calls) Create(from string, to string, URL string) (*CallsResponse, *ErrorResponse) {
	apiURL := fmt.Sprintf("%s/Calls.json", c.client.getAPIBaseURL())
	values := url.Values{"Url": {URL}, "From": {from}, "To": {to}}

	response, err := c.client.post(apiURL, values)
	if err != nil {
		log.Panic(err) // Not optimal, will see in the future
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		callsResponse, err := c.getResponseStruct(response.Body)
		if err != nil {
			log.Print(err)
		}
		return callsResponse, nil
	}

	return nil, NewErrorResponse(response.Body)
}
