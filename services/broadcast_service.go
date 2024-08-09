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
	existingEntity, _ := s.GetBroadcastByName(broadcast.Name)

	if len(existingEntity.DocId) > 1 {
		return models.Broadcast{}, errors.New(global_constant.ERROR_BROAD_CAST_ALREADY_EXISTS)
	}

	broadcast.DocId = models.Generate16DigitUUID()
	broadcast.StatusCode = global_constant.BROADCAST_ACTIVE
	_, err := s.broadcastGnoSQL.Create(broadcast.ToDocument())

	return broadcast, err
}

func (s *broadcastService) GetBroadcasts() ([]models.Broadcast, error) {

	result, err := s.broadcastGnoSQL.Filter(gnosql_client.MapInterface{})
	var broadcasts = make([]models.Broadcast, 0)

	if err != nil {
		return broadcasts, err
	}

	for _, document := range result.Data {
		broadcasts = append(broadcasts, models.ToBroadcastModel(document))
	}

	return broadcasts, nil
}

func (s *broadcastService) GetBroadcastByID(docId string) (models.Broadcast, error) {
	result, err := s.broadcastGnoSQL.Read(docId)
	return models.ToBroadcastModel(result.Data), err
}

func (s *broadcastService) GetBroadcastByName(broadcastName string) (models.Broadcast, error) {
	var broadcast models.Broadcast

	filter := gnosql_client.MapInterface{
		"name": broadcastName,
	}

	result, err := s.broadcastGnoSQL.Filter(filter)

	if err != nil {
		return models.Broadcast{}, err
	}

	if len(result.Data) > 0 {
		broadcast = models.ToBroadcastModel(result.Data[0])
	} else {
		return models.Broadcast{}, errors.New(global_constant.ERROR_BROADCAST_NOT_FOUND)
	}

	return broadcast, nil

}

func (s *broadcastService) UpdateBroadcast(updateBroadcast models.Broadcast) error {
	_, err := s.broadcastGnoSQL.Update(updateBroadcast.DocId, updateBroadcast.ToDocument())

	return err
}

func (s *broadcastService) DeleteBroadcast(docId string) error {
	_, err := s.broadcastGnoSQL.Delete(docId)

	return err
}
