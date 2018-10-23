package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/posttul/grpc-lot-service/backend/protos"
	"github.com/posttul/grpc-lot-service/backend/storage"
	"github.com/posttul/grpc-lot-service/backend/storage/postgres"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

// Server of the current lotus backend
type Server struct {
	storage storage.Service
}

// GetLotByID use this to get the lot form the server storage
func (s *Server) GetLotByID(context context.Context, in *pb.Lot) (*pb.Lot, error) {
	logrus.Infof("Get lot by id %d", in.GetID())
	return s.storage.GetLotByID(in.GetID())
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	sto, err := postgres.New("postgres", "", "lotes", "127.0.0.1")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLotusServer(grpcServer, &Server{
		storage: sto,
	})
	reflection.Register(grpcServer)
	logrus.Infof(" Serving gPRC on the port %s", port)
	grpcServer.Serve(listen)
}
