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

func (g *AliasGateway) Url() string {
	return "https://secure.ogone.com/ncol/prod/alias_gateway_utf8.asp"
}

func (g *AliasGateway) SandboxUrl() string {
	return "https://secure.ogone.com/ncol/test/alias_gateway_utf8.asp"
}

func (g *AliasGateway) Send(r *AliasRequest) (*AliasResponse, error) {

	values := url.Values{}

	for k, v := range r.Data() {
		values.Add(k, v)
	}

	req, err := http.NewRequest("GET", r.Url()+"?"+values.Encode(), bytes.NewBufferString(""))

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	redirectUrl, _ := res.Location()

	return &AliasResponse{redirectUrl: redirectUrl}, nil
}
