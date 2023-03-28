package main

import (
	"Blockchain/blockchain/blockchainpb"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	blockchainpb.BlockchainServer
	db *badger.DB
}

func NewServer(db *badger.DB) *Server {
	return &Server{db: db}
}

func main() {
	fmt.Println("Server started...")

	db, _ := badger.Open(badger.DefaultOptions("/Users/aless/Desktop/Go/Blockchain/DB"))
	defer db.Close()

	server := NewServer(db)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blockchainpb.RegisterBlockchainServer(s, server)

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
