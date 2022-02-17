package my_http

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCallGoogle(t *testing.T) {
	msgs := []string{}
	var msg string
	fmt.Println("Starting Test of CallGoogleAndUsingReaderInterface")
	msg = testCallGoogleAndUsingReaderInterface(t)
	msgs = append(msgs, "when we want to get from google "+msg)

	msg = testCallGoogleAndUsingReaderInterfaceMoreSimple(t)
	msgs = append(msgs, "when we want to get from google "+msg)

	PrintTestsMessages(msgs)
}

func testCallGoogleAndUsingReaderInterfaceMoreSimple(t *testing.T) string {
	msg := "should ..."
	// Given
	someResult := CallGoogleAndUsingReaderInterfaceUsingMyWriter()
	// Then
	require.Condition(t, func() (success bool) {
		return someResult == MyFakeWrite
	})
	return msg
}

func testCallGoogleAndUsingReaderInterface(t *testing.T) string {
	msg := "should return the body as byte slice"
	// Given
	someResult := CallGoogleAndUsingReaderInterface()
	// Then
	require.Condition(t, func() (success bool) {
		return len(someResult) == 99999
	})
	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
