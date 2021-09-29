// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/genesis_account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type GenesisAccount struct {
	ChainID uint64                                   `protobuf:"varint,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Address string                                   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Coins   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=Coins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"Coins"`
}

func (m *GenesisAccount) Reset()         { *m = GenesisAccount{} }
func (m *GenesisAccount) String() string { return proto.CompactTextString(m) }
func (*GenesisAccount) ProtoMessage()    {}
func (*GenesisAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_44cd378858027518, []int{0}
}
func (m *GenesisAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccount.Merge(m, src)
}
func (m *GenesisAccount) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccount proto.InternalMessageInfo

func (m *GenesisAccount) GetChainID() uint64 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *GenesisAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GenesisAccount) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisAccount)(nil), "tendermint.spn.launch.GenesisAccount")
}

func init() { proto.RegisterFile("launch/genesis_account.proto", fileDescriptor_44cd378858027518) }

var fileDescriptor_44cd378858027518 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4a, 0xc5, 0x30,
	0x14, 0x87, 0x1b, 0xaf, 0x7f, 0xb0, 0x82, 0x43, 0x51, 0xa8, 0x17, 0x49, 0x8b, 0x8b, 0x5d, 0x4c,
	0xb8, 0xfa, 0x04, 0x56, 0x41, 0x5c, 0x3b, 0xea, 0x20, 0x69, 0x1a, 0xda, 0xa0, 0x4d, 0x4a, 0x4f,
	0x2a, 0xfa, 0x0e, 0x0e, 0x3e, 0x87, 0x4f, 0x72, 0xc7, 0x0b, 0x2e, 0x4e, 0x57, 0x69, 0xdf, 0xc2,
	0x49, 0xda, 0xf4, 0xa2, 0xa3, 0x53, 0x72, 0x38, 0xdf, 0x39, 0xdf, 0xe1, 0xe7, 0x1e, 0x3e, 0xb0,
	0x46, 0xf1, 0x82, 0xe6, 0x42, 0x09, 0x90, 0x70, 0xc7, 0x38, 0xd7, 0x8d, 0x32, 0xa4, 0xaa, 0xb5,
	0xd1, 0xde, 0xbe, 0x11, 0x2a, 0x13, 0x75, 0x29, 0x95, 0x21, 0x50, 0x29, 0x62, 0xe1, 0xe9, 0x5e,
	0xae, 0x73, 0x3d, 0x10, 0xb4, 0xff, 0x59, 0x78, 0x8a, 0xb9, 0x86, 0x52, 0x03, 0x4d, 0x19, 0x08,
	0xfa, 0x38, 0x4b, 0x85, 0x61, 0x33, 0xca, 0xb5, 0x54, 0xb6, 0x7f, 0xf4, 0x8e, 0xdc, 0xdd, 0x2b,
	0xab, 0x39, 0xb7, 0x16, 0xcf, 0x77, 0xb7, 0x78, 0xc1, 0xa4, 0xba, 0xbe, 0xf4, 0x51, 0x88, 0xa2,
	0xf5, 0x64, 0x55, 0xf6, 0x1d, 0x96, 0x65, 0xb5, 0x00, 0xf0, 0xd7, 0x42, 0x14, 0x6d, 0x27, 0xab,
	0xd2, 0x7b, 0x41, 0xee, 0xc6, 0x85, 0x96, 0x0a, 0xfc, 0x49, 0x38, 0x89, 0x76, 0x4e, 0x0f, 0x88,
	0xf5, 0x92, 0xde, 0x4b, 0x46, 0x2f, 0xe9, 0x89, 0xf8, 0x76, 0xbe, 0x0c, 0x9c, 0xef, 0x65, 0x70,
	0x9c, 0x4b, 0x53, 0x34, 0x29, 0xe1, 0xba, 0xa4, 0xe3, 0x91, 0xf6, 0x39, 0x81, 0xec, 0x9e, 0x9a,
	0xe7, 0x4a, 0xc0, 0x30, 0xf0, 0xf6, 0x19, 0x44, 0xff, 0x44, 0x21, 0xb1, 0x47, 0xc4, 0xf1, 0xbc,
	0xc5, 0x68, 0xd1, 0x62, 0xf4, 0xd5, 0x62, 0xf4, 0xda, 0x61, 0x67, 0xd1, 0x61, 0xe7, 0xa3, 0xc3,
	0xce, 0xcd, 0xdf, 0x55, 0xbf, 0x39, 0x52, 0xa8, 0x14, 0x7d, 0xa2, 0x63, 0xec, 0xc3, 0xc2, 0x74,
	0x73, 0x08, 0xe8, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff, 0x91, 0x4e, 0x38, 0xe2, 0x8d, 0x01, 0x00,
	0x00,
}

func (m *GenesisAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesisAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesisAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainID != 0 {
		i = encodeVarintGenesisAccount(dAtA, i, uint64(m.ChainID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesisAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesisAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainID != 0 {
		n += 1 + sovGenesisAccount(uint64(m.ChainID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesisAccount(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGenesisAccount(uint64(l))
		}
	}
	return n
}

func sovGenesisAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesisAccount(x uint64) (n int) {
	return sovGenesisAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisAccount
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
			return fmt.Errorf("proto: GenesisAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			m.ChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisAccount
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
				return ErrInvalidLengthGenesisAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisAccount
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
				return ErrInvalidLengthGenesisAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisAccount
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
func skipGenesisAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesisAccount
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
					return 0, ErrIntOverflowGenesisAccount
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
					return 0, ErrIntOverflowGenesisAccount
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
				return 0, ErrInvalidLengthGenesisAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesisAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesisAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesisAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesisAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesisAccount = fmt.Errorf("proto: unexpected end of group")
)