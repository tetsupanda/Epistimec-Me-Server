// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/models/intervention.proto

package models

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

// An Intervention is a data model for representing
// a prescribed action. An Intervention can be bound
// to a Habit as a "prescription" for a Causal Belief.
//
// Example:
// id: 1
// name: "drinking water"
// description: "Drinking water upon waking."
type Intervention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                  // Unique identifier for the intervention
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // Name of the intervention
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // Description of the intervention
}

func (x *Intervention) Reset() {
	*x = Intervention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_intervention_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intervention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intervention) ProtoMessage() {}

func (x *Intervention) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_intervention_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intervention.ProtoReflect.Descriptor instead.
func (*Intervention) Descriptor() ([]byte, []int) {
	return file_proto_models_intervention_proto_rawDescGZIP(), []int{0}
}

func (x *Intervention) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Intervention) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Intervention) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Message representing a habit
// A habit represents a relationship between
// some causal belief (e.g., "I believe that drinking water improves my mental clarity.")
// and an intervention (e.g., "Drink 10 cups of water a day.").
//
// Example:
// id: 1
// name: "Drink Water Upon Waking"
// intervention_id: 1
// intervention_name: "drinking water"
// causal_belief_id: 1
// causal_belief_description: "If I drink water, I will have more mental clarity."
// description: "Drinking water upon waking will improve hydration and mental clarity throughout the day."
type Habit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                                           // Unique identifier for the habit
	Name                    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                                        // Name of the habit
	InterventionId          int32  `protobuf:"varint,3,opt,name=intervention_id,json=interventionId,proto3" json:"intervention_id,omitempty"`                             // ID of the related intervention
	InterventionName        string `protobuf:"bytes,4,opt,name=intervention_name,json=interventionName,proto3" json:"intervention_name,omitempty"`                        // External name for the intervention
	CausalBeliefId          int32  `protobuf:"varint,5,opt,name=causal_belief_id,json=causalBeliefId,proto3" json:"causal_belief_id,omitempty"`                           // ID of the related causal belief
	CausalBeliefDescription string `protobuf:"bytes,6,opt,name=causal_belief_description,json=causalBeliefDescription,proto3" json:"causal_belief_description,omitempty"` // External description for the causal belief
	Description             string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`                                                          // Detailed description of the habit
}

func (x *Habit) Reset() {
	*x = Habit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_intervention_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habit) ProtoMessage() {}

func (x *Habit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_intervention_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habit.ProtoReflect.Descriptor instead.
func (*Habit) Descriptor() ([]byte, []int) {
	return file_proto_models_intervention_proto_rawDescGZIP(), []int{1}
}

func (x *Habit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Habit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Habit) GetInterventionId() int32 {
	if x != nil {
		return x.InterventionId
	}
	return 0
}

func (x *Habit) GetInterventionName() string {
	if x != nil {
		return x.InterventionName
	}
	return ""
}

func (x *Habit) GetCausalBeliefId() int32 {
	if x != nil {
		return x.CausalBeliefId
	}
	return 0
}

func (x *Habit) GetCausalBeliefDescription() string {
	if x != nil {
		return x.CausalBeliefDescription
	}
	return ""
}

func (x *Habit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_models_intervention_proto protoreflect.FileDescriptor

var file_proto_models_intervention_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x02, 0x0a, 0x05, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63,
	0x61, 0x75, 0x73, 0x61, 0x6c, 0x5f, 0x62, 0x65, 0x6c, 0x69, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x42, 0x65, 0x6c,
	0x69, 0x65, 0x66, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x63, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x5f,
	0x62, 0x65, 0x6c, 0x69, 0x65, 0x66, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x61, 0x75, 0x73, 0x61, 0x6c,
	0x42, 0x65, 0x6c, 0x69, 0x65, 0x66, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x20, 0x5a, 0x1e, 0x65, 0x70, 0x69, 0x73, 0x74, 0x65, 0x6d, 0x69, 0x63,
	0x2d, 0x6d, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_models_intervention_proto_rawDescOnce sync.Once
	file_proto_models_intervention_proto_rawDescData = file_proto_models_intervention_proto_rawDesc
)

func file_proto_models_intervention_proto_rawDescGZIP() []byte {
	file_proto_models_intervention_proto_rawDescOnce.Do(func() {
		file_proto_models_intervention_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_models_intervention_proto_rawDescData)
	})
	return file_proto_models_intervention_proto_rawDescData
}

var file_proto_models_intervention_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_models_intervention_proto_goTypes = []interface{}{
	(*Intervention)(nil), // 0: Intervention
	(*Habit)(nil),        // 1: Habit
}
var file_proto_models_intervention_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_models_intervention_proto_init() }
func file_proto_models_intervention_proto_init() {
	if File_proto_models_intervention_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_models_intervention_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intervention); i {
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
		file_proto_models_intervention_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Habit); i {
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
			RawDescriptor: file_proto_models_intervention_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_models_intervention_proto_goTypes,
		DependencyIndexes: file_proto_models_intervention_proto_depIdxs,
		MessageInfos:      file_proto_models_intervention_proto_msgTypes,
	}.Build()
	File_proto_models_intervention_proto = out.File
	file_proto_models_intervention_proto_rawDesc = nil
	file_proto_models_intervention_proto_goTypes = nil
	file_proto_models_intervention_proto_depIdxs = nil
}
