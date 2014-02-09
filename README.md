Twilio NON-OFFICIAL Golang wrapper
==================================

**This project is in a really alpha state for now. Feel free of provide pull requests or add [new issues](https://github.com/agonzalezro/twilio-go/issues).**

[![GoDoc](https://godoc.org/github.com/agonzalezro/twilio-go?status.png)](https://godoc.org/github.com/agonzalezro/twilio-go)

How to use?
-----------

On the import section of your Go project:

    import (
        ...
        "github.com/agonzalezro/twilio-go"
        ...
    )

Now, if you want to send an SMS just do:

    client = twilio.NewTwilioRestClient("Your account here", "Your token here")
    client.Messsages.Create("+44fromnumber", "+31tonumber", "Hi world!")

If you want to check more than this tiny example, you must check the [generated
documentation at godoc](https://godoc.org/github.com/agonzalezro/twilio-go).

Supported methods
-----------------

- You can send SMS.
- You can do calls.

How to develop easily?
----------------------

What I do is create a symbolic link of my code (if you know a better way,
please let me know :)

If your go path is in your folder ~/go, and you cloned the repo to ~/twilio-go,
you can just copy paste this commands:

    $ export GOPATH=/Users/alex/go
    $ mkdir -p $GOPATH/src/github.com/agonzalezro
    $ ln -s $HOME/twilio-go $GOPATH/src/github.com/agonzalezro

Examples
--------

I've included few examples on the `examples/` folder.

For example, if you want to send an sms, you can use this command (filling the
needed information, of course):

	go run examples/sms/sms.go \
		-account XXXXYOURACCOUNTHEREXXXX \
		-token YYYYOURTOKENHEREYYYY \
		-from +15005550006 \
		-to +447449601002 \
		-body "Hi world\!"

If you want to make a call, you will need an adittional step. You should check
the documentation of [TwiML](https://www.twilio.com/docs/api/twiml) and provide
the XML in an endpoint that is accessible from the world.

After having that, you can use this command (note that I am using environment
variables now):

	TWILIO_ACCOUNT=AC0f9f0c5025b665ed8d68e649e442b0ff \
	TWILIO_TOKEN=77a081faa8bebb1487827b527f87a0e5 \
	TWILIO_NUMBER=+15005550006 \
	go run examples/voice/voice.go \
		-to 447449601002 \
		-url http://YOURENDPOINTXML/MUSTALLOWPOST
