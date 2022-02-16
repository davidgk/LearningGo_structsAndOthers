package interfaces

type EnglishBot struct {
}

func (EnglishBot) GetGreeting() string {
	return "Hi there"
}
