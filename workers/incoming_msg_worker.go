package workers

import (
	"fmt"

	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
	"github.com/nanda03dev/gque/services"
)

func StartIncomingMsgWorker() {
	for {
		message := <-common.IncomeMsgChannel
		fmt.Printf(" Msg arrived at worker %v ", message)

		var messageService = services.AppServices.Message
		var newMessage = models.Message{
			Name:        message.Name,
			MessageType: message.MessageType,
			Data:        message.Data,
		}

		messageService.CreateMessage(newMessage)

		fmt.Printf(" Msg created at worker %v ", message)

		if newMessage.MessageType == global_constant.MESSAGE_TYPE_QUEUE {
			channel := common.QueueChannelMap[newMessage.Name]

			if channel != nil {
				channel <- newMessage.Data
				fmt.Printf(" Msg send to queu chan at worker %v ", message)

			}
		}
	}
}
