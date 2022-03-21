package helper

import (
	"encoding/json"
	"net/http"
)

type TransactionResult struct {
	Success bool
	Message string
	Data    interface{}
}
type TransactionBool struct {
	Success bool
}
type Transaction interface {
}

func JsonReturn(jsonReturn Transaction, responseWriter http.ResponseWriter) {

	jsonResult, conversionError := json.Marshal(jsonReturn)

	if conversionError != nil {
		failure := TransactionBool{Success: false}
		failureResult, _ := json.Marshal(failure)
		responseWriter.Write(failureResult)
		return
	}

	responseWriter.Write(jsonResult)
}