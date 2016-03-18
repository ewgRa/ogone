package ogone

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDirectLink(t *testing.T) {
	orderID := "GOLANGTEST" + time.Now().Format("20060102150405_") + strconv.Itoa(rand.Intn(100000))

	aResp := NewTestAliasResponse(t, orderID)

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
