package robotitos

type EnglishBot struct {
	BotHumor
}

func (EnglishBot) GetGreeting() string {
	return "Hi there"
}
