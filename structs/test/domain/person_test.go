package domain

import (
	"fmt"
	d "structs/src/domain"
	"structs/test"
	"testing"
)

func TestPersonForUpdate(t *testing.T) {
	fmt.Println("Starting Test of Person for update")
	msg := testCheckReceiverMethod(t)
	msgs := []string{"When created Person needs to show its data " + msg}

	msg = testWrongWayUpdateName(t)
	msgs = append(msgs, "When update name wrong way  "+msg)

	msg = testCorrectWayUpdateNameUsingPointer(t)
	msgs = append(msgs, "When update name correct way with pointers  "+msg)

	msg = testCorrectWayUpdateNameUsingOriginalObject(t)
	msgs = append(msgs, "When update name correct way with original Object  "+msg)

	msg = testWayUpdateContact(t)
	msgs = append(msgs, "When update the contact with correct way with pointers  "+msg)

	msg = testWaySlicesAreDifferentFromStructs(t)
	msgs = append(msgs, "When update one slice without pointers "+msg)
	test.PrintTestsMessages(msgs)
}

func TestPersonForCreate(t *testing.T) {
	fmt.Println("Starting Test of Person for create")
	msg := testCreatePerson01(t)
	msgs := []string{"when want to create a Person " + msg}

	msg = testCreatePerson02(t)
	msgs = append(msgs, "when we want change person names "+msg)

	test.PrintTestsMessages(msgs)
}

func testWaySlicesAreDifferentFromStructs(t *testing.T) string {
	msg := "should update it"
	someList := []string{"carlos", "julian"}
	david := "David"
	updateFirstPositionLists(someList, david)
	if someList[0] != david {
		t.Error("error: list should be updated")
		msg = " and FAIL"
	}
	return msg
}

func updateFirstPositionLists(list []string, david string) {
	list[0] = david
}

func testCorrectWayUpdateNameUsingOriginalObject(t *testing.T) string {
	msg := "should update the name"
	firstName := "Alex"
	lastName := "Anderson"
	person := d.Person{FirstName: firstName, LastName: lastName}
	updatedName := "jimmy"
	person.CorrectUpdateName(updatedName)
	if person.FirstName != updatedName {
		t.Error("error: name should be updated")
		msg = " and FAIL"
	}
	return msg
}

func testWayUpdateContact(t *testing.T) string {
	msg := "should update the contact"
	firstName := "Alex"
	lastName := "Anderson"
	info := d.ContactInfo{Email: "mail", ZipCode: 12345}
	infoUpdated := d.ContactInfo{Email: "amail", ZipCode: 67891}
	person := d.Person{FirstName: firstName, LastName: lastName, ContactInfo: info}
	person.CorrectUpdateContact(infoUpdated)
	//fmt.Printf("%+v", person)
	if (person.Email != infoUpdated.Email) || (person.ZipCode != infoUpdated.ZipCode) {
		t.Error("error: contact should be updated")
		msg = " and FAIL"
	}
	return msg
}

func testCorrectWayUpdateNameUsingPointer(t *testing.T) string {
	msg := "should update the name"
	firstName := "Alex"
	lastName := "Anderson"
	person := d.Person{FirstName: firstName, LastName: lastName}
	updatedName := "jimmy"
	jimPointer := &person
	jimPointer.CorrectUpdateName(updatedName)
	//fmt.Printf("%+v", person)
	if person.FirstName != updatedName {
		t.Error("error: name should be updated")
		msg = " and FAIL"
	}
	return msg
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
