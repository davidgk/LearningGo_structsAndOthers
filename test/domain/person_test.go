package domain

import (
	"fmt"
	d "structs/src/domain"
	"structs/test"
	"testing"
)

func TestPerson(t *testing.T) {
	fmt.Println("Starting Test of Person")
	msg := testCreatePerson01(t)
	msgs := []string{"when want to create a Person " + msg}

	msg = testCreatePerson02(t)
	msgs = append(msgs, "when we want change person names "+msg)

	test.PrintTestsMessages(msgs)
}

func testCreatePerson02(t *testing.T) string {
	msg := "should update them"
	firstName := "Alex"
	lastName := "Anderson"
	var alex d.Person
	alex.FirstName = firstName
	alex.LastName = lastName

	if alex.FirstName != firstName {
		t.Errorf("error: expected firstName was %vand got %v", firstName, alex.FirstName)
		msg = " and FAIL"
	}
	if alex.LastName != lastName {
		t.Errorf("error: expected lastName was %vand got %v", lastName, alex.LastName)
		msg = " and FAIL"
	}
	return msg
}

func testCreatePerson01(t *testing.T) string {
	msg := "should be initialized name alex anderson"
	firstName := "Alex"
	lastName := "Anderson"
	person := d.Person{FirstName: firstName, LastName: lastName}
	if person.FirstName != firstName {
		t.Errorf("error: expected firstName was %vand got %v", firstName, person.FirstName)
		msg = " and FAIL"
	}
	if person.LastName != lastName {
		t.Errorf("error: expected lastName was %vand got %v", lastName, person.LastName)
		msg = " and FAIL"
	}
	return msg
}
