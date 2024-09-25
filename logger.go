package mpesa_go_sdk

import (
	"log/slog"
	"os"
)

type CustomLogger struct {
}

var log *CustomLogger
var slogger *slog.Logger

func init() {
	log = &CustomLogger{}
	slogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func (Logger *CustomLogger) log(requestType, transactionId string, errorCode int, reqBody interface{}) {
	switch errorCode {
	case 400:
		slogger.Warn(requestType, "transactionId", transactionId, "Body", reqBody)
		break
	case 500:
		slogger.Error(requestType, "transactionId", transactionId, "Body", reqBody)
		break
	default:
		slogger.Info(requestType, "transactionId", transactionId, "Body", reqBody)
		break
	}

}
