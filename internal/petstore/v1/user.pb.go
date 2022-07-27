// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: petstore/v1/user.proto

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

type UserStatus int32

const (
	UserStatus_UserStatus_UNDEFINED UserStatus = 0
	UserStatus_UserStatus_ACTIVE    UserStatus = 1
	UserStatus_UserStatus_INACTIVE  UserStatus = 2
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "UserStatus_UNDEFINED",
		1: "UserStatus_ACTIVE",
		2: "UserStatus_INACTIVE",
	}
	UserStatus_value = map[string]int32{
		"UserStatus_UNDEFINED": 0,
		"UserStatus_ACTIVE":    1,
		"UserStatus_INACTIVE":  2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_petstore_v1_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_petstore_v1_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{0}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LoginResponse_OneOf:
	//	*LoginResponse_Successful
	//	*LoginResponse_Error
	LoginResponse_OneOf isLoginResponse_LoginResponse_OneOf `protobuf_oneof:"LoginResponse_OneOf"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{1}
}

func (m *LoginResponse) GetLoginResponse_OneOf() isLoginResponse_LoginResponse_OneOf {
	if m != nil {
		return m.LoginResponse_OneOf
	}
	return nil
}

func (x *LoginResponse) GetSuccessful() *LoginResponse_LoginSuccessful {
	if x, ok := x.GetLoginResponse_OneOf().(*LoginResponse_Successful); ok {
		return x.Successful
	}
	return nil
}

func (x *LoginResponse) GetError() *Response {
	if x, ok := x.GetLoginResponse_OneOf().(*LoginResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isLoginResponse_LoginResponse_OneOf interface {
	isLoginResponse_LoginResponse_OneOf()
}

type LoginResponse_Successful struct {
	Successful *LoginResponse_LoginSuccessful `protobuf:"bytes,1,opt,name=successful,proto3,oneof"`
}

type LoginResponse_Error struct {
	Error *Response `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*LoginResponse_Successful) isLoginResponse_LoginResponse_OneOf() {}

func (*LoginResponse_Error) isLoginResponse_LoginResponse_OneOf() {}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{2}
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CreateUserRequest_OneOf:
	//	*CreateUserRequest_User
	//	*CreateUserRequest_WithArray_
	//	*CreateUserRequest_WithList
	CreateUserRequest_OneOf isCreateUserRequest_CreateUserRequest_OneOf `protobuf_oneof:"CreateUserRequest_OneOf"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{3}
}

func (m *CreateUserRequest) GetCreateUserRequest_OneOf() isCreateUserRequest_CreateUserRequest_OneOf {
	if m != nil {
		return m.CreateUserRequest_OneOf
	}
	return nil
}

func (x *CreateUserRequest) GetUser() *User {
	if x, ok := x.GetCreateUserRequest_OneOf().(*CreateUserRequest_User); ok {
		return x.User
	}
	return nil
}

func (x *CreateUserRequest) GetWithArray() *CreateUserRequest_WithArray {
	if x, ok := x.GetCreateUserRequest_OneOf().(*CreateUserRequest_WithArray_); ok {
		return x.WithArray
	}
	return nil
}

func (x *CreateUserRequest) GetWithList() *CreateUserRequest_WithArray {
	if x, ok := x.GetCreateUserRequest_OneOf().(*CreateUserRequest_WithList); ok {
		return x.WithList
	}
	return nil
}

type isCreateUserRequest_CreateUserRequest_OneOf interface {
	isCreateUserRequest_CreateUserRequest_OneOf()
}

type CreateUserRequest_User struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type CreateUserRequest_WithArray_ struct {
	WithArray *CreateUserRequest_WithArray `protobuf:"bytes,2,opt,name=withArray,proto3,oneof"`
}

type CreateUserRequest_WithList struct {
	WithList *CreateUserRequest_WithArray `protobuf:"bytes,3,opt,name=withList,proto3,oneof"`
}

func (*CreateUserRequest_User) isCreateUserRequest_CreateUserRequest_OneOf() {}

func (*CreateUserRequest_WithArray_) isCreateUserRequest_CreateUserRequest_OneOf() {}

func (*CreateUserRequest_WithList) isCreateUserRequest_CreateUserRequest_OneOf() {}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GetUserRequest_OneOf:
	//	*GetUserRequest_ByUsername_
	GetUserRequest_OneOf isGetUserRequest_GetUserRequest_OneOf `protobuf_oneof:"GetUserRequest_OneOf"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{4}
}

func (m *GetUserRequest) GetGetUserRequest_OneOf() isGetUserRequest_GetUserRequest_OneOf {
	if m != nil {
		return m.GetUserRequest_OneOf
	}
	return nil
}

func (x *GetUserRequest) GetByUsername() *GetUserRequest_ByUsername {
	if x, ok := x.GetGetUserRequest_OneOf().(*GetUserRequest_ByUsername_); ok {
		return x.ByUsername
	}
	return nil
}

type isGetUserRequest_GetUserRequest_OneOf interface {
	isGetUserRequest_GetUserRequest_OneOf()
}

type GetUserRequest_ByUsername_ struct {
	ByUsername *GetUserRequest_ByUsername `protobuf:"bytes,1,opt,name=byUsername,proto3,oneof"`
}

func (*GetUserRequest_ByUsername_) isGetUserRequest_GetUserRequest_OneOf() {}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GetUserResponse_OneOf:
	//	*GetUserResponse_User
	//	*GetUserResponse_Error
	GetUserResponse_OneOf isGetUserResponse_GetUserResponse_OneOf `protobuf_oneof:"GetUserResponse_OneOf"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{5}
}

func (m *GetUserResponse) GetGetUserResponse_OneOf() isGetUserResponse_GetUserResponse_OneOf {
	if m != nil {
		return m.GetUserResponse_OneOf
	}
	return nil
}

func (x *GetUserResponse) GetUser() *User {
	if x, ok := x.GetGetUserResponse_OneOf().(*GetUserResponse_User); ok {
		return x.User
	}
	return nil
}

func (x *GetUserResponse) GetError() *Response {
	if x, ok := x.GetGetUserResponse_OneOf().(*GetUserResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetUserResponse_GetUserResponse_OneOf interface {
	isGetUserResponse_GetUserResponse_OneOf()
}

type GetUserResponse_User struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type GetUserResponse_Error struct {
	Error *Response `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetUserResponse_User) isGetUserResponse_GetUserResponse_OneOf() {}

func (*GetUserResponse_Error) isGetUserResponse_GetUserResponse_OneOf() {}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName   string     `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string     `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email       string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password    string     `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber string     `protobuf:"bytes,7,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	UserStatus  UserStatus `protobuf:"varint,8,opt,name=userStatus,proto3,enum=petstore.v1.UserStatus" json:"userStatus,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetUserStatus() UserStatus {
	if x != nil {
		return x.UserStatus
	}
	return UserStatus_UserStatus_UNDEFINED
}

type LoginResponse_LoginSuccessful struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenExpires string `protobuf:"bytes,1,opt,name=tokenExpires,proto3" json:"tokenExpires,omitempty"`
	RateLimit    int32  `protobuf:"varint,2,opt,name=rateLimit,proto3" json:"rateLimit,omitempty"`
}

func (x *LoginResponse_LoginSuccessful) Reset() {
	*x = LoginResponse_LoginSuccessful{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse_LoginSuccessful) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse_LoginSuccessful) ProtoMessage() {}

func (x *LoginResponse_LoginSuccessful) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse_LoginSuccessful.ProtoReflect.Descriptor instead.
func (*LoginResponse_LoginSuccessful) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LoginResponse_LoginSuccessful) GetTokenExpires() string {
	if x != nil {
		return x.TokenExpires
	}
	return ""
}

func (x *LoginResponse_LoginSuccessful) GetRateLimit() int32 {
	if x != nil {
		return x.RateLimit
	}
	return 0
}

type CreateUserRequest_WithArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *CreateUserRequest_WithArray) Reset() {
	*x = CreateUserRequest_WithArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest_WithArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest_WithArray) ProtoMessage() {}

func (x *CreateUserRequest_WithArray) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest_WithArray.ProtoReflect.Descriptor instead.
func (*CreateUserRequest_WithArray) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CreateUserRequest_WithArray) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserRequest_ByUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetUserRequest_ByUsername) Reset() {
	*x = GetUserRequest_ByUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest_ByUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest_ByUsername) ProtoMessage() {}

func (x *GetUserRequest_ByUsername) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest_ByUsername.ProtoReflect.Descriptor instead.
func (*GetUserRequest_ByUsername) Descriptor() ([]byte, []int) {
	return file_petstore_v1_user_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetUserRequest_ByUsername) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_petstore_v1_user_proto protoreflect.FileDescriptor

var file_petstore_v1_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xf8, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x53, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x15, 0x0a, 0x13, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4f, 0x6e, 0x65,
	0x4f, 0x66, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x48, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x41, 0x72, 0x72, 0x61, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48,
	0x00, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x46, 0x0a, 0x08,
	0x77, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57,
	0x69, 0x74, 0x68, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x34, 0x0a, 0x09, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x19, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x62, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x28, 0x0a, 0x0a, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x16, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4f,
	0x6e, 0x65, 0x4f, 0x66, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x22, 0xf9, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x56, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x32, 0x91, 0x03,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x65,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x61, 0x79, 0x6e, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f,
	0x70, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petstore_v1_user_proto_rawDescOnce sync.Once
	file_petstore_v1_user_proto_rawDescData = file_petstore_v1_user_proto_rawDesc
)

func file_petstore_v1_user_proto_rawDescGZIP() []byte {
	file_petstore_v1_user_proto_rawDescOnce.Do(func() {
		file_petstore_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_petstore_v1_user_proto_rawDescData)
	})
	return file_petstore_v1_user_proto_rawDescData
}

var file_petstore_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_petstore_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_petstore_v1_user_proto_goTypes = []interface{}{
	(UserStatus)(0),                       // 0: petstore.v1.UserStatus
	(*LoginRequest)(nil),                  // 1: petstore.v1.LoginRequest
	(*LoginResponse)(nil),                 // 2: petstore.v1.LoginResponse
	(*LogoutRequest)(nil),                 // 3: petstore.v1.LogoutRequest
	(*CreateUserRequest)(nil),             // 4: petstore.v1.CreateUserRequest
	(*GetUserRequest)(nil),                // 5: petstore.v1.GetUserRequest
	(*GetUserResponse)(nil),               // 6: petstore.v1.GetUserResponse
	(*User)(nil),                          // 7: petstore.v1.User
	(*LoginResponse_LoginSuccessful)(nil), // 8: petstore.v1.LoginResponse.LoginSuccessful
	(*CreateUserRequest_WithArray)(nil),   // 9: petstore.v1.CreateUserRequest.WithArray
	(*GetUserRequest_ByUsername)(nil),     // 10: petstore.v1.GetUserRequest.ByUsername
	(*Response)(nil),                      // 11: petstore.v1.Response
}
var file_petstore_v1_user_proto_depIdxs = []int32{
	8,  // 0: petstore.v1.LoginResponse.successful:type_name -> petstore.v1.LoginResponse.LoginSuccessful
	11, // 1: petstore.v1.LoginResponse.error:type_name -> petstore.v1.Response
	7,  // 2: petstore.v1.CreateUserRequest.user:type_name -> petstore.v1.User
	9,  // 3: petstore.v1.CreateUserRequest.withArray:type_name -> petstore.v1.CreateUserRequest.WithArray
	9,  // 4: petstore.v1.CreateUserRequest.withList:type_name -> petstore.v1.CreateUserRequest.WithArray
	10, // 5: petstore.v1.GetUserRequest.byUsername:type_name -> petstore.v1.GetUserRequest.ByUsername
	7,  // 6: petstore.v1.GetUserResponse.user:type_name -> petstore.v1.User
	11, // 7: petstore.v1.GetUserResponse.error:type_name -> petstore.v1.Response
	0,  // 8: petstore.v1.User.userStatus:type_name -> petstore.v1.UserStatus
	7,  // 9: petstore.v1.CreateUserRequest.WithArray.users:type_name -> petstore.v1.User
	4,  // 10: petstore.v1.UserService.CreateUser:input_type -> petstore.v1.CreateUserRequest
	5,  // 11: petstore.v1.UserService.GetUser:input_type -> petstore.v1.GetUserRequest
	7,  // 12: petstore.v1.UserService.UpdateUser:input_type -> petstore.v1.User
	7,  // 13: petstore.v1.UserService.DeleteUser:input_type -> petstore.v1.User
	1,  // 14: petstore.v1.UserService.Login:input_type -> petstore.v1.LoginRequest
	3,  // 15: petstore.v1.UserService.Logout:input_type -> petstore.v1.LogoutRequest
	11, // 16: petstore.v1.UserService.CreateUser:output_type -> petstore.v1.Response
	6,  // 17: petstore.v1.UserService.GetUser:output_type -> petstore.v1.GetUserResponse
	11, // 18: petstore.v1.UserService.UpdateUser:output_type -> petstore.v1.Response
	11, // 19: petstore.v1.UserService.DeleteUser:output_type -> petstore.v1.Response
	2,  // 20: petstore.v1.UserService.Login:output_type -> petstore.v1.LoginResponse
	11, // 21: petstore.v1.UserService.Logout:output_type -> petstore.v1.Response
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_petstore_v1_user_proto_init() }
func file_petstore_v1_user_proto_init() {
	if File_petstore_v1_user_proto != nil {
		return
	}
	file_petstore_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_petstore_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_petstore_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_petstore_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_petstore_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_petstore_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_petstore_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_petstore_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_petstore_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse_LoginSuccessful); i {
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
		file_petstore_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest_WithArray); i {
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
		file_petstore_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest_ByUsername); i {
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
	file_petstore_v1_user_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LoginResponse_Successful)(nil),
		(*LoginResponse_Error)(nil),
	}
	file_petstore_v1_user_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CreateUserRequest_User)(nil),
		(*CreateUserRequest_WithArray_)(nil),
		(*CreateUserRequest_WithList)(nil),
	}
	file_petstore_v1_user_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetUserRequest_ByUsername_)(nil),
	}
	file_petstore_v1_user_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GetUserResponse_User)(nil),
		(*GetUserResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_petstore_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petstore_v1_user_proto_goTypes,
		DependencyIndexes: file_petstore_v1_user_proto_depIdxs,
		EnumInfos:         file_petstore_v1_user_proto_enumTypes,
		MessageInfos:      file_petstore_v1_user_proto_msgTypes,
	}.Build()
	File_petstore_v1_user_proto = out.File
	file_petstore_v1_user_proto_rawDesc = nil
	file_petstore_v1_user_proto_goTypes = nil
	file_petstore_v1_user_proto_depIdxs = nil
}
