// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/redis/pb/redis.proto

package redis

import (
	resource "github.com/infraboard/cmdb/apps/resource"
	request "github.com/infraboard/mcube/http/request"
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

type STATUS int32

const (
	// 未知状态
	STATUS_UNKNOW STATUS = 0
	// 表示创建中
	STATUS_PENDING STATUS = 5
	// 表示创建失败
	STATUS_CREATE_FAILED STATUS = 6
	// 表示运行中
	STATUS_RUNNING STATUS = 11
	// 实例配置变更生效中
	STATUS_MODIFYING STATUS = 20
	// 表示重启中
	STATUS_REBOOTING STATUS = 30
	// 备份恢复中
	STATUS_RESTORING STATUS = 31
	// 迁移中
	STATUS_TRANSING STATUS = 32
	// 缓存实例数据清空中的状态
	STATUS_FLUSHING STATUS = 33
	// 缓存实例处于正在扩容的状态
	STATUS_EXTENDING STATUS = 34
	// 迁移版本中
	STATUS_UPGRADING STATUS = 35
	// 表示实例正在进行备份
	STATUS_BACKING_UP STATUS = 36
	// 内外网切换中
	STATUS_NET_CHANGING STATUS = 37
	// 状态异常
	STATUS_ERROR STATUS = 50
	// 表示实例已经锁定
	STATUS_LOCKED STATUS = 70
	// 隔离中
	STATUS_ISOLATIONING STATUS = 71
	// 已隔中
	STATUS_ISOLATIONED STATUS = 72
	// 表示停止待销毁
	STATUS_SHUTDOWN STATUS = 80
	// 表示销毁中
	STATUS_DELETING STATUS = 81
	// 已销毁
	STATUS_DESTROYED STATUS = 90
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0:  "UNKNOW",
		5:  "PENDING",
		6:  "CREATE_FAILED",
		11: "RUNNING",
		20: "MODIFYING",
		30: "REBOOTING",
		31: "RESTORING",
		32: "TRANSING",
		33: "FLUSHING",
		34: "EXTENDING",
		35: "UPGRADING",
		36: "BACKING_UP",
		37: "NET_CHANGING",
		50: "ERROR",
		70: "LOCKED",
		71: "ISOLATIONING",
		72: "ISOLATIONED",
		80: "SHUTDOWN",
		81: "DELETING",
		90: "DESTROYED",
	}
	STATUS_value = map[string]int32{
		"UNKNOW":        0,
		"PENDING":       5,
		"CREATE_FAILED": 6,
		"RUNNING":       11,
		"MODIFYING":     20,
		"REBOOTING":     30,
		"RESTORING":     31,
		"TRANSING":      32,
		"FLUSHING":      33,
		"EXTENDING":     34,
		"UPGRADING":     35,
		"BACKING_UP":    36,
		"NET_CHANGING":  37,
		"ERROR":         50,
		"LOCKED":        70,
		"ISOLATIONING":  71,
		"ISOLATIONED":   72,
		"SHUTDOWN":      80,
		"DELETING":      81,
		"DESTROYED":     90,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_redis_pb_redis_proto_enumTypes[0].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_apps_redis_pb_redis_proto_enumTypes[0]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_apps_redis_pb_redis_proto_rawDescGZIP(), []int{0}
}

type Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"resource"
	Resource *resource.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource"`
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe"`
}

func (x *Redis) Reset() {
	*x = Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_redis_pb_redis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redis) ProtoMessage() {}

func (x *Redis) ProtoReflect() protoreflect.Message {
	mi := &file_apps_redis_pb_redis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redis.ProtoReflect.Descriptor instead.
func (*Redis) Descriptor() ([]byte, []int) {
	return file_apps_redis_pb_redis_proto_rawDescGZIP(), []int{0}
}

func (x *Redis) GetResource() *resource.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Redis) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 链接地址
	// @gotags: json:"connect_addr"
	ConnectAddr string `protobuf:"bytes,2,opt,name=connect_addr,json=connectAddr,proto3" json:"connect_addr"`
	// 链接端口
	// @gotags: json:"connect_port"
	ConnectPort int64 `protobuf:"varint,3,opt,name=connect_port,json=connectPort,proto3" json:"connect_port"`
	// 指定架构类型返回实例列表：cluster（集群版）standard（标准版） SplitRW（读写分离版）NULL（所有类型，默认值）
	// @gotags: json:"architecture_type"
	ArchitectureType string `protobuf:"bytes,4,opt,name=architecture_type,json=architectureType,proto3" json:"architecture_type"`
	// 理论最大QPS值
	// @gotags: json:"qps"
	Qps int64 `protobuf:"varint,5,opt,name=qps,proto3" json:"qps"`
	// 实例连接数限制，单位：个
	// @gotags: json:"max_connection"
	MaxConnection int64 `protobuf:"varint,7,opt,name=max_connection,json=maxConnection,proto3" json:"max_connection"`
	// 配置
	// @gotags: json:"config"
	Config string `protobuf:"bytes,8,opt,name=config,proto3" json:"config"`
	// 节点类型：double（双节点） single（单节点）
	// @gotags: json:"node_type"
	NodeType string `protobuf:"bytes,9,opt,name=node_type,json=nodeType,proto3" json:"node_type"`
	// 网络类型：CLASSIC（经典网络）VPC（VPC网络）
	// @gotags: json:"network_type"
	NetworkType string `protobuf:"bytes,10,opt,name=network_type,json=networkType,proto3" json:"network_type"`
	// 数据库类型。
	// @gotags: json:"engine_type"
	EngineType string `protobuf:"bytes,11,opt,name=engine_type,json=engineType,proto3" json:"engine_type"`
	// 数据库版本：2.8 4.0 5.0
	// @gotags: json:"engine_version"
	EngineVersion string `protobuf:"bytes,12,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version"`
	// 副本架构：master-slave（包括主从版和单节点版）cluster（包括读写分离版与集群版）
	// @gotags: json:"replication_mode"
	ReplicationMode string `protobuf:"bytes,13,opt,name=replication_mode,json=replicationMode,proto3" json:"replication_mode"`
	// 副本ID。
	// @gotags: json:"replica_id"
	ReplicaId string `protobuf:"bytes,14,opt,name=replica_id,json=replicaId,proto3" json:"replica_id"`
	// IP白名单
	// @gotags: json:"security_ip_list"
	SecurityIpList string `protobuf:"bytes,15,opt,name=security_ip_list,json=securityIpList,proto3" json:"security_ip_list"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_redis_pb_redis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_redis_pb_redis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Describe.ProtoReflect.Descriptor instead.
func (*Describe) Descriptor() ([]byte, []int) {
	return file_apps_redis_pb_redis_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetConnectAddr() string {
	if x != nil {
		return x.ConnectAddr
	}
	return ""
}

func (x *Describe) GetConnectPort() int64 {
	if x != nil {
		return x.ConnectPort
	}
	return 0
}

func (x *Describe) GetArchitectureType() string {
	if x != nil {
		return x.ArchitectureType
	}
	return ""
}

func (x *Describe) GetQps() int64 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *Describe) GetMaxConnection() int64 {
	if x != nil {
		return x.MaxConnection
	}
	return 0
}

func (x *Describe) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Describe) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *Describe) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

func (x *Describe) GetEngineType() string {
	if x != nil {
		return x.EngineType
	}
	return ""
}

func (x *Describe) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *Describe) GetReplicationMode() string {
	if x != nil {
		return x.ReplicationMode
	}
	return ""
}

func (x *Describe) GetReplicaId() string {
	if x != nil {
		return x.ReplicaId
	}
	return ""
}

func (x *Describe) GetSecurityIpList() string {
	if x != nil {
		return x.SecurityIpList
	}
	return ""
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表项
	// @gotags: json:"items"
	Items []*Redis `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_redis_pb_redis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_redis_pb_redis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_apps_redis_pb_redis_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Redis {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueryRedisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryRedisRequest) Reset() {
	*x = QueryRedisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_redis_pb_redis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRedisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRedisRequest) ProtoMessage() {}

func (x *QueryRedisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_redis_pb_redis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRedisRequest.ProtoReflect.Descriptor instead.
func (*QueryRedisRequest) Descriptor() ([]byte, []int) {
	return file_apps_redis_pb_redis_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRedisRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_apps_redis_pb_redis_proto protoreflect.FileDescriptor

var file_apps_redis_pb_redis_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x3e, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0xca, 0x03, 0x0a, 0x08, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2b, 0x0a,
	0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x70,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x70, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x49, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4b, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x2a, 0xaf, 0x02, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x44,
	0x49, 0x46, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x42, 0x4f,
	0x4f, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x1e, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x53, 0x54, 0x4f,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x1f, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x20, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4c, 0x55, 0x53, 0x48, 0x49, 0x4e, 0x47,
	0x10, 0x21, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x22, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x23,
	0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50, 0x10, 0x24,
	0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x49, 0x4e, 0x47,
	0x10, 0x25, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x32, 0x12, 0x0a, 0x0a,
	0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x46, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x53, 0x4f,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x47, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x53, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x48, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x50, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x51, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x53, 0x54,
	0x52, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x5a, 0x32, 0xaa, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x1c,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x28, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x53,
	0x65, 0x74, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d,
	0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_redis_pb_redis_proto_rawDescOnce sync.Once
	file_apps_redis_pb_redis_proto_rawDescData = file_apps_redis_pb_redis_proto_rawDesc
)

func file_apps_redis_pb_redis_proto_rawDescGZIP() []byte {
	file_apps_redis_pb_redis_proto_rawDescOnce.Do(func() {
		file_apps_redis_pb_redis_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_redis_pb_redis_proto_rawDescData)
	})
	return file_apps_redis_pb_redis_proto_rawDescData
}

var file_apps_redis_pb_redis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_redis_pb_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_redis_pb_redis_proto_goTypes = []interface{}{
	(STATUS)(0),                 // 0: infraboard.cmdb.redis.STATUS
	(*Redis)(nil),               // 1: infraboard.cmdb.redis.Redis
	(*Describe)(nil),            // 2: infraboard.cmdb.redis.Describe
	(*Set)(nil),                 // 3: infraboard.cmdb.redis.Set
	(*QueryRedisRequest)(nil),   // 4: infraboard.cmdb.redis.QueryRedisRequest
	(*resource.Resource)(nil),   // 5: infraboard.cmdb.resource.Resource
	(*request.PageRequest)(nil), // 6: infraboard.mcube.page.PageRequest
}
var file_apps_redis_pb_redis_proto_depIdxs = []int32{
	5, // 0: infraboard.cmdb.redis.Redis.resource:type_name -> infraboard.cmdb.resource.Resource
	2, // 1: infraboard.cmdb.redis.Redis.describe:type_name -> infraboard.cmdb.redis.Describe
	1, // 2: infraboard.cmdb.redis.Set.items:type_name -> infraboard.cmdb.redis.Redis
	6, // 3: infraboard.cmdb.redis.QueryRedisRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // 4: infraboard.cmdb.redis.Service.SyncRedis:input_type -> infraboard.cmdb.redis.Redis
	4, // 5: infraboard.cmdb.redis.Service.QueryRedis:input_type -> infraboard.cmdb.redis.QueryRedisRequest
	1, // 6: infraboard.cmdb.redis.Service.SyncRedis:output_type -> infraboard.cmdb.redis.Redis
	3, // 7: infraboard.cmdb.redis.Service.QueryRedis:output_type -> infraboard.cmdb.redis.Set
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_redis_pb_redis_proto_init() }
func file_apps_redis_pb_redis_proto_init() {
	if File_apps_redis_pb_redis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_redis_pb_redis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redis); i {
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
		file_apps_redis_pb_redis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Describe); i {
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
		file_apps_redis_pb_redis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_apps_redis_pb_redis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRedisRequest); i {
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
			RawDescriptor: file_apps_redis_pb_redis_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_redis_pb_redis_proto_goTypes,
		DependencyIndexes: file_apps_redis_pb_redis_proto_depIdxs,
		EnumInfos:         file_apps_redis_pb_redis_proto_enumTypes,
		MessageInfos:      file_apps_redis_pb_redis_proto_msgTypes,
	}.Build()
	File_apps_redis_pb_redis_proto = out.File
	file_apps_redis_pb_redis_proto_rawDesc = nil
	file_apps_redis_pb_redis_proto_goTypes = nil
	file_apps_redis_pb_redis_proto_depIdxs = nil
}
