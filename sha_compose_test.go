package ogone

import (
	"testing"
)

func TestShaCompose(t *testing.T) {
	expected := "F4CC376CD7A834D997B91598FA747825A238BE0A"

	params := map[string]string{
		"lANGUAGE": "en_US",
		"AMOUnT": "1500",
		"CURReNCY": "EUR",
		"ORDERID": "1234",
		"PSPID": "MyPSPID",
	}

	sign := ShaCompose(params, "Mysecretsig1875!?")

	if (sign != expected) {
		t.Fatalf("Sha compose failed, expected %s, got %s", expected, sign)
	}
}