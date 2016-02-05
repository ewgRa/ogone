package ogone

import "net/url"

type AliasResponse struct {
	redirectUrl *url.URL
}

func NewAliasResponse(redirectUrl *url.URL) *AliasResponse {
	return &AliasResponse{redirectUrl: redirectUrl}
}

func (r *AliasResponse) Alias() string {
	return r.redirectUrl.Query().Get("Alias")
}
