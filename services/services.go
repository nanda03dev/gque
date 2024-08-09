package services

import (
	"errors"

	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/config"
	"github.com/nanda03dev/gque/global_constant"
)

type Services struct {
	Message   MessageService
	Queue     QueueService
	Broadcast BroadcastService
}

var AppServices Services

func InitializeServices() Services {
	AppServices = Services{
		Message:   NewMessageService(config.GnoSQLDB),
		Queue:     NewQueueService(config.GnoSQLDB),
		Broadcast: NewBroadcastService(config.GnoSQLDB),
	}

	AppServices.Queue.InitializeChannels()

	return AppServices
}

func GetQueueChannel(queueName string) (chan string, error) {
	queueChan := common.QueueChannelMap[queueName]
	if queueChan != nil {
		return queueChan, nil
	}
	return nil, errors.New(global_constant.ERROR_QUEUE_NOT_FOUND)
}
