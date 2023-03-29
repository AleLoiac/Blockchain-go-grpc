// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0--rc3
// source: blockchain/blockchainpb/blockchain.proto

package blockchainpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VideoGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Publisher   string `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	ReleaseDate string `protobuf:"bytes,4,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
}

func (x *VideoGame) Reset() {
	*x = VideoGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoGame) ProtoMessage() {}

func (x *VideoGame) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoGame.ProtoReflect.Descriptor instead.
func (*VideoGame) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *VideoGame) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VideoGame) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoGame) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *VideoGame) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

type AddVideoGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Publisher   string `protobuf:"bytes,2,opt,name=publisher,proto3" json:"publisher,omitempty"`
	ReleaseDate string `protobuf:"bytes,3,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
}

func (x *AddVideoGameRequest) Reset() {
	*x = AddVideoGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVideoGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVideoGameRequest) ProtoMessage() {}

func (x *AddVideoGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVideoGameRequest.ProtoReflect.Descriptor instead.
func (*AddVideoGameRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *AddVideoGameRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddVideoGameRequest) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *AddVideoGameRequest) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

type GameCheckout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId       string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	User         string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	CheckoutDate string `protobuf:"bytes,3,opt,name=checkout_date,json=checkoutDate,proto3" json:"checkout_date,omitempty"`
	IsGenesis    bool   `protobuf:"varint,4,opt,name=is_genesis,json=isGenesis,proto3" json:"is_genesis,omitempty"`
}

func (x *GameCheckout) Reset() {
	*x = GameCheckout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameCheckout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameCheckout) ProtoMessage() {}

func (x *GameCheckout) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameCheckout.ProtoReflect.Descriptor instead.
func (*GameCheckout) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *GameCheckout) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *GameCheckout) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GameCheckout) GetCheckoutDate() string {
	if x != nil {
		return x.CheckoutDate
	}
	return ""
}

func (x *GameCheckout) GetIsGenesis() bool {
	if x != nil {
		return x.IsGenesis
	}
	return false
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position      int32         `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Data          *GameCheckout `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	PrevBlockHash string        `protobuf:"bytes,3,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	Hash          string        `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{3}
}

func (x *Block) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Block) GetData() *GameCheckout {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Block) GetPrevBlockHash() string {
	if x != nil {
		return x.PrevBlockHash
	}
	return ""
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Blockchain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *Blockchain) Reset() {
	*x = Blockchain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blockchain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blockchain) ProtoMessage() {}

func (x *Blockchain) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blockchain.ProtoReflect.Descriptor instead.
func (*Blockchain) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *Blockchain) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type AddBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GameCheckout `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddBlockRequest) Reset() {
	*x = AddBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlockRequest) ProtoMessage() {}

func (x *AddBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlockRequest.ProtoReflect.Descriptor instead.
func (*AddBlockRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *AddBlockRequest) GetData() *GameCheckout {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *AddBlockResponse) Reset() {
	*x = AddBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlockResponse) ProtoMessage() {}

func (x *AddBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlockResponse.ProtoReflect.Descriptor instead.
func (*AddBlockResponse) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{6}
}

func (x *AddBlockResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int32 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlockRequest) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type GetBlockChainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBlockChainRequest) Reset() {
	*x = GetBlockChainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockChainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockChainRequest) ProtoMessage() {}

func (x *GetBlockChainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_blockchainpb_blockchain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockChainRequest.ProtoReflect.Descriptor instead.
func (*GetBlockChainRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP(), []int{8}
}

var File_blockchain_blockchainpb_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_blockchainpb_blockchain_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x72, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x6c, 0x0a, 0x13, 0x41, 0x64,
	0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x7f, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a,
	0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x37, 0x0a, 0x0a, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x22, 0x3f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2d, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x32, 0xae, 0x02, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x20, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x0c, 0x41, 0x64, 0x64,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x61, 0x6d,
	0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_blockchainpb_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_blockchainpb_blockchain_proto_rawDescData = file_blockchain_blockchainpb_blockchain_proto_rawDesc
)

func file_blockchain_blockchainpb_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_blockchainpb_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_blockchainpb_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_blockchainpb_blockchain_proto_rawDescData)
	})
	return file_blockchain_blockchainpb_blockchain_proto_rawDescData
}

var file_blockchain_blockchainpb_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_blockchain_blockchainpb_blockchain_proto_goTypes = []interface{}{
	(*VideoGame)(nil),            // 0: blockchain.VideoGame
	(*AddVideoGameRequest)(nil),  // 1: blockchain.AddVideoGameRequest
	(*GameCheckout)(nil),         // 2: blockchain.gameCheckout
	(*Block)(nil),                // 3: blockchain.Block
	(*Blockchain)(nil),           // 4: blockchain.Blockchain
	(*AddBlockRequest)(nil),      // 5: blockchain.AddBlockRequest
	(*AddBlockResponse)(nil),     // 6: blockchain.AddBlockResponse
	(*GetBlockRequest)(nil),      // 7: blockchain.GetBlockRequest
	(*GetBlockChainRequest)(nil), // 8: blockchain.GetBlockChainRequest
}
var file_blockchain_blockchainpb_blockchain_proto_depIdxs = []int32{
	2, // 0: blockchain.Block.data:type_name -> blockchain.gameCheckout
	3, // 1: blockchain.Blockchain.blocks:type_name -> blockchain.Block
	2, // 2: blockchain.AddBlockRequest.data:type_name -> blockchain.gameCheckout
	5, // 3: blockchain.BlockchainService.AddBlock:input_type -> blockchain.AddBlockRequest
	7, // 4: blockchain.BlockchainService.GetBlock:input_type -> blockchain.GetBlockRequest
	8, // 5: blockchain.BlockchainService.GetBlockchain:input_type -> blockchain.GetBlockChainRequest
	1, // 6: blockchain.BlockchainService.AddVideoGame:input_type -> blockchain.AddVideoGameRequest
	6, // 7: blockchain.BlockchainService.AddBlock:output_type -> blockchain.AddBlockResponse
	3, // 8: blockchain.BlockchainService.GetBlock:output_type -> blockchain.Block
	3, // 9: blockchain.BlockchainService.GetBlockchain:output_type -> blockchain.Block
	0, // 10: blockchain.BlockchainService.AddVideoGame:output_type -> blockchain.VideoGame
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blockchain_blockchainpb_blockchain_proto_init() }
func file_blockchain_blockchainpb_blockchain_proto_init() {
	if File_blockchain_blockchainpb_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoGame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddVideoGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameCheckout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blockchain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchain_blockchainpb_blockchain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockChainRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_blockchainpb_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockchain_blockchainpb_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_blockchainpb_blockchain_proto_depIdxs,
		MessageInfos:      file_blockchain_blockchainpb_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_blockchainpb_blockchain_proto = out.File
	file_blockchain_blockchainpb_blockchain_proto_rawDesc = nil
	file_blockchain_blockchainpb_blockchain_proto_goTypes = nil
	file_blockchain_blockchainpb_blockchain_proto_depIdxs = nil
}
