// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gravity/v1/attestation.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// Attestation is an aggregate of `claims` that eventually becomes `observed` by
// all orchestrators
// EVENT_NONCE:
// EventNonce a nonce provided by the gravity contract that is unique per event fired
// These event nonces must be relayed in order. This is a correctness issue,
// if relaying out of order transaction replay attacks become possible
// OBSERVED:
// Observed indicates that >67% of validators have attested to the event,
// and that the event should be executed by the gravity state machine
//
// The actual content of the claims is passed in with the transaction making the claim
// and then passed through the call stack alongside the attestation while it is processed
// the key in which the attestation is stored is keyed on the exact details of the claim
// but there is no reason to store those exact details becuause the next message sender
// will kindly provide you with them.
type Attestation struct {
	Observed bool       `protobuf:"varint,1,opt,name=observed,proto3" json:"observed,omitempty"`
	Votes    []string   `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	Height   uint64     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Claim    *types.Any `protobuf:"bytes,4,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3205613bbab7525, []int{0}
}
func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return m.Size()
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetObserved() bool {
	if m != nil {
		return m.Observed
	}
	return false
}

func (m *Attestation) GetVotes() []string {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *Attestation) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Attestation) GetClaim() *types.Any {
	if m != nil {
		return m.Claim
	}
	return nil
}

// ERC20Token unique identifier for an Ethereum ERC20 token.
// CONTRACT:
// The contract address on ETH of the token, this could be a Cosmos
// originated token, if so it will be the ERC20 address of the representation
// (note: developers should look up the token symbol using the address on ETH to display for UI)
type ERC20Token struct {
	Contract string                                 `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Amount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *ERC20Token) Reset()         { *m = ERC20Token{} }
func (m *ERC20Token) String() string { return proto.CompactTextString(m) }
func (*ERC20Token) ProtoMessage()    {}
func (*ERC20Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3205613bbab7525, []int{1}
}
func (m *ERC20Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20Token.Merge(m, src)
}
func (m *ERC20Token) XXX_Size() int {
	return m.Size()
}
func (m *ERC20Token) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20Token.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20Token proto.InternalMessageInfo

func (m *ERC20Token) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*Attestation)(nil), "gravity.v1.Attestation")
	proto.RegisterType((*ERC20Token)(nil), "gravity.v1.ERC20Token")
}

func init() { proto.RegisterFile("gravity/v1/attestation.proto", fileDescriptor_e3205613bbab7525) }

var fileDescriptor_e3205613bbab7525 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x6e, 0xe2, 0x40,
	0x10, 0xc6, 0xbd, 0xfc, 0x13, 0x2c, 0xba, 0xe2, 0x2c, 0x74, 0xf2, 0xa1, 0x93, 0xb1, 0x28, 0x4e,
	0x6e, 0xd8, 0x3d, 0xb8, 0xe2, 0xaa, 0x2b, 0x00, 0x11, 0x29, 0x65, 0xac, 0x54, 0x69, 0x22, 0xff,
	0xd9, 0xac, 0x2d, 0xb0, 0x07, 0xd9, 0x63, 0x2b, 0x7e, 0x8b, 0x54, 0x79, 0x92, 0x3c, 0x04, 0x4a,
	0x45, 0x19, 0xa5, 0x40, 0x11, 0xbc, 0x48, 0x84, 0x6d, 0x48, 0x91, 0x22, 0xd5, 0xee, 0x6f, 0xbe,
	0xf9, 0x46, 0xdf, 0xee, 0xd0, 0x5f, 0x32, 0xb6, 0xb3, 0x00, 0x73, 0x9e, 0x8d, 0xb9, 0x8d, 0x28,
	0x12, 0xb4, 0x31, 0x80, 0x88, 0xad, 0x63, 0x40, 0x50, 0x69, 0xa5, 0xb2, 0x6c, 0xdc, 0xef, 0x49,
	0x90, 0x50, 0x94, 0xf9, 0xf1, 0x56, 0x76, 0xf4, 0x7f, 0x4a, 0x00, 0xb9, 0x12, 0xbc, 0x20, 0x27,
	0xbd, 0xe3, 0x76, 0x94, 0x9f, 0x24, 0x17, 0x92, 0x10, 0x92, 0xdb, 0xd2, 0x53, 0xc2, 0x97, 0xae,
	0xe1, 0x23, 0xa1, 0xdd, 0xe9, 0x47, 0x10, 0xb5, 0x4f, 0xdb, 0xe0, 0x24, 0x22, 0xce, 0x84, 0xa7,
	0x11, 0x83, 0x98, 0x6d, 0xeb, 0xcc, 0x6a, 0x8f, 0x36, 0x33, 0x40, 0x91, 0x68, 0x35, 0xa3, 0x6e,
	0x76, 0xac, 0x12, 0xd4, 0x1f, 0xb4, 0xe5, 0x8b, 0x40, 0xfa, 0xa8, 0xd5, 0x0d, 0x62, 0x36, 0xac,
	0x8a, 0xd4, 0xff, 0xb4, 0xe9, 0xae, 0xec, 0x20, 0xd4, 0x1a, 0x06, 0x31, 0xbb, 0x93, 0x1e, 0x2b,
	0x43, 0xb0, 0x53, 0x08, 0x36, 0x8d, 0xf2, 0xd9, 0xf7, 0xe7, 0xa7, 0xd1, 0xb7, 0x05, 0xfa, 0x22,
	0x16, 0x69, 0x38, 0x3f, 0xb6, 0x5b, 0xa5, 0x6b, 0xb8, 0xa6, 0x74, 0x61, 0xcd, 0x27, 0x7f, 0xae,
	0x61, 0x29, 0x8a, 0x58, 0x2e, 0x44, 0x18, 0xdb, 0x2e, 0x16, 0xb1, 0x3a, 0xd6, 0x99, 0xd5, 0x0b,
	0xda, 0xb2, 0x43, 0x48, 0x23, 0xd4, 0x6a, 0x47, 0x65, 0xc6, 0x36, 0xbb, 0x81, 0xf2, 0xba, 0x1b,
	0xfc, 0x96, 0x01, 0xfa, 0xa9, 0xc3, 0x5c, 0x08, 0xab, 0xef, 0xa8, 0x8e, 0x51, 0xe2, 0x2d, 0x39,
	0xe6, 0x6b, 0x91, 0xb0, 0xcb, 0x08, 0xad, 0xca, 0x3d, 0xbb, 0xda, 0xec, 0x75, 0xb2, 0xdd, 0xeb,
	0xe4, 0x6d, 0xaf, 0x93, 0x87, 0x83, 0xae, 0x6c, 0x0f, 0xba, 0xf2, 0x72, 0xd0, 0x95, 0x9b, 0x7f,
	0x9f, 0x27, 0x55, 0x9b, 0x1a, 0x39, 0x71, 0xe0, 0x49, 0xc1, 0x43, 0xf0, 0xd2, 0x95, 0xe0, 0xf7,
	0xa7, 0x7a, 0x39, 0xde, 0x69, 0x15, 0x8f, 0xfd, 0xfb, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x77, 0x87,
	0x7f, 0x3f, 0xf7, 0x01, 0x00, 0x00,
}

func (m *Attestation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attestation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attestation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttestation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintAttestation(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Votes[iNdEx])
			copy(dAtA[i:], m.Votes[iNdEx])
			i = encodeVarintAttestation(dAtA, i, uint64(len(m.Votes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Observed {
		i--
		if m.Observed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ERC20Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAttestation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Attestation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Observed {
		n += 2
	}
	if len(m.Votes) > 0 {
		for _, s := range m.Votes {
			l = len(s)
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovAttestation(uint64(m.Height))
	}
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *ERC20Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAttestation(uint64(l))
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attestation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: Attestation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attestation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Observed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
			m.Observed = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Any{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func (m *ERC20Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: ERC20Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func skipAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
				return 0, ErrInvalidLengthAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttestation = fmt.Errorf("proto: unexpected end of group")
)
