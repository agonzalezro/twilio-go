package twilio

import (
	"fmt"
	"strings"
	"testing"
)

// TestCallsResponseOk will test that giving the proper response we will
// create an struct from it
func TestCallsResponseOk(t *testing.T) {
	calls := Calls{}

	expected := "expected"
	responseBody := fmt.Sprintf("{\"sid\": \"%s\"}", expected)
	responseStruct, _ := calls.getResponseStruct(strings.NewReader(responseBody))

	if responseStruct.Sid != expected {
		t.Errorf("Expecting the sid: '%s' and received '%s'", expected, responseStruct.Sid)
	}

}

// TestCallsResponseErr will test that if an error occurs creating the struct
// we will be aware because it's going to be returned
func TestCallsResponseErr(t *testing.T) {
	calls := Calls{}

	_, err := calls.getResponseStruct(strings.NewReader(""))
	if err == nil {
		t.Error("Expecting an error non received")
	}
}
