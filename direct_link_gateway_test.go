package ogone

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDirectLinkSend(t *testing.T) {
	a := NewAliasGateway(&Config{
		pspId:        "ewgraogone",
		inPassPhrase: "qwdqwoidj29812d9",
	})

	ar := NewAliasRequest(a.SandboxUrl())

	ar.SetAcceptUrl("http://testsuc.com")
	ar.SetExceptionUrl("http://tesat.com")

	order := "GOLANGTEST" + time.Now().Format("20060102150405_") + strconv.Itoa(rand.Intn(100000))

	ar.SetOrderId(order)
	ar.SetCardNumber("4111111111111111")
	ar.SetCardHolderName("Sökolov Evgenii")
	ar.SetCvc("123")
	ar.SetCardExpireMonth("01")
	ar.SetCardExpireYear(string(strconv.Itoa(time.Now().Year() + 1)))

	aresp, _ := a.Send(ar)

	dl := NewDirectLinkGateway(&Config{
		pspId:        "ewgraogone",
		userId:       "ewgragolang",
		password:     "123123aa",
		inPassPhrase: "qwdqwoidj29812d9",
	})

	dlr := NewDirectLinkRequest(dl.SandboxUrl())

	dlr.SetAlias(aresp.Alias())
	dlr.SetAmount("100")
	dlr.SetReserveOperation()
	dlr.SetCurrency("EUR")
	dlr.SetOrderId(order)

	resp := dl.Send(dlr)
	fmt.Println(resp)
}
