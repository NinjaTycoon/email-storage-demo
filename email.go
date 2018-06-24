package main

import (
	"fmt"
	"ninja/email/db"
	"time"
	"ninja/email/message"
	"strconv"
)

// If you point to a file, then you can just keep running it and watch it go from 1 to 5 rows, then back down to 1.
// If you use default of in-memory DB, it will always list 1 row.
func main() {
	fmt.Println("Running...", time.Now())
	edb := db.NewDatabase()

	// It defaults to an in-memory database.
	// Enable this line and point to a file to use file based instead of memory.  The file will be
	// created if it does not exist.  When using a file instead of memory, it will add a row every time you run it
	// demonstrating that the previous inserts persisted.

	edb.DSN = "/home/erik/emaildb"

	InitData(&edb)
	TestDB(&edb)

	// delete all rows if max reached
	Truncate(&edb, 5)

}

func InitData(edb *db.EmailDB) {
	msg := message.Email{
		From: "here@there.com",
		To: []string{"them@there.com", "you@there.com"},
		Subject: "Hi, there",
		Body: "Drink milk!",
	}
	edb.Insert(msg)
}


func TestDB(edb *db.EmailDB) {
	fmt.Println("Start TestDB.  Emails...")
	var emails []message.Email
	emails = edb.FetchAll()
	fmt.Println("Len:", len(emails))
	for _, msg := range emails {
		fmt.Println(strconv.Itoa(msg.Id) + ": From: " + msg.From + ", To: " + msg.ToAsJsonString() + ", Subject: " + msg.Subject)
		fmt.Println("\t" + msg.Body)
	}

	fmt.Println("Finished")
}

// Truncates if max rows reached
func Truncate(edb *db.EmailDB, max int) {
	emails := edb.FetchAll()
	if len(emails) >= max {
		for _, email := range emails {
			edb.Delete(email.Id)
		}
	}
}