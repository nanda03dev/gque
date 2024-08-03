package services

import "github.com/nanda03dev/gque/config"

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
	return AppServices
}
