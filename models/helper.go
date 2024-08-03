package models

import (
	"github.com/google/uuid"
	"github.com/nanda03dev/gnosql_client"
)

func Generate16DigitUUID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func GetAllGnosqlCollections() []gnosql_client.CollectionInput {
	return []gnosql_client.CollectionInput{
		MessageGnosql,
		QueueGnosql,
		BroadcastGnosql,
	}
}

func GetStringValue(document gnosql_client.Document, key string) string {
	value := document[key]
	return value.(string)
}
func GetIntegerValue(document gnosql_client.Document, key string) int {
	value := document[key]
	return value.(int)
}
func GetBoolValue(document gnosql_client.Document, key string) bool {
	value := document[key]
	return value.(bool)
}
func GetValue[T any](document gnosql_client.Document, key string) T {
	value := document[key]
	return value.(T)
}
