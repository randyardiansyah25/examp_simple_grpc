// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package param

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Nama                 string   `protobuf:"bytes,1,opt,name=nama,proto3" json:"nama,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

func (m *Request) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Response struct {
	ResponseCode         string   `protobuf:"bytes,1,opt,name=responseCode,proto3" json:"responseCode,omitempty"`
	ResponseMessage      string   `protobuf:"bytes,2,opt,name=responseMessage,proto3" json:"responseMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseCode() string {
	if m != nil {
		return m.ResponseCode
	}
	return ""
}

func (m *Response) GetResponseMessage() string {
	if m != nil {
		return m.ResponseMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "param.Request")
	proto.RegisterType((*Response)(nil), "param.Response")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x48, 0x2c, 0x4a, 0xcc, 0x55, 0x32, 0xe6, 0x62,
	0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x73, 0x13, 0x33,
	0x73, 0x24, 0x98, 0xc0, 0x82, 0x10, 0x8e, 0x52, 0x04, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x12, 0x17, 0x4f, 0x11, 0x94, 0xed, 0x9c, 0x9f, 0x92, 0x0a, 0xd5, 0x8d,
	0x22, 0x26, 0xa4, 0xc1, 0xc5, 0x0f, 0xe3, 0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0xcd,
	0x43, 0x17, 0x36, 0xb2, 0xe7, 0xe2, 0x0e, 0x49, 0x2d, 0x2e, 0x81, 0x39, 0xc9, 0x80, 0x8b, 0xdb,
	0x3d, 0xb5, 0x04, 0x6e, 0x17, 0x9f, 0x1e, 0xd8, 0xd1, 0x7a, 0x50, 0x69, 0x29, 0x7e, 0x38, 0x1f,
	0xa2, 0x40, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x3b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84,
	0x3e, 0xa8, 0x57, 0xea, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestRequestClient is the client API for TestRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestRequestClient interface {
	GetResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testRequestClient struct {
	cc *grpc.ClientConn
}

func NewTestRequestClient(cc *grpc.ClientConn) TestRequestClient {
	return &testRequestClient{cc}
}

func (c *testRequestClient) GetResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/param.TestRequest/GetResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestRequestServer is the server API for TestRequest service.
type TestRequestServer interface {
	GetResponse(context.Context, *Request) (*Response, error)
}

func RegisterTestRequestServer(s *grpc.Server, srv TestRequestServer) {
	s.RegisterService(&_TestRequest_serviceDesc, srv)
}

func _TestRequest_GetResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRequestServer).GetResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/param.TestRequest/GetResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRequestServer).GetResponse(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestRequest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "param.TestRequest",
	HandlerType: (*TestRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResponse",
			Handler:    _TestRequest_GetResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}
