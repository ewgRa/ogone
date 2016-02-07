package ogone

import (
	"bytes"
	"net/http"
	"net/url"
)

// AliasGateway for sending request to Alias API
type AliasGateway struct {
}

// NewAliasGateway create new AliasGateway
func NewAliasGateway() *AliasGateway {
	return &AliasGateway{}
}

// Send Alias request to Alias API
func (g *AliasGateway) Send(r *AliasRequest) (*AliasResponse, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/prod/alias_gateway_utf8.asp")
}

// SandboxSend for send Alias request to sandbox Alias API
func (g *AliasGateway) SandboxSend(r *AliasRequest) (*AliasResponse, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/test/alias_gateway_utf8.asp")
}

func (g *AliasGateway) sendRequest(r *AliasRequest, gatewayURL string) (*AliasResponse, error) {

	values := url.Values{}

	for k, v := range r.Data() {
		values.Add(k, v)
	}

	req, err := http.NewRequest("GET", gatewayURL +"?"+values.Encode(), bytes.NewBufferString(""))

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	redirectURL, err := res.Location()

	if err != nil {
		return nil, err
	}

	return newAliasResponse(redirectURL), nil
}
