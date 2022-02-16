package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTranslatorBots(t *testing.T) {
	msgs := []string{}
	var msg string
	fmt.Println("Starting Test of ")
	msg = testEnglishBot(t)
	msgs = append(msgs, "when we want to test english bot "+msg)

	msg = testSpanishBot(t)
	msgs = append(msgs, "when we want to test spanish bot  "+msg)

	msg = testTranslatorInit(t)
	msgs = append(msgs, "when we want to test the translator  "+msg)

	PrintTestsMessages(msgs)
}

func testTranslatorInit(t *testing.T) string {
	msg := "should translator prepare greetings"
	// Given
	botEng := EnglishBot{}
	botSp := SpanishBot{}
	// When
	someResultEng := PrepareGreeting(botEng)
	someResultSp := PrepareGreeting(botSp)
	// Then
	require.Equal(t, "Hey!Hi there", someResultEng)
	require.Equal(t, "Hey!Hola", someResultSp)

	return msg
}

func testEnglishBot(t *testing.T) string {
	msg := "should obtain greetings"
	// Given
	bot := EnglishBot{}
	// when
	result01 := bot.GetGreeting()
	// Then
	require.Equal(t, "Hi there", result01)
	return msg
}

func testSpanishBot(t *testing.T) string {
	msg := "should obtain greetings"
	// Given
	bot := SpanishBot{}
	// when
	result01 := bot.GetGreeting()
	// Then
	require.Equal(t, "Hola", result01)
	return msg
}

func PrintTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
