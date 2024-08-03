package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/common"
)

type Message struct {
	DocId       string             `json:"docId" bson:"docId"`
	Name        string             `json:"name"`
	MessageType common.MessageType `json:"messageType"`
	Data        string             `json:"data"`
	StatusCode  common.StatusCode  `json:"statusCode"`
}

func (message Message) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":       message.DocId,
		"name":        message.Name,
		"messageType": message.MessageType,
		"data":        message.Data,
		"statusCode":  message.StatusCode,
	}
}

func ToMessageModel(messageDocument gnosql_client.Document) Message {
	return Message{
		DocId:       GetStringValue(messageDocument, "docId"),
		Name:        GetStringValue(messageDocument, "name"),
		MessageType: GetValue[common.MessageType](messageDocument, "messageType"),
		Data:        GetStringValue(messageDocument, "data"),
		StatusCode:  GetValue[common.StatusCode](messageDocument, "statusCode"),
	}
}

var MessageGnosql = gnosql_client.CollectionInput{
	CollectionName: "messages",
	IndexKeys:      []string{"statusCode"},
}
