package ogone

import "net/url"

type AliasResponse struct {
	redirectUrl *url.URL
}

func (r *AliasResponse) Alias() string {
	return r.redirectUrl.Query().Get("Alias")
}
