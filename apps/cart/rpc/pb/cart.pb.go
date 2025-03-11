// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: cart.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Userid        int64                  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`     // 用户 ID
	Proid         int64                  `protobuf:"varint,2,opt,name=proid,proto3" json:"proid,omitempty"`       // 商品 ID
	Quantity      int64                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"` // 商品数量
	Checked       int64                  `protobuf:"varint,4,opt,name=checked,proto3" json:"checked,omitempty"`   // 是否勾选
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	mi := &file_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *AddReq) GetProid() int64 {
	if x != nil {
		return x.Proid
	}
	return 0
}

func (x *AddReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddReq) GetChecked() int64 {
	if x != nil {
		return x.Checked
	}
	return 0
}

type AddResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddResp) Reset() {
	*x = AddResp{}
	mi := &file_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResp) ProtoMessage() {}

func (x *AddResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResp.ProtoReflect.Descriptor instead.
func (*AddResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 购物车记录 ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveReq) Reset() {
	*x = RemoveReq{}
	mi := &file_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReq) ProtoMessage() {}

func (x *RemoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReq.ProtoReflect.Descriptor instead.
func (*RemoveReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveResp) Reset() {
	*x = RemoveResp{}
	mi := &file_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResp) ProtoMessage() {}

func (x *RemoveResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResp.ProtoReflect.Descriptor instead.
func (*RemoveResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RemoveResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Userid        int64                  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"` // 用户 ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	mi := &file_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *ListReq) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type ListResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	Data          []*CartItem            `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`       // 购物车列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	mi := &file_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *ListResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListResp) GetData() []*CartItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`             // 购物车记录 ID
	Quantity      int64                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"` // 新的商品数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	mi := &file_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	mi := &file_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  // 购物车记录 ID
	Userid        int64                  `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`                          // 用户 ID
	Proid         int64                  `protobuf:"varint,3,opt,name=proid,proto3" json:"proid,omitempty"`                            // 商品 ID
	Quantity      int64                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`                      // 商品数量
	Checked       int64                  `protobuf:"varint,5,opt,name=checked,proto3" json:"checked,omitempty"`                        // 是否勾选
	CreateTime    string                 `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间
	UpdateTime    string                 `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"` // 更新时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{8}
}

func (x *CartItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartItem) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *CartItem) GetProid() int64 {
	if x != nil {
		return x.Proid
	}
	return 0
}

func (x *CartItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetChecked() int64 {
	if x != nil {
		return x.Checked
	}
	return 0
}

func (x *CartItem) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *CartItem) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x6c, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x37,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x21, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x37, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x9b, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x1e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x27, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData []byte
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cart_proto_rawDesc), len(file_cart_proto_rawDesc)))
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cart_proto_goTypes = []any{
	(*AddReq)(nil),     // 0: pb.AddReq
	(*AddResp)(nil),    // 1: pb.AddResp
	(*RemoveReq)(nil),  // 2: pb.RemoveReq
	(*RemoveResp)(nil), // 3: pb.RemoveResp
	(*ListReq)(nil),    // 4: pb.ListReq
	(*ListResp)(nil),   // 5: pb.ListResp
	(*UpdateReq)(nil),  // 6: pb.UpdateReq
	(*UpdateResp)(nil), // 7: pb.UpdateResp
	(*CartItem)(nil),   // 8: pb.CartItem
}
var file_cart_proto_depIdxs = []int32{
	8, // 0: pb.ListResp.data:type_name -> pb.CartItem
	0, // 1: pb.Cart.Add:input_type -> pb.AddReq
	2, // 2: pb.Cart.Remove:input_type -> pb.RemoveReq
	4, // 3: pb.Cart.List:input_type -> pb.ListReq
	6, // 4: pb.Cart.Update:input_type -> pb.UpdateReq
	1, // 5: pb.Cart.Add:output_type -> pb.AddResp
	3, // 6: pb.Cart.Remove:output_type -> pb.RemoveResp
	5, // 7: pb.Cart.List:output_type -> pb.ListResp
	7, // 8: pb.Cart.Update:output_type -> pb.UpdateResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cart_proto_rawDesc), len(file_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
