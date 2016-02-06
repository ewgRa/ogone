package ogone

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DirectLinkGateway struct {
}

func NewDirectLinkGateway() *DirectLinkGateway {
	return &DirectLinkGateway{}
}

func (g *DirectLinkGateway) Send(r *DirectLinkRequest) (*DirectLinkResponse, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/prod/orderdirect_utf8.asp")
}

func (g *DirectLinkGateway) SandboxSend(r *DirectLinkRequest) (*DirectLinkResponse, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/test/orderdirect_utf8.asp")
}

func (g *DirectLinkGateway) sendRequest(r *DirectLinkRequest, gatewayUrl string) (*DirectLinkResponse, error) {

	values := url.Values{}

	for k, v := range r.Data() {
		values.Add(k, v)
	}

	req, err := http.NewRequest("POST", gatewayUrl+"?"+values.Encode(), bytes.NewBufferString(""))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return errors.New("No redirect")
	},
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Wrong http answer")
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, err
	}

	dlResp, err := NewDirectLinkResponse(string(content))

	if err != nil {
		return nil, err
	}

	return dlResp, nil
}
