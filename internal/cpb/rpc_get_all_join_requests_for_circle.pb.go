// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_get_all_join_requests_for_circle.proto

package cpb

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JoinRequestId int64  `protobuf:"varint,1,opt,name=join_request_id,json=joinRequestId,proto3" json:"join_request_id,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CircleId      string `protobuf:"bytes,3,opt,name=circle_id,json=circleId,proto3" json:"circle_id,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_all_join_requests_for_circle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_all_join_requests_for_circle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_all_join_requests_for_circle_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetJoinRequestId() int64 {
	if x != nil {
		return x.JoinRequestId
	}
	return 0
}

func (x *JoinRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *JoinRequest) GetCircleId() string {
	if x != nil {
		return x.CircleId
	}
	return ""
}

type GetCircleJoinRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CircleId string `protobuf:"bytes,1,opt,name=circle_id,json=circleId,proto3" json:"circle_id,omitempty"`
	Page     int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetCircleJoinRequestsRequest) Reset() {
	*x = GetCircleJoinRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_all_join_requests_for_circle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCircleJoinRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCircleJoinRequestsRequest) ProtoMessage() {}

func (x *GetCircleJoinRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_all_join_requests_for_circle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCircleJoinRequestsRequest.ProtoReflect.Descriptor instead.
func (*GetCircleJoinRequestsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_all_join_requests_for_circle_proto_rawDescGZIP(), []int{1}
}

func (x *GetCircleJoinRequestsRequest) GetCircleId() string {
	if x != nil {
		return x.CircleId
	}
	return ""
}

func (x *GetCircleJoinRequestsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetCircleJoinRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JoinRequests []*JoinRequest `protobuf:"bytes,1,rep,name=join_requests,json=joinRequests,proto3" json:"join_requests,omitempty"`
}

func (x *GetCircleJoinRequestsResponse) Reset() {
	*x = GetCircleJoinRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_all_join_requests_for_circle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCircleJoinRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCircleJoinRequestsResponse) ProtoMessage() {}

func (x *GetCircleJoinRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_all_join_requests_for_circle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCircleJoinRequestsResponse.ProtoReflect.Descriptor instead.
func (*GetCircleJoinRequestsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_all_join_requests_for_circle_proto_rawDescGZIP(), []int{2}
}

func (x *GetCircleJoinRequestsResponse) GetJoinRequests() []*JoinRequest {
	if x != nil {
		return x.JoinRequests
	}
	return nil
}

var File_rpc_get_all_join_requests_for_circle_proto protoreflect.FileDescriptor

var file_rpc_get_all_join_requests_for_circle_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6a, 0x6f,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f,
	0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x70,
	0x62, 0x22, 0x6b, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x4f,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x56, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x70, 0x62, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x6a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x76, 0x61, 0x73, 0x68, 0x2f, 0x62, 0x67, 0x6d, 0x6f,
	0x6f, 0x64, 0x2d, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_all_join_requests_for_circle_proto_rawDescOnce sync.Once
	file_rpc_get_all_join_requests_for_circle_proto_rawDescData = file_rpc_get_all_join_requests_for_circle_proto_rawDesc
)

func file_rpc_get_all_join_requests_for_circle_proto_rawDescGZIP() []byte {
	file_rpc_get_all_join_requests_for_circle_proto_rawDescOnce.Do(func() {
		file_rpc_get_all_join_requests_for_circle_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_all_join_requests_for_circle_proto_rawDescData)
	})
	return file_rpc_get_all_join_requests_for_circle_proto_rawDescData
}

var file_rpc_get_all_join_requests_for_circle_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_get_all_join_requests_for_circle_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),                   // 0: cpb.JoinRequest
	(*GetCircleJoinRequestsRequest)(nil),  // 1: cpb.GetCircleJoinRequestsRequest
	(*GetCircleJoinRequestsResponse)(nil), // 2: cpb.GetCircleJoinRequestsResponse
}
var file_rpc_get_all_join_requests_for_circle_proto_depIdxs = []int32{
	0, // 0: cpb.GetCircleJoinRequestsResponse.join_requests:type_name -> cpb.JoinRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_get_all_join_requests_for_circle_proto_init() }
func file_rpc_get_all_join_requests_for_circle_proto_init() {
	if File_rpc_get_all_join_requests_for_circle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_all_join_requests_for_circle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_rpc_get_all_join_requests_for_circle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCircleJoinRequestsRequest); i {
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
		file_rpc_get_all_join_requests_for_circle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCircleJoinRequestsResponse); i {
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
			RawDescriptor: file_rpc_get_all_join_requests_for_circle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_all_join_requests_for_circle_proto_goTypes,
		DependencyIndexes: file_rpc_get_all_join_requests_for_circle_proto_depIdxs,
		MessageInfos:      file_rpc_get_all_join_requests_for_circle_proto_msgTypes,
	}.Build()
	File_rpc_get_all_join_requests_for_circle_proto = out.File
	file_rpc_get_all_join_requests_for_circle_proto_rawDesc = nil
	file_rpc_get_all_join_requests_for_circle_proto_goTypes = nil
	file_rpc_get_all_join_requests_for_circle_proto_depIdxs = nil
}
