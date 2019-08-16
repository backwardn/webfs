// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cellspecs.proto

package models

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

type HTTPCellSpec struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	AuthHeader           string   `protobuf:"bytes,2,opt,name=auth_header,json=authHeader,proto3" json:"auth_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPCellSpec) Reset()         { *m = HTTPCellSpec{} }
func (m *HTTPCellSpec) String() string { return proto.CompactTextString(m) }
func (*HTTPCellSpec) ProtoMessage()    {}
func (*HTTPCellSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellspecs_1df46cd1462615db, []int{0}
}
func (m *HTTPCellSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPCellSpec.Unmarshal(m, b)
}
func (m *HTTPCellSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPCellSpec.Marshal(b, m, deterministic)
}
func (dst *HTTPCellSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPCellSpec.Merge(dst, src)
}
func (m *HTTPCellSpec) XXX_Size() int {
	return xxx_messageInfo_HTTPCellSpec.Size(m)
}
func (m *HTTPCellSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPCellSpec.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPCellSpec proto.InternalMessageInfo

func (m *HTTPCellSpec) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HTTPCellSpec) GetAuthHeader() string {
	if m != nil {
		return m.AuthHeader
	}
	return ""
}

func init() {
	proto.RegisterType((*HTTPCellSpec)(nil), "HTTPCellSpec")
}

func init() { proto.RegisterFile("cellspecs.proto", fileDescriptor_cellspecs_1df46cd1462615db) }

var fileDescriptor_cellspecs_1df46cd1462615db = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4e, 0xcd, 0xc9,
	0x29, 0x2e, 0x48, 0x4d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe4, 0xe2, 0xf1,
	0x08, 0x09, 0x09, 0x70, 0x4e, 0xcd, 0xc9, 0x09, 0x2e, 0x48, 0x4d, 0x16, 0x12, 0xe0, 0x62, 0x2e,
	0x2d, 0xca, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0xe4, 0xb9, 0xb8, 0x13,
	0x4b, 0x4b, 0x32, 0xe2, 0x33, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x98, 0xc0, 0x32, 0x5c, 0x20,
	0x21, 0x0f, 0xb0, 0x48, 0x12, 0x1b, 0xd8, 0x24, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0xb3, 0x92, 0x5c, 0x5c, 0x00, 0x00, 0x00,
}
