// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package iproto

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

type CreateRawTransferRequest struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient" json:"recipient,omitempty"`
	Amount               []byte   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Nonce                uint64   `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRawTransferRequest) Reset()         { *m = CreateRawTransferRequest{} }
func (m *CreateRawTransferRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRawTransferRequest) ProtoMessage()    {}
func (*CreateRawTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{0}
}
func (m *CreateRawTransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRawTransferRequest.Unmarshal(m, b)
}
func (m *CreateRawTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRawTransferRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRawTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRawTransferRequest.Merge(dst, src)
}
func (m *CreateRawTransferRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRawTransferRequest.Size(m)
}
func (m *CreateRawTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRawTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRawTransferRequest proto.InternalMessageInfo

func (m *CreateRawTransferRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *CreateRawTransferRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *CreateRawTransferRequest) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CreateRawTransferRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CreateRawTransferRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateRawTransferResponse struct {
	SerializedTransfer   []byte   `protobuf:"bytes,1,opt,name=serialized_transfer,json=serializedTransfer,proto3" json:"serialized_transfer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRawTransferResponse) Reset()         { *m = CreateRawTransferResponse{} }
func (m *CreateRawTransferResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRawTransferResponse) ProtoMessage()    {}
func (*CreateRawTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{1}
}
func (m *CreateRawTransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRawTransferResponse.Unmarshal(m, b)
}
func (m *CreateRawTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRawTransferResponse.Marshal(b, m, deterministic)
}
func (dst *CreateRawTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRawTransferResponse.Merge(dst, src)
}
func (m *CreateRawTransferResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRawTransferResponse.Size(m)
}
func (m *CreateRawTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRawTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRawTransferResponse proto.InternalMessageInfo

func (m *CreateRawTransferResponse) GetSerializedTransfer() []byte {
	if m != nil {
		return m.SerializedTransfer
	}
	return nil
}

type CreateRawVoteRequest struct {
	Voter                []byte   `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Votee                []byte   `protobuf:"bytes,2,opt,name=votee,proto3" json:"votee,omitempty"`
	Nonce                uint64   `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRawVoteRequest) Reset()         { *m = CreateRawVoteRequest{} }
func (m *CreateRawVoteRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRawVoteRequest) ProtoMessage()    {}
func (*CreateRawVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{2}
}
func (m *CreateRawVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRawVoteRequest.Unmarshal(m, b)
}
func (m *CreateRawVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRawVoteRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRawVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRawVoteRequest.Merge(dst, src)
}
func (m *CreateRawVoteRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRawVoteRequest.Size(m)
}
func (m *CreateRawVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRawVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRawVoteRequest proto.InternalMessageInfo

func (m *CreateRawVoteRequest) GetVoter() []byte {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *CreateRawVoteRequest) GetVotee() []byte {
	if m != nil {
		return m.Votee
	}
	return nil
}

func (m *CreateRawVoteRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type CreateRawVoteResponse struct {
	SerializedVote       []byte   `protobuf:"bytes,1,opt,name=serialized_vote,json=serializedVote,proto3" json:"serialized_vote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRawVoteResponse) Reset()         { *m = CreateRawVoteResponse{} }
func (m *CreateRawVoteResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRawVoteResponse) ProtoMessage()    {}
func (*CreateRawVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{3}
}
func (m *CreateRawVoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRawVoteResponse.Unmarshal(m, b)
}
func (m *CreateRawVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRawVoteResponse.Marshal(b, m, deterministic)
}
func (dst *CreateRawVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRawVoteResponse.Merge(dst, src)
}
func (m *CreateRawVoteResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRawVoteResponse.Size(m)
}
func (m *CreateRawVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRawVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRawVoteResponse proto.InternalMessageInfo

func (m *CreateRawVoteResponse) GetSerializedVote() []byte {
	if m != nil {
		return m.SerializedVote
	}
	return nil
}

type SendTransferRequest struct {
	SerializedTransfer   []byte   `protobuf:"bytes,1,opt,name=serialized_transfer,json=serializedTransfer,proto3" json:"serialized_transfer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendTransferRequest) Reset()         { *m = SendTransferRequest{} }
func (m *SendTransferRequest) String() string { return proto.CompactTextString(m) }
func (*SendTransferRequest) ProtoMessage()    {}
func (*SendTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{4}
}
func (m *SendTransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTransferRequest.Unmarshal(m, b)
}
func (m *SendTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTransferRequest.Marshal(b, m, deterministic)
}
func (dst *SendTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTransferRequest.Merge(dst, src)
}
func (m *SendTransferRequest) XXX_Size() int {
	return xxx_messageInfo_SendTransferRequest.Size(m)
}
func (m *SendTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendTransferRequest proto.InternalMessageInfo

func (m *SendTransferRequest) GetSerializedTransfer() []byte {
	if m != nil {
		return m.SerializedTransfer
	}
	return nil
}

type SendTransferResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendTransferResponse) Reset()         { *m = SendTransferResponse{} }
func (m *SendTransferResponse) String() string { return proto.CompactTextString(m) }
func (*SendTransferResponse) ProtoMessage()    {}
func (*SendTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{5}
}
func (m *SendTransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTransferResponse.Unmarshal(m, b)
}
func (m *SendTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTransferResponse.Marshal(b, m, deterministic)
}
func (dst *SendTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTransferResponse.Merge(dst, src)
}
func (m *SendTransferResponse) XXX_Size() int {
	return xxx_messageInfo_SendTransferResponse.Size(m)
}
func (m *SendTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendTransferResponse proto.InternalMessageInfo

type SendVoteRequest struct {
	SerializedVote       []byte   `protobuf:"bytes,1,opt,name=serialized_vote,json=serializedVote,proto3" json:"serialized_vote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVoteRequest) Reset()         { *m = SendVoteRequest{} }
func (m *SendVoteRequest) String() string { return proto.CompactTextString(m) }
func (*SendVoteRequest) ProtoMessage()    {}
func (*SendVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{6}
}
func (m *SendVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVoteRequest.Unmarshal(m, b)
}
func (m *SendVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVoteRequest.Marshal(b, m, deterministic)
}
func (dst *SendVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVoteRequest.Merge(dst, src)
}
func (m *SendVoteRequest) XXX_Size() int {
	return xxx_messageInfo_SendVoteRequest.Size(m)
}
func (m *SendVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendVoteRequest proto.InternalMessageInfo

func (m *SendVoteRequest) GetSerializedVote() []byte {
	if m != nil {
		return m.SerializedVote
	}
	return nil
}

type SendVoteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVoteResponse) Reset()         { *m = SendVoteResponse{} }
func (m *SendVoteResponse) String() string { return proto.CompactTextString(m) }
func (*SendVoteResponse) ProtoMessage()    {}
func (*SendVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_3a4b7e1b850fd7c2, []int{7}
}
func (m *SendVoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVoteResponse.Unmarshal(m, b)
}
func (m *SendVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVoteResponse.Marshal(b, m, deterministic)
}
func (dst *SendVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVoteResponse.Merge(dst, src)
}
func (m *SendVoteResponse) XXX_Size() int {
	return xxx_messageInfo_SendVoteResponse.Size(m)
}
func (m *SendVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendVoteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRawTransferRequest)(nil), "iproto.CreateRawTransferRequest")
	proto.RegisterType((*CreateRawTransferResponse)(nil), "iproto.CreateRawTransferResponse")
	proto.RegisterType((*CreateRawVoteRequest)(nil), "iproto.CreateRawVoteRequest")
	proto.RegisterType((*CreateRawVoteResponse)(nil), "iproto.CreateRawVoteResponse")
	proto.RegisterType((*SendTransferRequest)(nil), "iproto.SendTransferRequest")
	proto.RegisterType((*SendTransferResponse)(nil), "iproto.SendTransferResponse")
	proto.RegisterType((*SendVoteRequest)(nil), "iproto.SendVoteRequest")
	proto.RegisterType((*SendVoteResponse)(nil), "iproto.SendVoteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChainService service

type ChainServiceClient interface {
	CreateRawTransfer(ctx context.Context, in *CreateRawTransferRequest, opts ...grpc.CallOption) (*CreateRawTransferResponse, error)
	CreateRawVote(ctx context.Context, in *CreateRawVoteRequest, opts ...grpc.CallOption) (*CreateRawVoteResponse, error)
	SendTransfer(ctx context.Context, in *SendTransferRequest, opts ...grpc.CallOption) (*SendTransferResponse, error)
	SendVote(ctx context.Context, in *SendVoteRequest, opts ...grpc.CallOption) (*SendVoteResponse, error)
}

type chainServiceClient struct {
	cc *grpc.ClientConn
}

func NewChainServiceClient(cc *grpc.ClientConn) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) CreateRawTransfer(ctx context.Context, in *CreateRawTransferRequest, opts ...grpc.CallOption) (*CreateRawTransferResponse, error) {
	out := new(CreateRawTransferResponse)
	err := grpc.Invoke(ctx, "/iproto.ChainService/CreateRawTransfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) CreateRawVote(ctx context.Context, in *CreateRawVoteRequest, opts ...grpc.CallOption) (*CreateRawVoteResponse, error) {
	out := new(CreateRawVoteResponse)
	err := grpc.Invoke(ctx, "/iproto.ChainService/CreateRawVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) SendTransfer(ctx context.Context, in *SendTransferRequest, opts ...grpc.CallOption) (*SendTransferResponse, error) {
	out := new(SendTransferResponse)
	err := grpc.Invoke(ctx, "/iproto.ChainService/SendTransfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) SendVote(ctx context.Context, in *SendVoteRequest, opts ...grpc.CallOption) (*SendVoteResponse, error) {
	out := new(SendVoteResponse)
	err := grpc.Invoke(ctx, "/iproto.ChainService/SendVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChainService service

type ChainServiceServer interface {
	CreateRawTransfer(context.Context, *CreateRawTransferRequest) (*CreateRawTransferResponse, error)
	CreateRawVote(context.Context, *CreateRawVoteRequest) (*CreateRawVoteResponse, error)
	SendTransfer(context.Context, *SendTransferRequest) (*SendTransferResponse, error)
	SendVote(context.Context, *SendVoteRequest) (*SendVoteResponse, error)
}

func RegisterChainServiceServer(s *grpc.Server, srv ChainServiceServer) {
	s.RegisterService(&_ChainService_serviceDesc, srv)
}

func _ChainService_CreateRawTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRawTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).CreateRawTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.ChainService/CreateRawTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).CreateRawTransfer(ctx, req.(*CreateRawTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_CreateRawVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRawVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).CreateRawVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.ChainService/CreateRawVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).CreateRawVote(ctx, req.(*CreateRawVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_SendTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).SendTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.ChainService/SendTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).SendTransfer(ctx, req.(*SendTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_SendVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).SendVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.ChainService/SendVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).SendVote(ctx, req.(*SendVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iproto.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRawTransfer",
			Handler:    _ChainService_CreateRawTransfer_Handler,
		},
		{
			MethodName: "CreateRawVote",
			Handler:    _ChainService_CreateRawVote_Handler,
		},
		{
			MethodName: "SendTransfer",
			Handler:    _ChainService_SendTransfer_Handler,
		},
		{
			MethodName: "SendVote",
			Handler:    _ChainService_SendVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_3a4b7e1b850fd7c2) }

var fileDescriptor_rpc_3a4b7e1b850fd7c2 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x7d, 0xe9, 0x17, 0xaf, 0x97, 0xbc, 0x57, 0xbd, 0x8d, 0x35, 0xd6, 0x0a, 0x31, 0x1b, 0xbb,
	0xaa, 0xa0, 0x3b, 0x37, 0x0a, 0x05, 0x37, 0x8a, 0x8b, 0x54, 0x44, 0xdc, 0xc8, 0x98, 0x5c, 0x31,
	0xa0, 0x93, 0x38, 0x99, 0x56, 0xf0, 0x67, 0xf8, 0xab, 0xfc, 0x59, 0x92, 0x64, 0x92, 0x4c, 0xbf,
	0x44, 0x57, 0xc9, 0xbd, 0xe7, 0x9e, 0x33, 0xe7, 0xdc, 0x19, 0x68, 0x8b, 0xd8, 0x1f, 0xc5, 0x22,
	0x92, 0x11, 0xb6, 0xc2, 0xec, 0xeb, 0x7e, 0x18, 0x60, 0x8f, 0x05, 0x31, 0x49, 0x1e, 0x7b, 0xbb,
	0x16, 0x8c, 0x27, 0x8f, 0x24, 0x3c, 0x7a, 0x9d, 0x52, 0x22, 0xb1, 0x07, 0xad, 0x84, 0x78, 0x40,
	0xc2, 0x36, 0x1c, 0x63, 0xd8, 0xf6, 0x54, 0x85, 0x03, 0x68, 0x0b, 0xf2, 0xc3, 0x38, 0x24, 0x2e,
	0xed, 0x5a, 0x06, 0x55, 0x8d, 0x94, 0xc5, 0x5e, 0xa2, 0x29, 0x97, 0x76, 0xdd, 0x31, 0x86, 0xa6,
	0xa7, 0x2a, 0xb4, 0xa0, 0xc9, 0x23, 0xee, 0x93, 0xdd, 0x70, 0x8c, 0x61, 0xc3, 0xcb, 0x0b, 0x44,
	0x68, 0x04, 0x4c, 0x32, 0xbb, 0x99, 0xcd, 0x66, 0xff, 0xee, 0x25, 0xec, 0xac, 0xf0, 0x94, 0xc4,
	0x11, 0x4f, 0x08, 0x0f, 0xa1, 0x9b, 0x90, 0x08, 0xd9, 0x73, 0xf8, 0x4e, 0xc1, 0xbd, 0x54, 0x70,
	0xe6, 0xd0, 0xf4, 0xb0, 0x82, 0x0a, 0xa2, 0x7b, 0x0b, 0x56, 0xa9, 0x76, 0x13, 0x49, 0x2a, 0xd2,
	0x59, 0xd0, 0x9c, 0x45, 0xb2, 0xa4, 0xe6, 0x45, 0xd1, 0xa5, 0x2c, 0x97, 0xea, 0x52, 0xe5, 0xbd,
	0xae, 0x79, 0x77, 0xcf, 0x60, 0x6b, 0x41, 0x59, 0x79, 0x3c, 0x80, 0x8e, 0xe6, 0x31, 0x95, 0x50,
	0x87, 0xfc, 0xaf, 0xda, 0x29, 0xc1, 0x3d, 0x87, 0xee, 0x84, 0x78, 0xb0, 0xb8, 0xf8, 0x5f, 0x67,
	0xec, 0x81, 0x35, 0xaf, 0x93, 0x1b, 0x71, 0x4f, 0xa0, 0x93, 0xf6, 0xf5, 0xd8, 0x3f, 0xf6, 0x86,
	0xb0, 0x51, 0x71, 0x73, 0xbd, 0xa3, 0xcf, 0x1a, 0x98, 0xe3, 0x27, 0x16, 0xf2, 0x09, 0x89, 0x59,
	0xe8, 0x13, 0xde, 0xc1, 0xe6, 0xd2, 0x55, 0xa1, 0x33, 0xca, 0x5f, 0xd7, 0x68, 0xdd, 0xcb, 0xea,
	0xef, 0x7f, 0x33, 0xa1, 0xac, 0xff, 0xc1, 0x2b, 0xf8, 0x37, 0xb7, 0x5e, 0x1c, 0x2c, 0xb1, 0xb4,
	0x60, 0xfd, 0xbd, 0x35, 0x68, 0xa9, 0x77, 0x01, 0xa6, 0xbe, 0x24, 0xdc, 0x2d, 0x08, 0x2b, 0xae,
	0xa0, 0x3f, 0x58, 0x0d, 0x96, 0x62, 0xa7, 0xf0, 0xb7, 0xd8, 0x0e, 0x6e, 0xeb, 0xb3, 0xba, 0x25,
	0x7b, 0x19, 0x28, 0x04, 0x1e, 0x5a, 0x19, 0x72, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xf5,
	0x1c, 0x9e, 0x95, 0x03, 0x00, 0x00,
}
