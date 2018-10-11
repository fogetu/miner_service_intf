// Code generated by protoc-gen-go. DO NOT EDIT.
// source: PoolStruct.proto

package mine_intf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 矿池列表请求参数
type PoolListRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolListRequest) Reset()         { *m = PoolListRequest{} }
func (m *PoolListRequest) String() string { return proto.CompactTextString(m) }
func (*PoolListRequest) ProtoMessage()    {}
func (*PoolListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_314003dc9bc46c44, []int{0}
}

func (m *PoolListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolListRequest.Unmarshal(m, b)
}
func (m *PoolListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolListRequest.Marshal(b, m, deterministic)
}
func (m *PoolListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolListRequest.Merge(m, src)
}
func (m *PoolListRequest) XXX_Size() int {
	return xxx_messageInfo_PoolListRequest.Size(m)
}
func (m *PoolListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PoolListRequest proto.InternalMessageInfo

func (m *PoolListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PoolListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type PoolListRequestFilterField struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	UserId               int32    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PoolTypeId           string   `protobuf:"bytes,5,opt,name=pool_type_id,json=poolTypeId,proto3" json:"pool_type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolListRequestFilterField) Reset()         { *m = PoolListRequestFilterField{} }
func (m *PoolListRequestFilterField) String() string { return proto.CompactTextString(m) }
func (*PoolListRequestFilterField) ProtoMessage()    {}
func (*PoolListRequestFilterField) Descriptor() ([]byte, []int) {
	return fileDescriptor_314003dc9bc46c44, []int{0, 0}
}

func (m *PoolListRequestFilterField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolListRequestFilterField.Unmarshal(m, b)
}
func (m *PoolListRequestFilterField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolListRequestFilterField.Marshal(b, m, deterministic)
}
func (m *PoolListRequestFilterField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolListRequestFilterField.Merge(m, src)
}
func (m *PoolListRequestFilterField) XXX_Size() int {
	return xxx_messageInfo_PoolListRequestFilterField.Size(m)
}
func (m *PoolListRequestFilterField) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolListRequestFilterField.DiscardUnknown(m)
}

var xxx_messageInfo_PoolListRequestFilterField proto.InternalMessageInfo

func (m *PoolListRequestFilterField) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PoolListRequestFilterField) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PoolListRequestFilterField) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PoolListRequestFilterField) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PoolListRequestFilterField) GetPoolTypeId() string {
	if m != nil {
		return m.PoolTypeId
	}
	return ""
}

// 矿池列表返回
type PoolListResponse struct {
	Code                 ErrorCode   `protobuf:"varint,1,opt,name=code,proto3,enum=mine_intf.ErrorCode" json:"code,omitempty"`
	Data                 []*PoolItem `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PoolListResponse) Reset()         { *m = PoolListResponse{} }
func (m *PoolListResponse) String() string { return proto.CompactTextString(m) }
func (*PoolListResponse) ProtoMessage()    {}
func (*PoolListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_314003dc9bc46c44, []int{1}
}

func (m *PoolListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolListResponse.Unmarshal(m, b)
}
func (m *PoolListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolListResponse.Marshal(b, m, deterministic)
}
func (m *PoolListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolListResponse.Merge(m, src)
}
func (m *PoolListResponse) XXX_Size() int {
	return xxx_messageInfo_PoolListResponse.Size(m)
}
func (m *PoolListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PoolListResponse proto.InternalMessageInfo

func (m *PoolListResponse) GetCode() ErrorCode {
	if m != nil {
		return m.Code
	}
	return ErrorCode_SUCCESS
}

func (m *PoolListResponse) GetData() []*PoolItem {
	if m != nil {
		return m.Data
	}
	return nil
}

type PoolItem struct {
	AutoId               string   `protobuf:"bytes,1,opt,name=auto_id,json=autoId,proto3" json:"auto_id,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	UserId               uint32   `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PoolTypeId           string   `protobuf:"bytes,6,opt,name=pool_type_id,json=poolTypeId,proto3" json:"pool_type_id,omitempty"`
	FriendlyId           uint32   `protobuf:"varint,7,opt,name=friendly_id,json=friendlyId,proto3" json:"friendly_id,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Class                uint32   `protobuf:"varint,9,opt,name=class,proto3" json:"class,omitempty"`
	Abbreviation         string   `protobuf:"bytes,10,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	MinerCount           uint32   `protobuf:"varint,11,opt,name=miner_count,json=minerCount,proto3" json:"miner_count,omitempty"`
	MinerAmount          float64  `protobuf:"fixed64,12,opt,name=miner_amount,json=minerAmount,proto3" json:"miner_amount,omitempty"`
	AmountCoin           uint32   `protobuf:"varint,13,opt,name=amount_coin,json=amountCoin,proto3" json:"amount_coin,omitempty"`
	MinerCoin            float64  `protobuf:"fixed64,14,opt,name=miner_coin,json=minerCoin,proto3" json:"miner_coin,omitempty"`
	BasicYield           float64  `protobuf:"fixed64,15,opt,name=basic_yield,json=basicYield,proto3" json:"basic_yield,omitempty"`
	PresentedYield       float64  `protobuf:"fixed64,16,opt,name=presented_yield,json=presentedYield,proto3" json:"presented_yield,omitempty"`
	AnnualYield          float64  `protobuf:"fixed64,17,opt,name=annual_yield,json=annualYield,proto3" json:"annual_yield,omitempty"`
	Sold                 uint32   `protobuf:"varint,18,opt,name=sold,proto3" json:"sold,omitempty"`
	Expires              uint64   `protobuf:"varint,19,opt,name=expires,proto3" json:"expires,omitempty"`
	ExitAt               uint64   `protobuf:"varint,20,opt,name=exit_at,json=exitAt,proto3" json:"exit_at,omitempty"`
	Duration             uint32   `protobuf:"varint,21,opt,name=duration,proto3" json:"duration,omitempty"`
	MiningDate           string   `protobuf:"bytes,22,opt,name=mining_date,json=miningDate,proto3" json:"mining_date,omitempty"`
	MiningLeft           uint32   `protobuf:"varint,23,opt,name=mining_left,json=miningLeft,proto3" json:"mining_left,omitempty"`
	MiningCount          uint32   `protobuf:"varint,24,opt,name=mining_count,json=miningCount,proto3" json:"mining_count,omitempty"`
	CreatedAt            string   `protobuf:"bytes,25,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,26,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolItem) Reset()         { *m = PoolItem{} }
func (m *PoolItem) String() string { return proto.CompactTextString(m) }
func (*PoolItem) ProtoMessage()    {}
func (*PoolItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_314003dc9bc46c44, []int{2}
}

func (m *PoolItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolItem.Unmarshal(m, b)
}
func (m *PoolItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolItem.Marshal(b, m, deterministic)
}
func (m *PoolItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolItem.Merge(m, src)
}
func (m *PoolItem) XXX_Size() int {
	return xxx_messageInfo_PoolItem.Size(m)
}
func (m *PoolItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolItem.DiscardUnknown(m)
}

var xxx_messageInfo_PoolItem proto.InternalMessageInfo

func (m *PoolItem) GetAutoId() string {
	if m != nil {
		return m.AutoId
	}
	return ""
}

func (m *PoolItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PoolItem) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PoolItem) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PoolItem) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PoolItem) GetPoolTypeId() string {
	if m != nil {
		return m.PoolTypeId
	}
	return ""
}

func (m *PoolItem) GetFriendlyId() uint32 {
	if m != nil {
		return m.FriendlyId
	}
	return 0
}

func (m *PoolItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PoolItem) GetClass() uint32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *PoolItem) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

func (m *PoolItem) GetMinerCount() uint32 {
	if m != nil {
		return m.MinerCount
	}
	return 0
}

func (m *PoolItem) GetMinerAmount() float64 {
	if m != nil {
		return m.MinerAmount
	}
	return 0
}

func (m *PoolItem) GetAmountCoin() uint32 {
	if m != nil {
		return m.AmountCoin
	}
	return 0
}

func (m *PoolItem) GetMinerCoin() float64 {
	if m != nil {
		return m.MinerCoin
	}
	return 0
}

func (m *PoolItem) GetBasicYield() float64 {
	if m != nil {
		return m.BasicYield
	}
	return 0
}

func (m *PoolItem) GetPresentedYield() float64 {
	if m != nil {
		return m.PresentedYield
	}
	return 0
}

func (m *PoolItem) GetAnnualYield() float64 {
	if m != nil {
		return m.AnnualYield
	}
	return 0
}

func (m *PoolItem) GetSold() uint32 {
	if m != nil {
		return m.Sold
	}
	return 0
}

func (m *PoolItem) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *PoolItem) GetExitAt() uint64 {
	if m != nil {
		return m.ExitAt
	}
	return 0
}

func (m *PoolItem) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PoolItem) GetMiningDate() string {
	if m != nil {
		return m.MiningDate
	}
	return ""
}

func (m *PoolItem) GetMiningLeft() uint32 {
	if m != nil {
		return m.MiningLeft
	}
	return 0
}

func (m *PoolItem) GetMiningCount() uint32 {
	if m != nil {
		return m.MiningCount
	}
	return 0
}

func (m *PoolItem) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *PoolItem) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*PoolListRequest)(nil), "mine_intf.PoolListRequest")
	proto.RegisterType((*PoolListRequestFilterField)(nil), "mine_intf.PoolListRequest.filter_field")
	proto.RegisterType((*PoolListResponse)(nil), "mine_intf.PoolListResponse")
	proto.RegisterType((*PoolItem)(nil), "mine_intf.PoolItem")
}

func init() { proto.RegisterFile("PoolStruct.proto", fileDescriptor_314003dc9bc46c44) }

var fileDescriptor_314003dc9bc46c44 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb5, 0xe9, 0x26, 0x4d, 0xa6, 0x69, 0x52, 0xdc, 0xd2, 0x9a, 0xa0, 0x8a, 0x34, 0x97,
	0xe6, 0xd4, 0x43, 0x79, 0x82, 0x10, 0x38, 0x44, 0xea, 0x01, 0x6d, 0xb9, 0x70, 0x5a, 0x39, 0xeb,
	0x49, 0x65, 0x69, 0x63, 0x2f, 0xb6, 0x17, 0x35, 0x3d, 0x71, 0xe1, 0xe5, 0x78, 0x2a, 0xe4, 0xf1,
	0x6e, 0x54, 0x28, 0x9c, 0xe2, 0xf9, 0xe7, 0x9b, 0x19, 0xfb, 0x9f, 0x55, 0xe0, 0xe4, 0xb3, 0x31,
	0xe5, 0xbd, 0xb7, 0x75, 0xe1, 0x6f, 0x2a, 0x6b, 0xbc, 0x61, 0x83, 0xad, 0xd2, 0x98, 0x2b, 0xed,
	0x37, 0x93, 0xf1, 0x07, 0xe1, 0x70, 0x69, 0xb4, 0x6b, 0x72, 0xb3, 0x5f, 0x09, 0x8c, 0x43, 0xc1,
	0x9d, 0x72, 0x3e, 0xc3, 0x6f, 0x35, 0x3a, 0xcf, 0x18, 0xa4, 0x95, 0x78, 0x40, 0x9e, 0x4c, 0x93,
	0x79, 0x37, 0xa3, 0x33, 0x7b, 0x0b, 0x83, 0xf0, 0x9b, 0x3b, 0xf5, 0x84, 0xbc, 0x43, 0x89, 0x7e,
	0x10, 0xee, 0xd5, 0x13, 0x4e, 0x7e, 0x26, 0x30, 0xdc, 0xa8, 0xd2, 0xa3, 0xcd, 0x37, 0x0a, 0x4b,
	0xc9, 0x46, 0xd0, 0x51, 0x92, 0xea, 0x07, 0x59, 0x47, 0x49, 0x76, 0x0e, 0x3d, 0xe7, 0x85, 0xaf,
	0x5d, 0x53, 0xda, 0x44, 0x61, 0x92, 0xdf, 0x55, 0xc8, 0x0f, 0xe2, 0xa4, 0x70, 0x66, 0x17, 0x70,
	0x58, 0x3b, 0xb4, 0xb9, 0x92, 0x3c, 0x8d, 0x70, 0x08, 0x57, 0x92, 0x4d, 0x61, 0x58, 0x19, 0x53,
	0xe6, 0x81, 0x0a, 0xd9, 0x2e, 0xb5, 0x87, 0xa0, 0x7d, 0xd9, 0x55, 0xb8, 0x92, 0x33, 0x8c, 0x8f,
	0x8f, 0x6f, 0x71, 0x95, 0xd1, 0x0e, 0xd9, 0x1c, 0xd2, 0xc2, 0xc8, 0xf8, 0x98, 0xd1, 0xed, 0xd9,
	0xcd, 0xde, 0x8b, 0x9b, 0x4f, 0xd6, 0x1a, 0xbb, 0x34, 0x12, 0x33, 0x22, 0xd8, 0x35, 0xa4, 0x52,
	0x78, 0xc1, 0x3b, 0xd3, 0x83, 0xf9, 0xd1, 0xed, 0xe9, 0x33, 0x32, 0x34, 0x5d, 0x79, 0xdc, 0x66,
	0x04, 0xcc, 0x7e, 0xf4, 0xa0, 0xdf, 0x4a, 0xe1, 0xba, 0xa2, 0xf6, 0x26, 0xdf, 0xbf, 0xb7, 0x17,
	0xc2, 0x55, 0xeb, 0x41, 0xe7, 0x1f, 0x1e, 0x84, 0xd7, 0x1e, 0xbf, 0xf0, 0x20, 0x25, 0xf5, 0x85,
	0x07, 0xdd, 0x08, 0xff, 0xc7, 0x83, 0xde, 0xdf, 0x1e, 0xb0, 0x77, 0x70, 0xb4, 0xb1, 0x0a, 0xb5,
	0x2c, 0x77, 0x01, 0x38, 0xa4, 0x72, 0x68, 0xa5, 0x95, 0x0c, 0xf3, 0xb4, 0xd8, 0x22, 0xef, 0x53,
	0x29, 0x9d, 0xd9, 0x19, 0x74, 0x8b, 0x52, 0x38, 0xc7, 0x07, 0x84, 0xc7, 0x80, 0xcd, 0x60, 0x28,
	0xd6, 0x6b, 0x8b, 0xdf, 0x95, 0xf0, 0xca, 0x68, 0x0e, 0x54, 0xf1, 0x87, 0x16, 0xc6, 0x05, 0x9f,
	0x6c, 0x5e, 0x98, 0x5a, 0x7b, 0x7e, 0x14, 0xc7, 0x91, 0xb4, 0x0c, 0x0a, 0xbb, 0x82, 0x61, 0x04,
	0xc4, 0x96, 0x88, 0xe1, 0x34, 0x99, 0x27, 0x59, 0x2c, 0x5a, 0x90, 0x14, 0x7a, 0xc4, 0x64, 0x5e,
	0x18, 0xa5, 0xf9, 0x71, 0xec, 0x11, 0xa5, 0xa5, 0x51, 0x9a, 0x5d, 0x02, 0xb4, 0x43, 0x94, 0xe6,
	0x23, 0xea, 0x30, 0x68, 0x66, 0x28, 0xba, 0xc3, 0x5a, 0x38, 0x55, 0xe4, 0xbb, 0xf0, 0xf1, 0xf1,
	0x31, 0xe5, 0x81, 0xa4, 0xaf, 0xf4, 0x39, 0x5e, 0xc3, 0xb8, 0xb2, 0xe8, 0x50, 0x7b, 0x94, 0x0d,
	0x74, 0x42, 0xd0, 0x68, 0x2f, 0x47, 0xf0, 0x0a, 0x86, 0x42, 0xeb, 0x5a, 0x94, 0x0d, 0xf5, 0x2a,
	0x5e, 0x36, 0x6a, 0x11, 0x61, 0x90, 0x3a, 0x53, 0x4a, 0xce, 0xe2, 0xba, 0xc2, 0x99, 0x71, 0x38,
	0xc4, 0xc7, 0x4a, 0x59, 0x74, 0xfc, 0x74, 0x9a, 0xcc, 0xd3, 0xac, 0x0d, 0xc3, 0x22, 0xf1, 0x51,
	0xf9, 0x5c, 0x78, 0x7e, 0x46, 0x99, 0x5e, 0x08, 0x17, 0x9e, 0x4d, 0xa0, 0x2f, 0x6b, 0x1b, 0x7d,
	0x7d, 0x4d, 0xad, 0xf6, 0x71, 0xe3, 0xa9, 0xd2, 0x0f, 0xb9, 0x14, 0x1e, 0xf9, 0x79, 0xdc, 0x71,
	0x94, 0x3e, 0x0a, 0x8f, 0xcf, 0x80, 0x12, 0x37, 0x9e, 0x5f, 0xec, 0x4d, 0x57, 0xfa, 0xe1, 0x0e,
	0x37, 0xad, 0xe9, 0x01, 0x88, 0x6b, 0xe1, 0x44, 0x34, 0x45, 0x71, 0x2f, 0x97, 0x00, 0x85, 0x45,
	0x11, 0x1c, 0x11, 0x9e, 0xbf, 0xa1, 0x19, 0x83, 0x46, 0x59, 0x50, 0xba, 0xae, 0x64, 0x9b, 0x9e,
	0xc4, 0x74, 0xa3, 0x2c, 0xfc, 0xba, 0x47, 0xff, 0x1e, 0xef, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x39, 0x42, 0x0f, 0xe6, 0x6d, 0x04, 0x00, 0x00,
}
