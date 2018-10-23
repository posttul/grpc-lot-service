package postgres

import (
	"database/sql"
	"fmt" // Handle postgres stuff

	_ "github.com/lib/pq"
	pb "github.com/posttul/grpc-lot-service/backend/protos"
	"github.com/posttul/grpc-lot-service/backend/storage"
)

// Postgres is use to storage data to postgres
type postgres struct {
	db *sql.DB
}

// New returns a new postgres storage
func New(user, password, dbName, host string) (storage.Service, error) {
	fmt.Printf("Starting a storage.Service with postgres\n")
	cnn, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbName))
	if err != nil {
		return nil, err
	}
	err = cnn.Ping()
	if err != nil {
		return nil, err
	}
	return &postgres{
		db: cnn,
	}, nil
}

// GetLots returns all lots
func (p *postgres) GetLots() (*pb.Lots, error) {
	st, err := p.db.Query("SELECT id,name,address,country FROM lot;")
	if err != nil {
		return nil, err
	}
	lots := []*pb.Lot{}
	for st.Next() {
		lot := &pb.Lot{}
		if err := st.Scan(
			&lot.ID,
			&lot.Name,
			&lot.Address,
			&lot.Country); err != nil {
			return nil, err
		}
		lots = append(lots, lot)
	}
	return &pb.Lots{Lots: lots}, nil
}

// GetLots returns all lots
func (p *postgres) GetLotByID(id int64) (*pb.Lot, error) {
	lot := pb.Lot{}
	err := p.db.QueryRow("SELECT id,name,address,country FROM lot WHERE id=$1;", id).
		Scan(
			&lot.ID,
			&lot.Name,
			&lot.Address,
			&lot.Country)
	if err != nil {
		return nil, err
	}
	return &lot, nil
}
