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

func (g *DirectLinkGateway) Send(r *DirectLinkRequest) (string, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/prod/orderdirect_utf8.asp")
}

func (g *DirectLinkGateway) SandboxSend(r *DirectLinkRequest) (string, error) {
	return g.sendRequest(r, "https://secure.ogone.com/ncol/test/orderdirect_utf8.asp")
}

func (g *DirectLinkGateway) sendRequest(r *DirectLinkRequest, gatewayUrl string) (string, error) {

	values := url.Values{}

	for k, v := range r.Data() {
		values.Add(k, v)
	}

	req, err := http.NewRequest("POST", gatewayUrl+"?"+values.Encode(), bytes.NewBufferString(""))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return "", err
	}

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return errors.New("No redirect")
	},
	}

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("Wrong http answer")
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return "", err
	}

	return string(content), nil
}
