// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/mail.proto

package api

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{0}
}

type MailboxId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MailboxId) Reset() {
	*x = MailboxId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxId) ProtoMessage() {}

func (x *MailboxId) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxId.ProtoReflect.Descriptor instead.
func (*MailboxId) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{1}
}

func (x *MailboxId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MailboxIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Login    string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *MailboxIn) Reset() {
	*x = MailboxIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxIn) ProtoMessage() {}

func (x *MailboxIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxIn.ProtoReflect.Descriptor instead.
func (*MailboxIn) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{2}
}

func (x *MailboxIn) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *MailboxIn) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *MailboxIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MailboxOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Server string `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Login  string `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *MailboxOut) Reset() {
	*x = MailboxOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxOut) ProtoMessage() {}

func (x *MailboxOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxOut.ProtoReflect.Descriptor instead.
func (*MailboxOut) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{3}
}

func (x *MailboxOut) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MailboxOut) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *MailboxOut) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type Mailboxes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mailboxes []*MailboxOut `protobuf:"bytes,1,rep,name=mailboxes,proto3" json:"mailboxes,omitempty"`
}

func (x *Mailboxes) Reset() {
	*x = Mailboxes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mailboxes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mailboxes) ProtoMessage() {}

func (x *Mailboxes) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mailboxes.ProtoReflect.Descriptor instead.
func (*Mailboxes) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{4}
}

func (x *Mailboxes) GetMailboxes() []*MailboxOut {
	if x != nil {
		return x.Mailboxes
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	From      string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{5}
}

func (x *Message) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Messages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Messages) Reset() {
	*x = Messages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Messages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Messages) ProtoMessage() {}

func (x *Messages) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Messages.ProtoReflect.Descriptor instead.
func (*Messages) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{6}
}

func (x *Messages) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type MailboxCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mailbox *MailboxIn `protobuf:"bytes,2,opt,name=mailbox,proto3" json:"mailbox,omitempty"`
}

func (x *MailboxCreate) Reset() {
	*x = MailboxCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxCreate) ProtoMessage() {}

func (x *MailboxCreate) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxCreate.ProtoReflect.Descriptor instead.
func (*MailboxCreate) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{7}
}

func (x *MailboxCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MailboxCreate) GetMailbox() *MailboxIn {
	if x != nil {
		return x.Mailbox
	}
	return nil
}

type MailboxGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *MailboxGet) Reset() {
	*x = MailboxGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxGet) ProtoMessage() {}

func (x *MailboxGet) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxGet.ProtoReflect.Descriptor instead.
func (*MailboxGet) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{8}
}

func (x *MailboxGet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MailboxDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mailbox *MailboxId `protobuf:"bytes,2,opt,name=mailbox,proto3" json:"mailbox,omitempty"`
}

func (x *MailboxDelete) Reset() {
	*x = MailboxDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mail_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxDelete) ProtoMessage() {}

func (x *MailboxDelete) ProtoReflect() protoreflect.Message {
	mi := &file_api_mail_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxDelete.ProtoReflect.Descriptor instead.
func (*MailboxDelete) Descriptor() ([]byte, []int) {
	return file_api_mail_proto_rawDescGZIP(), []int{9}
}

func (x *MailboxDelete) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MailboxDelete) GetMailbox() *MailboxId {
	if x != nil {
		return x.Mailbox
	}
	return nil
}

var File_api_mail_proto protoreflect.FileDescriptor

var file_api_mail_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b,
	0x0a, 0x09, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x09, 0x4d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x4a, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4f, 0x75, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x3a,
	0x0a, 0x09, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x6d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4f, 0x75, 0x74, 0x52,
	0x09, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x34, 0x0a,
	0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x0d, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x49, 0x6e, 0x52, 0x07,
	0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x22, 0x25, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x47, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52,
	0x0a, 0x0d, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x61, 0x69, 0x6c,
	0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x52, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x32, 0xbb, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x29, 0x0a, 0x06, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x47, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c,
	0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x47, 0x65,
	0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x63, 0x61, 0x70, 0x63, 0x6f, 0x6d, 0x36, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_mail_proto_rawDescOnce sync.Once
	file_api_mail_proto_rawDescData = file_api_mail_proto_rawDesc
)

func file_api_mail_proto_rawDescGZIP() []byte {
	file_api_mail_proto_rawDescOnce.Do(func() {
		file_api_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mail_proto_rawDescData)
	})
	return file_api_mail_proto_rawDescData
}

var file_api_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_mail_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: api.Empty
	(*MailboxId)(nil),     // 1: api.MailboxId
	(*MailboxIn)(nil),     // 2: api.MailboxIn
	(*MailboxOut)(nil),    // 3: api.MailboxOut
	(*Mailboxes)(nil),     // 4: api.Mailboxes
	(*Message)(nil),       // 5: api.Message
	(*Messages)(nil),      // 6: api.Messages
	(*MailboxCreate)(nil), // 7: api.MailboxCreate
	(*MailboxGet)(nil),    // 8: api.MailboxGet
	(*MailboxDelete)(nil), // 9: api.MailboxDelete
}
var file_api_mail_proto_depIdxs = []int32{
	3, // 0: api.Mailboxes.mailboxes:type_name -> api.MailboxOut
	5, // 1: api.Messages.messages:type_name -> api.Message
	2, // 2: api.MailboxCreate.mailbox:type_name -> api.MailboxIn
	1, // 3: api.MailboxDelete.mailbox:type_name -> api.MailboxId
	7, // 4: api.MailAggregator.Create:input_type -> api.MailboxCreate
	8, // 5: api.MailAggregator.Select:input_type -> api.MailboxGet
	9, // 6: api.MailAggregator.Delete:input_type -> api.MailboxDelete
	8, // 7: api.MailAggregator.Pull:input_type -> api.MailboxGet
	0, // 8: api.MailAggregator.Create:output_type -> api.Empty
	4, // 9: api.MailAggregator.Select:output_type -> api.Mailboxes
	4, // 10: api.MailAggregator.Delete:output_type -> api.Mailboxes
	6, // 11: api.MailAggregator.Pull:output_type -> api.Messages
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_mail_proto_init() }
func file_api_mail_proto_init() {
	if File_api_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxId); i {
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
		file_api_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxIn); i {
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
		file_api_mail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxOut); i {
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
		file_api_mail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mailboxes); i {
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
		file_api_mail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_api_mail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Messages); i {
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
		file_api_mail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxCreate); i {
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
		file_api_mail_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxGet); i {
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
		file_api_mail_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxDelete); i {
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
			RawDescriptor: file_api_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mail_proto_goTypes,
		DependencyIndexes: file_api_mail_proto_depIdxs,
		MessageInfos:      file_api_mail_proto_msgTypes,
	}.Build()
	File_api_mail_proto = out.File
	file_api_mail_proto_rawDesc = nil
	file_api_mail_proto_goTypes = nil
	file_api_mail_proto_depIdxs = nil
}