// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0--rc1
// source: oss/v1/oss.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 响应 - OSS前端直传信息
type OssStsTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// access_key
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// access_secret
	AccessSecret string `protobuf:"bytes,2,opt,name=access_secret,json=accessSecret,proto3" json:"access_secret,omitempty"`
	// 过期时间
	Expiration string `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// 安全令牌
	SecurityToken string `protobuf:"bytes,4,opt,name=security_token,json=securityToken,proto3" json:"security_token,omitempty"`
	// 终端
	EndPoint string `protobuf:"bytes,5,opt,name=end_point,json=endPoint,proto3" json:"end_point,omitempty"`
	// 存储桶
	BucketName string `protobuf:"bytes,6,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	// 区域
	Region string `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	// url
	Url string `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OssStsTokenResponse) Reset() {
	*x = OssStsTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_v1_oss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OssStsTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OssStsTokenResponse) ProtoMessage() {}

func (x *OssStsTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oss_v1_oss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OssStsTokenResponse.ProtoReflect.Descriptor instead.
func (*OssStsTokenResponse) Descriptor() ([]byte, []int) {
	return file_oss_v1_oss_proto_rawDescGZIP(), []int{0}
}

func (x *OssStsTokenResponse) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *OssStsTokenResponse) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

func (x *OssStsTokenResponse) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *OssStsTokenResponse) GetSecurityToken() string {
	if x != nil {
		return x.SecurityToken
	}
	return ""
}

func (x *OssStsTokenResponse) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *OssStsTokenResponse) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *OssStsTokenResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *OssStsTokenResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// 请求 - 文件上传
type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileType string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	Content  []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_v1_oss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oss_v1_oss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_oss_v1_oss_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *UploadFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// 响应 - 文件上传url
type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_v1_oss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oss_v1_oss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_oss_v1_oss_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_oss_v1_oss_proto protoreflect.FileDescriptor

var file_oss_v1_oss_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a,
	0x13, 0x4f, 0x73, 0x73, 0x53, 0x74, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x12, 0x72, 0x10,
	0x52, 0x03, 0x6a, 0x70, 0x67, 0x52, 0x03, 0x70, 0x6e, 0x67, 0x52, 0x04, 0x6a, 0x70, 0x65, 0x67,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x7a, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a,
	0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xc5, 0x01, 0x0a, 0x03, 0x4f, 0x53, 0x53, 0x12, 0x5e, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4f, 0x73, 0x73, 0x53, 0x74, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x73, 0x73, 0x53, 0x74, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x6f, 0x73, 0x73, 0x2f, 0x73, 0x74, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x5e, 0x0a,
	0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22,
	0x0c, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x56, 0xba,
	0x47, 0x32, 0x12, 0x30, 0x0a, 0x0c, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0x9c, 0x8d, 0xe5,
	0x8a, 0xa1, 0x12, 0x0c, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1,
	0x22, 0x0b, 0x0a, 0x09, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x05, 0x30,
	0x2e, 0x30, 0x2e, 0x31, 0x5a, 0x1f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x73, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oss_v1_oss_proto_rawDescOnce sync.Once
	file_oss_v1_oss_proto_rawDescData = file_oss_v1_oss_proto_rawDesc
)

func file_oss_v1_oss_proto_rawDescGZIP() []byte {
	file_oss_v1_oss_proto_rawDescOnce.Do(func() {
		file_oss_v1_oss_proto_rawDescData = protoimpl.X.CompressGZIP(file_oss_v1_oss_proto_rawDescData)
	})
	return file_oss_v1_oss_proto_rawDescData
}

var file_oss_v1_oss_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oss_v1_oss_proto_goTypes = []interface{}{
	(*OssStsTokenResponse)(nil), // 0: file.v1.OssStsTokenResponse
	(*UploadFileRequest)(nil),   // 1: file.v1.UploadFileRequest
	(*UploadFileResponse)(nil),  // 2: file.v1.UploadFileResponse
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_oss_v1_oss_proto_depIdxs = []int32{
	3, // 0: file.v1.OSS.GetOssStsToken:input_type -> google.protobuf.Empty
	1, // 1: file.v1.OSS.UploadFile:input_type -> file.v1.UploadFileRequest
	0, // 2: file.v1.OSS.GetOssStsToken:output_type -> file.v1.OssStsTokenResponse
	2, // 3: file.v1.OSS.UploadFile:output_type -> file.v1.UploadFileResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oss_v1_oss_proto_init() }
func file_oss_v1_oss_proto_init() {
	if File_oss_v1_oss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oss_v1_oss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OssStsTokenResponse); i {
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
		file_oss_v1_oss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_oss_v1_oss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
			RawDescriptor: file_oss_v1_oss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oss_v1_oss_proto_goTypes,
		DependencyIndexes: file_oss_v1_oss_proto_depIdxs,
		MessageInfos:      file_oss_v1_oss_proto_msgTypes,
	}.Build()
	File_oss_v1_oss_proto = out.File
	file_oss_v1_oss_proto_rawDesc = nil
	file_oss_v1_oss_proto_goTypes = nil
	file_oss_v1_oss_proto_depIdxs = nil
}
