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
		log.Printf("Message processing and stored %v ", message)

		var messageService = services.AppServices.Message
		var newMessage = models.Message{
			Name:        message.QueueName,
			MessageType: message.MessageType,
			Data:        message.Data,
		}

		messageCreateResult, messageCreateError := messageService.CreateMessage(newMessage)

		log.Printf("Message stored %v ", message)

		if messageCreateError == nil {
			MsgProducerChannel <- messageCreateResult
			log.Printf("Message send to prodcer  %v ", messageCreateResult)

		} else {
			log.Printf("Message cannot processd and stored %v ", message)
		}

	}
}
