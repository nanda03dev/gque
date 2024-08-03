package global_constant

import "github.com/nanda03dev/gque/common"

const (
	MESSAGE_TYPE_QUEUE     common.MessageType = "MESSAGE_TYPE_QUEUE"
	MESSAGE_TYPE_BROADCAST common.MessageType = "MESSAGE_TYPE_BROADCAST"
)

const (
	STATUS_CODE_UNPUBLISHED common.StatusCode = "200"
	STATUS_CODE_PUBLISHED   common.StatusCode = "201"
)

const (
	BROADCAST_ACTIVE   common.StatusCode = "300"
	BROADCAST_INACTIVE common.StatusCode = "305"
)

const (
	QUEUE_ACTIVE   common.StatusCode = "400"
	QUEUE_INACTIVE common.StatusCode = "405"
)

const (
	SUCCESS_MSG_PUSH = "Message successfully pushed"
)

const (
	ERROR_WHILE_BINDING_JSON   = "Request JSON binding failed"
	ERROR_WHILE_UNMARSHAL_JSON = "Request JSON Unmarhsall failed"
	ERROR_WHILE_MARSHAL_JSON   = "Request JSON Marhsall failed"
)
