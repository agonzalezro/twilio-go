package twilio

import (
	"fmt"
	"strings"
	"testing"
)

// TestMessagesResponseOk will test that giving the proper response we will
// create an struct from it
func TestMessagesResponseOk(t *testing.T) {
	messages := Messages{}

	expected := "expected"
	responseBody := fmt.Sprintf("{\"sid\": \"%s\"}", expected)
	responseStruct, _ := messages.getResponseStruct(strings.NewReader(responseBody))

	if responseStruct.Sid != expected {
		t.Errorf("Expecting the sid: '%s' and received '%s'", expected, responseStruct.Sid)
	}

}

// TestMessageResponseErr will test that if an error occurs creating the struct
// we will be aware because it's going to be returned
func TestMessageResponseErr(t *testing.T) {
	messages := Messages{}

	_, err := messages.getResponseStruct(strings.NewReader(""))
	if err == nil {
		t.Error("Expecting an error non received")
	}
}
