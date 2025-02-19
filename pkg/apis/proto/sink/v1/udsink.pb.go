// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pkg/apis/proto/sink/v1/udsink.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type EventTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// event_time is the time associated with each datum.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
}

func (x *EventTime) Reset() {
	*x = EventTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTime) ProtoMessage() {}

func (x *EventTime) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTime.ProtoReflect.Descriptor instead.
func (*EventTime) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{0}
}

func (x *EventTime) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

type Watermark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// watermark is the monotonically increasing time which denotes completeness for the given time for the given vertex.
	// This watermark can be used to track completeness before persisting to sink.
	Watermark *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=watermark,proto3" json:"watermark,omitempty"`
}

func (x *Watermark) Reset() {
	*x = Watermark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watermark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watermark) ProtoMessage() {}

func (x *Watermark) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watermark.ProtoReflect.Descriptor instead.
func (*Watermark) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{1}
}

func (x *Watermark) GetWatermark() *timestamppb.Timestamp {
	if x != nil {
		return x.Watermark
	}
	return nil
}

// *
// Datum represents a datum element.
type Datum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *EventTime `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Watermark *Watermark `protobuf:"bytes,4,opt,name=watermark,proto3" json:"watermark,omitempty"`
	Id        string     `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Datum) Reset() {
	*x = Datum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datum) ProtoMessage() {}

func (x *Datum) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datum.ProtoReflect.Descriptor instead.
func (*Datum) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{2}
}

func (x *Datum) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Datum) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Datum) GetEventTime() *EventTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *Datum) GetWatermark() *Watermark {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *Datum) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// DatumList represents a list of datum elements.
type DatumList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*Datum `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *DatumList) Reset() {
	*x = DatumList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumList) ProtoMessage() {}

func (x *DatumList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumList.ProtoReflect.Descriptor instead.
func (*DatumList) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{3}
}

func (x *DatumList) GetElements() []*Datum {
	if x != nil {
		return x.Elements
	}
	return nil
}

// *
// ReadyResponse is the health check result.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{4}
}

func (x *ReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

// *
// Response is the individual response of each message written to the sink.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the message, can be used to uniquely identify the message.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// success denotes the status of persisting to disk. if set to false, it means writing to sink for the message failed.
	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// err_msg is the error message, set it if success is set to false.
	ErrMsg string `protobuf:"bytes,3,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

// *
// ResponseList is the list of responses. The number of elements in this list will be equal to the number of Datum
// elements passed to the SinkFn.
type ResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*Response `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *ResponseList) Reset() {
	*x = ResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseList) ProtoMessage() {}

func (x *ResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseList.ProtoReflect.Descriptor instead.
func (*ResponseList) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseList) GetResponses() []*Response {
	if x != nil {
		return x.Responses
	}
	return nil
}

var File_pkg_apis_proto_sink_v1_udsink_proto protoreflect.FileDescriptor

var file_pkg_apis_proto_sink_v1_udsink_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x64, 0x73, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x38, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xa4, 0x01, 0x0a, 0x05,
	0x44, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x30, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61,
	0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75,
	0x6d, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x22, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73,
	0x67, 0x22, 0x3f, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x32, 0x81, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x6b, 0x46, 0x6e,
	0x12, 0x12, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x75, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x49,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6e, 0x75,
	0x6d, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_proto_sink_v1_udsink_proto_rawDescOnce sync.Once
	file_pkg_apis_proto_sink_v1_udsink_proto_rawDescData = file_pkg_apis_proto_sink_v1_udsink_proto_rawDesc
)

func file_pkg_apis_proto_sink_v1_udsink_proto_rawDescGZIP() []byte {
	file_pkg_apis_proto_sink_v1_udsink_proto_rawDescOnce.Do(func() {
		file_pkg_apis_proto_sink_v1_udsink_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_proto_sink_v1_udsink_proto_rawDescData)
	})
	return file_pkg_apis_proto_sink_v1_udsink_proto_rawDescData
}

var file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_apis_proto_sink_v1_udsink_proto_goTypes = []interface{}{
	(*EventTime)(nil),             // 0: sink.v1.EventTime
	(*Watermark)(nil),             // 1: sink.v1.Watermark
	(*Datum)(nil),                 // 2: sink.v1.Datum
	(*DatumList)(nil),             // 3: sink.v1.DatumList
	(*ReadyResponse)(nil),         // 4: sink.v1.ReadyResponse
	(*Response)(nil),              // 5: sink.v1.Response
	(*ResponseList)(nil),          // 6: sink.v1.ResponseList
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_pkg_apis_proto_sink_v1_udsink_proto_depIdxs = []int32{
	7, // 0: sink.v1.EventTime.event_time:type_name -> google.protobuf.Timestamp
	7, // 1: sink.v1.Watermark.watermark:type_name -> google.protobuf.Timestamp
	0, // 2: sink.v1.Datum.event_time:type_name -> sink.v1.EventTime
	1, // 3: sink.v1.Datum.watermark:type_name -> sink.v1.Watermark
	2, // 4: sink.v1.DatumList.elements:type_name -> sink.v1.Datum
	5, // 5: sink.v1.ResponseList.responses:type_name -> sink.v1.Response
	3, // 6: sink.v1.UserDefinedSink.SinkFn:input_type -> sink.v1.DatumList
	8, // 7: sink.v1.UserDefinedSink.IsReady:input_type -> google.protobuf.Empty
	6, // 8: sink.v1.UserDefinedSink.SinkFn:output_type -> sink.v1.ResponseList
	4, // 9: sink.v1.UserDefinedSink.IsReady:output_type -> sink.v1.ReadyResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_sink_v1_udsink_proto_init() }
func file_pkg_apis_proto_sink_v1_udsink_proto_init() {
	if File_pkg_apis_proto_sink_v1_udsink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTime); i {
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
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watermark); i {
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
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datum); i {
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
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumList); i {
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
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyResponse); i {
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
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseList); i {
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
			RawDescriptor: file_pkg_apis_proto_sink_v1_udsink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_proto_sink_v1_udsink_proto_goTypes,
		DependencyIndexes: file_pkg_apis_proto_sink_v1_udsink_proto_depIdxs,
		MessageInfos:      file_pkg_apis_proto_sink_v1_udsink_proto_msgTypes,
	}.Build()
	File_pkg_apis_proto_sink_v1_udsink_proto = out.File
	file_pkg_apis_proto_sink_v1_udsink_proto_rawDesc = nil
	file_pkg_apis_proto_sink_v1_udsink_proto_goTypes = nil
	file_pkg_apis_proto_sink_v1_udsink_proto_depIdxs = nil
}
