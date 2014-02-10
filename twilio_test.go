package twilio

import (
	"fmt"
	"strings"
	"testing"
)

// TestNewErrorResponse will test that the error response struct is created
// properly
func TestNewErrorResponse(t *testing.T) {
	expectedCode := 401

	responseBody := fmt.Sprintf("{\"Code\": %d}", expectedCode)
	errStruct := NewErrorResponse(strings.NewReader(responseBody))

	if errStruct.Code != 401 {
		t.Errorf("Expecting error code: '%d' and received '%d'", expectedCode, errStruct.Code)
	}
}
