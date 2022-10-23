package controllers

import (
	"github.com/MKwann7/server-ping-portal/app/core/src/code/libraries/helper"
	"net/http"
)

func HandleHealthCheck(responseWriter http.ResponseWriter, webRequest *http.Request) {

	healthCheck := helper.TransactionBool{Success: true}
	helper.JsonReturn(healthCheck, responseWriter)
}
