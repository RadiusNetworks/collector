// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: compact_snapshot.proto

package pganalyze_collector

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompactSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Basic information about this snapshot
	SnapshotVersionMajor int32                     `protobuf:"varint,1,opt,name=snapshot_version_major,json=snapshotVersionMajor,proto3" json:"snapshot_version_major,omitempty"`
	SnapshotVersionMinor int32                     `protobuf:"varint,2,opt,name=snapshot_version_minor,json=snapshotVersionMinor,proto3" json:"snapshot_version_minor,omitempty"`
	CollectorVersion     string                    `protobuf:"bytes,3,opt,name=collector_version,json=collectorVersion,proto3" json:"collector_version,omitempty"`
	SnapshotUuid         string                    `protobuf:"bytes,4,opt,name=snapshot_uuid,json=snapshotUuid,proto3" json:"snapshot_uuid,omitempty"`
	CollectedAt          *timestamppb.Timestamp    `protobuf:"bytes,5,opt,name=collected_at,json=collectedAt,proto3" json:"collected_at,omitempty"`
	BaseRefs             *CompactSnapshot_BaseRefs `protobuf:"bytes,6,opt,name=base_refs,json=baseRefs,proto3" json:"base_refs,omitempty"`
	// Types that are assignable to Data:
	//
	//	*CompactSnapshot_LogSnapshot
	//	*CompactSnapshot_SystemSnapshot
	//	*CompactSnapshot_ActivitySnapshot
	Data isCompactSnapshot_Data `protobuf_oneof:"data"`
}

func (x *CompactSnapshot) Reset() {
	*x = CompactSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compact_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactSnapshot) ProtoMessage() {}

func (x *CompactSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_compact_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactSnapshot.ProtoReflect.Descriptor instead.
func (*CompactSnapshot) Descriptor() ([]byte, []int) {
	return file_compact_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *CompactSnapshot) GetSnapshotVersionMajor() int32 {
	if x != nil {
		return x.SnapshotVersionMajor
	}
	return 0
}

func (x *CompactSnapshot) GetSnapshotVersionMinor() int32 {
	if x != nil {
		return x.SnapshotVersionMinor
	}
	return 0
}

func (x *CompactSnapshot) GetCollectorVersion() string {
	if x != nil {
		return x.CollectorVersion
	}
	return ""
}

func (x *CompactSnapshot) GetSnapshotUuid() string {
	if x != nil {
		return x.SnapshotUuid
	}
	return ""
}

func (x *CompactSnapshot) GetCollectedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CollectedAt
	}
	return nil
}

func (x *CompactSnapshot) GetBaseRefs() *CompactSnapshot_BaseRefs {
	if x != nil {
		return x.BaseRefs
	}
	return nil
}

func (m *CompactSnapshot) GetData() isCompactSnapshot_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CompactSnapshot) GetLogSnapshot() *CompactLogSnapshot {
	if x, ok := x.GetData().(*CompactSnapshot_LogSnapshot); ok {
		return x.LogSnapshot
	}
	return nil
}

func (x *CompactSnapshot) GetSystemSnapshot() *CompactSystemSnapshot {
	if x, ok := x.GetData().(*CompactSnapshot_SystemSnapshot); ok {
		return x.SystemSnapshot
	}
	return nil
}

func (x *CompactSnapshot) GetActivitySnapshot() *CompactActivitySnapshot {
	if x, ok := x.GetData().(*CompactSnapshot_ActivitySnapshot); ok {
		return x.ActivitySnapshot
	}
	return nil
}

type isCompactSnapshot_Data interface {
	isCompactSnapshot_Data()
}

type CompactSnapshot_LogSnapshot struct {
	LogSnapshot *CompactLogSnapshot `protobuf:"bytes,10,opt,name=log_snapshot,json=logSnapshot,proto3,oneof"`
}

type CompactSnapshot_SystemSnapshot struct {
	SystemSnapshot *CompactSystemSnapshot `protobuf:"bytes,11,opt,name=system_snapshot,json=systemSnapshot,proto3,oneof"`
}

type CompactSnapshot_ActivitySnapshot struct {
	ActivitySnapshot *CompactActivitySnapshot `protobuf:"bytes,12,opt,name=activity_snapshot,json=activitySnapshot,proto3,oneof"`
}

func (*CompactSnapshot_LogSnapshot) isCompactSnapshot_Data() {}

func (*CompactSnapshot_SystemSnapshot) isCompactSnapshot_Data() {}

func (*CompactSnapshot_ActivitySnapshot) isCompactSnapshot_Data() {}

type CompactSnapshot_BaseRefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleReferences     []*RoleReference     `protobuf:"bytes,1,rep,name=role_references,json=roleReferences,proto3" json:"role_references,omitempty"`
	DatabaseReferences []*DatabaseReference `protobuf:"bytes,2,rep,name=database_references,json=databaseReferences,proto3" json:"database_references,omitempty"`
	QueryReferences    []*QueryReference    `protobuf:"bytes,3,rep,name=query_references,json=queryReferences,proto3" json:"query_references,omitempty"`
	QueryInformations  []*QueryInformation  `protobuf:"bytes,4,rep,name=query_informations,json=queryInformations,proto3" json:"query_informations,omitempty"`
	RelationReferences []*RelationReference `protobuf:"bytes,5,rep,name=relation_references,json=relationReferences,proto3" json:"relation_references,omitempty"`
}

func (x *CompactSnapshot_BaseRefs) Reset() {
	*x = CompactSnapshot_BaseRefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compact_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactSnapshot_BaseRefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactSnapshot_BaseRefs) ProtoMessage() {}

func (x *CompactSnapshot_BaseRefs) ProtoReflect() protoreflect.Message {
	mi := &file_compact_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactSnapshot_BaseRefs.ProtoReflect.Descriptor instead.
func (*CompactSnapshot_BaseRefs) Descriptor() ([]byte, []int) {
	return file_compact_snapshot_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CompactSnapshot_BaseRefs) GetRoleReferences() []*RoleReference {
	if x != nil {
		return x.RoleReferences
	}
	return nil
}

func (x *CompactSnapshot_BaseRefs) GetDatabaseReferences() []*DatabaseReference {
	if x != nil {
		return x.DatabaseReferences
	}
	return nil
}

func (x *CompactSnapshot_BaseRefs) GetQueryReferences() []*QueryReference {
	if x != nil {
		return x.QueryReferences
	}
	return nil
}

func (x *CompactSnapshot_BaseRefs) GetQueryInformations() []*QueryInformation {
	if x != nil {
		return x.QueryInformations
	}
	return nil
}

func (x *CompactSnapshot_BaseRefs) GetRelationReferences() []*RelationReference {
	if x != nil {
		return x.RelationReferences
	}
	return nil
}

var File_compact_snapshot_proto protoreflect.FileDescriptor

var file_compact_snapshot_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x08, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a,
	0x6f, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x14, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4a, 0x0a, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70,
	0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x73, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x66, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x67,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x55, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70,
	0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x5b, 0x0a, 0x11, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x48, 0x00, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x1a, 0xaf, 0x03, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x66, 0x73, 0x12, 0x4b, 0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x57, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x12, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x57, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x67, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x70, 0x67, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compact_snapshot_proto_rawDescOnce sync.Once
	file_compact_snapshot_proto_rawDescData = file_compact_snapshot_proto_rawDesc
)

func file_compact_snapshot_proto_rawDescGZIP() []byte {
	file_compact_snapshot_proto_rawDescOnce.Do(func() {
		file_compact_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_compact_snapshot_proto_rawDescData)
	})
	return file_compact_snapshot_proto_rawDescData
}

var file_compact_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_compact_snapshot_proto_goTypes = []interface{}{
	(*CompactSnapshot)(nil),          // 0: pganalyze.collector.CompactSnapshot
	(*CompactSnapshot_BaseRefs)(nil), // 1: pganalyze.collector.CompactSnapshot.BaseRefs
	(*timestamppb.Timestamp)(nil),    // 2: google.protobuf.Timestamp
	(*CompactLogSnapshot)(nil),       // 3: pganalyze.collector.CompactLogSnapshot
	(*CompactSystemSnapshot)(nil),    // 4: pganalyze.collector.CompactSystemSnapshot
	(*CompactActivitySnapshot)(nil),  // 5: pganalyze.collector.CompactActivitySnapshot
	(*RoleReference)(nil),            // 6: pganalyze.collector.RoleReference
	(*DatabaseReference)(nil),        // 7: pganalyze.collector.DatabaseReference
	(*QueryReference)(nil),           // 8: pganalyze.collector.QueryReference
	(*QueryInformation)(nil),         // 9: pganalyze.collector.QueryInformation
	(*RelationReference)(nil),        // 10: pganalyze.collector.RelationReference
}
var file_compact_snapshot_proto_depIdxs = []int32{
	2,  // 0: pganalyze.collector.CompactSnapshot.collected_at:type_name -> google.protobuf.Timestamp
	1,  // 1: pganalyze.collector.CompactSnapshot.base_refs:type_name -> pganalyze.collector.CompactSnapshot.BaseRefs
	3,  // 2: pganalyze.collector.CompactSnapshot.log_snapshot:type_name -> pganalyze.collector.CompactLogSnapshot
	4,  // 3: pganalyze.collector.CompactSnapshot.system_snapshot:type_name -> pganalyze.collector.CompactSystemSnapshot
	5,  // 4: pganalyze.collector.CompactSnapshot.activity_snapshot:type_name -> pganalyze.collector.CompactActivitySnapshot
	6,  // 5: pganalyze.collector.CompactSnapshot.BaseRefs.role_references:type_name -> pganalyze.collector.RoleReference
	7,  // 6: pganalyze.collector.CompactSnapshot.BaseRefs.database_references:type_name -> pganalyze.collector.DatabaseReference
	8,  // 7: pganalyze.collector.CompactSnapshot.BaseRefs.query_references:type_name -> pganalyze.collector.QueryReference
	9,  // 8: pganalyze.collector.CompactSnapshot.BaseRefs.query_informations:type_name -> pganalyze.collector.QueryInformation
	10, // 9: pganalyze.collector.CompactSnapshot.BaseRefs.relation_references:type_name -> pganalyze.collector.RelationReference
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_compact_snapshot_proto_init() }
func file_compact_snapshot_proto_init() {
	if File_compact_snapshot_proto != nil {
		return
	}
	file_compact_activity_snapshot_proto_init()
	file_compact_log_snapshot_proto_init()
	file_compact_system_snapshot_proto_init()
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_compact_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactSnapshot); i {
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
		file_compact_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactSnapshot_BaseRefs); i {
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
	file_compact_snapshot_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CompactSnapshot_LogSnapshot)(nil),
		(*CompactSnapshot_SystemSnapshot)(nil),
		(*CompactSnapshot_ActivitySnapshot)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_compact_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compact_snapshot_proto_goTypes,
		DependencyIndexes: file_compact_snapshot_proto_depIdxs,
		MessageInfos:      file_compact_snapshot_proto_msgTypes,
	}.Build()
	File_compact_snapshot_proto = out.File
	file_compact_snapshot_proto_rawDesc = nil
	file_compact_snapshot_proto_goTypes = nil
	file_compact_snapshot_proto_depIdxs = nil
}
