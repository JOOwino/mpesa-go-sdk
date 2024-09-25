package mpesa_go_sdk

type StkRequest struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Passkey           string `json:"Passkey"`
	Amount            int64  `json:"Amount"`
	PhoneNumber       string `json:"PhoneNumber"`
	TransactionType   string `json:"TransactionType"`
	AccountReference  string `json:"AccountReference"`
	CallBackUrl       string `json:"CallBackUrl"`
}

type AuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	Duration    string `json:"expires_in"`
}

type StkPush struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            int64  `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

type StkAcknowledgement struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
	RequestId           string `json:"requestId"`
	ErrorCode           string `json:"errorCode"`
	ErrorMessage        string `json:"errorMessage"`
}

type DynamicQRReq struct {
	CPI          string `json:"CPI"`
	MerchantName string `json:"MerchantName"`
	RefNo        string `json:"RefNo"`
	Amount       int    `json:"Amount"`
	Size         int8   `json:"Size"`
}

type DynamicQRRes struct {
	ResponseCode        string `json:"ResponseCode"`
	RequestId           string `json:"RequestId"`
	ResponseDescription string `json:"ResponseDescription"`
	QrCode              []byte `json:"QrCode"`
}
