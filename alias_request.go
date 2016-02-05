package ogone

type AliasRequest struct {
	*BaseRequest
}

func NewAliasRequest() *AliasRequest {
	return &AliasRequest{NewBaseRequest()}
}

func (r *AliasRequest) SetPspId(pspId string) *AliasRequest {
	r.data["PSPID"] = pspId
	return r
}

func (r *AliasRequest) SetAcceptUrl(url string) *AliasRequest {
	r.data["ACCEPTURL"] = url
	return r
}

func (r *AliasRequest) SetExceptionUrl(url string) *AliasRequest {
	r.data["EXCEPTIONURL"] = url
	return r
}

func (r *AliasRequest) SetOrderId(orderId string) *AliasRequest {
	r.data["ORDERID"] = orderId
	return r
}

func (r *AliasRequest) SetCardNumber(number string) *AliasRequest {
	r.data["CARDNO"] = number
	return r
}

func (r *AliasRequest) SetCardHolderName(name string) *AliasRequest {
	r.data["CN"] = name
	return r
}

func (r *AliasRequest) SetCvc(cvc string) *AliasRequest {
	r.data["CVC"] = cvc
	return r
}

func (r *AliasRequest) SetCardExpireMonth(month string) *AliasRequest {
	r.data["ECOM_CARDINFO_EXPDATE_MONTH"] = month
	return r
}

func (r *AliasRequest) SetCardExpireYear(year string) *AliasRequest {
	r.data["ECOM_CARDINFO_EXPDATE_YEAR"] = year
	return r
}

func (r *AliasRequest) Sign(passPhrase string) *AliasRequest {
	r.data["SHASIGN"] = shaInAliasCompose(r.data, passPhrase)
	return r
}
