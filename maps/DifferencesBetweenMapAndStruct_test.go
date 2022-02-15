package maps

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var red string
var green string
var white string

func TestDifferenceBetweenMapAndStruct(t *testing.T) {
	msgs := []string{}
	red = "red"
	green = "green"
	white = "white"
	fmt.Println("Starting Test of DifferenceBetweenMapAndStruct ")
	msg := testCreateAMap(t)
	msgs = append(msgs, "when we want to create a map "+msg)

	msg = testRemoveValues(t)
	msgs = append(msgs, "when we want to remove values "+msg)

	msg = testIterateOverAMap(t)
	msgs = append(msgs, "when we want to get all values of map "+msg)

	msg = testAnotherAtts(t)
	msgs = append(msgs, "another atts on map "+msg)

	PrintTestsMessages(msgs)
}

func testAnotherAtts(t *testing.T) string {
	msg := "should work"
	// Given
	aMap := map[string]string{
		red:   red,
		green: green,
		white: white,
	}
	// When
	someResult := len(aMap)
	// Then
	require.Equal(t, 3, someResult)
	return msg
}

func testIterateOverAMap(t *testing.T) string {
	msg := "should get them"
	// Given
	aMap := map[string]string{
		red:   red,
		green: green,
		white: white,
	}
	// When
	someResult := getMapPrinted(aMap)
	values := strings.Split(strings.Trim(someResult, " "), " ")
	// Then
	require.Equal(t, 3, len(values))
	return msg
}

func getMapPrinted(aMap map[string]string) string {
	var result string
	for key, _ := range aMap {
		result = result + " " + aMap[key]
	}
	return result
}

func testRemoveValues(t *testing.T) string {
	msg := "should be removed"
	// Given
	aMap := map[string]string{
		red:   red,
		green: green,
	}
	// When
	delete(aMap, green)
	deletedValue := aMap[green]
	// Then
	require.Equal(t, red, aMap[red])
	require.Equal(t, deletedValue, "")
	return msg
}

func testCreateAMap(t *testing.T) string {
	msg := "should create a map of values"
	// Given
	aSecondMap := map[string]string{}
	aThirdMap := make(map[string]string)
	// When
	aMap := map[string]string{
		red:   red,
		green: green,
	}

	aSecondMap[red] = red
	aSecondMap[green] = green

	aThirdMap[red] = red
	aThirdMap[green] = green
	// Then
	require.Equal(t, red, aMap[red], "Expected red to have red")
	require.Equal(t, red, aSecondMap[red], "Expected green to have green")
	require.Equal(t, red, aThirdMap[red], "Expected green to have green")
	require.Equal(t, green, aMap[green], "Expected green to have green")
	require.Equal(t, green, aSecondMap[green], "Expected green to have green")
	require.Equal(t, green, aThirdMap[green], "Expected green to have green")

	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
