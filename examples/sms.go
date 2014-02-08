package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/agonzalezro/twilio-go"
)

var (
	account string
	body    string
	from    string
	to      string
	token   string
)

func init() {
	flag.StringVar(&account, "account", "", "the Twilio account")
	flag.StringVar(&token, "token", "", "the Twilio secret token")
	flag.StringVar(&from, "from", "", "the Twilio number")
	flag.StringVar(&to, "to", "", "the recipient number")
	flag.StringVar(&body, "body", "", "the text on the message")
}

func main() {
	flag.Parse()
	if account == "" || token == "" || from == "" || to == "" || body == "" {
		fmt.Println("All the options are mandatory! Check the help with -h.")
		os.Exit(1)
	}

	client := twilio.NewTwilioRestClient(account, token)
	client.Messages.Create(from, to, body)
}
