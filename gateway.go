package ogone

type Request interface {
	Data() map[string]string
}

type BaseRequest struct {
	data map[string]string
}

func (r *BaseRequest) Data() map[string]string {
	return r.data
}

func NewBaseRequest() *BaseRequest {
	return &BaseRequest{data: make(map[string]string)}
}

type Gateway interface {
	Url() string
	SandboxUrl() string
	Config() *Config
}

type BaseGateway struct {
	c *Config
}

func (g *BaseGateway) Config() *Config {
	return g.c
}
