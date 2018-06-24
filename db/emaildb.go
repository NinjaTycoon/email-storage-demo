package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"ninja/email/message"
)

type EmailDB struct {
	DSN string
	DriverName string
	db *sql.DB
}

// Obtains a new default instance of EmailDB, setup for in-memory with sqlite3 driver.
// You can override these values before you do your first operation.
func NewDatabase() EmailDB {
	edb := EmailDB{
		DSN: ":memory:",
		DriverName: "sqlite3",
	}
	return edb
}

func (edb *EmailDB) Db() *sql.DB {
	if edb.db == nil {
		edb.Open()
		edb.InitSchema()
	}
	return edb.db;
}

func (edb *EmailDB) Open() {
	edb.db, _ = sql.Open(edb.DriverName, edb.DSN)
}

func (edb *EmailDB) InitSchema() {
	statement, _ := edb.Db().Prepare("CREATE TABLE IF NOT EXISTS emails (id INTEGER PRIMARY KEY, header_from TEXT, header_to TEXT, header_subject TEXT, body TEXT)")
	statement.Exec()
}

func (edb *EmailDB) Insert(msg message.Email) (sql.Result, error) {
	statement, _ := edb.Db().Prepare("INSERT INTO emails (header_from, header_to, header_subject, body) VALUES (?, ?, ?, ?)")
	return statement.Exec(msg.From, msg.ToAsJsonString(), msg.Subject, msg.Body)
}

func (edb *EmailDB) Delete(id int) (sql.Result, error) {
	statement, _ := edb.Db().Prepare("DELETE FROM emails WHERE id = ?")
	return statement.Exec(uint64(id))
}

// Can scan Row or Rows to create a single Email instance
func scan(row scannable) message.Email {
	var toStr string
	msg := message.Email{}
	row.Scan(&msg.Id, &msg.From, &toStr, &msg.Subject, &msg.Body)
	msg.SetToUsingJson(toStr)
	return msg
}

// Need to scan both Row and Rows
type scannable interface {
	Scan(dest ...interface{}) error
}

func (edb *EmailDB) FetchAll() []message.Email {
	rows, _ := edb.Db().Query("SELECT id, header_from, header_to, header_subject, body FROM emails")
	//fmt.Printf("Error:%v\n", error)
	var msg message.Email
	emails := []message.Email{}
	for rows.Next() {
		msg = scan(rows)
		emails = append(emails, msg)
	}
	return emails
}

func (edb *EmailDB) Fetch(id int) message.Email {
	statement, _ := edb.Db().Prepare("SELECT id, header_from, header_to, header_subject, body FROM emails WHERE id = ?")
	row := statement.QueryRow(id)
	msg := scan(row)
	return msg
}