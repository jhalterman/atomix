// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: runtime/map/v1/maps.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/atomix/atomix/api/runtime/v1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Config struct {
	Cache CacheConfig `protobuf:"bytes,1,opt,name=cache,proto3" json:"cache"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc6c084a686856c, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetCache() CacheConfig {
	if m != nil {
		return m.Cache
	}
	return CacheConfig{}
}

type CacheConfig struct {
	Enabled bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Size_   uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (m *CacheConfig) Reset()         { *m = CacheConfig{} }
func (m *CacheConfig) String() string { return proto.CompactTextString(m) }
func (*CacheConfig) ProtoMessage()    {}
func (*CacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc6c084a686856c, []int{1}
}
func (m *CacheConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig.Merge(m, src)
}
func (m *CacheConfig) XXX_Size() int {
	return m.Size()
}
func (m *CacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig proto.InternalMessageInfo

func (m *CacheConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *CacheConfig) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

type CreateRequest struct {
	ID   v1.PrimitiveID `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Tags []string       `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc6c084a686856c, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetID() v1.PrimitiveID {
	if m != nil {
		return m.ID
	}
	return v1.PrimitiveID{}
}

func (m *CreateRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateResponse struct {
	Config Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc6c084a686856c, []int{3}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetConfig() Config {
	if m != nil {
		return m.Config
	}
	return Config{}
}

type CloseRequest struct {
	ID v1.PrimitiveID `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc6c084a686856c, []int{4}
}
func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return m.Size()
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetID() v1.PrimitiveID {
	if m != nil {
		return m.ID
	}
	return v1.PrimitiveID{}
}

type CloseResponse struct {
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dc6c084a686856c, []int{5}
}
func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return m.Size()
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Config)(nil), "atomix.runtime.map.v1.Config")
	proto.RegisterType((*CacheConfig)(nil), "atomix.runtime.map.v1.CacheConfig")
	proto.RegisterType((*CreateRequest)(nil), "atomix.runtime.map.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "atomix.runtime.map.v1.CreateResponse")
	proto.RegisterType((*CloseRequest)(nil), "atomix.runtime.map.v1.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "atomix.runtime.map.v1.CloseResponse")
}

func init() { proto.RegisterFile("runtime/map/v1/maps.proto", fileDescriptor_5dc6c084a686856c) }

var fileDescriptor_5dc6c084a686856c = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x18, 0x85, 0x33, 0x31, 0x8d, 0xfa, 0xd7, 0x2a, 0x0c, 0x15, 0x62, 0xc0, 0xb4, 0xc4, 0x0a, 0x5d,
	0x25, 0xa4, 0x6e, 0xc4, 0x82, 0x8b, 0xa4, 0x0b, 0x2b, 0x14, 0x4a, 0x40, 0xb7, 0x32, 0x6d, 0xc7,
	0x38, 0xd0, 0x64, 0x62, 0x26, 0x0d, 0xe2, 0x53, 0xf8, 0x2a, 0xbe, 0x45, 0x97, 0x5d, 0xba, 0x2a,
	0x92, 0xbe, 0x88, 0x64, 0x92, 0x40, 0x90, 0xdb, 0x7b, 0x37, 0x77, 0x95, 0x3f, 0x99, 0x73, 0xfe,
	0xf3, 0x9d, 0x30, 0xf0, 0x22, 0x3b, 0x24, 0x39, 0x8b, 0xa9, 0x1b, 0x93, 0xd4, 0x2d, 0xbc, 0xea,
	0x21, 0x9c, 0x34, 0xe3, 0x39, 0xc7, 0xcf, 0x49, 0xce, 0x63, 0xf6, 0xc3, 0x69, 0x14, 0x4e, 0x4c,
	0x52, 0xa7, 0xf0, 0x4c, 0xa3, 0x75, 0x14, 0x9e, 0xdb, 0x1e, 0x49, 0x83, 0x39, 0x8c, 0x78, 0xc4,
	0xe5, 0xe8, 0x56, 0x53, 0xfd, 0xd5, 0xfe, 0x00, 0x7a, 0xc0, 0x93, 0xaf, 0x2c, 0xc2, 0xef, 0xa1,
	0xb7, 0x25, 0xdb, 0x6f, 0xd4, 0x40, 0x63, 0x34, 0xed, 0xcf, 0x6c, 0xe7, 0xc6, 0x00, 0x27, 0xa8,
	0x34, 0xb5, 0xc5, 0xd7, 0x8e, 0xe7, 0x91, 0x12, 0xd6, 0x36, 0x7b, 0x0e, 0xfd, 0xce, 0x19, 0x36,
	0xe0, 0x21, 0x4d, 0xc8, 0x66, 0x4f, 0x77, 0x72, 0xe1, 0xa3, 0xb0, 0x7d, 0xc5, 0x18, 0x34, 0xc1,
	0x7e, 0x52, 0x43, 0x1d, 0xa3, 0xa9, 0x16, 0xca, 0xd9, 0xfe, 0x02, 0x83, 0x20, 0xa3, 0x24, 0xa7,
	0x21, 0xfd, 0x7e, 0xa0, 0x22, 0xc7, 0xef, 0x40, 0x65, 0xbb, 0x06, 0xc5, 0xfa, 0x1f, 0xa5, 0xf0,
	0x9c, 0x75, 0xc6, 0x62, 0x96, 0xb3, 0x82, 0x2e, 0x17, 0x3e, 0x54, 0x18, 0xe5, 0x79, 0xa4, 0x2e,
	0x17, 0xa1, 0xca, 0x64, 0x40, 0x4e, 0x22, 0x61, 0xa8, 0xe3, 0x07, 0xd3, 0xc7, 0xa1, 0x9c, 0xed,
	0x15, 0x3c, 0x6d, 0x03, 0x44, 0xca, 0x13, 0x41, 0xf1, 0x1c, 0xf4, 0xad, 0x44, 0x6d, 0x52, 0x5e,
	0x5e, 0x2b, 0xdc, 0xed, 0xda, 0x58, 0xec, 0x8f, 0xf0, 0x24, 0xd8, 0x73, 0x71, 0x1f, 0xb8, 0xf6,
	0x33, 0x18, 0x34, 0xbb, 0x6a, 0xb2, 0xd9, 0x6f, 0x04, 0xda, 0x8a, 0xa4, 0x02, 0x7f, 0x02, 0xbd,
	0x86, 0xc6, 0x93, 0x6b, 0x70, 0xdd, 0x9f, 0x66, 0xbe, 0xbe, 0x43, 0xd5, 0x34, 0x0f, 0xa1, 0x27,
	0x03, 0xf1, 0xab, 0x6b, 0xfa, 0x4e, 0x35, 0x73, 0x72, 0xbb, 0xa8, 0xde, 0xe9, 0xbf, 0x3d, 0x96,
	0x16, 0x3a, 0x95, 0x16, 0xfa, 0x5b, 0x5a, 0xe8, 0xd7, 0xc5, 0x52, 0x4e, 0x17, 0x4b, 0xf9, 0x73,
	0xb1, 0x14, 0x18, 0x32, 0xde, 0x6e, 0x20, 0x29, 0x6b, 0xdc, 0xbe, 0x5e, 0x15, 0xfc, 0xec, 0xad,
	0xd1, 0x46, 0x97, 0x17, 0xf1, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xa0, 0xae, 0xac,
	0xec, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MapsClient is the client API for Maps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MapsClient interface {
	// Create creates the map
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Close closes the map
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
}

type mapsClient struct {
	cc *grpc.ClientConn
}

func NewMapsClient(cc *grpc.ClientConn) MapsClient {
	return &mapsClient{cc}
}

func (c *mapsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/atomix.runtime.map.v1.Maps/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/atomix.runtime.map.v1.Maps/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapsServer is the server API for Maps service.
type MapsServer interface {
	// Create creates the map
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Close closes the map
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
}

// UnimplementedMapsServer can be embedded to have forward compatible implementations.
type UnimplementedMapsServer struct {
}

func (*UnimplementedMapsServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMapsServer) Close(ctx context.Context, req *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterMapsServer(s *grpc.Server, srv MapsServer) {
	s.RegisterService(&_Maps_serviceDesc, srv)
}

func _Maps_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.runtime.map.v1.Maps/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.runtime.map.v1.Maps/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.map.v1.Maps",
	HandlerType: (*MapsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Maps_Create_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Maps_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runtime/map/v1/maps.proto",
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Cache.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMaps(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CacheConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Size_ != 0 {
		i = encodeVarintMaps(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintMaps(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMaps(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CreateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMaps(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CloseRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloseRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloseRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMaps(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CloseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMaps(dAtA []byte, offset int, v uint64) int {
	offset -= sovMaps(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Cache.Size()
	n += 1 + l + sovMaps(uint64(l))
	return n
}

func (m *CacheConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.Size_ != 0 {
		n += 1 + sovMaps(uint64(m.Size_))
	}
	return n
}

func (m *CreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovMaps(uint64(l))
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovMaps(uint64(l))
		}
	}
	return n
}

func (m *CreateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Config.Size()
	n += 1 + l + sovMaps(uint64(l))
	return n
}

func (m *CloseRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovMaps(uint64(l))
	return n
}

func (m *CloseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMaps(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMaps(x uint64) (n int) {
	return sovMaps(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cache", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaps
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cache.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CacheConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CacheConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaps
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMaps
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaps
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloseRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloseRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloseRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaps
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMaps(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMaps
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMaps
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMaps
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMaps
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMaps
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMaps        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMaps          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMaps = fmt.Errorf("proto: unexpected end of group")
)
