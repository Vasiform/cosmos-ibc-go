// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/middleware/fee/v1/fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Fee interface
// See Fee Payment Middleware spec:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-029-fee-payment#fee-middleware-contract
type Fee struct {
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a4c0273395d2c1a, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

// Fee associated with a packet_id
type IdentifiedPacketFee struct {
	PacketId *types1.PacketId `protobuf:"bytes,1,opt,name=packet_id,json=packetId,proto3" json:"packet_id,omitempty"`
	Fee      *Fee             `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *IdentifiedPacketFee) Reset()         { *m = IdentifiedPacketFee{} }
func (m *IdentifiedPacketFee) String() string { return proto.CompactTextString(m) }
func (*IdentifiedPacketFee) ProtoMessage()    {}
func (*IdentifiedPacketFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a4c0273395d2c1a, []int{1}
}
func (m *IdentifiedPacketFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedPacketFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedPacketFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedPacketFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedPacketFee.Merge(m, src)
}
func (m *IdentifiedPacketFee) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedPacketFee) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedPacketFee.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedPacketFee proto.InternalMessageInfo

func (m *IdentifiedPacketFee) GetPacketId() *types1.PacketId {
	if m != nil {
		return m.PacketId
	}
	return nil
}

func (m *IdentifiedPacketFee) GetFee() *Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}

func init() {
	proto.RegisterType((*Fee)(nil), "ibc.applications.middleware.fee.v1.Fee")
	proto.RegisterType((*IdentifiedPacketFee)(nil), "ibc.applications.middleware.fee.v1.IdentifiedPacketFee")
}

func init() {
	proto.RegisterFile("ibc/applications/middleware/fee/v1/fee.proto", fileDescriptor_9a4c0273395d2c1a)
}

var fileDescriptor_9a4c0273395d2c1a = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xb1, 0x4e, 0xeb, 0x40,
	0x10, 0xb4, 0x5f, 0xa4, 0xe8, 0x3d, 0xa7, 0xcb, 0xa3, 0x08, 0x91, 0x70, 0x42, 0x1a, 0x52, 0x90,
	0x3d, 0x6c, 0xaa, 0x50, 0x06, 0x29, 0x52, 0x24, 0x0a, 0x94, 0x92, 0x06, 0x9d, 0xef, 0x36, 0xce,
	0x11, 0xdb, 0x6b, 0xe5, 0x2e, 0x41, 0xfc, 0x03, 0x05, 0xdf, 0xc1, 0x97, 0xa4, 0x4c, 0x49, 0x05,
	0x28, 0xf9, 0x11, 0x74, 0xb6, 0x11, 0x48, 0x14, 0x54, 0x3b, 0x77, 0x9a, 0xd9, 0x99, 0xdd, 0xf5,
	0x4e, 0x55, 0x24, 0x18, 0xcf, 0xf3, 0x44, 0x09, 0x6e, 0x14, 0x65, 0x9a, 0xa5, 0x4a, 0xca, 0x04,
	0xef, 0xf9, 0x12, 0xd9, 0x0c, 0x91, 0xad, 0x03, 0x5b, 0x20, 0x5f, 0x92, 0xa1, 0x66, 0x4f, 0x45,
	0x02, 0xbe, 0xb3, 0xe1, 0x8b, 0x0d, 0x96, 0xb6, 0x0e, 0xda, 0xbe, 0x20, 0x9d, 0x92, 0x66, 0x11,
	0xd7, 0x56, 0x1d, 0xa1, 0xe1, 0x01, 0x13, 0xa4, 0xb2, 0xb2, 0x47, 0xfb, 0x20, 0xa6, 0x98, 0x0a,
	0xc8, 0x2c, 0xaa, 0x7e, 0x8f, 0x6d, 0x0e, 0x41, 0x4b, 0x64, 0x62, 0xce, 0xb3, 0x0c, 0x13, 0x6b,
	0x5c, 0xc1, 0x92, 0xd2, 0xbb, 0xf3, 0x6a, 0x63, 0xc4, 0xa6, 0xf0, 0xea, 0x3c, 0xa5, 0x55, 0x66,
	0x5a, 0x6e, 0xb7, 0xd6, 0x6f, 0x84, 0x87, 0x50, 0x1a, 0x82, 0x35, 0x84, 0xca, 0x10, 0x2e, 0x49,
	0x65, 0xa3, 0xb3, 0xcd, 0x6b, 0xc7, 0x79, 0x7e, 0xeb, 0xf4, 0x63, 0x65, 0xe6, 0xab, 0x08, 0x04,
	0xa5, 0xac, 0x4a, 0x57, 0x96, 0x81, 0x96, 0x0b, 0x66, 0x1e, 0x72, 0xd4, 0x85, 0x40, 0x4f, 0xab,
	0xd6, 0xbd, 0x47, 0xd7, 0xfb, 0x3f, 0x91, 0x98, 0x19, 0x35, 0x53, 0x28, 0xaf, 0xb9, 0x58, 0xa0,
	0xb1, 0xe6, 0x17, 0xde, 0xbf, 0xbc, 0x78, 0xdc, 0x2a, 0xd9, 0x72, 0xbb, 0x6e, 0xbf, 0x11, 0x1e,
	0x81, 0x5d, 0x8a, 0x8d, 0x0e, 0x9f, 0x79, 0xd7, 0x01, 0x94, 0x92, 0x89, 0x9c, 0xfe, 0xcd, 0x2b,
	0xd4, 0x1c, 0x7a, 0xb5, 0x19, 0x62, 0xeb, 0x4f, 0xa1, 0x3a, 0x81, 0xdf, 0x57, 0x09, 0x63, 0xc4,
	0xa9, 0xd5, 0x8c, 0xae, 0x36, 0x3b, 0xdf, 0xdd, 0xee, 0x7c, 0xf7, 0x7d, 0xe7, 0xbb, 0x4f, 0x7b,
	0xdf, 0xd9, 0xee, 0x7d, 0xe7, 0x65, 0xef, 0x3b, 0x37, 0xe1, 0xcf, 0xd1, 0x54, 0x24, 0x06, 0x31,
	0xb1, 0x94, 0xe4, 0x2a, 0x41, 0x6d, 0x8f, 0xab, 0x59, 0x38, 0x1c, 0xd8, 0x63, 0x16, 0xa3, 0x46,
	0xf5, 0x62, 0x9f, 0xe7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x07, 0x16, 0x98, 0xc3, 0xfc, 0x01,
	0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedPacketFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedPacketFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedPacketFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PacketId != nil {
		{
			size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *IdentifiedPacketFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketId != nil {
		l = m.PacketId.Size()
		n += 1 + l + sovFee(uint64(l))
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovFee(uint64(l))
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *IdentifiedPacketFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: IdentifiedPacketFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedPacketFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PacketId == nil {
				m.PacketId = &types1.PacketId{}
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &Fee{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)
