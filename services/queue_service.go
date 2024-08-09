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
	existingEntity, _ := s.GetQueueByName(queue.Name)

	if len(existingEntity.DocId) > 1 {
		return models.Queue{}, errors.New(global_constant.ERROR_QUEUE_ALREADY_EXISTS)
	}

	queue.DocId = models.Generate16DigitUUID()
	queue.StatusCode = global_constant.QUEUE_ACTIVE

	_, err := s.queueGnoSQL.Create(queue.ToDocument())

	if err == nil {
		s.InitializeChannel(queue)
	}

	return queue, err
}

func (s *queueService) GetQueues() ([]models.Queue, error) {

	result, err := s.queueGnoSQL.Filter(gnosql_client.MapInterface{})
	var queues = make([]models.Queue, 0)

	if err != nil {
		return queues, err
	}

	for _, document := range result.Data {
		queues = append(queues, models.ToQueueModel(document))
	}
	return queues, err
}

func (s *queueService) GetQueueByID(docId string) (models.Queue, error) {
	result, err := s.queueGnoSQL.Read(docId)
	return models.ToQueueModel(result.Data), err
}

func (s *queueService) GetQueueByName(queueName string) (models.Queue, error) {
	var queue models.Queue

	filter := gnosql_client.MapInterface{
		"name": queueName,
	}

	result, err := s.queueGnoSQL.Filter(filter)

	if err != nil {
		return models.Queue{}, err
	}

	if len(result.Data) > 0 {
		queue = models.ToQueueModel(result.Data[0])
	} else {
		return models.Queue{}, errors.New(global_constant.ERROR_QUEUE_NOT_FOUND)
	}

	return queue, nil
}

func (s *queueService) UpdateQueue(updateQueue models.Queue) error {
	_, err := s.queueGnoSQL.Update(updateQueue.DocId, updateQueue.ToDocument())

	return err
}

func (s *queueService) DeleteQueue(docId string) error {
	queue, _ := s.GetQueueByID(docId)
	s.DeleteChannel(queue)

	_, err := s.queueGnoSQL.Delete(docId)

	return err
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
		common.QueueChannelMap[queue.Name] = make(chan string, global_constant.COMMON_QUEUE_SIZE)
	}
}

func (s *queueService) DeleteChannel(queue models.Queue) {
	delete(common.QueueChannelMap, queue.Name)
}
