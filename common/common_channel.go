package common

var IncomeMsgChannel chan IncomeMessage

func InitializeChannels() {
	IncomeMsgChannel = make(chan IncomeMessage)
}

func AddToIncomeMsgChannel(event IncomeMessage) {
	IncomeMsgChannel <- event
}
