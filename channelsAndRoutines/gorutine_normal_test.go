package channelsAndRoutines

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var links []string

func TestChannels(t *testing.T) {
	msgs := []string{}
	var msg string
	links = []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	fmt.Println("Starting Test of Channels and routines")
	msg = testCheckLinksWithoutGoRoutine(t)
	msgs = append(msgs, "when we sent a list of 5 links without go routine and all of them work"+msg)

	msg = testCheckLinksWithGoRoutine(t)
	msgs = append(msgs, "when we sent a list of 5 links with go routine and all of them work"+msg)

	PrintTestsMessages(msgs)
}
func testCheckLinksWithoutGoRoutine(t *testing.T) string {
	msg := "should return 5"
	// When
	someResult := checkTheLinks(links)
	// Then
	require.Equal(t, 5, someResult)
	return msg
}
func testCheckLinksWithGoRoutine(t *testing.T) string {
	msg := "should return 5"
	// When
	someResult := checkTheLinksEnhanced(links)
	// Then
	require.Equal(t, 5, someResult)
	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
