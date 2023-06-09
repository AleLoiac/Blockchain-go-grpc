// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc3
// source: blockchain.proto

package blockchainpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlockchainServiceClient is the client API for BlockchainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainServiceClient interface {
	AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*Block, error)
	GetBlockchain(ctx context.Context, in *GetBlockChainRequest, opts ...grpc.CallOption) (BlockchainService_GetBlockchainClient, error)
	AddVideoGame(ctx context.Context, in *AddVideoGameRequest, opts ...grpc.CallOption) (*VideoGame, error)
	GetVideoGame(ctx context.Context, in *GetVideoGameRequest, opts ...grpc.CallOption) (*VideoGame, error)
	ListVideoGames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BlockchainService_ListVideoGamesClient, error)
}

type blockchainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainServiceClient(cc grpc.ClientConnInterface) BlockchainServiceClient {
	return &blockchainServiceClient{cc}
}

func (c *blockchainServiceClient) AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error) {
	out := new(AddBlockResponse)
	err := c.cc.Invoke(ctx, "/blockchain.BlockchainService/AddBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/blockchain.BlockchainService/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetBlockchain(ctx context.Context, in *GetBlockChainRequest, opts ...grpc.CallOption) (BlockchainService_GetBlockchainClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[0], "/blockchain.BlockchainService/GetBlockchain", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceGetBlockchainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockchainService_GetBlockchainClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type blockchainServiceGetBlockchainClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceGetBlockchainClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) AddVideoGame(ctx context.Context, in *AddVideoGameRequest, opts ...grpc.CallOption) (*VideoGame, error) {
	out := new(VideoGame)
	err := c.cc.Invoke(ctx, "/blockchain.BlockchainService/AddVideoGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetVideoGame(ctx context.Context, in *GetVideoGameRequest, opts ...grpc.CallOption) (*VideoGame, error) {
	out := new(VideoGame)
	err := c.cc.Invoke(ctx, "/blockchain.BlockchainService/GetVideoGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) ListVideoGames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BlockchainService_ListVideoGamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[1], "/blockchain.BlockchainService/ListVideoGames", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceListVideoGamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockchainService_ListVideoGamesClient interface {
	Recv() (*VideoGame, error)
	grpc.ClientStream
}

type blockchainServiceListVideoGamesClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceListVideoGamesClient) Recv() (*VideoGame, error) {
	m := new(VideoGame)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockchainServiceServer is the server API for BlockchainService service.
// All implementations must embed UnimplementedBlockchainServiceServer
// for forward compatibility
type BlockchainServiceServer interface {
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*Block, error)
	GetBlockchain(*GetBlockChainRequest, BlockchainService_GetBlockchainServer) error
	AddVideoGame(context.Context, *AddVideoGameRequest) (*VideoGame, error)
	GetVideoGame(context.Context, *GetVideoGameRequest) (*VideoGame, error)
	ListVideoGames(*Empty, BlockchainService_ListVideoGamesServer) error
	mustEmbedUnimplementedBlockchainServiceServer()
}

// UnimplementedBlockchainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainServiceServer struct {
}

func (UnimplementedBlockchainServiceServer) AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
}
func (UnimplementedBlockchainServiceServer) GetBlock(context.Context, *GetBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBlockchainServiceServer) GetBlockchain(*GetBlockChainRequest, BlockchainService_GetBlockchainServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockchain not implemented")
}
func (UnimplementedBlockchainServiceServer) AddVideoGame(context.Context, *AddVideoGameRequest) (*VideoGame, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideoGame not implemented")
}
func (UnimplementedBlockchainServiceServer) GetVideoGame(context.Context, *GetVideoGameRequest) (*VideoGame, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoGame not implemented")
}
func (UnimplementedBlockchainServiceServer) ListVideoGames(*Empty, BlockchainService_ListVideoGamesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListVideoGames not implemented")
}
func (UnimplementedBlockchainServiceServer) mustEmbedUnimplementedBlockchainServiceServer() {}

// UnsafeBlockchainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServiceServer will
// result in compilation errors.
type UnsafeBlockchainServiceServer interface {
	mustEmbedUnimplementedBlockchainServiceServer()
}

func RegisterBlockchainServiceServer(s grpc.ServiceRegistrar, srv BlockchainServiceServer) {
	s.RegisterService(&BlockchainService_ServiceDesc, srv)
}

func _BlockchainService_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.BlockchainService/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).AddBlock(ctx, req.(*AddBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.BlockchainService/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetBlockchain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlockChainRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockchainServiceServer).GetBlockchain(m, &blockchainServiceGetBlockchainServer{stream})
}

type BlockchainService_GetBlockchainServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type blockchainServiceGetBlockchainServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceGetBlockchainServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockchainService_AddVideoGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVideoGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).AddVideoGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.BlockchainService/AddVideoGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).AddVideoGame(ctx, req.(*AddVideoGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetVideoGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetVideoGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.BlockchainService/GetVideoGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetVideoGame(ctx, req.(*GetVideoGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_ListVideoGames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockchainServiceServer).ListVideoGames(m, &blockchainServiceListVideoGamesServer{stream})
}

type BlockchainService_ListVideoGamesServer interface {
	Send(*VideoGame) error
	grpc.ServerStream
}

type blockchainServiceListVideoGamesServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceListVideoGamesServer) Send(m *VideoGame) error {
	return x.ServerStream.SendMsg(m)
}

// BlockchainService_ServiceDesc is the grpc.ServiceDesc for BlockchainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockchainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blockchain.BlockchainService",
	HandlerType: (*BlockchainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlock",
			Handler:    _BlockchainService_AddBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BlockchainService_GetBlock_Handler,
		},
		{
			MethodName: "AddVideoGame",
			Handler:    _BlockchainService_AddVideoGame_Handler,
		},
		{
			MethodName: "GetVideoGame",
			Handler:    _BlockchainService_GetVideoGame_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlockchain",
			Handler:       _BlockchainService_GetBlockchain_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListVideoGames",
			Handler:       _BlockchainService_ListVideoGames_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blockchain.proto",
}
