package ogone

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DirectLinkGateway struct {
	*BaseGateway
}

func NewDirectLinkGateway(c *Config) *DirectLinkGateway {
	return &DirectLinkGateway{&BaseGateway{c: c}}
}

func (g *DirectLinkGateway) Url() string {
	return "https://secure.ogone.com/ncol/prod/orderdirect_utf8.asp"
}

func (g *DirectLinkGateway) SandboxUrl() string {
	return "https://secure.ogone.com/ncol/test/orderdirect_utf8.asp"
}

// FIXME XXX: errors
func (g *DirectLinkGateway) Send(r *DirectLinkRequest) string {
	data := make(map[string]string, len(r.Data())+3)

	for k, v := range r.Data() {
		data[k] = v
	}

	data["PSPID"] = g.Config().pspId
	data["USERID"] = g.Config().userId
	data["PSWD"] = g.Config().password

	values := url.Values{}

	for k, v := range data {
		values.Add(k, v)
	}

	values.Add("SHASIGN", shaInCompose(data, g.Config().inPassPhrase))

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
