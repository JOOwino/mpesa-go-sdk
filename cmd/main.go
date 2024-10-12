package main

import (
	"context"
	"fmt"
	mpesa "github.com/JOOwino/mpesa-go-sdk"
	"os"
	"time"
)

const (
	apiKey    = "xHCvJADoytWq8OL8PFqYXcTiDbmKDZ9q8cNStzI2q2CV6iiU"
	apiSecret = "5W9egqxcAVawsc7C42kwYqxIWcVoguob7HxBl0OkUGIF6xRay11aL2uC9wgQ3Gdt"
)

// A test for the Initiating NI Push
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("URL:  \n" + os.Getenv("MPESA_CALL_BACK"))

	//Create A STK Request
	req := mpesa.StkRequest{
		BusinessShortCode: "174379",
		TransactionType:   "CustomerPayBillOnline",
		PhoneNumber:       "254710119383",
		AccountReference:  "Test13",
		Amount:            1,
		Passkey:           "",
		CallBackUrl:       os.Getenv("MPESA_CALL_BACK"),
	}

	fmt.Printf("REQ: %v", req)

	mpesaCall := mpesa.New(apiKey, apiSecret, false)
	fmt.Println("After Making MPESA CALL")
	stkRes, err := mpesaCall.SendStkPush(ctx, req)
	if err != nil {
		fmt.Printf("Error while doing STkPush: %v \n", err)
	}

	fmt.Printf("Stk AckRes: %v \n", stkRes)
}
