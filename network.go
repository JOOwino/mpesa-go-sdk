package mpesa_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net"
	"net/http"
	"os"
	"time"
)

func init() {
	ipAddr, err := GetIpAddress()
	if err != nil {
		fmt.Printf("Error while getting IP Address : %v", err)
	}
	uri := "http://" + ipAddr + ":8080" + CALLBACK_URI
	os.Setenv("MPESA_CALL_BACK", uri)

}

const (
	CALLBACK_URI        = "/stk-push/callback"
	PRODUCTION_BASE_URL = "https://api.safaricom.co.ke"
	SANDBOX_BASE_URL    = "https://sandbox.safaricom.co.ke"
	API_PASSWORD        = "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTYwMjE2MTY1NjI3"
)

type ApiCall struct {
	apiClient     HttpClient
	apiKey        string
	apiSecret     string
	token         string
	baseUrl       string
	tokenDuration time.Time
	isProd        bool
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(key, secret string, isProduction bool) *ApiCall {

	//At the point of Initializing a ApiClient Struct The ApiCall should have a token Embedded in It
	//Hence, at the point of the Initialization *ApiCall has a token attached to it

	//Return a type in ApiCall that implements the Do interface
	//We can pick the whole http.Client table but we only want to use the Do method Belonging to http.Client
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	baseUrl := SANDBOX_BASE_URL
	if isProduction {
		baseUrl = PRODUCTION_BASE_URL
	}

	return &ApiCall{
		apiClient: client,
		apiKey:    key,
		apiSecret: secret,
		baseUrl:   baseUrl,
		isProd:    isProduction,
	}

}

func (apiCall *ApiCall) generateToken(ctx context.Context) error {
	//Checks if Token Is Set and Not Yet Expired
	if apiCall.token != "" && apiCall.tokenDuration.After(time.Now()) {
		log.log("TOKEN ALREADY EXIST", uuid.NewString(), 400, "Token Still Valid")
		return nil
	}

	res, err := apiCall.MakeHttpRequest(ctx, apiCall.baseUrl+"/oauth/v1/generate?grant_type=client_credentials", http.MethodGet, nil)
	if err != nil {
		return err
	}

	var authRes *AuthorizationResponse
	if err = json.NewDecoder(res.Body).Decode(&authRes); err != nil {
		log.log("GENERATE TOKEN", uuid.NewString(), 500, err)
		return err
	}

	apiCall.token = authRes.AccessToken
	apiCall.tokenDuration = time.Now().Local().Add(time.Second * 3559)
	return nil
}

// MakeApiRequest This ie the Entry Call For APi

func (apiCall *ApiCall) MakeApiRequest(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	if err := apiCall.generateToken(ctx); err != nil {
		return nil, err
	}
	res, err := apiCall.MakeHttpRequest(ctx, url, http.MethodPost, body)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (apiCall *ApiCall) MakeHttpRequest(ctx context.Context, url, method string, body interface{}) (*http.Response, error) {

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.log("ERROR MARSHALLING DATA", uuid.NewString(), 500, err)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.log("MAKE HTTP REQUEST", apiCall.apiKey, 400, err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	if body != nil {
		req.Header.Set("Authorization", "Bearer "+apiCall.token)
	} else {
		req.SetBasicAuth(apiCall.apiKey, apiCall.apiSecret)
	}

	res, err := apiCall.apiClient.Do(req)
	if err != nil {
		fmt.Printf("Error While Sending Request for AppKey: %s \n. Error occured : %v \n", apiCall.apiKey, err)
		log.log("MAKE HTTP REQUEST", apiCall.apiKey, 400, err)
		return nil, err
	}

	return res, nil

}

func GetIpAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	var publicIp string
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				publicIp = ipnet.IP.To4().String()
				fmt.Printf("IP ADDRESS: %s \n", publicIp)
			}
		}

	}

	return publicIp, nil
}
