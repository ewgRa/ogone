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

func (r *AliasResponse) Sign() string {
	return r.redirectUrl.Query().Get("SHASign")
}

func (r *AliasResponse) IsOk() bool {
	return r.redirectUrl.Query().Get("status") == "0"
}

func (r *AliasResponse) CheckSign(passPhrase string) bool {
	params := make(map[string]string)

	for k, v := range r.redirectUrl.Query() {
		for _, iv := range v {
			params[k] = iv
		}
	}

	return r.Sign() == shaOutAliasCompose(params, passPhrase)
}
