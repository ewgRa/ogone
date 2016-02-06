package ogone

import (
	"bytes"
"net/http"
	"net/url"
)

type AliasGateway struct {
}

func NewAliasGateway() *AliasGateway {
	return &AliasGateway{}
}

func (g *AliasGateway) SandboxSend(r *AliasRequest) (*AliasResponse, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/test/alias_gateway_utf8.asp")
}

func (g *AliasGateway) Send(r *AliasRequest) (*AliasResponse, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/prod/alias_gateway_utf8.asp")
}

func (g *AliasGateway) sendRequest(r *AliasRequest, gatewayUrl string) (*AliasResponse, error) {

	values := url.Values{}

	for k, v := range r.Data() {
		values.Add(k, v)
	}

	req, err := http.NewRequest("GET", gatewayUrl+"?"+values.Encode(), bytes.NewBufferString(""))

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	redirectUrl, err := res.Location()

	if err != nil {
		return nil, err
	}

	return NewAliasResponse(redirectUrl), nil
}
