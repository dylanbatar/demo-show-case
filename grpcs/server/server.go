package main

import (
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"
	pb "github.com/dylanbatar/demo-show-case/grpcs/proto"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:50051"

type Server struct {
	pb.UploadServiceServer
}

func (s *Server) Upload(in *pb.PersonRequest, stream pb.UploadService_UploadServer) error {
	for i, person := range in.Persons {
		stream.Send(&pb.PersonResponse{Processed: spew.Sprintf("%d: %v processed", i, person.Name)})
	}

	stream.Send(&pb.PersonResponse{Processed: spew.Sprintf("total processed %d", len(in.Persons))})
	return nil
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("error with tcp connection %v", err)
	}

	log.Printf("server listening %s\n", addr)

	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(20 * 1024 * 1024),
	)
	pb.RegisterUploadServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
