// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

package communication

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LogText struct {
	LoggedAt             *timestamp.Timestamp `protobuf:"bytes,1,opt,name=logged_at,json=loggedAt,proto3" json:"logged_at,omitempty"`
	LogMessage           string               `protobuf:"bytes,2,opt,name=log_message,json=logMessage,proto3" json:"log_message,omitempty"`
	LogFile              *LogPath             `protobuf:"bytes,3,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogText) Reset()         { *m = LogText{} }
func (m *LogText) String() string { return proto.CompactTextString(m) }
func (*LogText) ProtoMessage()    {}
func (*LogText) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

func (m *LogText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogText.Unmarshal(m, b)
}
func (m *LogText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogText.Marshal(b, m, deterministic)
}
func (m *LogText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogText.Merge(m, src)
}
func (m *LogText) XXX_Size() int {
	return xxx_messageInfo_LogText.Size(m)
}
func (m *LogText) XXX_DiscardUnknown() {
	xxx_messageInfo_LogText.DiscardUnknown(m)
}

var xxx_messageInfo_LogText proto.InternalMessageInfo

func (m *LogText) GetLoggedAt() *timestamp.Timestamp {
	if m != nil {
		return m.LoggedAt
	}
	return nil
}

func (m *LogText) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

func (m *LogText) GetLogFile() *LogPath {
	if m != nil {
		return m.LogFile
	}
	return nil
}

type LogFile struct {
	Path                 *LogPath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogFile) Reset()         { *m = LogFile{} }
func (m *LogFile) String() string { return proto.CompactTextString(m) }
func (*LogFile) ProtoMessage()    {}
func (*LogFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}

func (m *LogFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogFile.Unmarshal(m, b)
}
func (m *LogFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogFile.Marshal(b, m, deterministic)
}
func (m *LogFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogFile.Merge(m, src)
}
func (m *LogFile) XXX_Size() int {
	return xxx_messageInfo_LogFile.Size(m)
}
func (m *LogFile) XXX_DiscardUnknown() {
	xxx_messageInfo_LogFile.DiscardUnknown(m)
}

var xxx_messageInfo_LogFile proto.InternalMessageInfo

func (m *LogFile) GetPath() *LogPath {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *LogFile) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type FileList struct {
	Paths                []*LogPath `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FileList) Reset()         { *m = FileList{} }
func (m *FileList) String() string { return proto.CompactTextString(m) }
func (*FileList) ProtoMessage()    {}
func (*FileList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{2}
}

func (m *FileList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileList.Unmarshal(m, b)
}
func (m *FileList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileList.Marshal(b, m, deterministic)
}
func (m *FileList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileList.Merge(m, src)
}
func (m *FileList) XXX_Size() int {
	return xxx_messageInfo_FileList.Size(m)
}
func (m *FileList) XXX_DiscardUnknown() {
	xxx_messageInfo_FileList.DiscardUnknown(m)
}

var xxx_messageInfo_FileList proto.InternalMessageInfo

func (m *FileList) GetPaths() []*LogPath {
	if m != nil {
		return m.Paths
	}
	return nil
}

type LogPath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogPath) Reset()         { *m = LogPath{} }
func (m *LogPath) String() string { return proto.CompactTextString(m) }
func (*LogPath) ProtoMessage()    {}
func (*LogPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{3}
}

func (m *LogPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogPath.Unmarshal(m, b)
}
func (m *LogPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogPath.Marshal(b, m, deterministic)
}
func (m *LogPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogPath.Merge(m, src)
}
func (m *LogPath) XXX_Size() int {
	return xxx_messageInfo_LogPath.Size(m)
}
func (m *LogPath) XXX_DiscardUnknown() {
	xxx_messageInfo_LogPath.DiscardUnknown(m)
}

var xxx_messageInfo_LogPath proto.InternalMessageInfo

func (m *LogPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*LogText)(nil), "communication.LogText")
	proto.RegisterType((*LogFile)(nil), "communication.LogFile")
	proto.RegisterType((*FileList)(nil), "communication.FileList")
	proto.RegisterType((*LogPath)(nil), "communication.LogPath")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886) }

var fileDescriptor_a153da538f858886 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4e, 0xea, 0x50,
	0x10, 0xa5, 0x8f, 0xf7, 0x1e, 0x30, 0x7d, 0x2f, 0xc6, 0x1b, 0xa3, 0xb5, 0x89, 0x81, 0x74, 0x45,
	0x8c, 0x29, 0x11, 0x17, 0xba, 0x53, 0x24, 0xd1, 0x4d, 0x8d, 0xa6, 0x61, 0xe5, 0x86, 0x5c, 0xca,
	0x70, 0x69, 0x72, 0xdb, 0x21, 0x74, 0x30, 0xf2, 0x19, 0x6e, 0xfc, 0x5e, 0xd3, 0x5b, 0x6a, 0x04,
	0xc1, 0xdd, 0xf4, 0xcc, 0x39, 0x67, 0x4e, 0x67, 0x2e, 0x34, 0x34, 0x29, 0x7f, 0x36, 0x27, 0x26,
	0xf1, 0x3f, 0xa2, 0x24, 0x59, 0xa4, 0x71, 0x24, 0x39, 0xa6, 0xd4, 0x6d, 0x2a, 0x22, 0xa5, 0xb1,
	0x63, 0x9a, 0xa3, 0xc5, 0xa4, 0xc3, 0x71, 0x82, 0x19, 0xcb, 0x64, 0x56, 0xf0, 0xdd, 0xe3, 0x4d,
	0x82, 0x4c, 0x97, 0x45, 0xcb, 0x7b, 0xb7, 0xa0, 0x16, 0x90, 0x1a, 0xe0, 0x2b, 0x8b, 0x4b, 0x33,
	0x43, 0xe1, 0x78, 0x28, 0xd9, 0xb1, 0x5a, 0x56, 0xdb, 0xee, 0xba, 0x7e, 0x21, 0xf5, 0x4b, 0xa9,
	0x3f, 0x28, 0xbd, 0xc3, 0x7a, 0x41, 0xee, 0xb1, 0x68, 0x82, 0xad, 0x49, 0x0d, 0x13, 0xcc, 0x32,
	0xa9, 0xd0, 0xf9, 0xd5, 0xb2, 0xda, 0x8d, 0x10, 0x34, 0xa9, 0x87, 0x02, 0x11, 0xe7, 0x90, 0x93,
	0x87, 0x93, 0x58, 0xa3, 0x53, 0x35, 0xc6, 0x87, 0xfe, 0xda, 0x3f, 0xf8, 0x01, 0xa9, 0x27, 0xc9,
	0xd3, 0xb0, 0xa6, 0x49, 0xdd, 0xc5, 0x1a, 0xbd, 0x47, 0x93, 0x2b, 0x2f, 0xc5, 0x29, 0xfc, 0x9e,
	0x49, 0x9e, 0xae, 0x22, 0xed, 0x52, 0x1a, 0x8e, 0x70, 0xa0, 0x16, 0x51, 0xca, 0x98, 0xf2, 0x2a,
	0x46, 0xf9, 0xe9, 0x5d, 0x41, 0x3d, 0x77, 0x0b, 0xe2, 0x8c, 0xc5, 0x19, 0xfc, 0xc9, 0xd9, 0x99,
	0x63, 0xb5, 0xaa, 0x3f, 0x58, 0x16, 0x24, 0xef, 0xc4, 0x44, 0xc9, 0x11, 0x21, 0xbe, 0x44, 0x69,
	0x14, 0x23, 0xbb, 0x6f, 0x16, 0xfc, 0x0b, 0x48, 0xf5, 0x49, 0x6b, 0x8c, 0x98, 0xe6, 0xe2, 0x06,
	0xec, 0x7b, 0xe4, 0xcf, 0x61, 0x3b, 0xdc, 0xdd, 0xa3, 0x0d, 0xbc, 0x14, 0x78, 0x15, 0x71, 0x0d,
	0x76, 0x88, 0x72, 0x5c, 0x2e, 0x60, 0x97, 0xc3, 0x16, 0xdc, 0xec, 0xae, 0xd2, 0x0d, 0xc1, 0x0e,
	0x48, 0x85, 0x18, 0x61, 0xfc, 0x82, 0x73, 0xd1, 0x87, 0xfd, 0x55, 0x1d, 0x98, 0x9b, 0x99, 0x73,
	0x6f, 0x51, 0xe7, 0xb8, 0x7b, 0xf0, 0xed, 0xe6, 0xbd, 0x74, 0xe9, 0x55, 0x6e, 0xf7, 0x9e, 0xd7,
	0xdf, 0xdd, 0xe8, 0xaf, 0x21, 0x5c, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xb3, 0xbf, 0x20,
	0x9a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogCollectorClient is the client API for LogCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogCollectorClient interface {
	GetFileList(ctx context.Context, in *LogPath, opts ...grpc.CallOption) (*FileList, error)
	ReadLogFile(ctx context.Context, in *LogPath, opts ...grpc.CallOption) (*LogFile, error)
}

type logCollectorClient struct {
	cc *grpc.ClientConn
}

func NewLogCollectorClient(cc *grpc.ClientConn) LogCollectorClient {
	return &logCollectorClient{cc}
}

func (c *logCollectorClient) GetFileList(ctx context.Context, in *LogPath, opts ...grpc.CallOption) (*FileList, error) {
	out := new(FileList)
	err := c.cc.Invoke(ctx, "/communication.LogCollector/GetFileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logCollectorClient) ReadLogFile(ctx context.Context, in *LogPath, opts ...grpc.CallOption) (*LogFile, error) {
	out := new(LogFile)
	err := c.cc.Invoke(ctx, "/communication.LogCollector/ReadLogFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogCollectorServer is the server API for LogCollector service.
type LogCollectorServer interface {
	GetFileList(context.Context, *LogPath) (*FileList, error)
	ReadLogFile(context.Context, *LogPath) (*LogFile, error)
}

// UnimplementedLogCollectorServer can be embedded to have forward compatible implementations.
type UnimplementedLogCollectorServer struct {
}

func (*UnimplementedLogCollectorServer) GetFileList(ctx context.Context, req *LogPath) (*FileList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileList not implemented")
}
func (*UnimplementedLogCollectorServer) ReadLogFile(ctx context.Context, req *LogPath) (*LogFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadLogFile not implemented")
}

func RegisterLogCollectorServer(s *grpc.Server, srv LogCollectorServer) {
	s.RegisterService(&_LogCollector_serviceDesc, srv)
}

func _LogCollector_GetFileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogCollectorServer).GetFileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.LogCollector/GetFileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogCollectorServer).GetFileList(ctx, req.(*LogPath))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogCollector_ReadLogFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogCollectorServer).ReadLogFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.LogCollector/ReadLogFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogCollectorServer).ReadLogFile(ctx, req.(*LogPath))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogCollector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "communication.LogCollector",
	HandlerType: (*LogCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileList",
			Handler:    _LogCollector_GetFileList_Handler,
		},
		{
			MethodName: "ReadLogFile",
			Handler:    _LogCollector_ReadLogFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}

// LogReceiverClient is the client API for LogReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogReceiverClient interface {
	ReceiveLoggedText(ctx context.Context, in *LogText, opts ...grpc.CallOption) (*any.Any, error)
}

type logReceiverClient struct {
	cc *grpc.ClientConn
}

func NewLogReceiverClient(cc *grpc.ClientConn) LogReceiverClient {
	return &logReceiverClient{cc}
}

func (c *logReceiverClient) ReceiveLoggedText(ctx context.Context, in *LogText, opts ...grpc.CallOption) (*any.Any, error) {
	out := new(any.Any)
	err := c.cc.Invoke(ctx, "/communication.LogReceiver/ReceiveLoggedText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogReceiverServer is the server API for LogReceiver service.
type LogReceiverServer interface {
	ReceiveLoggedText(context.Context, *LogText) (*any.Any, error)
}

// UnimplementedLogReceiverServer can be embedded to have forward compatible implementations.
type UnimplementedLogReceiverServer struct {
}

func (*UnimplementedLogReceiverServer) ReceiveLoggedText(ctx context.Context, req *LogText) (*any.Any, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveLoggedText not implemented")
}

func RegisterLogReceiverServer(s *grpc.Server, srv LogReceiverServer) {
	s.RegisterService(&_LogReceiver_serviceDesc, srv)
}

func _LogReceiver_ReceiveLoggedText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogText)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogReceiverServer).ReceiveLoggedText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.LogReceiver/ReceiveLoggedText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogReceiverServer).ReceiveLoggedText(ctx, req.(*LogText))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogReceiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "communication.LogReceiver",
	HandlerType: (*LogReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveLoggedText",
			Handler:    _LogReceiver_ReceiveLoggedText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}