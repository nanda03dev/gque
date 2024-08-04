package services

import (
	"errors"
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
)

type BroadcastService interface {
	CreateBroadcast(broadcast models.Broadcast) (models.Broadcast, error)
	GetBroadcasts() ([]models.Broadcast, error)
	GetBroadcastByID(docId string) (models.Broadcast, error)
	GetBroadcastByName(broadcastName string) (models.Broadcast, error)
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
	_, err := s.GetBroadcastByName(broadcast.Name)

	if err == nil {
		return models.Broadcast{}, errors.New(global_constant.ERROR_BROAD_CAST_ALREADY_EXISTS)
	}

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

func (s *broadcastService) GetBroadcastByName(broadcastName string) (models.Broadcast, error) {
	var broadcast models.Broadcast

	filter := gnosql_client.MapInterface{
		"name": broadcastName,
	}

	result := s.broadcastGnoSQL.Filter(filter)

	if result.Error != nil {
		return models.Broadcast{}, result.Error
	}

	if len(result.Data) > 0 {
		broadcast = models.ToBroadcastModel(result.Data[0])
	} else {
		return models.Broadcast{}, errors.New(global_constant.ERROR_BROADCAST_NOT_FOUND)
	}

	return broadcast, nil

}

func (s *broadcastService) UpdateBroadcast(updateBroadcast models.Broadcast) error {
	result := s.broadcastGnoSQL.Update(updateBroadcast.DocId, updateBroadcast.ToDocument())

	return result.Error
}

func (s *broadcastService) DeleteBroadcast(docId string) error {
	result := s.broadcastGnoSQL.Delete(docId)

	return result.Error
}
