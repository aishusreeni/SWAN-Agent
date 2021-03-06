// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sl_route_ipv4.proto

package service_layer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// IPv4 route
type SLRoutev4 struct {
	// IPv4 Prefix
	// Valid addresses:
	//     0.0.0.0
	//     1.0.0.0 - 223.255.255.255
	Prefix uint32 `protobuf:"varint,1,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	// IPv4 prefix length, [0-32]
	PrefixLen uint32 `protobuf:"varint,2,opt,name=PrefixLen,proto3" json:"PrefixLen,omitempty"`
	// Common route attributes
	RouteCommon *SLRouteCommon `protobuf:"bytes,3,opt,name=RouteCommon,proto3" json:"RouteCommon,omitempty"`
	// List of route paths for a particular route.
	// Specifying more than one path is allowed for ECMP/UCMP cases
	PathList             []*SLRoutePath `protobuf:"bytes,4,rep,name=PathList,proto3" json:"PathList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SLRoutev4) Reset()         { *m = SLRoutev4{} }
func (m *SLRoutev4) String() string { return proto.CompactTextString(m) }
func (*SLRoutev4) ProtoMessage()    {}
func (*SLRoutev4) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25f7b0c2338e8af, []int{0}
}

func (m *SLRoutev4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SLRoutev4.Unmarshal(m, b)
}
func (m *SLRoutev4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SLRoutev4.Marshal(b, m, deterministic)
}
func (m *SLRoutev4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SLRoutev4.Merge(m, src)
}
func (m *SLRoutev4) XXX_Size() int {
	return xxx_messageInfo_SLRoutev4.Size(m)
}
func (m *SLRoutev4) XXX_DiscardUnknown() {
	xxx_messageInfo_SLRoutev4.DiscardUnknown(m)
}

var xxx_messageInfo_SLRoutev4 proto.InternalMessageInfo

func (m *SLRoutev4) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *SLRoutev4) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

func (m *SLRoutev4) GetRouteCommon() *SLRouteCommon {
	if m != nil {
		return m.RouteCommon
	}
	return nil
}

func (m *SLRoutev4) GetPathList() []*SLRoutePath {
	if m != nil {
		return m.PathList
	}
	return nil
}

// List of routes for bulk download
type SLRoutev4Msg struct {
	// Route Object Operations
	Oper SLObjectOp `protobuf:"varint,1,opt,name=Oper,proto3,enum=service_layer.SLObjectOp" json:"Oper,omitempty"`
	// Correlator. This can be used to correlate replies with requests.
	// The Server simply reflects this field back in the reply.
	Correlator uint64 `protobuf:"varint,2,opt,name=Correlator,proto3" json:"Correlator,omitempty"`
	// VRF name.
	VrfName string `protobuf:"bytes,3,opt,name=VrfName,proto3" json:"VrfName,omitempty"`
	// List of routes for the VRF specified above
	Routes               []*SLRoutev4 `protobuf:"bytes,4,rep,name=Routes,proto3" json:"Routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SLRoutev4Msg) Reset()         { *m = SLRoutev4Msg{} }
func (m *SLRoutev4Msg) String() string { return proto.CompactTextString(m) }
func (*SLRoutev4Msg) ProtoMessage()    {}
func (*SLRoutev4Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25f7b0c2338e8af, []int{1}
}

func (m *SLRoutev4Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SLRoutev4Msg.Unmarshal(m, b)
}
func (m *SLRoutev4Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SLRoutev4Msg.Marshal(b, m, deterministic)
}
func (m *SLRoutev4Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SLRoutev4Msg.Merge(m, src)
}
func (m *SLRoutev4Msg) XXX_Size() int {
	return xxx_messageInfo_SLRoutev4Msg.Size(m)
}
func (m *SLRoutev4Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_SLRoutev4Msg.DiscardUnknown(m)
}

var xxx_messageInfo_SLRoutev4Msg proto.InternalMessageInfo

func (m *SLRoutev4Msg) GetOper() SLObjectOp {
	if m != nil {
		return m.Oper
	}
	return SLObjectOp_SL_OBJOP_RESERVED
}

func (m *SLRoutev4Msg) GetCorrelator() uint64 {
	if m != nil {
		return m.Correlator
	}
	return 0
}

func (m *SLRoutev4Msg) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SLRoutev4Msg) GetRoutes() []*SLRoutev4 {
	if m != nil {
		return m.Routes
	}
	return nil
}

// IPv4 route result, uniquely identified by the Prefix/PrefixLen pair
type SLRoutev4Res struct {
	// Corresponding error code
	ErrStatus *SLErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus,proto3" json:"ErrStatus,omitempty"`
	// IPv4 Prefix
	Prefix uint32 `protobuf:"varint,2,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	// IPv4 prefix length, [0-32]
	PrefixLen            uint32   `protobuf:"varint,3,opt,name=PrefixLen,proto3" json:"PrefixLen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SLRoutev4Res) Reset()         { *m = SLRoutev4Res{} }
func (m *SLRoutev4Res) String() string { return proto.CompactTextString(m) }
func (*SLRoutev4Res) ProtoMessage()    {}
func (*SLRoutev4Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25f7b0c2338e8af, []int{2}
}

func (m *SLRoutev4Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SLRoutev4Res.Unmarshal(m, b)
}
func (m *SLRoutev4Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SLRoutev4Res.Marshal(b, m, deterministic)
}
func (m *SLRoutev4Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SLRoutev4Res.Merge(m, src)
}
func (m *SLRoutev4Res) XXX_Size() int {
	return xxx_messageInfo_SLRoutev4Res.Size(m)
}
func (m *SLRoutev4Res) XXX_DiscardUnknown() {
	xxx_messageInfo_SLRoutev4Res.DiscardUnknown(m)
}

var xxx_messageInfo_SLRoutev4Res proto.InternalMessageInfo

func (m *SLRoutev4Res) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLRoutev4Res) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *SLRoutev4Res) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

// IPv4 bulk route result status
type SLRoutev4MsgRsp struct {
	// Correlator. This can be used to correlate replies with requests.
	// The Server simply reflects this field back in the reply.
	Correlator uint64 `protobuf:"varint,1,opt,name=Correlator,proto3" json:"Correlator,omitempty"`
	// VRF name (matches the VRF name of the original operation)
	VrfName string `protobuf:"bytes,2,opt,name=VrfName,proto3" json:"VrfName,omitempty"`
	// Summary result of the bulk operation (refer to enum SLErrorStatus)
	// In general, the StatusSummary implies one of 3 things:
	// 1. SL_SUCCESS: signifies that the entire bulk operation was successful.
	//         In this case, the Results list is empty.
	// 2. SL_SOME_ERR: signifies that the operation failed for one or more
	//         entries. In this case, Results holds the result for
	//         each individual entry in the bulk.
	// 3. SL_RPC_XXX: signifies that the entire bulk operation failed.
	//         In this case, the Results list is empty.
	StatusSummary *SLErrorStatus `protobuf:"bytes,3,opt,name=StatusSummary,proto3" json:"StatusSummary,omitempty"`
	// In case of errors, this field indicates which entry in the bulk was
	// erroneous.
	Results              []*SLRoutev4Res `protobuf:"bytes,4,rep,name=Results,proto3" json:"Results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SLRoutev4MsgRsp) Reset()         { *m = SLRoutev4MsgRsp{} }
func (m *SLRoutev4MsgRsp) String() string { return proto.CompactTextString(m) }
func (*SLRoutev4MsgRsp) ProtoMessage()    {}
func (*SLRoutev4MsgRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25f7b0c2338e8af, []int{3}
}

func (m *SLRoutev4MsgRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SLRoutev4MsgRsp.Unmarshal(m, b)
}
func (m *SLRoutev4MsgRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SLRoutev4MsgRsp.Marshal(b, m, deterministic)
}
func (m *SLRoutev4MsgRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SLRoutev4MsgRsp.Merge(m, src)
}
func (m *SLRoutev4MsgRsp) XXX_Size() int {
	return xxx_messageInfo_SLRoutev4MsgRsp.Size(m)
}
func (m *SLRoutev4MsgRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SLRoutev4MsgRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SLRoutev4MsgRsp proto.InternalMessageInfo

func (m *SLRoutev4MsgRsp) GetCorrelator() uint64 {
	if m != nil {
		return m.Correlator
	}
	return 0
}

func (m *SLRoutev4MsgRsp) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SLRoutev4MsgRsp) GetStatusSummary() *SLErrorStatus {
	if m != nil {
		return m.StatusSummary
	}
	return nil
}

func (m *SLRoutev4MsgRsp) GetResults() []*SLRoutev4Res {
	if m != nil {
		return m.Results
	}
	return nil
}

// Used to retrieve route attributes
type SLRoutev4GetMsg struct {
	// Correlator. This can be used to correlate stream replies with requests.
	// The Server simply reflects this field back in the reply.
	Correlator uint64 `protobuf:"varint,1,opt,name=Correlator,proto3" json:"Correlator,omitempty"`
	// VRF name.
	// If the Key is not specified, then request up to the first
	// 'EntriesCount' entries.
	VrfName string `protobuf:"bytes,2,opt,name=VrfName,proto3" json:"VrfName,omitempty"`
	// IPv4 Prefix
	Prefix uint32 `protobuf:"varint,3,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	// IPv4 prefix length, [0-32]
	PrefixLen uint32 `protobuf:"varint,4,opt,name=PrefixLen,proto3" json:"PrefixLen,omitempty"`
	// Number of entries requested
	EntriesCount uint32 `protobuf:"varint,5,opt,name=EntriesCount,proto3" json:"EntriesCount,omitempty"`
	// if GetNext is FALSE:
	//     request up to 'EntriesCount' entries starting from the key
	// If GetNext is TRUE, or if the key exact match is not found:
	//     request up to 'EntriesCount' entries starting from the key's next
	GetNext              bool     `protobuf:"varint,6,opt,name=GetNext,proto3" json:"GetNext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SLRoutev4GetMsg) Reset()         { *m = SLRoutev4GetMsg{} }
func (m *SLRoutev4GetMsg) String() string { return proto.CompactTextString(m) }
func (*SLRoutev4GetMsg) ProtoMessage()    {}
func (*SLRoutev4GetMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25f7b0c2338e8af, []int{4}
}

func (m *SLRoutev4GetMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SLRoutev4GetMsg.Unmarshal(m, b)
}
func (m *SLRoutev4GetMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SLRoutev4GetMsg.Marshal(b, m, deterministic)
}
func (m *SLRoutev4GetMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SLRoutev4GetMsg.Merge(m, src)
}
func (m *SLRoutev4GetMsg) XXX_Size() int {
	return xxx_messageInfo_SLRoutev4GetMsg.Size(m)
}
func (m *SLRoutev4GetMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SLRoutev4GetMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SLRoutev4GetMsg proto.InternalMessageInfo

func (m *SLRoutev4GetMsg) GetCorrelator() uint64 {
	if m != nil {
		return m.Correlator
	}
	return 0
}

func (m *SLRoutev4GetMsg) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SLRoutev4GetMsg) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *SLRoutev4GetMsg) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

func (m *SLRoutev4GetMsg) GetEntriesCount() uint32 {
	if m != nil {
		return m.EntriesCount
	}
	return 0
}

func (m *SLRoutev4GetMsg) GetGetNext() bool {
	if m != nil {
		return m.GetNext
	}
	return false
}

// Gt Route message response
type SLRoutev4GetMsgRsp struct {
	// Correlator. This can be used to correlate replies with requests.
	// The Server simply reflects this field back in the reply.
	Correlator uint64 `protobuf:"varint,1,opt,name=Correlator,proto3" json:"Correlator,omitempty"`
	// End Of File.
	// When set to True, it indicates that the server has returned M, where
	// M < N, of the original N requested Entries.
	Eof bool `protobuf:"varint,2,opt,name=Eof,proto3" json:"Eof,omitempty"`
	// VRF name.
	VrfName string `protobuf:"bytes,3,opt,name=VrfName,proto3" json:"VrfName,omitempty"`
	// Status of the Get operation
	ErrStatus *SLErrorStatus `protobuf:"bytes,4,opt,name=ErrStatus,proto3" json:"ErrStatus,omitempty"`
	// Returned entries as requested in the Get operation.
	// if ErrStatus is SL_SUCCESS, Entries contains the info requested
	Entries              []*SLRoutev4 `protobuf:"bytes,5,rep,name=Entries,proto3" json:"Entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SLRoutev4GetMsgRsp) Reset()         { *m = SLRoutev4GetMsgRsp{} }
func (m *SLRoutev4GetMsgRsp) String() string { return proto.CompactTextString(m) }
func (*SLRoutev4GetMsgRsp) ProtoMessage()    {}
func (*SLRoutev4GetMsgRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25f7b0c2338e8af, []int{5}
}

func (m *SLRoutev4GetMsgRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SLRoutev4GetMsgRsp.Unmarshal(m, b)
}
func (m *SLRoutev4GetMsgRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SLRoutev4GetMsgRsp.Marshal(b, m, deterministic)
}
func (m *SLRoutev4GetMsgRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SLRoutev4GetMsgRsp.Merge(m, src)
}
func (m *SLRoutev4GetMsgRsp) XXX_Size() int {
	return xxx_messageInfo_SLRoutev4GetMsgRsp.Size(m)
}
func (m *SLRoutev4GetMsgRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SLRoutev4GetMsgRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SLRoutev4GetMsgRsp proto.InternalMessageInfo

func (m *SLRoutev4GetMsgRsp) GetCorrelator() uint64 {
	if m != nil {
		return m.Correlator
	}
	return 0
}

func (m *SLRoutev4GetMsgRsp) GetEof() bool {
	if m != nil {
		return m.Eof
	}
	return false
}

func (m *SLRoutev4GetMsgRsp) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SLRoutev4GetMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLRoutev4GetMsgRsp) GetEntries() []*SLRoutev4 {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*SLRoutev4)(nil), "service_layer.SLRoutev4")
	proto.RegisterType((*SLRoutev4Msg)(nil), "service_layer.SLRoutev4Msg")
	proto.RegisterType((*SLRoutev4Res)(nil), "service_layer.SLRoutev4Res")
	proto.RegisterType((*SLRoutev4MsgRsp)(nil), "service_layer.SLRoutev4MsgRsp")
	proto.RegisterType((*SLRoutev4GetMsg)(nil), "service_layer.SLRoutev4GetMsg")
	proto.RegisterType((*SLRoutev4GetMsgRsp)(nil), "service_layer.SLRoutev4GetMsgRsp")
}

func init() { proto.RegisterFile("sl_route_ipv4.proto", fileDescriptor_c25f7b0c2338e8af) }

var fileDescriptor_c25f7b0c2338e8af = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x8f, 0x12, 0x41,
	0x10, 0xcd, 0x2c, 0x2c, 0x1f, 0xc5, 0xe2, 0x47, 0xaf, 0x1f, 0x23, 0xae, 0x1b, 0x1c, 0x13, 0x25,
	0x26, 0x92, 0x0d, 0xa2, 0x07, 0x0f, 0x1e, 0x24, 0x48, 0x62, 0x58, 0x58, 0x1b, 0xdd, 0xc4, 0x78,
	0x20, 0x03, 0x69, 0x10, 0x33, 0x43, 0x4f, 0xba, 0x1b, 0xb2, 0xdc, 0xfc, 0x3b, 0xfe, 0x02, 0xe3,
	0xc9, 0x9b, 0x67, 0x7f, 0x92, 0xe9, 0x9e, 0x19, 0xe8, 0x01, 0x66, 0x16, 0xdd, 0x1b, 0x53, 0xf5,
	0xea, 0xd5, 0x7b, 0x55, 0x35, 0x03, 0x1c, 0x72, 0xa7, 0xcf, 0xe8, 0x4c, 0x90, 0xfe, 0xc4, 0x9b,
	0xd7, 0xab, 0x1e, 0xa3, 0x82, 0xa2, 0x22, 0x27, 0x6c, 0x3e, 0x19, 0x92, 0xbe, 0x63, 0x2f, 0x08,
	0x2b, 0xdd, 0xe6, 0x4e, 0x7f, 0x48, 0x5d, 0x97, 0x4e, 0xfb, 0x62, 0xe1, 0x11, 0xee, 0xa3, 0x54,
	0xd8, 0x2f, 0xf5, 0x93, 0x7e, 0xd8, 0xfa, 0x69, 0x40, 0xbe, 0xd7, 0xc6, 0x32, 0x31, 0xaf, 0xa3,
	0x3b, 0x90, 0x39, 0x63, 0x64, 0x34, 0xb9, 0x30, 0x8d, 0xb2, 0x51, 0x29, 0xe2, 0xe0, 0x09, 0x1d,
	0x41, 0xde, 0xff, 0xd5, 0x26, 0x53, 0x73, 0x4f, 0xa5, 0x56, 0x01, 0xf4, 0x1a, 0x0a, 0x8a, 0xa0,
	0xa1, 0x88, 0xcd, 0x54, 0xd9, 0xa8, 0x14, 0x6a, 0x47, 0xd5, 0x88, 0xac, 0x6a, 0xd0, 0xc4, 0xc7,
	0x60, 0xbd, 0x00, 0xbd, 0x84, 0xdc, 0x99, 0x2d, 0xbe, 0xb4, 0x27, 0x5c, 0x98, 0xe9, 0x72, 0xaa,
	0x52, 0xa8, 0x95, 0xb6, 0x17, 0x4b, 0x14, 0x5e, 0x62, 0xad, 0xef, 0x06, 0x1c, 0x2c, 0xb5, 0x9f,
	0xf2, 0x31, 0x7a, 0x06, 0xe9, 0xae, 0x47, 0x98, 0x12, 0x7f, 0xad, 0x76, 0x6f, 0x83, 0xa4, 0x3b,
	0xf8, 0x4a, 0x86, 0xa2, 0xeb, 0x61, 0x05, 0x43, 0xc7, 0x00, 0x0d, 0xca, 0x18, 0x71, 0x6c, 0x41,
	0x99, 0xb2, 0x95, 0xc6, 0x5a, 0x04, 0x99, 0x90, 0x3d, 0x67, 0xa3, 0x8e, 0xed, 0x12, 0xe5, 0x29,
	0x8f, 0xc3, 0x47, 0x74, 0x02, 0x19, 0xd5, 0x96, 0x07, 0x7a, 0xcd, 0xed, 0x7a, 0xe7, 0x75, 0x1c,
	0xe0, 0xac, 0x6f, 0xba, 0x56, 0x4c, 0x38, 0x7a, 0x05, 0xf9, 0x26, 0x63, 0x3d, 0x61, 0x8b, 0x19,
	0x57, 0x82, 0xb7, 0x8d, 0xac, 0xc9, 0x18, 0x0d, 0x30, 0x78, 0x05, 0xd7, 0xd6, 0xb4, 0x17, 0xbf,
	0xa6, 0xd4, 0xda, 0x9a, 0xac, 0xdf, 0x06, 0x5c, 0xd7, 0xc7, 0x85, 0xb9, 0xb7, 0x36, 0x02, 0x23,
	0x69, 0x04, 0x7b, 0xd1, 0x11, 0xbc, 0x81, 0xa2, 0xaf, 0xa6, 0x37, 0x73, 0x5d, 0x9b, 0x2d, 0x62,
	0xd7, 0xae, 0x7b, 0x88, 0x96, 0xa0, 0x17, 0x90, 0xc5, 0x84, 0xcf, 0x1c, 0x11, 0xce, 0xf1, 0x7e,
	0xec, 0x1c, 0x09, 0xc7, 0x21, 0xd6, 0xfa, 0xa5, 0x1b, 0x69, 0x11, 0x21, 0x57, 0xff, 0xff, 0x46,
	0x56, 0xc3, 0x4c, 0xc5, 0x0f, 0x33, 0xbd, 0x7e, 0xf3, 0x16, 0x1c, 0x34, 0xa7, 0x82, 0x4d, 0x08,
	0x6f, 0xd0, 0xd9, 0x54, 0x98, 0xfb, 0x0a, 0x10, 0x89, 0xc9, 0x9e, 0x2d, 0x22, 0x3a, 0xe4, 0x42,
	0x98, 0x99, 0xb2, 0x51, 0xc9, 0xe1, 0xf0, 0xd1, 0xfa, 0x63, 0x00, 0x5a, 0x73, 0xb0, 0xcb, 0x36,
	0x6e, 0x40, 0xaa, 0x49, 0x47, 0xca, 0x40, 0x0e, 0xcb, 0x9f, 0x09, 0x27, 0x1a, 0xb9, 0xaf, 0xf4,
	0xbf, 0xdd, 0x57, 0x0d, 0xb2, 0x81, 0x11, 0x73, 0xff, 0x92, 0xfb, 0x0e, 0x81, 0xb5, 0x1f, 0x19,
	0x28, 0x2e, 0xc3, 0xea, 0xf5, 0x1a, 0xc0, 0xe1, 0xca, 0xa3, 0x43, 0x07, 0xb6, 0xc3, 0x5b, 0x44,
	0xa0, 0x47, 0xdb, 0xb9, 0x56, 0x88, 0x53, 0x3e, 0x2e, 0x3d, 0xde, 0x01, 0x24, 0x27, 0x36, 0x85,
	0xbb, 0x6b, 0x3d, 0xa4, 0x05, 0xd5, 0xe7, 0x49, 0x12, 0x45, 0x88, 0x92, 0xbd, 0x9e, 0xee, 0x08,
	0x94, 0xfd, 0x3a, 0x70, 0x73, 0xd9, 0xef, 0x9c, 0x8d, 0x30, 0x19, 0x77, 0x3d, 0xb4, 0xf9, 0xb5,
	0xf2, 0x53, 0x92, 0xfc, 0x41, 0x7c, 0x4e, 0xf2, 0x7d, 0xd4, 0xee, 0xc0, 0x4f, 0x48, 0xe9, 0x71,
	0x45, 0x81, 0xe0, 0x72, 0x62, 0x5a, 0xd2, 0x7e, 0x82, 0x5b, 0x3a, 0x6d, 0x8b, 0x08, 0xe5, 0xe3,
	0x32, 0x62, 0x6b, 0x33, 0x8d, 0xdf, 0x86, 0xc5, 0x01, 0xf5, 0x3b, 0x28, 0x68, 0x6b, 0x46, 0xb1,
	0x6f, 0xac, 0xe4, 0x3b, 0x4e, 0x48, 0x4a, 0xae, 0xf7, 0xda, 0x37, 0x51, 0xfa, 0x8e, 0xc5, 0x07,
	0xfa, 0x1e, 0x26, 0xe7, 0x25, 0xe5, 0x07, 0x6d, 0x41, 0x5d, 0xaf, 0x27, 0x18, 0xb1, 0xdd, 0x2b,
	0x89, 0xac, 0x18, 0x27, 0x06, 0xfa, 0x1c, 0x7d, 0x5d, 0x03, 0xda, 0xab, 0xcb, 0x95, 0xe4, 0x83,
	0x8c, 0xfa, 0x27, 0x7e, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x6f, 0xa0, 0x35, 0xdd, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SLRoutev4OperClient is the client API for SLRoutev4Oper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SLRoutev4OperClient interface {
	// Used to retrieve Global Route information
	SLRoutev4GlobalsGet(ctx context.Context, in *SLRouteGlobalsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalsGetMsgRsp, error)
	// Used to retrieve Global Route Stats
	SLRoutev4GlobalStatsGet(ctx context.Context, in *SLRouteGlobalStatsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalStatsGetMsgRsp, error)
	// SLVrfRegMsg.Oper = SL_REGOP_REGISTER:
	//     VRF registration: Sends a list of VRF registration messages
	//     and expects a list of registration responses.
	//     A client Must Register a VRF BEFORE routes can be added/modified in
	//    the associated VRF.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_UNREGISTER:
	//     VRF Un-registeration: Sends a list of VRF un-registration messages
	//     and expects a list of un-registration responses.
	//     This can be used to convey that the client is no longer interested
	//     in this VRF. All previously installed routes would be lost.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_EOF:
	//     VRF End Of File message.
	//     After Registration, the client is expected to send an EOF
	//     message to convey the end of replay of the client's known objects.
	//     This is especially useful under certain restart scenarios when the
	//     client and the server are trying to synchronize their Routes.
	SLRoutev4VrfRegOp(ctx context.Context, in *SLVrfRegMsg, opts ...grpc.CallOption) (*SLVrfRegMsgRsp, error)
	// VRF get. Used to retrieve VRF attributes from the server.
	SLRoutev4VrfRegGet(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVrfRegGetMsgRsp, error)
	// Used to retrieve VRF Stats from the server.
	SLRoutev4VrfGetStats(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVRFGetStatsMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//     Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//     Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//     Route delete. The route path is not necessary to delete the route.
	SLRoutev4Op(ctx context.Context, in *SLRoutev4Msg, opts ...grpc.CallOption) (*SLRoutev4MsgRsp, error)
	// Retrieves route attributes.
	SLRoutev4Get(ctx context.Context, in *SLRoutev4GetMsg, opts ...grpc.CallOption) (*SLRoutev4GetMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//     Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//     Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//     Route delete. The route path is not necessary to delete the route.
	SLRoutev4OpStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4OpStreamClient, error)
	// Retrieves route attributes.
	SLRoutev4GetStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4GetStreamClient, error)
}

type sLRoutev4OperClient struct {
	cc *grpc.ClientConn
}

func NewSLRoutev4OperClient(cc *grpc.ClientConn) SLRoutev4OperClient {
	return &sLRoutev4OperClient{cc}
}

func (c *sLRoutev4OperClient) SLRoutev4GlobalsGet(ctx context.Context, in *SLRouteGlobalsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalsGetMsgRsp, error) {
	out := new(SLRouteGlobalsGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4GlobalsGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4GlobalStatsGet(ctx context.Context, in *SLRouteGlobalStatsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalStatsGetMsgRsp, error) {
	out := new(SLRouteGlobalStatsGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4GlobalStatsGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4VrfRegOp(ctx context.Context, in *SLVrfRegMsg, opts ...grpc.CallOption) (*SLVrfRegMsgRsp, error) {
	out := new(SLVrfRegMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4VrfRegGet(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVrfRegGetMsgRsp, error) {
	out := new(SLVrfRegGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4VrfGetStats(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVRFGetStatsMsgRsp, error) {
	out := new(SLVRFGetStatsMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4VrfGetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4Op(ctx context.Context, in *SLRoutev4Msg, opts ...grpc.CallOption) (*SLRoutev4MsgRsp, error) {
	out := new(SLRoutev4MsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4Op", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4Get(ctx context.Context, in *SLRoutev4GetMsg, opts ...grpc.CallOption) (*SLRoutev4GetMsgRsp, error) {
	out := new(SLRoutev4GetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4OpStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4OpStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SLRoutev4Oper_serviceDesc.Streams[0], "/service_layer.SLRoutev4Oper/SLRoutev4OpStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLRoutev4OperSLRoutev4OpStreamClient{stream}
	return x, nil
}

type SLRoutev4Oper_SLRoutev4OpStreamClient interface {
	Send(*SLRoutev4Msg) error
	Recv() (*SLRoutev4MsgRsp, error)
	grpc.ClientStream
}

type sLRoutev4OperSLRoutev4OpStreamClient struct {
	grpc.ClientStream
}

func (x *sLRoutev4OperSLRoutev4OpStreamClient) Send(m *SLRoutev4Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4OpStreamClient) Recv() (*SLRoutev4MsgRsp, error) {
	m := new(SLRoutev4MsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLRoutev4OperClient) SLRoutev4GetStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SLRoutev4Oper_serviceDesc.Streams[1], "/service_layer.SLRoutev4Oper/SLRoutev4GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLRoutev4OperSLRoutev4GetStreamClient{stream}
	return x, nil
}

type SLRoutev4Oper_SLRoutev4GetStreamClient interface {
	Send(*SLRoutev4GetMsg) error
	Recv() (*SLRoutev4GetMsgRsp, error)
	grpc.ClientStream
}

type sLRoutev4OperSLRoutev4GetStreamClient struct {
	grpc.ClientStream
}

func (x *sLRoutev4OperSLRoutev4GetStreamClient) Send(m *SLRoutev4GetMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4GetStreamClient) Recv() (*SLRoutev4GetMsgRsp, error) {
	m := new(SLRoutev4GetMsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SLRoutev4OperServer is the server API for SLRoutev4Oper service.
type SLRoutev4OperServer interface {
	// Used to retrieve Global Route information
	SLRoutev4GlobalsGet(context.Context, *SLRouteGlobalsGetMsg) (*SLRouteGlobalsGetMsgRsp, error)
	// Used to retrieve Global Route Stats
	SLRoutev4GlobalStatsGet(context.Context, *SLRouteGlobalStatsGetMsg) (*SLRouteGlobalStatsGetMsgRsp, error)
	// SLVrfRegMsg.Oper = SL_REGOP_REGISTER:
	//     VRF registration: Sends a list of VRF registration messages
	//     and expects a list of registration responses.
	//     A client Must Register a VRF BEFORE routes can be added/modified in
	//    the associated VRF.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_UNREGISTER:
	//     VRF Un-registeration: Sends a list of VRF un-registration messages
	//     and expects a list of un-registration responses.
	//     This can be used to convey that the client is no longer interested
	//     in this VRF. All previously installed routes would be lost.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_EOF:
	//     VRF End Of File message.
	//     After Registration, the client is expected to send an EOF
	//     message to convey the end of replay of the client's known objects.
	//     This is especially useful under certain restart scenarios when the
	//     client and the server are trying to synchronize their Routes.
	SLRoutev4VrfRegOp(context.Context, *SLVrfRegMsg) (*SLVrfRegMsgRsp, error)
	// VRF get. Used to retrieve VRF attributes from the server.
	SLRoutev4VrfRegGet(context.Context, *SLVrfRegGetMsg) (*SLVrfRegGetMsgRsp, error)
	// Used to retrieve VRF Stats from the server.
	SLRoutev4VrfGetStats(context.Context, *SLVrfRegGetMsg) (*SLVRFGetStatsMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//     Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//     Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//     Route delete. The route path is not necessary to delete the route.
	SLRoutev4Op(context.Context, *SLRoutev4Msg) (*SLRoutev4MsgRsp, error)
	// Retrieves route attributes.
	SLRoutev4Get(context.Context, *SLRoutev4GetMsg) (*SLRoutev4GetMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//     Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//     Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//     Route delete. The route path is not necessary to delete the route.
	SLRoutev4OpStream(SLRoutev4Oper_SLRoutev4OpStreamServer) error
	// Retrieves route attributes.
	SLRoutev4GetStream(SLRoutev4Oper_SLRoutev4GetStreamServer) error
}

func RegisterSLRoutev4OperServer(s *grpc.Server, srv SLRoutev4OperServer) {
	s.RegisterService(&_SLRoutev4Oper_serviceDesc, srv)
}

func _SLRoutev4Oper_SLRoutev4GlobalsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRouteGlobalsGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4GlobalsGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalsGet(ctx, req.(*SLRouteGlobalsGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4GlobalStatsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRouteGlobalStatsGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalStatsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4GlobalStatsGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalStatsGet(ctx, req.(*SLRouteGlobalStatsGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4VrfRegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLVrfRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegOp(ctx, req.(*SLVrfRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4VrfRegGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLVrfRegGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegGet(ctx, req.(*SLVrfRegGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4VrfGetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLVrfRegGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfGetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4VrfGetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfGetStats(ctx, req.(*SLVrfRegGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4Op_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRoutev4Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4Op(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4Op",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4Op(ctx, req.(*SLRoutev4Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRoutev4GetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4Get(ctx, req.(*SLRoutev4GetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4OpStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLRoutev4OperServer).SLRoutev4OpStream(&sLRoutev4OperSLRoutev4OpStreamServer{stream})
}

type SLRoutev4Oper_SLRoutev4OpStreamServer interface {
	Send(*SLRoutev4MsgRsp) error
	Recv() (*SLRoutev4Msg, error)
	grpc.ServerStream
}

type sLRoutev4OperSLRoutev4OpStreamServer struct {
	grpc.ServerStream
}

func (x *sLRoutev4OperSLRoutev4OpStreamServer) Send(m *SLRoutev4MsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4OpStreamServer) Recv() (*SLRoutev4Msg, error) {
	m := new(SLRoutev4Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SLRoutev4Oper_SLRoutev4GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLRoutev4OperServer).SLRoutev4GetStream(&sLRoutev4OperSLRoutev4GetStreamServer{stream})
}

type SLRoutev4Oper_SLRoutev4GetStreamServer interface {
	Send(*SLRoutev4GetMsgRsp) error
	Recv() (*SLRoutev4GetMsg, error)
	grpc.ServerStream
}

type sLRoutev4OperSLRoutev4GetStreamServer struct {
	grpc.ServerStream
}

func (x *sLRoutev4OperSLRoutev4GetStreamServer) Send(m *SLRoutev4GetMsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4GetStreamServer) Recv() (*SLRoutev4GetMsg, error) {
	m := new(SLRoutev4GetMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SLRoutev4Oper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service_layer.SLRoutev4Oper",
	HandlerType: (*SLRoutev4OperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SLRoutev4GlobalsGet",
			Handler:    _SLRoutev4Oper_SLRoutev4GlobalsGet_Handler,
		},
		{
			MethodName: "SLRoutev4GlobalStatsGet",
			Handler:    _SLRoutev4Oper_SLRoutev4GlobalStatsGet_Handler,
		},
		{
			MethodName: "SLRoutev4VrfRegOp",
			Handler:    _SLRoutev4Oper_SLRoutev4VrfRegOp_Handler,
		},
		{
			MethodName: "SLRoutev4VrfRegGet",
			Handler:    _SLRoutev4Oper_SLRoutev4VrfRegGet_Handler,
		},
		{
			MethodName: "SLRoutev4VrfGetStats",
			Handler:    _SLRoutev4Oper_SLRoutev4VrfGetStats_Handler,
		},
		{
			MethodName: "SLRoutev4Op",
			Handler:    _SLRoutev4Oper_SLRoutev4Op_Handler,
		},
		{
			MethodName: "SLRoutev4Get",
			Handler:    _SLRoutev4Oper_SLRoutev4Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SLRoutev4OpStream",
			Handler:       _SLRoutev4Oper_SLRoutev4OpStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SLRoutev4GetStream",
			Handler:       _SLRoutev4Oper_SLRoutev4GetStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sl_route_ipv4.proto",
}
