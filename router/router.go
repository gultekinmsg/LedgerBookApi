package router

import (
	"github.com/gorilla/mux"
	"guestLedger/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	var obj middleware.Handle
	router.HandleFunc("/list", obj.GetAllGuestMessages).Methods("GET", "OPTIONS")
	router.HandleFunc("/create", obj.CreateGuestMessage).Methods("POST", "OPTIONS")
	return router
}
