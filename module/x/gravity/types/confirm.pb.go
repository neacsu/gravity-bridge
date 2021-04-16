// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gravity/v1/confirm.proto

package types

import (
	fmt "fmt"
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

// ConfirmType is the cosmos type of a confirmation of a given object from the eth signers
type ConfirmType int32

const (
	ConfirmType_CONFIRM_TYPE_UNSPECIFIED ConfirmType = 0
	ConfirmType_CONFIRM_TYPE_VALSET      ConfirmType = 1
	ConfirmType_CONFIRM_TYPE_BATCH       ConfirmType = 2
	ConfirmType_CONFIRM_TYPE_LOGIC       ConfirmType = 3
)

var ConfirmType_name = map[int32]string{
	0: "CONFIRM_TYPE_UNSPECIFIED",
	1: "CONFIRM_TYPE_VALSET",
	2: "CONFIRM_TYPE_BATCH",
	3: "CONFIRM_TYPE_LOGIC",
}

var ConfirmType_value = map[string]int32{
	"CONFIRM_TYPE_UNSPECIFIED": 0,
	"CONFIRM_TYPE_VALSET":      1,
	"CONFIRM_TYPE_BATCH":       2,
	"CONFIRM_TYPE_LOGIC":       3,
}

func (x ConfirmType) String() string {
	return proto.EnumName(ConfirmType_name, int32(x))
}

func (ConfirmType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fd753ee695c4d9b, []int{0}
}

// ConfirmLogicCall
// When validators observe a MsgRequestBatch they form a batch by ordering
// transactions currently in the txqueue in order of highest to lowest fee,
// cutting off when the batch either reaches a hardcoded maximum size (to be
// decided, probably around 100) or when transactions stop being profitable
// (TODO determine this without nondeterminism) This message includes the batch
// as well as an Ethereum signature over this batch by the validator
// -------------
type ConfirmLogicCall struct {
	InvalidationId      string `protobuf:"bytes,1,opt,name=invalidation_id,json=invalidationId,proto3" json:"invalidation_id,omitempty"`
	InvalidationNonce   uint64 `protobuf:"varint,2,opt,name=invalidation_nonce,json=invalidationNonce,proto3" json:"invalidation_nonce,omitempty"`
	EthSigner           string `protobuf:"bytes,3,opt,name=eth_signer,json=ethSigner,proto3" json:"eth_signer,omitempty"`
	OrchestratorAddress string `protobuf:"bytes,4,opt,name=orchestrator_address,json=orchestratorAddress,proto3" json:"orchestrator_address,omitempty"`
	Signature           string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfirmLogicCall) Reset()         { *m = ConfirmLogicCall{} }
func (m *ConfirmLogicCall) String() string { return proto.CompactTextString(m) }
func (*ConfirmLogicCall) ProtoMessage()    {}
func (*ConfirmLogicCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fd753ee695c4d9b, []int{0}
}
func (m *ConfirmLogicCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfirmLogicCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfirmLogicCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfirmLogicCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmLogicCall.Merge(m, src)
}
func (m *ConfirmLogicCall) XXX_Size() int {
	return m.Size()
}
func (m *ConfirmLogicCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmLogicCall.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmLogicCall proto.InternalMessageInfo

func (m *ConfirmLogicCall) GetInvalidationId() string {
	if m != nil {
		return m.InvalidationId
	}
	return ""
}

func (m *ConfirmLogicCall) GetInvalidationNonce() uint64 {
	if m != nil {
		return m.InvalidationNonce
	}
	return 0
}

func (m *ConfirmLogicCall) GetEthSigner() string {
	if m != nil {
		return m.EthSigner
	}
	return ""
}

func (m *ConfirmLogicCall) GetOrchestratorAddress() string {
	if m != nil {
		return m.OrchestratorAddress
	}
	return ""
}

func (m *ConfirmLogicCall) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

// ConfirmBatch
// When validators observe a MsgRequestBatch they form a batch by ordering
// transactions currently in the txqueue in order of highest to lowest fee,
// cutting off when the batch either reaches a hardcoded maximum size (to be
// decided, probably around 100) or when transactions stop being profitable
// (TODO determine this without nondeterminism) This message includes the batch
// as well as an Ethereum signature over this batch by the validator
// -------------
type ConfirmBatch struct {
	Nonce               uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TokenContract       string `protobuf:"bytes,2,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
	EthSigner           string `protobuf:"bytes,3,opt,name=eth_signer,json=ethSigner,proto3" json:"eth_signer,omitempty"`
	OrchestratorAddress string `protobuf:"bytes,4,opt,name=orchestrator_address,json=orchestratorAddress,proto3" json:"orchestrator_address,omitempty"`
	Signature           string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfirmBatch) Reset()         { *m = ConfirmBatch{} }
func (m *ConfirmBatch) String() string { return proto.CompactTextString(m) }
func (*ConfirmBatch) ProtoMessage()    {}
func (*ConfirmBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fd753ee695c4d9b, []int{1}
}
func (m *ConfirmBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfirmBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfirmBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfirmBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmBatch.Merge(m, src)
}
func (m *ConfirmBatch) XXX_Size() int {
	return m.Size()
}
func (m *ConfirmBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmBatch proto.InternalMessageInfo

func (m *ConfirmBatch) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *ConfirmBatch) GetTokenContract() string {
	if m != nil {
		return m.TokenContract
	}
	return ""
}

func (m *ConfirmBatch) GetEthSigner() string {
	if m != nil {
		return m.EthSigner
	}
	return ""
}

func (m *ConfirmBatch) GetOrchestratorAddress() string {
	if m != nil {
		return m.OrchestratorAddress
	}
	return ""
}

func (m *ConfirmBatch) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

// ValsetConfirm
// this is the message sent by the validators when they wish to submit their
// signatures over the validator set at a given block height. A validator must
// first call MsgSetEthAddress to set their Ethereum address to be used for
// signing. Then someone (anyone) must make a ValsetRequest, the request is
// essentially a messaging mechanism to determine which block all validators
// should submit signatures over. Finally validators sign the validator set,
// powers, and Ethereum addresses of the entire validator set at the height of a
// ValsetRequest and submit that signature with this message.
//
// If a sufficient number of validators (66% of voting power) (A) have set
// Ethereum addresses and (B) submit ValsetConfirm messages with their
// signatures it is then possible for anyone to view these signatures in the
// chain store and submit them to Ethereum to update the validator set
// -------------
type ValsetConfirm struct {
	Nonce               uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	OrchestratorAddress string `protobuf:"bytes,2,opt,name=orchestrator_address,json=orchestratorAddress,proto3" json:"orchestrator_address,omitempty"`
	EthAddress          string `protobuf:"bytes,3,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	Signature           string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ValsetConfirm) Reset()         { *m = ValsetConfirm{} }
func (m *ValsetConfirm) String() string { return proto.CompactTextString(m) }
func (*ValsetConfirm) ProtoMessage()    {}
func (*ValsetConfirm) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fd753ee695c4d9b, []int{2}
}
func (m *ValsetConfirm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValsetConfirm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValsetConfirm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValsetConfirm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValsetConfirm.Merge(m, src)
}
func (m *ValsetConfirm) XXX_Size() int {
	return m.Size()
}
func (m *ValsetConfirm) XXX_DiscardUnknown() {
	xxx_messageInfo_ValsetConfirm.DiscardUnknown(m)
}

var xxx_messageInfo_ValsetConfirm proto.InternalMessageInfo

func (m *ValsetConfirm) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *ValsetConfirm) GetOrchestratorAddress() string {
	if m != nil {
		return m.OrchestratorAddress
	}
	return ""
}

func (m *ValsetConfirm) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *ValsetConfirm) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterEnum("gravity.v1.ConfirmType", ConfirmType_name, ConfirmType_value)
	proto.RegisterType((*ConfirmLogicCall)(nil), "gravity.v1.ConfirmLogicCall")
	proto.RegisterType((*ConfirmBatch)(nil), "gravity.v1.ConfirmBatch")
	proto.RegisterType((*ValsetConfirm)(nil), "gravity.v1.ValsetConfirm")
}

func init() { proto.RegisterFile("gravity/v1/confirm.proto", fileDescriptor_9fd753ee695c4d9b) }

var fileDescriptor_9fd753ee695c4d9b = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x69, 0x8a, 0x94, 0x29, 0x2d, 0x66, 0x5b, 0x81, 0x0f, 0xc5, 0x54, 0x95, 0x10,
	0x15, 0x52, 0x63, 0x45, 0x1c, 0x38, 0x27, 0x26, 0x05, 0x4b, 0x21, 0x2d, 0x49, 0xa8, 0x04, 0x17,
	0x6b, 0xb3, 0x5e, 0xec, 0x15, 0x8e, 0x37, 0xda, 0x9d, 0x44, 0xe4, 0x2d, 0xb8, 0xf1, 0x30, 0xbc,
	0x00, 0xc7, 0x1e, 0x11, 0x27, 0x94, 0xbc, 0x08, 0xf2, 0x9f, 0xa0, 0x04, 0xc8, 0xb5, 0x47, 0x7f,
	0xbf, 0xd1, 0xcc, 0xef, 0x93, 0xbc, 0x60, 0x47, 0x9a, 0xcd, 0x24, 0xce, 0xdd, 0x59, 0xd3, 0xe5,
	0x2a, 0xfd, 0x28, 0xf5, 0xb8, 0x31, 0xd1, 0x0a, 0x15, 0x85, 0x92, 0x34, 0x66, 0xcd, 0xd3, 0x9f,
	0x04, 0x2c, 0xaf, 0xa0, 0x5d, 0x15, 0x49, 0xee, 0xb1, 0x24, 0xa1, 0x4f, 0xe1, 0x9e, 0x4c, 0x67,
	0x2c, 0x91, 0x21, 0x43, 0xa9, 0xd2, 0x40, 0x86, 0x36, 0x39, 0x21, 0x67, 0xf5, 0xfe, 0xc1, 0x7a,
	0xec, 0x87, 0xf4, 0x1c, 0xe8, 0xc6, 0x60, 0xaa, 0x52, 0x2e, 0xec, 0xea, 0x09, 0x39, 0xab, 0xf5,
	0xef, 0xaf, 0x93, 0x5e, 0x06, 0xe8, 0x23, 0x00, 0x81, 0x71, 0x60, 0x64, 0x94, 0x0a, 0x6d, 0xef,
	0xe4, 0x2b, 0xeb, 0x02, 0xe3, 0x41, 0x1e, 0xd0, 0x26, 0x1c, 0x29, 0xcd, 0x63, 0x61, 0x50, 0x33,
	0x54, 0x3a, 0x60, 0x61, 0xa8, 0x85, 0x31, 0x76, 0x2d, 0x1f, 0x3c, 0x5c, 0x67, 0xad, 0x02, 0xd1,
	0x63, 0xa8, 0x67, 0xdb, 0x18, 0x4e, 0xb5, 0xb0, 0x77, 0x8b, 0x85, 0x7f, 0x82, 0xd3, 0x6f, 0x04,
	0xee, 0x96, 0xe5, 0xda, 0x0c, 0x79, 0x4c, 0x8f, 0x60, 0xb7, 0x50, 0x24, 0xb9, 0x62, 0xf1, 0x41,
	0x9f, 0xc0, 0x01, 0xaa, 0x4f, 0x22, 0x0d, 0xb8, 0x4a, 0x51, 0x33, 0x8e, 0x79, 0x83, 0x7a, 0x7f,
	0x3f, 0x4f, 0xbd, 0x32, 0xbc, 0x75, 0xfb, 0xaf, 0x04, 0xf6, 0xaf, 0x59, 0x62, 0x04, 0x96, 0x1d,
	0xb6, 0xe8, 0x6f, 0x3b, 0x5c, 0xdd, 0x7e, 0xf8, 0x31, 0xec, 0x65, 0x55, 0x56, 0x93, 0x45, 0x97,
	0xac, 0xdd, 0x7f, 0xcd, 0x6a, 0x7f, 0x99, 0x3d, 0x43, 0xd8, 0x2b, 0x95, 0x86, 0xf3, 0x89, 0xa0,
	0xc7, 0x60, 0x7b, 0x97, 0xbd, 0x0b, 0xbf, 0xff, 0x26, 0x18, 0xbe, 0xbf, 0xea, 0x04, 0xef, 0x7a,
	0x83, 0xab, 0x8e, 0xe7, 0x5f, 0xf8, 0x9d, 0x97, 0x56, 0x85, 0x3e, 0x84, 0xc3, 0x0d, 0x7a, 0xdd,
	0xea, 0x0e, 0x3a, 0x43, 0x8b, 0xd0, 0x07, 0x40, 0x37, 0x40, 0xbb, 0x35, 0xf4, 0x5e, 0x5b, 0xd5,
	0x7f, 0xf2, 0xee, 0xe5, 0x2b, 0xdf, 0xb3, 0x76, 0xda, 0x6f, 0xbf, 0x2f, 0x1c, 0x72, 0xb3, 0x70,
	0xc8, 0xaf, 0x85, 0x43, 0xbe, 0x2c, 0x9d, 0xca, 0xcd, 0xd2, 0xa9, 0xfc, 0x58, 0x3a, 0x95, 0x0f,
	0x2f, 0x22, 0x89, 0xf1, 0x74, 0xd4, 0xe0, 0x6a, 0xec, 0x72, 0x65, 0xc6, 0xca, 0xb8, 0xe5, 0x2f,
	0x7e, 0x3e, 0xd2, 0x32, 0x8c, 0x84, 0x3b, 0x56, 0xe1, 0x34, 0x11, 0xee, 0xe7, 0x55, 0xee, 0xe2,
	0x7c, 0x22, 0xcc, 0xe8, 0x4e, 0xfe, 0x20, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x52,
	0x47, 0xa3, 0x2c, 0x03, 0x00, 0x00,
}

func (m *ConfirmLogicCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfirmLogicCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfirmLogicCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrchestratorAddress) > 0 {
		i -= len(m.OrchestratorAddress)
		copy(dAtA[i:], m.OrchestratorAddress)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.OrchestratorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EthSigner) > 0 {
		i -= len(m.EthSigner)
		copy(dAtA[i:], m.EthSigner)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.EthSigner)))
		i--
		dAtA[i] = 0x1a
	}
	if m.InvalidationNonce != 0 {
		i = encodeVarintConfirm(dAtA, i, uint64(m.InvalidationNonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.InvalidationId) > 0 {
		i -= len(m.InvalidationId)
		copy(dAtA[i:], m.InvalidationId)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.InvalidationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfirmBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfirmBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfirmBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrchestratorAddress) > 0 {
		i -= len(m.OrchestratorAddress)
		copy(dAtA[i:], m.OrchestratorAddress)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.OrchestratorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EthSigner) > 0 {
		i -= len(m.EthSigner)
		copy(dAtA[i:], m.EthSigner)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.EthSigner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenContract) > 0 {
		i -= len(m.TokenContract)
		copy(dAtA[i:], m.TokenContract)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.TokenContract)))
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintConfirm(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValsetConfirm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValsetConfirm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValsetConfirm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OrchestratorAddress) > 0 {
		i -= len(m.OrchestratorAddress)
		copy(dAtA[i:], m.OrchestratorAddress)
		i = encodeVarintConfirm(dAtA, i, uint64(len(m.OrchestratorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintConfirm(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfirm(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfirm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfirmLogicCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InvalidationId)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	if m.InvalidationNonce != 0 {
		n += 1 + sovConfirm(uint64(m.InvalidationNonce))
	}
	l = len(m.EthSigner)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.OrchestratorAddress)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	return n
}

func (m *ConfirmBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovConfirm(uint64(m.Nonce))
	}
	l = len(m.TokenContract)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.EthSigner)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.OrchestratorAddress)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	return n
}

func (m *ValsetConfirm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovConfirm(uint64(m.Nonce))
	}
	l = len(m.OrchestratorAddress)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovConfirm(uint64(l))
	}
	return n
}

func sovConfirm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfirm(x uint64) (n int) {
	return sovConfirm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfirmLogicCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfirm
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
			return fmt.Errorf("proto: ConfirmLogicCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfirmLogicCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvalidationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvalidationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvalidationNonce", wireType)
			}
			m.InvalidationNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InvalidationNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthSigner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthSigner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrchestratorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrchestratorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfirm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfirm
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
func (m *ConfirmBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfirm
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
			return fmt.Errorf("proto: ConfirmBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfirmBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthSigner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthSigner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrchestratorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrchestratorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfirm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfirm
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
func (m *ValsetConfirm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfirm
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
			return fmt.Errorf("proto: ValsetConfirm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValsetConfirm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrchestratorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrchestratorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfirm
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
				return ErrInvalidLengthConfirm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfirm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfirm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfirm
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
func skipConfirm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfirm
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
					return 0, ErrIntOverflowConfirm
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
					return 0, ErrIntOverflowConfirm
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
				return 0, ErrInvalidLengthConfirm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfirm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfirm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfirm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfirm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfirm = fmt.Errorf("proto: unexpected end of group")
)
