// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: api.proto

// 定义包名

package v1

import (
	context "context"
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

// 定义MapperRegisterRequest消息
type MapperRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name字段
	Mapper *MapperInfo `protobuf:"bytes,1,opt,name=mapper,proto3" json:"mapper,omitempty"`
}

func (x *MapperRegisterRequest) Reset() {
	*x = MapperRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapperRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapperRegisterRequest) ProtoMessage() {}

func (x *MapperRegisterRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MapperRegisterRequest.ProtoReflect.Descriptor instead.
func (*MapperRegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *MapperRegisterRequest) GetMapper() *MapperInfo {
	if x != nil {
		return x.Mapper
	}
	return nil
}

// 定义MapperRegisterResponse消息
type MapperRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MapperRegisterResponse) Reset() {
	*x = MapperRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapperRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapperRegisterResponse) ProtoMessage() {}

func (x *MapperRegisterResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MapperRegisterResponse.ProtoReflect.Descriptor instead.
func (*MapperRegisterResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type MapperInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	ApiVersion string `protobuf:"bytes,3,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Protocol   string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address    []byte `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *MapperInfo) Reset() {
	*x = MapperInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapperInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapperInfo) ProtoMessage() {}

func (x *MapperInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MapperInfo.ProtoReflect.Descriptor instead.
func (*MapperInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *MapperInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MapperInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MapperInfo) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *MapperInfo) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *MapperInfo) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *MapperInfo) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type ReportDeviceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName     string        `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	ReportedDevice *DeviceStatus `protobuf:"bytes,2,opt,name=reportedDevice,proto3" json:"reportedDevice,omitempty"`
}

func (x *ReportDeviceStatusRequest) Reset() {
	*x = ReportDeviceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDeviceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDeviceStatusRequest) ProtoMessage() {}

func (x *ReportDeviceStatusRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReportDeviceStatusRequest.ProtoReflect.Descriptor instead.
func (*ReportDeviceStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *ReportDeviceStatusRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *ReportDeviceStatusRequest) GetReportedDevice() *DeviceStatus {
	if x != nil {
		return x.ReportedDevice
	}
	return nil
}

type DeviceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Twins []*Twin `protobuf:"bytes,1,rep,name=twins,proto3" json:"twins,omitempty"`
	State string  `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *DeviceStatus) Reset() {
	*x = DeviceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatus) ProtoMessage() {}

func (x *DeviceStatus) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeviceStatus.ProtoReflect.Descriptor instead.
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceStatus) GetTwins() []*Twin {
	if x != nil {
		return x.Twins
	}
	return nil
}

func (x *DeviceStatus) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type Twin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyName string        `protobuf:"bytes,1,opt,name=propertyName,proto3" json:"propertyName,omitempty"`
	Desired      *TwinProperty `protobuf:"bytes,2,opt,name=desired,proto3" json:"desired,omitempty"`
	Reported     *TwinProperty `protobuf:"bytes,3,opt,name=reported,proto3" json:"reported,omitempty"`
}

func (x *Twin) Reset() {
	*x = Twin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Twin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Twin) ProtoMessage() {}

func (x *Twin) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Twin.ProtoReflect.Descriptor instead.
func (*Twin) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Twin) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *Twin) GetDesired() *TwinProperty {
	if x != nil {
		return x.Desired
	}
	return nil
}

func (x *Twin) GetReported() *TwinProperty {
	if x != nil {
		return x.Reported
	}
	return nil
}

type TwinProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    string            `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TwinProperty) Reset() {
	*x = TwinProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwinProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwinProperty) ProtoMessage() {}

func (x *TwinProperty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TwinProperty.ProtoReflect.Descriptor instead.
func (*TwinProperty) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *TwinProperty) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TwinProperty) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ReportDeviceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportDeviceStatusResponse) Reset() {
	*x = ReportDeviceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDeviceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDeviceStatusResponse) ProtoMessage() {}

func (x *ReportDeviceStatusResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReportDeviceStatusResponse.ProtoReflect.Descriptor instead.
func (*ReportDeviceStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22,
	0x3f, 0x0a, 0x15, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x22, 0x18, 0x0a, 0x16, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0a, 0x4d,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x75, 0x0a, 0x19, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x0c, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x74,
	0x77, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x77, 0x69, 0x6e, 0x52, 0x05, 0x74, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x84, 0x01, 0x0a, 0x04, 0x54, 0x77, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x77, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x77, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x0c, 0x54, 0x77, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x77, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1c, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb8, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x49, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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
	(*MapperRegisterRequest)(nil),      // 0: v1.MapperRegisterRequest
	(*MapperRegisterResponse)(nil),     // 1: v1.MapperRegisterResponse
	(*MapperInfo)(nil),                 // 2: v1.MapperInfo
	(*ReportDeviceStatusRequest)(nil),  // 3: v1.ReportDeviceStatusRequest
	(*DeviceStatus)(nil),               // 4: v1.DeviceStatus
	(*Twin)(nil),                       // 5: v1.Twin
	(*TwinProperty)(nil),               // 6: v1.TwinProperty
	(*ReportDeviceStatusResponse)(nil), // 7: v1.ReportDeviceStatusResponse
	nil,                                // 8: v1.TwinProperty.MetadataEntry
}
var file_api_proto_depIdxs = []int32{
	2, // 0: v1.MapperRegisterRequest.mapper:type_name -> v1.MapperInfo
	4, // 1: v1.ReportDeviceStatusRequest.reportedDevice:type_name -> v1.DeviceStatus
	5, // 2: v1.DeviceStatus.twins:type_name -> v1.Twin
	6, // 3: v1.Twin.desired:type_name -> v1.TwinProperty
	6, // 4: v1.Twin.reported:type_name -> v1.TwinProperty
	8, // 5: v1.TwinProperty.metadata:type_name -> v1.TwinProperty.MetadataEntry
	0, // 6: v1.DeviceManagerService.MapperRegister:input_type -> v1.MapperRegisterRequest
	3, // 7: v1.DeviceManagerService.ReportDeviceStatus:input_type -> v1.ReportDeviceStatusRequest
	1, // 8: v1.DeviceManagerService.MapperRegister:output_type -> v1.MapperRegisterResponse
	7, // 9: v1.DeviceManagerService.ReportDeviceStatus:output_type -> v1.ReportDeviceStatusResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapperRegisterRequest); i {
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
			switch v := v.(*MapperRegisterResponse); i {
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
			switch v := v.(*MapperInfo); i {
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
			switch v := v.(*ReportDeviceStatusRequest); i {
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
			switch v := v.(*DeviceStatus); i {
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
			switch v := v.(*Twin); i {
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
			switch v := v.(*TwinProperty); i {
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
			switch v := v.(*ReportDeviceStatusResponse); i {
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceManagerServiceClient is the client API for DeviceManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceManagerServiceClient interface {
	// 定义MapperRegister方法，接受MapperRegister消息， 并返回MapperRegister消息
	MapperRegister(ctx context.Context, in *MapperRegisterRequest, opts ...grpc.CallOption) (*MapperRegisterResponse, error)
	ReportDeviceStatus(ctx context.Context, in *ReportDeviceStatusRequest, opts ...grpc.CallOption) (*ReportDeviceStatusResponse, error)
}

type deviceManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceManagerServiceClient(cc grpc.ClientConnInterface) DeviceManagerServiceClient {
	return &deviceManagerServiceClient{cc}
}

func (c *deviceManagerServiceClient) MapperRegister(ctx context.Context, in *MapperRegisterRequest, opts ...grpc.CallOption) (*MapperRegisterResponse, error) {
	out := new(MapperRegisterResponse)
	err := c.cc.Invoke(ctx, "/v1.DeviceManagerService/MapperRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManagerServiceClient) ReportDeviceStatus(ctx context.Context, in *ReportDeviceStatusRequest, opts ...grpc.CallOption) (*ReportDeviceStatusResponse, error) {
	out := new(ReportDeviceStatusResponse)
	err := c.cc.Invoke(ctx, "/v1.DeviceManagerService/ReportDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceManagerServiceServer is the server API for DeviceManagerService service.
type DeviceManagerServiceServer interface {
	// 定义MapperRegister方法，接受MapperRegister消息， 并返回MapperRegister消息
	MapperRegister(context.Context, *MapperRegisterRequest) (*MapperRegisterResponse, error)
	ReportDeviceStatus(context.Context, *ReportDeviceStatusRequest) (*ReportDeviceStatusResponse, error)
}

// UnimplementedDeviceManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceManagerServiceServer struct {
}

func (*UnimplementedDeviceManagerServiceServer) MapperRegister(context.Context, *MapperRegisterRequest) (*MapperRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapperRegister not implemented")
}
func (*UnimplementedDeviceManagerServiceServer) ReportDeviceStatus(context.Context, *ReportDeviceStatusRequest) (*ReportDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportDeviceStatus not implemented")
}

func RegisterDeviceManagerServiceServer(s *grpc.Server, srv DeviceManagerServiceServer) {
	s.RegisterService(&_DeviceManagerService_serviceDesc, srv)
}

func _DeviceManagerService_MapperRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapperRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServiceServer).MapperRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DeviceManagerService/MapperRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServiceServer).MapperRegister(ctx, req.(*MapperRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManagerService_ReportDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServiceServer).ReportDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DeviceManagerService/ReportDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServiceServer).ReportDeviceStatus(ctx, req.(*ReportDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DeviceManagerService",
	HandlerType: (*DeviceManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MapperRegister",
			Handler:    _DeviceManagerService_MapperRegister_Handler,
		},
		{
			MethodName: "ReportDeviceStatus",
			Handler:    _DeviceManagerService_ReportDeviceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func NewMapperClient(cc grpc.ClientConnInterface) DeviceManagerServiceClient {
	return &deviceManagerServiceClient{cc}
}