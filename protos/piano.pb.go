// Code generated by protoc-gen-go. DO NOT EDIT.
// source: piano.proto

package piano

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

type Note struct {
	Pitch                uint32   `protobuf:"varint,1,opt,name=pitch,proto3" json:"pitch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_385ee2a16a45e920, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetPitch() uint32 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func init() {
	proto.RegisterType((*Note)(nil), "piano.Note")
}

func init() { proto.RegisterFile("piano.proto", fileDescriptor_385ee2a16a45e920) }

var fileDescriptor_385ee2a16a45e920 = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc8, 0x4c, 0xcc,
	0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x64, 0xb8, 0x58, 0xfc,
	0xf2, 0x4b, 0x52, 0x85, 0x44, 0xb8, 0x58, 0x0b, 0x32, 0x4b, 0x92, 0x33, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x78, 0x83, 0x20, 0x1c, 0x23, 0x23, 0x2e, 0xd6, 0x00, 0x90, 0x32, 0x21, 0x4d, 0x2e, 0x76,
	0xe7, 0xfc, 0xbc, 0xbc, 0xd4, 0xe4, 0x12, 0x21, 0x6e, 0x3d, 0x88, 0x31, 0x20, 0x6d, 0x52, 0xc8,
	0x1c, 0x25, 0x06, 0x0d, 0x46, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0xf9, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x6e, 0x32, 0x0a, 0x6e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PianoClient is the client API for Piano service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PianoClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Piano_ConnectClient, error)
}

type pianoClient struct {
	cc *grpc.ClientConn
}

func NewPianoClient(cc *grpc.ClientConn) PianoClient {
	return &pianoClient{cc}
}

func (c *pianoClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Piano_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Piano_serviceDesc.Streams[0], "/piano.Piano/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &pianoConnectClient{stream}
	return x, nil
}

type Piano_ConnectClient interface {
	Send(*Note) error
	Recv() (*Note, error)
	grpc.ClientStream
}

type pianoConnectClient struct {
	grpc.ClientStream
}

func (x *pianoConnectClient) Send(m *Note) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pianoConnectClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PianoServer is the server API for Piano service.
type PianoServer interface {
	Connect(Piano_ConnectServer) error
}

// UnimplementedPianoServer can be embedded to have forward compatible implementations.
type UnimplementedPianoServer struct {
}

func (*UnimplementedPianoServer) Connect(srv Piano_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterPianoServer(s *grpc.Server, srv PianoServer) {
	s.RegisterService(&_Piano_serviceDesc, srv)
}

func _Piano_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PianoServer).Connect(&pianoConnectServer{stream})
}

type Piano_ConnectServer interface {
	Send(*Note) error
	Recv() (*Note, error)
	grpc.ServerStream
}

type pianoConnectServer struct {
	grpc.ServerStream
}

func (x *pianoConnectServer) Send(m *Note) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pianoConnectServer) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Piano_serviceDesc = grpc.ServiceDesc{
	ServiceName: "piano.Piano",
	HandlerType: (*PianoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Piano_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "piano.proto",
}