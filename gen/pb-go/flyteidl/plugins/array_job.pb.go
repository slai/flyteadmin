// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: flyteidl/plugins/array_job.proto

package plugins

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

// Describes a job that can process independent pieces of data concurrently. Multiple copies of the runnable component
// will be executed concurrently.
type ArrayJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the minimum number of instances to bring up concurrently at any given point. Note that this is an
	// optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
	// running instances might be more. This has to be a positive number if assigned. Default value is size.
	Parallelism int64 `protobuf:"varint,1,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	// Defines the number of instances to launch at most. This number should match the size of the input if the job
	// requires processing of all input data. This has to be a positive number.
	// In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// Types that are assignable to SuccessCriteria:
	//	*ArrayJob_MinSuccesses
	//	*ArrayJob_MinSuccessRatio
	SuccessCriteria isArrayJob_SuccessCriteria `protobuf_oneof:"success_criteria"`
}

func (x *ArrayJob) Reset() {
	*x = ArrayJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_array_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayJob) ProtoMessage() {}

func (x *ArrayJob) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_array_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayJob.ProtoReflect.Descriptor instead.
func (*ArrayJob) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_array_job_proto_rawDescGZIP(), []int{0}
}

func (x *ArrayJob) GetParallelism() int64 {
	if x != nil {
		return x.Parallelism
	}
	return 0
}

func (x *ArrayJob) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (m *ArrayJob) GetSuccessCriteria() isArrayJob_SuccessCriteria {
	if m != nil {
		return m.SuccessCriteria
	}
	return nil
}

func (x *ArrayJob) GetMinSuccesses() int64 {
	if x, ok := x.GetSuccessCriteria().(*ArrayJob_MinSuccesses); ok {
		return x.MinSuccesses
	}
	return 0
}

func (x *ArrayJob) GetMinSuccessRatio() float32 {
	if x, ok := x.GetSuccessCriteria().(*ArrayJob_MinSuccessRatio); ok {
		return x.MinSuccessRatio
	}
	return 0
}

type isArrayJob_SuccessCriteria interface {
	isArrayJob_SuccessCriteria()
}

type ArrayJob_MinSuccesses struct {
	// An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
	// the array job will be marked as successful and outputs will be computed. This has to be a non-negative number if
	// assigned. Default value is size (if specified).
	MinSuccesses int64 `protobuf:"varint,3,opt,name=min_successes,json=minSuccesses,proto3,oneof"`
}

type ArrayJob_MinSuccessRatio struct {
	// If the array job size is not known beforehand, the min_success_ratio can instead be used to determine when an array
	// job can be marked successful.
	MinSuccessRatio float32 `protobuf:"fixed32,4,opt,name=min_success_ratio,json=minSuccessRatio,proto3,oneof"`
}

func (*ArrayJob_MinSuccesses) isArrayJob_SuccessCriteria() {}

func (*ArrayJob_MinSuccessRatio) isArrayJob_SuccessCriteria() {}

var File_flyteidl_plugins_array_job_proto protoreflect.FileDescriptor

var file_flyteidl_plugins_array_job_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x08, 0x41, 0x72, 0x72, 0x61, 0x79, 0x4a, 0x6f,
	0x62, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c,
	0x69, 0x73, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2c,
	0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x69, 0x6e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x12, 0x0a, 0x10,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x42, 0xbf, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x42, 0x0d, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d,
	0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x58, 0xaa, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0xca, 0x02, 0x10, 0x46,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0xe2,
	0x02, 0x1c, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_plugins_array_job_proto_rawDescOnce sync.Once
	file_flyteidl_plugins_array_job_proto_rawDescData = file_flyteidl_plugins_array_job_proto_rawDesc
)

func file_flyteidl_plugins_array_job_proto_rawDescGZIP() []byte {
	file_flyteidl_plugins_array_job_proto_rawDescOnce.Do(func() {
		file_flyteidl_plugins_array_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_plugins_array_job_proto_rawDescData)
	})
	return file_flyteidl_plugins_array_job_proto_rawDescData
}

var file_flyteidl_plugins_array_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyteidl_plugins_array_job_proto_goTypes = []interface{}{
	(*ArrayJob)(nil), // 0: flyteidl.plugins.ArrayJob
}
var file_flyteidl_plugins_array_job_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flyteidl_plugins_array_job_proto_init() }
func file_flyteidl_plugins_array_job_proto_init() {
	if File_flyteidl_plugins_array_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_plugins_array_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayJob); i {
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
	file_flyteidl_plugins_array_job_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ArrayJob_MinSuccesses)(nil),
		(*ArrayJob_MinSuccessRatio)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_plugins_array_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_plugins_array_job_proto_goTypes,
		DependencyIndexes: file_flyteidl_plugins_array_job_proto_depIdxs,
		MessageInfos:      file_flyteidl_plugins_array_job_proto_msgTypes,
	}.Build()
	File_flyteidl_plugins_array_job_proto = out.File
	file_flyteidl_plugins_array_job_proto_rawDesc = nil
	file_flyteidl_plugins_array_job_proto_goTypes = nil
	file_flyteidl_plugins_array_job_proto_depIdxs = nil
}
