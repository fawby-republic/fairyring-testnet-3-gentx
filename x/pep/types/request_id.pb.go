// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/request_id.proto

package types

import (
	fmt "fmt"
	types "github.com/Fairblock/fairyring/x/common/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type RequestId struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ReqId   string `protobuf:"bytes,2,opt,name=reqId,proto3" json:"reqId,omitempty"`
}

func (m *RequestId) Reset()         { *m = RequestId{} }
func (m *RequestId) String() string { return proto.CompactTextString(m) }
func (*RequestId) ProtoMessage()    {}
func (*RequestId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e457d2e8ff0411e, []int{0}
}
func (m *RequestId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestId.Merge(m, src)
}
func (m *RequestId) XXX_Size() int {
	return m.Size()
}
func (m *RequestId) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestId.DiscardUnknown(m)
}

var xxx_messageInfo_RequestId proto.InternalMessageInfo

func (m *RequestId) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RequestId) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

type PrivateRequest struct {
	Creator            string                     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ReqId              string                     `protobuf:"bytes,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	Pubkey             string                     `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	EncryptedKeyshares []*types.EncryptedKeyshare `protobuf:"bytes,5,rep,name=encrypted_keyshares,json=encryptedKeyshares,proto3" json:"encrypted_keyshares,omitempty"`
}

func (m *PrivateRequest) Reset()         { *m = PrivateRequest{} }
func (m *PrivateRequest) String() string { return proto.CompactTextString(m) }
func (*PrivateRequest) ProtoMessage()    {}
func (*PrivateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e457d2e8ff0411e, []int{1}
}
func (m *PrivateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrivateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrivateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrivateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateRequest.Merge(m, src)
}
func (m *PrivateRequest) XXX_Size() int {
	return m.Size()
}
func (m *PrivateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateRequest proto.InternalMessageInfo

func (m *PrivateRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PrivateRequest) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func (m *PrivateRequest) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *PrivateRequest) GetEncryptedKeyshares() []*types.EncryptedKeyshare {
	if m != nil {
		return m.EncryptedKeyshares
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestId)(nil), "fairyring.pep.RequestId")
	proto.RegisterType((*PrivateRequest)(nil), "fairyring.pep.PrivateRequest")
}

func init() { proto.RegisterFile("fairyring/pep/request_id.proto", fileDescriptor_3e457d2e8ff0411e) }

var fileDescriptor_3e457d2e8ff0411e = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x6b, 0xaa, 0x16, 0xd5, 0x08, 0x06, 0x53, 0x50, 0xd4, 0xc1, 0xaa, 0xda, 0xa5, 0x62,
	0x88, 0x25, 0x18, 0xd9, 0x10, 0x20, 0x55, 0x2c, 0xa8, 0x62, 0x62, 0x89, 0xf2, 0xf3, 0x91, 0x5a,
	0xa5, 0xb1, 0xf3, 0xc5, 0x41, 0xe4, 0x16, 0xdc, 0x84, 0x6b, 0x30, 0x76, 0x64, 0x44, 0xc9, 0x45,
	0x50, 0x9d, 0x94, 0x08, 0x06, 0x36, 0xbf, 0x3f, 0x7a, 0x64, 0xbf, 0xa6, 0xfc, 0xc9, 0x97, 0x58,
	0xa0, 0x4c, 0x62, 0xa1, 0x41, 0x0b, 0x84, 0x34, 0x87, 0xcc, 0x78, 0x32, 0x72, 0x35, 0x2a, 0xa3,
	0xd8, 0xe1, 0x4f, 0xee, 0x6a, 0xd0, 0xa3, 0x61, 0xac, 0x62, 0x65, 0x13, 0xb1, 0x3d, 0xd5, 0xa5,
	0xd1, 0xb4, 0x85, 0x84, 0x6a, 0xbd, 0x56, 0x89, 0xc8, 0x96, 0x3e, 0x42, 0xe4, 0x99, 0x42, 0x43,
	0x56, 0x97, 0x26, 0x97, 0x74, 0xb0, 0xa8, 0xe9, 0xf3, 0x88, 0x39, 0x74, 0x3f, 0x44, 0xf0, 0x8d,
	0x42, 0x87, 0x8c, 0xc9, 0x6c, 0xb0, 0xd8, 0x49, 0x36, 0xa4, 0x3d, 0x84, 0x74, 0x1e, 0x39, 0x7b,
	0xd6, 0xaf, 0xc5, 0xe4, 0x9d, 0xd0, 0xa3, 0x7b, 0x94, 0x2f, 0xbe, 0x81, 0x06, 0xf2, 0x0f, 0xe2,
	0x84, 0xf6, 0x11, 0x52, 0x4f, 0xfe, 0x66, 0xb0, 0x53, 0xda, 0xd7, 0x79, 0xb0, 0x82, 0xc2, 0xe9,
	0x5a, 0xbb, 0x51, 0xec, 0x81, 0x1e, 0x43, 0x12, 0x62, 0xa1, 0x0d, 0x44, 0xde, 0x0a, 0x0a, 0x7b,
	0xf7, 0xcc, 0xe9, 0x8d, 0xbb, 0xb3, 0x83, 0xf3, 0xa9, 0xdb, 0x0e, 0x50, 0xbf, 0xcd, 0xbd, 0xd9,
	0x95, 0xef, 0x9a, 0xee, 0x82, 0xc1, 0x5f, 0x2b, 0xbb, 0xba, 0xfe, 0x28, 0x39, 0xd9, 0x94, 0x9c,
	0x7c, 0x95, 0x9c, 0xbc, 0x55, 0xbc, 0xb3, 0xa9, 0x78, 0xe7, 0xb3, 0xe2, 0x9d, 0xc7, 0xb3, 0x58,
	0x9a, 0x65, 0x1e, 0x6c, 0x71, 0xe2, 0xd6, 0x97, 0x18, 0x3c, 0xab, 0x70, 0x25, 0xda, 0x09, 0x5f,
	0xed, 0x4f, 0xd8, 0xe9, 0x82, 0xbe, 0xdd, 0xee, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x14, 0x1c,
	0xef, 0xb3, 0xa7, 0x01, 0x00, 0x00,
}

func (m *RequestId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReqId) > 0 {
		i -= len(m.ReqId)
		copy(dAtA[i:], m.ReqId)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.ReqId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PrivateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrivateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedKeyshares) > 0 {
		for iNdEx := len(m.EncryptedKeyshares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EncryptedKeyshares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRequestId(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ReqId) > 0 {
		i -= len(m.ReqId)
		copy(dAtA[i:], m.ReqId)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.ReqId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequestId(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequestId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = len(m.ReqId)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	return n
}

func (m *PrivateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = len(m.ReqId)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	if len(m.EncryptedKeyshares) > 0 {
		for _, e := range m.EncryptedKeyshares {
			l = e.Size()
			n += 1 + l + sovRequestId(uint64(l))
		}
	}
	return n
}

func sovRequestId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequestId(x uint64) (n int) {
	return sovRequestId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestId
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
			return fmt.Errorf("proto: RequestId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestId
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
func (m *PrivateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestId
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
			return fmt.Errorf("proto: PrivateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKeyshares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedKeyshares = append(m.EncryptedKeyshares, &types.EncryptedKeyshare{})
			if err := m.EncryptedKeyshares[len(m.EncryptedKeyshares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestId
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
func skipRequestId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequestId
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
					return 0, ErrIntOverflowRequestId
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
					return 0, ErrIntOverflowRequestId
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
				return 0, ErrInvalidLengthRequestId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequestId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequestId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequestId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequestId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequestId = fmt.Errorf("proto: unexpected end of group")
)