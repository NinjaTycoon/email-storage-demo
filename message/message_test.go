package message

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestToAsJsonString(t *testing.T) {
	to := []string{"them@there.com", "you@there.com"}
	msg := Email{
		From: "here@there.com",
		To: to,
		Subject: "Hi, there",
		Body: "Drink milk!",
	}

	sJsonTo := msg.ToAsJsonString()
	assert.NotNil(t, sJsonTo)
	fmt.Println("sJsonTo:", sJsonTo)
	msg.To = nil
	assert.Nil(t, msg.To)
	to2 := msg.SetToUsingJson(sJsonTo)
	assert.NotNil(t, msg.To)
	assert.Equal(t, to, msg.To)
	fmt.Println("sJsonTo 2:", to2)
}

// Test json marshalling
func TestToJson(t *testing.T) {
	to := []string{"them@there.com", "you@there.com"}
	msg := Email{
		From: "here@there.com",
		To: to,
		Subject: "Hi, there",
		Body: "Drink milk!",
	}

	// Convert Email to json
	jMsg, error := msg.ToJson()

	require.Nil(t, error)
	sMsg := string(jMsg)
	assert.True(t, strings.Contains(sMsg, to[0]))
	fmt.Printf("json: %q\n", jMsg)

	// Convert from json to Email
	msg2 := Email{}
	error = msg2.FromJson(jMsg)
	require.Nil(t, error)
	assert.Equal(t, msg, msg2)
	jMsg, error = msg2.ToJson()
	fmt.Printf("json: %q\n", jMsg)
}