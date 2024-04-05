// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: expression_solver.proto

package expression_solver_v1

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

type ExpressionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expression string `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ExpressionRequest) Reset() {
	*x = ExpressionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expression_solver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionRequest) ProtoMessage() {}

func (x *ExpressionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expression_solver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionRequest.ProtoReflect.Descriptor instead.
func (*ExpressionRequest) Descriptor() ([]byte, []int) {
	return file_expression_solver_proto_rawDescGZIP(), []int{0}
}

func (x *ExpressionRequest) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *ExpressionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ExpressionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpressionId string `protobuf:"bytes,1,opt,name=expressionId,proto3" json:"expressionId,omitempty"`
}

func (x *ExpressionResponse) Reset() {
	*x = ExpressionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expression_solver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionResponse) ProtoMessage() {}

func (x *ExpressionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expression_solver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionResponse.ProtoReflect.Descriptor instead.
func (*ExpressionResponse) Descriptor() ([]byte, []int) {
	return file_expression_solver_proto_rawDescGZIP(), []int{1}
}

func (x *ExpressionResponse) GetExpressionId() string {
	if x != nil {
		return x.ExpressionId
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expression_solver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_expression_solver_proto_msgTypes[2]
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
	return file_expression_solver_proto_rawDescGZIP(), []int{2}
}

type CalculatorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculators []*Calculator `protobuf:"bytes,1,rep,name=calculators,proto3" json:"calculators,omitempty"`
}

func (x *CalculatorList) Reset() {
	*x = CalculatorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expression_solver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorList) ProtoMessage() {}

func (x *CalculatorList) ProtoReflect() protoreflect.Message {
	mi := &file_expression_solver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorList.ProtoReflect.Descriptor instead.
func (*CalculatorList) Descriptor() ([]byte, []int) {
	return file_expression_solver_proto_rawDescGZIP(), []int{3}
}

func (x *CalculatorList) GetCalculators() []*Calculator {
	if x != nil {
		return x.Calculators
	}
	return nil
}

type Calculator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId      string           `protobuf:"bytes,1,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	LeastExpression *LeastExpression `protobuf:"bytes,2,opt,name=leastExpression,proto3" json:"leastExpression,omitempty"`
}

func (x *Calculator) Reset() {
	*x = Calculator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expression_solver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calculator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calculator) ProtoMessage() {}

func (x *Calculator) ProtoReflect() protoreflect.Message {
	mi := &file_expression_solver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calculator.ProtoReflect.Descriptor instead.
func (*Calculator) Descriptor() ([]byte, []int) {
	return file_expression_solver_proto_rawDescGZIP(), []int{4}
}

func (x *Calculator) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Calculator) GetLeastExpression() *LeastExpression {
	if x != nil {
		return x.LeastExpression
	}
	return nil
}

type LeastExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number1          float64 `protobuf:"fixed64,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2          float64 `protobuf:"fixed64,2,opt,name=number2,proto3" json:"number2,omitempty"`
	Operator         string  `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	IdExpression     int32   `protobuf:"varint,4,opt,name=idExpression,proto3" json:"idExpression,omitempty"`
	DurationInSecond int32   `protobuf:"varint,5,opt,name=durationInSecond,proto3" json:"durationInSecond,omitempty"`
}

func (x *LeastExpression) Reset() {
	*x = LeastExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expression_solver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeastExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeastExpression) ProtoMessage() {}

func (x *LeastExpression) ProtoReflect() protoreflect.Message {
	mi := &file_expression_solver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeastExpression.ProtoReflect.Descriptor instead.
func (*LeastExpression) Descriptor() ([]byte, []int) {
	return file_expression_solver_proto_rawDescGZIP(), []int{5}
}

func (x *LeastExpression) GetNumber1() float64 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *LeastExpression) GetNumber2() float64 {
	if x != nil {
		return x.Number2
	}
	return 0
}

func (x *LeastExpression) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *LeastExpression) GetIdExpression() int32 {
	if x != nil {
		return x.IdExpression
	}
	return 0
}

func (x *LeastExpression) GetDurationInSecond() int32 {
	if x != nil {
		return x.DurationInSecond
	}
	return 0
}

var File_expression_solver_proto protoreflect.FileDescriptor

var file_expression_solver_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x11,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x51, 0x0a, 0x0e,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f,
	0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22,
	0x7a, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4c, 0x0a,
	0x0f, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb1, 0x01, 0x0a, 0x0f,
	0x6c, 0x65, 0x61, 0x73, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x64, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x64, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x32,
	0xc7, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x0f, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_expression_solver_proto_rawDescOnce sync.Once
	file_expression_solver_proto_rawDescData = file_expression_solver_proto_rawDesc
)

func file_expression_solver_proto_rawDescGZIP() []byte {
	file_expression_solver_proto_rawDescOnce.Do(func() {
		file_expression_solver_proto_rawDescData = protoimpl.X.CompressGZIP(file_expression_solver_proto_rawDescData)
	})
	return file_expression_solver_proto_rawDescData
}

var file_expression_solver_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_expression_solver_proto_goTypes = []interface{}{
	(*ExpressionRequest)(nil),  // 0: expression_solver.ExpressionRequest
	(*ExpressionResponse)(nil), // 1: expression_solver.ExpressionResponse
	(*Empty)(nil),              // 2: expression_solver.Empty
	(*CalculatorList)(nil),     // 3: expression_solver.CalculatorList
	(*Calculator)(nil),         // 4: expression_solver.Calculator
	(*LeastExpression)(nil),    // 5: expression_solver.leastExpression
}
var file_expression_solver_proto_depIdxs = []int32{
	4, // 0: expression_solver.CalculatorList.calculators:type_name -> expression_solver.Calculator
	5, // 1: expression_solver.Calculator.leastExpression:type_name -> expression_solver.leastExpression
	0, // 2: expression_solver.ExpressionSolver.SolveExpression:input_type -> expression_solver.ExpressionRequest
	2, // 3: expression_solver.ExpressionSolver.GetCalculatorsStatus:input_type -> expression_solver.Empty
	1, // 4: expression_solver.ExpressionSolver.SolveExpression:output_type -> expression_solver.ExpressionResponse
	3, // 5: expression_solver.ExpressionSolver.GetCalculatorsStatus:output_type -> expression_solver.CalculatorList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_expression_solver_proto_init() }
func file_expression_solver_proto_init() {
	if File_expression_solver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_expression_solver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionRequest); i {
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
		file_expression_solver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionResponse); i {
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
		file_expression_solver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_expression_solver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorList); i {
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
		file_expression_solver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calculator); i {
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
		file_expression_solver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeastExpression); i {
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
			RawDescriptor: file_expression_solver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_expression_solver_proto_goTypes,
		DependencyIndexes: file_expression_solver_proto_depIdxs,
		MessageInfos:      file_expression_solver_proto_msgTypes,
	}.Build()
	File_expression_solver_proto = out.File
	file_expression_solver_proto_rawDesc = nil
	file_expression_solver_proto_goTypes = nil
	file_expression_solver_proto_depIdxs = nil
}
