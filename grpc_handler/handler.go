package grpc_handler

import (
	"context"

	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/global_constant"
	"github.com/nanda03dev/gque/models"
	pb "github.com/nanda03dev/gque/proto"
	"github.com/nanda03dev/gque/services"
)

type GqueServer struct {
	pb.UnimplementedGqueServiceServer

	Services services.Services
}

func (gRPC *GqueServer) CreateQueue(ctx context.Context,
	req *pb.QueueCreateRequest) (*pb.SuccessResponse, error) {

	var newQueue = models.Queue{
		Name: req.QueueName,
		Time: req.Time,
	}
	result, err := gRPC.Services.Queue.CreateQueue(newQueue)

	var response = &pb.SuccessResponse{
		Data: result.DocId,
	}

	return response, err
}

func (gRPC *GqueServer) CreateBroadcast(ctx context.Context,
	req *pb.BroadcastCreateRequest) (*pb.SuccessResponse, error) {

	var response = &pb.SuccessResponse{}

	var newBroadcast = models.Broadcast{
		Name:       req.BroadcastName,
		QueueNames: req.QueueNames,
	}

	result, err := gRPC.Services.Broadcast.CreateBroadcast(newBroadcast)

	response.Data = result.DocId

	return response, err
}

func (gRPC *GqueServer) PushMessage(ctx context.Context,
	req *pb.PushMessageRequest) (*pb.SuccessResponse, error) {
	var response = &pb.SuccessResponse{}

	var newMessage = common.IncomeMessage{
		QueueName:   req.QueueName,
		MessageType: global_constant.MESSAGE_TYPE_QUEUE,
		Data:        req.Message,
	}

	common.AddToIncomeMsgChannel(newMessage)

	response.Data = global_constant.SUCCESS_MSG_PUSH

	return response, nil
}

func (gRPC *GqueServer) BroadcastMessage(ctx context.Context,
	req *pb.BroadcastMessageRequest) (*pb.SuccessResponse, error) {
	var response = &pb.SuccessResponse{}

	var newMessage = common.IncomeMessage{
		QueueName:   req.BroadcastName,
		MessageType: global_constant.MESSAGE_TYPE_BROADCAST,
		Data:        req.Message,
	}

	common.IncomeMsgChannel <- newMessage

	response.Data = global_constant.SUCCESS_MSG_PUSH

	return response, nil
}

func (gRPC *GqueServer) ConsumeQueueMessages(req *pb.ConsumerRequest, stream pb.GqueService_ConsumeQueueMessagesServer) error {
	queueChan, err := services.GetQueueChannel(req.QueueName)
	if err != nil {
		stream.Context().Done()
		return nil
	}

	for {
		select {
		case <-stream.Context().Done():
			// Client has disconnected
			return nil
		case message := <-queueChan:
			if err := stream.Send(&pb.ConsumerMessage{
				Message: message,
			}); err != nil {
				return err
			}
		}
	}
}
