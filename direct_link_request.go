package ogone

type DirectLinkRequest struct {
	*BaseRequest
}

func NewDirectLinkRequest(url string) *DirectLinkRequest {
	return &DirectLinkRequest{NewBaseRequest(url)}
}

func (r *DirectLinkRequest) SetReserveOperation() *DirectLinkRequest {
	r.SetOperation("RES")
	return r
}

func (r *DirectLinkRequest) SetOperation(operation string) *DirectLinkRequest {
	r.data["OPERATION"] = operation
	return r
}

func (r *DirectLinkRequest) SetCurrency(currency string) *DirectLinkRequest {
	r.data["CURRENCY"] = currency
	return r
}

func (r *DirectLinkRequest) SetAlias(alias string) *DirectLinkRequest {
	r.data["ALIAS"] = alias
	return r
}

func (r *DirectLinkRequest) SetOrderId(orderId string) *DirectLinkRequest {
	r.data["ORDERID"] = orderId
	return r
}

func (r *DirectLinkRequest) SetAmount(amount string) *DirectLinkRequest {
	r.data["AMOUNT"] = amount
	return r
}
