// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: petstore/v1/pet.proto

package v1

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

type PetStatus int32

const (
	PetStatus_PetStatus_UNDEFINED PetStatus = 0
	PetStatus_PetStatus_AVAILABLE PetStatus = 1
	PetStatus_PetStatus_PENDING   PetStatus = 2
	PetStatus_PetStatus_SOLD      PetStatus = 3
)

// Enum value maps for PetStatus.
var (
	PetStatus_name = map[int32]string{
		0: "PetStatus_UNDEFINED",
		1: "PetStatus_AVAILABLE",
		2: "PetStatus_PENDING",
		3: "PetStatus_SOLD",
	}
	PetStatus_value = map[string]int32{
		"PetStatus_UNDEFINED": 0,
		"PetStatus_AVAILABLE": 1,
		"PetStatus_PENDING":   2,
		"PetStatus_SOLD":      3,
	}
)

func (x PetStatus) Enum() *PetStatus {
	p := new(PetStatus)
	*p = x
	return p
}

func (x PetStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PetStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_petstore_v1_pet_proto_enumTypes[0].Descriptor()
}

func (PetStatus) Type() protoreflect.EnumType {
	return &file_petstore_v1_pet_proto_enumTypes[0]
}

func (x PetStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PetStatus.Descriptor instead.
func (PetStatus) EnumDescriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{0}
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId              int64  `protobuf:"varint,1,opt,name=petId,proto3" json:"petId,omitempty"`
	AdditionalMetadata string `protobuf:"bytes,2,opt,name=additionalMetadata,proto3" json:"additionalMetadata,omitempty"`
	File               []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{0}
}

func (x *UploadImageRequest) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

func (x *UploadImageRequest) GetAdditionalMetadata() string {
	if x != nil {
		return x.AdditionalMetadata
	}
	return ""
}

func (x *UploadImageRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type FindPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FindPet_OneOf:
	//	*FindPetRequest_ByStatus_
	//	*FindPetRequest_ById_
	//	*FindPetRequest_ByTags_
	FindPet_OneOf isFindPetRequest_FindPet_OneOf `protobuf_oneof:"FindPet_OneOf"`
}

func (x *FindPetRequest) Reset() {
	*x = FindPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetRequest) ProtoMessage() {}

func (x *FindPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetRequest.ProtoReflect.Descriptor instead.
func (*FindPetRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{1}
}

func (m *FindPetRequest) GetFindPet_OneOf() isFindPetRequest_FindPet_OneOf {
	if m != nil {
		return m.FindPet_OneOf
	}
	return nil
}

func (x *FindPetRequest) GetByStatus() *FindPetRequest_ByStatus {
	if x, ok := x.GetFindPet_OneOf().(*FindPetRequest_ByStatus_); ok {
		return x.ByStatus
	}
	return nil
}

func (x *FindPetRequest) GetById() *FindPetRequest_ByStatus {
	if x, ok := x.GetFindPet_OneOf().(*FindPetRequest_ById_); ok {
		return x.ById
	}
	return nil
}

// Deprecated: Do not use.
func (x *FindPetRequest) GetByTags() *FindPetRequest_ByTags {
	if x, ok := x.GetFindPet_OneOf().(*FindPetRequest_ByTags_); ok {
		return x.ByTags
	}
	return nil
}

type isFindPetRequest_FindPet_OneOf interface {
	isFindPetRequest_FindPet_OneOf()
}

type FindPetRequest_ByStatus_ struct {
	ByStatus *FindPetRequest_ByStatus `protobuf:"bytes,1,opt,name=byStatus,proto3,oneof"`
}

type FindPetRequest_ById_ struct {
	ById *FindPetRequest_ByStatus `protobuf:"bytes,2,opt,name=byId,proto3,oneof"`
}

type FindPetRequest_ByTags_ struct {
	// Deprecated: Do not use.
	ByTags *FindPetRequest_ByTags `protobuf:"bytes,3,opt,name=byTags,proto3,oneof"`
}

func (*FindPetRequest_ByStatus_) isFindPetRequest_FindPet_OneOf() {}

func (*FindPetRequest_ById_) isFindPetRequest_FindPet_OneOf() {}

func (*FindPetRequest_ByTags_) isFindPetRequest_FindPet_OneOf() {}

type FindPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FindPetResponse_OneOf:
	//	*FindPetResponse_Pets
	//	*FindPetResponse_Error
	FindPetResponse_OneOf isFindPetResponse_FindPetResponse_OneOf `protobuf_oneof:"FindPetResponse_OneOf"`
}

func (x *FindPetResponse) Reset() {
	*x = FindPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetResponse) ProtoMessage() {}

func (x *FindPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetResponse.ProtoReflect.Descriptor instead.
func (*FindPetResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{2}
}

func (m *FindPetResponse) GetFindPetResponse_OneOf() isFindPetResponse_FindPetResponse_OneOf {
	if m != nil {
		return m.FindPetResponse_OneOf
	}
	return nil
}

func (x *FindPetResponse) GetPets() *FindPetResponse_FindPetSuccessful {
	if x, ok := x.GetFindPetResponse_OneOf().(*FindPetResponse_Pets); ok {
		return x.Pets
	}
	return nil
}

func (x *FindPetResponse) GetError() *Response {
	if x, ok := x.GetFindPetResponse_OneOf().(*FindPetResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isFindPetResponse_FindPetResponse_OneOf interface {
	isFindPetResponse_FindPetResponse_OneOf()
}

type FindPetResponse_Pets struct {
	Pets *FindPetResponse_FindPetSuccessful `protobuf:"bytes,1,opt,name=pets,proto3,oneof"`
}

type FindPetResponse_Error struct {
	Error *Response `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*FindPetResponse_Pets) isFindPetResponse_FindPetResponse_OneOf() {}

func (*FindPetResponse_Error) isFindPetResponse_FindPetResponse_OneOf() {}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{3}
}

func (x *Category) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{4}
}

func (x *Tag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category  *Category `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	PhotoUrls []string  `protobuf:"bytes,3,rep,name=photoUrls,proto3" json:"photoUrls,omitempty"`
	Tags      []*Tag    `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Status    PetStatus `protobuf:"varint,5,opt,name=status,proto3,enum=petstore.v1.PetStatus" json:"status,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{5}
}

func (x *Pet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pet) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Pet) GetPhotoUrls() []string {
	if x != nil {
		return x.PhotoUrls
	}
	return nil
}

func (x *Pet) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Pet) GetStatus() PetStatus {
	if x != nil {
		return x.Status
	}
	return PetStatus_PetStatus_UNDEFINED
}

type FindPetRequest_ByStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status PetStatus `protobuf:"varint,1,opt,name=status,proto3,enum=petstore.v1.PetStatus" json:"status,omitempty"`
}

func (x *FindPetRequest_ByStatus) Reset() {
	*x = FindPetRequest_ByStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetRequest_ByStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetRequest_ByStatus) ProtoMessage() {}

func (x *FindPetRequest_ByStatus) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetRequest_ByStatus.ProtoReflect.Descriptor instead.
func (*FindPetRequest_ByStatus) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindPetRequest_ByStatus) GetStatus() PetStatus {
	if x != nil {
		return x.Status
	}
	return PetStatus_PetStatus_UNDEFINED
}

type FindPetRequest_ById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId int64 `protobuf:"varint,1,opt,name=petId,proto3" json:"petId,omitempty"`
}

func (x *FindPetRequest_ById) Reset() {
	*x = FindPetRequest_ById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetRequest_ById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetRequest_ById) ProtoMessage() {}

func (x *FindPetRequest_ById) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetRequest_ById.ProtoReflect.Descriptor instead.
func (*FindPetRequest_ById) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{1, 1}
}

func (x *FindPetRequest_ById) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

type FindPetRequest_ByTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *FindPetRequest_ByTags) Reset() {
	*x = FindPetRequest_ByTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetRequest_ByTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetRequest_ByTags) ProtoMessage() {}

func (x *FindPetRequest_ByTags) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetRequest_ByTags.ProtoReflect.Descriptor instead.
func (*FindPetRequest_ByTags) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{1, 2}
}

func (x *FindPetRequest_ByTags) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type FindPetResponse_FindPetSuccessful struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *FindPetResponse_FindPetSuccessful) Reset() {
	*x = FindPetResponse_FindPetSuccessful{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_pet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPetResponse_FindPetSuccessful) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPetResponse_FindPetSuccessful) ProtoMessage() {}

func (x *FindPetResponse_FindPetSuccessful) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_pet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPetResponse_FindPetSuccessful.ProtoReflect.Descriptor instead.
func (*FindPetResponse_FindPetSuccessful) Descriptor() ([]byte, []int) {
	return file_petstore_v1_pet_proto_rawDescGZIP(), []int{2, 0}
}

func (x *FindPetResponse_FindPetSuccessful) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

var File_petstore_v1_pet_proto protoreflect.FileDescriptor

var file_petstore_v1_pet_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xdb,
	0x02, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x42, 0x0a, 0x08, 0x62, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x08, 0x62, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x04, 0x62, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x04, 0x62, 0x79, 0x49,
	0x64, 0x12, 0x40, 0x0a, 0x06, 0x62, 0x79, 0x54, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42,
	0x79, 0x54, 0x61, 0x67, 0x73, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x06, 0x62, 0x79, 0x54,
	0x61, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x08, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x1c, 0x0a, 0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x1c, 0x0a,
	0x06, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x46,
	0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x5f, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x22, 0xda, 0x01, 0x0a,
	0x0f, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x04, 0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x39, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73,
	0x42, 0x17, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x24, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0x68, 0x0a, 0x09, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x13, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x53, 0x4f, 0x4c, 0x44, 0x10, 0x03, 0x32, 0xc5, 0x02,
	0x0a, 0x0a, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0b,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x07, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x79, 0x6e, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x2f, 0x70, 0x65,
	0x74, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petstore_v1_pet_proto_rawDescOnce sync.Once
	file_petstore_v1_pet_proto_rawDescData = file_petstore_v1_pet_proto_rawDesc
)

func file_petstore_v1_pet_proto_rawDescGZIP() []byte {
	file_petstore_v1_pet_proto_rawDescOnce.Do(func() {
		file_petstore_v1_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_petstore_v1_pet_proto_rawDescData)
	})
	return file_petstore_v1_pet_proto_rawDescData
}

var file_petstore_v1_pet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_petstore_v1_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_petstore_v1_pet_proto_goTypes = []interface{}{
	(PetStatus)(0),                            // 0: petstore.v1.PetStatus
	(*UploadImageRequest)(nil),                // 1: petstore.v1.UploadImageRequest
	(*FindPetRequest)(nil),                    // 2: petstore.v1.FindPetRequest
	(*FindPetResponse)(nil),                   // 3: petstore.v1.FindPetResponse
	(*Category)(nil),                          // 4: petstore.v1.Category
	(*Tag)(nil),                               // 5: petstore.v1.Tag
	(*Pet)(nil),                               // 6: petstore.v1.Pet
	(*FindPetRequest_ByStatus)(nil),           // 7: petstore.v1.FindPetRequest.ByStatus
	(*FindPetRequest_ById)(nil),               // 8: petstore.v1.FindPetRequest.ById
	(*FindPetRequest_ByTags)(nil),             // 9: petstore.v1.FindPetRequest.ByTags
	(*FindPetResponse_FindPetSuccessful)(nil), // 10: petstore.v1.FindPetResponse.FindPetSuccessful
	(*Response)(nil),                          // 11: petstore.v1.Response
}
var file_petstore_v1_pet_proto_depIdxs = []int32{
	7,  // 0: petstore.v1.FindPetRequest.byStatus:type_name -> petstore.v1.FindPetRequest.ByStatus
	7,  // 1: petstore.v1.FindPetRequest.byId:type_name -> petstore.v1.FindPetRequest.ByStatus
	9,  // 2: petstore.v1.FindPetRequest.byTags:type_name -> petstore.v1.FindPetRequest.ByTags
	10, // 3: petstore.v1.FindPetResponse.pets:type_name -> petstore.v1.FindPetResponse.FindPetSuccessful
	11, // 4: petstore.v1.FindPetResponse.error:type_name -> petstore.v1.Response
	4,  // 5: petstore.v1.Pet.category:type_name -> petstore.v1.Category
	5,  // 6: petstore.v1.Pet.tags:type_name -> petstore.v1.Tag
	0,  // 7: petstore.v1.Pet.status:type_name -> petstore.v1.PetStatus
	0,  // 8: petstore.v1.FindPetRequest.ByStatus.status:type_name -> petstore.v1.PetStatus
	6,  // 9: petstore.v1.FindPetResponse.FindPetSuccessful.pets:type_name -> petstore.v1.Pet
	1,  // 10: petstore.v1.PetService.UploadImage:input_type -> petstore.v1.UploadImageRequest
	6,  // 11: petstore.v1.PetService.CreatePet:input_type -> petstore.v1.Pet
	6,  // 12: petstore.v1.PetService.UpdatePet:input_type -> petstore.v1.Pet
	6,  // 13: petstore.v1.PetService.DeletePet:input_type -> petstore.v1.Pet
	2,  // 14: petstore.v1.PetService.FindPet:input_type -> petstore.v1.FindPetRequest
	11, // 15: petstore.v1.PetService.UploadImage:output_type -> petstore.v1.Response
	11, // 16: petstore.v1.PetService.CreatePet:output_type -> petstore.v1.Response
	11, // 17: petstore.v1.PetService.UpdatePet:output_type -> petstore.v1.Response
	11, // 18: petstore.v1.PetService.DeletePet:output_type -> petstore.v1.Response
	3,  // 19: petstore.v1.PetService.FindPet:output_type -> petstore.v1.FindPetResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_petstore_v1_pet_proto_init() }
func file_petstore_v1_pet_proto_init() {
	if File_petstore_v1_pet_proto != nil {
		return
	}
	file_petstore_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_petstore_v1_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_petstore_v1_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetRequest); i {
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
		file_petstore_v1_pet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetResponse); i {
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
		file_petstore_v1_pet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_petstore_v1_pet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_petstore_v1_pet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_petstore_v1_pet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetRequest_ByStatus); i {
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
		file_petstore_v1_pet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetRequest_ById); i {
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
		file_petstore_v1_pet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetRequest_ByTags); i {
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
		file_petstore_v1_pet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPetResponse_FindPetSuccessful); i {
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
	file_petstore_v1_pet_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FindPetRequest_ByStatus_)(nil),
		(*FindPetRequest_ById_)(nil),
		(*FindPetRequest_ByTags_)(nil),
	}
	file_petstore_v1_pet_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FindPetResponse_Pets)(nil),
		(*FindPetResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_petstore_v1_pet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petstore_v1_pet_proto_goTypes,
		DependencyIndexes: file_petstore_v1_pet_proto_depIdxs,
		EnumInfos:         file_petstore_v1_pet_proto_enumTypes,
		MessageInfos:      file_petstore_v1_pet_proto_msgTypes,
	}.Build()
	File_petstore_v1_pet_proto = out.File
	file_petstore_v1_pet_proto_rawDesc = nil
	file_petstore_v1_pet_proto_goTypes = nil
	file_petstore_v1_pet_proto_depIdxs = nil
}
