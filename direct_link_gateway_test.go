package ogone

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDirectLinkSend(t *testing.T) {
	a := NewAliasGateway()

	ar := NewAliasRequest(a.SandboxUrl())

	ar.SetAcceptUrl("http://testsuc.com")
	ar.SetExceptionUrl("http://tesat.com")

	order := "GOLANGTEST" + time.Now().Format("20060102150405_") + strconv.Itoa(rand.Intn(100000))

	ar.SetPspId("ewgraogone")
	ar.SetOrderId(order)
	ar.SetCardNumber("4111111111111111")
	ar.SetCardHolderName("Sökolov Evgenii")
	ar.SetCvc("123")
	ar.SetCardExpireMonth("01")
	ar.SetCardExpireYear(string(strconv.Itoa(time.Now().Year() + 1)))

	ar.Sign("qwdqwoidj29812d9")
	aresp, _ := a.Send(ar)

	dl := NewDirectLinkGateway()

	dlr := NewDirectLinkRequest(dl.SandboxUrl())

	dlr.SetPspId("ewgraogone")
	dlr.SetUserId("ewgragolang")
	dlr.SetPassword("123123aa")
	dlr.SetAlias(aresp.Alias())
	dlr.SetAmount("100")
	dlr.SetReserveOperation()
	dlr.SetCurrency("EUR")
	dlr.SetOrderId(order)

	dlr.Sign("qwdqwoidj29812d9")
	resp, _ := dl.Send(dlr)
	fmt.Println(resp)
}
