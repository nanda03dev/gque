package services

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
)

type MessageService interface {
	CreateMessage(message models.Message) (models.Message, error)
	GetMessages() ([]models.Message, error)
	GetMessageByID(docId string) (models.Message, error)
	UpdateMessage(message models.Message) error
	UpdateToPublishedMessage(message models.Message) error
	DeleteMessage(docId string) error
}

type messageService struct {
	messageGnoSQL *gnosql_client.Collection
}

func NewMessageService(gnosql *gnosql_client.Database) MessageService {
	return &messageService{gnosql.Collections[models.MessageGnosql.CollectionName]}
}

func (s *messageService) CreateMessage(message models.Message) (models.Message, error) {
	message.DocId = models.Generate16DigitUUID()
	message.StatusCode = global_constant.STATUS_CODE_UNPUBLISHED
	_, err := s.messageGnoSQL.Create(message.ToDocument())

	return message, err
}

func (s *messageService) GetMessages() ([]models.Message, error) {
	var limit int = 1000

	var filter gnosql_client.MapInterface = gnosql_client.MapInterface{
		"statusCode": global_constant.STATUS_CODE_UNPUBLISHED,
		"limit":      limit,
	}
	result, err := s.messageGnoSQL.Filter(filter)
	var messages = make([]models.Message, 0)

	if err != nil {
		return messages, err
	}

	for _, document := range result.Data {
		messages = append(messages, models.ToMessageModel(document))
	}

	return messages, nil
}

func (s *messageService) GetMessageByID(docId string) (models.Message, error) {
	result, err := s.messageGnoSQL.Read(docId)
	return models.ToMessageModel(result.Data), err
}

func (s *messageService) UpdateMessage(updateMessage models.Message) error {
	_, err := s.messageGnoSQL.Update(updateMessage.DocId, updateMessage.ToDocument())

	return err
}

func (s *messageService) UpdateToPublishedMessage(updateMessage models.Message) error {
	updateMessage.StatusCode = global_constant.STATUS_CODE_PUBLISHED
	_, err := s.messageGnoSQL.Update(updateMessage.DocId, updateMessage.ToDocument())

	return err
}

func (s *messageService) DeleteMessage(docId string) error {
	_, err := s.messageGnoSQL.Delete(docId)

	return err
}
