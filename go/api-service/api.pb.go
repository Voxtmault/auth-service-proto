// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api.proto

package api_service

import (
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

type GenericResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response uint32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GenericResponse) GetResponse() uint32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// API Sub-Services
type API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiId     uint32 `protobuf:"varint,1,opt,name=apiId,proto3" json:"apiId,omitempty"`
	ApiName   string `protobuf:"bytes,2,opt,name=apiName,proto3" json:"apiName,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *API) Reset() {
	*x = API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*API) ProtoMessage() {}

func (x *API) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use API.ProtoReflect.Descriptor instead.
func (*API) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *API) GetApiId() uint32 {
	if x != nil {
		return x.ApiId
	}
	return 0
}

func (x *API) GetApiName() string {
	if x != nil {
		return x.ApiName
	}
	return ""
}

func (x *API) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type RegisterAuthorizedAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiName string `protobuf:"bytes,1,opt,name=apiName,proto3" json:"apiName,omitempty"`
}

func (x *RegisterAuthorizedAPIRequest) Reset() {
	*x = RegisterAuthorizedAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAuthorizedAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAuthorizedAPIRequest) ProtoMessage() {}

func (x *RegisterAuthorizedAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAuthorizedAPIRequest.ProtoReflect.Descriptor instead.
func (*RegisterAuthorizedAPIRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterAuthorizedAPIRequest) GetApiName() string {
	if x != nil {
		return x.ApiName
	}
	return ""
}

type RegisterAuthorizedAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response uint32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ApiToken string `protobuf:"bytes,3,opt,name=apiToken,proto3" json:"apiToken,omitempty"` // Base64 Encoded (Use as is)
}

func (x *RegisterAuthorizedAPIResponse) Reset() {
	*x = RegisterAuthorizedAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAuthorizedAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAuthorizedAPIResponse) ProtoMessage() {}

func (x *RegisterAuthorizedAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAuthorizedAPIResponse.ProtoReflect.Descriptor instead.
func (*RegisterAuthorizedAPIResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterAuthorizedAPIResponse) GetResponse() uint32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *RegisterAuthorizedAPIResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterAuthorizedAPIResponse) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

type RollingAPIKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiToken string `protobuf:"bytes,1,opt,name=apiToken,proto3" json:"apiToken,omitempty"` // Base64 Encoded
}

func (x *RollingAPIKeyRequest) Reset() {
	*x = RollingAPIKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollingAPIKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollingAPIKeyRequest) ProtoMessage() {}

func (x *RollingAPIKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollingAPIKeyRequest.ProtoReflect.Descriptor instead.
func (*RollingAPIKeyRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *RollingAPIKeyRequest) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

type RollingAPIKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response uint32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ApiToken string `protobuf:"bytes,3,opt,name=apiToken,proto3" json:"apiToken,omitempty"` // Base64 Encoded (Use as is)
}

func (x *RollingAPIKeyResponse) Reset() {
	*x = RollingAPIKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollingAPIKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollingAPIKeyResponse) ProtoMessage() {}

func (x *RollingAPIKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollingAPIKeyResponse.ProtoReflect.Descriptor instead.
func (*RollingAPIKeyResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *RollingAPIKeyResponse) GetResponse() uint32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *RollingAPIKeyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RollingAPIKeyResponse) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

type DisableAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	ApiId  uint32 `protobuf:"varint,2,opt,name=apiId,proto3" json:"apiId,omitempty"`
}

func (x *DisableAPIRequest) Reset() {
	*x = DisableAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableAPIRequest) ProtoMessage() {}

func (x *DisableAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableAPIRequest.ProtoReflect.Descriptor instead.
func (*DisableAPIRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *DisableAPIRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *DisableAPIRequest) GetApiId() uint32 {
	if x != nil {
		return x.ApiId
	}
	return 0
}

type GetRegisteredAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetRegisteredAPIRequest) Reset() {
	*x = GetRegisteredAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegisteredAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegisteredAPIRequest) ProtoMessage() {}

func (x *GetRegisteredAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegisteredAPIRequest.ProtoReflect.Descriptor instead.
func (*GetRegisteredAPIRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetRegisteredAPIRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetRegisteredAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response uint32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Api      []*API `protobuf:"bytes,2,rep,name=api,proto3" json:"api,omitempty"`
}

func (x *GetRegisteredAPIResponse) Reset() {
	*x = GetRegisteredAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegisteredAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegisteredAPIResponse) ProtoMessage() {}

func (x *GetRegisteredAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegisteredAPIResponse.ProtoReflect.Descriptor instead.
func (*GetRegisteredAPIResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetRegisteredAPIResponse) GetResponse() uint32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *GetRegisteredAPIResponse) GetApi() []*API {
	if x != nil {
		return x.Api
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x53, 0x0a, 0x03, 0x41, 0x50,
	0x49, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x38, 0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x71, 0x0a, 0x1d, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x14,
	0x52, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x69, 0x0a, 0x15, 0x52, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x41, 0x0a, 0x11, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x69, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x22, 0x2f,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x53, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x50, 0x49, 0x52,
	0x03, 0x61, 0x70, 0x69, 0x32, 0xd1, 0x02, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x50, 0x49, 0x12, 0x22, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x50,
	0x49, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x41, 0x50, 0x49, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x50, 0x49, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x50, 0x49, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_goTypes = []interface{}{
	(*GenericResponse)(nil),               // 0: auth.GenericResponse
	(*API)(nil),                           // 1: auth.API
	(*RegisterAuthorizedAPIRequest)(nil),  // 2: auth.RegisterAuthorizedAPIRequest
	(*RegisterAuthorizedAPIResponse)(nil), // 3: auth.RegisterAuthorizedAPIResponse
	(*RollingAPIKeyRequest)(nil),          // 4: auth.RollingAPIKeyRequest
	(*RollingAPIKeyResponse)(nil),         // 5: auth.RollingAPIKeyResponse
	(*DisableAPIRequest)(nil),             // 6: auth.DisableAPIRequest
	(*GetRegisteredAPIRequest)(nil),       // 7: auth.GetRegisteredAPIRequest
	(*GetRegisteredAPIResponse)(nil),      // 8: auth.GetRegisteredAPIResponse
}
var file_api_proto_depIdxs = []int32{
	1, // 0: auth.GetRegisteredAPIResponse.api:type_name -> auth.API
	2, // 1: auth.APIService.RegisterAuthorizedAPI:input_type -> auth.RegisterAuthorizedAPIRequest
	4, // 2: auth.APIService.RollingAPIKey:input_type -> auth.RollingAPIKeyRequest
	6, // 3: auth.APIService.DisableAPI:input_type -> auth.DisableAPIRequest
	7, // 4: auth.APIService.GetRegisteredAPI:input_type -> auth.GetRegisteredAPIRequest
	3, // 5: auth.APIService.RegisterAuthorizedAPI:output_type -> auth.RegisterAuthorizedAPIResponse
	5, // 6: auth.APIService.RollingAPIKey:output_type -> auth.RollingAPIKeyResponse
	0, // 7: auth.APIService.DisableAPI:output_type -> auth.GenericResponse
	8, // 8: auth.APIService.GetRegisteredAPI:output_type -> auth.GetRegisteredAPIResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericResponse); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*API); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAuthorizedAPIRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAuthorizedAPIResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollingAPIKeyRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollingAPIKeyResponse); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableAPIRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegisteredAPIRequest); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegisteredAPIResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
