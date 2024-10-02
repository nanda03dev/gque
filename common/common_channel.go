package common

var IncomeMsgChannel chan IncomeMessage
var QueueChannelMap QueueChannels

func InitializeChannels() {
	IncomeMsgChannel = make(chan IncomeMessage, 1000000)
	QueueChannelMap = make(QueueChannels, 1000000)
}

func AddToIncomeMsgChannel(event IncomeMessage) {
	IncomeMsgChannel <- event
}
