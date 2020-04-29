// Code generated by protoc-gen-go. DO NOT EDIT.
// source: output/streaming/badger.proto

package streaming

import (
	fmt "fmt"
	execution "github.com/cube2222/octosql/execution"
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

type RecordData struct {
	Ids                  []*execution.RecordID `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	IsUndo               bool                  `protobuf:"varint,2,opt,name=isUndo,proto3" json:"isUndo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RecordData) Reset()         { *m = RecordData{} }
func (m *RecordData) String() string { return proto.CompactTextString(m) }
func (*RecordData) ProtoMessage()    {}
func (*RecordData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bd670ef2db0007a, []int{0}
}

func (m *RecordData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordData.Unmarshal(m, b)
}
func (m *RecordData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordData.Marshal(b, m, deterministic)
}
func (m *RecordData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordData.Merge(m, src)
}
func (m *RecordData) XXX_Size() int {
	return xxx_messageInfo_RecordData.Size(m)
}
func (m *RecordData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordData proto.InternalMessageInfo

func (m *RecordData) GetIds() []*execution.RecordID {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *RecordData) GetIsUndo() bool {
	if m != nil {
		return m.IsUndo
	}
	return false
}

func init() {
	proto.RegisterType((*RecordData)(nil), "execution.RecordData")
}

func init() { proto.RegisterFile("output/streaming/badger.proto", fileDescriptor_9bd670ef2db0007a) }

var fileDescriptor_9bd670ef2db0007a = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x2f, 0x2d, 0x29,
	0x28, 0x2d, 0xd1, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0xd7, 0x4f, 0x4a, 0x4c,
	0x49, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xad, 0x48, 0x4d, 0x2e,
	0x2d, 0xc9, 0xcc, 0xcf, 0x93, 0x12, 0x83, 0x33, 0xf5, 0x8b, 0x52, 0x93, 0xf3, 0x8b, 0x52, 0x20,
	0x4a, 0x94, 0xbc, 0xb9, 0xb8, 0x82, 0xc0, 0x7c, 0x97, 0xc4, 0x92, 0x44, 0x21, 0x55, 0x2e, 0xe6,
	0xcc, 0x94, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x61, 0x3d, 0xb8, 0x1e, 0x3d, 0x88,
	0x1a, 0x4f, 0x97, 0x20, 0x90, 0xbc, 0x90, 0x18, 0x17, 0x5b, 0x66, 0x71, 0x68, 0x5e, 0x4a, 0xbe,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x94, 0xe7, 0xa4, 0x17, 0xa5, 0x93, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x5c, 0x9a, 0x94, 0x6a, 0x64, 0x64, 0x64, 0xa4,
	0x9f, 0x9f, 0x5c, 0x92, 0x5f, 0x5c, 0x98, 0xa3, 0x8f, 0xee, 0xd8, 0x24, 0x36, 0xb0, 0x1b, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x93, 0x65, 0x65, 0xc7, 0x00, 0x00, 0x00,
}
