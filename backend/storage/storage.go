package storage

import pb "github.com/posttul/grpc-lot-service/backend/protos"

// Service interface for a storage backend
type Service interface {
	GetLots() ([]pb.Lot, error)
	GetLotByID(string int64) (*pb.Lot, error)
}
