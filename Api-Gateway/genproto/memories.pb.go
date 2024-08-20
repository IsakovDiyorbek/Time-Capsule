// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: memories.proto

package genproto

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

type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Date        string   `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Tags        []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Latitude    float32  `protobuf:"fixed32,7,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float32  `protobuf:"fixed32,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	PlaceName   string   `protobuf:"bytes,9,opt,name=place_name,json=placeName,proto3" json:"place_name,omitempty"`
	Privacy     string   `protobuf:"bytes,10,opt,name=privacy,proto3" json:"privacy,omitempty"`
	CreatedAt   string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{0}
}

func (x *Memory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Memory) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Memory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Memory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Memory) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Memory) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Memory) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Memory) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Memory) GetPlaceName() string {
	if x != nil {
		return x.PlaceName
	}
	return ""
}

func (x *Memory) GetPrivacy() string {
	if x != nil {
		return x.Privacy
	}
	return ""
}

func (x *Memory) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Memory) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AddMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Date        string   `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Tags        []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Latitude    float32  `protobuf:"fixed32,7,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float32  `protobuf:"fixed32,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	PlaceName   string   `protobuf:"bytes,9,opt,name=place_name,json=placeName,proto3" json:"place_name,omitempty"`
	Privacy     string   `protobuf:"bytes,10,opt,name=privacy,proto3" json:"privacy,omitempty"`
}

func (x *AddMemoryRequest) Reset() {
	*x = AddMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMemoryRequest) ProtoMessage() {}

func (x *AddMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMemoryRequest.ProtoReflect.Descriptor instead.
func (*AddMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{1}
}

func (x *AddMemoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddMemoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddMemoryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddMemoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddMemoryRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *AddMemoryRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AddMemoryRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *AddMemoryRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *AddMemoryRequest) GetPlaceName() string {
	if x != nil {
		return x.PlaceName
	}
	return ""
}

func (x *AddMemoryRequest) GetPrivacy() string {
	if x != nil {
		return x.Privacy
	}
	return ""
}

type AddMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddMemoryResponse) Reset() {
	*x = AddMemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMemoryResponse) ProtoMessage() {}

func (x *AddMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMemoryResponse.ProtoReflect.Descriptor instead.
func (*AddMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{2}
}

type GetAllMemoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit     string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	StartDate string `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Tags      string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetAllMemoriesRequest) Reset() {
	*x = GetAllMemoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMemoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMemoriesRequest) ProtoMessage() {}

func (x *GetAllMemoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMemoriesRequest.ProtoReflect.Descriptor instead.
func (*GetAllMemoriesRequest) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMemoriesRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetAllMemoriesRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *GetAllMemoriesRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetAllMemoriesRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *GetAllMemoriesRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type GetAllMemoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memories []*Memory `protobuf:"bytes,1,rep,name=memories,proto3" json:"memories,omitempty"`
}

func (x *GetAllMemoriesResponse) Reset() {
	*x = GetAllMemoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMemoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMemoriesResponse) ProtoMessage() {}

func (x *GetAllMemoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMemoriesResponse.ProtoReflect.Descriptor instead.
func (*GetAllMemoriesResponse) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllMemoriesResponse) GetMemories() []*Memory {
	if x != nil {
		return x.Memories
	}
	return nil
}

type GetMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMemoryRequest) Reset() {
	*x = GetMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoryRequest) ProtoMessage() {}

func (x *GetMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoryRequest.ProtoReflect.Descriptor instead.
func (*GetMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{5}
}

func (x *GetMemoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memory *Memory `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *GetMemoryResponse) Reset() {
	*x = GetMemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoryResponse) ProtoMessage() {}

func (x *GetMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoryResponse.ProtoReflect.Descriptor instead.
func (*GetMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{6}
}

func (x *GetMemoryResponse) GetMemory() *Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

type UpdateMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tags  []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateMemoryRequest) Reset() {
	*x = UpdateMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoryRequest) ProtoMessage() {}

func (x *UpdateMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMemoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMemoryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMemoryRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMemoryResponse) Reset() {
	*x = UpdateMemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoryResponse) ProtoMessage() {}

func (x *UpdateMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{8}
}

type DeleteMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMemoryRequest) Reset() {
	*x = DeleteMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoryRequest) ProtoMessage() {}

func (x *DeleteMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMemoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteMemoryResponse) Reset() {
	*x = DeleteMemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoryResponse) ProtoMessage() {}

func (x *DeleteMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteMemoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ShareMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId   string   `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	SharedWith []string `protobuf:"bytes,2,rep,name=shared_with,json=sharedWith,proto3" json:"shared_with,omitempty"`
}

func (x *ShareMemoryRequest) Reset() {
	*x = ShareMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareMemoryRequest) ProtoMessage() {}

func (x *ShareMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareMemoryRequest.ProtoReflect.Descriptor instead.
func (*ShareMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{11}
}

func (x *ShareMemoryRequest) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *ShareMemoryRequest) GetSharedWith() []string {
	if x != nil {
		return x.SharedWith
	}
	return nil
}

type ShareMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShareMemoryResponse) Reset() {
	*x = ShareMemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memories_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareMemoryResponse) ProtoMessage() {}

func (x *ShareMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memories_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareMemoryResponse.ProtoReflect.Descriptor instead.
func (*ShareMemoryResponse) Descriptor() ([]byte, []int) {
	return file_memories_proto_rawDescGZIP(), []int{12}
}

var File_memories_proto protoreflect.FileDescriptor

var file_memories_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xc2, 0x02, 0x0a, 0x06, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x8e, 0x02, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x22, 0x13, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x22, 0x4f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xda, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1a,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x1c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_memories_proto_rawDescOnce sync.Once
	file_memories_proto_rawDescData = file_memories_proto_rawDesc
)

func file_memories_proto_rawDescGZIP() []byte {
	file_memories_proto_rawDescOnce.Do(func() {
		file_memories_proto_rawDescData = protoimpl.X.CompressGZIP(file_memories_proto_rawDescData)
	})
	return file_memories_proto_rawDescData
}

var file_memories_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_memories_proto_goTypes = []interface{}{
	(*Memory)(nil),                 // 0: timeline.Memory
	(*AddMemoryRequest)(nil),       // 1: timeline.AddMemoryRequest
	(*AddMemoryResponse)(nil),      // 2: timeline.AddMemoryResponse
	(*GetAllMemoriesRequest)(nil),  // 3: timeline.GetAllMemoriesRequest
	(*GetAllMemoriesResponse)(nil), // 4: timeline.GetAllMemoriesResponse
	(*GetMemoryRequest)(nil),       // 5: timeline.GetMemoryRequest
	(*GetMemoryResponse)(nil),      // 6: timeline.GetMemoryResponse
	(*UpdateMemoryRequest)(nil),    // 7: timeline.UpdateMemoryRequest
	(*UpdateMemoryResponse)(nil),   // 8: timeline.UpdateMemoryResponse
	(*DeleteMemoryRequest)(nil),    // 9: timeline.DeleteMemoryRequest
	(*DeleteMemoryResponse)(nil),   // 10: timeline.DeleteMemoryResponse
	(*ShareMemoryRequest)(nil),     // 11: timeline.ShareMemoryRequest
	(*ShareMemoryResponse)(nil),    // 12: timeline.ShareMemoryResponse
}
var file_memories_proto_depIdxs = []int32{
	0,  // 0: timeline.GetAllMemoriesResponse.memories:type_name -> timeline.Memory
	0,  // 1: timeline.GetMemoryResponse.memory:type_name -> timeline.Memory
	1,  // 2: timeline.MemoryService.AddMemory:input_type -> timeline.AddMemoryRequest
	5,  // 3: timeline.MemoryService.GetMemory:input_type -> timeline.GetMemoryRequest
	3,  // 4: timeline.MemoryService.GetAllMemories:input_type -> timeline.GetAllMemoriesRequest
	7,  // 5: timeline.MemoryService.UpdateMemory:input_type -> timeline.UpdateMemoryRequest
	9,  // 6: timeline.MemoryService.DeleteMemory:input_type -> timeline.DeleteMemoryRequest
	11, // 7: timeline.MemoryService.ShareMemory:input_type -> timeline.ShareMemoryRequest
	2,  // 8: timeline.MemoryService.AddMemory:output_type -> timeline.AddMemoryResponse
	6,  // 9: timeline.MemoryService.GetMemory:output_type -> timeline.GetMemoryResponse
	4,  // 10: timeline.MemoryService.GetAllMemories:output_type -> timeline.GetAllMemoriesResponse
	8,  // 11: timeline.MemoryService.UpdateMemory:output_type -> timeline.UpdateMemoryResponse
	10, // 12: timeline.MemoryService.DeleteMemory:output_type -> timeline.DeleteMemoryResponse
	12, // 13: timeline.MemoryService.ShareMemory:output_type -> timeline.ShareMemoryResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_memories_proto_init() }
func file_memories_proto_init() {
	if File_memories_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memory); i {
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
		file_memories_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMemoryRequest); i {
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
		file_memories_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMemoryResponse); i {
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
		file_memories_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMemoriesRequest); i {
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
		file_memories_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMemoriesResponse); i {
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
		file_memories_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemoryRequest); i {
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
		file_memories_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemoryResponse); i {
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
		file_memories_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemoryRequest); i {
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
		file_memories_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemoryResponse); i {
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
		file_memories_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemoryRequest); i {
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
		file_memories_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemoryResponse); i {
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
		file_memories_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareMemoryRequest); i {
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
		file_memories_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareMemoryResponse); i {
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
			RawDescriptor: file_memories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memories_proto_goTypes,
		DependencyIndexes: file_memories_proto_depIdxs,
		MessageInfos:      file_memories_proto_msgTypes,
	}.Build()
	File_memories_proto = out.File
	file_memories_proto_rawDesc = nil
	file_memories_proto_goTypes = nil
	file_memories_proto_depIdxs = nil
}