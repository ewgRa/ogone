package ogone

import (
	"testing"
)

func TestShaInAliasCompose(t *testing.T) {
	expected := "F4CC376CD7A834D997B91598FA747825A238BE0A"

	params := map[string]string{
		"lANGUAGE":    "en_US",
		"CARDNO":      "DO_NOT_INCLUDE_ME",
		"NOT_ALLOWED": "foo",
		"AMOUnT":      "1500",
		"CURReNCY":    "EUR",
		"ORDERID":     "1234",
		"PSPID":       "MyPSPID",
	}

	sign := shaInAliasCompose(params, "Mysecretsig1875!?")

	if sign != expected {
		t.Fatalf("ShaIn compose failed, expected %s, got %s", expected, sign)
	}
}

func TestShaInCompose(t *testing.T) {
	expected := "F4CC376CD7A834D997B91598FA747825A238BE0A"

	params := map[string]string{
		"lANGUAGE":    "en_US",
		"NOT_ALLOWED": "foo",
		"AMOUnT":      "1500",
		"CURReNCY":    "EUR",
		"ORDERID":     "1234",
		"PSPID":       "MyPSPID",
	}

	sign := shaInCompose(params, "Mysecretsig1875!?")

	if sign != expected {
		t.Fatalf("ShaIn compose failed, expected %s, got %s", expected, sign)
	}
}

func TestShaOutCompose(t *testing.T) {
	expected := "209113288F93A9AB8E474EA78D899AFDBB874355"

	params := map[string]string{
		"amount":      "15",
		"BRAND":       "VISA",
		"ACCEPTANCE":  "1234",
		"CARDNO":      "XXXXXXXXXXXX1111",
		"NOT_ALLOWED": "foo",
		"currency":    "EUR",
		"NCERROR":     "0",
		"orderID":     "12",
		"PAYID":       "32100123",
		"PM":          "CreditCard",
		"STATUS":      "9",
	}

	sign := shaOutCompose(params, "Mysecretsig1875!?")

	if sign != expected {
		t.Fatalf("ShaOut compose failed, expected %s, got %s", expected, sign)
	}
}

func BenchmarkShaInCompose(b *testing.B) {
	params := map[string]string{
		"lANGUAGE":    "en_US",
		"NOT_ALLOWED": "foo",
		"AMOUnT":      "1500",
		"CURReNCY":    "EUR",
		"ORDERID":     "1234",
		"PSPID":       "MyPSPID",
	}

	for n := 0; n < b.N; n++ {
		shaInCompose(params, "Mysecretsig1875!?")
	}
}
