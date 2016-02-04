package ogone

const operationReserve = "RES"

type DirectLinkRequest struct {
	data map[string]string
}

func NewDirectLinkRequest() *DirectLinkRequest {
	return &DirectLinkRequest{data: make(map[string]string)}
}

func (dlr *DirectLinkRequest) SetReserveOperation() *DirectLinkRequest{
	dlr.SetOperation(operationReserve)
	return dlr
}

func (dlr *DirectLinkRequest) SetOperation(operation string) *DirectLinkRequest{
	dlr.data["OPERATION"] = operation
	return dlr
}

func (dlr *DirectLinkRequest) SetCurrency(currency string) *DirectLinkRequest{
	dlr.data["CURRENCY"] = currency
	return dlr
}

func (dlr *DirectLinkRequest) SetAlias(alias string) *DirectLinkRequest{
	dlr.data["ALIAS"] = alias
	return dlr
}

func (dlr *DirectLinkRequest) SetOrderId(orderId string) *DirectLinkRequest{
	dlr.data["ORDERID"] = orderId
	return dlr
}

func (dlr *DirectLinkRequest) SetAmount(amount string) *DirectLinkRequest{
	dlr.data["AMOUNT"] = amount
	return dlr
}
