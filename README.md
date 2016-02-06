# Golang library for Ogone integration

This library is an [Ingenico Payment Services](https://payment-services.ingenico.com/int/en) integration library for the
[Go](http://www.golang.org/) programming language.

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ewgra/ogone/master/LICENSE)
[![GoReportCard](http://goreportcard.com/badge/ewgra/ogone)](http://goreportcard.com/report/ewgra/ogone)

## Releases

Library version | Package URL
----------------------|------------------
1.x                   | [`gopkg.in/ewgra/ogone.v1`](https://gopkg.in/ewgra/ogone.v1)

**Installation:**


```sh
$ go get gopkg.in/ewgra/ogone.v1
```

And then import it in your code:

```go
import "gopkg.in/ewgra/ogone.v1"
```

## Example

### [DirectLink](https://payment-services.ingenico.com/int/en/ogone/support/guides/integration%20guides/directlink) request with alias

```go
	dlr := NewDirectLinkRequest()

	dlr.
		SetAlias("ALIAS").
		SetAmount("100").
		SetReserveOperation().
		SetCurrency("EUR").
		SetOrderId("ORDERID")

	dlg := NewDirectLinkGateway()

	dlr.
		SetPspId("ewgraogone").
		SetUserId("ewgragolang").
		SetPassword("123123aa").
		Sign("qwdqwoidj29812d9")

	dlResp, _ := dlg.Send(dlr) // Use SandboxSend for send it to sandbox

	if !dlResp.IsAuthorised() {
		// .. STATUS in response is not Authorized
	}

	// Request was authorized successfully

```

### Perform server-to-server [Alias](https://payment-services.ingenico.com/int/en/ogone/support/guides/integration%20guides/alias-gateway) request

```go
	ar := NewAliasRequest()

	ar.
		SetAcceptUrl("https://github.com/ewgRa/ogone/success").
		SetExceptionUrl("https://github.com/ewgRa/ogone/exception").
		SetOrderId("ORDERID").
		SetCardNumber("4111111111111111").
		SetCardHolderName("Sökolov Evgenii").
		SetCvc("123").
		SetCardExpireMonth("01").
		SetCardExpireYear("2020")

	ag := NewAliasGateway()

	ar.
		SetPspId("ewgraogone").
		Sign("qwdqwoidj29812d9")

	aResp, _ := ag.Send(ar) // Use SandboxSend for send it to sandbox

	if !aResp.CheckSign("qwdqwoidj29812d9") {
		// ... signuture check failed, do not trust this request
	}

	if !aResp.IsOk() {
		// ... STATUS at response is not OK
	}

	// Alias created, you can use it
```

## Status

### DirectLink

- [x] With Alias
- [ ] another API

### Alias

- [x] Server-to-server request
- [ ] another API

## How to contribute

Open issue or pull requests

## LICENSE

MIT-LICENSE. See [LICENSE](http://ewgra.mit-license.org/)
or the LICENSE file provided in the repository for details.
