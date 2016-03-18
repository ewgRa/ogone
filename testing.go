package ogone

import (
	"time"
	"strconv"
	"testing"
)

// NewTestAliasResponse perform alias request and return response. Use it when you need alias
func NewTestAliasResponse(orderID string, t *testing.T) *AliasResponse {
	ar := NewAliasRequest()

	ar.
	SetAcceptURL("https://github.com/ewgRa/ogone/success").
	SetExceptionURL("https://github.com/ewgRa/ogone/exception").
	SetOrderID(orderID).
	SetCardNumber("4111111111111111").
	SetCardHolderName("Sökolov Evgenii").
	SetCardCVC("123").
	SetCardExpireMonth("01").
	SetCardExpireYear(string(strconv.Itoa(time.Now().Year() + 1)))

	ag := NewAliasGateway()

	ar.
	SetPspID("ewgraogone").
	Sign("qwdqwoidj29812d9")

	aResp, _ := ag.SandboxSend(ar)

	if !aResp.CheckSign("qwdqwoidj29812d9") {
		t.Fatalf("Check Alias response signature failed")
	}

	if !aResp.IsOk() {
		t.Fatalf("Status is not okey at Alias response")
	}

	return aResp
}