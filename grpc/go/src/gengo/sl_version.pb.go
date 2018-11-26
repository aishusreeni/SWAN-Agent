// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sl_version.proto

package service_layer

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Service Layer API version.
// This is used in the Global init message exchange to handshake client/server
// Version numbers.
type SLVersion int32

const (
	SLVersion_SL_VERSION_UNUSED SLVersion = 0
	SLVersion_SL_MAJOR_VERSION  SLVersion = 0
	SLVersion_SL_MINOR_VERSION  SLVersion = 0
	SLVersion_SL_SUB_VERSION    SLVersion = 1
)

var SLVersion_name = map[int32]string{
	0: "SL_VERSION_UNUSED",
	// Duplicate value: 0: "SL_MAJOR_VERSION",
	// Duplicate value: 0: "SL_MINOR_VERSION",
	1: "SL_SUB_VERSION",
}

var SLVersion_value = map[string]int32{
	"SL_VERSION_UNUSED": 0,
	"SL_MAJOR_VERSION":  0,
	"SL_MINOR_VERSION":  0,
	"SL_SUB_VERSION":    1,
}

func (x SLVersion) String() string {
	return proto.EnumName(SLVersion_name, int32(x))
}

func (SLVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df2f827a2739d5ca, []int{0}
}

func init() {
	proto.RegisterEnum("service_layer.SLVersion", SLVersion_name, SLVersion_value)
}

func init() { proto.RegisterFile("sl_version.proto", fileDescriptor_df2f827a2739d5ca) }

var fileDescriptor_df2f827a2739d5ca = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xce, 0x89, 0x2f,
	0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d, 0xcf, 0x49, 0xac, 0x4c, 0x2d, 0xd2, 0x4a, 0xe3, 0xe2, 0x0c,
	0xf6, 0x09, 0x83, 0xa8, 0x10, 0x12, 0xe5, 0x12, 0x0c, 0xf6, 0x89, 0x0f, 0x73, 0x0d, 0x0a, 0xf6,
	0xf4, 0xf7, 0x8b, 0x0f, 0xf5, 0x0b, 0x0d, 0x76, 0x75, 0x11, 0x60, 0x10, 0x12, 0xe1, 0x12, 0x08,
	0xf6, 0x89, 0xf7, 0x75, 0xf4, 0xf2, 0x0f, 0x82, 0x49, 0x22, 0x44, 0x3d, 0xfd, 0x50, 0x44, 0x85,
	0xb8, 0xf8, 0x82, 0x7d, 0xe2, 0x83, 0x43, 0x9d, 0xe0, 0x62, 0x8c, 0x52, 0x4c, 0x02, 0x8c, 0x49,
	0x6c, 0x60, 0xdb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xef, 0x8c, 0xa8, 0x91, 0x00,
	0x00, 0x00,
}