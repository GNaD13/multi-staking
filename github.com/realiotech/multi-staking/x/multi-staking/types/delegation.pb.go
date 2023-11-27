// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multistaking/v1/delegation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MultiStakingLock struct {
	ConversionRatio  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=conversion_ratio,json=conversionRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"conversion_ratio"`
	LockedAmount     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=locked_amount,json=lockedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"locked_amount"`
	DelegatorAddress string                                 `protobuf:"bytes,3,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	ValidatorAddress string                                 `protobuf:"bytes,4,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *MultiStakingLock) Reset()         { *m = MultiStakingLock{} }
func (m *MultiStakingLock) String() string { return proto.CompactTextString(m) }
func (*MultiStakingLock) ProtoMessage()    {}
func (*MultiStakingLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dade11e2944357, []int{0}
}
func (m *MultiStakingLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiStakingLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiStakingLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiStakingLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiStakingLock.Merge(m, src)
}
func (m *MultiStakingLock) XXX_Size() int {
	return m.Size()
}
func (m *MultiStakingLock) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiStakingLock.DiscardUnknown(m)
}

var xxx_messageInfo_MultiStakingLock proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MultiStakingLock)(nil), "multistaking.v1.MultiStakingLock")
}

func init() { proto.RegisterFile("multistaking/v1/delegation.proto", fileDescriptor_e9dade11e2944357) }

var fileDescriptor_e9dade11e2944357 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4, 0x4f, 0x49, 0xcd, 0x49,
	0x4d, 0x4f, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x56,
	0xa1, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd3, 0x07, 0xb1, 0x20, 0xca,
	0xa4, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21, 0x12, 0x10, 0x0e, 0x44, 0x4a, 0xe9,
	0x10, 0x13, 0x97, 0x80, 0x2f, 0xc8, 0x90, 0x60, 0x88, 0x21, 0x3e, 0xf9, 0xc9, 0xd9, 0x42, 0xe9,
	0x5c, 0x02, 0xc9, 0xf9, 0x79, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0xf1, 0x45, 0x20, 0x1b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x6c, 0x4e, 0xdc, 0x93, 0x67, 0xb8, 0x75, 0x4f, 0x5e,
	0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17, 0x6a, 0x1e, 0x94, 0xd2, 0x2d,
	0x4e, 0xc9, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x73, 0x49, 0x4d, 0xbe, 0xb4, 0x45, 0x97,
	0x0b, 0x6a, 0x9d, 0x4b, 0x6a, 0x72, 0x10, 0x3f, 0xc2, 0xd4, 0x20, 0x90, 0xa1, 0x42, 0x89, 0x5c,
	0xbc, 0x39, 0xf9, 0xc9, 0xd9, 0xa9, 0x29, 0xf1, 0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0x4c,
	0x24, 0xdb, 0xe2, 0x99, 0x57, 0x82, 0x64, 0x8b, 0x67, 0x5e, 0x49, 0x10, 0x0f, 0xc4, 0x48, 0x47,
	0xb0, 0x89, 0x42, 0xda, 0x5c, 0x82, 0xd0, 0x60, 0xcb, 0x2f, 0x8a, 0x4f, 0x4c, 0x49, 0x29, 0x4a,
	0x2d, 0x2e, 0x96, 0x60, 0x06, 0x59, 0x13, 0x24, 0x00, 0x97, 0x70, 0x84, 0x88, 0x83, 0x14, 0x97,
	0x25, 0xe6, 0x64, 0xa6, 0xa0, 0x28, 0x66, 0x81, 0x28, 0x86, 0x4b, 0x40, 0x15, 0x5b, 0x71, 0x74,
	0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40, 0x9e, 0xc1, 0x29, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0x2c, 0x91, 0x7c, 0x50, 0x94, 0x9a, 0x98, 0x93, 0x99, 0x5f, 0x92, 0x9a,
	0x9c, 0xa1, 0x0f, 0x8e, 0x36, 0x5d, 0x58, 0xcc, 0x56, 0xa0, 0xf1, 0xc1, 0x1e, 0x4b, 0x62, 0x03,
	0x47, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x67, 0x2c, 0x88, 0x06, 0x02, 0x00, 0x00,
}

func (m *MultiStakingLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiStakingLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiStakingLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.LockedAmount.Size()
		i -= size
		if _, err := m.LockedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ConversionRatio.Size()
		i -= size
		if _, err := m.ConversionRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDelegation(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MultiStakingLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ConversionRatio.Size()
	n += 1 + l + sovDelegation(uint64(l))
	l = m.LockedAmount.Size()
	n += 1 + l + sovDelegation(uint64(l))
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	return n
}

func sovDelegation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegation(x uint64) (n int) {
	return sovDelegation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiStakingLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegation
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
			return fmt.Errorf("proto: MultiStakingLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiStakingLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConversionRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConversionRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LockedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegation
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
func skipDelegation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
				return 0, ErrInvalidLengthDelegation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegation = fmt.Errorf("proto: unexpected end of group")
)
