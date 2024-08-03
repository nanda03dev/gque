package common

var IncomeMsgChannel chan IncomeMessage
var QueueChannelMap QueueChannels

func InitializeChannels() {

	IncomeMsgChannel = make(chan IncomeMessage)
	QueueChannelMap = make(QueueChannels)
}

func AddToIncomeMsgChannel(event IncomeMessage) {
	IncomeMsgChannel <- event
}
