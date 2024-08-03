package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/common"
)

type Broadcast struct {
	DocId      string            `json:"docId" bson:"docId"`
	Name       string            `json:"name"`
	StatusCode common.StatusCode `json:"statusCode"`
	QueueNames []string          `json:"queueNames"`
}

func (broadcast Broadcast) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":      broadcast.DocId,
		"name":       broadcast.Name,
		"statusCode": broadcast.StatusCode,
		"queueNames": broadcast.QueueNames,
	}
}

func ToBroadcastModel(broadcastDocument gnosql_client.Document) Broadcast {
	return Broadcast{
		DocId:      GetStringValue(broadcastDocument, "docId"),
		Name:       GetStringValue(broadcastDocument, "name"),
		StatusCode: GetValue[common.StatusCode](broadcastDocument, "statusCode"),
		QueueNames: GetValue[[]string](broadcastDocument, "queueNames"),
	}
}

var BroadcastGnosql = gnosql_client.CollectionInput{
	CollectionName: "broadcasts",
	IndexKeys:      []string{},
}
