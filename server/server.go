package main

import (
	"Blockchain/blockchain/blockchainpb"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"strconv"
)

func (s *Server) GetLastBlock() *blockchainpb.Block {

	var lastBlock *blockchainpb.Block

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 1
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		// Iterate in reverse order to get the last block
		for it.Rewind(); it.Valid(); {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				block := &blockchainpb.Block{}
				err := proto.Unmarshal(v, block)
				if err != nil {
					return err
				}
				lastBlock = block
				return nil
			})
			if err != nil {
				return err
			}
			// Stop iterating once the last block is found
			break
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to get last block: %v", err)
	}
	return lastBlock
}

func (s *Server) AddBlock(ctx context.Context, req *blockchainpb.AddBlockRequest) (*blockchainpb.AddBlockResponse, error) {

	fmt.Printf("AddBlock function is invoked with: %v\n", req)

	lastBlock := s.GetLastBlock()

	//create the hash and encode it to a string
	hash := sha256.Sum256([]byte(lastBlock.Hash + req.GetGameId() + req.GetUser() + req.GetCheckoutDate()))
	strHash := hex.EncodeToString(hash[:])

	data := &blockchainpb.GameCheckout{
		GameId:       req.GetGameId(),
		User:         req.GetUser(),
		CheckoutDate: req.GetCheckoutDate(),
		IsGenesis:    false,
	}

	err := s.dbGames.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(req.GetGameId()))
		if err != nil {
			return status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find game with id: %v", req.GetGameId()))
		}
		return err
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find game with id: %v", req.GetGameId()))
	}

	block := &blockchainpb.Block{
		Position:      lastBlock.Position + 1,
		Data:          data,
		Timestamp:     timestamppb.Now(),
		PrevBlockHash: lastBlock.Hash,
		Hash:          strHash,
	}

	err = s.db.Update(func(txn *badger.Txn) error {
		blockData, err := proto.Marshal(block)
		if err != nil {
			return err
		}
		return txn.Set([]byte(strconv.Itoa(int(block.Position))), blockData)
	})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to create block: %v", err),
		)
	}

	return &blockchainpb.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) NewBlock(checkout *blockchainpb.GameCheckout) *blockchainpb.Block {

	hash := sha256.Sum256([]byte(checkout.GetGameId() + checkout.GetUser() + checkout.GetCheckoutDate()))
	strHash := hex.EncodeToString(hash[:])

	block := &blockchainpb.Block{
		Position:      0,
		Data:          checkout,
		Timestamp:     timestamppb.Now(),
		PrevBlockHash: "",
		Hash:          strHash,
	}

	err := s.db.Update(func(txn *badger.Txn) error {
		blockData, err := proto.Marshal(block)
		if err != nil {
			return err
		}
		return txn.Set([]byte(strconv.Itoa(int(block.Position))), blockData)
	})

	if err != nil {
		status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to create first block: %v", err),
		)
		return nil
	}

	return block
}

func (s *Server) GetBlock(ctx context.Context, req *blockchainpb.GetBlockRequest) (*blockchainpb.Block, error) {

	fmt.Printf("GetBlock function is invoked with block at position: %v\n", req.GetPosition())

	var block blockchainpb.Block

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(strconv.Itoa(int(req.GetPosition()))))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			err = proto.Unmarshal(val, &block)
			if err != nil {
				fmt.Printf("Error unmarshaling block data: %v\n", err)
				return err
			}
			return nil
		})
		return err
	})

	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("Block with position '%s' not found", strconv.Itoa(int(req.GetPosition()))),
			)
		} else {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Failed to get block: %v", err),
			)
		}
	}

	return &blockchainpb.Block{
		Position:      block.GetPosition(),
		Data:          block.GetData(),
		Timestamp:     block.GetTimestamp(),
		PrevBlockHash: block.GetPrevBlockHash(),
		Hash:          block.GetHash(),
	}, nil
}

func (s *Server) GetBlockchain(req *blockchainpb.GetBlockChainRequest, stream blockchainpb.BlockchainService_GetBlockchainServer) error {

	fmt.Println("GetBlockchain function is invoked with an empty request")

	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = 10
	opts.PrefetchValues = false

	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			k, _ := strconv.Atoi(string(key))
			block, err := s.GetBlock(stream.Context(), &blockchainpb.GetBlockRequest{Position: int32(k)})
			if err != nil {
				return err
			}
			if err = stream.Send(block); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) NewGenesisBlock() *blockchainpb.Block {

	emptyCheckout := &blockchainpb.GameCheckout{
		GameId:       "",
		User:         "",
		CheckoutDate: "",
		IsGenesis:    true,
	}

	return s.NewBlock(emptyCheckout)
}

func (s *Server) CheckGenesis() bool {

	req := &blockchainpb.GetBlockRequest{Position: 0}

	block, _ := s.GetBlock(context.Background(), req)

	if block.GetData().GetIsGenesis() == true {
		return true
	} else {
		return false
	}
}

func (s *Server) NewBlockchain() *blockchainpb.Blockchain {

	return &blockchainpb.Blockchain{
		Blocks: []*blockchainpb.Block{
			s.NewGenesisBlock(),
		},
	}
}

func (s *Server) AddVideoGame(ctx context.Context, req *blockchainpb.AddVideoGameRequest) (*blockchainpb.VideoGame, error) {

	fmt.Printf("AddVideoGame function is invoked with: %v\n", req)

	game := &blockchainpb.VideoGame{
		Id:          uuid.New().String(),
		Title:       req.GetTitle(),
		Publisher:   req.GetPublisher(),
		ReleaseDate: req.GetReleaseDate(),
	}

	err := s.dbGames.Update(func(txn *badger.Txn) error {
		gameData, err := proto.Marshal(game)
		if err != nil {
			return err
		}
		return txn.Set([]byte(game.Id), gameData)
	})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to create game: %v", err),
		)
	}

	return &blockchainpb.VideoGame{
		Id:          game.GetId(),
		Title:       game.GetTitle(),
		Publisher:   game.GetPublisher(),
		ReleaseDate: game.GetReleaseDate(),
	}, nil
}

func (s *Server) GetVideoGame(ctx context.Context, req *blockchainpb.GetVideoGameRequest) (*blockchainpb.VideoGame, error) {

	fmt.Printf("GetVideoGame function is invoked with: %v\n", req)

	var game blockchainpb.VideoGame

	err := s.dbGames.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(req.GetGameId()))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			err = proto.Unmarshal(val, &game)
			if err != nil {
				fmt.Printf("Error unmarshaling game data: %v\n", err)
				return err
			}
			return nil
		})
		return err
	})

	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("Game with ID '%s' not found", req.GameId),
			)
		} else {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Failed to get game: %v", err),
			)
		}
	}

	return &blockchainpb.VideoGame{
		Id:          game.GetId(),
		Title:       game.GetTitle(),
		Publisher:   game.GetPublisher(),
		ReleaseDate: game.GetReleaseDate(),
	}, nil
}

func (s *Server) ListVideoGames(req *blockchainpb.Empty, stream blockchainpb.BlockchainService_ListVideoGamesServer) error {

	fmt.Println("ListVideoGames function is invoked with an empty request")

	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = 10
	opts.PrefetchValues = false

	err := s.dbGames.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			game, err := s.GetVideoGame(stream.Context(), &blockchainpb.GetVideoGameRequest{GameId: string(key)})
			if err != nil {
				return err
			}
			if err = stream.Send(game); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

type Server struct {
	blockchainpb.BlockchainServiceServer
	db      *badger.DB
	dbGames *badger.DB
}

func NewServer(db *badger.DB, dbGames *badger.DB) *Server {
	return &Server{db: db, dbGames: dbGames}
}

func main() {
	fmt.Println("Server started...")

	db, _ := badger.Open(badger.DefaultOptions("/Users/aless/Desktop/Go/Blockchain/DB"))
	defer db.Close()

	dbGames, _ := badger.Open(badger.DefaultOptions("/Users/aless/Desktop/Go/Blockchain/games_DB"))
	defer dbGames.Close()

	server := NewServer(db, dbGames)

	if server.CheckGenesis() == false {
		server.NewBlockchain()
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blockchainpb.RegisterBlockchainServiceServer(s, server)

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
