package db

import (
	"testing"
	"fmt"
	"ninja/email/message"
	"github.com/stretchr/testify/assert"
)

var to = []string{"them1@there.com", "you1@there.com"}
var msg = message.Email{
	From: "here1@there.com",
	To: to,
	Subject: "Hi, there1",
	Body: "Drink milk!",
}
var msg2 = message.Email{
	From: "here1@tubs.com",
	To: to,
	Subject: "Hi, Tubs!",
	Body: "Drink tea!",
}

func TestEmailDB_Insert(t *testing.T) {
	fmt.Println("TestEmailDB_Insert")
	edb := NewDatabase()
	result, error := edb.Insert(msg)
	assert.Nil(t, error)
	assert.NotNil(t, result)
	id, _ := result.LastInsertId()
	re, _ := result.RowsAffected()
	assert.True(t, id > 0)
	assert.True(t, re == 1)
	msg.Id = int(id)

	jMsg, _ := msg.ToJson()
	fmt.Printf("Email %v:%q\n", id, jMsg)
}

func TestEmailDB_FetchAll(t *testing.T) {
	edb := NewDatabase()
	edb.Insert(msg)
	edb.Insert(msg2)

	emails := edb.FetchAll()
	assert.Equal(t, 2, len(emails))
}

func TestEmailDB_Fetch(t *testing.T) {
	edb := NewDatabase()
	edb.Insert(msg)
	result, _ := edb.Insert(msg2)
	id, _ := result.LastInsertId()
	assert.True(t, id > 0)

	msg2b := edb.Fetch(int(id))
	msg2.Id = int(id)
	assert.Equal(t, msg2, msg2b)
}

func TestEmailDB_Delete(t *testing.T) {
	edb := NewDatabase()
	edb.Insert(msg)
	edb.Insert(msg2)

	emails := edb.FetchAll()
	assert.Equal(t, 2, len(emails))

	edb.Delete(emails[0].Id)
	assert.Equal(t, 1, len(edb.FetchAll()))

	edb.Delete(emails[1].Id)
	assert.Equal(t, 0, len(edb.FetchAll()))
}

