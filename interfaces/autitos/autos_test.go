package autitos

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAutosAndBoats(t *testing.T) {
	msgs := []string{}
	var msg string
	fmt.Println("Starting Test of Autos")
	msg = testDifferentCars(t)
	msgs = append(msgs, "when we want to create cars using interfaces "+msg)

	msg = testDifferentCars2Changes(t)
	msgs = append(msgs, "when we want to Change Car data "+msg)

	msg = testDifferentBoats(t)
	msgs = append(msgs, "when we want to create boats using interfaces "+msg)

	PrintTestsMessages(msgs)
}

func testDifferentBoats(t *testing.T) string {
	msg := "should be created"
	// Given
	barracuda := Boat{Info: &Information{Type: "Barracuda", Fuel: "Diesel"}}
	// When
	barracudaResult := barracuda.description()
	// Then
	require.Equal(t, barracudaResult, "Barracuda - Diesel")

	return msg
}

func testDifferentCars2Changes(t *testing.T) string {
	msg := "should change it"
	// Given
	peugeot := Car{Info: &Information{Type: "hatchback", Fuel: "Diesel"}}
	// When
	peugeot.changeType("5Puertas")
	peugeotResult := peugeot.description()

	// Then
	require.Equal(t, peugeotResult, "5Puertas - Diesel")
	return msg
}

func testDifferentCars(t *testing.T) string {
	msg := "should be created"
	// Given
	peugeot := Car{Info: &Information{Type: "hatchback", Fuel: "Diesel"}}
	duster := Car{Info: &Information{Type: "truck", Fuel: "Nafta"}}
	// When
	peugeotResult := peugeot.description()
	dusterResult := duster.description()
	// Then
	require.Equal(t, peugeotResult, "hatchback - Diesel")
	require.Equal(t, dusterResult, "truck - Nafta")
	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
