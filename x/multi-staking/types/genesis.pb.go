// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multistaking/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GenesisState struct {
	MultiStakingLocks         []MultiStakingLock          `protobuf:"bytes,1,rep,name=multi_staking_locks,json=multiStakingLocks,proto3" json:"multi_staking_locks"`
	ValidatorMultiStakingCoin []ValidatorMultiStakingCoin `protobuf:"bytes,2,rep,name=validator_multi_staking_coin,json=validatorMultiStakingCoin,proto3" json:"validator_multi_staking_coin"`
	StakingGenesisState       *types.GenesisState         `protobuf:"bytes,3,opt,name=staking_genesis_state,json=stakingGenesisState,proto3" json:"staking_genesis_state,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f95a201ebed173c, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetMultiStakingLocks() []MultiStakingLock {
	if m != nil {
		return m.MultiStakingLocks
	}
	return nil
}

func (m *GenesisState) GetValidatorMultiStakingCoin() []ValidatorMultiStakingCoin {
	if m != nil {
		return m.ValidatorMultiStakingCoin
	}
	return nil
}

func (m *GenesisState) GetStakingGenesisState() *types.GenesisState {
	if m != nil {
		return m.StakingGenesisState
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "multistaking.v1.GenesisState")
}

func init() { proto.RegisterFile("multistaking/v1/genesis.proto", fileDescriptor_8f95a201ebed173c) }

var fileDescriptor_8f95a201ebed173c = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x5b, 0xb8, 0xb9, 0x8b, 0x72, 0x93, 0x9b, 0x5b, 0xae, 0x11, 0x88, 0x16, 0x54, 0x16,
	0xc4, 0xc4, 0x4e, 0x8a, 0x2b, 0xb7, 0xb8, 0x70, 0xa3, 0x1b, 0x48, 0xd4, 0xb8, 0x69, 0xa6, 0x75,
	0x1c, 0x26, 0x74, 0x7a, 0x90, 0x19, 0x1a, 0x79, 0x0b, 0x1f, 0xc6, 0x87, 0x60, 0xc9, 0xd2, 0x95,
	0x31, 0xf0, 0x22, 0xa6, 0x9d, 0x69, 0x28, 0x35, 0xee, 0xe6, 0x3f, 0xdf, 0xdf, 0xf3, 0xf7, 0x9c,
	0x63, 0x1d, 0xf2, 0x79, 0x24, 0x99, 0x90, 0x78, 0xc2, 0x62, 0x8a, 0x12, 0x0f, 0x51, 0x12, 0x13,
	0xc1, 0x84, 0x3b, 0x9d, 0x81, 0x04, 0xfb, 0x6f, 0x11, 0xbb, 0x89, 0xd7, 0x6a, 0x52, 0x00, 0x1a,
	0x11, 0x94, 0xe1, 0x60, 0xfe, 0x84, 0x70, 0xbc, 0x50, 0xde, 0x56, 0xbb, 0x8c, 0x24, 0xe3, 0x44,
	0x48, 0xcc, 0xa7, 0xda, 0xf0, 0x9f, 0x02, 0x85, 0xec, 0x89, 0xd2, 0x97, 0xae, 0x36, 0x43, 0x10,
	0x1c, 0x84, 0xaf, 0x80, 0x12, 0x1a, 0x39, 0x4a, 0xa1, 0x00, 0x0b, 0x82, 0x12, 0x2f, 0x20, 0x12,
	0x7b, 0x28, 0x04, 0x16, 0x6b, 0xde, 0xd5, 0x7c, 0xfb, 0xfb, 0xca, 0xb2, 0x33, 0x43, 0xeb, 0xa4,
	0x3c, 0x62, 0xa6, 0xfd, 0x7c, 0x28, 0x65, 0xda, 0xd7, 0xad, 0xb8, 0x50, 0x16, 0xa1, 0xc1, 0xf1,
	0x5b, 0xc5, 0xfa, 0x73, 0xa5, 0xfa, 0x8d, 0x24, 0x96, 0xc4, 0xbe, 0xb3, 0xea, 0x3b, 0x0d, 0xfc,
	0x08, 0xc2, 0x89, 0x68, 0x98, 0x9d, 0x6a, 0xaf, 0xd6, 0x3f, 0x72, 0x4b, 0x0b, 0x73, 0x6f, 0x52,
	0x3d, 0x52, 0xfa, 0x1a, 0xc2, 0xc9, 0xe0, 0xd7, 0xf2, 0xa3, 0x6d, 0x0c, 0xff, 0xf1, 0x52, 0x5d,
	0xd8, 0xcf, 0xd6, 0x41, 0x82, 0x23, 0xf6, 0x88, 0x25, 0xcc, 0xfc, 0xdd, 0x88, 0x74, 0xe6, 0x46,
	0x25, 0x4b, 0x38, 0xfd, 0x96, 0x70, 0x9b, 0x7f, 0x54, 0x8c, 0xba, 0x04, 0x16, 0xeb, 0xa8, 0x66,
	0xf2, 0x93, 0xc1, 0xbe, 0xb7, 0xf6, 0xf2, 0x08, 0xbd, 0xb3, 0x34, 0x52, 0x92, 0x46, 0xb5, 0x63,
	0xf6, 0x6a, 0xfd, 0xae, 0xab, 0xcf, 0xb1, 0x4d, 0xcb, 0x16, 0xec, 0x16, 0x17, 0x32, 0xac, 0x6b,
	0x5a, 0x2c, 0x0e, 0x46, 0xcb, 0xb5, 0x63, 0xae, 0xd6, 0x8e, 0xf9, 0xb9, 0x76, 0xcc, 0xd7, 0x8d,
	0x63, 0xac, 0x36, 0x8e, 0xf1, 0xbe, 0x71, 0x8c, 0x87, 0x0b, 0xca, 0xe4, 0x78, 0x1e, 0xb8, 0x21,
	0x70, 0x34, 0x23, 0x38, 0x62, 0x20, 0x49, 0x38, 0x56, 0x47, 0x39, 0xcb, 0xaf, 0xf4, 0x52, 0xd2,
	0x72, 0x31, 0x25, 0x22, 0xf8, 0x9d, 0x9d, 0xe4, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x3e,
	0x77, 0xb6, 0xb5, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StakingGenesisState != nil {
		{
			size, err := m.StakingGenesisState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorMultiStakingCoin) > 0 {
		for iNdEx := len(m.ValidatorMultiStakingCoin) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorMultiStakingCoin[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MultiStakingLocks) > 0 {
		for iNdEx := len(m.MultiStakingLocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MultiStakingLocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MultiStakingLocks) > 0 {
		for _, e := range m.MultiStakingLocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorMultiStakingCoin) > 0 {
		for _, e := range m.ValidatorMultiStakingCoin {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.StakingGenesisState != nil {
		l = m.StakingGenesisState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiStakingLocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultiStakingLocks = append(m.MultiStakingLocks, MultiStakingLock{})
			if err := m.MultiStakingLocks[len(m.MultiStakingLocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorMultiStakingCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorMultiStakingCoin = append(m.ValidatorMultiStakingCoin, ValidatorMultiStakingCoin{})
			if err := m.ValidatorMultiStakingCoin[len(m.ValidatorMultiStakingCoin)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingGenesisState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StakingGenesisState == nil {
				m.StakingGenesisState = &types.GenesisState{}
			}
			if err := m.StakingGenesisState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
