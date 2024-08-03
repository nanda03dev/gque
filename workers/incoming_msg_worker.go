package workers

import (
	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/models"
	"github.com/nanda03dev/gque/services"
)

func StartIncomingMsgWorker() {
	for {
		message := <-common.IncomeMsgChannel

		var messageService = services.AppServices.Message
		var newMessage = models.Message{
			Name:        message.Name,
			MessageType: message.MessageType,
			Data:        message.Data,
		}
		messageService.CreateMessage(newMessage)
	}
}
