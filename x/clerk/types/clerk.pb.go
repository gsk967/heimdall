// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/clerk/v1beta1/clerk.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
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

type EventRecord struct {
	Id         uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id" yaml:"id"`
	Contract   string    `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Data       []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	RecordTime time.Time `protobuf:"bytes,4,opt,name=record_time,json=recordTime,proto3,stdtime" json:"record_time" yaml:"record_time"`
	LogIndex   uint64    `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty" yaml:"log_index"`
	TxHash     string    `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"tx_hash"`
	ChainId    string    `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
}

func (m *EventRecord) Reset()         { *m = EventRecord{} }
func (m *EventRecord) String() string { return proto.CompactTextString(m) }
func (*EventRecord) ProtoMessage()    {}
func (*EventRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_88fd2b9ce5955508, []int{0}
}
func (m *EventRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRecord.Merge(m, src)
}
func (m *EventRecord) XXX_Size() int {
	return m.Size()
}
func (m *EventRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRecord.DiscardUnknown(m)
}

var xxx_messageInfo_EventRecord proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRecord)(nil), "heimdall.clerk.v1beta1.EventRecord")
}

func init() {
	proto.RegisterFile("heimdall/clerk/v1beta1/clerk.proto", fileDescriptor_88fd2b9ce5955508)
}

var fileDescriptor_88fd2b9ce5955508 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x50, 0xfa, 0xc7, 0x45, 0x80, 0x7c, 0x27, 0x14, 0x75, 0x88, 0x2b, 0xb3, 0x54,
	0x42, 0x72, 0x54, 0xd8, 0x6e, 0xac, 0x84, 0x74, 0xb7, 0x5a, 0x4c, 0x30, 0x54, 0x4e, 0x6c, 0x12,
	0xeb, 0x92, 0xb8, 0x4a, 0x7c, 0x47, 0xee, 0x1b, 0x30, 0xde, 0x47, 0xe0, 0xe3, 0xdc, 0xd8, 0x05,
	0x89, 0x29, 0xa0, 0x76, 0x63, 0xcc, 0x27, 0x40, 0xb1, 0x9b, 0x8a, 0xed, 0x79, 0xdf, 0xf7, 0xf7,
	0xf8, 0x7d, 0xe4, 0x17, 0x92, 0x4c, 0xaa, 0x42, 0xf0, 0x3c, 0x8f, 0x92, 0x5c, 0x56, 0xb7, 0xd1,
	0xfd, 0x3a, 0x96, 0x86, 0xaf, 0x5d, 0x45, 0x77, 0x95, 0x36, 0x1a, 0xbd, 0x19, 0x18, 0xea, 0xba,
	0x27, 0x66, 0x71, 0x99, 0xea, 0x54, 0x5b, 0x24, 0xea, 0x95, 0xa3, 0x17, 0x38, 0xd5, 0x3a, 0xcd,
	0x65, 0x64, 0xab, 0xf8, 0xee, 0x6b, 0x64, 0x54, 0x21, 0x6b, 0xc3, 0x8b, 0x9d, 0x03, 0xc8, 0x4f,
	0x1f, 0xce, 0x3f, 0xde, 0xcb, 0xd2, 0x30, 0x99, 0xe8, 0x4a, 0xa0, 0xb7, 0xd0, 0x57, 0x22, 0x00,
	0x4b, 0xb0, 0x1a, 0x6d, 0x2e, 0xfe, 0xb6, 0xd8, 0x57, 0xa2, 0x6b, 0xf1, 0xec, 0x81, 0x17, 0xf9,
	0x15, 0x51, 0x82, 0x30, 0x5f, 0x09, 0xb4, 0x80, 0xd3, 0x44, 0x97, 0xa6, 0xe2, 0x89, 0x09, 0xfc,
	0x25, 0x58, 0xcd, 0xd8, 0xb9, 0x46, 0x08, 0x8e, 0x04, 0x37, 0x3c, 0x78, 0xb6, 0x04, 0xab, 0x17,
	0xcc, 0x6a, 0xf4, 0x05, 0xce, 0x2b, 0xfb, 0xfc, 0xb6, 0x5f, 0x1f, 0x8c, 0x96, 0x60, 0x35, 0x7f,
	0xbf, 0xa0, 0x2e, 0x1b, 0x1d, 0xb2, 0xd1, 0x4f, 0x43, 0xb6, 0x4d, 0xf8, 0xd4, 0x62, 0xaf, 0x6b,
	0x31, 0x72, 0x7b, 0xff, 0x33, 0x93, 0xc7, 0xdf, 0x18, 0x30, 0xe8, 0x3a, 0xbd, 0x01, 0xad, 0xe1,
	0x2c, 0xd7, 0xe9, 0x56, 0x95, 0x42, 0x36, 0xc1, 0x73, 0x1b, 0xfc, 0xb2, 0x6b, 0xf1, 0x6b, 0x67,
	0x3d, 0x8f, 0x08, 0x9b, 0xe6, 0x3a, 0xbd, 0xe9, 0x25, 0x7a, 0x07, 0x27, 0xa6, 0xd9, 0x66, 0xbc,
	0xce, 0x82, 0x71, 0x1f, 0x7f, 0x83, 0xba, 0x16, 0xbf, 0x74, 0x86, 0xd3, 0x80, 0xb0, 0xb1, 0x69,
	0xae, 0x79, 0x9d, 0x21, 0x0a, 0xa7, 0x49, 0xc6, 0x55, 0xb9, 0x55, 0x22, 0x98, 0x58, 0xfa, 0xa2,
	0x6b, 0xf1, 0x2b, 0x47, 0x0f, 0x13, 0xc2, 0x26, 0x56, 0xde, 0x88, 0xab, 0xd1, 0xf7, 0x1f, 0xd8,
	0xdb, 0x5c, 0x3f, 0x1d, 0x42, 0xb0, 0x3f, 0x84, 0xe0, 0xcf, 0x21, 0x04, 0x8f, 0xc7, 0xd0, 0xdb,
	0x1f, 0x43, 0xef, 0xd7, 0x31, 0xf4, 0x3e, 0xd3, 0x54, 0x99, 0xec, 0x2e, 0xa6, 0x89, 0x2e, 0xa2,
	0x82, 0x1b, 0x95, 0x94, 0xd2, 0x7c, 0xd3, 0xd5, 0x6d, 0x74, 0x3e, 0x7e, 0x73, 0x3a, 0xbf, 0x79,
	0xd8, 0xc9, 0x3a, 0x1e, 0xdb, 0xff, 0xf9, 0xf0, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xfb, 0xc4,
	0xad, 0x1d, 0x02, 0x00, 0x00,
}

func (m *EventRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.LogIndex != 0 {
		i = encodeVarintClerk(dAtA, i, uint64(m.LogIndex))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RecordTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RecordTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClerk(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintClerk(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClerk(dAtA []byte, offset int, v uint64) int {
	offset -= sovClerk(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClerk(uint64(m.Id))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RecordTime)
	n += 1 + l + sovClerk(uint64(l))
	if m.LogIndex != 0 {
		n += 1 + sovClerk(uint64(m.LogIndex))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	return n
}

func sovClerk(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClerk(x uint64) (n int) {
	return sovClerk(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClerk
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
			return fmt.Errorf("proto: EventRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
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
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
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
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RecordTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogIndex", wireType)
			}
			m.LogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
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
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
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
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClerk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClerk
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClerk
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
func skipClerk(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClerk
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
					return 0, ErrIntOverflowClerk
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
					return 0, ErrIntOverflowClerk
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
				return 0, ErrInvalidLengthClerk
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClerk
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClerk
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClerk        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClerk          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClerk = fmt.Errorf("proto: unexpected end of group")
)