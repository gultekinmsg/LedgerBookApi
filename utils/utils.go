package utils

import (
	"database/sql"
	"guestLedger/models"
	"log"
	"net/http"
)

func CheckErrorFatal(err error, message string) {
	if err != nil {
		log.Fatalf(message)
	}
}
func CheckErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func Response(didSuccessful bool) string {
	if didSuccessful {
		return "200"
	} else {
		return "500"
	}
}
func SqlToObject(result *sql.Rows) []models.GuestLedger {
	var list []models.GuestLedger
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result)
	for result.Next() {
		var ledger models.GuestLedger
		err := result.Scan(&ledger.EMAIL, &ledger.MESSAGE)
		CheckErrorFatal(err, "Can not iterate on rows")
		list = append(list, ledger)
	}
	return list
}
func GetListPage(r *http.Request) string {
	keys, ok := r.URL.Query()["page"]
	if !ok {
		log.Fatalf("Cannot retrieve page from request")
	}
	return keys[0]
}
