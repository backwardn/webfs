// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tree.proto

package wrds // import "github.com/brendoncarroll/webfs/pkg/wrds"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import webref "github.com/brendoncarroll/webfs/pkg/webref"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tree struct {
	Level                uint32       `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Entries              []*TreeEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}
func (*Tree) Descriptor() ([]byte, []int) {
	return fileDescriptor_tree_e38f47230227e6c0, []int{0}
}
func (m *Tree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tree.Unmarshal(m, b)
}
func (m *Tree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tree.Marshal(b, m, deterministic)
}
func (dst *Tree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tree.Merge(dst, src)
}
func (m *Tree) XXX_Size() int {
	return xxx_messageInfo_Tree.Size(m)
}
func (m *Tree) XXX_DiscardUnknown() {
	xxx_messageInfo_Tree.DiscardUnknown(m)
}

var xxx_messageInfo_Tree proto.InternalMessageInfo

func (m *Tree) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Tree) GetEntries() []*TreeEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TreeEntry struct {
	Key                  []byte      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Ref                  *webref.Ref `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TreeEntry) Reset()         { *m = TreeEntry{} }
func (m *TreeEntry) String() string { return proto.CompactTextString(m) }
func (*TreeEntry) ProtoMessage()    {}
func (*TreeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_tree_e38f47230227e6c0, []int{1}
}
func (m *TreeEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeEntry.Unmarshal(m, b)
}
func (m *TreeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeEntry.Marshal(b, m, deterministic)
}
func (dst *TreeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeEntry.Merge(dst, src)
}
func (m *TreeEntry) XXX_Size() int {
	return xxx_messageInfo_TreeEntry.Size(m)
}
func (m *TreeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TreeEntry proto.InternalMessageInfo

func (m *TreeEntry) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TreeEntry) GetRef() *webref.Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func init() {
	proto.RegisterType((*Tree)(nil), "wrds.Tree")
	proto.RegisterType((*TreeEntry)(nil), "wrds.TreeEntry")
}

func init() { proto.RegisterFile("tree.proto", fileDescriptor_tree_e38f47230227e6c0) }

var fileDescriptor_tree_e38f47230227e6c0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x3f, 0x4b, 0xc7, 0x30,
	0x10, 0x86, 0xe9, 0x1f, 0x95, 0x5e, 0x15, 0x25, 0x38, 0x14, 0x41, 0x28, 0x9d, 0xa2, 0x43, 0x02,
	0x75, 0x75, 0x12, 0xc4, 0x3d, 0x38, 0xb9, 0x99, 0xf6, 0x52, 0x4b, 0x63, 0x52, 0xae, 0xd1, 0xd2,
	0x6f, 0x2f, 0x69, 0xf1, 0xb7, 0xdd, 0x3d, 0xc7, 0xbd, 0xef, 0x03, 0x10, 0x08, 0x51, 0xcc, 0xe4,
	0x83, 0x67, 0xf9, 0x4a, 0xfd, 0x72, 0x57, 0x10, 0x9a, 0x03, 0x34, 0x6f, 0x90, 0xbf, 0x13, 0x22,
	0xbb, 0x85, 0x33, 0x8b, 0xbf, 0x68, 0xab, 0xa4, 0x4e, 0xf8, 0x95, 0x3a, 0x16, 0xf6, 0x00, 0x17,
	0xe8, 0x02, 0x8d, 0xb8, 0x54, 0x69, 0x9d, 0xf1, 0xb2, 0xbd, 0x16, 0x31, 0x40, 0xc4, 0x97, 0x57,
	0x17, 0x68, 0x53, 0xff, 0xf7, 0xe6, 0x19, 0x8a, 0x13, 0x65, 0x37, 0x90, 0x4d, 0xb8, 0xed, 0x59,
	0x97, 0x2a, 0x8e, 0xec, 0x1e, 0x32, 0x42, 0x53, 0xa5, 0x75, 0xc2, 0xcb, 0xb6, 0x14, 0x2b, 0xea,
	0xe8, 0xa0, 0xd0, 0xa8, 0xc8, 0x5f, 0x1e, 0x3f, 0xf8, 0x30, 0x86, 0xaf, 0x1f, 0x2d, 0x3a, 0xff,
	0x2d, 0x35, 0xa1, 0xeb, 0xbd, 0xeb, 0x3e, 0x89, 0xbc, 0xb5, 0x72, 0x45, 0x6d, 0x16, 0x39, 0x4f,
	0x83, 0x8c, 0xe5, 0xfa, 0x7c, 0x37, 0x7f, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x86, 0xb5, 0x9b,
	0x3b, 0xd8, 0x00, 0x00, 0x00,
}