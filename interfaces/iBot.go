package interfaces

type HumorStatus interface {
	RetrieveHumor() string
}

type BotHumor interface {
	HumorStatus
	Humor() HumorStatus
}

type Bot interface {
	BotHumor
	GetGreeting() string
}
