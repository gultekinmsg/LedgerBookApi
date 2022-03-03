package main

import (
	"github.com/joho/godotenv"
	"guestLedger/router"
	"guestLedger/utils"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	utils.CheckErrorFatal(err, "Error loading .env file")
	r := router.Router()
	err = http.ListenAndServe(os.Getenv("PORT"), r)
	utils.CheckErrorPanic(err)
}
