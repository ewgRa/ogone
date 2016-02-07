package ogone

type request interface {
	Data() map[string]string
}

type baseRequest struct {
	data map[string]string
}

func (r *baseRequest) Data() map[string]string {
	return r.data
}

func newBaseRequest() *baseRequest {
	return &baseRequest{data: make(map[string]string)}
}
