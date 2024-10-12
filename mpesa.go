package mpesa_go_sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCall *ApiCall) SendStkPush(ctx context.Context, request StkRequest) (*StkAcknowledgement, error) {

	tStamp := time.Now().Format("20060102150405")
	stkReq := &StkPush{
		BusinessShortCode: request.BusinessShortCode,
		PhoneNumber:       request.PhoneNumber,
		PartyA:            request.PhoneNumber,
		PartyB:            request.BusinessShortCode,
		Amount:            request.Amount,
		Timestamp:         "20160216165627",
		TransactionType:   request.TransactionType,
		AccountReference:  request.AccountReference,
		CallBackURL:       request.CallBackUrl,
		TransactionDesc:   "Initiate Transaction",
		Password:          generatePassword(request, tStamp, apiCall.isProd),
	}

	log.log("STK-PUSH REQUEST", uuid.NewString(), 200, stkReq)

	var stkAck *StkAcknowledgement

	res, err := apiCall.MakeApiRequest(ctx, apiCall.baseUrl+"/mpesa/stkpush/v1/processrequest", stkReq)
	if err != nil {
		log.log("STK-PUSH REQUEST", uuid.NewString(), 400, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.log("STK-PUSH REQUEST", uuid.NewString(), 400, res)
	}

	if err = json.NewDecoder(res.Body).Decode(&stkAck); err != nil {
		log.log("STK-PUSH REQUEST", uuid.NewString(), 400, err)
		return nil, err
	}
	return stkAck, nil

}

func generatePassword(stkReq StkRequest, tStamp string, isProduction bool) string {
	if !isProduction {
		return API_PASSWORD
	}

	return base64.StdEncoding.EncodeToString([]byte(stkReq.Passkey + tStamp + stkReq.BusinessShortCode))
}
