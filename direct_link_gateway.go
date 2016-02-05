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

func (g *DirectLinkGateway) Url() string {
	return "https://secure.ogone.com/ncol/prod/orderdirect_utf8.asp"
}

func (g *DirectLinkGateway) SandboxUrl() string {
	return "https://secure.ogone.com/ncol/test/orderdirect_utf8.asp"
}

// FIXME XXX: errors
func (g *DirectLinkGateway) Send(r *DirectLinkRequest) string {

	values := url.Values{}

	for k, v := range r.Data() {
		values.Add(k, v)
	}

	req, err := http.NewRequest("POST", r.Url()+"?"+values.Encode(), bytes.NewBufferString(""))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return ""
	}

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return errors.New("No redirect")
	},
	}

	res, err := client.Do(req)

	if err != nil {
		return ""
	}

	if res.StatusCode != http.StatusOK {

	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return string(content)
}
