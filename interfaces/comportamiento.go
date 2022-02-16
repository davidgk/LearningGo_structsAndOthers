package interfaces

type SenseHumor struct {
	Humor string
}

type FunnyRobot struct {
	HumorStatus
}

type SadRobot struct {
	HumorStatus
}

func (s SenseHumor) RetrieveHumor() string {
	return s.Humor
}

func (bot FunnyRobot) Humor() HumorStatus {
	return bot.HumorStatus
}

func (bot SadRobot) Humor() HumorStatus {
	return bot.HumorStatus
}
