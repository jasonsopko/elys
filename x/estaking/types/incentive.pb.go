// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/estaking/incentive.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Incentive Info
type IncentiveInfo struct {
	// reward amount in eden for 1 year
	EdenAmountPerYear github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=eden_amount_per_year,json=edenAmountPerYear,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eden_amount_per_year"`
	// starting block height of the distribution
	DistributionStartBlock github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=distribution_start_block,json=distributionStartBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"distribution_start_block"`
	// distribution duration - block number per year
	TotalBlocksPerYear github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=total_blocks_per_year,json=totalBlocksPerYear,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_blocks_per_year"`
	// blocks distributed
	BlocksDistributed github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=blocks_distributed,json=blocksDistributed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"blocks_distributed"`
}

func (m *IncentiveInfo) Reset()         { *m = IncentiveInfo{} }
func (m *IncentiveInfo) String() string { return proto.CompactTextString(m) }
func (*IncentiveInfo) ProtoMessage()    {}
func (*IncentiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9de1d3bcda4768b, []int{0}
}
func (m *IncentiveInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveInfo.Merge(m, src)
}
func (m *IncentiveInfo) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IncentiveInfo)(nil), "elys.estaking.IncentiveInfo")
}

func init() { proto.RegisterFile("elys/estaking/incentive.proto", fileDescriptor_e9de1d3bcda4768b) }

var fileDescriptor_e9de1d3bcda4768b = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0x00, 0xe0, 0xe4, 0xef, 0xcf, 0x0f, 0x7f, 0xa0, 0x07, 0x43, 0x95, 0x50, 0x30, 0x15, 0x0f,
	0xe2, 0xa5, 0xd9, 0x83, 0x4f, 0x60, 0x11, 0xa4, 0x37, 0xd1, 0x93, 0x82, 0x84, 0x4d, 0x32, 0x4d,
	0x97, 0x24, 0x3b, 0x61, 0x77, 0xa2, 0xf6, 0x2d, 0x7c, 0x00, 0x1f, 0xa8, 0xc7, 0x1e, 0xc5, 0x43,
	0x91, 0xf6, 0x45, 0x64, 0x37, 0xad, 0xed, 0xb9, 0xa7, 0x64, 0x98, 0xc9, 0x37, 0x99, 0xdd, 0xf1,
	0x4e, 0xa1, 0x9c, 0x69, 0x06, 0x9a, 0x78, 0x21, 0x64, 0xce, 0x84, 0x4c, 0x41, 0x92, 0x78, 0x81,
	0xa8, 0x56, 0x48, 0xe8, 0x77, 0x4d, 0x3a, 0xda, 0xa6, 0xfb, 0xbd, 0x1c, 0x73, 0xb4, 0x19, 0x66,
	0xde, 0xda, 0xa2, 0xfe, 0x20, 0x47, 0xcc, 0x4b, 0x60, 0x36, 0x4a, 0x9a, 0x09, 0x23, 0x51, 0x99,
	0x4f, 0xaa, 0xba, 0x2d, 0x38, 0xff, 0xe8, 0x78, 0xdd, 0xf1, 0x56, 0x1e, 0xcb, 0x09, 0xfa, 0xb1,
	0xd7, 0x83, 0x0c, 0x64, 0xcc, 0x2b, 0x6c, 0x24, 0xc5, 0x35, 0xa8, 0x78, 0x06, 0x5c, 0x05, 0xee,
	0x99, 0x7b, 0xf9, 0x7f, 0x14, 0xcd, 0x97, 0x03, 0xe7, 0x6b, 0x39, 0xb8, 0xc8, 0x05, 0x4d, 0x9b,
	0x24, 0x4a, 0xb1, 0x62, 0x29, 0xea, 0x0a, 0xf5, 0xe6, 0x31, 0xd4, 0x59, 0xc1, 0x68, 0x56, 0x83,
	0x8e, 0xc6, 0x92, 0xee, 0x8f, 0x8c, 0x75, 0x6d, 0xa9, 0x3b, 0x50, 0x8f, 0xc0, 0x95, 0x3f, 0xf5,
	0x82, 0x4c, 0x68, 0x52, 0x22, 0x69, 0x48, 0xa0, 0x8c, 0x35, 0x71, 0x45, 0x71, 0x52, 0x62, 0x5a,
	0x04, 0x7f, 0x0e, 0x6a, 0x72, 0xb2, 0xef, 0x3d, 0x18, 0x6e, 0x64, 0x34, 0x9f, 0x7b, 0xc7, 0x84,
	0xc4, 0xcb, 0x16, 0xd7, 0xbb, 0x59, 0x3a, 0x07, 0xb5, 0xf1, 0x2d, 0x66, 0x69, 0xbd, 0x1d, 0xe6,
	0xd9, 0xf3, 0x37, 0xf8, 0xef, 0x3f, 0x40, 0x16, 0xfc, 0x3d, 0xec, 0xac, 0x5a, 0xe9, 0x66, 0x07,
	0x8d, 0x6e, 0xe7, 0xab, 0xd0, 0x5d, 0xac, 0x42, 0xf7, 0x7b, 0x15, 0xba, 0xef, 0xeb, 0xd0, 0x59,
	0xac, 0x43, 0xe7, 0x73, 0x1d, 0x3a, 0x4f, 0xc3, 0x3d, 0xd4, 0x6c, 0xc2, 0x50, 0x02, 0xbd, 0xa2,
	0x2a, 0x6c, 0xc0, 0xde, 0x76, 0x7b, 0x63, 0xfd, 0xe4, 0x9f, 0xbd, 0xee, 0xab, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc2, 0x59, 0x9d, 0x93, 0x55, 0x02, 0x00, 0x00,
}

func (m *IncentiveInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BlocksDistributed.Size()
		i -= size
		if _, err := m.BlocksDistributed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TotalBlocksPerYear.Size()
		i -= size
		if _, err := m.TotalBlocksPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DistributionStartBlock.Size()
		i -= size
		if _, err := m.DistributionStartBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EdenAmountPerYear.Size()
		i -= size
		if _, err := m.EdenAmountPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IncentiveInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EdenAmountPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.DistributionStartBlock.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.TotalBlocksPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.BlocksDistributed.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IncentiveInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenAmountPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenAmountPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStartBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionStartBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBlocksPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBlocksPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksDistributed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlocksDistributed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
