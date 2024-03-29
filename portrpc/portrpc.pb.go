// Code generated by protoc-gen-go. DO NOT EDIT.
// source: portrpc.proto

package portrpc

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

type PutResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca1483cacdcc443, []int{0}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type Shortcode struct {
	Shortcode            string   `protobuf:"bytes,1,opt,name=Shortcode,proto3" json:"Shortcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shortcode) Reset()         { *m = Shortcode{} }
func (m *Shortcode) String() string { return proto.CompactTextString(m) }
func (*Shortcode) ProtoMessage()    {}
func (*Shortcode) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca1483cacdcc443, []int{1}
}

func (m *Shortcode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shortcode.Unmarshal(m, b)
}
func (m *Shortcode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shortcode.Marshal(b, m, deterministic)
}
func (m *Shortcode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shortcode.Merge(m, src)
}
func (m *Shortcode) XXX_Size() int {
	return xxx_messageInfo_Shortcode.Size(m)
}
func (m *Shortcode) XXX_DiscardUnknown() {
	xxx_messageInfo_Shortcode.DiscardUnknown(m)
}

var xxx_messageInfo_Shortcode proto.InternalMessageInfo

func (m *Shortcode) GetShortcode() string {
	if m != nil {
		return m.Shortcode
	}
	return ""
}

// "Make a new message that simply has no fields. That way you can
// add new fields when you need to and not break any application code."
type GetShortcodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShortcodesRequest) Reset()         { *m = GetShortcodesRequest{} }
func (m *GetShortcodesRequest) String() string { return proto.CompactTextString(m) }
func (*GetShortcodesRequest) ProtoMessage()    {}
func (*GetShortcodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca1483cacdcc443, []int{2}
}

func (m *GetShortcodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShortcodesRequest.Unmarshal(m, b)
}
func (m *GetShortcodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShortcodesRequest.Marshal(b, m, deterministic)
}
func (m *GetShortcodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShortcodesRequest.Merge(m, src)
}
func (m *GetShortcodesRequest) XXX_Size() int {
	return xxx_messageInfo_GetShortcodesRequest.Size(m)
}
func (m *GetShortcodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShortcodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShortcodesRequest proto.InternalMessageInfo

// Mildly cheated by using `go2proto` for this to start with.
// Annoyingly it lowercases all the fieldnames. But it'll do
// as a start. Later versions could use `struct.proto` and
// work of a 'real' Go `struct` (if that gives any benefit.)
type Port struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float64 `protobuf:"fixed64,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string  `protobuf:"bytes,9,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	Shortcode            string    `protobuf:"bytes,11,opt,name=shortcode,proto3" json:"shortcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca1483cacdcc443, []int{3}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetCoordinates() []float64 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Port) GetShortcode() string {
	if m != nil {
		return m.Shortcode
	}
	return ""
}

func init() {
	proto.RegisterType((*PutResponse)(nil), "portrpc.PutResponse")
	proto.RegisterType((*Shortcode)(nil), "portrpc.Shortcode")
	proto.RegisterType((*GetShortcodesRequest)(nil), "portrpc.GetShortcodesRequest")
	proto.RegisterType((*Port)(nil), "portrpc.Port")
}

func init() { proto.RegisterFile("portrpc.proto", fileDescriptor_3ca1483cacdcc443) }

var fileDescriptor_3ca1483cacdcc443 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0x71, 0xfe, 0x39, 0x9e, 0x6c, 0x2e, 0x43, 0x08, 0xc2, 0xec, 0x82, 0xf1, 0x29, 0xb9,
	0x84, 0x65, 0xf7, 0xd0, 0x7b, 0x09, 0xe4, 0x1a, 0xdc, 0x27, 0x50, 0x14, 0xd1, 0x1a, 0x12, 0x8d,
	0x2b, 0xc9, 0x85, 0xf4, 0x59, 0xfa, 0x2a, 0x7d, 0xb7, 0x22, 0xd9, 0x52, 0x12, 0xc8, 0x6d, 0x7e,
	0xdf, 0x8c, 0xc7, 0xcc, 0xf7, 0x09, 0xe6, 0x0d, 0x69, 0xab, 0x1b, 0xb1, 0x69, 0x34, 0x59, 0xc2,
	0xb4, 0xc7, 0x72, 0x0d, 0xb3, 0x7d, 0x6b, 0x2b, 0x69, 0x1a, 0x52, 0x46, 0x62, 0x0e, 0xd3, 0x50,
	0xb3, 0xa4, 0x48, 0x56, 0x59, 0x15, 0xb9, 0x5c, 0x43, 0xf6, 0xf2, 0x46, 0xda, 0x0a, 0x3a, 0x4a,
	0xfc, 0x7d, 0x03, 0xfd, 0xe4, 0x55, 0x28, 0x97, 0xb0, 0xd8, 0x49, 0x1b, 0xd9, 0x54, 0xf2, 0xbd,
	0x95, 0xc6, 0x96, 0x5f, 0x03, 0x18, 0xed, 0x49, 0x5b, 0x44, 0x18, 0x29, 0x7e, 0x0e, 0x5f, 0xfa,
	0xda, 0x69, 0xa2, 0xb6, 0x17, 0x36, 0xe8, 0x34, 0x57, 0x23, 0x83, 0x54, 0x50, 0xab, 0xac, 0xbe,
	0xb0, 0xa1, 0x97, 0x03, 0xe2, 0x02, 0xc6, 0xfc, 0x54, 0x73, 0xc3, 0x46, 0xc5, 0x70, 0x95, 0x55,
	0x1d, 0xb8, 0x79, 0x2d, 0x5f, 0x6b, 0x52, 0x86, 0x8d, 0xbd, 0x1e, 0x10, 0x0b, 0x98, 0x09, 0x22,
	0x7d, 0xac, 0x15, 0xb7, 0xd2, 0xb0, 0x49, 0x31, 0x5c, 0x25, 0xd5, 0xad, 0xe4, 0x6e, 0x6f, 0x34,
	0x7d, 0xd4, 0x4a, 0x48, 0x96, 0x76, 0xb7, 0x07, 0x76, 0x3d, 0x5b, 0x9f, 0xe5, 0x27, 0x29, 0xc9,
	0xa6, 0x5d, 0x2f, 0x30, 0x2e, 0x61, 0xd2, 0xaa, 0x13, 0x09, 0xc3, 0x32, 0xff, 0xcb, 0x9e, 0xfc,
	0x3d, 0xce, 0x1d, 0xe8, 0xef, 0xe9, 0x6d, 0x33, 0xd1, 0xb6, 0x59, 0x67, 0x5b, 0x14, 0xfe, 0x7d,
	0x27, 0xf0, 0xcb, 0xd9, 0xb3, 0xe5, 0x96, 0x1f, 0xb8, 0x91, 0xb8, 0x81, 0x74, 0xdf, 0x5a, 0xef,
	0xd8, 0x7c, 0x13, 0x12, 0x74, 0x98, 0x2f, 0xae, 0x78, 0x13, 0xdf, 0x13, 0xe0, 0x4e, 0xfa, 0xf9,
	0xe7, 0xcb, 0x35, 0x2b, 0x8c, 0xb3, 0x51, 0xcb, 0xef, 0xd7, 0xe1, 0x16, 0xe6, 0x77, 0x81, 0xe1,
	0x9f, 0xd8, 0x7f, 0x14, 0x64, 0xfe, 0x60, 0xe5, 0xdf, 0xe4, 0x30, 0xf1, 0x8f, 0xeb, 0xff, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x5a, 0xce, 0x14, 0x6d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortDatabaseClient is the client API for PortDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortDatabaseClient interface {
	// `PutPort` adds a single port to the database. Primary key is the shortcode.
	PutPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*PutResponse, error)
	// `GetPortByShortcode` returns a single port matching the given shortcode.
	GetPortByShortcode(ctx context.Context, in *Shortcode, opts ...grpc.CallOption) (*Port, error)
	// `GetShortcodes` returns a sorted list of all the known shortcodes.
	GetShortcodes(ctx context.Context, in *GetShortcodesRequest, opts ...grpc.CallOption) (PortDatabase_GetShortcodesClient, error)
}

type portDatabaseClient struct {
	cc *grpc.ClientConn
}

func NewPortDatabaseClient(cc *grpc.ClientConn) PortDatabaseClient {
	return &portDatabaseClient{cc}
}

func (c *portDatabaseClient) PutPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/portrpc.PortDatabase/PutPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDatabaseClient) GetPortByShortcode(ctx context.Context, in *Shortcode, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/portrpc.PortDatabase/GetPortByShortcode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDatabaseClient) GetShortcodes(ctx context.Context, in *GetShortcodesRequest, opts ...grpc.CallOption) (PortDatabase_GetShortcodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PortDatabase_serviceDesc.Streams[0], "/portrpc.PortDatabase/GetShortcodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &portDatabaseGetShortcodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PortDatabase_GetShortcodesClient interface {
	Recv() (*Shortcode, error)
	grpc.ClientStream
}

type portDatabaseGetShortcodesClient struct {
	grpc.ClientStream
}

func (x *portDatabaseGetShortcodesClient) Recv() (*Shortcode, error) {
	m := new(Shortcode)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortDatabaseServer is the server API for PortDatabase service.
type PortDatabaseServer interface {
	// `PutPort` adds a single port to the database. Primary key is the shortcode.
	PutPort(context.Context, *Port) (*PutResponse, error)
	// `GetPortByShortcode` returns a single port matching the given shortcode.
	GetPortByShortcode(context.Context, *Shortcode) (*Port, error)
	// `GetShortcodes` returns a sorted list of all the known shortcodes.
	GetShortcodes(*GetShortcodesRequest, PortDatabase_GetShortcodesServer) error
}

// UnimplementedPortDatabaseServer can be embedded to have forward compatible implementations.
type UnimplementedPortDatabaseServer struct {
}

func (*UnimplementedPortDatabaseServer) PutPort(ctx context.Context, req *Port) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPort not implemented")
}
func (*UnimplementedPortDatabaseServer) GetPortByShortcode(ctx context.Context, req *Shortcode) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortByShortcode not implemented")
}
func (*UnimplementedPortDatabaseServer) GetShortcodes(req *GetShortcodesRequest, srv PortDatabase_GetShortcodesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetShortcodes not implemented")
}

func RegisterPortDatabaseServer(s *grpc.Server, srv PortDatabaseServer) {
	s.RegisterService(&_PortDatabase_serviceDesc, srv)
}

func _PortDatabase_PutPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDatabaseServer).PutPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portrpc.PortDatabase/PutPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDatabaseServer).PutPort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDatabase_GetPortByShortcode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shortcode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDatabaseServer).GetPortByShortcode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portrpc.PortDatabase/GetPortByShortcode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDatabaseServer).GetPortByShortcode(ctx, req.(*Shortcode))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDatabase_GetShortcodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetShortcodesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PortDatabaseServer).GetShortcodes(m, &portDatabaseGetShortcodesServer{stream})
}

type PortDatabase_GetShortcodesServer interface {
	Send(*Shortcode) error
	grpc.ServerStream
}

type portDatabaseGetShortcodesServer struct {
	grpc.ServerStream
}

func (x *portDatabaseGetShortcodesServer) Send(m *Shortcode) error {
	return x.ServerStream.SendMsg(m)
}

var _PortDatabase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portrpc.PortDatabase",
	HandlerType: (*PortDatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutPort",
			Handler:    _PortDatabase_PutPort_Handler,
		},
		{
			MethodName: "GetPortByShortcode",
			Handler:    _PortDatabase_GetPortByShortcode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetShortcodes",
			Handler:       _PortDatabase_GetShortcodes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "portrpc.proto",
}
