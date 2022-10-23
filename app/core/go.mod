module github.com/MKwann7/server-ping-portal

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.2.0
	github.com/urfave/negroni v1.0.0
)

replace github.com/MKwann7/server-ping-portal/app/core/src/code/controllers => ../core/src/code/controllers

replace github.com/MKwann7/server-ping-portal/app/core/src/code/dtos => ../core/src/code/dtos

replace github.com/MKwann7/server-ping-portal/app/core/src/code/libraries/builder => ../core/src/code/libraries/builder

replace github.com/MKwann7/server-ping-portal/app/core/src/code/libraries/db => ../core/src/code/libraries/db

replace github.com/MKwann7/server-ping-portal/app/core/src/code/libraries/exceptions => ../core/src/code/libraries/exceptions

replace github.com/MKwann7/server-ping-portal/app/core/src/code/libraries/helper => ../core/src/code/libraries/helper
