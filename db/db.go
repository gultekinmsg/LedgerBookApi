package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"guestLedger/models"
	"guestLedger/utils"
	"log"
	"os"
	"strconv"
	"time"
)

type DatabaseActions interface {
	createConnection() *sql.DB
	InsertGuestMessage(message models.GuestLedger) bool
	GetAllMessages(page string) ([]models.GuestLedger, error)
}
type Database struct {
}

func (d *Database) createConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	utils.CheckErrorPanic(err)
	err = db.Ping()
	utils.CheckErrorPanic(err)
	return db
}
func (d *Database) InsertGuestMessage(message models.GuestLedger) bool {
	db := d.createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Can not close db connection after insert")
		}
	}(db)
	t := time.Now()
	sqlStatement := `insert into "ledger"."Messages"("Email", "Message","Time") values($1, $2,$3)`
	_, err := db.Exec(sqlStatement, message.EMAIL, message.MESSAGE, t)
	utils.CheckErrorPanic(err)
	return true
}
func (d *Database) GetAllMessages(page string) ([]models.GuestLedger, error) {
	db := d.createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Can not close db connection after get")
		}
	}(db)
	pageNum, err := strconv.Atoi(page)
	utils.CheckErrorFatal(err, "Unable to retrieve page")
	offset := (pageNum - 1) * 10
	sqlStatement := `select "Email","Message" from ledger."Messages" order by "Time" desc limit 10 offset $1`
	rows, err := db.Query(sqlStatement, offset)
	utils.CheckErrorFatal(err, "Unable to execute the select query")
	list := utils.SqlToObject(rows)
	return list, err
}
