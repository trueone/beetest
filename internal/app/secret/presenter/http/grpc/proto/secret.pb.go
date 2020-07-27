// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: internal/app/secret/presenter/http/grpc/proto/secret.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret           string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	ExpireAfterViews int64  `protobuf:"varint,2,opt,name=expireAfterViews,proto3" json:"expireAfterViews,omitempty"`
	ExpireAfter      int64  `protobuf:"varint,3,opt,name=expireAfter,proto3" json:"expireAfter,omitempty"`
}

func (x *SecretRequest) Reset() {
	*x = SecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretRequest) ProtoMessage() {}

func (x *SecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretRequest.ProtoReflect.Descriptor instead.
func (*SecretRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescGZIP(), []int{0}
}

func (x *SecretRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *SecretRequest) GetExpireAfterViews() int64 {
	if x != nil {
		return x.ExpireAfterViews
	}
	return 0
}

func (x *SecretRequest) GetExpireAfter() int64 {
	if x != nil {
		return x.ExpireAfter
	}
	return 0
}

type HashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *HashRequest) Reset() {
	*x = HashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashRequest) ProtoMessage() {}

func (x *HashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashRequest.ProtoReflect.Descriptor instead.
func (*HashRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescGZIP(), []int{1}
}

func (x *HashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type SecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash           string               `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	SecretText     string               `protobuf:"bytes,2,opt,name=secretText,proto3" json:"secretText,omitempty"`
	CreatedAt      *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	RemainingViews int64                `protobuf:"varint,5,opt,name=remainingViews,proto3" json:"remainingViews,omitempty"`
}

func (x *SecretResponse) Reset() {
	*x = SecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretResponse) ProtoMessage() {}

func (x *SecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretResponse.ProtoReflect.Descriptor instead.
func (*SecretResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescGZIP(), []int{2}
}

func (x *SecretResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *SecretResponse) GetSecretText() string {
	if x != nil {
		return x.SecretText
	}
	return ""
}

func (x *SecretResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SecretResponse) GetExpiresAt() *timestamp.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *SecretResponse) GetRemainingViews() int64 {
	if x != nil {
		return x.RemainingViews
	}
	return 0
}

var File_internal_app_secret_presenter_http_grpc_proto_secret_proto protoreflect.FileDescriptor

var file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0b, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xe0,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6d,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x69, 0x65, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x69, 0x65, 0x77,
	0x73, 0x32, 0x75, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x75, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x62,
	0x65, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescOnce sync.Once
	file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescData = file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDesc
)

func file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescGZIP() []byte {
	file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescOnce.Do(func() {
		file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescData)
	})
	return file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDescData
}

var file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_app_secret_presenter_http_grpc_proto_secret_proto_goTypes = []interface{}{
	(*SecretRequest)(nil),       // 0: proto.SecretRequest
	(*HashRequest)(nil),         // 1: proto.HashRequest
	(*SecretResponse)(nil),      // 2: proto.SecretResponse
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_internal_app_secret_presenter_http_grpc_proto_secret_proto_depIdxs = []int32{
	3, // 0: proto.SecretResponse.createdAt:type_name -> google.protobuf.Timestamp
	3, // 1: proto.SecretResponse.expiresAt:type_name -> google.protobuf.Timestamp
	0, // 2: proto.Secret.Create:input_type -> proto.SecretRequest
	1, // 3: proto.Secret.Get:input_type -> proto.HashRequest
	2, // 4: proto.Secret.Create:output_type -> proto.SecretResponse
	2, // 5: proto.Secret.Get:output_type -> proto.SecretResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_app_secret_presenter_http_grpc_proto_secret_proto_init() }
func file_internal_app_secret_presenter_http_grpc_proto_secret_proto_init() {
	if File_internal_app_secret_presenter_http_grpc_proto_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_secret_presenter_http_grpc_proto_secret_proto_goTypes,
		DependencyIndexes: file_internal_app_secret_presenter_http_grpc_proto_secret_proto_depIdxs,
		MessageInfos:      file_internal_app_secret_presenter_http_grpc_proto_secret_proto_msgTypes,
	}.Build()
	File_internal_app_secret_presenter_http_grpc_proto_secret_proto = out.File
	file_internal_app_secret_presenter_http_grpc_proto_secret_proto_rawDesc = nil
	file_internal_app_secret_presenter_http_grpc_proto_secret_proto_goTypes = nil
	file_internal_app_secret_presenter_http_grpc_proto_secret_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SecretClient is the client API for Secret service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretClient interface {
	Create(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretResponse, error)
	Get(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*SecretResponse, error)
}

type secretClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretClient(cc grpc.ClientConnInterface) SecretClient {
	return &secretClient{cc}
}

func (c *secretClient) Create(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) Get(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServer is the server API for Secret service.
type SecretServer interface {
	Create(context.Context, *SecretRequest) (*SecretResponse, error)
	Get(context.Context, *HashRequest) (*SecretResponse, error)
}

// UnimplementedSecretServer can be embedded to have forward compatible implementations.
type UnimplementedSecretServer struct {
}

func (*UnimplementedSecretServer) Create(context.Context, *SecretRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSecretServer) Get(context.Context, *HashRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterSecretServer(s *grpc.Server, srv SecretServer) {
	s.RegisterService(&_Secret_serviceDesc, srv)
}

func _Secret_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Create(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Get(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Secret_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Secret",
	HandlerType: (*SecretServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Secret_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Secret_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/secret/presenter/http/grpc/proto/secret.proto",
}
