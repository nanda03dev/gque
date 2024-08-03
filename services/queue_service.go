package services

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
)

type QueueService interface {
	CreateQueue(queue models.Queue) (models.Queue, error)
	GetQueues() ([]models.Queue, error)
	GetQueueByID(docId string) (models.Queue, error)
	UpdateQueue(queue models.Queue) error
	DeleteQueue(docId string) error
}

type queueService struct {
	queueGnoSQL *gnosql_client.Collection
}

func NewQueueService(gnosql *gnosql_client.Database) QueueService {
	return &queueService{gnosql.Collections[models.QueueGnosql.CollectionName]}
}

func (s *queueService) CreateQueue(queue models.Queue) (models.Queue, error) {
	queue.DocId = models.Generate16DigitUUID()
	queue.StatusCode = global_constant.QUEUE_ACTIVE
	result := s.queueGnoSQL.Create(queue.ToDocument())

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

func (s *queueService) UpdateQueue(updateQueue models.Queue) error {
	result := s.queueGnoSQL.Update(updateQueue.DocId, updateQueue.ToDocument())

	return result.Error
}

func (s *queueService) DeleteQueue(docId string) error {
	result := s.queueGnoSQL.Delete(docId)

	return result.Error
}
