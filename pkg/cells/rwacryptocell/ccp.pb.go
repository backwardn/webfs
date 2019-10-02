// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ccp.proto

package rwacryptocell

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

type CellContents struct {
	Who                  *Who           `protobuf:"bytes,1,opt,name=who,proto3" json:"who,omitempty"`
	WhoAuthor            int32          `protobuf:"varint,2,opt,name=who_author,json=whoAuthor,proto3" json:"who_author,omitempty"`
	WhoSigs              map[int32]*Sig `protobuf:"bytes,3,rep,name=who_sigs,json=whoSigs,proto3" json:"who_sigs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	What                 *What          `protobuf:"bytes,4,opt,name=what,proto3" json:"what,omitempty"`
	WhatAuthor           int32          `protobuf:"varint,5,opt,name=what_author,json=whatAuthor,proto3" json:"what_author,omitempty"`
	WhatSig              *Sig           `protobuf:"bytes,6,opt,name=what_sig,json=whatSig,proto3" json:"what_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CellContents) Reset()         { *m = CellContents{} }
func (m *CellContents) String() string { return proto.CompactTextString(m) }
func (*CellContents) ProtoMessage()    {}
func (*CellContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d78a8e7a0223a, []int{0}
}

func (m *CellContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellContents.Unmarshal(m, b)
}
func (m *CellContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellContents.Marshal(b, m, deterministic)
}
func (m *CellContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellContents.Merge(m, src)
}
func (m *CellContents) XXX_Size() int {
	return xxx_messageInfo_CellContents.Size(m)
}
func (m *CellContents) XXX_DiscardUnknown() {
	xxx_messageInfo_CellContents.DiscardUnknown(m)
}

var xxx_messageInfo_CellContents proto.InternalMessageInfo

func (m *CellContents) GetWho() *Who {
	if m != nil {
		return m.Who
	}
	return nil
}

func (m *CellContents) GetWhoAuthor() int32 {
	if m != nil {
		return m.WhoAuthor
	}
	return 0
}

func (m *CellContents) GetWhoSigs() map[int32]*Sig {
	if m != nil {
		return m.WhoSigs
	}
	return nil
}

func (m *CellContents) GetWhat() *What {
	if m != nil {
		return m.What
	}
	return nil
}

func (m *CellContents) GetWhatAuthor() int32 {
	if m != nil {
		return m.WhatAuthor
	}
	return 0
}

func (m *CellContents) GetWhatSig() *Sig {
	if m != nil {
		return m.WhatSig
	}
	return nil
}

type Who struct {
	Entities             []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	Admin                []int32   `protobuf:"varint,2,rep,packed,name=admin,proto3" json:"admin,omitempty"`
	Write                []int32   `protobuf:"varint,3,rep,packed,name=write,proto3" json:"write,omitempty"`
	Read                 []int32   `protobuf:"varint,4,rep,packed,name=read,proto3" json:"read,omitempty"`
	Gen                  int64     `protobuf:"varint,5,opt,name=gen,proto3" json:"gen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Who) Reset()         { *m = Who{} }
func (m *Who) String() string { return proto.CompactTextString(m) }
func (*Who) ProtoMessage()    {}
func (*Who) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d78a8e7a0223a, []int{1}
}

func (m *Who) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Who.Unmarshal(m, b)
}
func (m *Who) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Who.Marshal(b, m, deterministic)
}
func (m *Who) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Who.Merge(m, src)
}
func (m *Who) XXX_Size() int {
	return xxx_messageInfo_Who.Size(m)
}
func (m *Who) XXX_DiscardUnknown() {
	xxx_messageInfo_Who.DiscardUnknown(m)
}

var xxx_messageInfo_Who proto.InternalMessageInfo

func (m *Who) GetEntities() []*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *Who) GetAdmin() []int32 {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *Who) GetWrite() []int32 {
	if m != nil {
		return m.Write
	}
	return nil
}

func (m *Who) GetRead() []int32 {
	if m != nil {
		return m.Read
	}
	return nil
}

func (m *Who) GetGen() int64 {
	if m != nil {
		return m.Gen
	}
	return 0
}

type What struct {
	Payload              *EncMsg       `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Deks                 []*AsymEncMsg `protobuf:"bytes,2,rep,name=deks,proto3" json:"deks,omitempty"`
	Gen                  int64         `protobuf:"varint,3,opt,name=gen,proto3" json:"gen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *What) Reset()         { *m = What{} }
func (m *What) String() string { return proto.CompactTextString(m) }
func (*What) ProtoMessage()    {}
func (*What) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d78a8e7a0223a, []int{2}
}

func (m *What) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_What.Unmarshal(m, b)
}
func (m *What) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_What.Marshal(b, m, deterministic)
}
func (m *What) XXX_Merge(src proto.Message) {
	xxx_messageInfo_What.Merge(m, src)
}
func (m *What) XXX_Size() int {
	return xxx_messageInfo_What.Size(m)
}
func (m *What) XXX_DiscardUnknown() {
	xxx_messageInfo_What.DiscardUnknown(m)
}

var xxx_messageInfo_What proto.InternalMessageInfo

func (m *What) GetPayload() *EncMsg {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *What) GetDeks() []*AsymEncMsg {
	if m != nil {
		return m.Deks
	}
	return nil
}

func (m *What) GetGen() int64 {
	if m != nil {
		return m.Gen
	}
	return 0
}

func init() {
	proto.RegisterType((*CellContents)(nil), "ccp.CellContents")
	proto.RegisterMapType((map[int32]*Sig)(nil), "ccp.CellContents.WhoSigsEntry")
	proto.RegisterType((*Who)(nil), "ccp.Who")
	proto.RegisterType((*What)(nil), "ccp.What")
}

func init() { proto.RegisterFile("ccp.proto", fileDescriptor_879d78a8e7a0223a) }

var fileDescriptor_879d78a8e7a0223a = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x86, 0x71, 0x64, 0xef, 0x26, 0xe3, 0x40, 0x8b, 0xe8, 0xc1, 0x04, 0x76, 0x1b, 0xb2, 0x94,
	0xe6, 0x14, 0x43, 0x7a, 0x68, 0xb7, 0xb7, 0xed, 0x76, 0x8f, 0xbd, 0x28, 0x87, 0x85, 0x5e, 0x8a,
	0x22, 0xab, 0xb6, 0x88, 0x63, 0x19, 0x49, 0xa9, 0xf1, 0xad, 0x6f, 0xd4, 0x57, 0x2c, 0x33, 0x4e,
	0x82, 0x4f, 0x9e, 0xf9, 0x67, 0x34, 0xdf, 0xfc, 0xb2, 0x60, 0xa6, 0x54, 0xbb, 0x69, 0x9d, 0x0d,
	0x96, 0x33, 0xa5, 0xda, 0xc5, 0x5c, 0xb9, 0xbe, 0x0d, 0x76, 0x90, 0x56, 0xff, 0x26, 0x30, 0x7f,
	0xd6, 0x75, 0xfd, 0x6c, 0x9b, 0xa0, 0x9b, 0xe0, 0xf9, 0x02, 0x58, 0x57, 0xd9, 0x2c, 0x5a, 0x46,
	0xeb, 0x74, 0x3b, 0xdd, 0xe0, 0xe1, 0xd7, 0xca, 0x0a, 0x14, 0xf9, 0x1d, 0x40, 0x57, 0xd9, 0x5f,
	0xf2, 0x14, 0x2a, 0xeb, 0xb2, 0xc9, 0x32, 0x5a, 0x27, 0x62, 0xd6, 0x55, 0xf6, 0x89, 0x04, 0xfe,
	0x08, 0x53, 0x2c, 0x7b, 0x53, 0xfa, 0x8c, 0x2d, 0xd9, 0x3a, 0xdd, 0xde, 0xd3, 0xf9, 0xf1, 0x7c,
	0x1c, 0xb6, 0x33, 0xa5, 0x7f, 0x69, 0x82, 0xeb, 0xc5, 0x6d, 0x37, 0x64, 0xfc, 0x0e, 0xe2, 0xae,
	0x92, 0x21, 0x8b, 0x09, 0x3b, 0x3b, 0x63, 0x65, 0x10, 0x24, 0xf3, 0xf7, 0x90, 0xe2, 0xf7, 0x42,
	0x4e, 0x88, 0x0c, 0x28, 0x9d, 0xd1, 0x0f, 0x88, 0x96, 0x01, 0xd9, 0xd9, 0xcd, 0x68, 0xf5, 0x9d,
	0x29, 0x11, 0x22, 0xc3, 0xce, 0x94, 0x8b, 0xef, 0x30, 0x1f, 0xd3, 0xf9, 0x5b, 0x60, 0x07, 0xdd,
	0x93, 0xd5, 0x44, 0x60, 0xc8, 0xef, 0x21, 0xf9, 0x23, 0xeb, 0x93, 0x26, 0x6f, 0xe3, 0x19, 0x83,
	0xfc, 0x75, 0xf2, 0x25, 0x5a, 0xfd, 0x8d, 0x80, 0xbd, 0x56, 0x96, 0x7f, 0x84, 0xa9, 0x6e, 0x82,
	0x09, 0x46, 0xfb, 0x2c, 0x22, 0xb7, 0x29, 0xb5, 0xbf, 0xa0, 0xd8, 0x8b, 0x6b, 0x91, 0xbf, 0x83,
	0x44, 0x16, 0x47, 0xd3, 0x64, 0x93, 0x25, 0x5b, 0x27, 0x62, 0x48, 0x50, 0xed, 0x9c, 0x09, 0x9a,
	0x6e, 0x2a, 0x11, 0x43, 0xc2, 0x39, 0xc4, 0x4e, 0xcb, 0x22, 0x8b, 0x49, 0xa4, 0x18, 0xd7, 0x2c,
	0x75, 0x43, 0xa6, 0x99, 0xc0, 0x70, 0x55, 0x40, 0x8c, 0x97, 0xc3, 0x3f, 0xc0, 0x6d, 0x2b, 0xfb,
	0xda, 0xca, 0xe2, 0xfc, 0xbf, 0x2e, 0x1b, 0xa8, 0x1f, 0xbe, 0x14, 0x97, 0x1a, 0x7f, 0x80, 0xb8,
	0xd0, 0x07, 0x4f, 0xfc, 0x74, 0xfb, 0x86, 0x7a, 0x9e, 0x7c, 0x7f, 0x3c, 0xf7, 0x51, 0xf1, 0x42,
	0x61, 0x57, 0xca, 0xb7, 0xc7, 0x9f, 0x9f, 0x4b, 0x13, 0xaa, 0xd3, 0x7e, 0xa3, 0xec, 0x31, 0xdf,
	0x3b, 0xdd, 0x14, 0xb6, 0x51, 0xd2, 0x39, 0x5b, 0xd7, 0x79, 0xa7, 0xf7, 0xbf, 0x7d, 0xde, 0x1e,
	0xca, 0x5c, 0xe9, 0xba, 0xf6, 0xb9, 0xeb, 0xe4, 0xf0, 0xae, 0x30, 0xdd, 0xdf, 0xd0, 0xe3, 0xfa,
	0xf4, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x61, 0x14, 0xf4, 0x17, 0x7c, 0x02, 0x00, 0x00,
}