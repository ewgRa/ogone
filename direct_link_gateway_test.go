package ogone

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDirectLink(t *testing.T) {
	ar := NewAliasRequest()

	orderId := "GOLANGTEST" + time.Now().Format("20060102150405_") + strconv.Itoa(rand.Intn(100000))

	ar.
		SetAcceptUrl("https://github.com/ewgRa/ogone/success").
		SetExceptionUrl("https://github.com/ewgRa/ogone/exception").
		SetOrderId(orderId).
		SetCardNumber("4111111111111111").
		SetCardHolderName("Sökolov Evgenii").
		SetCvc("123").
		SetCardExpireMonth("01").
		SetCardExpireYear(string(strconv.Itoa(time.Now().Year() + 1)))

	ag := NewAliasGateway()

	ar.
		SetPspId("ewgraogone").
		Sign("qwdqwoidj29812d9")

	aResp, _ := ag.SandboxSend(ar)

	if !aResp.CheckSign("qwdqwoidj29812d9") {
		t.Fatalf("Check Alias response signature failed")
	}

	// FIXME XXX: assert response success

	dlr := NewDirectLinkRequest()

	dlr.
		SetAlias(aResp.Alias()).
		SetAmount("100").
		SetReserveOperation().
		SetCurrency("EUR").
		SetOrderId(orderId)

	dlg := NewDirectLinkGateway()

	dlr.
		SetPspId("ewgraogone").
		SetUserId("ewgragolang").
		SetPassword("123123aa").
		Sign("qwdqwoidj29812d9")

	dlResp, _ := dlg.SandboxSend(dlr)

	// FIXME XXX: check status
	fmt.Println(dlResp)
}
