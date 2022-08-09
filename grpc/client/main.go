package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/aronkst/grpc-vs-rest-go/data"
	"github.com/aronkst/grpc-vs-rest-go/grpc/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	countString := flag.String("count", "", "")
	flag.Parse()

	countInt, err := strconv.Atoi(*countString)
	if err != nil {
		log.Fatal(err)
	}

	count := int32(countInt)

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := message.NewMessageServiceClient(conn)

	for i := int32(1); i <= count; i++ {
		dataInput := data.Data{
			Value1: i,
			Value2: i,
			Value3: i,
			Value4: i,
			Value5: i,
		}
		message := &message.Message{
			Value1: dataInput.Value1,
			Value2: dataInput.Value2,
			Value3: dataInput.Value3,
			Value4: dataInput.Value4,
			Value5: dataInput.Value5,
		}

		res, err := client.SendMessage(context.Background(), message)
		if err != nil {
			log.Fatal(err)
		}

		data := data.Data{
			Value1: res.Value1,
			Value2: res.Value2,
			Value3: res.Value3,
			Value4: res.Value4,
			Value5: res.Value5,
		}

		fmt.Printf("data (GRPC Client): %v, i: %d\n", data, i)
	}
}
