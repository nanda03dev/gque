package common

type QueueChannels map[string]chan string

type QueueName string
type MessageType string
type Document map[string]interface{}
type StatusCode string

type FilterBodyType struct {
	Key   string
	Value interface{}
}

type FiltersBodyType []FilterBodyType

type SortBodyType struct {
	Key   string `json:"Key"`
	Order int    `json:"order"`
}

type RequestFilterBodyType struct {
	ListOfFilter FiltersBodyType `json:"filters"`
	Size         int             `json:"size"`
	SortBody     SortBodyType    `json:"sort"`
}

type IncomeMessage struct {
	Name        string      `json:"name"`
	MessageType MessageType `json:"messageType"`
	Data        string      `json:"data"`
}
