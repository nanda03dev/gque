package services

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
)

type BroadcastService interface {
	CreateBroadcast(broadcast models.Broadcast) (models.Broadcast, error)
	GetBroadcasts() ([]models.Broadcast, error)
	GetBroadcastByID(docId string) (models.Broadcast, error)
	UpdateBroadcast(broadcast models.Broadcast) error
	DeleteBroadcast(docId string) error
}

type broadcastService struct {
	broadcastGnoSQL *gnosql_client.Collection
}

func NewBroadcastService(gnosql *gnosql_client.Database) BroadcastService {
	return &broadcastService{gnosql.Collections[models.BroadcastGnosql.CollectionName]}
}

func (s *broadcastService) CreateBroadcast(broadcast models.Broadcast) (models.Broadcast, error) {
	broadcast.DocId = models.Generate16DigitUUID()
	broadcast.StatusCode = global_constant.BROADCAST_ACTIVE
	result := s.broadcastGnoSQL.Create(broadcast.ToDocument())

	return broadcast, result.Error
}

func (s *broadcastService) GetBroadcasts() ([]models.Broadcast, error) {

	result := s.broadcastGnoSQL.Filter(gnosql_client.MapInterface{})
	var broadcasts = make([]models.Broadcast, 0)

	for _, document := range result.Data {
		broadcasts = append(broadcasts, models.ToBroadcastModel(document))
	}

	return broadcasts, result.Error
}

func (s *broadcastService) GetBroadcastByID(docId string) (models.Broadcast, error) {
	result := s.broadcastGnoSQL.Read(docId)
	return models.ToBroadcastModel(result.Data), result.Error
}

func (s *broadcastService) UpdateBroadcast(updateBroadcast models.Broadcast) error {
	result := s.broadcastGnoSQL.Update(updateBroadcast.DocId, updateBroadcast.ToDocument())

	return result.Error
}

func (s *broadcastService) DeleteBroadcast(docId string) error {
	result := s.broadcastGnoSQL.Delete(docId)

	return result.Error
}
