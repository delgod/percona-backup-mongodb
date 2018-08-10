// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_30cec09fc0c5571e, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ClientStatus struct {
	ReplicaSetUUID       string   `protobuf:"bytes,1,opt,name=ReplicaSetUUID,proto3" json:"ReplicaSetUUID,omitempty"`
	ReplicaSetName       string   `protobuf:"bytes,2,opt,name=ReplicaSetName,proto3" json:"ReplicaSetName,omitempty"`
	ReplicaSetVersion    int64    `protobuf:"varint,3,opt,name=ReplicaSetVersion,proto3" json:"ReplicaSetVersion,omitempty"`
	LastOplogTime        int64    `protobuf:"varint,4,opt,name=LastOplogTime,proto3" json:"LastOplogTime,omitempty"`
	RunningDBBackup      bool     `protobuf:"varint,5,opt,name=RunningDBBackup,proto3" json:"RunningDBBackup,omitempty"`
	RunningOplogBackup   bool     `protobuf:"varint,6,opt,name=RunningOplogBackup,proto3" json:"RunningOplogBackup,omitempty"`
	Compression          string   `protobuf:"bytes,7,opt,name=Compression,proto3" json:"Compression,omitempty"`
	Encrypted            string   `protobuf:"bytes,8,opt,name=Encrypted,proto3" json:"Encrypted,omitempty"`
	Destination          string   `protobuf:"bytes,9,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Filename             string   `protobuf:"bytes,10,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Started              int64    `protobuf:"varint,11,opt,name=Started,proto3" json:"Started,omitempty"`
	Finished             int64    `protobuf:"varint,12,opt,name=Finished,proto3" json:"Finished,omitempty"`
	LastError            string   `protobuf:"bytes,13,opt,name=LastError,proto3" json:"LastError,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatus) Reset()         { *m = ClientStatus{} }
func (m *ClientStatus) String() string { return proto.CompactTextString(m) }
func (*ClientStatus) ProtoMessage()    {}
func (*ClientStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_30cec09fc0c5571e, []int{1}
}
func (m *ClientStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatus.Unmarshal(m, b)
}
func (m *ClientStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatus.Marshal(b, m, deterministic)
}
func (dst *ClientStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatus.Merge(dst, src)
}
func (m *ClientStatus) XXX_Size() int {
	return xxx_messageInfo_ClientStatus.Size(m)
}
func (m *ClientStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatus proto.InternalMessageInfo

func (m *ClientStatus) GetReplicaSetUUID() string {
	if m != nil {
		return m.ReplicaSetUUID
	}
	return ""
}

func (m *ClientStatus) GetReplicaSetName() string {
	if m != nil {
		return m.ReplicaSetName
	}
	return ""
}

func (m *ClientStatus) GetReplicaSetVersion() int64 {
	if m != nil {
		return m.ReplicaSetVersion
	}
	return 0
}

func (m *ClientStatus) GetLastOplogTime() int64 {
	if m != nil {
		return m.LastOplogTime
	}
	return 0
}

func (m *ClientStatus) GetRunningDBBackup() bool {
	if m != nil {
		return m.RunningDBBackup
	}
	return false
}

func (m *ClientStatus) GetRunningOplogBackup() bool {
	if m != nil {
		return m.RunningOplogBackup
	}
	return false
}

func (m *ClientStatus) GetCompression() string {
	if m != nil {
		return m.Compression
	}
	return ""
}

func (m *ClientStatus) GetEncrypted() string {
	if m != nil {
		return m.Encrypted
	}
	return ""
}

func (m *ClientStatus) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *ClientStatus) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ClientStatus) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *ClientStatus) GetFinished() int64 {
	if m != nil {
		return m.Finished
	}
	return 0
}

func (m *ClientStatus) GetLastError() string {
	if m != nil {
		return m.LastError
	}
	return ""
}

type Client struct {
	Version              int32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ClientID             string        `protobuf:"bytes,2,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	Status               *ClientStatus `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	LastSeen             int64         `protobuf:"varint,4,opt,name=LastSeen,proto3" json:"LastSeen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_30cec09fc0c5571e, []int{2}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (dst *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(dst, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Client) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Client) GetStatus() *ClientStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Client) GetLastSeen() int64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

type StartBackup struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackup) Reset()         { *m = StartBackup{} }
func (m *StartBackup) String() string { return proto.CompactTextString(m) }
func (*StartBackup) ProtoMessage()    {}
func (*StartBackup) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_30cec09fc0c5571e, []int{3}
}
func (m *StartBackup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackup.Unmarshal(m, b)
}
func (m *StartBackup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackup.Marshal(b, m, deterministic)
}
func (dst *StartBackup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackup.Merge(dst, src)
}
func (m *StartBackup) XXX_Size() int {
	return xxx_messageInfo_StartBackup.Size(m)
}
func (m *StartBackup) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackup.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackup proto.InternalMessageInfo

func (m *StartBackup) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*ClientStatus)(nil), "api.ClientStatus")
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*StartBackup)(nil), "api.StartBackup")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	GetClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Api_GetClientsClient, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Api_GetClientsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Api_serviceDesc.Streams[0], "/api.Api/GetClients", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiGetClientsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Api_GetClientsClient interface {
	Recv() (*Client, error)
	grpc.ClientStream
}

type apiGetClientsClient struct {
	grpc.ClientStream
}

func (x *apiGetClientsClient) Recv() (*Client, error) {
	m := new(Client)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	GetClients(*Empty, Api_GetClientsServer) error
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GetClients_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServer).GetClients(m, &apiGetClientsServer{stream})
}

type Api_GetClientsServer interface {
	Send(*Client) error
	grpc.ServerStream
}

type apiGetClientsServer struct {
	grpc.ServerStream
}

func (x *apiGetClientsServer) Send(m *Client) error {
	return x.ServerStream.SendMsg(m)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetClients",
			Handler:       _Api_GetClients_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_30cec09fc0c5571e) }

var fileDescriptor_api_30cec09fc0c5571e = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdb, 0x6e, 0xd4, 0x30,
	0x10, 0x25, 0x84, 0xbd, 0x64, 0xd2, 0x05, 0xd5, 0x4f, 0xa6, 0xe2, 0x21, 0x8a, 0x10, 0xa4, 0x52,
	0x15, 0x55, 0xe5, 0x0b, 0x48, 0xb3, 0xa0, 0x4a, 0x5c, 0xaa, 0x84, 0xf2, 0x6e, 0xd2, 0x51, 0xb0,
	0x48, 0x1c, 0xcb, 0xf6, 0x22, 0xfa, 0x01, 0xfc, 0x26, 0xdf, 0x82, 0xec, 0xb8, 0x69, 0xba, 0xd0,
	0xb7, 0x9c, 0x33, 0x67, 0xc6, 0x33, 0x73, 0x26, 0x10, 0x31, 0xc9, 0x73, 0xa9, 0x06, 0x33, 0x90,
	0x90, 0x49, 0x9e, 0xae, 0x60, 0xb1, 0xed, 0xa5, 0xb9, 0x49, 0xff, 0x84, 0x70, 0x70, 0xde, 0x71,
	0x14, 0xa6, 0x36, 0xcc, 0xec, 0x34, 0x79, 0x05, 0x4f, 0x2b, 0x94, 0x1d, 0x6f, 0x58, 0x8d, 0xe6,
	0xea, 0xea, 0xa2, 0xa4, 0x41, 0x12, 0x64, 0x51, 0xb5, 0xc7, 0xde, 0xd7, 0x7d, 0x62, 0x3d, 0xd2,
	0xc7, 0xfb, 0x3a, 0xcb, 0x92, 0x13, 0x38, 0xbc, 0x63, 0xbe, 0xa2, 0xd2, 0x7c, 0x10, 0x34, 0x4c,
	0x82, 0x2c, 0xac, 0xfe, 0x0d, 0x90, 0x97, 0xb0, 0xf9, 0xc0, 0xb4, 0xf9, 0x2c, 0xbb, 0xa1, 0xfd,
	0xc2, 0x7b, 0xa4, 0x4f, 0x9c, 0xf2, 0x3e, 0x49, 0x32, 0x78, 0x56, 0xed, 0x84, 0xe0, 0xa2, 0x2d,
	0x8b, 0x82, 0x35, 0x3f, 0x76, 0x92, 0x2e, 0x92, 0x20, 0x5b, 0x57, 0xfb, 0x34, 0xc9, 0x81, 0x78,
	0xca, 0x65, 0x7b, 0xf1, 0xd2, 0x89, 0xff, 0x13, 0x21, 0x09, 0xc4, 0xe7, 0x43, 0x2f, 0x15, 0x6a,
	0xd7, 0xe7, 0xca, 0x8d, 0x34, 0xa7, 0xc8, 0x0b, 0x88, 0xb6, 0xa2, 0x51, 0x37, 0xd2, 0xe0, 0x35,
	0x5d, 0xbb, 0xf8, 0x1d, 0x61, 0xf3, 0x4b, 0xd4, 0x86, 0x0b, 0x66, 0x6c, 0x7e, 0x34, 0xe6, 0xcf,
	0x28, 0x72, 0x04, 0xeb, 0x77, 0xbc, 0x43, 0x61, 0x37, 0x06, 0x2e, 0x3c, 0x61, 0x42, 0x61, 0x55,
	0x1b, 0xa6, 0x6c, 0xe5, 0xd8, 0xcd, 0x7d, 0x0b, 0xc7, 0x2c, 0xc1, 0xf5, 0x77, 0xbc, 0xa6, 0x07,
	0x2e, 0x34, 0x61, 0xdb, 0x91, 0x5d, 0xcf, 0x56, 0xa9, 0x41, 0xd1, 0xcd, 0xd8, 0xd1, 0x44, 0xa4,
	0xbf, 0x03, 0x58, 0x8e, 0x06, 0xdb, 0xf2, 0x3f, 0xbd, 0x01, 0xd6, 0xd3, 0x45, 0x75, 0x0b, 0x6d,
	0xf9, 0x51, 0x73, 0x51, 0x7a, 0x1b, 0x27, 0x4c, 0x8e, 0x61, 0x39, 0x9e, 0x86, 0x73, 0x2d, 0x3e,
	0x3b, 0xcc, 0xed, 0x2d, 0xcd, 0x6f, 0xa6, 0xf2, 0x02, 0x5b, 0xc6, 0x3e, 0x5c, 0x23, 0x0a, 0x6f,
	0xdc, 0x84, 0xd3, 0xd7, 0x10, 0xbb, 0x61, 0xfc, 0xa2, 0x1f, 0xec, 0xe5, 0xec, 0x14, 0xc2, 0xb7,
	0x92, 0x93, 0x63, 0x80, 0xf7, 0x68, 0xc6, 0x67, 0x34, 0x01, 0xf7, 0xa8, 0x3b, 0xd9, 0xa3, 0x78,
	0xd6, 0x40, 0xfa, 0xe8, 0x34, 0x28, 0x4e, 0xe0, 0x39, 0x1f, 0xf2, 0x56, 0xc9, 0x26, 0xc7, 0x5f,
	0xac, 0x97, 0x1d, 0xea, 0xbc, 0x47, 0xad, 0x59, 0x8b, 0xba, 0xd8, 0x7c, 0xf4, 0x5f, 0x97, 0xf6,
	0xfa, 0x2f, 0x83, 0x6f, 0x4b, 0xf7, 0x1b, 0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x14, 0x7d,
	0x48, 0x4d, 0x13, 0x03, 0x00, 0x00,
}
