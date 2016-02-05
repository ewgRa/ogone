package ogone

import (
	"bytes"
	"net/http"
	"net/url"
)

type AliasGateway struct {
	*BaseGateway
}

func NewAliasGateway(c *Config) *AliasGateway {
	return &AliasGateway{&BaseGateway{c: c}}
}

func (g *AliasGateway) Url() string {
	return "https://secure.ogone.com/ncol/prod/alias_gateway_utf8.asp"
}

func (g *AliasGateway) SandboxUrl() string {
	return "https://secure.ogone.com/ncol/test/alias_gateway_utf8.asp"
}

func (g *AliasGateway) Send(r *AliasRequest) (*AliasResponse, error) {
	data := make(map[string]string, len(r.Data())+1)

	for k, v := range r.Data() {
		data[k] = v
	}

	data["PSPID"] = g.Config().pspId

	values := url.Values{}

	for k, v := range data {
		values.Add(k, v)
	}

	values.Add("SHASIGN", shaInAliasCompose(data, g.Config().inPassPhrase))

	url := g.Url()

	if g.Config().sandbox {
		url = g.SandboxUrl()
	}

	req, err := http.NewRequest("GET", url+"?"+values.Encode(), bytes.NewBufferString(""))

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
