// This example is going to make a call using the Twilio API.
//
// You will need a public endpoint serving an XML file with the TwiML language.
// You will also need to set the environment variables:
//     TWILIO_ACCOUNT
//     TWILIO_TOKEN
//     TWILIO_NUMBER
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/agonzalezro/twilio-go"
)

var (
	to  string
	url string
)

func init() {
	flag.StringVar(&to, "to", "", "the recipient number")
	flag.StringVar(&url, "url", "", "the URL where Twilio is going to POST to")
}

func main() {
	flag.Parse()

	account := os.Getenv("TWILIO_ACCOUNT")
	token := os.Getenv("TWILIO_TOKEN")
	from := os.Getenv("TWILIO_NUMBER")

	if account == "" || token == "" || from == "" || to == "" || url == "" {
		fmt.Println("All the options are mandatory! Check the help with -h.")
		fmt.Println("Also, you must set the env variables: TWILIO_ACCOUNT, TWILIO_TOKEN & TWILIO_NUMBER.")
		os.Exit(1)
	}

	client := twilio.NewTwilioRestClient(account, token)
	client.Calls.Create(from, to, url)
}
