package main

import (
	"Blockchain/blockchain/blockchainpb"
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestAddVideoGame(t *testing.T) {
	// Mock database
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Server instance
	server := &Server{dbGames: db}

	req := &blockchainpb.AddVideoGameRequest{
		Title:       "Test Game",
		Publisher:   "Test Publisher",
		ReleaseDate: "01-01-2023",
	}

	got, err := server.AddVideoGame(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	if got.Title != req.Title {
		t.Errorf("Expected title %q, got %q", req.Title, got.Title)
	}
	if got.Publisher != req.Publisher {
		t.Errorf("Expected publisher %q, got %q", req.Publisher, got.Publisher)
	}
	if got.ReleaseDate != req.ReleaseDate {
		t.Errorf("Expected price %v, got %v", req.ReleaseDate, got.ReleaseDate)
	}
	if got.Id == "" {
		t.Error("Expected non-empty ID")
	}
}

/*
func TestAddVideoGame(t *testing.T) {
	// Mock database
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Server instance
	server := &Server{dbGames: db}

	want := &blockchainpb.AddVideoGameRequest{
		Title:       "Test Game",
		Publisher:   "Test Publisher",
		ReleaseDate: "01-01-2023",
	}
	got, err := server.AddVideoGame(context.Background(), want)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
*/

func TestGetVideoGame(t *testing.T) {

	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	server := &Server{dbGames: db}

	game := &blockchainpb.VideoGame{
		Id:          "test-id",
		Title:       "Test Title",
		Publisher:   "Test publisher",
		ReleaseDate: "01-01-2023",
	}

	gameData, err := proto.Marshal(game)
	if err != nil {
		t.Fatalf("Failed to marshal game: %v", err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(game.Id), gameData)
	})
	if err != nil {
		t.Fatalf("Failed to store game: %v", err)
	}

	req := &blockchainpb.GetVideoGameRequest{GameId: game.Id}
	res, err := server.GetVideoGame(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to get game: %v", err)
	}

	if res.Title != game.Title {
		t.Errorf("Expected title %q, got %q", game.Title, res.Title)
	}
	if res.Publisher != game.Publisher {
		t.Errorf("Expected publisher %q, got %q", game.Publisher, res.Publisher)
	}
	if res.ReleaseDate != game.ReleaseDate {
		t.Errorf("Expected release date %v, got %v", game.ReleaseDate, res.ReleaseDate)
	}
}

func TestAddBlock(t *testing.T) {
	// Mock database
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	dbGames, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer dbGames.Close()

	// Server instance
	server := &Server{db: db, dbGames: dbGames}

	server.NewBlockchain()

	req := &blockchainpb.AddBlockRequest{
		GameId:       "test-id",
		User:         "Test User",
		CheckoutDate: "01-01-2023",
	}

	game := &blockchainpb.VideoGame{
		Id:          "test-id",
		Title:       "Test Game",
		Publisher:   "Test Publisher",
		ReleaseDate: "01-01-2023",
	}

	err = server.dbGames.Update(func(txn *badger.Txn) error {
		gameData, err := proto.Marshal(game)
		if err != nil {
			return err
		}
		return txn.Set([]byte(game.Id), gameData)
	})

	if err != nil {
		status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to create game: %v", err),
		)
	}

	got, err := server.AddBlock(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create block: %v", err)
	}

	if got.Hash == "" {
		t.Error("Expected non-empty Hash")
	}
}

func TestGetBlock(t *testing.T) {

	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	server := &Server{db: db}

	server.NewBlockchain()

	req := &blockchainpb.GetBlockRequest{
		Position: 0,
	}

	got, err := server.GetBlock(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to get block: %v", err)
	}

	if got.Position != req.Position {
		t.Errorf("Expected position %q, got %q", req.Position, got.Position)
	}
	if got.PrevBlockHash != "" {
		t.Errorf("Expected previous block hash %q", "")
	}
	if got.Hash == "" {
		t.Errorf("Expected not empty hash")
	}
	if got.Data.GameId != "" {
		t.Errorf("Expected game ID %q", "")
	}
	if got.Data.User != "" {
		t.Errorf("Expected user %q", "")
	}
	if got.Data.CheckoutDate != "" {
		t.Errorf("Expected checkout date %q", "")
	}
	if got.Data.IsGenesis != true {
		t.Errorf("Expected genesis block")
	}
}
