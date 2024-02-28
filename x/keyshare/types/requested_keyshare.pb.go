// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/keyshare/requested_keyshare.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type KeyShareRequest struct {
	Identity     string               `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Pubkey       string               `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	IbcInfo      *IBCInfo             `protobuf:"bytes,3,opt,name=ibc_info,json=ibcInfo,proto3" json:"ibc_info,omitempty"`
	Counterparty *CounterPartyIBCInfo `protobuf:"bytes,4,opt,name=counterparty,proto3" json:"counterparty,omitempty"`
	AggrKeyshare string               `protobuf:"bytes,5,opt,name=aggr_keyshare,json=aggrKeyshare,proto3" json:"aggr_keyshare,omitempty"`
	ProposalId   string               `protobuf:"bytes,6,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	RequestId    string               `protobuf:"bytes,7,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Sent         bool                 `protobuf:"varint,8,opt,name=sent,proto3" json:"sent,omitempty"`
}

func (m *KeyShareRequest) Reset()         { *m = KeyShareRequest{} }
func (m *KeyShareRequest) String() string { return proto.CompactTextString(m) }
func (*KeyShareRequest) ProtoMessage()    {}
func (*KeyShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ed024b19ae59bd, []int{0}
}
func (m *KeyShareRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyShareRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyShareRequest.Merge(m, src)
}
func (m *KeyShareRequest) XXX_Size() int {
	return m.Size()
}
func (m *KeyShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyShareRequest proto.InternalMessageInfo

func (m *KeyShareRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *KeyShareRequest) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *KeyShareRequest) GetIbcInfo() *IBCInfo {
	if m != nil {
		return m.IbcInfo
	}
	return nil
}

func (m *KeyShareRequest) GetCounterparty() *CounterPartyIBCInfo {
	if m != nil {
		return m.Counterparty
	}
	return nil
}

func (m *KeyShareRequest) GetAggrKeyshare() string {
	if m != nil {
		return m.AggrKeyshare
	}
	return ""
}

func (m *KeyShareRequest) GetProposalId() string {
	if m != nil {
		return m.ProposalId
	}
	return ""
}

func (m *KeyShareRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *KeyShareRequest) GetSent() bool {
	if m != nil {
		return m.Sent
	}
	return false
}

type IBCInfo struct {
	ClientID     string `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ConnectionID string `protobuf:"bytes,2,opt,name=ConnectionID,proto3" json:"ConnectionID,omitempty"`
	ChannelID    string `protobuf:"bytes,3,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	PortID       string `protobuf:"bytes,4,opt,name=PortID,proto3" json:"PortID,omitempty"`
}

func (m *IBCInfo) Reset()         { *m = IBCInfo{} }
func (m *IBCInfo) String() string { return proto.CompactTextString(m) }
func (*IBCInfo) ProtoMessage()    {}
func (*IBCInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ed024b19ae59bd, []int{1}
}
func (m *IBCInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCInfo.Merge(m, src)
}
func (m *IBCInfo) XXX_Size() int {
	return m.Size()
}
func (m *IBCInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IBCInfo proto.InternalMessageInfo

func (m *IBCInfo) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *IBCInfo) GetConnectionID() string {
	if m != nil {
		return m.ConnectionID
	}
	return ""
}

func (m *IBCInfo) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *IBCInfo) GetPortID() string {
	if m != nil {
		return m.PortID
	}
	return ""
}

type CounterPartyIBCInfo struct {
	ClientID     string `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ConnectionID string `protobuf:"bytes,2,opt,name=ConnectionID,proto3" json:"ConnectionID,omitempty"`
	ChannelID    string `protobuf:"bytes,3,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	PortID       string `protobuf:"bytes,4,opt,name=PortID,proto3" json:"PortID,omitempty"`
}

func (m *CounterPartyIBCInfo) Reset()         { *m = CounterPartyIBCInfo{} }
func (m *CounterPartyIBCInfo) String() string { return proto.CompactTextString(m) }
func (*CounterPartyIBCInfo) ProtoMessage()    {}
func (*CounterPartyIBCInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ed024b19ae59bd, []int{2}
}
func (m *CounterPartyIBCInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CounterPartyIBCInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CounterPartyIBCInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CounterPartyIBCInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CounterPartyIBCInfo.Merge(m, src)
}
func (m *CounterPartyIBCInfo) XXX_Size() int {
	return m.Size()
}
func (m *CounterPartyIBCInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CounterPartyIBCInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CounterPartyIBCInfo proto.InternalMessageInfo

func (m *CounterPartyIBCInfo) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *CounterPartyIBCInfo) GetConnectionID() string {
	if m != nil {
		return m.ConnectionID
	}
	return ""
}

func (m *CounterPartyIBCInfo) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *CounterPartyIBCInfo) GetPortID() string {
	if m != nil {
		return m.PortID
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyShareRequest)(nil), "fairyring.keyshare.KeyShareRequest")
	proto.RegisterType((*IBCInfo)(nil), "fairyring.keyshare.IBCInfo")
	proto.RegisterType((*CounterPartyIBCInfo)(nil), "fairyring.keyshare.CounterPartyIBCInfo")
}

func init() {
	proto.RegisterFile("fairyring/keyshare/requested_keyshare.proto", fileDescriptor_e8ed024b19ae59bd)
}

var fileDescriptor_e8ed024b19ae59bd = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xcd, 0xea, 0xd3, 0x40,
	0x14, 0xc5, 0x3b, 0x7f, 0x6b, 0x9b, 0xdc, 0x56, 0x84, 0x11, 0x24, 0xf8, 0x11, 0x4b, 0x5c, 0x58,
	0x10, 0x12, 0x50, 0xf0, 0x01, 0x9a, 0x22, 0x0c, 0xdd, 0x94, 0xb8, 0x73, 0x53, 0xf2, 0x31, 0x4d,
	0x87, 0xc6, 0x99, 0x38, 0x99, 0x80, 0x59, 0xf9, 0x00, 0x6e, 0x7c, 0x27, 0x37, 0x2e, 0xbb, 0x74,
	0x29, 0xed, 0x8b, 0xc8, 0x4c, 0x92, 0x96, 0x62, 0xf7, 0xff, 0xdd, 0x9c, 0x73, 0xef, 0xb9, 0xb9,
	0xf9, 0x71, 0xe1, 0xed, 0x36, 0x66, 0xb2, 0x91, 0x8c, 0xe7, 0xc1, 0x9e, 0x36, 0xd5, 0x2e, 0x96,
	0x34, 0x90, 0xf4, 0x6b, 0x4d, 0x2b, 0x45, 0xb3, 0x4d, 0x6f, 0xf9, 0xa5, 0x14, 0x4a, 0x60, 0x7c,
	0x6e, 0xf6, 0xfb, 0x8a, 0xf7, 0xeb, 0x0e, 0x1e, 0xaf, 0x68, 0xf3, 0x49, 0x8b, 0xa8, 0x0d, 0xe2,
	0x67, 0x60, 0xb1, 0x8c, 0x72, 0xc5, 0x54, 0xe3, 0xa0, 0x19, 0x9a, 0xdb, 0xd1, 0x59, 0xe3, 0xa7,
	0x30, 0x2a, 0xeb, 0x64, 0x4f, 0x1b, 0xe7, 0xce, 0x54, 0x3a, 0x85, 0x3f, 0x80, 0xc5, 0x92, 0x74,
	0xc3, 0xf8, 0x56, 0x38, 0x0f, 0x66, 0x68, 0x3e, 0x79, 0xf7, 0xdc, 0xff, 0xff, 0x73, 0x3e, 0x59,
	0x84, 0x84, 0x6f, 0x45, 0x34, 0x66, 0x49, 0xaa, 0x1f, 0x78, 0x05, 0xd3, 0x54, 0xd4, 0x5c, 0x51,
	0x59, 0xc6, 0x52, 0x35, 0xce, 0xd0, 0x64, 0xdf, 0xdc, 0xca, 0x86, 0x6d, 0xdf, 0x5a, 0xf7, 0xf5,
	0x73, 0xae, 0xc2, 0xf8, 0x35, 0x3c, 0x8a, 0xf3, 0x5c, 0x9e, 0xff, 0xdb, 0x79, 0x68, 0x76, 0x9c,
	0x6a, 0x73, 0xd5, 0x79, 0xf8, 0x15, 0x4c, 0x4a, 0x29, 0x4a, 0x51, 0xc5, 0xc5, 0x86, 0x65, 0xce,
	0xc8, 0xb4, 0x40, 0x6f, 0x91, 0x0c, 0xbf, 0x04, 0xe8, 0x10, 0xea, 0xfa, 0xd8, 0xd4, 0xed, 0xce,
	0x21, 0x19, 0xc6, 0x30, 0xac, 0x28, 0x57, 0x8e, 0x35, 0x43, 0x73, 0x2b, 0x32, 0x6f, 0xef, 0x3b,
	0x8c, 0xbb, 0x8d, 0x34, 0xbc, 0xb0, 0x60, 0x94, 0x2b, 0xb2, 0xec, 0xe1, 0xf5, 0x1a, 0x7b, 0x30,
	0x0d, 0x05, 0xe7, 0x34, 0x55, 0x4c, 0x70, 0xb2, 0xec, 0x10, 0x5e, 0x79, 0xf8, 0x05, 0xd8, 0xe1,
	0x2e, 0xe6, 0x9c, 0x16, 0x64, 0x69, 0x48, 0xda, 0xd1, 0xc5, 0xd0, 0xf8, 0xd7, 0x42, 0xea, 0xd9,
	0xc3, 0x16, 0x7f, 0xab, 0xbc, 0x1f, 0x08, 0x9e, 0xdc, 0xe0, 0x73, 0x3f, 0xdb, 0x2c, 0xc8, 0xef,
	0xa3, 0x8b, 0x0e, 0x47, 0x17, 0xfd, 0x3d, 0xba, 0xe8, 0xe7, 0xc9, 0x1d, 0x1c, 0x4e, 0xee, 0xe0,
	0xcf, 0xc9, 0x1d, 0x7c, 0x0e, 0x72, 0xa6, 0x76, 0x75, 0xe2, 0xa7, 0xe2, 0x4b, 0xf0, 0x31, 0x66,
	0x32, 0x29, 0x44, 0xba, 0x0f, 0x2e, 0x47, 0xfc, 0xed, 0x72, 0xc6, 0xaa, 0x29, 0x69, 0x95, 0x8c,
	0xcc, 0xe9, 0xbe, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xcb, 0xdb, 0x7d, 0xe9, 0x02, 0x00,
	0x00,
}

func (m *KeyShareRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyShareRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyShareRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sent {
		i--
		if m.Sent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ProposalId) > 0 {
		i -= len(m.ProposalId)
		copy(dAtA[i:], m.ProposalId)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ProposalId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AggrKeyshare) > 0 {
		i -= len(m.AggrKeyshare)
		copy(dAtA[i:], m.AggrKeyshare)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.AggrKeyshare)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Counterparty != nil {
		{
			size, err := m.Counterparty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequestedKeyshare(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.IbcInfo != nil {
		{
			size, err := m.IbcInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequestedKeyshare(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PortID) > 0 {
		i -= len(m.PortID)
		copy(dAtA[i:], m.PortID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.PortID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionID) > 0 {
		i -= len(m.ConnectionID)
		copy(dAtA[i:], m.ConnectionID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ConnectionID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CounterPartyIBCInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CounterPartyIBCInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CounterPartyIBCInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PortID) > 0 {
		i -= len(m.PortID)
		copy(dAtA[i:], m.PortID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.PortID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionID) > 0 {
		i -= len(m.ConnectionID)
		copy(dAtA[i:], m.ConnectionID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ConnectionID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintRequestedKeyshare(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequestedKeyshare(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequestedKeyshare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyShareRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	if m.IbcInfo != nil {
		l = m.IbcInfo.Size()
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	if m.Counterparty != nil {
		l = m.Counterparty.Size()
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.AggrKeyshare)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.ProposalId)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	if m.Sent {
		n += 2
	}
	return n
}

func (m *IBCInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.ConnectionID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.PortID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	return n
}

func (m *CounterPartyIBCInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.ConnectionID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	l = len(m.PortID)
	if l > 0 {
		n += 1 + l + sovRequestedKeyshare(uint64(l))
	}
	return n
}

func sovRequestedKeyshare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequestedKeyshare(x uint64) (n int) {
	return sovRequestedKeyshare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyShareRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestedKeyshare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyShareRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyShareRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IbcInfo == nil {
				m.IbcInfo = &IBCInfo{}
			}
			if err := m.IbcInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counterparty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Counterparty == nil {
				m.Counterparty = &CounterPartyIBCInfo{}
			}
			if err := m.Counterparty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggrKeyshare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggrKeyshare = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sent = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRequestedKeyshare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestedKeyshare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestedKeyshare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CounterPartyIBCInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestedKeyshare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CounterPartyIBCInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CounterPartyIBCInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestedKeyshare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestedKeyshare
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRequestedKeyshare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequestedKeyshare
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRequestedKeyshare
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRequestedKeyshare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequestedKeyshare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequestedKeyshare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequestedKeyshare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequestedKeyshare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequestedKeyshare = fmt.Errorf("proto: unexpected end of group")
)
