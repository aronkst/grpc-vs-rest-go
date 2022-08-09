package message

import (
	context "context"
	"fmt"

	"github.com/aronkst/grpc-vs-rest-go/data"
)

type Server struct {
	UnimplementedMessageServiceServer
}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	data := data.Data{
		Value1: message.Value1,
		Value2: message.Value2,
		Value3: message.Value3,
		Value4: message.Value4,
		Value5: message.Value5,
	}

	fmt.Printf("data (GRPC Server): %v\n", data)

	messageOutput := &Message{
		Value1: data.Value1,
		Value2: data.Value2,
		Value3: data.Value3,
		Value4: data.Value4,
		Value5: data.Value5,
	}

	return messageOutput, nil
}
