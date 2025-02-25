// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type PostProductRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostProductRequest) Reset()         { *m = PostProductRequest{} }
func (m *PostProductRequest) String() string { return proto.CompactTextString(m) }
func (*PostProductRequest) ProtoMessage()    {}
func (*PostProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1}
}

func (m *PostProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostProductRequest.Unmarshal(m, b)
}
func (m *PostProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostProductRequest.Marshal(b, m, deterministic)
}
func (m *PostProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostProductRequest.Merge(m, src)
}
func (m *PostProductRequest) XXX_Size() int {
	return xxx_messageInfo_PostProductRequest.Size(m)
}
func (m *PostProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostProductRequest proto.InternalMessageInfo

func (m *PostProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PostProductRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PostProductRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type PostProductResponse struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostProductResponse) Reset()         { *m = PostProductResponse{} }
func (m *PostProductResponse) String() string { return proto.CompactTextString(m) }
func (*PostProductResponse) ProtoMessage()    {}
func (*PostProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{2}
}

func (m *PostProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostProductResponse.Unmarshal(m, b)
}
func (m *PostProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostProductResponse.Marshal(b, m, deterministic)
}
func (m *PostProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostProductResponse.Merge(m, src)
}
func (m *PostProductResponse) XXX_Size() int {
	return xxx_messageInfo_PostProductResponse.Size(m)
}
func (m *PostProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostProductResponse proto.InternalMessageInfo

func (m *PostProductResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type GetProductRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{3}
}

func (m *GetProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRequest.Unmarshal(m, b)
}
func (m *GetProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRequest.Marshal(b, m, deterministic)
}
func (m *GetProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRequest.Merge(m, src)
}
func (m *GetProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRequest.Size(m)
}
func (m *GetProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRequest proto.InternalMessageInfo

func (m *GetProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetProductResponse struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductResponse) Reset()         { *m = GetProductResponse{} }
func (m *GetProductResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductResponse) ProtoMessage()    {}
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{4}
}

func (m *GetProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductResponse.Unmarshal(m, b)
}
func (m *GetProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductResponse.Marshal(b, m, deterministic)
}
func (m *GetProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductResponse.Merge(m, src)
}
func (m *GetProductResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductResponse.Size(m)
}
func (m *GetProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductResponse proto.InternalMessageInfo

func (m *GetProductResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type GetProductsRequest struct {
	Skip                 uint64   `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take                 uint64   `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	Ids                  []string `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	Query                string   `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductsRequest) Reset()         { *m = GetProductsRequest{} }
func (m *GetProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductsRequest) ProtoMessage()    {}
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{5}
}

func (m *GetProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsRequest.Unmarshal(m, b)
}
func (m *GetProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsRequest.Marshal(b, m, deterministic)
}
func (m *GetProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsRequest.Merge(m, src)
}
func (m *GetProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductsRequest.Size(m)
}
func (m *GetProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsRequest proto.InternalMessageInfo

func (m *GetProductsRequest) GetSkip() uint64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *GetProductsRequest) GetTake() uint64 {
	if m != nil {
		return m.Take
	}
	return 0
}

func (m *GetProductsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GetProductsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type GetProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProductsResponse) Reset()         { *m = GetProductsResponse{} }
func (m *GetProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductsResponse) ProtoMessage()    {}
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{6}
}

func (m *GetProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsResponse.Unmarshal(m, b)
}
func (m *GetProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsResponse.Marshal(b, m, deterministic)
}
func (m *GetProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsResponse.Merge(m, src)
}
func (m *GetProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductsResponse.Size(m)
}
func (m *GetProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsResponse proto.InternalMessageInfo

func (m *GetProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "pb.Product")
	proto.RegisterType((*PostProductRequest)(nil), "pb.PostProductRequest")
	proto.RegisterType((*PostProductResponse)(nil), "pb.PostProductResponse")
	proto.RegisterType((*GetProductRequest)(nil), "pb.GetProductRequest")
	proto.RegisterType((*GetProductResponse)(nil), "pb.GetProductResponse")
	proto.RegisterType((*GetProductsRequest)(nil), "pb.GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "pb.GetProductsResponse")
}

func init() { proto.RegisterFile("catalog.proto", fileDescriptor_0abbfcf058acdf89) }

var fileDescriptor_0abbfcf058acdf89 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0x32, 0x31,
	0x10, 0xc6, 0xe9, 0xee, 0xbe, 0x2f, 0x32, 0x1b, 0x89, 0x0e, 0x8a, 0x1b, 0x4e, 0x9b, 0x1a, 0x23,
	0x27, 0x0e, 0x78, 0xf4, 0x4f, 0x48, 0x3c, 0x78, 0x25, 0xf5, 0x0b, 0xb8, 0xec, 0x36, 0xa6, 0x41,
	0x69, 0x69, 0x8b, 0x89, 0x9f, 0xd3, 0x2f, 0x64, 0xb6, 0x05, 0x2c, 0x8b, 0x9a, 0x78, 0x9b, 0x3e,
	0x33, 0x93, 0xdf, 0xd3, 0xa7, 0x85, 0xc3, 0xb2, 0xb0, 0xc5, 0x8b, 0x7c, 0x1e, 0x29, 0x2d, 0xad,
	0xc4, 0x48, 0xcd, 0x28, 0x87, 0xf6, 0x54, 0xcb, 0x6a, 0x55, 0x5a, 0xec, 0x42, 0x24, 0xaa, 0x8c,
	0xe4, 0x64, 0xd8, 0x61, 0x91, 0xa8, 0x10, 0x21, 0x59, 0x14, 0xaf, 0x3c, 0x8b, 0x9c, 0xe2, 0x6a,
	0xcc, 0x21, 0xad, 0xb8, 0x29, 0xb5, 0x50, 0x56, 0xc8, 0x45, 0x16, 0xbb, 0x56, 0x28, 0xe1, 0x09,
	0xfc, 0x53, 0x5a, 0x94, 0x3c, 0x4b, 0x72, 0x32, 0x24, 0xcc, 0x1f, 0xe8, 0x13, 0xe0, 0x54, 0x1a,
	0xbb, 0x46, 0x31, 0xbe, 0x5c, 0x71, 0x63, 0xb7, 0x04, 0xf2, 0x33, 0x21, 0xfa, 0x85, 0x10, 0x87,
	0x84, 0x1b, 0xe8, 0xed, 0x10, 0x8c, 0x92, 0x0b, 0xc3, 0xf1, 0x02, 0xda, 0xca, 0x4b, 0x8e, 0x92,
	0x8e, 0xd3, 0x91, 0x9a, 0x8d, 0x36, 0x53, 0x9b, 0x1e, 0x3d, 0x87, 0xe3, 0x07, 0xde, 0xb4, 0xd7,
	0x08, 0x84, 0x5e, 0x03, 0x86, 0x43, 0x7f, 0x23, 0x54, 0xe1, 0xb2, 0x09, 0x12, 0x30, 0x73, 0xa1,
	0xdc, 0x66, 0xc2, 0x5c, 0x5d, 0x6b, 0xb6, 0x98, 0xfb, 0xdc, 0x13, 0xe6, 0x6a, 0x3c, 0x82, 0x58,
	0x54, 0x26, 0x8b, 0xf3, 0x78, 0xd8, 0x61, 0x75, 0x59, 0xa7, 0xb0, 0x5c, 0x71, 0xfd, 0xee, 0x72,
	0xee, 0x30, 0x7f, 0xa0, 0x77, 0xd0, 0xdb, 0xa1, 0xac, 0x3d, 0x5e, 0xc2, 0xc1, 0xda, 0x87, 0xc9,
	0x48, 0x1e, 0x37, 0x4d, 0x6e, 0x9b, 0xe3, 0x0f, 0x02, 0xdd, 0x7b, 0xff, 0x49, 0x1e, 0xb9, 0x7e,
	0x13, 0x25, 0xc7, 0x09, 0xa4, 0x41, 0xb0, 0xd8, 0x77, 0x8b, 0x7b, 0x6f, 0x39, 0x38, 0xdb, 0xd3,
	0x3d, 0x9b, 0xb6, 0xf0, 0x16, 0xe0, 0xcb, 0x14, 0x9e, 0xd6, 0x83, 0x7b, 0x61, 0x0f, 0xfa, 0x4d,
	0x79, 0xbb, 0x3e, 0x81, 0x34, 0xb8, 0x13, 0x36, 0x06, 0xcd, 0x8e, 0x81, 0x6f, 0x2e, 0x4f, 0x5b,
	0xb3, 0xff, 0xee, 0xbf, 0x5f, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xdd, 0x14, 0xe5, 0x00,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	PostProduct(ctx context.Context, in *PostProductRequest, opts ...grpc.CallOption) (*PostProductResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
}

type catalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatalogServiceClient(cc *grpc.ClientConn) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) PostProduct(ctx context.Context, in *PostProductRequest, opts ...grpc.CallOption) (*PostProductResponse, error) {
	out := new(PostProductResponse)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/PostProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	PostProduct(context.Context, *PostProductRequest) (*PostProductResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) PostProduct(ctx context.Context, req *PostProductRequest) (*PostProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostProduct not implemented")
}
func (*UnimplementedCatalogServiceServer) GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedCatalogServiceServer) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_PostProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).PostProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/PostProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).PostProduct(ctx, req.(*PostProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostProduct",
			Handler:    _CatalogService_PostProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _CatalogService_GetProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _CatalogService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}