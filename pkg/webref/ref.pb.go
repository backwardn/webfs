// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ref.proto

package webref

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EncAlgo int32

const (
	EncAlgo_UNKNOWN_EA EncAlgo = 0
	EncAlgo_AES256CTR  EncAlgo = 1
	EncAlgo_CHACHA20   EncAlgo = 2
)

var EncAlgo_name = map[int32]string{
	0: "UNKNOWN_EA",
	1: "AES256CTR",
	2: "CHACHA20",
}

var EncAlgo_value = map[string]int32{
	"UNKNOWN_EA": 0,
	"AES256CTR":  1,
	"CHACHA20":   2,
}

func (x EncAlgo) String() string {
	return proto.EnumName(EncAlgo_name, int32(x))
}

func (EncAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{0}
}

type CompressAlgo int32

const (
	CompressAlgo_UNKNOWN_CA CompressAlgo = 0
	CompressAlgo_SNAPPY     CompressAlgo = 1
)

var CompressAlgo_name = map[int32]string{
	0: "UNKNOWN_CA",
	1: "SNAPPY",
}

var CompressAlgo_value = map[string]int32{
	"UNKNOWN_CA": 0,
	"SNAPPY":     1,
}

func (x CompressAlgo) String() string {
	return proto.EnumName(CompressAlgo_name, int32(x))
}

func (CompressAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{1}
}

type Crypto struct {
	Ref                  *Ref     `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	EncAlgo              EncAlgo  `protobuf:"varint,2,opt,name=enc_algo,json=encAlgo,proto3,enum=webref.EncAlgo" json:"enc_algo,omitempty"`
	Dek                  []byte   `protobuf:"bytes,3,opt,name=dek,proto3" json:"dek,omitempty"`
	Length               int32    `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crypto) Reset()         { *m = Crypto{} }
func (m *Crypto) String() string { return proto.CompactTextString(m) }
func (*Crypto) ProtoMessage()    {}
func (*Crypto) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{0}
}

func (m *Crypto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crypto.Unmarshal(m, b)
}
func (m *Crypto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crypto.Marshal(b, m, deterministic)
}
func (m *Crypto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crypto.Merge(m, src)
}
func (m *Crypto) XXX_Size() int {
	return xxx_messageInfo_Crypto.Size(m)
}
func (m *Crypto) XXX_DiscardUnknown() {
	xxx_messageInfo_Crypto.DiscardUnknown(m)
}

var xxx_messageInfo_Crypto proto.InternalMessageInfo

func (m *Crypto) GetRef() *Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *Crypto) GetEncAlgo() EncAlgo {
	if m != nil {
		return m.EncAlgo
	}
	return EncAlgo_UNKNOWN_EA
}

func (m *Crypto) GetDek() []byte {
	if m != nil {
		return m.Dek
	}
	return nil
}

func (m *Crypto) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Slice struct {
	Ref                  *Ref     `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               int32    `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slice) Reset()         { *m = Slice{} }
func (m *Slice) String() string { return proto.CompactTextString(m) }
func (*Slice) ProtoMessage()    {}
func (*Slice) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{1}
}

func (m *Slice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slice.Unmarshal(m, b)
}
func (m *Slice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slice.Marshal(b, m, deterministic)
}
func (m *Slice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slice.Merge(m, src)
}
func (m *Slice) XXX_Size() int {
	return xxx_messageInfo_Slice.Size(m)
}
func (m *Slice) XXX_DiscardUnknown() {
	xxx_messageInfo_Slice.DiscardUnknown(m)
}

var xxx_messageInfo_Slice proto.InternalMessageInfo

func (m *Slice) GetRef() *Ref {
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
	Replicas             []*Ref   `protobuf:"bytes,1,rep,name=replicas,proto3" json:"replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mirror) Reset()         { *m = Mirror{} }
func (m *Mirror) String() string { return proto.CompactTextString(m) }
func (*Mirror) ProtoMessage()    {}
func (*Mirror) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{2}
}

func (m *Mirror) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mirror.Unmarshal(m, b)
}
func (m *Mirror) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mirror.Marshal(b, m, deterministic)
}
func (m *Mirror) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mirror.Merge(m, src)
}
func (m *Mirror) XXX_Size() int {
	return xxx_messageInfo_Mirror.Size(m)
}
func (m *Mirror) XXX_DiscardUnknown() {
	xxx_messageInfo_Mirror.DiscardUnknown(m)
}

var xxx_messageInfo_Mirror proto.InternalMessageInfo

func (m *Mirror) GetReplicas() []*Ref {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type ReedSolomon struct {
	Pieces               []*Ref   `protobuf:"bytes,1,rep,name=pieces,proto3" json:"pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReedSolomon) Reset()         { *m = ReedSolomon{} }
func (m *ReedSolomon) String() string { return proto.CompactTextString(m) }
func (*ReedSolomon) ProtoMessage()    {}
func (*ReedSolomon) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{3}
}

func (m *ReedSolomon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReedSolomon.Unmarshal(m, b)
}
func (m *ReedSolomon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReedSolomon.Marshal(b, m, deterministic)
}
func (m *ReedSolomon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReedSolomon.Merge(m, src)
}
func (m *ReedSolomon) XXX_Size() int {
	return xxx_messageInfo_ReedSolomon.Size(m)
}
func (m *ReedSolomon) XXX_DiscardUnknown() {
	xxx_messageInfo_ReedSolomon.DiscardUnknown(m)
}

var xxx_messageInfo_ReedSolomon proto.InternalMessageInfo

func (m *ReedSolomon) GetPieces() []*Ref {
	if m != nil {
		return m.Pieces
	}
	return nil
}

type Annotated struct {
	Ref                  *Ref              `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Annotated) Reset()         { *m = Annotated{} }
func (m *Annotated) String() string { return proto.CompactTextString(m) }
func (*Annotated) ProtoMessage()    {}
func (*Annotated) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{4}
}

func (m *Annotated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Annotated.Unmarshal(m, b)
}
func (m *Annotated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Annotated.Marshal(b, m, deterministic)
}
func (m *Annotated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Annotated.Merge(m, src)
}
func (m *Annotated) XXX_Size() int {
	return xxx_messageInfo_Annotated.Size(m)
}
func (m *Annotated) XXX_DiscardUnknown() {
	xxx_messageInfo_Annotated.DiscardUnknown(m)
}

var xxx_messageInfo_Annotated proto.InternalMessageInfo

func (m *Annotated) GetRef() *Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *Annotated) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type Compress struct {
	Ref                  *Ref         `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Algo                 CompressAlgo `protobuf:"varint,2,opt,name=algo,proto3,enum=webref.CompressAlgo" json:"algo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Compress) Reset()         { *m = Compress{} }
func (m *Compress) String() string { return proto.CompactTextString(m) }
func (*Compress) ProtoMessage()    {}
func (*Compress) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{5}
}

func (m *Compress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Compress.Unmarshal(m, b)
}
func (m *Compress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Compress.Marshal(b, m, deterministic)
}
func (m *Compress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Compress.Merge(m, src)
}
func (m *Compress) XXX_Size() int {
	return xxx_messageInfo_Compress.Size(m)
}
func (m *Compress) XXX_DiscardUnknown() {
	xxx_messageInfo_Compress.DiscardUnknown(m)
}

var xxx_messageInfo_Compress proto.InternalMessageInfo

func (m *Compress) GetRef() *Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *Compress) GetAlgo() CompressAlgo {
	if m != nil {
		return m.Algo
	}
	return CompressAlgo_UNKNOWN_CA
}

type Ref struct {
	// Types that are valid to be assigned to Ref:
	//	*Ref_Url
	//	*Ref_Crypto
	//	*Ref_Slice
	//	*Ref_Mirror
	//	*Ref_Annotated
	//	*Ref_ReedSolomon
	//	*Ref_Compress
	Ref                  isRef_Ref `protobuf_oneof:"ref"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}
func (*Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_65d958559ea81b29, []int{6}
}

func (m *Ref) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ref.Unmarshal(m, b)
}
func (m *Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ref.Marshal(b, m, deterministic)
}
func (m *Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ref.Merge(m, src)
}
func (m *Ref) XXX_Size() int {
	return xxx_messageInfo_Ref.Size(m)
}
func (m *Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_Ref proto.InternalMessageInfo

type isRef_Ref interface {
	isRef_Ref()
}

type Ref_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type Ref_Crypto struct {
	Crypto *Crypto `protobuf:"bytes,2,opt,name=crypto,proto3,oneof"`
}

type Ref_Slice struct {
	Slice *Slice `protobuf:"bytes,3,opt,name=slice,proto3,oneof"`
}

type Ref_Mirror struct {
	Mirror *Mirror `protobuf:"bytes,4,opt,name=mirror,proto3,oneof"`
}

type Ref_Annotated struct {
	Annotated *Annotated `protobuf:"bytes,5,opt,name=annotated,proto3,oneof"`
}

type Ref_ReedSolomon struct {
	ReedSolomon *ReedSolomon `protobuf:"bytes,6,opt,name=reed_solomon,json=reedSolomon,proto3,oneof"`
}

type Ref_Compress struct {
	Compress *Compress `protobuf:"bytes,7,opt,name=compress,proto3,oneof"`
}

func (*Ref_Url) isRef_Ref() {}

func (*Ref_Crypto) isRef_Ref() {}

func (*Ref_Slice) isRef_Ref() {}

func (*Ref_Mirror) isRef_Ref() {}

func (*Ref_Annotated) isRef_Ref() {}

func (*Ref_ReedSolomon) isRef_Ref() {}

func (*Ref_Compress) isRef_Ref() {}

func (m *Ref) GetRef() isRef_Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *Ref) GetUrl() string {
	if x, ok := m.GetRef().(*Ref_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Ref) GetCrypto() *Crypto {
	if x, ok := m.GetRef().(*Ref_Crypto); ok {
		return x.Crypto
	}
	return nil
}

func (m *Ref) GetSlice() *Slice {
	if x, ok := m.GetRef().(*Ref_Slice); ok {
		return x.Slice
	}
	return nil
}

func (m *Ref) GetMirror() *Mirror {
	if x, ok := m.GetRef().(*Ref_Mirror); ok {
		return x.Mirror
	}
	return nil
}

func (m *Ref) GetAnnotated() *Annotated {
	if x, ok := m.GetRef().(*Ref_Annotated); ok {
		return x.Annotated
	}
	return nil
}

func (m *Ref) GetReedSolomon() *ReedSolomon {
	if x, ok := m.GetRef().(*Ref_ReedSolomon); ok {
		return x.ReedSolomon
	}
	return nil
}

func (m *Ref) GetCompress() *Compress {
	if x, ok := m.GetRef().(*Ref_Compress); ok {
		return x.Compress
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Ref) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Ref_Url)(nil),
		(*Ref_Crypto)(nil),
		(*Ref_Slice)(nil),
		(*Ref_Mirror)(nil),
		(*Ref_Annotated)(nil),
		(*Ref_ReedSolomon)(nil),
		(*Ref_Compress)(nil),
	}
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
	return fileDescriptor_65d958559ea81b29, []int{7}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
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
	return EncAlgo_UNKNOWN_EA
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
	proto.RegisterEnum("webref.EncAlgo", EncAlgo_name, EncAlgo_value)
	proto.RegisterEnum("webref.CompressAlgo", CompressAlgo_name, CompressAlgo_value)
	proto.RegisterType((*Crypto)(nil), "webref.Crypto")
	proto.RegisterType((*Slice)(nil), "webref.Slice")
	proto.RegisterType((*Mirror)(nil), "webref.Mirror")
	proto.RegisterType((*ReedSolomon)(nil), "webref.ReedSolomon")
	proto.RegisterType((*Annotated)(nil), "webref.Annotated")
	proto.RegisterMapType((map[string]string)(nil), "webref.Annotated.AnnotationsEntry")
	proto.RegisterType((*Compress)(nil), "webref.Compress")
	proto.RegisterType((*Ref)(nil), "webref.Ref")
	proto.RegisterType((*Options)(nil), "webref.Options")
	proto.RegisterMapType((map[string]string)(nil), "webref.Options.AttrsEntry")
	proto.RegisterMapType((map[string]int32)(nil), "webref.Options.ReplicasEntry")
}

func init() { proto.RegisterFile("ref.proto", fileDescriptor_65d958559ea81b29) }

var fileDescriptor_65d958559ea81b29 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0x8d, 0x93, 0xda, 0x49, 0xae, 0xd3, 0xd6, 0x3f, 0xfd, 0x4a, 0x31, 0x81, 0xb2, 0xe0, 0x31,
	0x96, 0x85, 0x91, 0xb4, 0x19, 0x2b, 0xdd, 0x06, 0x03, 0x37, 0x0b, 0x18, 0xb6, 0xa5, 0x45, 0xd9,
	0x1f, 0xb6, 0x97, 0xe0, 0x38, 0xd7, 0x69, 0xa8, 0x63, 0x05, 0xd9, 0x59, 0xe9, 0xc3, 0xbe, 0xd3,
	0x1e, 0xf7, 0x85, 0xf6, 0x3d, 0x86, 0x25, 0xdb, 0x75, 0xb3, 0x41, 0xd8, 0x9b, 0x74, 0xef, 0x39,
	0x47, 0xf2, 0xb9, 0x47, 0x86, 0x3a, 0x47, 0xbf, 0xbb, 0xe2, 0x2c, 0x66, 0x44, 0xbb, 0xc1, 0x29,
	0x47, 0xdf, 0xfa, 0x0e, 0xda, 0x80, 0xdf, 0xae, 0x62, 0x46, 0x8e, 0xa0, 0xc2, 0xd1, 0x37, 0x95,
	0x96, 0xd2, 0xd6, 0xfb, 0x7a, 0x57, 0xf6, 0xbb, 0x14, 0x7d, 0x9a, 0xd4, 0x49, 0x07, 0x6a, 0x18,
	0x7a, 0x13, 0x37, 0x98, 0x33, 0xb3, 0xdc, 0x52, 0xda, 0x7b, 0xfd, 0xfd, 0x0c, 0x33, 0x0c, 0x3d,
	0x3b, 0x98, 0x33, 0x5a, 0x45, 0xb9, 0x20, 0x06, 0x54, 0x66, 0x78, 0x6d, 0x56, 0x5a, 0x4a, 0xbb,
	0x41, 0x93, 0x25, 0x39, 0x04, 0x2d, 0xc0, 0x70, 0x1e, 0x5f, 0x99, 0x3b, 0x2d, 0xa5, 0xad, 0xd2,
	0x74, 0x67, 0x7d, 0x02, 0x75, 0x1c, 0x2c, 0x3c, 0xdc, 0x76, 0xfa, 0x21, 0x68, 0xcc, 0xf7, 0x23,
	0x8c, 0xc5, 0xd9, 0x2a, 0x4d, 0x77, 0x05, 0xdd, 0xca, 0x3d, 0xdd, 0x13, 0xd0, 0xde, 0x2f, 0x38,
	0x67, 0x9c, 0x3c, 0x86, 0x1a, 0xc7, 0x55, 0xb0, 0xf0, 0xdc, 0xc8, 0x54, 0x5a, 0x95, 0x4d, 0xf5,
	0xbc, 0x69, 0xf5, 0x41, 0xa7, 0x88, 0xb3, 0x31, 0x0b, 0xd8, 0x92, 0x85, 0xe4, 0x21, 0x68, 0xab,
	0x05, 0x7a, 0xf8, 0x57, 0x56, 0xda, 0xb2, 0x7e, 0x28, 0x50, 0xb7, 0xc3, 0x90, 0xc5, 0x6e, 0x8c,
	0xb3, 0x6d, 0xdf, 0xf0, 0x06, 0x74, 0x57, 0x62, 0x17, 0x2c, 0x8c, 0xcc, 0xb2, 0x90, 0xb5, 0x32,
	0x58, 0x2e, 0x93, 0xad, 0x12, 0xd0, 0x30, 0x8c, 0xf9, 0x2d, 0x2d, 0xd2, 0x9a, 0xaf, 0xc1, 0xd8,
	0x04, 0x24, 0x7e, 0x5f, 0xe3, 0xad, 0x38, 0xb8, 0x4e, 0x93, 0x25, 0x39, 0x00, 0xf5, 0x9b, 0x1b,
	0xac, 0x51, 0xd8, 0x55, 0xa7, 0x72, 0xf3, 0xb2, 0x7c, 0xa6, 0x58, 0x63, 0xa8, 0x0d, 0xd8, 0x72,
	0xc5, 0x31, 0x8a, 0xb6, 0x5d, 0xb8, 0x0d, 0x3b, 0x85, 0x71, 0x1f, 0x64, 0xfd, 0x8c, 0x2e, 0x66,
	0x2e, 0x10, 0xd6, 0xcf, 0x32, 0x54, 0x28, 0xfa, 0x84, 0x40, 0x65, 0xcd, 0x03, 0x79, 0x11, 0xa7,
	0x44, 0x93, 0x0d, 0x69, 0x83, 0xe6, 0x89, 0x84, 0x09, 0x1d, 0xbd, 0xbf, 0x97, 0xeb, 0x88, 0xaa,
	0x53, 0xa2, 0x69, 0x9f, 0x3c, 0x02, 0x35, 0x4a, 0xc2, 0x20, 0x66, 0xa9, 0xf7, 0x77, 0x33, 0xa0,
	0x48, 0x88, 0x53, 0xa2, 0xb2, 0x9b, 0x08, 0x2e, 0xc5, 0x6c, 0x45, 0x96, 0x0a, 0x82, 0x72, 0xe2,
	0x89, 0xa0, 0xec, 0x93, 0x13, 0xa8, 0xbb, 0x99, 0xad, 0xa6, 0x2a, 0xc0, 0xff, 0xfd, 0xe1, 0xb7,
	0x53, 0xa2, 0x77, 0x28, 0x72, 0x06, 0x0d, 0x8e, 0x38, 0x9b, 0x44, 0x32, 0x06, 0xa6, 0x26, 0x58,
	0xff, 0xdf, 0x79, 0x93, 0x27, 0xc4, 0x29, 0x51, 0x9d, 0x17, 0x02, 0xd3, 0x85, 0x9a, 0x97, 0x3a,
	0x63, 0x56, 0x05, 0xcb, 0xd8, 0x74, 0xcc, 0x29, 0xd1, 0x1c, 0x73, 0xae, 0x0a, 0xf3, 0xad, 0x5f,
	0x65, 0xa8, 0x5e, 0xac, 0xc4, 0x30, 0xef, 0xbd, 0x31, 0x65, 0xcb, 0x1b, 0x7b, 0x00, 0x7a, 0x84,
	0x1e, 0xc7, 0x78, 0x12, 0x21, 0xce, 0x84, 0xb7, 0x0d, 0x0a, 0xb2, 0x34, 0x46, 0x9c, 0x91, 0x27,
	0x60, 0xb0, 0xa9, 0xbf, 0x8e, 0x3c, 0x37, 0xc6, 0x49, 0xe1, 0x91, 0xd4, 0xe8, 0x7e, 0x5e, 0x7f,
	0x27, 0xca, 0xe4, 0x45, 0xe1, 0x8d, 0xec, 0x88, 0x58, 0x1e, 0x65, 0xe7, 0xa6, 0x57, 0xeb, 0xd2,
	0xb4, 0x2f, 0x13, 0x99, 0xc3, 0xc9, 0x31, 0xa8, 0x6e, 0x1c, 0xf3, 0xc8, 0x54, 0x05, 0xaf, 0xb9,
	0xc9, 0xb3, 0x93, 0xa6, 0x24, 0x49, 0x60, 0xf3, 0x15, 0xec, 0xde, 0x13, 0xdb, 0x96, 0x5e, 0xb5,
	0x90, 0xde, 0xe6, 0x19, 0xc0, 0x9d, 0xe2, 0xbf, 0xe4, 0xbe, 0x73, 0x0a, 0xd5, 0xd4, 0x43, 0xb2,
	0x07, 0xf0, 0x71, 0xf4, 0x76, 0x74, 0xf1, 0x79, 0x34, 0x19, 0xda, 0x46, 0x89, 0xec, 0x42, 0xdd,
	0x1e, 0x8e, 0xfb, 0xcf, 0x4f, 0x07, 0x1f, 0xa8, 0xa1, 0x90, 0x06, 0xd4, 0x06, 0x8e, 0x3d, 0x70,
	0xec, 0xfe, 0xb1, 0x51, 0xee, 0x74, 0xa0, 0x51, 0x0c, 0x7c, 0x91, 0x3c, 0x48, 0xc8, 0x00, 0xda,
	0x78, 0x64, 0x5f, 0x5e, 0x7e, 0x31, 0x94, 0xf3, 0xa7, 0x5f, 0x3b, 0xf3, 0x45, 0x7c, 0xb5, 0x9e,
	0x76, 0x3d, 0xb6, 0xec, 0x4d, 0x39, 0x86, 0x33, 0x16, 0x7a, 0x2e, 0xe7, 0x2c, 0x08, 0x7a, 0x37,
	0x38, 0xf5, 0xa3, 0xde, 0xea, 0x7a, 0xde, 0x93, 0x16, 0x4d, 0x35, 0xf1, 0x27, 0x7e, 0xf6, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x36, 0x08, 0xbd, 0xd9, 0x96, 0x05, 0x00, 0x00,
}
