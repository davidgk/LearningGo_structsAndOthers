package robotitos

type SpanishBot struct {
	BotHumor
}

func (SpanishBot) GetGreeting() string {
	return "Hola"
}
