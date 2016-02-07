package ogone

// AliasRequest for perform alias server-to-server request
// right before you will send request to gateway, you need to sign it by calling Sign
type AliasRequest struct {
	*baseRequest
}

// NewAliasRequest create AliasRequest
func NewAliasRequest() *AliasRequest {
	return &AliasRequest{newBaseRequest()}
}

// SetPspID for set PSPID request parameter
func (r *AliasRequest) SetPspID(pspID string) *AliasRequest {
	r.data["PSPID"] = pspID
	return r
}

// SetAcceptURL for set ACCEPTURL request parameter
func (r *AliasRequest) SetAcceptURL(url string) *AliasRequest {
	r.data["ACCEPTURL"] = url
	return r
}

// SetExceptionURL for set EXCEPTIONURL request parameter
func (r *AliasRequest) SetExceptionURL(url string) *AliasRequest {
	r.data["EXCEPTIONURL"] = url
	return r
}

// SetOrderID for set ORDERID request parameter
func (r *AliasRequest) SetOrderID(orderID string) *AliasRequest {
	r.data["ORDERID"] = orderID
	return r
}

// SetCardNumber for set card number request parameter
func (r *AliasRequest) SetCardNumber(number string) *AliasRequest {
	r.data["CARDNO"] = number
	return r
}

// SetCardHolderName for set card holder name request parameter
func (r *AliasRequest) SetCardHolderName(name string) *AliasRequest {
	r.data["CN"] = name
	return r
}

// SetCardCVC for set card CVC request parameter
func (r *AliasRequest) SetCardCVC(cvc string) *AliasRequest {
	r.data["CVC"] = cvc
	return r
}

// SetCardExpireMonth for set card month expiration request parameter
func (r *AliasRequest) SetCardExpireMonth(month string) *AliasRequest {
	r.data["ECOM_CARDINFO_EXPDATE_MONTH"] = month
	return r
}

// SetCardExpireYear for set card year expiration request parameter
func (r *AliasRequest) SetCardExpireYear(year string) *AliasRequest {
	r.data["ECOM_CARDINFO_EXPDATE_YEAR"] = year
	return r
}

// Sign for sign request by SHASIGN
func (r *AliasRequest) Sign(passPhrase string) *AliasRequest {
	r.data["SHASIGN"] = shaInAliasCompose(r.data, passPhrase)
	return r
}
