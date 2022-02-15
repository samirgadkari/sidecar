// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: messages/sidecar.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_OK                Status = 0
	Status_ERR_SENDING_MSG   Status = 1
	Status_ERR_RECEIVING_MSG Status = 2
	Status_ERR_SUBSCRIBING   Status = 3
	Status_ERR_PUBLISHING    Status = 4
	Status_ERR_LOGGING       Status = 5
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "ERR_SENDING_MSG",
		2: "ERR_RECEIVING_MSG",
		3: "ERR_SUBSCRIBING",
		4: "ERR_PUBLISHING",
		5: "ERR_LOGGING",
	}
	Status_value = map[string]int32{
		"OK":                0,
		"ERR_SENDING_MSG":   1,
		"ERR_RECEIVING_MSG": 2,
		"ERR_SUBSCRIBING":   3,
		"ERR_PUBLISHING":    4,
		"ERR_LOGGING":       5,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_sidecar_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_messages_sidecar_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{0}
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcServType string `protobuf:"bytes,1,opt,name=srcServType,proto3" json:"srcServType,omitempty"`
	DstServType string `protobuf:"bytes,2,opt,name=dstServType,proto3" json:"dstServType,omitempty"`
	ServId      []byte `protobuf:"bytes,3,opt,name=servId,proto3" json:"servId,omitempty"`
	MsgId       uint32 `protobuf:"varint,4,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetSrcServType() string {
	if x != nil {
		return x.SrcServType
	}
	return ""
}

func (x *Header) GetDstServType() string {
	if x != nil {
		return x.DstServType
	}
	return ""
}

func (x *Header) GetServId() []byte {
	if x != nil {
		return x.ServId
	}
	return nil
}

func (x *Header) GetMsgId() uint32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type ResponseHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeader) ProtoMessage() {}

func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHeader.ProtoReflect.Descriptor instead.
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseHeader) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{2}
}

type RegistrationMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header                  *Header              `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	CircuitFailureThreshold *uint32              `protobuf:"varint,2,opt,name=circuitFailureThreshold,proto3,oneof" json:"circuitFailureThreshold,omitempty"`
	DebounceDelay           *durationpb.Duration `protobuf:"bytes,3,opt,name=debounceDelay,proto3,oneof" json:"debounceDelay,omitempty"`
	RetryNum                *uint32              `protobuf:"varint,4,opt,name=retryNum,proto3,oneof" json:"retryNum,omitempty"`
	RetryDelay              *durationpb.Duration `protobuf:"bytes,5,opt,name=retryDelay,proto3,oneof" json:"retryDelay,omitempty"`
}

func (x *RegistrationMsg) Reset() {
	*x = RegistrationMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationMsg) ProtoMessage() {}

func (x *RegistrationMsg) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationMsg.ProtoReflect.Descriptor instead.
func (*RegistrationMsg) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{3}
}

func (x *RegistrationMsg) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RegistrationMsg) GetCircuitFailureThreshold() uint32 {
	if x != nil && x.CircuitFailureThreshold != nil {
		return *x.CircuitFailureThreshold
	}
	return 0
}

func (x *RegistrationMsg) GetDebounceDelay() *durationpb.Duration {
	if x != nil {
		return x.DebounceDelay
	}
	return nil
}

func (x *RegistrationMsg) GetRetryNum() uint32 {
	if x != nil && x.RetryNum != nil {
		return *x.RetryNum
	}
	return 0
}

func (x *RegistrationMsg) GetRetryDelay() *durationpb.Duration {
	if x != nil {
		return x.RetryDelay
	}
	return nil
}

type PubMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PubMsg) Reset() {
	*x = PubMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubMsg) ProtoMessage() {}

func (x *PubMsg) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubMsg.ProtoReflect.Descriptor instead.
func (*PubMsg) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{4}
}

func (x *PubMsg) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PubMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SubMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SubMsg) Reset() {
	*x = SubMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMsg) ProtoMessage() {}

func (x *SubMsg) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMsg.ProtoReflect.Descriptor instead.
func (*SubMsg) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{5}
}

func (x *SubMsg) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SubMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LogMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LogMsg) Reset() {
	*x = LogMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMsg) ProtoMessage() {}

func (x *LogMsg) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMsg.ProtoReflect.Descriptor instead.
func (*LogMsg) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{6}
}

func (x *LogMsg) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *LogMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *Header         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	RspHeader *ResponseHeader `protobuf:"bytes,2,opt,name=rspHeader,proto3" json:"rspHeader,omitempty"`
	Msg       string          `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{7}
}

func (x *RegistrationResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RegistrationResponse) GetRspHeader() *ResponseHeader {
	if x != nil {
		return x.RspHeader
	}
	return nil
}

func (x *RegistrationResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msg    string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SubResponse) Reset() {
	*x = SubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubResponse) ProtoMessage() {}

func (x *SubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubResponse.ProtoReflect.Descriptor instead.
func (*SubResponse) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{8}
}

func (x *SubResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SubResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RecvResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RecvResponse) Reset() {
	*x = RecvResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecvResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecvResponse) ProtoMessage() {}

func (x *RecvResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecvResponse.ProtoReflect.Descriptor instead.
func (*RecvResponse) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{9}
}

func (x *RecvResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RecvResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msg    string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PubResponse) Reset() {
	*x = PubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubResponse) ProtoMessage() {}

func (x *PubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubResponse.ProtoReflect.Descriptor instead.
func (*PubResponse) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{10}
}

func (x *PubResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PubResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Msg    string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_sidecar_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_sidecar_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_messages_sidecar_proto_rawDescGZIP(), []int{11}
}

func (x *LogResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *LogResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_messages_sidecar_proto protoreflect.FileDescriptor

var file_messages_sidecar_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x72, 0x63, 0x53, 0x65, 0x72, 0x76, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x53, 0x65, 0x72, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x28,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0xeb, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x3d, 0x0a, 0x17, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x17, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x44,
	0x0a, 0x0d, 0x64, 0x65, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4e, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4e,
	0x75, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x88, 0x01, 0x01, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x65, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4e, 0x75, 0x6d,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22,
	0x44, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x44, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x4d, 0x73, 0x67, 0x12,
	0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x44, 0x0a, 0x06, 0x4c,
	0x6f, 0x67, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x09, 0x72, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x51,
	0x0a, 0x0b, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x4c, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x51, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x51, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x76, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x53,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x53,
	0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x42, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x5f,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b,
	0x45, 0x52, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x32, 0x91, 0x02,
	0x0a, 0x07, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67,
	0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x03, 0x53, 0x75, 0x62, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x73, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x52, 0x65, 0x63, 0x76, 0x12, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x03, 0x50, 0x75, 0x62, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x4d, 0x73, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x61, 0x6d, 0x69, 0x72, 0x67, 0x61, 0x64, 0x6b, 0x61, 0x72, 0x69, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_sidecar_proto_rawDescOnce sync.Once
	file_messages_sidecar_proto_rawDescData = file_messages_sidecar_proto_rawDesc
)

func file_messages_sidecar_proto_rawDescGZIP() []byte {
	file_messages_sidecar_proto_rawDescOnce.Do(func() {
		file_messages_sidecar_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_sidecar_proto_rawDescData)
	})
	return file_messages_sidecar_proto_rawDescData
}

var file_messages_sidecar_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_sidecar_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_messages_sidecar_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: messages.Status
	(*Header)(nil),               // 1: messages.Header
	(*ResponseHeader)(nil),       // 2: messages.ResponseHeader
	(*Empty)(nil),                // 3: messages.Empty
	(*RegistrationMsg)(nil),      // 4: messages.RegistrationMsg
	(*PubMsg)(nil),               // 5: messages.PubMsg
	(*SubMsg)(nil),               // 6: messages.SubMsg
	(*LogMsg)(nil),               // 7: messages.LogMsg
	(*RegistrationResponse)(nil), // 8: messages.RegistrationResponse
	(*SubResponse)(nil),          // 9: messages.SubResponse
	(*RecvResponse)(nil),         // 10: messages.RecvResponse
	(*PubResponse)(nil),          // 11: messages.PubResponse
	(*LogResponse)(nil),          // 12: messages.LogResponse
	(*durationpb.Duration)(nil),  // 13: google.protobuf.Duration
}
var file_messages_sidecar_proto_depIdxs = []int32{
	1,  // 0: messages.RegistrationMsg.header:type_name -> messages.Header
	13, // 1: messages.RegistrationMsg.debounceDelay:type_name -> google.protobuf.Duration
	13, // 2: messages.RegistrationMsg.retryDelay:type_name -> google.protobuf.Duration
	1,  // 3: messages.PubMsg.header:type_name -> messages.Header
	1,  // 4: messages.SubMsg.header:type_name -> messages.Header
	1,  // 5: messages.LogMsg.header:type_name -> messages.Header
	1,  // 6: messages.RegistrationResponse.header:type_name -> messages.Header
	2,  // 7: messages.RegistrationResponse.rspHeader:type_name -> messages.ResponseHeader
	2,  // 8: messages.SubResponse.header:type_name -> messages.ResponseHeader
	1,  // 9: messages.RecvResponse.header:type_name -> messages.Header
	2,  // 10: messages.PubResponse.header:type_name -> messages.ResponseHeader
	2,  // 11: messages.LogResponse.header:type_name -> messages.ResponseHeader
	4,  // 12: messages.Sidecar.Register:input_type -> messages.RegistrationMsg
	6,  // 13: messages.Sidecar.Sub:input_type -> messages.SubMsg
	3,  // 14: messages.Sidecar.Recv:input_type -> messages.Empty
	5,  // 15: messages.Sidecar.Pub:input_type -> messages.PubMsg
	7,  // 16: messages.Sidecar.Log:input_type -> messages.LogMsg
	8,  // 17: messages.Sidecar.Register:output_type -> messages.RegistrationResponse
	9,  // 18: messages.Sidecar.Sub:output_type -> messages.SubResponse
	10, // 19: messages.Sidecar.Recv:output_type -> messages.RecvResponse
	11, // 20: messages.Sidecar.Pub:output_type -> messages.PubResponse
	12, // 21: messages.Sidecar.Log:output_type -> messages.LogResponse
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_messages_sidecar_proto_init() }
func file_messages_sidecar_proto_init() {
	if File_messages_sidecar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_sidecar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_messages_sidecar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseHeader); i {
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
		file_messages_sidecar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_messages_sidecar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationMsg); i {
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
		file_messages_sidecar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubMsg); i {
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
		file_messages_sidecar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMsg); i {
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
		file_messages_sidecar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMsg); i {
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
		file_messages_sidecar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationResponse); i {
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
		file_messages_sidecar_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubResponse); i {
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
		file_messages_sidecar_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecvResponse); i {
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
		file_messages_sidecar_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubResponse); i {
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
		file_messages_sidecar_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
	file_messages_sidecar_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_sidecar_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messages_sidecar_proto_goTypes,
		DependencyIndexes: file_messages_sidecar_proto_depIdxs,
		EnumInfos:         file_messages_sidecar_proto_enumTypes,
		MessageInfos:      file_messages_sidecar_proto_msgTypes,
	}.Build()
	File_messages_sidecar_proto = out.File
	file_messages_sidecar_proto_rawDesc = nil
	file_messages_sidecar_proto_goTypes = nil
	file_messages_sidecar_proto_depIdxs = nil
}
