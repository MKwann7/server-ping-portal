package main

import (
	"github.com/MKwann7/server-ping-portal/src/app/controllers/api/ping"
	"github.com/MKwann7/server-ping-portal/src/app/controllers/healthcheck"
	"github.com/MKwann7/server-ping-portal/src/app/libraries/auth"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

var (
	clients = make(map[*websocket.Conn]bool)
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := router()
	middleware := negroni.Classic()
	middleware.UseHandler(router)

	if os.Getenv("ENV") == "local" {
		router.Headers("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Methods")
		router.Headers("Access-Control-Allow-Origin", "*")
		router.Headers("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		router.Headers("Content-Type", "application/json charset=utf-8")
	}

	http.ListenAndServe(":"+os.Getenv("PORT_NUM"), corseHandler(middleware))
}

func router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/health-check").
		HandlerFunc(healthcheck.HandleHealthCheck)
	router.
		Methods("GET").
		Path("/api/v1/ping/server").
		HandlerFunc(validateAuth(ping.HandleServerRequest))

	return router
}

func corseHandler(handler http.Handler) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, webRequest *http.Request) {
		if webRequest.Method == "OPTIONS" {
			responseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Methods")
			responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
			responseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
			responseWriter.Header().Set("Content-Type", "application/json charset=utf-8")
		} else {
			handler.ServeHTTP(responseWriter,webRequest)
		}
	}
}

func validateAuth(f func(responseWriter http.ResponseWriter, webRequest *http.Request)) func(http.ResponseWriter, *http.Request)  {

	user, validationError := auth.ValidateJwt(webRequest)

	//if validationError != nil {
	//	errorManagement.HandleErr(responseWriter, validationError, http.StatusBadRequest)
	//	log.Println(validationError.Error())
	//	return
	//}

	return f

}