package ogone

type Request interface {
	// FIXME XXX: visibility
	Data() map[string]string
	Url() string
}

type BaseRequest struct {
	data map[string]string
	url  string
}

func (r *BaseRequest) Data() map[string]string {
	return r.data
}

func (r *BaseRequest) Url() string {
	return r.url
}

func NewBaseRequest(url string) *BaseRequest {
	return &BaseRequest{data: make(map[string]string), url: url}
}

type BaseGateway struct {
	c *Config
}

func (g *BaseGateway) Config() *Config {
	return g.c
}
