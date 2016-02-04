package ogone
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
	"strings"
)

const dlGateway = "https://secure.ogone.com/ncol/prod/orderdirect_utf8.asp"
const sandboxDlGateway = "https://secure.ogone.com/ncol/test/orderdirect_utf8.asp"

type DirectLink struct {
	c *Config
}

func NewDirectLink(c *Config) *DirectLink {
	return &DirectLink{c: c}
}

func (dl *DirectLink) Send(r *DirectLinkRequest) string {

	data := make(map[string]string, len(r.data)+3)

	for k, v := range r.data {
		data[k] = v
	}

	data["PSPID"] = dl.c.pspId
	data["USERID"] = dl.c.userId
	data["PSWD"] = dl.c.password

	values := url.Values{}

	for k, v := range data {
		values.Add(k, v)
	}

	values.Add("SHASIGN", shaInCompose(data, dl.c.inPassPhrase))

	url := dlGateway

	if (dl.c.sandbox) {
		url = sandboxDlGateway
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return ""
	}

	client := &http.Client{}

	fmt.Println(req)
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
