// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: p2p/proto/class.proto

package spec

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

type EntryPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector *Felt252 `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	Offset   uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *EntryPoint) Reset() {
	*x = EntryPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryPoint) ProtoMessage() {}

func (x *EntryPoint) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryPoint.ProtoReflect.Descriptor instead.
func (*EntryPoint) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{0}
}

func (x *EntryPoint) GetSelector() *Felt252 {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *EntryPoint) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Cairo0Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abi          string        `protobuf:"bytes,1,opt,name=abi,proto3" json:"abi,omitempty"`
	Externals    []*EntryPoint `protobuf:"bytes,2,rep,name=externals,proto3" json:"externals,omitempty"`
	L1Handlers   []*EntryPoint `protobuf:"bytes,3,rep,name=l1_handlers,json=l1Handlers,proto3" json:"l1_handlers,omitempty"`
	Constructors []*EntryPoint `protobuf:"bytes,4,rep,name=constructors,proto3" json:"constructors,omitempty"`
	// Compressed in base64 representation.
	Program string `protobuf:"bytes,5,opt,name=program,proto3" json:"program,omitempty"`
}

func (x *Cairo0Class) Reset() {
	*x = Cairo0Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cairo0Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cairo0Class) ProtoMessage() {}

func (x *Cairo0Class) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cairo0Class.ProtoReflect.Descriptor instead.
func (*Cairo0Class) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{1}
}

func (x *Cairo0Class) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

func (x *Cairo0Class) GetExternals() []*EntryPoint {
	if x != nil {
		return x.Externals
	}
	return nil
}

func (x *Cairo0Class) GetL1Handlers() []*EntryPoint {
	if x != nil {
		return x.L1Handlers
	}
	return nil
}

func (x *Cairo0Class) GetConstructors() []*EntryPoint {
	if x != nil {
		return x.Constructors
	}
	return nil
}

func (x *Cairo0Class) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

type SierraEntryPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Selector *Felt252 `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *SierraEntryPoint) Reset() {
	*x = SierraEntryPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SierraEntryPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SierraEntryPoint) ProtoMessage() {}

func (x *SierraEntryPoint) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SierraEntryPoint.ProtoReflect.Descriptor instead.
func (*SierraEntryPoint) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{2}
}

func (x *SierraEntryPoint) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SierraEntryPoint) GetSelector() *Felt252 {
	if x != nil {
		return x.Selector
	}
	return nil
}

type Cairo1EntryPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Externals    []*SierraEntryPoint `protobuf:"bytes,1,rep,name=externals,proto3" json:"externals,omitempty"`
	L1Handlers   []*SierraEntryPoint `protobuf:"bytes,2,rep,name=l1_handlers,json=l1Handlers,proto3" json:"l1_handlers,omitempty"`
	Constructors []*SierraEntryPoint `protobuf:"bytes,3,rep,name=constructors,proto3" json:"constructors,omitempty"`
}

func (x *Cairo1EntryPoints) Reset() {
	*x = Cairo1EntryPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cairo1EntryPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cairo1EntryPoints) ProtoMessage() {}

func (x *Cairo1EntryPoints) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cairo1EntryPoints.ProtoReflect.Descriptor instead.
func (*Cairo1EntryPoints) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{3}
}

func (x *Cairo1EntryPoints) GetExternals() []*SierraEntryPoint {
	if x != nil {
		return x.Externals
	}
	return nil
}

func (x *Cairo1EntryPoints) GetL1Handlers() []*SierraEntryPoint {
	if x != nil {
		return x.L1Handlers
	}
	return nil
}

func (x *Cairo1EntryPoints) GetConstructors() []*SierraEntryPoint {
	if x != nil {
		return x.Constructors
	}
	return nil
}

type Cairo1Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abi                  string             `protobuf:"bytes,1,opt,name=abi,proto3" json:"abi,omitempty"`
	EntryPoints          *Cairo1EntryPoints `protobuf:"bytes,2,opt,name=entry_points,json=entryPoints,proto3" json:"entry_points,omitempty"`
	Program              []*Felt252         `protobuf:"bytes,3,rep,name=program,proto3" json:"program,omitempty"`
	ContractClassVersion string             `protobuf:"bytes,4,opt,name=contract_class_version,json=contractClassVersion,proto3" json:"contract_class_version,omitempty"`
}

func (x *Cairo1Class) Reset() {
	*x = Cairo1Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cairo1Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cairo1Class) ProtoMessage() {}

func (x *Cairo1Class) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cairo1Class.ProtoReflect.Descriptor instead.
func (*Cairo1Class) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{4}
}

func (x *Cairo1Class) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

func (x *Cairo1Class) GetEntryPoints() *Cairo1EntryPoints {
	if x != nil {
		return x.EntryPoints
	}
	return nil
}

func (x *Cairo1Class) GetProgram() []*Felt252 {
	if x != nil {
		return x.Program
	}
	return nil
}

func (x *Cairo1Class) GetContractClassVersion() string {
	if x != nil {
		return x.ContractClassVersion
	}
	return ""
}

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Class:
	//
	//	*Class_Cairo0
	//	*Class_Cairo1
	Class     isClass_Class `protobuf_oneof:"class"`
	Domain    uint32        `protobuf:"varint,3,opt,name=domain,proto3" json:"domain,omitempty"`
	ClassHash *Hash         `protobuf:"bytes,4,opt,name=class_hash,json=classHash,proto3" json:"class_hash,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{5}
}

func (m *Class) GetClass() isClass_Class {
	if m != nil {
		return m.Class
	}
	return nil
}

func (x *Class) GetCairo0() *Cairo0Class {
	if x, ok := x.GetClass().(*Class_Cairo0); ok {
		return x.Cairo0
	}
	return nil
}

func (x *Class) GetCairo1() *Cairo1Class {
	if x, ok := x.GetClass().(*Class_Cairo1); ok {
		return x.Cairo1
	}
	return nil
}

func (x *Class) GetDomain() uint32 {
	if x != nil {
		return x.Domain
	}
	return 0
}

func (x *Class) GetClassHash() *Hash {
	if x != nil {
		return x.ClassHash
	}
	return nil
}

type isClass_Class interface {
	isClass_Class()
}

type Class_Cairo0 struct {
	Cairo0 *Cairo0Class `protobuf:"bytes,1,opt,name=cairo0,proto3,oneof"`
}

type Class_Cairo1 struct {
	Cairo1 *Cairo1Class `protobuf:"bytes,2,opt,name=cairo1,proto3,oneof"`
}

func (*Class_Cairo0) isClass_Class() {}

func (*Class_Cairo1) isClass_Class() {}

type ClassesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iteration *Iteration `protobuf:"bytes,1,opt,name=iteration,proto3" json:"iteration,omitempty"`
}

func (x *ClassesRequest) Reset() {
	*x = ClassesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassesRequest) ProtoMessage() {}

func (x *ClassesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassesRequest.ProtoReflect.Descriptor instead.
func (*ClassesRequest) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{6}
}

func (x *ClassesRequest) GetIteration() *Iteration {
	if x != nil {
		return x.Iteration
	}
	return nil
}

// Responses are sent ordered by the order given in the request.
type ClassesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClassMessage:
	//
	//	*ClassesResponse_Class
	//	*ClassesResponse_Fin
	ClassMessage isClassesResponse_ClassMessage `protobuf_oneof:"class_message"`
}

func (x *ClassesResponse) Reset() {
	*x = ClassesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_class_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassesResponse) ProtoMessage() {}

func (x *ClassesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_class_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassesResponse.ProtoReflect.Descriptor instead.
func (*ClassesResponse) Descriptor() ([]byte, []int) {
	return file_p2p_proto_class_proto_rawDescGZIP(), []int{7}
}

func (m *ClassesResponse) GetClassMessage() isClassesResponse_ClassMessage {
	if m != nil {
		return m.ClassMessage
	}
	return nil
}

func (x *ClassesResponse) GetClass() *Class {
	if x, ok := x.GetClassMessage().(*ClassesResponse_Class); ok {
		return x.Class
	}
	return nil
}

func (x *ClassesResponse) GetFin() *Fin {
	if x, ok := x.GetClassMessage().(*ClassesResponse_Fin); ok {
		return x.Fin
	}
	return nil
}

type isClassesResponse_ClassMessage interface {
	isClassesResponse_ClassMessage()
}

type ClassesResponse_Class struct {
	Class *Class `protobuf:"bytes,1,opt,name=class,proto3,oneof"`
}

type ClassesResponse_Fin struct {
	Fin *Fin `protobuf:"bytes,2,opt,name=fin,proto3,oneof"` // Fin is sent after the peer sent all the data or when it encountered a block that it doesn't have its classes.
}

func (*ClassesResponse_Class) isClassesResponse_ClassMessage() {}

func (*ClassesResponse_Fin) isClassesResponse_ClassMessage() {}

var File_p2p_proto_class_proto protoreflect.FileDescriptor

var file_p2p_proto_class_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x46, 0x65, 0x6c, 0x74, 0x32, 0x35, 0x32, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x69, 0x72, 0x6f, 0x30, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x62, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x62, 0x69, 0x12, 0x29, 0x0a,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x0b, 0x6c, 0x31, 0x5f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x6c, 0x31, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x22, 0x4e, 0x0a, 0x10, 0x53, 0x69, 0x65, 0x72, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x08, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x46, 0x65, 0x6c, 0x74, 0x32, 0x35, 0x32, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x69, 0x72, 0x6f, 0x31, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x69, 0x65,
	0x72, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x0b, 0x6c, 0x31, 0x5f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x53, 0x69, 0x65, 0x72, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x0a, 0x6c, 0x31, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x69, 0x65, 0x72, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x69, 0x72, 0x6f, 0x31, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x62, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x62, 0x69, 0x12, 0x35, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x61,
	0x69, 0x72, 0x6f, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x46, 0x65, 0x6c, 0x74, 0x32, 0x35, 0x32, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x34, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x26, 0x0a, 0x06, 0x63, 0x61, 0x69, 0x72, 0x6f, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x43, 0x61, 0x69, 0x72, 0x6f, 0x30, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x61, 0x69, 0x72, 0x6f, 0x30, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x61, 0x69, 0x72,
	0x6f, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x61, 0x69, 0x72, 0x6f,
	0x31, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x48, 0x00, 0x52, 0x06, 0x63, 0x61, 0x69, 0x72, 0x6f, 0x31,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x48, 0x61, 0x73, 0x68, 0x42, 0x07,
	0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x3a, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x03, 0x66, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x46, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x66, 0x69, 0x6e,
	0x42, 0x0f, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x45, 0x74, 0x68, 0x2f, 0x6a, 0x75,
	0x6e, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2f,
	0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_proto_class_proto_rawDescOnce sync.Once
	file_p2p_proto_class_proto_rawDescData = file_p2p_proto_class_proto_rawDesc
)

func file_p2p_proto_class_proto_rawDescGZIP() []byte {
	file_p2p_proto_class_proto_rawDescOnce.Do(func() {
		file_p2p_proto_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_proto_class_proto_rawDescData)
	})
	return file_p2p_proto_class_proto_rawDescData
}

var file_p2p_proto_class_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_p2p_proto_class_proto_goTypes = []any{
	(*EntryPoint)(nil),        // 0: EntryPoint
	(*Cairo0Class)(nil),       // 1: Cairo0Class
	(*SierraEntryPoint)(nil),  // 2: SierraEntryPoint
	(*Cairo1EntryPoints)(nil), // 3: Cairo1EntryPoints
	(*Cairo1Class)(nil),       // 4: Cairo1Class
	(*Class)(nil),             // 5: Class
	(*ClassesRequest)(nil),    // 6: ClassesRequest
	(*ClassesResponse)(nil),   // 7: ClassesResponse
	(*Felt252)(nil),           // 8: Felt252
	(*Hash)(nil),              // 9: Hash
	(*Iteration)(nil),         // 10: Iteration
	(*Fin)(nil),               // 11: Fin
}
var file_p2p_proto_class_proto_depIdxs = []int32{
	8,  // 0: EntryPoint.selector:type_name -> Felt252
	0,  // 1: Cairo0Class.externals:type_name -> EntryPoint
	0,  // 2: Cairo0Class.l1_handlers:type_name -> EntryPoint
	0,  // 3: Cairo0Class.constructors:type_name -> EntryPoint
	8,  // 4: SierraEntryPoint.selector:type_name -> Felt252
	2,  // 5: Cairo1EntryPoints.externals:type_name -> SierraEntryPoint
	2,  // 6: Cairo1EntryPoints.l1_handlers:type_name -> SierraEntryPoint
	2,  // 7: Cairo1EntryPoints.constructors:type_name -> SierraEntryPoint
	3,  // 8: Cairo1Class.entry_points:type_name -> Cairo1EntryPoints
	8,  // 9: Cairo1Class.program:type_name -> Felt252
	1,  // 10: Class.cairo0:type_name -> Cairo0Class
	4,  // 11: Class.cairo1:type_name -> Cairo1Class
	9,  // 12: Class.class_hash:type_name -> Hash
	10, // 13: ClassesRequest.iteration:type_name -> Iteration
	5,  // 14: ClassesResponse.class:type_name -> Class
	11, // 15: ClassesResponse.fin:type_name -> Fin
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_p2p_proto_class_proto_init() }
func file_p2p_proto_class_proto_init() {
	if File_p2p_proto_class_proto != nil {
		return
	}
	file_p2p_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_p2p_proto_class_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EntryPoint); i {
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
		file_p2p_proto_class_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Cairo0Class); i {
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
		file_p2p_proto_class_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SierraEntryPoint); i {
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
		file_p2p_proto_class_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Cairo1EntryPoints); i {
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
		file_p2p_proto_class_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Cairo1Class); i {
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
		file_p2p_proto_class_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Class); i {
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
		file_p2p_proto_class_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ClassesRequest); i {
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
		file_p2p_proto_class_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ClassesResponse); i {
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
	file_p2p_proto_class_proto_msgTypes[5].OneofWrappers = []any{
		(*Class_Cairo0)(nil),
		(*Class_Cairo1)(nil),
	}
	file_p2p_proto_class_proto_msgTypes[7].OneofWrappers = []any{
		(*ClassesResponse_Class)(nil),
		(*ClassesResponse_Fin)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_proto_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_proto_class_proto_goTypes,
		DependencyIndexes: file_p2p_proto_class_proto_depIdxs,
		MessageInfos:      file_p2p_proto_class_proto_msgTypes,
	}.Build()
	File_p2p_proto_class_proto = out.File
	file_p2p_proto_class_proto_rawDesc = nil
	file_p2p_proto_class_proto_goTypes = nil
	file_p2p_proto_class_proto_depIdxs = nil
}