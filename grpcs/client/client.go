package main

import (
	"context"
	"io"
	"log"
	"os"

	pb "github.com/dylanbatar/demo-show-case/grpcs/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

const addr = "0.0.0.0:50051"

func readJson() *pb.PersonRequest {
	file, _ := os.ReadFile("./grpcs/client/persons.json")
	persons := pb.PersonRequest{}
	err := protojson.Unmarshal(file, &persons)

	if err != nil {
		log.Fatalln("error reading file", err)
	}
	return &persons
}

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("error with dial", err)
	}

	defer conn.Close()

	client := pb.NewUploadServiceClient(conn)

	//stream, err := client.Upload(context.Background(), &pb.PersonRequest{Person: []*pb.Person{
	//	{
	//		Name:      "test 1",
	//		Languages: "spanish",
	//		Bio:       "nose",
	//		Version:   1,
	//	},
	//	{
	//		Name:      "test 2",
	//		Languages: "spanish",
	//		Bio:       "nose",
	//		Version:   1,
	//	},
	//}})

	stream, err := client.Upload(context.Background(), readJson())

	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(msg.Processed)
	}

}
