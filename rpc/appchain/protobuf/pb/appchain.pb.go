// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: protobuf/appchain.proto

package pb

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

type L1StakerRewardsAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId       string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StakerAddress string `protobuf:"bytes,2,opt,name=staker_address,json=stakerAddress,proto3" json:"staker_address,omitempty"`
	Strategies    string `protobuf:"bytes,3,opt,name=strategies,proto3" json:"strategies,omitempty"`
}

func (x *L1StakerRewardsAmountRequest) Reset() {
	*x = L1StakerRewardsAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_appchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L1StakerRewardsAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L1StakerRewardsAmountRequest) ProtoMessage() {}

func (x *L1StakerRewardsAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_appchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L1StakerRewardsAmountRequest.ProtoReflect.Descriptor instead.
func (*L1StakerRewardsAmountRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_appchain_proto_rawDescGZIP(), []int{0}
}

func (x *L1StakerRewardsAmountRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *L1StakerRewardsAmountRequest) GetStakerAddress() string {
	if x != nil {
		return x.StakerAddress
	}
	return ""
}

func (x *L1StakerRewardsAmountRequest) GetStrategies() string {
	if x != nil {
		return x.Strategies
	}
	return ""
}

type L1StakerRewardsAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Income  string `protobuf:"bytes,3,opt,name=income,proto3" json:"income,omitempty"`
}

func (x *L1StakerRewardsAmountResponse) Reset() {
	*x = L1StakerRewardsAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_appchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L1StakerRewardsAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L1StakerRewardsAmountResponse) ProtoMessage() {}

func (x *L1StakerRewardsAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_appchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L1StakerRewardsAmountResponse.ProtoReflect.Descriptor instead.
func (*L1StakerRewardsAmountResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_appchain_proto_rawDescGZIP(), []int{1}
}

func (x *L1StakerRewardsAmountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *L1StakerRewardsAmountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *L1StakerRewardsAmountResponse) GetIncome() string {
	if x != nil {
		return x.Income
	}
	return ""
}

type L2StakerRewardsAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId       string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StakerAddress string `protobuf:"bytes,2,opt,name=staker_address,json=stakerAddress,proto3" json:"staker_address,omitempty"`
	Strategy      string `protobuf:"bytes,3,opt,name=strategy,proto3" json:"strategy,omitempty"`
}

func (x *L2StakerRewardsAmountRequest) Reset() {
	*x = L2StakerRewardsAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_appchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L2StakerRewardsAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L2StakerRewardsAmountRequest) ProtoMessage() {}

func (x *L2StakerRewardsAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_appchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L2StakerRewardsAmountRequest.ProtoReflect.Descriptor instead.
func (*L2StakerRewardsAmountRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_appchain_proto_rawDescGZIP(), []int{2}
}

func (x *L2StakerRewardsAmountRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *L2StakerRewardsAmountRequest) GetStakerAddress() string {
	if x != nil {
		return x.StakerAddress
	}
	return ""
}

func (x *L2StakerRewardsAmountRequest) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

type L2StakerRewardsAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Income  string `protobuf:"bytes,3,opt,name=income,proto3" json:"income,omitempty"`
}

func (x *L2StakerRewardsAmountResponse) Reset() {
	*x = L2StakerRewardsAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_appchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L2StakerRewardsAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L2StakerRewardsAmountResponse) ProtoMessage() {}

func (x *L2StakerRewardsAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_appchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L2StakerRewardsAmountResponse.ProtoReflect.Descriptor instead.
func (*L2StakerRewardsAmountResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_appchain_proto_rawDescGZIP(), []int{3}
}

func (x *L2StakerRewardsAmountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *L2StakerRewardsAmountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *L2StakerRewardsAmountResponse) GetIncome() string {
	if x != nil {
		return x.Income
	}
	return ""
}

var File_protobuf_appchain_proto protoreflect.FileDescriptor

var file_protobuf_appchain_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x63, 0x6f, 0x72, 0x75,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x80,
	0x01, 0x0a, 0x1c, 0x4c, 0x31, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65,
	0x73, 0x22, 0x6b, 0x0a, 0x1d, 0x4c, 0x31, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x7c,
	0x0a, 0x1c, 0x4c, 0x32, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x6b, 0x0a, 0x1d,
	0x4c, 0x32, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x32, 0x91, 0x02, 0x0a, 0x0f, 0x41, 0x70,
	0x70, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a,
	0x15, 0x4c, 0x31, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x61, 0x63, 0x6f, 0x72, 0x75, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x31, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x6f, 0x72,
	0x75, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x4c, 0x31, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a,
	0x15, 0x4c, 0x32, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x61, 0x63, 0x6f, 0x72, 0x75, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x32, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x6f, 0x72,
	0x75, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x4c, 0x32, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_appchain_proto_rawDescOnce sync.Once
	file_protobuf_appchain_proto_rawDescData = file_protobuf_appchain_proto_rawDesc
)

func file_protobuf_appchain_proto_rawDescGZIP() []byte {
	file_protobuf_appchain_proto_rawDescOnce.Do(func() {
		file_protobuf_appchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_appchain_proto_rawDescData)
	})
	return file_protobuf_appchain_proto_rawDescData
}

var file_protobuf_appchain_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_appchain_proto_goTypes = []interface{}{
	(*L1StakerRewardsAmountRequest)(nil),  // 0: acorus.rpc.appchain.L1StakerRewardsAmountRequest
	(*L1StakerRewardsAmountResponse)(nil), // 1: acorus.rpc.appchain.L1StakerRewardsAmountResponse
	(*L2StakerRewardsAmountRequest)(nil),  // 2: acorus.rpc.appchain.L2StakerRewardsAmountRequest
	(*L2StakerRewardsAmountResponse)(nil), // 3: acorus.rpc.appchain.L2StakerRewardsAmountResponse
}
var file_protobuf_appchain_proto_depIdxs = []int32{
	0, // 0: acorus.rpc.appchain.AppChainService.L1StakerRewardsAmount:input_type -> acorus.rpc.appchain.L1StakerRewardsAmountRequest
	2, // 1: acorus.rpc.appchain.AppChainService.L2StakerRewardsAmount:input_type -> acorus.rpc.appchain.L2StakerRewardsAmountRequest
	1, // 2: acorus.rpc.appchain.AppChainService.L1StakerRewardsAmount:output_type -> acorus.rpc.appchain.L1StakerRewardsAmountResponse
	3, // 3: acorus.rpc.appchain.AppChainService.L2StakerRewardsAmount:output_type -> acorus.rpc.appchain.L2StakerRewardsAmountResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_appchain_proto_init() }
func file_protobuf_appchain_proto_init() {
	if File_protobuf_appchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_appchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L1StakerRewardsAmountRequest); i {
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
		file_protobuf_appchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L1StakerRewardsAmountResponse); i {
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
		file_protobuf_appchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L2StakerRewardsAmountRequest); i {
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
		file_protobuf_appchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L2StakerRewardsAmountResponse); i {
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
			RawDescriptor: file_protobuf_appchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_appchain_proto_goTypes,
		DependencyIndexes: file_protobuf_appchain_proto_depIdxs,
		MessageInfos:      file_protobuf_appchain_proto_msgTypes,
	}.Build()
	File_protobuf_appchain_proto = out.File
	file_protobuf_appchain_proto_rawDesc = nil
	file_protobuf_appchain_proto_goTypes = nil
	file_protobuf_appchain_proto_depIdxs = nil
}
