package ogone

// DirectLinkRequest for perform DirectLink request
// right before you will send request to gateway, you need to sign it by calling Sign
type DirectLinkRequest struct {
	*baseRequest
}

// NewDirectLinkRequest for create DirectLinkRequest
func NewDirectLinkRequest() *DirectLinkRequest {
	return &DirectLinkRequest{newBaseRequest()}
}

// SetPspID for set PSPID request parameter
func (r *DirectLinkRequest) SetPspID(pspID string) *DirectLinkRequest {
	r.data["PSPID"] = pspID
	return r
}

// SetUserID for set user ID request parameter
func (r *DirectLinkRequest) SetUserID(userID string) *DirectLinkRequest {
	r.data["USERID"] = userID
	return r
}

// SetPassword for set password request parameter
func (r *DirectLinkRequest) SetPassword(password string) *DirectLinkRequest {
	r.data["PSWD"] = password
	return r
}

// SetReserveOperation mark request as reserve operation
func (r *DirectLinkRequest) SetReserveOperation() *DirectLinkRequest {
	r.SetOperation("RES")
	return r
}

// SetOperation for set operation request parameter
func (r *DirectLinkRequest) SetOperation(operation string) *DirectLinkRequest {
	r.data["OPERATION"] = operation
	return r
}

// SetCurrency for set currecy request parameter
func (r *DirectLinkRequest) SetCurrency(currency string) *DirectLinkRequest {
	r.data["CURRENCY"] = currency
	return r
}

// SetAlias for set alias request parameter
func (r *DirectLinkRequest) SetAlias(alias string) *DirectLinkRequest {
	r.data["ALIAS"] = alias
	return r
}

// SetOrderID for set order ID request parameter
func (r *DirectLinkRequest) SetOrderID(orderID string) *DirectLinkRequest {
	r.data["ORDERID"] = orderID
	return r
}

// SetAmount for set amount request parameter
func (r *DirectLinkRequest) SetAmount(amount string) *DirectLinkRequest {
	r.data["AMOUNT"] = amount
	return r
}

// Sign for sign request by SHASIGN
func (r *DirectLinkRequest) Sign(passPhrase string) *DirectLinkRequest {
	r.data["SHASIGN"] = shaInCompose(r.data, passPhrase)
	return r
}
