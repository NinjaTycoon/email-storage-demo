package message

import (
	"encoding/json"
)

type Email struct {
	Id int
	From string
	To []string
	Subject string
	Body string
}

func (e *Email) ToJson() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Email) FromJson(bJson []byte) error {
	return json.Unmarshal(bJson, e)
}

func (e *Email) ToAsJson() ([]byte, error) {
	return json.Marshal(e.To)
}

func (e *Email) ToAsJsonString() string {
	be, _ := e.ToAsJson()
	return string(be)
}

// Set 'To' from a json string
func (e *Email) SetToUsingJson(sJson string) []string {
	json.Unmarshal([]byte(sJson), &e.To)
	return e.To
}


//type IEmail interface {
//	From() string
//}


