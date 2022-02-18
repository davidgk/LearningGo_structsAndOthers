package channelsAndRoutines

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
	t1 := time.Now()
	fmt.Println(t1)

	msg = testCheckLinksWithoutGoRoutine(t)
	msgs = append(msgs, "when we sent a list of 5 links without go routine and all of them work"+msg)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1), "testCheckLinksWithoutGoRoutine")

	msg = testCheckLinksWithSyncGrouping(t)
	msgs = append(msgs, "when we sent a list of 5 links with sync Groups and go routines and all of them work"+msg)
	t3 := time.Now()
	fmt.Println(t3.Sub(t2), "testCheckLinksWithSyncGrouping")

	msg = testCheckLinksWithGoRoutineChannelsWay(t)
	msgs = append(msgs, "when we sent a list of 5 links with go routine and channels all of them work"+msg)
	t4 := time.Now()
	fmt.Println(t4.Sub(t3), "testCheckLinksWithGoRoutineChannelsWay")

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

func testCheckLinksWithSyncGrouping(t *testing.T) string {
	msg := "should return 5"
	// When
	someResult := checkTheLinksEnhanced(links)
	// Then
	require.Equal(t, 5, someResult)
	return msg
}

func testCheckLinksWithGoRoutineChannelsWay(t *testing.T) string {
	msg := "should return 5"
	// When
	someResult := checkTheLinksUsingChannels(links)
	// Then
	require.Equal(t, 5, someResult)
	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println("	", m)
	}
}
