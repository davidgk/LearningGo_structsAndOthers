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

	msg = testHumorBots(t)
	msgs = append(msgs, "when we want create bots With Humor  "+msg)

	PrintTestsMessages(msgs)
}
func testHumorBots(t *testing.T) string {
	msg := "should answer they humor"
	// Given
	sadness := SenseHumor{Humor: "Sad"}
	happiness := SenseHumor{Humor: "Happy"}

	sadRobot := SadRobot{HumorStatus: sadness}
	funnyRobot := FunnyRobot{HumorStatus: happiness}

	engBot := EnglishBot{BotHumor: funnyRobot}
	spanBot := SpanishBot{BotHumor: sadRobot}

	// When
	humorEnglish := engBot.RetrieveHumor()
	humorSpanish := spanBot.RetrieveHumor()
	// Then
	require.Equal(t, "Sad", humorSpanish, "Expected false but was true")
	require.Equal(t, "Happy", humorEnglish, "Expected false but was true")
	return msg
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
