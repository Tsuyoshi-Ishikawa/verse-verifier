// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: message.proto

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

type PubSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*PubSub_Misc
	//	*PubSub_OptimismSignatureExchange
	Body isPubSub_Body `protobuf_oneof:"body"`
}

func (x *PubSub) Reset() {
	*x = PubSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubSub) ProtoMessage() {}

func (x *PubSub) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubSub.ProtoReflect.Descriptor instead.
func (*PubSub) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (m *PubSub) GetBody() isPubSub_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *PubSub) GetMisc() []byte {
	if x, ok := x.GetBody().(*PubSub_Misc); ok {
		return x.Misc
	}
	return nil
}

func (x *PubSub) GetOptimismSignatureExchange() *OptimismSignatureExchange {
	if x, ok := x.GetBody().(*PubSub_OptimismSignatureExchange); ok {
		return x.OptimismSignatureExchange
	}
	return nil
}

type isPubSub_Body interface {
	isPubSub_Body()
}

type PubSub_Misc struct {
	Misc []byte `protobuf:"bytes,1,opt,name=misc,proto3,oneof"`
}

type PubSub_OptimismSignatureExchange struct {
	OptimismSignatureExchange *OptimismSignatureExchange `protobuf:"bytes,2,opt,name=optimism_signature_exchange,json=optimismSignatureExchange,proto3,oneof"`
}

func (*PubSub_Misc) isPubSub_Body() {}

func (*PubSub_OptimismSignatureExchange) isPubSub_Body() {}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*Stream_Misc
	//	*Stream_Eom
	//	*Stream_OptimismSignatureExchange
	Body isStream_Body `protobuf_oneof:"body"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (m *Stream) GetBody() isStream_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Stream) GetMisc() []byte {
	if x, ok := x.GetBody().(*Stream_Misc); ok {
		return x.Misc
	}
	return nil
}

func (x *Stream) GetEom() []byte {
	if x, ok := x.GetBody().(*Stream_Eom); ok {
		return x.Eom
	}
	return nil
}

func (x *Stream) GetOptimismSignatureExchange() *OptimismSignatureExchange {
	if x, ok := x.GetBody().(*Stream_OptimismSignatureExchange); ok {
		return x.OptimismSignatureExchange
	}
	return nil
}

type isStream_Body interface {
	isStream_Body()
}

type Stream_Misc struct {
	Misc []byte `protobuf:"bytes,1,opt,name=misc,proto3,oneof"`
}

type Stream_Eom struct {
	Eom []byte `protobuf:"bytes,2,opt,name=eom,proto3,oneof"`
}

type Stream_OptimismSignatureExchange struct {
	OptimismSignatureExchange *OptimismSignatureExchange `protobuf:"bytes,3,opt,name=optimism_signature_exchange,json=optimismSignatureExchange,proto3,oneof"`
}

func (*Stream_Misc) isStream_Body() {}

func (*Stream_Eom) isStream_Body() {}

func (*Stream_OptimismSignatureExchange) isStream_Body() {}

type OptimismSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PreviousId        string `protobuf:"bytes,2,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	Signer            []byte `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	Scc               []byte `protobuf:"bytes,4,opt,name=scc,proto3" json:"scc,omitempty"`
	BatchIndex        uint64 `protobuf:"varint,5,opt,name=batch_index,json=batchIndex,proto3" json:"batch_index,omitempty"`
	BatchRoot         []byte `protobuf:"bytes,6,opt,name=batch_root,json=batchRoot,proto3" json:"batch_root,omitempty"`
	BatchSize         uint64 `protobuf:"varint,7,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	PrevTotalElements uint64 `protobuf:"varint,8,opt,name=prev_total_elements,json=prevTotalElements,proto3" json:"prev_total_elements,omitempty"`
	ExtraData         []byte `protobuf:"bytes,9,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	Approved          bool   `protobuf:"varint,10,opt,name=approved,proto3" json:"approved,omitempty"`
	Signature         []byte `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *OptimismSignature) Reset() {
	*x = OptimismSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimismSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimismSignature) ProtoMessage() {}

func (x *OptimismSignature) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimismSignature.ProtoReflect.Descriptor instead.
func (*OptimismSignature) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *OptimismSignature) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OptimismSignature) GetPreviousId() string {
	if x != nil {
		return x.PreviousId
	}
	return ""
}

func (x *OptimismSignature) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *OptimismSignature) GetScc() []byte {
	if x != nil {
		return x.Scc
	}
	return nil
}

func (x *OptimismSignature) GetBatchIndex() uint64 {
	if x != nil {
		return x.BatchIndex
	}
	return 0
}

func (x *OptimismSignature) GetBatchRoot() []byte {
	if x != nil {
		return x.BatchRoot
	}
	return nil
}

func (x *OptimismSignature) GetBatchSize() uint64 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *OptimismSignature) GetPrevTotalElements() uint64 {
	if x != nil {
		return x.PrevTotalElements
	}
	return 0
}

func (x *OptimismSignature) GetExtraData() []byte {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

func (x *OptimismSignature) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

func (x *OptimismSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type OptimismSignatureExchange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latests   []*OptimismSignature                 `protobuf:"bytes,1,rep,name=latests,proto3" json:"latests,omitempty"`
	Requests  []*OptimismSignatureExchange_Request `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
	Responses []*OptimismSignature                 `protobuf:"bytes,3,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *OptimismSignatureExchange) Reset() {
	*x = OptimismSignatureExchange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimismSignatureExchange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimismSignatureExchange) ProtoMessage() {}

func (x *OptimismSignatureExchange) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimismSignatureExchange.ProtoReflect.Descriptor instead.
func (*OptimismSignatureExchange) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *OptimismSignatureExchange) GetLatests() []*OptimismSignature {
	if x != nil {
		return x.Latests
	}
	return nil
}

func (x *OptimismSignatureExchange) GetRequests() []*OptimismSignatureExchange_Request {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *OptimismSignatureExchange) GetResponses() []*OptimismSignature {
	if x != nil {
		return x.Responses
	}
	return nil
}

type OptimismSignatureExchange_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer  []byte `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	IdAfter string `protobuf:"bytes,2,opt,name=id_after,json=idAfter,proto3" json:"id_after,omitempty"`
}

func (x *OptimismSignatureExchange_Request) Reset() {
	*x = OptimismSignatureExchange_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimismSignatureExchange_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimismSignatureExchange_Request) ProtoMessage() {}

func (x *OptimismSignatureExchange_Request) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimismSignatureExchange_Request.ProtoReflect.Descriptor instead.
func (*OptimismSignatureExchange_Request) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3, 0}
}

func (x *OptimismSignatureExchange_Request) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *OptimismSignatureExchange_Request) GetIdAfter() string {
	if x != nil {
		return x.IdAfter
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x06, 0x50, 0x75, 0x62,
	0x53, 0x75, 0x62, 0x12, 0x14, 0x0a, 0x04, 0x6d, 0x69, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x69, 0x73, 0x63, 0x12, 0x64, 0x0a, 0x1b, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x73, 0x6d, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73,
	0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x6d, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x14, 0x0a, 0x04, 0x6d, 0x69, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x04, 0x6d, 0x69, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x03, 0x65, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x65, 0x6f, 0x6d, 0x12, 0x64, 0x0a, 0x1b,
	0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x6d, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x73, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73,
	0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xd6, 0x02, 0x0a, 0x11, 0x4f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x63, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x63, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72,
	0x65, 0x76, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x19, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x6d,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x73, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x6d, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x38, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x73, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(*PubSub)(nil),                            // 0: message.PubSub
	(*Stream)(nil),                            // 1: message.Stream
	(*OptimismSignature)(nil),                 // 2: message.OptimismSignature
	(*OptimismSignatureExchange)(nil),         // 3: message.OptimismSignatureExchange
	(*OptimismSignatureExchange_Request)(nil), // 4: message.OptimismSignatureExchange.Request
}
var file_message_proto_depIdxs = []int32{
	3, // 0: message.PubSub.optimism_signature_exchange:type_name -> message.OptimismSignatureExchange
	3, // 1: message.Stream.optimism_signature_exchange:type_name -> message.OptimismSignatureExchange
	2, // 2: message.OptimismSignatureExchange.latests:type_name -> message.OptimismSignature
	4, // 3: message.OptimismSignatureExchange.requests:type_name -> message.OptimismSignatureExchange.Request
	2, // 4: message.OptimismSignatureExchange.responses:type_name -> message.OptimismSignature
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubSub); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimismSignature); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimismSignatureExchange); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimismSignatureExchange_Request); i {
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
	file_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PubSub_Misc)(nil),
		(*PubSub_OptimismSignatureExchange)(nil),
	}
	file_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Stream_Misc)(nil),
		(*Stream_Eom)(nil),
		(*Stream_OptimismSignatureExchange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
