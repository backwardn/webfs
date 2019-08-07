// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ref.proto

package webref // import "github.com/brendoncarroll/webfs/pkg/webref"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EncAlgo int32

const (
	EncAlgo_UNKNOWN   EncAlgo = 0
	EncAlgo_AES256CTR EncAlgo = 1
	EncAlgo_CHACHA20  EncAlgo = 2
)

var EncAlgo_name = map[int32]string{
	0: "UNKNOWN",
	1: "AES256CTR",
	2: "CHACHA20",
}
var EncAlgo_value = map[string]int32{
	"UNKNOWN":   0,
	"AES256CTR": 1,
	"CHACHA20":  2,
}

func (x EncAlgo) String() string {
	return proto.EnumName(EncAlgo_name, int32(x))
}
func (EncAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{0}
}

type CryptoRef struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	EncAlgo              EncAlgo  `protobuf:"varint,2,opt,name=enc_algo,json=encAlgo,proto3,enum=webref.EncAlgo" json:"enc_algo,omitempty"`
	Dek                  []byte   `protobuf:"bytes,3,opt,name=dek,proto3" json:"dek,omitempty"`
	Length               int32    `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoRef) Reset()         { *m = CryptoRef{} }
func (m *CryptoRef) String() string { return proto.CompactTextString(m) }
func (*CryptoRef) ProtoMessage()    {}
func (*CryptoRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{0}
}
func (m *CryptoRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoRef.Unmarshal(m, b)
}
func (m *CryptoRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoRef.Marshal(b, m, deterministic)
}
func (dst *CryptoRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoRef.Merge(dst, src)
}
func (m *CryptoRef) XXX_Size() int {
	return xxx_messageInfo_CryptoRef.Size(m)
}
func (m *CryptoRef) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoRef.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoRef proto.InternalMessageInfo

func (m *CryptoRef) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CryptoRef) GetEncAlgo() EncAlgo {
	if m != nil {
		return m.EncAlgo
	}
	return EncAlgo_UNKNOWN
}

func (m *CryptoRef) GetDek() []byte {
	if m != nil {
		return m.Dek
	}
	return nil
}

func (m *CryptoRef) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Slice struct {
	Ref                  *CryptoRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Offset               int32      `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               int32      `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Slice) Reset()         { *m = Slice{} }
func (m *Slice) String() string { return proto.CompactTextString(m) }
func (*Slice) ProtoMessage()    {}
func (*Slice) Descriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{1}
}
func (m *Slice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slice.Unmarshal(m, b)
}
func (m *Slice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slice.Marshal(b, m, deterministic)
}
func (dst *Slice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slice.Merge(dst, src)
}
func (m *Slice) XXX_Size() int {
	return xxx_messageInfo_Slice.Size(m)
}
func (m *Slice) XXX_DiscardUnknown() {
	xxx_messageInfo_Slice.DiscardUnknown(m)
}

var xxx_messageInfo_Slice proto.InternalMessageInfo

func (m *Slice) GetRef() *CryptoRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *Slice) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Slice) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Mirror struct {
	Replicas             []*RepPackRef `protobuf:"bytes,1,rep,name=replicas,proto3" json:"replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Mirror) Reset()         { *m = Mirror{} }
func (m *Mirror) String() string { return proto.CompactTextString(m) }
func (*Mirror) ProtoMessage()    {}
func (*Mirror) Descriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{2}
}
func (m *Mirror) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mirror.Unmarshal(m, b)
}
func (m *Mirror) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mirror.Marshal(b, m, deterministic)
}
func (dst *Mirror) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mirror.Merge(dst, src)
}
func (m *Mirror) XXX_Size() int {
	return xxx_messageInfo_Mirror.Size(m)
}
func (m *Mirror) XXX_DiscardUnknown() {
	xxx_messageInfo_Mirror.DiscardUnknown(m)
}

var xxx_messageInfo_Mirror proto.InternalMessageInfo

func (m *Mirror) GetReplicas() []*RepPackRef {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type RepPackRef struct {
	// Types that are valid to be assigned to Ref:
	//	*RepPackRef_Single
	//	*RepPackRef_Mirror
	//	*RepPackRef_Slice
	Ref                  isRepPackRef_Ref `protobuf_oneof:"ref"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RepPackRef) Reset()         { *m = RepPackRef{} }
func (m *RepPackRef) String() string { return proto.CompactTextString(m) }
func (*RepPackRef) ProtoMessage()    {}
func (*RepPackRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{3}
}
func (m *RepPackRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepPackRef.Unmarshal(m, b)
}
func (m *RepPackRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepPackRef.Marshal(b, m, deterministic)
}
func (dst *RepPackRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepPackRef.Merge(dst, src)
}
func (m *RepPackRef) XXX_Size() int {
	return xxx_messageInfo_RepPackRef.Size(m)
}
func (m *RepPackRef) XXX_DiscardUnknown() {
	xxx_messageInfo_RepPackRef.DiscardUnknown(m)
}

var xxx_messageInfo_RepPackRef proto.InternalMessageInfo

type isRepPackRef_Ref interface {
	isRepPackRef_Ref()
}

type RepPackRef_Single struct {
	Single *CryptoRef `protobuf:"bytes,1,opt,name=single,proto3,oneof"`
}

type RepPackRef_Mirror struct {
	Mirror *Mirror `protobuf:"bytes,2,opt,name=mirror,proto3,oneof"`
}

type RepPackRef_Slice struct {
	Slice *Slice `protobuf:"bytes,3,opt,name=slice,proto3,oneof"`
}

func (*RepPackRef_Single) isRepPackRef_Ref() {}

func (*RepPackRef_Mirror) isRepPackRef_Ref() {}

func (*RepPackRef_Slice) isRepPackRef_Ref() {}

func (m *RepPackRef) GetRef() isRepPackRef_Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *RepPackRef) GetSingle() *CryptoRef {
	if x, ok := m.GetRef().(*RepPackRef_Single); ok {
		return x.Single
	}
	return nil
}

func (m *RepPackRef) GetMirror() *Mirror {
	if x, ok := m.GetRef().(*RepPackRef_Mirror); ok {
		return x.Mirror
	}
	return nil
}

func (m *RepPackRef) GetSlice() *Slice {
	if x, ok := m.GetRef().(*RepPackRef_Slice); ok {
		return x.Slice
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RepPackRef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RepPackRef_OneofMarshaler, _RepPackRef_OneofUnmarshaler, _RepPackRef_OneofSizer, []interface{}{
		(*RepPackRef_Single)(nil),
		(*RepPackRef_Mirror)(nil),
		(*RepPackRef_Slice)(nil),
	}
}

func _RepPackRef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RepPackRef)
	// ref
	switch x := m.Ref.(type) {
	case *RepPackRef_Single:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Single); err != nil {
			return err
		}
	case *RepPackRef_Mirror:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mirror); err != nil {
			return err
		}
	case *RepPackRef_Slice:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Slice); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RepPackRef.Ref has unexpected type %T", x)
	}
	return nil
}

func _RepPackRef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RepPackRef)
	switch tag {
	case 1: // ref.single
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoRef)
		err := b.DecodeMessage(msg)
		m.Ref = &RepPackRef_Single{msg}
		return true, err
	case 2: // ref.mirror
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mirror)
		err := b.DecodeMessage(msg)
		m.Ref = &RepPackRef_Mirror{msg}
		return true, err
	case 3: // ref.slice
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Slice)
		err := b.DecodeMessage(msg)
		m.Ref = &RepPackRef_Slice{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RepPackRef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RepPackRef)
	// ref
	switch x := m.Ref.(type) {
	case *RepPackRef_Single:
		s := proto.Size(x.Single)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RepPackRef_Mirror:
		s := proto.Size(x.Mirror)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RepPackRef_Slice:
		s := proto.Size(x.Slice)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Ref struct {
	Ref                  *RepPackRef       `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Attrs                map[string]string `protobuf:"bytes,2,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}
func (*Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{4}
}
func (m *Ref) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ref.Unmarshal(m, b)
}
func (m *Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ref.Marshal(b, m, deterministic)
}
func (dst *Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ref.Merge(dst, src)
}
func (m *Ref) XXX_Size() int {
	return xxx_messageInfo_Ref.Size(m)
}
func (m *Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_Ref proto.InternalMessageInfo

func (m *Ref) GetRef() *RepPackRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *Ref) GetAttrs() map[string]string {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Options struct {
	EncAlgo              EncAlgo           `protobuf:"varint,1,opt,name=enc_algo,json=encAlgo,proto3,enum=webref.EncAlgo" json:"enc_algo,omitempty"`
	SecretSeed           []byte            `protobuf:"bytes,2,opt,name=secret_seed,json=secretSeed,proto3" json:"secret_seed,omitempty"`
	ObfuscateLength      bool              `protobuf:"varint,3,opt,name=obfuscate_length,json=obfuscateLength,proto3" json:"obfuscate_length,omitempty"`
	Replicas             map[string]int32  `protobuf:"bytes,4,rep,name=replicas,proto3" json:"replicas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Attrs                map[string]string `protobuf:"bytes,5,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_ref_76af531ce5e4e576, []int{5}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (dst *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(dst, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetEncAlgo() EncAlgo {
	if m != nil {
		return m.EncAlgo
	}
	return EncAlgo_UNKNOWN
}

func (m *Options) GetSecretSeed() []byte {
	if m != nil {
		return m.SecretSeed
	}
	return nil
}

func (m *Options) GetObfuscateLength() bool {
	if m != nil {
		return m.ObfuscateLength
	}
	return false
}

func (m *Options) GetReplicas() map[string]int32 {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *Options) GetAttrs() map[string]string {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoRef)(nil), "webref.CryptoRef")
	proto.RegisterType((*Slice)(nil), "webref.Slice")
	proto.RegisterType((*Mirror)(nil), "webref.Mirror")
	proto.RegisterType((*RepPackRef)(nil), "webref.RepPackRef")
	proto.RegisterType((*Ref)(nil), "webref.Ref")
	proto.RegisterMapType((map[string]string)(nil), "webref.Ref.AttrsEntry")
	proto.RegisterType((*Options)(nil), "webref.Options")
	proto.RegisterMapType((map[string]string)(nil), "webref.Options.AttrsEntry")
	proto.RegisterMapType((map[string]int32)(nil), "webref.Options.ReplicasEntry")
	proto.RegisterEnum("webref.EncAlgo", EncAlgo_name, EncAlgo_value)
}

func init() { proto.RegisterFile("ref.proto", fileDescriptor_ref_76af531ce5e4e576) }

var fileDescriptor_ref_76af531ce5e4e576 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x6d, 0x8b, 0xd3, 0x5c,
	0x10, 0x6d, 0x9a, 0x4d, 0xda, 0x4e, 0xda, 0xdd, 0x3e, 0x97, 0x87, 0x25, 0x2c, 0x88, 0xa5, 0x2a,
	0xc4, 0xba, 0xa4, 0x4b, 0x16, 0xa5, 0xea, 0xa7, 0x6e, 0x29, 0x14, 0xd4, 0xae, 0xdc, 0x2a, 0x82,
	0x08, 0x25, 0x49, 0x27, 0xd9, 0xd0, 0x6c, 0x6e, 0xb8, 0x49, 0x5d, 0xfa, 0x33, 0xc4, 0xdf, 0xe8,
	0xff, 0x90, 0x7b, 0xd3, 0xf4, 0xc5, 0x17, 0x16, 0xfc, 0x76, 0x67, 0xe6, 0xcc, 0xcc, 0x39, 0x67,
	0x42, 0xa0, 0xc1, 0x31, 0xb0, 0x53, 0xce, 0x72, 0x46, 0xf4, 0x3b, 0xf4, 0x38, 0x06, 0xdd, 0x0c,
	0x1a, 0x23, 0xbe, 0x4e, 0x73, 0x46, 0x31, 0x20, 0x6d, 0x50, 0x57, 0x3c, 0x36, 0x95, 0x8e, 0x62,
	0x35, 0xa8, 0x78, 0x92, 0x1e, 0xd4, 0x31, 0xf1, 0xe7, 0x6e, 0x1c, 0x32, 0xb3, 0xda, 0x51, 0xac,
	0x63, 0xe7, 0xc4, 0x2e, 0x3a, 0xed, 0x71, 0xe2, 0x0f, 0xe3, 0x90, 0xd1, 0x1a, 0x16, 0x0f, 0xd1,
	0xbd, 0xc0, 0xa5, 0xa9, 0x76, 0x14, 0xab, 0x49, 0xc5, 0x93, 0x9c, 0x82, 0x1e, 0x63, 0x12, 0xe6,
	0x37, 0xe6, 0x51, 0x47, 0xb1, 0x34, 0xba, 0x89, 0xba, 0x5f, 0x40, 0x9b, 0xc5, 0x91, 0x8f, 0xe4,
	0x11, 0xa8, 0x1c, 0x03, 0xb9, 0xd0, 0x70, 0xfe, 0x2b, 0x27, 0x6f, 0x09, 0x51, 0x51, 0x15, 0x53,
	0x58, 0x10, 0x64, 0x98, 0x4b, 0x06, 0x1a, 0xdd, 0x44, 0x7b, 0xd3, 0xd5, 0x83, 0xe9, 0x03, 0xd0,
	0xdf, 0x45, 0x9c, 0x33, 0x4e, 0x6c, 0xa8, 0x73, 0x4c, 0xe3, 0xc8, 0x77, 0x33, 0x53, 0xe9, 0xa8,
	0x96, 0xe1, 0x90, 0x72, 0x07, 0xc5, 0xf4, 0xbd, 0xeb, 0x2f, 0xc5, 0x92, 0x2d, 0xa6, 0xfb, 0x4d,
	0x01, 0xd8, 0x15, 0xc8, 0x33, 0xd0, 0xb3, 0x28, 0x09, 0x63, 0xfc, 0x2b, 0xc1, 0x49, 0x85, 0x6e,
	0x20, 0xc4, 0x02, 0xfd, 0x56, 0x6e, 0x95, 0x2c, 0x0d, 0xe7, 0xb8, 0x04, 0x17, 0x5c, 0x04, 0xb2,
	0xa8, 0x93, 0x27, 0xa0, 0x65, 0x42, 0xbd, 0xa4, 0x6d, 0x38, 0xad, 0x12, 0x28, 0x2d, 0x99, 0x54,
	0x68, 0x51, 0xbd, 0xd2, 0xa4, 0x37, 0xdd, 0xef, 0x0a, 0xa8, 0x82, 0xcc, 0xe3, 0x7d, 0xab, 0xfe,
	0x24, 0x43, 0x7a, 0x75, 0x0e, 0x9a, 0x9b, 0xe7, 0x3c, 0x33, 0xab, 0x52, 0xee, 0xe9, 0x0e, 0x17,
	0xd8, 0x43, 0x51, 0x18, 0x27, 0x39, 0x5f, 0xd3, 0x02, 0x74, 0x36, 0x00, 0xd8, 0x25, 0xc5, 0xfd,
	0x96, 0xb8, 0x2e, 0xaf, 0xbf, 0xc4, 0x35, 0xf9, 0x1f, 0xb4, 0xaf, 0x6e, 0xbc, 0x42, 0x29, 0xa9,
	0x41, 0x8b, 0xe0, 0x55, 0x75, 0xa0, 0x74, 0x7f, 0x54, 0xa1, 0x76, 0x9d, 0xe6, 0x11, 0x4b, 0xb2,
	0x83, 0x6f, 0x44, 0xb9, 0xe7, 0x1b, 0x79, 0x08, 0x46, 0x86, 0x3e, 0xc7, 0x7c, 0x9e, 0x21, 0x2e,
	0xe4, 0xdc, 0x26, 0x85, 0x22, 0x35, 0x43, 0x5c, 0x90, 0xa7, 0xd0, 0x66, 0x5e, 0xb0, 0xca, 0x7c,
	0x37, 0xc7, 0xf9, 0xde, 0x79, 0xeb, 0xf4, 0x64, 0x9b, 0x7f, 0x2b, 0xd3, 0xe4, 0xe5, 0xde, 0x75,
	0x8f, 0xa4, 0xdc, 0x07, 0xe5, 0xde, 0x0d, 0x35, 0x61, 0x8f, 0xac, 0x17, 0xaa, 0xb7, 0x70, 0x72,
	0x51, 0xda, 0xa4, 0xc9, 0xbe, 0xb3, 0x5f, 0xfb, 0x7e, 0xb7, 0xea, 0x35, 0xb4, 0x0e, 0x86, 0xdd,
	0xe7, 0x96, 0xb6, 0xe7, 0xd6, 0xbf, 0xfb, 0xdc, 0xbb, 0x84, 0xda, 0xc6, 0x43, 0x62, 0x40, 0xed,
	0xe3, 0xf4, 0xcd, 0xf4, 0xfa, 0xd3, 0xb4, 0x5d, 0x21, 0x2d, 0x68, 0x0c, 0xc7, 0x33, 0xe7, 0xf9,
	0x8b, 0xd1, 0x07, 0xda, 0x56, 0x48, 0x13, 0xea, 0xa3, 0xc9, 0x70, 0x34, 0x19, 0x3a, 0x17, 0xed,
	0xea, 0xd5, 0xf9, 0xe7, 0x5e, 0x18, 0xe5, 0x37, 0x2b, 0xcf, 0xf6, 0xd9, 0x6d, 0xdf, 0xe3, 0x98,
	0x2c, 0x58, 0xe2, 0xbb, 0x9c, 0xb3, 0x38, 0xee, 0xdf, 0xa1, 0x17, 0x64, 0xfd, 0x74, 0x19, 0xf6,
	0x0b, 0xcd, 0x9e, 0x2e, 0x7f, 0x08, 0x97, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x85, 0x69, 0xeb,
	0xc9, 0x1d, 0x04, 0x00, 0x00,
}
