package models

import (
	"encoding/json"

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
	entityString, _ := json.Marshal(broadcastDocument)

	var parsedEntity Broadcast
	json.Unmarshal(entityString, &parsedEntity)

	return parsedEntity
}

var BroadcastGnosql = gnosql_client.CollectionInput{
	CollectionName: "broadcasts",
	IndexKeys:      []string{},
}
