// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/event_feed/event_feed.proto

package event_feed // import "github.com/chef/automate/api/interservice/event_feed"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventFeedServiceClient is the client API for EventFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventFeedServiceClient interface {
}

type eventFeedServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventFeedServiceClient(cc *grpc.ClientConn) EventFeedServiceClient {
	return &eventFeedServiceClient{cc}
}

// EventFeedServiceServer is the server API for EventFeedService service.
type EventFeedServiceServer interface {
}

func RegisterEventFeedServiceServer(s *grpc.Server, srv EventFeedServiceServer) {
	s.RegisterService(&_EventFeedService_serviceDesc, srv)
}

var _EventFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.event_feed.api.EventFeedService",
	HandlerType: (*EventFeedServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/interservice/event_feed/event_feed.proto",
}

func init() {
	proto.RegisterFile("api/interservice/event_feed/event_feed.proto", fileDescriptor_event_feed_b1bddf94924b1ebf)
}

var fileDescriptor_event_feed_b1bddf94924b1ebf = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x2c, 0xc8, 0xd4,
	0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0x89, 0x4f, 0x4b, 0x4d, 0x4d, 0x41, 0x62, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0x29, 0x27, 0x67, 0xa4, 0xa6, 0xe9, 0x25, 0x96, 0x96, 0xe4, 0xe7, 0x26, 0x96, 0xa4, 0xea, 0xa5,
	0xe4, 0xe7, 0x26, 0x66, 0xe6, 0xe9, 0x21, 0x29, 0x4b, 0x2c, 0xc8, 0x34, 0x12, 0xe2, 0x12, 0x70,
	0x05, 0x89, 0xb8, 0xa5, 0xa6, 0xa6, 0x04, 0x43, 0x0c, 0x75, 0x32, 0x8b, 0x32, 0x49, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x99, 0xa2, 0x0f, 0x33, 0x45, 0x1f, 0x8f,
	0x0b, 0x92, 0xd8, 0xc0, 0xf6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff, 0x0d, 0x9a, 0x52,
	0xa7, 0x00, 0x00, 0x00,
}