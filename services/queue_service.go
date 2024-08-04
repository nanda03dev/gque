package services

import (
	"errors"

	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
)

type QueueService interface {
	CreateQueue(queue models.Queue) (models.Queue, error)
	GetQueues() ([]models.Queue, error)
	GetQueueByID(docId string) (models.Queue, error)
	UpdateQueue(queue models.Queue) error
	DeleteQueue(docId string) error
	InitializeChannels() error
}

type queueService struct {
	queueGnoSQL *gnosql_client.Collection
}

func NewQueueService(gnosql *gnosql_client.Database) QueueService {
	return &queueService{gnosql.Collections[models.QueueGnosql.CollectionName]}
}

func (s *queueService) CreateQueue(queue models.Queue) (models.Queue, error) {
	_, err := s.GetQueueByName(queue.Name)

	if err == nil {
		return models.Queue{}, errors.New(global_constant.ERROR_QUEUE_ALREADY_EXISTS)
	}

	queue.DocId = models.Generate16DigitUUID()
	queue.StatusCode = global_constant.QUEUE_ACTIVE
	result := s.queueGnoSQL.Create(queue.ToDocument())

	s.InitializeChannel(queue)

	return queue, result.Error
}

func (s *queueService) GetQueues() ([]models.Queue, error) {

	result := s.queueGnoSQL.Filter(gnosql_client.MapInterface{})
	var queues = make([]models.Queue, 0)

	for _, document := range result.Data {
		queues = append(queues, models.ToQueueModel(document))
	}
	return queues, result.Error
}

func (s *queueService) GetQueueByID(docId string) (models.Queue, error) {
	result := s.queueGnoSQL.Read(docId)
	return models.ToQueueModel(result.Data), result.Error
}

func (s *queueService) GetQueueByName(queueName string) (models.Queue, error) {
	var queue models.Queue

	filter := gnosql_client.MapInterface{
		"name": queueName,
	}

	result := s.queueGnoSQL.Filter(filter)

	if result.Error != nil {
		return models.Queue{}, result.Error
	}

	if len(result.Data) > 0 {
		queue = models.ToQueueModel(result.Data[0])
	} else {
		return models.Queue{}, errors.New(global_constant.ERROR_QUEUE_NOT_FOUND)
	}

	return queue, nil
}

func (s *queueService) UpdateQueue(updateQueue models.Queue) error {
	result := s.queueGnoSQL.Update(updateQueue.DocId, updateQueue.ToDocument())

	return result.Error
}

func (s *queueService) DeleteQueue(docId string) error {
	queue, _ := s.GetQueueByID(docId)
	s.DeleteChannel(queue)

	result := s.queueGnoSQL.Delete(docId)

	return result.Error
}

func (s *queueService) InitializeChannels() error {

	AllQueue, getError := s.GetQueues()

	for _, value := range AllQueue {
		s.InitializeChannel(value)
	}
	return getError
}

func (s *queueService) InitializeChannel(queue models.Queue) {
	if common.QueueChannelMap[queue.Name] == nil {
		common.QueueChannelMap[queue.Name] = make(chan string)
	}
}

func (s *queueService) DeleteChannel(queue models.Queue) {
	delete(common.QueueChannelMap, queue.Name)
}
