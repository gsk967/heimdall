// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkpoint/v1beta/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	github_com_maticnetwork_heimdall_types "github.com/maticnetwork/heimdall/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	CheckpointBufferTime time.Duration `protobuf:"bytes,1,opt,name=checkpoint_buffer_time,json=checkpointBufferTime,proto3,stdduration" json:"checkpoint_buffer_time"`
	AvgCheckpointLength  uint64        `protobuf:"varint,2,opt,name=avg_checkpoint_length,json=avgCheckpointLength,proto3" json:"avg_checkpoint_length,omitempty"`
	MaxCheckpointLength  uint64        `protobuf:"varint,3,opt,name=max_checkpoint_length,json=maxCheckpointLength,proto3" json:"max_checkpoint_length,omitempty"`
	ChildBlockInterval   uint64        `protobuf:"varint,4,opt,name=child_block_interval,json=childBlockInterval,proto3" json:"child_block_interval,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a783d62b74a2d0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// GenesisState defines the checkpoint module's genesis state.
type GenesisState struct {
	Params             Params                                               `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BufferedCheckpoint *github_com_maticnetwork_heimdall_types.Checkpoint   `protobuf:"bytes,2,opt,name=buffered_checkpoint,json=bufferedCheckpoint,proto3,casttype=github.com/maticnetwork/heimdall/types.Checkpoint" json:"buffered_checkpoint,omitempty"`
	LastNoAck          uint64                                               `protobuf:"varint,3,opt,name=last_no_ack,json=lastNoAck,proto3" json:"last_no_ack,omitempty"`
	AckCount           uint64                                               `protobuf:"varint,4,opt,name=ack_count,json=ackCount,proto3" json:"ack_count,omitempty"`
	Checkpoints        []*github_com_maticnetwork_heimdall_types.Checkpoint `protobuf:"bytes,5,rep,name=checkpoints,proto3,casttype=github.com/maticnetwork/heimdall/types.Checkpoint" json:"checkpoints,omitempty"`
}

func (m *GenesisState) Reset()      { *m = GenesisState{} }
func (*GenesisState) ProtoMessage() {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a783d62b74a2d0, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisState.Unmarshal(m, b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return xxx_messageInfo_GenesisState.Size(m)
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "heimdall.checkpoint.x.checkpoint.types.Params")
	proto.RegisterType((*GenesisState)(nil), "heimdall.checkpoint.x.checkpoint.types.GenesisState")
}

func init() { proto.RegisterFile("checkpoint/v1beta/genesis.proto", fileDescriptor_99a783d62b74a2d0) }

var fileDescriptor_99a783d62b74a2d0 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x9d, 0x26, 0x44, 0xe9, 0x85, 0xe9, 0x1a, 0x50, 0x5a, 0x24, 0x3b, 0xea, 0x80, 0x3a,
	0xdd, 0xd1, 0x54, 0x2c, 0x6c, 0x75, 0x91, 0x10, 0x52, 0x41, 0x28, 0xb0, 0xc0, 0x62, 0x9d, 0x2f,
	0x97, 0xf3, 0xc9, 0xbe, 0xbb, 0xc8, 0x3e, 0x87, 0x44, 0x62, 0x87, 0xb1, 0x63, 0xc7, 0xfe, 0x39,
	0x1d, 0x3b, 0x32, 0x15, 0x94, 0xfc, 0x17, 0x4c, 0xc8, 0x67, 0x07, 0x5b, 0x14, 0x09, 0x06, 0x36,
	0xbf, 0xfb, 0xfc, 0x7d, 0xef, 0xdd, 0xef, 0xe9, 0x80, 0x47, 0x23, 0x46, 0xe3, 0xb9, 0x16, 0xca,
	0xe0, 0xc5, 0x71, 0xc8, 0x0c, 0xc1, 0x9c, 0x29, 0x96, 0x89, 0x0c, 0xcd, 0x53, 0x6d, 0x34, 0x7c,
	0x1c, 0x31, 0x21, 0xa7, 0x24, 0x49, 0x50, 0xfd, 0x27, 0x5a, 0x36, 0x0b, 0xb3, 0x9a, 0xb3, 0xec,
	0x60, 0xc0, 0x35, 0xd7, 0xd6, 0x82, 0x8b, 0xaf, 0xd2, 0x7d, 0xb0, 0xcf, 0xb5, 0xe6, 0x09, 0xc3,
	0xb6, 0x0a, 0xf3, 0x19, 0x26, 0x6a, 0x55, 0x49, 0xee, 0xef, 0xd2, 0x34, 0x4f, 0x89, 0x11, 0x5a,
	0x95, 0xfa, 0xe1, 0xe7, 0x1d, 0xd0, 0x7d, 0x43, 0x52, 0x22, 0x33, 0xf8, 0x1e, 0x3c, 0xac, 0xfb,
	0x05, 0x61, 0x3e, 0x9b, 0xb1, 0x34, 0x30, 0x42, 0xb2, 0x61, 0x6b, 0xd4, 0x3a, 0xea, 0x8f, 0xf7,
	0x51, 0x99, 0x85, 0xb6, 0x59, 0xe8, 0x79, 0x95, 0xe5, 0xf7, 0xae, 0x6f, 0x3d, 0xe7, 0xf2, 0x9b,
	0xd7, 0x9a, 0x0c, 0xea, 0x08, 0xdf, 0x26, 0xbc, 0x13, 0x92, 0xc1, 0x31, 0x78, 0x40, 0x16, 0x3c,
	0x68, 0xc4, 0x27, 0x4c, 0x71, 0x13, 0x0d, 0x77, 0x46, 0xad, 0xa3, 0xce, 0x64, 0x8f, 0x2c, 0xf8,
	0xd9, 0x2f, 0xed, 0xdc, 0x4a, 0x85, 0x47, 0x92, 0xe5, 0x1f, 0x3c, 0xed, 0xd2, 0x23, 0xc9, 0xf2,
	0x8e, 0xe7, 0x09, 0x18, 0xd0, 0x48, 0x24, 0xd3, 0x20, 0x4c, 0x34, 0x8d, 0x03, 0xa1, 0x0c, 0x4b,
	0x17, 0x24, 0x19, 0x76, 0xac, 0x05, 0x5a, 0xcd, 0x2f, 0xa4, 0x97, 0x95, 0xf2, 0xac, 0xf7, 0xe5,
	0xca, 0x73, 0x2e, 0xaf, 0x3c, 0xe7, 0xf0, 0xa2, 0x0d, 0xee, 0xbf, 0x28, 0x97, 0xf2, 0xd6, 0x10,
	0xc3, 0xe0, 0x39, 0xe8, 0xce, 0x2d, 0x99, 0xea, 0xfe, 0x08, 0xfd, 0xdb, 0x92, 0x50, 0xc9, 0xd3,
	0xef, 0x14, 0x50, 0x26, 0x55, 0x06, 0xfc, 0x04, 0xf6, 0x4a, 0xa4, 0x6c, 0xda, 0xb8, 0x93, 0x05,
	0xd0, 0x1f, 0x0f, 0xee, 0xa0, 0x3d, 0x55, 0x2b, 0xff, 0xe9, 0x8f, 0x5b, 0xef, 0x98, 0x0b, 0x13,
	0xe5, 0x21, 0xa2, 0x5a, 0x62, 0x49, 0x8c, 0xa0, 0x8a, 0x99, 0x8f, 0x3a, 0x8d, 0xf1, 0x76, 0x16,
	0x5c, 0xf6, 0xac, 0x51, 0x4c, 0xe0, 0xb6, 0x4f, 0x7d, 0x06, 0x5d, 0xd0, 0x4f, 0x48, 0x66, 0x02,
	0xa5, 0x03, 0x42, 0xe3, 0x0a, 0xe1, 0x6e, 0x71, 0xf4, 0x5a, 0x9f, 0xd2, 0x18, 0x3e, 0x02, 0xbb,
	0x84, 0xc6, 0x01, 0xd5, 0xb9, 0x32, 0x15, 0xad, 0x1e, 0xa1, 0xf1, 0x59, 0x51, 0x43, 0x09, 0xfa,
	0xf5, 0xc4, 0xd9, 0xf0, 0xde, 0xa8, 0xfd, 0xbf, 0x47, 0x6e, 0xe6, 0xd7, 0x2b, 0xf1, 0x5f, 0x5d,
	0xaf, 0x5d, 0xe7, 0x66, 0xed, 0x3a, 0xdf, 0xd7, 0xae, 0x73, 0xb1, 0x71, 0x9d, 0x9b, 0x8d, 0xeb,
	0x7c, 0xdd, 0xb8, 0xce, 0x87, 0x93, 0xbf, 0xf6, 0x58, 0xe2, 0xc6, 0x9b, 0xb3, 0x0d, 0xc3, 0xae,
	0x1d, 0xf5, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xfe, 0x46, 0xb0, 0x8e, 0x03, 0x00,
	0x00,
}