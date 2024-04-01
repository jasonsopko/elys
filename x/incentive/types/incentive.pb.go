// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/incentive/incentive.proto

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
	// unused
	EpochNumBlocks github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=epoch_num_blocks,json=epochNumBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"epoch_num_blocks"`
	// maximum eden allocation per day that won't exceed 30% apr
	MaxEdenPerAllocation github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=max_eden_per_allocation,json=maxEdenPerAllocation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_eden_per_allocation"`
	// unused
	DistributionEpochInBlocks github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=distribution_epoch_in_blocks,json=distributionEpochInBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"distribution_epoch_in_blocks"`
	// current epoch in block number
	CurrentEpochInBlocks github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=current_epoch_in_blocks,json=currentEpochInBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_epoch_in_blocks"`
}

func (m *IncentiveInfo) Reset()         { *m = IncentiveInfo{} }
func (m *IncentiveInfo) String() string { return proto.CompactTextString(m) }
func (*IncentiveInfo) ProtoMessage()    {}
func (*IncentiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0e67c7f36f3313, []int{0}
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
	proto.RegisterType((*IncentiveInfo)(nil), "elys.incentive.IncentiveInfo")
}

func init() { proto.RegisterFile("elys/incentive/incentive.proto", fileDescriptor_ed0e67c7f36f3313) }

var fileDescriptor_ed0e67c7f36f3313 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x86, 0x13, 0x75, 0x57, 0x6c, 0x70, 0xd1, 0x30, 0x6a, 0x5c, 0x24, 0x23, 0x1e, 0xc4, 0xcb,
	0x26, 0x07, 0x9f, 0x60, 0x07, 0x16, 0xcc, 0x45, 0x16, 0xbd, 0xa8, 0x97, 0xa6, 0x93, 0xd4, 0x66,
	0x9a, 0x49, 0x77, 0x85, 0xee, 0x8a, 0xce, 0xbc, 0x85, 0x8f, 0xb5, 0xc7, 0x3d, 0x89, 0x78, 0x58,
	0x64, 0xe6, 0x45, 0xa4, 0x3b, 0x33, 0x3b, 0x19, 0xbc, 0xe5, 0x94, 0xee, 0x54, 0xf1, 0x7d, 0xfd,
	0x53, 0x14, 0x4b, 0xa0, 0x59, 0xd9, 0x4c, 0xea, 0x12, 0x34, 0xc9, 0xef, 0xb0, 0x3f, 0xa5, 0xad,
	0x41, 0xc2, 0xe8, 0xc4, 0xd5, 0xd3, 0xbb, 0xbf, 0xa7, 0x93, 0x1a, 0x6b, 0xf4, 0xa5, 0xcc, 0x9d,
	0xfa, 0xae, 0xd3, 0x69, 0x8d, 0x58, 0x37, 0x90, 0xf9, 0x5b, 0xd1, 0x5d, 0x65, 0x24, 0x15, 0x58,
	0x12, 0xaa, 0xed, 0x1b, 0xde, 0xfc, 0x3a, 0x62, 0x8f, 0xf3, 0x1d, 0x24, 0xd7, 0x57, 0x18, 0x71,
	0x36, 0x81, 0x0a, 0x34, 0x17, 0x0a, 0x3b, 0x4d, 0xbc, 0x05, 0xc3, 0x57, 0x20, 0x4c, 0x1c, 0xbe,
	0x0e, 0xdf, 0x3d, 0x9a, 0xa5, 0xd7, 0xb7, 0xd3, 0xe0, 0xcf, 0xed, 0xf4, 0x6d, 0x2d, 0x69, 0xde,
	0x15, 0x69, 0x89, 0x2a, 0x2b, 0xd1, 0x2a, 0xb4, 0xdb, 0xcf, 0x99, 0xad, 0x16, 0x19, 0xad, 0x5a,
	0xb0, 0x69, 0xae, 0xe9, 0xd3, 0x53, 0xc7, 0x3a, 0xf7, 0xa8, 0x4b, 0x30, 0x5f, 0x41, 0x98, 0x68,
	0xce, 0xe2, 0x4a, 0x5a, 0x32, 0xb2, 0xe8, 0x48, 0xa2, 0xe6, 0x96, 0x84, 0x21, 0x5e, 0x34, 0x58,
	0x2e, 0xe2, 0x7b, 0xa3, 0x24, 0xcf, 0x87, 0xbc, 0xcf, 0x0e, 0x37, 0x73, 0xb4, 0x48, 0xb0, 0x67,
	0x84, 0x24, 0x9a, 0x1e, 0x6e, 0xf7, 0x59, 0xee, 0x8f, 0xd2, 0x44, 0x1e, 0xe6, 0xd1, 0x76, 0x17,
	0xe6, 0x0b, 0x7b, 0x02, 0x2d, 0x96, 0x73, 0xae, 0x3b, 0xb5, 0xd5, 0xc4, 0x0f, 0x46, 0xd1, 0x4f,
	0x3c, 0xe7, 0x63, 0xa7, 0x7a, 0x41, 0x04, 0xec, 0x85, 0x12, 0x4b, 0xee, 0x67, 0xe1, 0x1e, 0x2e,
	0x9a, 0x06, 0x4b, 0xe1, 0x12, 0xc6, 0x47, 0xa3, 0x04, 0x13, 0x25, 0x96, 0x17, 0x15, 0xe8, 0x4b,
	0x30, 0xe7, 0x77, 0xac, 0x08, 0xd9, 0xab, 0x83, 0x69, 0xf4, 0x69, 0xa4, 0xde, 0x85, 0x39, 0x1e,
	0xe5, 0x7a, 0x39, 0x64, 0x5e, 0x38, 0x64, 0xae, 0xf7, 0xb9, 0xca, 0xce, 0x18, 0xd0, 0xf4, 0x9f,
	0xeb, 0xe1, 0xb8, 0x5c, 0x5b, 0xdc, 0x81, 0x66, 0xf6, 0xe1, 0x7a, 0x9d, 0x84, 0x37, 0xeb, 0x24,
	0xfc, 0xbb, 0x4e, 0xc2, 0x9f, 0x9b, 0x24, 0xb8, 0xd9, 0x24, 0xc1, 0xef, 0x4d, 0x12, 0x7c, 0x4b,
	0x07, 0x5c, 0xb7, 0x44, 0x67, 0x1a, 0xe8, 0x07, 0x9a, 0x85, 0xbf, 0x64, 0xcb, 0xc1, 0xce, 0x79,
	0x47, 0x71, 0xec, 0x37, 0xe5, 0xfd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x65, 0x62, 0xc6,
	0x92, 0x03, 0x00, 0x00,
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
		size := m.CurrentEpochInBlocks.Size()
		i -= size
		if _, err := m.CurrentEpochInBlocks.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.DistributionEpochInBlocks.Size()
		i -= size
		if _, err := m.DistributionEpochInBlocks.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.MaxEdenPerAllocation.Size()
		i -= size
		if _, err := m.MaxEdenPerAllocation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.EpochNumBlocks.Size()
		i -= size
		if _, err := m.EpochNumBlocks.MarshalTo(dAtA[i:]); err != nil {
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
	l = m.EpochNumBlocks.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.MaxEdenPerAllocation.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.DistributionEpochInBlocks.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.CurrentEpochInBlocks.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumBlocks", wireType)
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
			if err := m.EpochNumBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEdenPerAllocation", wireType)
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
			if err := m.MaxEdenPerAllocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionEpochInBlocks", wireType)
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
			if err := m.DistributionEpochInBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochInBlocks", wireType)
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
			if err := m.CurrentEpochInBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
