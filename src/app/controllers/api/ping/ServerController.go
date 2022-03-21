package ping

import (
	"github.com/MKwann7/server-ping-portal/src/app/libraries/helper"
	"net/http"
)

func HandleServerRequest(responseWriter http.ResponseWriter, webRequest *http.Request) {

	healthCheck := helper.TransactionBool{Success: false}
	helper.JsonReturn(healthCheck, responseWriter)
}