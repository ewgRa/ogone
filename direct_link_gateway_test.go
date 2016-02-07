package ogone

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDirectLink(t *testing.T) {
	ar := NewAliasRequest()

	orderID := "GOLANGTEST" + time.Now().Format("20060102150405_") + strconv.Itoa(rand.Intn(100000))

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

	dlr := NewDirectLinkRequest()

	dlr.
		SetAlias(aResp.Alias()).
		SetAmount("100").
		SetReserveOperation().
		SetCurrency("EUR").
		SetOrderID(orderID)

	dlg := NewDirectLinkGateway()

	dlr.
		SetPspID("ewgraogone").
		SetUserID("ewgragolang").
		SetPassword("123123aa").
		Sign("qwdqwoidj29812d9")

	dlResp, _ := dlg.SandboxSend(dlr)

	if !dlResp.IsAuthorised() {
		t.Fatalf("Status is not authorized at DirectLink response")
	}
}
