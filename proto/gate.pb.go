// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gate.proto

package gate

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//define node status
type NodeStatus int32

const (
	NodeStatus_NODE_DOWN     NodeStatus = 0
	NodeStatus_NODE_UP       NodeStatus = 1
	NodeStatus_NODE_MAINTAIN NodeStatus = 2
)

var NodeStatus_name = map[int32]string{
	0: "NODE_DOWN",
	1: "NODE_UP",
	2: "NODE_MAINTAIN",
}

var NodeStatus_value = map[string]int32{
	"NODE_DOWN":     0,
	"NODE_UP":       1,
	"NODE_MAINTAIN": 2,
}

func (x NodeStatus) String() string {
	return proto.EnumName(NodeStatus_name, int32(x))
}

func (NodeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}

//auth info
type AccessAuth struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessAuth) Reset()         { *m = AccessAuth{} }
func (m *AccessAuth) String() string { return proto.CompactTextString(m) }
func (*AccessAuth) ProtoMessage()    {}
func (*AccessAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}

func (m *AccessAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessAuth.Unmarshal(m, b)
}
func (m *AccessAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessAuth.Marshal(b, m, deterministic)
}
func (m *AccessAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessAuth.Merge(m, src)
}
func (m *AccessAuth) XXX_Size() int {
	return xxx_messageInfo_AccessAuth.Size(m)
}
func (m *AccessAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessAuth.DiscardUnknown(m)
}

var xxx_messageInfo_AccessAuth proto.InternalMessageInfo

func (m *AccessAuth) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *AccessAuth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//byte message data
type ByteMessage struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MessageId            uint32   `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	ConnId               uint32   `protobuf:"varint,3,opt,name=connId,proto3" json:"connId,omitempty"`
	PlayerId             int64    `protobuf:"varint,4,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	CastConnIds          []uint32 `protobuf:"varint,6,rep,packed,name=castConnIds,proto3" json:"castConnIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByteMessage) Reset()         { *m = ByteMessage{} }
func (m *ByteMessage) String() string { return proto.CompactTextString(m) }
func (*ByteMessage) ProtoMessage()    {}
func (*ByteMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{1}
}

func (m *ByteMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByteMessage.Unmarshal(m, b)
}
func (m *ByteMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByteMessage.Marshal(b, m, deterministic)
}
func (m *ByteMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByteMessage.Merge(m, src)
}
func (m *ByteMessage) XXX_Size() int {
	return xxx_messageInfo_ByteMessage.Size(m)
}
func (m *ByteMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ByteMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ByteMessage proto.InternalMessageInfo

func (m *ByteMessage) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ByteMessage) GetMessageId() uint32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *ByteMessage) GetConnId() uint32 {
	if m != nil {
		return m.ConnId
	}
	return 0
}

func (m *ByteMessage) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *ByteMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ByteMessage) GetCastConnIds() []uint32 {
	if m != nil {
		return m.CastConnIds
	}
	return nil
}

//general request
type GateReq struct {
	Service              string      `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MessageId            uint32      `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Address              string      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Data                 []byte      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	IsAsync              bool        `protobuf:"varint,5,opt,name=isAsync,proto3" json:"isAsync,omitempty"`
	Auth                 *AccessAuth `protobuf:"bytes,6,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GateReq) Reset()         { *m = GateReq{} }
func (m *GateReq) String() string { return proto.CompactTextString(m) }
func (*GateReq) ProtoMessage()    {}
func (*GateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{2}
}

func (m *GateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GateReq.Unmarshal(m, b)
}
func (m *GateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GateReq.Marshal(b, m, deterministic)
}
func (m *GateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GateReq.Merge(m, src)
}
func (m *GateReq) XXX_Size() int {
	return xxx_messageInfo_GateReq.Size(m)
}
func (m *GateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GateReq.DiscardUnknown(m)
}

var xxx_messageInfo_GateReq proto.InternalMessageInfo

func (m *GateReq) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GateReq) GetMessageId() uint32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *GateReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GateReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GateReq) GetIsAsync() bool {
	if m != nil {
		return m.IsAsync
	}
	return false
}

func (m *GateReq) GetAuth() *AccessAuth {
	if m != nil {
		return m.Auth
	}
	return nil
}

//general response
type GateResp struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MessageId            uint32   `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	ErrorCode            int32    `protobuf:"varint,3,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,4,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GateResp) Reset()         { *m = GateResp{} }
func (m *GateResp) String() string { return proto.CompactTextString(m) }
func (*GateResp) ProtoMessage()    {}
func (*GateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{3}
}

func (m *GateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GateResp.Unmarshal(m, b)
}
func (m *GateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GateResp.Marshal(b, m, deterministic)
}
func (m *GateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GateResp.Merge(m, src)
}
func (m *GateResp) XXX_Size() int {
	return xxx_messageInfo_GateResp.Size(m)
}
func (m *GateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GateResp.DiscardUnknown(m)
}

var xxx_messageInfo_GateResp proto.InternalMessageInfo

func (m *GateResp) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GateResp) GetMessageId() uint32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *GateResp) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GateResp) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *GateResp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("gate.NodeStatus", NodeStatus_name, NodeStatus_value)
	proto.RegisterType((*AccessAuth)(nil), "gate.AccessAuth")
	proto.RegisterType((*ByteMessage)(nil), "gate.ByteMessage")
	proto.RegisterType((*GateReq)(nil), "gate.GateReq")
	proto.RegisterType((*GateResp)(nil), "gate.GateResp")
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor_743bb58a714d8b7d) }

var fileDescriptor_743bb58a714d8b7d = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9d, 0x36, 0x4d, 0x9a, 0x97, 0x8d, 0x64, 0x1f, 0x22, 0x43, 0xd9, 0x43, 0x08, 0x0b,
	0x06, 0x0f, 0x45, 0x56, 0xf1, 0xe2, 0x29, 0xdd, 0x95, 0x25, 0x87, 0xcd, 0xca, 0x54, 0xf1, 0x28,
	0x63, 0xe6, 0xb1, 0x5b, 0xb4, 0x49, 0xcc, 0x4c, 0x85, 0x7e, 0x13, 0xbf, 0x85, 0x07, 0xbf, 0xa0,
	0x74, 0x92, 0x6e, 0x23, 0xee, 0xa9, 0xb7, 0xf7, 0xfb, 0xcf, 0x4c, 0xf8, 0x65, 0xfe, 0x09, 0xc0,
	0x9d, 0x34, 0x34, 0x6f, 0xda, 0xda, 0xd4, 0xe8, 0xec, 0xe6, 0xe4, 0x0d, 0x40, 0x56, 0x96, 0xa4,
	0x75, 0xb6, 0x31, 0xf7, 0x18, 0xc1, 0x58, 0x36, 0x0d, 0x67, 0x31, 0x4b, 0x7d, 0xb1, 0x1b, 0xf1,
	0x19, 0x4c, 0x4c, 0xfd, 0x8d, 0x2a, 0x3e, 0xb2, 0x59, 0x07, 0xc9, 0x1f, 0x06, 0xc1, 0x62, 0x6b,
	0xe8, 0x86, 0xb4, 0x96, 0x77, 0x84, 0x1c, 0x3c, 0x4d, 0xed, 0xcf, 0x55, 0x49, 0xfd, 0xd9, 0x3d,
	0xe2, 0x19, 0xf8, 0xeb, 0x6e, 0x53, 0xae, 0xec, 0x33, 0x42, 0x71, 0x08, 0xf0, 0x39, 0xb8, 0x65,
	0x5d, 0x55, 0xb9, 0xe2, 0x63, 0xbb, 0xd4, 0x13, 0xce, 0x60, 0xda, 0x7c, 0x97, 0x5b, 0x6a, 0x73,
	0xc5, 0x9d, 0x98, 0xa5, 0x63, 0xf1, 0xc0, 0x88, 0xe0, 0x28, 0x69, 0x24, 0x9f, 0xc4, 0x2c, 0x3d,
	0x11, 0x76, 0xc6, 0x73, 0x08, 0x4a, 0xa9, 0xcd, 0xa5, 0x3d, 0xad, 0xb9, 0x1b, 0x8f, 0xd3, 0x70,
	0x31, 0x8a, 0x98, 0x18, 0xc6, 0xc9, 0x6f, 0x06, 0xde, 0xb5, 0x34, 0x24, 0xe8, 0xc7, 0xd1, 0xc6,
	0x1c, 0x3c, 0xa9, 0x54, 0x4b, 0x5a, 0x5b, 0x65, 0x5f, 0xec, 0xf1, 0xc1, 0xcb, 0x19, 0x78, 0x71,
	0xf0, 0x56, 0x3a, 0xd3, 0xdb, 0xaa, 0xb4, 0xba, 0x53, 0xb1, 0x47, 0x3c, 0x07, 0x47, 0x6e, 0xcc,
	0x3d, 0x77, 0x63, 0x96, 0x06, 0x17, 0xd1, 0xdc, 0x16, 0x73, 0x68, 0x42, 0xd8, 0xd5, 0xe4, 0x17,
	0x83, 0x69, 0x67, 0xac, 0x9b, 0xa3, 0x95, 0xcf, 0xc0, 0xa7, 0xb6, 0xad, 0xdb, 0xcb, 0x5a, 0x91,
	0x95, 0x9e, 0x88, 0x43, 0x80, 0x09, 0x9c, 0x58, 0xe8, 0xab, 0xb4, 0xfa, 0xbe, 0xf8, 0x27, 0x7b,
	0xec, 0xca, 0x5f, 0xbe, 0x03, 0x28, 0x6a, 0x45, 0x4b, 0x23, 0xcd, 0x46, 0x63, 0x08, 0x7e, 0x71,
	0x7b, 0xf5, 0xfe, 0xcb, 0xd5, 0xed, 0xe7, 0x22, 0x7a, 0x82, 0x01, 0x78, 0x16, 0x3f, 0x7d, 0x88,
	0x18, 0x9e, 0x42, 0x68, 0xe1, 0x26, 0xcb, 0x8b, 0x8f, 0x59, 0x5e, 0x44, 0xa3, 0x8b, 0x0a, 0x82,
	0xdd, 0x6b, 0x2d, 0x7b, 0xff, 0xb7, 0x00, 0x8b, 0x55, 0xa5, 0x96, 0xa6, 0x25, 0xb9, 0xc6, 0xd3,
	0xee, 0x32, 0x06, 0xdf, 0xd7, 0xec, 0xff, 0x28, 0x65, 0xaf, 0x18, 0xbe, 0x00, 0xf7, 0x9a, 0xaa,
	0x5d, 0x9d, 0x61, 0xb7, 0xa1, 0x6f, 0x77, 0xf6, 0x74, 0x88, 0xba, 0x59, 0x84, 0x10, 0x94, 0xf5,
	0x7a, 0x6e, 0x4a, 0x9b, 0x7f, 0x75, 0xed, 0x1f, 0xf0, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0xaf, 0xe8, 0xf3, 0x0f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GateServiceClient is the client API for GateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GateServiceClient interface {
	//stream mode
	//cast or receive data between gate and target service
	BindStream(ctx context.Context, opts ...grpc.CallOption) (GateService_BindStreamClient, error)
	//general sync mode
	GenReq(ctx context.Context, in *GateReq, opts ...grpc.CallOption) (*GateResp, error)
}

type gateServiceClient struct {
	cc *grpc.ClientConn
}

func NewGateServiceClient(cc *grpc.ClientConn) GateServiceClient {
	return &gateServiceClient{cc}
}

func (c *gateServiceClient) BindStream(ctx context.Context, opts ...grpc.CallOption) (GateService_BindStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GateService_serviceDesc.Streams[0], "/gate.GateService/BindStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gateServiceBindStreamClient{stream}
	return x, nil
}

type GateService_BindStreamClient interface {
	Send(*ByteMessage) error
	Recv() (*ByteMessage, error)
	grpc.ClientStream
}

type gateServiceBindStreamClient struct {
	grpc.ClientStream
}

func (x *gateServiceBindStreamClient) Send(m *ByteMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gateServiceBindStreamClient) Recv() (*ByteMessage, error) {
	m := new(ByteMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gateServiceClient) GenReq(ctx context.Context, in *GateReq, opts ...grpc.CallOption) (*GateResp, error) {
	out := new(GateResp)
	err := c.cc.Invoke(ctx, "/gate.GateService/GenReq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateServiceServer is the server API for GateService service.
type GateServiceServer interface {
	//stream mode
	//cast or receive data between gate and target service
	BindStream(GateService_BindStreamServer) error
	//general sync mode
	GenReq(context.Context, *GateReq) (*GateResp, error)
}

// UnimplementedGateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGateServiceServer struct {
}

func (*UnimplementedGateServiceServer) BindStream(srv GateService_BindStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BindStream not implemented")
}
func (*UnimplementedGateServiceServer) GenReq(ctx context.Context, req *GateReq) (*GateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenReq not implemented")
}

func RegisterGateServiceServer(s *grpc.Server, srv GateServiceServer) {
	s.RegisterService(&_GateService_serviceDesc, srv)
}

func _GateService_BindStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GateServiceServer).BindStream(&gateServiceBindStreamServer{stream})
}

type GateService_BindStreamServer interface {
	Send(*ByteMessage) error
	Recv() (*ByteMessage, error)
	grpc.ServerStream
}

type gateServiceBindStreamServer struct {
	grpc.ServerStream
}

func (x *gateServiceBindStreamServer) Send(m *ByteMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gateServiceBindStreamServer) Recv() (*ByteMessage, error) {
	m := new(ByteMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GateService_GenReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServiceServer).GenReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gate.GateService/GenReq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServiceServer).GenReq(ctx, req.(*GateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gate.GateService",
	HandlerType: (*GateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenReq",
			Handler:    _GateService_GenReq_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BindStream",
			Handler:       _GateService_BindStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gate.proto",
}
