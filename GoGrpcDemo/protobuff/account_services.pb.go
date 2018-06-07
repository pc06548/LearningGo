// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account_services.proto

package protobuff

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_services_73951d16d8e8a80a, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type RequestAccountDetails struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAccountDetails) Reset()         { *m = RequestAccountDetails{} }
func (m *RequestAccountDetails) String() string { return proto.CompactTextString(m) }
func (*RequestAccountDetails) ProtoMessage()    {}
func (*RequestAccountDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_services_73951d16d8e8a80a, []int{1}
}
func (m *RequestAccountDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAccountDetails.Unmarshal(m, b)
}
func (m *RequestAccountDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAccountDetails.Marshal(b, m, deterministic)
}
func (dst *RequestAccountDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAccountDetails.Merge(dst, src)
}
func (m *RequestAccountDetails) XXX_Size() int {
	return xxx_messageInfo_RequestAccountDetails.Size(m)
}
func (m *RequestAccountDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAccountDetails.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAccountDetails proto.InternalMessageInfo

func (m *RequestAccountDetails) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type AccountId struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountId) Reset()         { *m = AccountId{} }
func (m *AccountId) String() string { return proto.CompactTextString(m) }
func (*AccountId) ProtoMessage()    {}
func (*AccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_services_73951d16d8e8a80a, []int{2}
}
func (m *AccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountId.Unmarshal(m, b)
}
func (m *AccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountId.Marshal(b, m, deterministic)
}
func (dst *AccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountId.Merge(dst, src)
}
func (m *AccountId) XXX_Size() int {
	return xxx_messageInfo_AccountId.Size(m)
}
func (m *AccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountId.DiscardUnknown(m)
}

var xxx_messageInfo_AccountId proto.InternalMessageInfo

func (m *AccountId) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "protobuff.Account")
	proto.RegisterType((*RequestAccountDetails)(nil), "protobuff.RequestAccountDetails")
	proto.RegisterType((*AccountId)(nil), "protobuff.AccountId")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServicesClient is the client API for AccountServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServicesClient interface {
	AddMoneyToAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
	GetAccount(ctx context.Context, in *RequestAccountDetails, opts ...grpc.CallOption) (*Account, error)
	CreateAccount(ctx context.Context, in *AccountId, opts ...grpc.CallOption) (*Account, error)
}

type accountServicesClient struct {
	cc *grpc.ClientConn
}

func NewAccountServicesClient(cc *grpc.ClientConn) AccountServicesClient {
	return &accountServicesClient{cc}
}

func (c *accountServicesClient) AddMoneyToAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/protobuff.AccountServices/AddMoneyToAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServicesClient) GetAccount(ctx context.Context, in *RequestAccountDetails, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/protobuff.AccountServices/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServicesClient) CreateAccount(ctx context.Context, in *AccountId, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/protobuff.AccountServices/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServicesServer is the server API for AccountServices service.
type AccountServicesServer interface {
	AddMoneyToAccount(context.Context, *Account) (*Account, error)
	GetAccount(context.Context, *RequestAccountDetails) (*Account, error)
	CreateAccount(context.Context, *AccountId) (*Account, error)
}

func RegisterAccountServicesServer(s *grpc.Server, srv AccountServicesServer) {
	s.RegisterService(&_AccountServices_serviceDesc, srv)
}

func _AccountServices_AddMoneyToAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServicesServer).AddMoneyToAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuff.AccountServices/AddMoneyToAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServicesServer).AddMoneyToAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountServices_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccountDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServicesServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuff.AccountServices/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServicesServer).GetAccount(ctx, req.(*RequestAccountDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountServices_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServicesServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuff.AccountServices/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServicesServer).CreateAccount(ctx, req.(*AccountId))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuff.AccountServices",
	HandlerType: (*AccountServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMoneyToAccount",
			Handler:    _AccountServices_AddMoneyToAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountServices_GetAccount_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _AccountServices_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_services.proto",
}

func init() {
	proto.RegisterFile("account_services.proto", fileDescriptor_account_services_73951d16d8e8a80a)
}

var fileDescriptor_account_services_73951d16d8e8a80a = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0x49, 0xa5, 0x69, 0x69, 0x4a, 0xf6, 0x5c, 0xec, 0x8e, 0x10,
	0x45, 0x42, 0x32, 0x5c, 0x9c, 0x50, 0xf5, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x08, 0x01, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0x5c, 0x10, 0x5b, 0x82, 0x49, 0x81, 0x51, 0x83, 0x29,
	0x08, 0xca, 0x53, 0x32, 0xe5, 0x12, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x81, 0x9a, 0xe3,
	0x92, 0x5a, 0x92, 0x98, 0x99, 0x53, 0x8c, 0xdf, 0x38, 0x25, 0x4d, 0x2e, 0x4e, 0x47, 0xb8, 0xd9,
	0x78, 0x95, 0x1a, 0x5d, 0x66, 0xe4, 0xe2, 0x87, 0xaa, 0x0d, 0x86, 0xfa, 0x43, 0xc8, 0x96, 0x4b,
	0xd0, 0x31, 0x25, 0xc5, 0x37, 0x3f, 0x2f, 0xb5, 0x32, 0x24, 0x1f, 0xe6, 0x01, 0x21, 0x3d, 0xb8,
	0xbf, 0xf4, 0xa0, 0x62, 0x52, 0x58, 0xc4, 0x94, 0x18, 0x84, 0x5c, 0xb8, 0xb8, 0xdc, 0x53, 0x61,
	0x0e, 0x16, 0x52, 0x40, 0x52, 0x83, 0xd5, 0x2f, 0x38, 0x4c, 0xb1, 0xe6, 0xe2, 0x75, 0x2e, 0x4a,
	0x4d, 0x2c, 0x49, 0x85, 0x19, 0x24, 0x82, 0xa9, 0xcc, 0x33, 0x05, 0xbb, 0xe6, 0x24, 0x36, 0xb0,
	0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x79, 0x70, 0xc4, 0x1c, 0xa4, 0x01, 0x00, 0x00,
}
