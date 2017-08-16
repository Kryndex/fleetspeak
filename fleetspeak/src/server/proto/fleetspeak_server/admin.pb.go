// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/proto/fleetspeak_server/admin.proto

/*
Package fleetspeak_server is a generated protocol buffer package.

It is generated from these files:
	fleetspeak/src/server/proto/fleetspeak_server/admin.proto
	fleetspeak/src/server/proto/fleetspeak_server/broadcasts.proto
	fleetspeak/src/server/proto/fleetspeak_server/server.proto
	fleetspeak/src/server/proto/fleetspeak_server/services.proto

It has these top-level messages:
	CreateBroadcastRequest
	ListActiveBroadcastsRequest
	ListActiveBroadcastsResponse
	ListClientsResponse
	Client
	GetMessageStatusRequest
	GetMessageStatusResponse
	StoreFileRequest
	Broadcast
	ServerConfig
	ServiceConfig
*/
package fleetspeak_server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import fleetspeak "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"

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

type CreateBroadcastRequest struct {
	Broadcast *Broadcast `protobuf:"bytes,1,opt,name=broadcast" json:"broadcast,omitempty"`
	Limit     uint64     `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *CreateBroadcastRequest) Reset()                    { *m = CreateBroadcastRequest{} }
func (m *CreateBroadcastRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBroadcastRequest) ProtoMessage()               {}
func (*CreateBroadcastRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateBroadcastRequest) GetBroadcast() *Broadcast {
	if m != nil {
		return m.Broadcast
	}
	return nil
}

func (m *CreateBroadcastRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListActiveBroadcastsRequest struct {
	// If set, only return broadcasts with the given service_name.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *ListActiveBroadcastsRequest) Reset()                    { *m = ListActiveBroadcastsRequest{} }
func (m *ListActiveBroadcastsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListActiveBroadcastsRequest) ProtoMessage()               {}
func (*ListActiveBroadcastsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListActiveBroadcastsRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type ListActiveBroadcastsResponse struct {
	Broadcasts []*Broadcast `protobuf:"bytes,1,rep,name=broadcasts" json:"broadcasts,omitempty"`
}

func (m *ListActiveBroadcastsResponse) Reset()                    { *m = ListActiveBroadcastsResponse{} }
func (m *ListActiveBroadcastsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListActiveBroadcastsResponse) ProtoMessage()               {}
func (*ListActiveBroadcastsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListActiveBroadcastsResponse) GetBroadcasts() []*Broadcast {
	if m != nil {
		return m.Broadcasts
	}
	return nil
}

type ListClientsResponse struct {
	Clients []*Client `protobuf:"bytes,1,rep,name=clients" json:"clients,omitempty"`
}

func (m *ListClientsResponse) Reset()                    { *m = ListClientsResponse{} }
func (m *ListClientsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListClientsResponse) ProtoMessage()               {}
func (*ListClientsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListClientsResponse) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

type Client struct {
	ClientId        []byte                     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Labels          []*fleetspeak.Label        `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	LastContactTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=last_contact_time,json=lastContactTime" json:"last_contact_time,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Client) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *Client) GetLabels() []*fleetspeak.Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Client) GetLastContactTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastContactTime
	}
	return nil
}

type GetMessageStatusRequest struct {
	MessageId []byte `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (m *GetMessageStatusRequest) Reset()                    { *m = GetMessageStatusRequest{} }
func (m *GetMessageStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMessageStatusRequest) ProtoMessage()               {}
func (*GetMessageStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetMessageStatusRequest) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

type GetMessageStatusResponse struct {
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	Result       *fleetspeak.MessageResult  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *GetMessageStatusResponse) Reset()                    { *m = GetMessageStatusResponse{} }
func (m *GetMessageStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMessageStatusResponse) ProtoMessage()               {}
func (*GetMessageStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetMessageStatusResponse) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *GetMessageStatusResponse) GetResult() *fleetspeak.MessageResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type StoreFileRequest struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	FileName    string `protobuf:"bytes,2,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	Data        []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *StoreFileRequest) Reset()                    { *m = StoreFileRequest{} }
func (m *StoreFileRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreFileRequest) ProtoMessage()               {}
func (*StoreFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StoreFileRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *StoreFileRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *StoreFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateBroadcastRequest)(nil), "fleetspeak.server.CreateBroadcastRequest")
	proto.RegisterType((*ListActiveBroadcastsRequest)(nil), "fleetspeak.server.ListActiveBroadcastsRequest")
	proto.RegisterType((*ListActiveBroadcastsResponse)(nil), "fleetspeak.server.ListActiveBroadcastsResponse")
	proto.RegisterType((*ListClientsResponse)(nil), "fleetspeak.server.ListClientsResponse")
	proto.RegisterType((*Client)(nil), "fleetspeak.server.Client")
	proto.RegisterType((*GetMessageStatusRequest)(nil), "fleetspeak.server.GetMessageStatusRequest")
	proto.RegisterType((*GetMessageStatusResponse)(nil), "fleetspeak.server.GetMessageStatusResponse")
	proto.RegisterType((*StoreFileRequest)(nil), "fleetspeak.server.StoreFileRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Admin service

type AdminClient interface {
	// CreateBroadcast creates a FS broadcast, potentially sending a message to
	// many machines in the fleet.
	CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// ListActiveBroadcasts lists the currently active FS broadcasts.
	ListActiveBroadcasts(ctx context.Context, in *ListActiveBroadcastsRequest, opts ...grpc.CallOption) (*ListActiveBroadcastsResponse, error)
	// ListClients lists the currently active FS clients.
	ListClients(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*ListClientsResponse, error)
	// GetMessageStatus retrieves the current status of a message.
	GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*GetMessageStatusResponse, error)
	// InsertMessage inserts a message into the Fleetspeak system to be processed
	// by the server or delivered to a client.
	// TODO: Have this method return the message that is written to the
	// datastore (or at least the id).
	InsertMessage(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// StoreFile inserts a file into the Fleetspeak system.
	StoreFile(ctx context.Context, in *StoreFileRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// KeepAlive does as little as possible.
	KeepAlive(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/CreateBroadcast", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListActiveBroadcasts(ctx context.Context, in *ListActiveBroadcastsRequest, opts ...grpc.CallOption) (*ListActiveBroadcastsResponse, error) {
	out := new(ListActiveBroadcastsResponse)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/ListActiveBroadcasts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListClients(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*ListClientsResponse, error) {
	out := new(ListClientsResponse)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/ListClients", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*GetMessageStatusResponse, error) {
	out := new(GetMessageStatusResponse)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/GetMessageStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InsertMessage(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/InsertMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StoreFile(ctx context.Context, in *StoreFileRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/StoreFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) KeepAlive(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := grpc.Invoke(ctx, "/fleetspeak.server.Admin/KeepAlive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	// CreateBroadcast creates a FS broadcast, potentially sending a message to
	// many machines in the fleet.
	CreateBroadcast(context.Context, *CreateBroadcastRequest) (*fleetspeak.EmptyMessage, error)
	// ListActiveBroadcasts lists the currently active FS broadcasts.
	ListActiveBroadcasts(context.Context, *ListActiveBroadcastsRequest) (*ListActiveBroadcastsResponse, error)
	// ListClients lists the currently active FS clients.
	ListClients(context.Context, *fleetspeak.EmptyMessage) (*ListClientsResponse, error)
	// GetMessageStatus retrieves the current status of a message.
	GetMessageStatus(context.Context, *GetMessageStatusRequest) (*GetMessageStatusResponse, error)
	// InsertMessage inserts a message into the Fleetspeak system to be processed
	// by the server or delivered to a client.
	// TODO: Have this method return the message that is written to the
	// datastore (or at least the id).
	InsertMessage(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error)
	// StoreFile inserts a file into the Fleetspeak system.
	StoreFile(context.Context, *StoreFileRequest) (*fleetspeak.EmptyMessage, error)
	// KeepAlive does as little as possible.
	KeepAlive(context.Context, *fleetspeak.EmptyMessage) (*fleetspeak.EmptyMessage, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_CreateBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/CreateBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateBroadcast(ctx, req.(*CreateBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListActiveBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveBroadcastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListActiveBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListActiveBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListActiveBroadcasts(ctx, req.(*ListActiveBroadcastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListClients(ctx, req.(*fleetspeak.EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/GetMessageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetMessageStatus(ctx, req.(*GetMessageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InsertMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InsertMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/InsertMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InsertMessage(ctx, req.(*fleetspeak.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StoreFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StoreFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/StoreFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StoreFile(ctx, req.(*StoreFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).KeepAlive(ctx, req.(*fleetspeak.EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fleetspeak.server.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBroadcast",
			Handler:    _Admin_CreateBroadcast_Handler,
		},
		{
			MethodName: "ListActiveBroadcasts",
			Handler:    _Admin_ListActiveBroadcasts_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _Admin_ListClients_Handler,
		},
		{
			MethodName: "GetMessageStatus",
			Handler:    _Admin_GetMessageStatus_Handler,
		},
		{
			MethodName: "InsertMessage",
			Handler:    _Admin_InsertMessage_Handler,
		},
		{
			MethodName: "StoreFile",
			Handler:    _Admin_StoreFile_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Admin_KeepAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fleetspeak/src/server/proto/fleetspeak_server/admin.proto",
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/proto/fleetspeak_server/admin.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xfb, 0x93, 0xaf, 0x9e, 0xa6, 0x6a, 0xbb, 0xad, 0x3e, 0x8c, 0x5b, 0x44, 0x31, 0x12,
	0x6a, 0x41, 0xb2, 0x45, 0x73, 0x03, 0x08, 0x41, 0x43, 0x44, 0x51, 0x20, 0x20, 0xe1, 0x72, 0xc1,
	0x05, 0x52, 0xb4, 0xb1, 0x27, 0xd1, 0x82, 0xff, 0xf0, 0x6e, 0x22, 0xf5, 0x25, 0xb8, 0xe6, 0x35,
	0x79, 0x03, 0x64, 0xef, 0x3a, 0x8e, 0x1c, 0x27, 0x0d, 0x77, 0xf6, 0x99, 0x73, 0xce, 0xcc, 0xee,
	0x9e, 0x81, 0xe7, 0xc3, 0x00, 0x51, 0xf0, 0x04, 0xe9, 0x0f, 0x87, 0xa7, 0x9e, 0xc3, 0x31, 0x9d,
	0x60, 0xea, 0x24, 0x69, 0x2c, 0x62, 0xa7, 0xac, 0xf5, 0x15, 0x4e, 0xfd, 0x90, 0x45, 0x76, 0x5e,
	0x25, 0x07, 0x65, 0xd9, 0x96, 0x65, 0xf3, 0xfe, 0x28, 0x8e, 0x47, 0x01, 0x4a, 0xf9, 0x60, 0x3c,
	0x74, 0x04, 0x0b, 0x91, 0x0b, 0x1a, 0x26, 0x52, 0x63, 0xb6, 0x2a, 0xed, 0xbc, 0x38, 0x0c, 0xe3,
	0x68, 0xae, 0x9d, 0xc2, 0x95, 0xe8, 0xd5, 0xbf, 0xcd, 0x38, 0x48, 0x63, 0xea, 0x7b, 0x94, 0x0b,
	0x2e, 0xf5, 0xd6, 0x77, 0xf8, 0xbf, 0x93, 0x22, 0x15, 0xf8, 0xa6, 0xa8, 0xb8, 0xf8, 0x73, 0x8c,
	0x5c, 0x90, 0x17, 0xa0, 0x4f, 0xd9, 0x86, 0x76, 0xaa, 0x9d, 0xed, 0x5c, 0x9c, 0xd8, 0x73, 0xc7,
	0xb2, 0x4b, 0x5d, 0x49, 0x27, 0x47, 0xb0, 0x15, 0xb0, 0x90, 0x09, 0x63, 0xfd, 0x54, 0x3b, 0xdb,
	0x74, 0xe5, 0x8f, 0x75, 0x09, 0xc7, 0x3d, 0xc6, 0x45, 0xdb, 0x13, 0x6c, 0x52, 0xf6, 0xe3, 0x45,
	0xc3, 0x07, 0xd0, 0xcc, 0x3c, 0x99, 0x87, 0xfd, 0x88, 0x86, 0x98, 0xf7, 0xd4, 0xdd, 0x1d, 0x85,
	0x7d, 0xa2, 0x21, 0x5a, 0xdf, 0xe0, 0xa4, 0xde, 0x81, 0x27, 0x71, 0xc4, 0x91, 0xbc, 0x04, 0x28,
	0x4f, 0x68, 0x68, 0xa7, 0x1b, 0xb7, 0x0e, 0x3d, 0xc3, 0xb7, 0xde, 0xc3, 0x61, 0xe6, 0xde, 0x09,
	0x18, 0x46, 0x33, 0xa6, 0x2d, 0xf8, 0xcf, 0x93, 0x90, 0x72, 0xbc, 0x5b, 0xe3, 0x28, 0x45, 0x6e,
	0xc1, 0xb4, 0x7e, 0x6b, 0xd0, 0x90, 0x18, 0x39, 0x06, 0x5d, 0xa2, 0x7d, 0xe6, 0xe7, 0x87, 0x6a,
	0xba, 0xdb, 0x12, 0xe8, 0xfa, 0xe4, 0x1c, 0x1a, 0x01, 0x1d, 0x60, 0xc0, 0x8d, 0xf5, 0xdc, 0xfb,
	0x60, 0xd6, 0xbb, 0x97, 0x55, 0x5c, 0x45, 0x20, 0x57, 0x70, 0x10, 0x50, 0x2e, 0xfa, 0x5e, 0x1c,
	0x09, 0xea, 0x89, 0x7e, 0x96, 0x1f, 0x63, 0x23, 0x7f, 0x18, 0xd3, 0x96, 0xe1, 0xb2, 0x8b, 0x70,
	0xd9, 0x5f, 0x8a, 0x70, 0xb9, 0x7b, 0x99, 0xa8, 0x23, 0x35, 0x19, 0x6a, 0x3d, 0x83, 0x3b, 0xef,
	0x50, 0x7c, 0x44, 0xce, 0xe9, 0x08, 0xaf, 0x05, 0x15, 0xe3, 0xe9, 0x13, 0xdc, 0x03, 0x08, 0x25,
	0x5e, 0xce, 0xaa, 0x2b, 0xa4, 0xeb, 0x5b, 0xbf, 0x34, 0x30, 0xe6, 0xa5, 0xea, 0x9a, 0x5e, 0xc3,
	0xae, 0x97, 0x25, 0x89, 0xc5, 0x91, 0x1c, 0x4d, 0xbb, 0x75, 0xb4, 0x66, 0x21, 0xc8, 0x20, 0xf2,
	0x14, 0x1a, 0x29, 0xf2, 0x71, 0x20, 0x53, 0x53, 0xb9, 0x66, 0xd5, 0xd3, 0xcd, 0x09, 0xae, 0x22,
	0x5a, 0x43, 0xd8, 0xbf, 0x16, 0x71, 0x8a, 0x57, 0x2c, 0xc0, 0xd5, 0x63, 0x94, 0xbd, 0xc8, 0x90,
	0x05, 0xaa, 0xbe, 0x9e, 0xd7, 0xb7, 0x33, 0x20, 0x2f, 0x12, 0xd8, 0xf4, 0xa9, 0xa0, 0xf9, 0xcd,
	0x36, 0xdd, 0xfc, 0xfb, 0xe2, 0xcf, 0x26, 0x6c, 0xb5, 0xb3, 0xf5, 0x26, 0x5f, 0x61, 0xaf, 0xb2,
	0x2f, 0xe4, 0xbc, 0x2e, 0x0e, 0xb5, 0x3b, 0x65, 0x1a, 0xb3, 0xd4, 0xb7, 0x61, 0x22, 0x6e, 0xd4,
	0xb9, 0xac, 0x35, 0x72, 0x03, 0x47, 0x75, 0xd9, 0x26, 0x76, 0x8d, 0xfd, 0x92, 0x35, 0x32, 0x9d,
	0x95, 0xf9, 0xf2, 0xe1, 0xac, 0x35, 0xf2, 0x19, 0x76, 0x66, 0x82, 0x4f, 0x16, 0x4e, 0x69, 0x3e,
	0x5a, 0xe0, 0x5d, 0x59, 0x19, 0x6b, 0x8d, 0x84, 0xb0, 0x5f, 0x4d, 0x0a, 0x79, 0x5c, 0xa3, 0x5e,
	0x90, 0x44, 0xf3, 0xc9, 0x4a, 0xdc, 0x69, 0xbb, 0x4b, 0xd8, 0xed, 0x46, 0x1c, 0xd3, 0x82, 0x40,
	0x0e, 0x6b, 0xc2, 0xb3, 0xf4, 0xfa, 0x7b, 0xa0, 0x4f, 0xa3, 0x44, 0x1e, 0xd6, 0x74, 0xaf, 0x06,
	0x6d, 0xa9, 0x5b, 0x1b, 0xf4, 0x0f, 0x88, 0x49, 0x3b, 0x60, 0x13, 0x5c, 0x72, 0x9f, 0x4b, 0x2c,
	0x06, 0x8d, 0x7c, 0x61, 0x5a, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x89, 0xfb, 0x6c, 0xe3, 0x86,
	0x06, 0x00, 0x00,
}
