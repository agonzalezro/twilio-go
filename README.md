Twilio NON-OFFICIAL Golang wrapper
==================================

**This project is in a really alpha state for now. I will try to keep it
up-to-date but I've just started it.**

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
