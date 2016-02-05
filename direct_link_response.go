package ogone

type DirectLinkResponse struct {
	data map[string]string
}

func NewDirectLinkResponse() *DirectLinkResponse {
	return &DirectLinkResponse{}
}
