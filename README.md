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
