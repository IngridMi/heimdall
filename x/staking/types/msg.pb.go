// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staking/v1beta/msg.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	github_com_maticnetwork_heimdall_types "github.com/maticnetwork/heimdall/types"
	github_com_maticnetwork_heimdall_types_common "github.com/maticnetwork/heimdall/types/common"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgValidatorJoin struct {
	From            github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,1,opt,name=from,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"from,omitempty"`
	ID              github_com_maticnetwork_heimdall_types.ValidatorID            `protobuf:"varint,2,opt,name=id,proto3,casttype=github.com/maticnetwork/heimdall/types.ValidatorID" json:"id,omitempty"`
	ActivationEpoch uint64                                                        `protobuf:"varint,3,opt,name=activation_epoch,json=activationEpoch,proto3" json:"activation_epoch,omitempty"`
	Amount          github_com_cosmos_cosmos_sdk_types.Int                        `protobuf:"bytes,4,opt,name=amount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount,omitempty"`
	SignerPubKey    github_com_maticnetwork_heimdall_types_common.PubKey          `protobuf:"bytes,5,opt,name=signer_pub_key,json=signerPubKey,proto3,casttype=github.com/maticnetwork/heimdall/types/common.PubKey" json:"signer_pub_key,omitempty"`
	TxHash          github_com_maticnetwork_heimdall_types_common.HeimdallHash    `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallHash" json:"tx_hash,omitempty"`
	LogIndex        uint64                                                        `protobuf:"varint,7,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	BlockNumber     uint64                                                        `protobuf:"varint,8,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Nonce           uint64                                                        `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MsgValidatorJoin) Reset()         { *m = MsgValidatorJoin{} }
func (m *MsgValidatorJoin) String() string { return proto.CompactTextString(m) }
func (*MsgValidatorJoin) ProtoMessage()    {}
func (*MsgValidatorJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b2771020bbd7c4, []int{0}
}
func (m *MsgValidatorJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgValidatorJoin.Unmarshal(m, b)
}
func (m *MsgValidatorJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgValidatorJoin.Marshal(b, m, deterministic)
}
func (m *MsgValidatorJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgValidatorJoin.Merge(m, src)
}
func (m *MsgValidatorJoin) XXX_Size() int {
	return xxx_messageInfo_MsgValidatorJoin.Size(m)
}
func (m *MsgValidatorJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgValidatorJoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgValidatorJoin proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgValidatorJoin)(nil), "heimdall.staking.x.staking.types.MsgValidatorJoin")
}

func init() { proto.RegisterFile("staking/v1beta/msg.proto", fileDescriptor_c6b2771020bbd7c4) }

var fileDescriptor_c6b2771020bbd7c4 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x75, 0x42, 0x9a, 0x36, 0x43, 0x05, 0xd5, 0xa8, 0x8b, 0xa1, 0x48, 0x76, 0x60, 0x81, 0x0a,
	0x12, 0x1e, 0x5e, 0x0b, 0x54, 0x01, 0x52, 0xa3, 0x22, 0x35, 0xe5, 0x21, 0x64, 0x09, 0x90, 0x58,
	0x60, 0x8d, 0xed, 0xe9, 0x78, 0x14, 0xcf, 0xdc, 0xc8, 0x33, 0x2e, 0xc9, 0x1f, 0xb0, 0xe4, 0x13,
	0xf8, 0x0f, 0x7e, 0x80, 0x65, 0x97, 0xac, 0x2c, 0x94, 0xfc, 0x45, 0x56, 0x28, 0x13, 0xb7, 0x20,
	0x36, 0x3c, 0x56, 0x33, 0xe7, 0x9c, 0xb9, 0xe7, 0x9e, 0x7b, 0x65, 0x23, 0x62, 0x2c, 0x1b, 0x49,
	0x2d, 0xe8, 0xc9, 0xdd, 0x84, 0x5b, 0x46, 0x95, 0x11, 0xe1, 0xb8, 0x04, 0x0b, 0xb8, 0x9f, 0x73,
	0xa9, 0x32, 0x56, 0x14, 0x61, 0xf3, 0x24, 0x9c, 0x9c, 0xdf, 0xec, 0x74, 0xcc, 0xcd, 0xce, 0xb6,
	0x00, 0x01, 0xee, 0x31, 0x5d, 0xde, 0x56, 0x75, 0x3b, 0x57, 0x04, 0x80, 0x28, 0x38, 0x75, 0x28,
	0xa9, 0x8e, 0x29, 0xd3, 0xd3, 0x46, 0xf2, 0x7f, 0x97, 0xb2, 0xaa, 0x64, 0x56, 0x82, 0x5e, 0xe9,
	0xd7, 0xbf, 0x74, 0xd0, 0xd6, 0x0b, 0x23, 0xde, 0xb0, 0x42, 0x66, 0xcc, 0x42, 0x79, 0x04, 0x52,
	0xe3, 0xd7, 0xa8, 0x73, 0x5c, 0x82, 0x22, 0xad, 0x7e, 0x6b, 0xb7, 0x37, 0xd8, 0x5f, 0xd4, 0xc1,
	0x63, 0x21, 0x6d, 0x5e, 0x25, 0x61, 0x0a, 0x8a, 0x2a, 0x66, 0x65, 0xaa, 0xb9, 0xfd, 0x00, 0xe5,
	0x88, 0x9e, 0x25, 0xa6, 0x2e, 0x1f, 0x4d, 0x41, 0x29, 0xd0, 0xe1, 0x61, 0xc3, 0xee, 0x67, 0x59,
	0xc9, 0x8d, 0x89, 0x9c, 0x1d, 0x7e, 0x8e, 0xda, 0x32, 0x23, 0xed, 0x7e, 0x6b, 0xb7, 0x33, 0x78,
	0x34, 0xab, 0x83, 0xf6, 0xf0, 0x60, 0x51, 0x07, 0xf7, 0xfe, 0xce, 0x3a, 0x3c, 0x8f, 0x38, 0x3c,
	0x88, 0xda, 0x32, 0xc3, 0x37, 0xd1, 0x16, 0x4b, 0xad, 0x3c, 0x71, 0xd3, 0xc4, 0x7c, 0x0c, 0x69,
	0x4e, 0x2e, 0x2c, 0xbd, 0xa3, 0xcb, 0x3f, 0xf9, 0xa7, 0x4b, 0x1a, 0x0f, 0x50, 0x97, 0x29, 0xa8,
	0xb4, 0x25, 0x1d, 0x37, 0xd1, 0xad, 0x45, 0x1d, 0xdc, 0xf8, 0xa5, 0x6d, 0x0a, 0x46, 0x81, 0x69,
	0x8e, 0xdb, 0x26, 0x1b, 0x35, 0x2d, 0x87, 0xda, 0x46, 0x4d, 0x25, 0x7e, 0x8f, 0x2e, 0x19, 0x29,
	0x34, 0x2f, 0xe3, 0x71, 0x95, 0xc4, 0x23, 0x3e, 0x25, 0x6b, 0xce, 0xeb, 0xe1, 0xa2, 0x0e, 0x1e,
	0xfc, 0xdb, 0x76, 0x5e, 0x55, 0xc9, 0x33, 0x3e, 0x8d, 0x36, 0x57, 0x7e, 0x2b, 0x84, 0xdf, 0xa2,
	0x75, 0x3b, 0x89, 0x73, 0x66, 0x72, 0xd2, 0x75, 0xc6, 0x4f, 0x16, 0x75, 0xb0, 0xf7, 0x7f, 0x6b,
	0x3f, 0x64, 0x26, 0x8f, 0xba, 0x76, 0xb2, 0x3c, 0xf1, 0x55, 0xd4, 0x2b, 0x40, 0xc4, 0x52, 0x67,
	0x7c, 0x42, 0xd6, 0xdd, 0x82, 0x36, 0x0a, 0x10, 0xc3, 0x25, 0xc6, 0xd7, 0xd0, 0x66, 0x52, 0x40,
	0x3a, 0x8a, 0x75, 0xa5, 0x12, 0x5e, 0x92, 0x0d, 0xa7, 0x5f, 0x74, 0xdc, 0x4b, 0x47, 0xe1, 0x6d,
	0xb4, 0xa6, 0x41, 0xa7, 0x9c, 0xf4, 0x9c, 0xb6, 0x02, 0x7b, 0x9d, 0x8f, 0x9f, 0x03, 0x6f, 0x70,
	0xf4, 0x75, 0xe6, 0x7b, 0xa7, 0x33, 0xdf, 0xfb, 0x3e, 0xf3, 0xbd, 0x4f, 0x73, 0xdf, 0x3b, 0x9d,
	0xfb, 0xde, 0xb7, 0xb9, 0xef, 0xbd, 0xbb, 0xf3, 0xc7, 0xe4, 0x13, 0x7a, 0xf6, 0x1f, 0xb8, 0x19,
	0x92, 0xae, 0xfb, 0x20, 0xef, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xf0, 0x50, 0xa9, 0x1f,
	0x03, 0x00, 0x00,
}
