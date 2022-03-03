package middleware

import (
	"encoding/json"
	"guestLedger/db"
	"guestLedger/models"
	"guestLedger/utils"
	"net/http"
)

type HandleActions interface {
	CreateGuestMessage(w http.ResponseWriter, r *http.Request)
	GetAllGuestMessages(w http.ResponseWriter, r *http.Request)
}
type Handle struct {
}

func (h *Handle) CreateGuestMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var guestMessage models.GuestLedger
	err := json.NewDecoder(r.Body).Decode(&guestMessage)
	utils.CheckErrorFatal(err, "Unable to decode the request body")
	var obj db.Database
	didInsert := obj.InsertGuestMessage(guestMessage)
	response := utils.Response(didInsert)
	err = json.NewEncoder(w).Encode(response)
	utils.CheckErrorFatal(err, "Unable to encode the response body")
}
func (h *Handle) GetAllGuestMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj db.Database
	users, err := obj.GetAllMessages(utils.GetListPage(r))
	utils.CheckErrorFatal(err, "Unable to get all messages")
	err = json.NewEncoder(w).Encode(users)
	utils.CheckErrorFatal(err, "Unable to encode the response body")
}
