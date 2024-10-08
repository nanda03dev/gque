package workers

import (
	"log"

	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/models"
	"github.com/nanda03dev/gque/services"
)

func StartIncomingMsgWorker() {
	for {
		message := <-common.IncomeMsgChannel

		var messageService = services.AppServices.Message
		var newMessage = models.Message{
			Name:        message.QueueName,
			MessageType: message.MessageType,
			Data:        message.Data,
		}

		messageCreateResult, messageCreateError := messageService.CreateMessage(newMessage)

		if messageCreateError == nil {
			MsgProducerChannel <- messageCreateResult

		} else {
			log.Printf("Message cannot processd and stored %v ", message)
		}

	}
}
