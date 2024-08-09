package workers

import (
	"log"

	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
	"github.com/nanda03dev/gque/services"
)

func GenerateMessages(message models.Message) map[string][]string {
	var result = make(map[string][]string)
	queueName := message.Name

	if message.MessageType == global_constant.MESSAGE_TYPE_QUEUE {
		result[queueName] = append(result[queueName], message.Data)
	}

	if message.MessageType == global_constant.MESSAGE_TYPE_BROADCAST {
		broadcast, _ := services.AppServices.Broadcast.GetBroadcastByName(queueName)

		if len(broadcast.DocId) < 1 {
			log.Printf("Error %v : %v ", broadcast.Name, global_constant.ERROR_BROADCAST_NOT_FOUND)
		}

		for _, queueName := range broadcast.QueueNames {
			result[queueName] = append(result[queueName], message.Data)
		}
	}

	return result
}
