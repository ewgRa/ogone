package ogone

import "net/url"

// AliasResponse from alias server-to-server request
type AliasResponse struct {
	redirectURL *url.URL
}

func newAliasResponse(redirectURL *url.URL) *AliasResponse {
	return &AliasResponse{redirectURL: redirectURL}
}

// Alias return alias response parameter
func (r *AliasResponse) Alias() string {
	return r.redirectURL.Query().Get("Alias")
}
// Sign return SHA signature response parameter
func (r *AliasResponse) Sign() string {
	return r.redirectURL.Query().Get("SHASign")
}

// IsOk for checking that alias was created
func (r *AliasResponse) IsOk() bool {
	return r.redirectURL.Query().Get("status") == "0"
}

// CheckSign for checking signature equality
func (r *AliasResponse) CheckSign(passPhrase string) bool {
	params := make(map[string]string)

	for k, v := range r.redirectURL.Query() {
		for _, iv := range v {
			params[k] = iv
		}
	}

	return r.Sign() == shaOutAliasCompose(params, passPhrase)
}
