package workers

import (
	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/models"
	"github.com/nanda03dev/gque/services"
)

var MsgProducerChannel chan models.Message = make(chan models.Message)

func StartMsgProducerWorker() {
	for {
		message := <-MsgProducerChannel

		var messageService = services.AppServices.Message

		var queuesAndMessages = GenerateMessages(message)

		for queueName, messagesToSend := range queuesAndMessages {
			channel := common.QueueChannelMap[queueName]

			if channel != nil {
				for _, messageToSend := range messagesToSend {
					channel <- messageToSend
				}
			}
		}
		messageService.UpdateToPublishedMessage(message)
	}
}
