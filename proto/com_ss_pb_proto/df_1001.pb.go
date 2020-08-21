// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: df_1001.proto

package com_ss_pb_proto

import (
	proto "github.com/golang/protobuf/proto"
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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *uint32  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"` // 结果值 1成功 | 其他数值参考错误码
	Param  []string `protobuf:"bytes,2,rep,name=param" json:"param,omitempty"`    // 消息参数
	Type   *uint32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`     // 1-飘字，2-弹框
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *Result) GetParam() []string {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *Result) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

type Server_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid   *uint32 `protobuf:"varint,1,opt,name=sid" json:"sid,omitempty"`     // 服务器ID
	Sname *string `protobuf:"bytes,2,opt,name=sname" json:"sname,omitempty"`  // 服务器名称
	Host  *string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`    // 服务器地址
	Port  *uint32 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`   // 服务器端口号
	Uid   *uint32 `protobuf:"varint,5,opt,name=uid" json:"uid,omitempty"`     // 该服角色ID, 0表示无角色
	Key   *string `protobuf:"bytes,6,opt,name=key" json:"key,omitempty"`      // 登录安全key
	State *uint32 `protobuf:"varint,7,opt,name=state" json:"state,omitempty"` // 服务器状态   (正常,爆满,维护)
}

func (x *Server_Info) Reset() {
	*x = Server_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Info) ProtoMessage() {}

func (x *Server_Info) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Info.ProtoReflect.Descriptor instead.
func (*Server_Info) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{1}
}

func (x *Server_Info) GetSid() uint32 {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return 0
}

func (x *Server_Info) GetSname() string {
	if x != nil && x.Sname != nil {
		return *x.Sname
	}
	return ""
}

func (x *Server_Info) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *Server_Info) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *Server_Info) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *Server_Info) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *Server_Info) GetState() uint32 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

// 错误码通用
type Sc_10010000 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"` // 返回结果
}

func (x *Sc_10010000) Reset() {
	*x = Sc_10010000{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_10010000) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_10010000) ProtoMessage() {}

func (x *Sc_10010000) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_10010000.ProtoReflect.Descriptor instead.
func (*Sc_10010000) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{2}
}

func (x *Sc_10010000) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

//登陆游戏
type Cs_10010001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    *uint32 `protobuf:"varint,1,opt,name=uuid" json:"uuid,omitempty"`                      // 账号ID
	Uid     *uint32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`                        // 角色ID
	Cid     *uint32 `protobuf:"varint,3,opt,name=cid" json:"cid,omitempty"`                        // 渠道ID
	Sid     *uint32 `protobuf:"varint,4,opt,name=sid" json:"sid,omitempty"`                        // 服务器ID
	KeyTime *uint32 `protobuf:"varint,5,opt,name=key_time,json=keyTime" json:"key_time,omitempty"` // 登录安全验证时间
	Key     *string `protobuf:"bytes,6,opt,name=key" json:"key,omitempty"`                         // 登录安全key
}

func (x *Cs_10010001) Reset() {
	*x = Cs_10010001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_10010001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_10010001) ProtoMessage() {}

func (x *Cs_10010001) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_10010001.ProtoReflect.Descriptor instead.
func (*Cs_10010001) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{3}
}

func (x *Cs_10010001) GetUuid() uint32 {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return 0
}

func (x *Cs_10010001) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *Cs_10010001) GetCid() uint32 {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return 0
}

func (x *Cs_10010001) GetSid() uint32 {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return 0
}

func (x *Cs_10010001) GetKeyTime() uint32 {
	if x != nil && x.KeyTime != nil {
		return *x.KeyTime
	}
	return 0
}

func (x *Cs_10010001) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

type Sc_10010001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginResult *bool `protobuf:"varint,1,opt,name=login_result,json=loginResult" json:"login_result,omitempty"` // false-需要创建角色，true-登录成功
}

func (x *Sc_10010001) Reset() {
	*x = Sc_10010001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_10010001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_10010001) ProtoMessage() {}

func (x *Sc_10010001) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_10010001.ProtoReflect.Descriptor instead.
func (*Sc_10010001) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{4}
}

func (x *Sc_10010001) GetLoginResult() bool {
	if x != nil && x.LoginResult != nil {
		return *x.LoginResult
	}
	return false
}

// 创建角色(必须先登录,再创角色)
type Cs_10010002 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uname *string `protobuf:"bytes,1,opt,name=uname" json:"uname,omitempty"` // 昵称
	Sex   *uint32 `protobuf:"varint,2,opt,name=sex" json:"sex,omitempty"`    // 性别1-男性，2-女性
}

func (x *Cs_10010002) Reset() {
	*x = Cs_10010002{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_10010002) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_10010002) ProtoMessage() {}

func (x *Cs_10010002) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_10010002.ProtoReflect.Descriptor instead.
func (*Cs_10010002) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{5}
}

func (x *Cs_10010002) GetUname() string {
	if x != nil && x.Uname != nil {
		return *x.Uname
	}
	return ""
}

func (x *Cs_10010002) GetSex() uint32 {
	if x != nil && x.Sex != nil {
		return *x.Sex
	}
	return 0
}

// 创建成功返回
type Sc_10010002 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid *uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"` // 角色ID 登录成功
}

func (x *Sc_10010002) Reset() {
	*x = Sc_10010002{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_10010002) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_10010002) ProtoMessage() {}

func (x *Sc_10010002) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_10010002.ProtoReflect.Descriptor instead.
func (*Sc_10010002) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{6}
}

func (x *Sc_10010002) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

// 服务器时间校正
type Sc_10010003 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *uint32 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"` // 服务器时间戳
}

func (x *Sc_10010003) Reset() {
	*x = Sc_10010003{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_10010003) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_10010003) ProtoMessage() {}

func (x *Sc_10010003) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_10010003.ProtoReflect.Descriptor instead.
func (*Sc_10010003) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{7}
}

func (x *Sc_10010003) GetTime() uint32 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

// 请求服务器列表
type Cs_10010004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SdkUuid    *string `protobuf:"bytes,1,opt,name=sdk_uuid,json=sdkUuid" json:"sdk_uuid,omitempty"`          // SDK登录成功后的用户ID
	Cid        *uint32 `protobuf:"varint,2,opt,name=cid" json:"cid,omitempty"`                                // 渠道ID
	LoginToken *string `protobuf:"bytes,3,opt,name=login_token,json=loginToken" json:"login_token,omitempty"` // SDK登录标记
}

func (x *Cs_10010004) Reset() {
	*x = Cs_10010004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_10010004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_10010004) ProtoMessage() {}

func (x *Cs_10010004) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_10010004.ProtoReflect.Descriptor instead.
func (*Cs_10010004) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{8}
}

func (x *Cs_10010004) GetSdkUuid() string {
	if x != nil && x.SdkUuid != nil {
		return *x.SdkUuid
	}
	return ""
}

func (x *Cs_10010004) GetCid() uint32 {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return 0
}

func (x *Cs_10010004) GetLoginToken() string {
	if x != nil && x.LoginToken != nil {
		return *x.LoginToken
	}
	return ""
}

// 请求服务器列表
type Sc_10010004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       *uint32        `protobuf:"varint,1,opt,name=uuid" json:"uuid,omitempty"`                              // 用户账号ID
	KeyTime    *uint32        `protobuf:"varint,2,opt,name=key_time,json=keyTime" json:"key_time,omitempty"`         // 登录安全验证时间
	ServerInfo []*Server_Info `protobuf:"bytes,3,rep,name=server_info,json=serverInfo" json:"server_info,omitempty"` // 服务器信息列表
}

func (x *Sc_10010004) Reset() {
	*x = Sc_10010004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sc_10010004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sc_10010004) ProtoMessage() {}

func (x *Sc_10010004) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sc_10010004.ProtoReflect.Descriptor instead.
func (*Sc_10010004) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{9}
}

func (x *Sc_10010004) GetUuid() uint32 {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return 0
}

func (x *Sc_10010004) GetKeyTime() uint32 {
	if x != nil && x.KeyTime != nil {
		return *x.KeyTime
	}
	return 0
}

func (x *Sc_10010004) GetServerInfo() []*Server_Info {
	if x != nil {
		return x.ServerInfo
	}
	return nil
}

// 登录成功后上报数据
type Cs_10010005 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OsType *uint32 `protobuf:"varint,1,opt,name=os_type,json=osType" json:"os_type,omitempty"` // 系统类型 CONST_SYS_OS_TYPE_*
}

func (x *Cs_10010005) Reset() {
	*x = Cs_10010005{}
	if protoimpl.UnsafeEnabled {
		mi := &file_df_1001_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cs_10010005) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cs_10010005) ProtoMessage() {}

func (x *Cs_10010005) ProtoReflect() protoreflect.Message {
	mi := &file_df_1001_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cs_10010005.ProtoReflect.Descriptor instead.
func (*Cs_10010005) Descriptor() ([]byte, []int) {
	return file_df_1001_proto_rawDescGZIP(), []int{10}
}

func (x *Cs_10010005) GetOsType() uint32 {
	if x != nil && x.OsType != nil {
		return *x.OsType
	}
	return 0
}

var File_df_1001_proto protoreflect.FileDescriptor

var file_df_1001_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x66, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0d, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x12, 0x0c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x22, 0x6e, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x12, 0x0d, 0x0a, 0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x12, 0x0c, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x12, 0x0c, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0b,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0b, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x22, 0x36, 0x0a, 0x0b, 0x53, 0x63, 0x5f, 0x31, 0x30,
	0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x2e,
	0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x61, 0x0a, 0x0b, 0x43, 0x73, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x31, 0x12, 0x0c,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0b, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0b, 0x0a, 0x03, 0x63, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0b, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x12, 0x10, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x0b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x22, 0x23, 0x0a, 0x0b, 0x53, 0x63, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30,
	0x31, 0x12, 0x14, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x22, 0x29, 0x0a, 0x0b, 0x43, 0x73, 0x5f, 0x31, 0x30,
	0x30, 0x31, 0x30, 0x30, 0x30, 0x32, 0x12, 0x0d, 0x0a, 0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0b, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x22, 0x1a, 0x0a, 0x0b, 0x53, 0x63, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30,
	0x32, 0x12, 0x0b, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x22, 0x1b,
	0x0a, 0x0b, 0x53, 0x63, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x33, 0x12, 0x0c, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x22, 0x41, 0x0a, 0x0b, 0x43,
	0x73, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x34, 0x12, 0x10, 0x0a, 0x08, 0x73, 0x64,
	0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x12, 0x0b, 0x0a, 0x03,
	0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x13, 0x0a, 0x0b, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x22, 0x60,
	0x0a, 0x0b, 0x53, 0x63, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x34, 0x12, 0x0c, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x10, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x12, 0x31, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x1e, 0x0a, 0x0b, 0x43, 0x73, 0x5f, 0x31, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x35, 0x12,
	0x0f, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
}

var (
	file_df_1001_proto_rawDescOnce sync.Once
	file_df_1001_proto_rawDescData = file_df_1001_proto_rawDesc
)

func file_df_1001_proto_rawDescGZIP() []byte {
	file_df_1001_proto_rawDescOnce.Do(func() {
		file_df_1001_proto_rawDescData = protoimpl.X.CompressGZIP(file_df_1001_proto_rawDescData)
	})
	return file_df_1001_proto_rawDescData
}

var file_df_1001_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_df_1001_proto_goTypes = []interface{}{
	(*Result)(nil),      // 0: com.ss.pb.proto.Result
	(*Server_Info)(nil), // 1: com.ss.pb.proto.Server_Info
	(*Sc_10010000)(nil), // 2: com.ss.pb.proto.Sc_10010000
	(*Cs_10010001)(nil), // 3: com.ss.pb.proto.Cs_10010001
	(*Sc_10010001)(nil), // 4: com.ss.pb.proto.Sc_10010001
	(*Cs_10010002)(nil), // 5: com.ss.pb.proto.Cs_10010002
	(*Sc_10010002)(nil), // 6: com.ss.pb.proto.Sc_10010002
	(*Sc_10010003)(nil), // 7: com.ss.pb.proto.Sc_10010003
	(*Cs_10010004)(nil), // 8: com.ss.pb.proto.Cs_10010004
	(*Sc_10010004)(nil), // 9: com.ss.pb.proto.Sc_10010004
	(*Cs_10010005)(nil), // 10: com.ss.pb.proto.Cs_10010005
}
var file_df_1001_proto_depIdxs = []int32{
	0, // 0: com.ss.pb.proto.Sc_10010000.result:type_name -> com.ss.pb.proto.Result
	1, // 1: com.ss.pb.proto.Sc_10010004.server_info:type_name -> com.ss.pb.proto.Server_Info
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_df_1001_proto_init() }
func file_df_1001_proto_init() {
	if File_df_1001_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_df_1001_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_df_1001_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Info); i {
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
		file_df_1001_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_10010000); i {
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
		file_df_1001_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_10010001); i {
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
		file_df_1001_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_10010001); i {
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
		file_df_1001_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_10010002); i {
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
		file_df_1001_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_10010002); i {
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
		file_df_1001_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_10010003); i {
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
		file_df_1001_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_10010004); i {
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
		file_df_1001_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sc_10010004); i {
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
		file_df_1001_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cs_10010005); i {
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
			RawDescriptor: file_df_1001_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_df_1001_proto_goTypes,
		DependencyIndexes: file_df_1001_proto_depIdxs,
		MessageInfos:      file_df_1001_proto_msgTypes,
	}.Build()
	File_df_1001_proto = out.File
	file_df_1001_proto_rawDesc = nil
	file_df_1001_proto_goTypes = nil
	file_df_1001_proto_depIdxs = nil
}