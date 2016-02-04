package ogone

import (
	"testing"
	"fmt"
)

func TestDirectLinkSend(t *testing.T) {
	dl := NewDirectLink(&Config{
		pspId: "ewgraogone",
		userId: "ewgragolang",
		password: "123123aa",
		inPassPhrase: "qwdqwoidj29812d9",
		sandbox: true,
	})

	dlr := NewDirectLinkRequest()

	dlr.SetAlias("test")
	dlr.SetAmount("100")
	dlr.SetReserveOperation()
	dlr.SetCurrency("EUR")
	dlr.SetOrderId("TESTORDER")

	resp := dl.Send(dlr)

	fmt.Println(resp)
}