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
	var statusCode = common.StatusCode(broadcastDocument["statusCode"].(string))

	existingQueueNames := broadcastDocument["queueNames"]

	var queueNames = []string{}
	for _, each := range existingQueueNames.([]interface{}) {
		queueNames = append(queueNames, each.(string))
	}

	return Broadcast{
		DocId:      GetStringValue(broadcastDocument, "docId"),
		Name:       GetStringValue(broadcastDocument, "name"),
		StatusCode: statusCode,
		QueueNames: queueNames,
	}
}

var BroadcastGnosql = gnosql_client.CollectionInput{
	CollectionName: "broadcasts",
	IndexKeys:      []string{},
}
