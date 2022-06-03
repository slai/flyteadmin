// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: flyteidl/core/compiler.proto

package core

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

// Adjacency list for the workflow. This is created as part of the compilation process. Every process after the compilation
// step uses this created ConnectionSet
type ConnectionSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of all the node ids that are downstream from a given node id
	Downstream map[string]*ConnectionSet_IdList `protobuf:"bytes,7,rep,name=downstream,proto3" json:"downstream,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of all the node ids, that are upstream of this node id
	Upstream map[string]*ConnectionSet_IdList `protobuf:"bytes,8,rep,name=upstream,proto3" json:"upstream,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectionSet) Reset() {
	*x = ConnectionSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_compiler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionSet) ProtoMessage() {}

func (x *ConnectionSet) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_compiler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionSet.ProtoReflect.Descriptor instead.
func (*ConnectionSet) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_compiler_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionSet) GetDownstream() map[string]*ConnectionSet_IdList {
	if x != nil {
		return x.Downstream
	}
	return nil
}

func (x *ConnectionSet) GetUpstream() map[string]*ConnectionSet_IdList {
	if x != nil {
		return x.Upstream
	}
	return nil
}

// Output of the compilation Step. This object represents one workflow. We store more metadata at this layer
type CompiledWorkflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Completely contained Workflow Template
	Template *WorkflowTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	// For internal use only! This field is used by the system and must not be filled in. Any values set will be ignored.
	Connections *ConnectionSet `protobuf:"bytes,2,opt,name=connections,proto3" json:"connections,omitempty"`
}

func (x *CompiledWorkflow) Reset() {
	*x = CompiledWorkflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_compiler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompiledWorkflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompiledWorkflow) ProtoMessage() {}

func (x *CompiledWorkflow) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_compiler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompiledWorkflow.ProtoReflect.Descriptor instead.
func (*CompiledWorkflow) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_compiler_proto_rawDescGZIP(), []int{1}
}

func (x *CompiledWorkflow) GetTemplate() *WorkflowTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *CompiledWorkflow) GetConnections() *ConnectionSet {
	if x != nil {
		return x.Connections
	}
	return nil
}

// Output of the Compilation step. This object represent one Task. We store more metadata at this layer
type CompiledTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Completely contained TaskTemplate
	Template *TaskTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *CompiledTask) Reset() {
	*x = CompiledTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_compiler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompiledTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompiledTask) ProtoMessage() {}

func (x *CompiledTask) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_compiler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompiledTask.ProtoReflect.Descriptor instead.
func (*CompiledTask) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_compiler_proto_rawDescGZIP(), []int{2}
}

func (x *CompiledTask) GetTemplate() *TaskTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

// A Compiled Workflow Closure contains all the information required to start a new execution, or to visualize a workflow
// and its details. The CompiledWorkflowClosure should always contain a primary workflow, that is the main workflow that
// will being the execution. All subworkflows are denormalized. WorkflowNodes refer to the workflow identifiers of
// compiled subworkflows.
type CompiledWorkflowClosure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//+required
	Primary *CompiledWorkflow `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	// Guaranteed that there will only exist one and only one workflow with a given id, i.e., every sub workflow has a
	// unique identifier. Also every enclosed subworkflow is used either by a primary workflow or by a subworkflow
	// as an inlined workflow
	//+optional
	SubWorkflows []*CompiledWorkflow `protobuf:"bytes,2,rep,name=sub_workflows,json=subWorkflows,proto3" json:"sub_workflows,omitempty"`
	// Guaranteed that there will only exist one and only one task with a given id, i.e., every task has a unique id
	//+required (at least 1)
	Tasks []*CompiledTask `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *CompiledWorkflowClosure) Reset() {
	*x = CompiledWorkflowClosure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_compiler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompiledWorkflowClosure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompiledWorkflowClosure) ProtoMessage() {}

func (x *CompiledWorkflowClosure) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_compiler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompiledWorkflowClosure.ProtoReflect.Descriptor instead.
func (*CompiledWorkflowClosure) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_compiler_proto_rawDescGZIP(), []int{3}
}

func (x *CompiledWorkflowClosure) GetPrimary() *CompiledWorkflow {
	if x != nil {
		return x.Primary
	}
	return nil
}

func (x *CompiledWorkflowClosure) GetSubWorkflows() []*CompiledWorkflow {
	if x != nil {
		return x.SubWorkflows
	}
	return nil
}

func (x *CompiledWorkflowClosure) GetTasks() []*CompiledTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ConnectionSet_IdList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ConnectionSet_IdList) Reset() {
	*x = ConnectionSet_IdList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_compiler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionSet_IdList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionSet_IdList) ProtoMessage() {}

func (x *ConnectionSet_IdList) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_compiler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionSet_IdList.ProtoReflect.Descriptor instead.
func (*ConnectionSet_IdList) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_compiler_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConnectionSet_IdList) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_flyteidl_core_compiler_proto protoreflect.FileDescriptor

var file_flyteidl_core_compiler_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1c, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x46, 0x0a, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1a,
	0x0a, 0x06, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x1a, 0x62, 0x0a, 0x0f, 0x44, 0x6f,
	0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x2e, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60,
	0x0a, 0x0d, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x2e, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x8f, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x3b, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x17,
	0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x43, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x42, 0xad, 0x01, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x42, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x46, 0x43, 0x58, 0xaa, 0x02,
	0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0xca, 0x02,
	0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0xe2, 0x02,
	0x19, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x46, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_core_compiler_proto_rawDescOnce sync.Once
	file_flyteidl_core_compiler_proto_rawDescData = file_flyteidl_core_compiler_proto_rawDesc
)

func file_flyteidl_core_compiler_proto_rawDescGZIP() []byte {
	file_flyteidl_core_compiler_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_compiler_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_compiler_proto_rawDescData)
	})
	return file_flyteidl_core_compiler_proto_rawDescData
}

var file_flyteidl_core_compiler_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flyteidl_core_compiler_proto_goTypes = []interface{}{
	(*ConnectionSet)(nil),           // 0: flyteidl.core.ConnectionSet
	(*CompiledWorkflow)(nil),        // 1: flyteidl.core.CompiledWorkflow
	(*CompiledTask)(nil),            // 2: flyteidl.core.CompiledTask
	(*CompiledWorkflowClosure)(nil), // 3: flyteidl.core.CompiledWorkflowClosure
	(*ConnectionSet_IdList)(nil),    // 4: flyteidl.core.ConnectionSet.IdList
	nil,                             // 5: flyteidl.core.ConnectionSet.DownstreamEntry
	nil,                             // 6: flyteidl.core.ConnectionSet.UpstreamEntry
	(*WorkflowTemplate)(nil),        // 7: flyteidl.core.WorkflowTemplate
	(*TaskTemplate)(nil),            // 8: flyteidl.core.TaskTemplate
}
var file_flyteidl_core_compiler_proto_depIdxs = []int32{
	5,  // 0: flyteidl.core.ConnectionSet.downstream:type_name -> flyteidl.core.ConnectionSet.DownstreamEntry
	6,  // 1: flyteidl.core.ConnectionSet.upstream:type_name -> flyteidl.core.ConnectionSet.UpstreamEntry
	7,  // 2: flyteidl.core.CompiledWorkflow.template:type_name -> flyteidl.core.WorkflowTemplate
	0,  // 3: flyteidl.core.CompiledWorkflow.connections:type_name -> flyteidl.core.ConnectionSet
	8,  // 4: flyteidl.core.CompiledTask.template:type_name -> flyteidl.core.TaskTemplate
	1,  // 5: flyteidl.core.CompiledWorkflowClosure.primary:type_name -> flyteidl.core.CompiledWorkflow
	1,  // 6: flyteidl.core.CompiledWorkflowClosure.sub_workflows:type_name -> flyteidl.core.CompiledWorkflow
	2,  // 7: flyteidl.core.CompiledWorkflowClosure.tasks:type_name -> flyteidl.core.CompiledTask
	4,  // 8: flyteidl.core.ConnectionSet.DownstreamEntry.value:type_name -> flyteidl.core.ConnectionSet.IdList
	4,  // 9: flyteidl.core.ConnectionSet.UpstreamEntry.value:type_name -> flyteidl.core.ConnectionSet.IdList
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_flyteidl_core_compiler_proto_init() }
func file_flyteidl_core_compiler_proto_init() {
	if File_flyteidl_core_compiler_proto != nil {
		return
	}
	file_flyteidl_core_workflow_proto_init()
	file_flyteidl_core_tasks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_compiler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionSet); i {
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
		file_flyteidl_core_compiler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompiledWorkflow); i {
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
		file_flyteidl_core_compiler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompiledTask); i {
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
		file_flyteidl_core_compiler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompiledWorkflowClosure); i {
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
		file_flyteidl_core_compiler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionSet_IdList); i {
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
			RawDescriptor: file_flyteidl_core_compiler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_compiler_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_compiler_proto_depIdxs,
		MessageInfos:      file_flyteidl_core_compiler_proto_msgTypes,
	}.Build()
	File_flyteidl_core_compiler_proto = out.File
	file_flyteidl_core_compiler_proto_rawDesc = nil
	file_flyteidl_core_compiler_proto_goTypes = nil
	file_flyteidl_core_compiler_proto_depIdxs = nil
}
