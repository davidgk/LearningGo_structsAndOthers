package domain

import (
	"fmt"
	"github.com/stretchr/testify/require"
	d "structs/src/domain"
	"testing"
)

func TestCreateCar(t *testing.T) {
	fmt.Println("Starting Test of CreateCar")
	msg := yourTestCreateCar(t)
	msgs := []string{"when want to create a car with brand " + msg}

	// msg = yourTestCreateCarConMotor(t)
	// msgs = append(msgs, "when we want to ... "+msg)

	PrintTestsMessages(msgs)
}

func yourTestCreateCar(t *testing.T) string {
	msg := "should be created"
	brand := "Ford"
	aCar := d.Car{Brand: brand}
	require.Equal(t, brand, aCar.Brand, "the brand is different than expected")
	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
