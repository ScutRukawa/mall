// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PageInfo struct {
	Pn                   uint32   `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	PSize                uint32   `protobuf:"varint,2,opt,name=pSize,proto3" json:"pSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *PageInfo) GetPSize() uint32 {
	if m != nil {
		return m.PSize
	}
	return 0
}

type MobileRequest struct {
	Moblie               string   `protobuf:"bytes,1,opt,name=moblie,proto3" json:"moblie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileRequest) Reset()         { *m = MobileRequest{} }
func (m *MobileRequest) String() string { return proto.CompactTextString(m) }
func (*MobileRequest) ProtoMessage()    {}
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *MobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileRequest.Unmarshal(m, b)
}
func (m *MobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileRequest.Marshal(b, m, deterministic)
}
func (m *MobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileRequest.Merge(m, src)
}
func (m *MobileRequest) XXX_Size() int {
	return xxx_messageInfo_MobileRequest.Size(m)
}
func (m *MobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MobileRequest proto.InternalMessageInfo

func (m *MobileRequest) GetMoblie() string {
	if m != nil {
		return m.Moblie
	}
	return ""
}

type CreateUserInfo struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInfo) Reset()         { *m = CreateUserInfo{} }
func (m *CreateUserInfo) String() string { return proto.CompactTextString(m) }
func (*CreateUserInfo) ProtoMessage()    {}
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CreateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInfo.Unmarshal(m, b)
}
func (m *CreateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInfo.Marshal(b, m, deterministic)
}
func (m *CreateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInfo.Merge(m, src)
}
func (m *CreateUserInfo) XXX_Size() int {
	return xxx_messageInfo_CreateUserInfo.Size(m)
}
func (m *CreateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInfo proto.InternalMessageInfo

func (m *CreateUserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CreateUserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UpdateUserInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday             string   `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfo) Reset()         { *m = UpdateUserInfo{} }
func (m *UpdateUserInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfo) ProtoMessage()    {}
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UpdateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfo.Unmarshal(m, b)
}
func (m *UpdateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfo.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfo.Merge(m, src)
}
func (m *UpdateUserInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfo.Size(m)
}
func (m *UpdateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfo proto.InternalMessageInfo

func (m *UpdateUserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UpdateUserInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UpdateUserInfo) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

type UserInfoResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Birthday             string   `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender               int32    `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Role                 int32    `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	NickName             string   `protobuf:"bytes,7,opt,name=nickName,proto3" json:"nickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfoResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfoResponse) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *UserInfoResponse) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfoResponse) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *UserInfoResponse) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type IDRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserListResponse struct {
	Data                 []*UserInfoResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total                int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetData() []*UserInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UserListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*PageInfo)(nil), "PageInfo")
	proto.RegisterType((*MobileRequest)(nil), "MobileRequest")
	proto.RegisterType((*CreateUserInfo)(nil), "CreateUserInfo")
	proto.RegisterType((*UpdateUserInfo)(nil), "UpdateUserInfo")
	proto.RegisterType((*UserInfoResponse)(nil), "UserInfoResponse")
	proto.RegisterType((*IDRequest)(nil), "IDRequest")
	proto.RegisterType((*UserListResponse)(nil), "UserListResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x8b, 0x13, 0x31,
	0x14, 0xc5, 0xe9, 0x6c, 0xdb, 0xdd, 0xb9, 0xcb, 0x4e, 0x35, 0x48, 0x19, 0x66, 0x5f, 0x96, 0x01,
	0x71, 0x41, 0xc9, 0x2e, 0xab, 0x3e, 0xf9, 0x56, 0x15, 0x29, 0xf8, 0x8f, 0xc8, 0xbe, 0xf8, 0x64,
	0x6a, 0xee, 0xd6, 0xe0, 0x74, 0x12, 0x93, 0x14, 0x19, 0x9f, 0xfc, 0x5c, 0x7e, 0x3a, 0x99, 0xa4,
	0x99, 0xed, 0xd4, 0x22, 0xec, 0x53, 0x7b, 0x92, 0xdc, 0xfb, 0xbb, 0x87, 0x39, 0x17, 0x60, 0x6d,
	0xd1, 0x50, 0x6d, 0x94, 0x53, 0xc5, 0xe9, 0x52, 0xa9, 0x65, 0x85, 0x17, 0x5e, 0x2d, 0xd6, 0x37,
	0x17, 0xb8, 0xd2, 0xae, 0x09, 0x97, 0xe5, 0x25, 0x1c, 0x7d, 0xe4, 0x4b, 0x9c, 0xd7, 0x37, 0x8a,
	0x64, 0x90, 0xe8, 0x3a, 0x1f, 0x9c, 0x0d, 0xce, 0x4f, 0x58, 0xa2, 0x6b, 0xf2, 0x00, 0x46, 0xfa,
	0x93, 0xfc, 0x85, 0x79, 0xe2, 0x8f, 0x82, 0x28, 0x1f, 0xc1, 0xc9, 0x3b, 0xb5, 0x90, 0x15, 0x32,
	0xfc, 0xb1, 0x46, 0xeb, 0xc8, 0x14, 0xc6, 0x2b, 0xb5, 0xa8, 0x24, 0xfa, 0xd2, 0x94, 0x6d, 0x54,
	0xf9, 0x05, 0xb2, 0x97, 0x06, 0xb9, 0xc3, 0x6b, 0x8b, 0xc6, 0x03, 0x0a, 0x38, 0xaa, 0xe5, 0xd7,
	0xef, 0xef, 0xf9, 0x2a, 0xbe, 0xed, 0x74, 0x7b, 0xa7, 0xb9, 0xb5, 0x3f, 0x95, 0x11, 0x9e, 0x97,
	0xb2, 0x4e, 0x6f, 0x08, 0xb2, 0xc2, 0xfc, 0xa0, 0x23, 0xc8, 0x0a, 0x4b, 0x0d, 0xd9, 0xb5, 0x16,
	0xdb, 0x84, 0x0c, 0x12, 0x29, 0x7c, 0xef, 0x11, 0x4b, 0xa4, 0xe8, 0x11, 0x93, 0x1d, 0xe2, 0x14,
	0xc6, 0x4b, 0xac, 0x05, 0x9a, 0xd8, 0x35, 0xa8, 0xb6, 0x66, 0x21, 0x8d, 0xfb, 0x26, 0x78, 0x93,
	0x0f, 0x43, 0x4d, 0xd4, 0xe5, 0x9f, 0x01, 0xdc, 0x8b, 0x30, 0x86, 0x56, 0xab, 0xda, 0xe2, 0x3e,
	0xe8, 0x5d, 0xad, 0xfc, 0x0f, 0xba, 0x35, 0xe8, 0xc8, 0x33, 0xe2, 0xa0, 0x04, 0x86, 0x46, 0x55,
	0x98, 0x8f, 0xfd, 0xa9, 0xff, 0xdf, 0x33, 0x7c, 0xd8, 0x37, 0x5c, 0x9e, 0x42, 0x3a, 0x7f, 0x15,
	0xbf, 0xda, 0xce, 0xd0, 0xe5, 0x87, 0x60, 0xec, 0xad, 0xb4, 0xae, 0x33, 0xf6, 0x10, 0x86, 0x82,
	0x3b, 0x9e, 0x0f, 0xce, 0x0e, 0xce, 0x8f, 0xaf, 0xee, 0xd3, 0x5d, 0xe7, 0xcc, 0x5f, 0xb7, 0x39,
	0x71, 0xca, 0xf1, 0xca, 0x9b, 0x1d, 0xb1, 0x20, 0xae, 0x7e, 0x27, 0x30, 0x6c, 0x0b, 0xc8, 0x63,
	0x38, 0x7e, 0x83, 0x2e, 0x36, 0x27, 0x29, 0x8d, 0x81, 0x2b, 0x42, 0xc7, 0x1e, 0xf2, 0x19, 0x4c,
	0x36, 0x8f, 0x67, 0x4d, 0x88, 0x19, 0xc9, 0x68, 0x2f, 0x6f, 0xc5, 0xbf, 0x73, 0x90, 0x27, 0x1d,
	0x62, 0xd6, 0xcc, 0x05, 0x01, 0xda, 0xf9, 0xdc, 0xf7, 0xfa, 0x12, 0xe0, 0x36, 0x98, 0x64, 0x42,
	0xfb, 0x29, 0xdd, 0x57, 0xf1, 0x1c, 0xe0, 0x36, 0x68, 0x64, 0x42, 0xfb, 0xa9, 0x2b, 0xa6, 0x34,
	0xac, 0x18, 0x8d, 0x2b, 0x46, 0x5f, 0xb7, 0x2b, 0x36, 0x4b, 0x3f, 0x1f, 0xd2, 0x17, 0xe1, 0x6c,
	0xec, 0x7f, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x39, 0x38, 0x1e, 0x99, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
	GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	GetUserById(context.Context, *IDRequest) (*UserInfoResponse, error)
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*empty.Empty, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserList(ctx context.Context, req *PageInfo) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedUserServer) GetUserByMobile(ctx context.Context, req *MobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (*UnimplementedUserServer) GetUserById(ctx context.Context, req *IDRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateUserInfo) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UpdateUserInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _User_GetUserByMobile_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
