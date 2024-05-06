// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/chat/v1/user.proto

package chatpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User_Type int32

const (
	// Default value for the enum. DO NOT USE.
	User_TYPE_UNSPECIFIED User_Type = 0
	// Human user.
	User_HUMAN User_Type = 1
	// Chat app user.
	User_BOT User_Type = 2
)

// Enum value maps for User_Type.
var (
	User_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "HUMAN",
		2: "BOT",
	}
	User_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"HUMAN":            1,
		"BOT":              2,
	}
)

func (x User_Type) Enum() *User_Type {
	p := new(User_Type)
	*p = x
	return p
}

func (x User_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_chat_v1_user_proto_enumTypes[0].Descriptor()
}

func (User_Type) Type() protoreflect.EnumType {
	return &file_google_chat_v1_user_proto_enumTypes[0]
}

func (x User_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Type.Descriptor instead.
func (User_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_chat_v1_user_proto_rawDescGZIP(), []int{0, 0}
}

// A user in Google Chat.
// When returned as an output from a request, if your Chat app [authenticates as
// a
// user](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user),
// the output for a `User` resource only populates the user's `name` and `type`.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name for a Google Chat [user][google.chat.v1.User].
	//
	// Format: `users/{user}`. `users/app` can be used as an alias for the calling
	// app [bot][google.chat.v1.User.Type.BOT] user.
	//
	// For [human users][google.chat.v1.User.Type.HUMAN], `{user}` is the same
	// user identifier as:
	//
	// - the `id` for the
	// [Person](https://developers.google.com/people/api/rest/v1/people) in the
	// People API. For example, `users/123456789` in Chat API represents the same
	// person as the `123456789` Person profile ID in People API.
	//
	// - the `id` for a
	// [user](https://developers.google.com/admin-sdk/directory/reference/rest/v1/users)
	// in the Admin SDK Directory API.
	//
	// - the user's email address can be used as an alias for `{user}` in API
	// requests. For example, if the People API Person profile ID for
	// `user@example.com` is `123456789`, you can use `users/user@example.com` as
	// an alias to reference `users/123456789`. Only the canonical resource name
	// (for example `users/123456789`) will be returned from the API.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The user's display name.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Unique identifier of the user's Google Workspace domain.
	DomainId string `protobuf:"bytes,6,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// User type.
	Type User_Type `protobuf:"varint,5,opt,name=type,proto3,enum=google.chat.v1.User_Type" json:"type,omitempty"`
	// Output only. When `true`, the user is deleted or their profile is not
	// visible.
	IsAnonymous bool `protobuf:"varint,7,opt,name=is_anonymous,json=isAnonymous,proto3" json:"is_anonymous,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_user_proto_msgTypes[0]
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
	return file_google_chat_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *User) GetType() User_Type {
	if x != nil {
		return x.Type
	}
	return User_TYPE_UNSPECIFIED
}

func (x *User) GetIsAnonymous() bool {
	if x != nil {
		return x.IsAnonymous
	}
	return false
}

var File_google_chat_v1_user_proto protoreflect.FileDescriptor

var file_google_chat_v1_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x22, 0x30, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4f, 0x54,
	0x10, 0x02, 0x42, 0x94, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x3b, 0x63, 0x68,
	0x61, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70,
	0x70, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a,
	0x3a, 0x43, 0x68, 0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_chat_v1_user_proto_rawDescOnce sync.Once
	file_google_chat_v1_user_proto_rawDescData = file_google_chat_v1_user_proto_rawDesc
)

func file_google_chat_v1_user_proto_rawDescGZIP() []byte {
	file_google_chat_v1_user_proto_rawDescOnce.Do(func() {
		file_google_chat_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_chat_v1_user_proto_rawDescData)
	})
	return file_google_chat_v1_user_proto_rawDescData
}

var file_google_chat_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_chat_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_chat_v1_user_proto_goTypes = []interface{}{
	(User_Type)(0), // 0: google.chat.v1.User.Type
	(*User)(nil),   // 1: google.chat.v1.User
}
var file_google_chat_v1_user_proto_depIdxs = []int32{
	0, // 0: google.chat.v1.User.type:type_name -> google.chat.v1.User.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_chat_v1_user_proto_init() }
func file_google_chat_v1_user_proto_init() {
	if File_google_chat_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_chat_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_chat_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_chat_v1_user_proto_goTypes,
		DependencyIndexes: file_google_chat_v1_user_proto_depIdxs,
		EnumInfos:         file_google_chat_v1_user_proto_enumTypes,
		MessageInfos:      file_google_chat_v1_user_proto_msgTypes,
	}.Build()
	File_google_chat_v1_user_proto = out.File
	file_google_chat_v1_user_proto_rawDesc = nil
	file_google_chat_v1_user_proto_goTypes = nil
	file_google_chat_v1_user_proto_depIdxs = nil
}