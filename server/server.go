package main

import (
	"Blockchain/blockchain/blockchainpb"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
		for it.Rewind(); it.Valid(); it.Next() {
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
	hash := sha256.Sum256([]byte(lastBlock.Hash + req.Data))
	strHash := hex.EncodeToString(hash[:])

	block := &blockchainpb.Block{
		Position:      lastBlock.Position + 1,
		Data:          req.Data,
		PrevBlockHash: lastBlock.Hash,
		Hash:          strHash,
	}

	err := s.db.Update(func(txn *badger.Txn) error {
		productData, err := proto.Marshal(block)
		if err != nil {
			return err
		}
		return txn.Set([]byte(strconv.Itoa(int(block.Position))), productData)
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

func (s *Server) NewBlock(data string) *blockchainpb.Block {
	hash := sha256.Sum256([]byte(data))
	strHash := hex.EncodeToString(hash[:])

	block := &blockchainpb.Block{
		Position:      0,
		Data:          data,
		PrevBlockHash: "",
		Hash:          strHash,
	}

	err := s.db.Update(func(txn *badger.Txn) error {
		productData, err := proto.Marshal(block)
		if err != nil {
			return err
		}
		return txn.Set([]byte(strconv.Itoa(int(block.Position))), productData)
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
	return s.NewBlock("Genesis Block")
}

func (s *Server) checkGenesis() bool {

	req := &blockchainpb.GetBlockRequest{Position: 0}

	block, _ := s.GetBlock(context.Background(), req)
	if block.GetData() == "Genesis Block" {
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

type Server struct {
	blockchainpb.BlockchainServiceServer
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

	if server.checkGenesis() == false {
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
