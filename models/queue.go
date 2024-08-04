package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque/common"
)

type Queue struct {
	DocId       string            `json:"docId" bson:"docId"`
	Name        string            `json:"name"`
	Time        int64             `json:"time"`
	BroadcastId string            `json:"broadcastId"`
	StatusCode  common.StatusCode `json:"statusCode"`
}

func (queue Queue) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":       queue.DocId,
		"name":        queue.Name,
		"time":        queue.Time,
		"broadcastId": queue.BroadcastId,
		"statusCode":  queue.StatusCode,
	}
}

func ToQueueModel(queueDocument gnosql_client.Document) Queue {
	var statusCode = common.StatusCode(queueDocument["statusCode"].(string))

	return Queue{
		DocId: GetStringValue(queueDocument, "docId"),
		Name:  GetStringValue(queueDocument, "name"),
		// Time:        time,
		BroadcastId: GetStringValue(queueDocument, "broadcastId"),
		StatusCode:  statusCode,
	}
}

var QueueGnosql = gnosql_client.CollectionInput{
	CollectionName: "queues",
	IndexKeys:      []string{"broadcastId"},
}
