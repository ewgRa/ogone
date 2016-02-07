# Golang library for Ogone integration

This library is an [Ingenico Payment Services](https://payment-services.ingenico.com/int/en) integration library for the
[Go](http://www.golang.org/) programming language.

[![Build Status](https://travis-ci.org/ewgRa/ogone.svg?branch=master)](https://travis-ci.org/ewgRa/ogone)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ewgra/ogone/master/LICENSE)
[![GoReportCard](http://goreportcard.com/badge/ewgra/ogone)](http://goreportcard.com/report/ewgra/ogone)
[![codecov.io](https://codecov.io/github/ewgRa/ogone/coverage.svg?branch=master)](https://codecov.io/github/ewgRa/ogone?branch=master)

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
		SetOrderID("ORDERID")

	dlg := NewDirectLinkGateway()

	dlr.
		SetPspID("ewgraogone").
		SetUserID("ewgragolang").
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
		SetAcceptURL("https://github.com/ewgRa/ogone/success").
		SetExceptionURL("https://github.com/ewgRa/ogone/exception").
		SetOrderID("ORDERID").
		SetCardNumber("4111111111111111").
		SetCardHolderName("Sökolov Evgenii").
		SetCardCvc("123").
		SetCardExpireMonth("01").
		SetCardExpireYear("2020")

	ag := NewAliasGateway()

	ar.
		SetPspID("ewgraogone").
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
