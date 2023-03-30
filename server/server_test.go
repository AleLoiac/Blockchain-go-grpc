package main

import (
	"Blockchain/blockchain/blockchainpb"
	"context"
	"github.com/dgraph-io/badger/v3"
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
