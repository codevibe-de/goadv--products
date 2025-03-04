// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.3
// source: product.proto

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

type ProductCategory int32

const (
	ProductCategory_PRODUCT_CATEGORY_UNSPECIFIED ProductCategory = 0
	ProductCategory_PRODUCT_CATEGORY_PIZZA       ProductCategory = 1
	ProductCategory_PRODUCT_CATEGORY_PASTA       ProductCategory = 2
	ProductCategory_PRODUCT_CATEGORY_SALAD       ProductCategory = 3
	ProductCategory_PRODUCT_CATEGORY_DESSERT     ProductCategory = 4
)

// Enum value maps for ProductCategory.
var (
	ProductCategory_name = map[int32]string{
		0: "PRODUCT_CATEGORY_UNSPECIFIED",
		1: "PRODUCT_CATEGORY_PIZZA",
		2: "PRODUCT_CATEGORY_PASTA",
		3: "PRODUCT_CATEGORY_SALAD",
		4: "PRODUCT_CATEGORY_DESSERT",
	}
	ProductCategory_value = map[string]int32{
		"PRODUCT_CATEGORY_UNSPECIFIED": 0,
		"PRODUCT_CATEGORY_PIZZA":       1,
		"PRODUCT_CATEGORY_PASTA":       2,
		"PRODUCT_CATEGORY_SALAD":       3,
		"PRODUCT_CATEGORY_DESSERT":     4,
	}
)

func (x ProductCategory) Enum() *ProductCategory {
	p := new(ProductCategory)
	*p = x
	return p
}

func (x ProductCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_product_proto_enumTypes[0].Descriptor()
}

func (ProductCategory) Type() protoreflect.EnumType {
	return &file_product_proto_enumTypes[0]
}

func (x ProductCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductCategory.Descriptor instead.
func (ProductCategory) EnumDescriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

// Request a product by its unique ID
type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type ProductCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProductCreateRequest) Reset() {
	*x = ProductCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreateRequest) ProtoMessage() {}

func (x *ProductCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreateRequest.ProtoReflect.Descriptor instead.
func (*ProductCreateRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductCreateRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCreateRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// Request multiple products by their IDs
type ProductListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []string `protobuf:"bytes,1,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
}

func (x *ProductListRequest) Reset() {
	*x = ProductListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListRequest) ProtoMessage() {}

func (x *ProductListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListRequest.ProtoReflect.Descriptor instead.
func (*ProductListRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductListRequest) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

// Response containing product details
type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string          `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float64         `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Category  ProductCategory `protobuf:"varint,5,opt,name=category,proto3,enum=product.ProductCategory" json:"category,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetCategory() ProductCategory {
	if x != nil {
		return x.Category
	}
	return ProductCategory_PRODUCT_CATEGORY_UNSPECIFIED
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x2f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x14, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x73, 0x22, 0x90, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2a, 0xa5, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50,
	0x49, 0x5a, 0x5a, 0x41, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x53, 0x54, 0x41,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x41, 0x4c, 0x41, 0x44, 0x10, 0x03, 0x12, 0x1c,
	0x0a, 0x18, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x53, 0x45, 0x52, 0x54, 0x10, 0x04, 0x32, 0xe4, 0x01, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_product_proto_goTypes = []interface{}{
	(ProductCategory)(0),         // 0: product.ProductCategory
	(*ProductRequest)(nil),       // 1: product.ProductRequest
	(*ProductCreateRequest)(nil), // 2: product.ProductCreateRequest
	(*ProductListRequest)(nil),   // 3: product.ProductListRequest
	(*ProductResponse)(nil),      // 4: product.ProductResponse
}
var file_product_proto_depIdxs = []int32{
	0, // 0: product.ProductResponse.category:type_name -> product.ProductCategory
	1, // 1: product.ProductService.GetProduct:input_type -> product.ProductRequest
	3, // 2: product.ProductService.ListProducts:input_type -> product.ProductListRequest
	2, // 3: product.ProductService.CreateProduct:input_type -> product.ProductCreateRequest
	4, // 4: product.ProductService.GetProduct:output_type -> product.ProductResponse
	4, // 5: product.ProductService.ListProducts:output_type -> product.ProductResponse
	4, // 6: product.ProductService.CreateProduct:output_type -> product.ProductResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreateRequest); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		EnumInfos:         file_product_proto_enumTypes,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
