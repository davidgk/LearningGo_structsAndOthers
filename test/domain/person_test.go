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

	msg = testCreatePersonWithContactInfo(t)
	msgs = append(msgs, "when we want create a person with contact info "+msg)

	msg = testCheckReceiverMethod(t)
	msgs = append(msgs, "When created Person needs to show its data "+msg)

	msg = testWrongWayUpdateName(t)
	msgs = append(msgs, "When update name wrong way  "+msg)

	test.PrintTestsMessages(msgs)
}

func testWrongWayUpdateName(t *testing.T) string {
	msg := "should not update name"
	firstName := "Alex"
	lastName := "Anderson"
	person := d.Person{FirstName: firstName, LastName: lastName}
	person.WrongUpdateName("jimmy")

	if person.FirstName != firstName {
		t.Error("error: name shouldn't be updated")
		msg = " and FAIL"
	}
	return msg
}

func testCheckReceiverMethod(t *testing.T) string {
	msg := "should be contains a contact with email and zipcode"
	firstName := "Alex"
	lastName := "Anderson"
	email := "myEmail"
	code := 12345
	contactInfo := d.ContactInfo{Email: email, ZipCode: code}
	person := d.Person{FirstName: firstName, LastName: lastName, ContactInfo: contactInfo}
	values := person.ToStringMainData()
	expected := "Anderson, Alex - ContactInfo: Email: myEmail, Zipcode: 12345"
	if len(values) != len(expected) {
		t.Error("error: not show the correct values")
		msg = " and FAIL"
	}
	return msg
}

func testCreatePersonWithContactInfo(t *testing.T) string {
	msg := "should be contains a contact with email and zipcode"
	firstName := "Alex"
	lastName := "Anderson"
	email := "myEmail"
	code := 12345
	contactInfo := d.ContactInfo{Email: email, ZipCode: code}
	person := d.Person{FirstName: firstName, LastName: lastName, ContactInfo: contactInfo}
	if person.ContactInfo.Email != email {
		t.Errorf("error: expected email was %vand got %v", email, contactInfo.Email)
		msg = " and FAIL"
	}
	if person.ContactInfo.ZipCode != code {
		t.Errorf("error: expected zipcode was %vand got %v", code, contactInfo.ZipCode)
		msg = " and FAIL"
	}
	return msg
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
