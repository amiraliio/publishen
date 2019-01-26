// Code generated by protoc-gen-go. DO NOT EDIT.
// source: publish.proto

package publishen_publish_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestClass int32

const (
	Request_I_WANT RequestClass = 0
	Request_I_HAVE RequestClass = 1
)

var RequestClass_name = map[int32]string{
	0: "I_WANT",
	1: "I_HAVE",
}
var RequestClass_value = map[string]int32{
	"I_WANT": 0,
	"I_HAVE": 1,
}

func (x RequestClass) String() string {
	return proto.EnumName(RequestClass_name, int32(x))
}
func (RequestClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{0, 0}
}

type RequestGroup int32

const (
	Request_ARTICLE RequestGroup = 0
	Request_PRODUCT RequestGroup = 1
	Request_GALLERY RequestGroup = 2
	Request_VIDEO   RequestGroup = 3
	Request_JOB     RequestGroup = 4
	Request_PROJECT RequestGroup = 5
)

var RequestGroup_name = map[int32]string{
	0: "ARTICLE",
	1: "PRODUCT",
	2: "GALLERY",
	3: "VIDEO",
	4: "JOB",
	5: "PROJECT",
}
var RequestGroup_value = map[string]int32{
	"ARTICLE": 0,
	"PRODUCT": 1,
	"GALLERY": 2,
	"VIDEO":   3,
	"JOB":     4,
	"PROJECT": 5,
}

func (x RequestGroup) String() string {
	return proto.EnumName(RequestGroup_name, int32(x))
}
func (RequestGroup) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{0, 1}
}

type PublishPublishStatus int32

const (
	Publish_ACTIVE     PublishPublishStatus = 0
	Publish_DEACTIVATE PublishPublishStatus = 1
)

var PublishPublishStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "DEACTIVATE",
}
var PublishPublishStatus_value = map[string]int32{
	"ACTIVE":     0,
	"DEACTIVATE": 1,
}

func (x PublishPublishStatus) String() string {
	return proto.EnumName(PublishPublishStatus_name, int32(x))
}
func (PublishPublishStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{1, 0}
}

type PublishShareState int32

const (
	// publish
	Publish_PUBLIC  PublishShareState = 0
	Publish_PRIVATE PublishShareState = 1
)

var PublishShareState_name = map[int32]string{
	0: "PUBLIC",
	1: "PRIVATE",
}
var PublishShareState_value = map[string]int32{
	"PUBLIC":  0,
	"PRIVATE": 1,
}

func (x PublishShareState) String() string {
	return proto.EnumName(PublishShareState_name, int32(x))
}
func (PublishShareState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{1, 1}
}

type PublishClass int32

const (
	Publish_I_WANT PublishClass = 0
	Publish_I_HAVE PublishClass = 1
)

var PublishClass_name = map[int32]string{
	0: "I_WANT",
	1: "I_HAVE",
}
var PublishClass_value = map[string]int32{
	"I_WANT": 0,
	"I_HAVE": 1,
}

func (x PublishClass) String() string {
	return proto.EnumName(PublishClass_name, int32(x))
}
func (PublishClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{1, 2}
}

type PublishGroup int32

const (
	Publish_ARTICLE PublishGroup = 0
	Publish_PRODUCT PublishGroup = 1
	Publish_GALLERY PublishGroup = 2
	Publish_VIDEO   PublishGroup = 3
	Publish_JOB     PublishGroup = 4
	Publish_PROJECT PublishGroup = 5
)

var PublishGroup_name = map[int32]string{
	0: "ARTICLE",
	1: "PRODUCT",
	2: "GALLERY",
	3: "VIDEO",
	4: "JOB",
	5: "PROJECT",
}
var PublishGroup_value = map[string]int32{
	"ARTICLE": 0,
	"PRODUCT": 1,
	"GALLERY": 2,
	"VIDEO":   3,
	"JOB":     4,
	"PROJECT": 5,
}

func (x PublishGroup) String() string {
	return proto.EnumName(PublishGroup_name, int32(x))
}
func (PublishGroup) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{1, 3}
}

type Request struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Request) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Request) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Publish struct {
	ID                   string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Banner               string               `protobuf:"bytes,4,opt,name=banner,proto3" json:"banner,omitempty"`
	Images               []string             `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
	Location             *Location            `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Attributes           []*Attribute         `protobuf:"bytes,7,rep,name=attributes,proto3" json:"attributes,omitempty"`
	ShareGroupID         string               `protobuf:"bytes,9,opt,name=shareGroupID,proto3" json:"shareGroupID,omitempty"`
	UserID               string               `protobuf:"bytes,10,opt,name=userID,proto3" json:"userID,omitempty"`
	Tags                 []string             `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	Key                  string               `protobuf:"bytes,12,opt,name=Key,proto3" json:"Key,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Publish) Reset()         { *m = Publish{} }
func (m *Publish) String() string { return proto.CompactTextString(m) }
func (*Publish) ProtoMessage()    {}
func (*Publish) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{1}
}
func (m *Publish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Publish.Unmarshal(m, b)
}
func (m *Publish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Publish.Marshal(b, m, deterministic)
}
func (dst *Publish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Publish.Merge(dst, src)
}
func (m *Publish) XXX_Size() int {
	return xxx_messageInfo_Publish.Size(m)
}
func (m *Publish) XXX_DiscardUnknown() {
	xxx_messageInfo_Publish.DiscardUnknown(m)
}

var xxx_messageInfo_Publish proto.InternalMessageInfo

func (m *Publish) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Publish) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Publish) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Publish) GetBanner() string {
	if m != nil {
		return m.Banner
	}
	return ""
}

func (m *Publish) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Publish) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Publish) GetAttributes() []*Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Publish) GetShareGroupID() string {
	if m != nil {
		return m.ShareGroupID
	}
	return ""
}

func (m *Publish) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Publish) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Publish) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Publish) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Publish) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Location struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{2}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Attribute struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{3}
}
func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute.Unmarshal(m, b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
}
func (dst *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(dst, src)
}
func (m *Attribute) XXX_Size() int {
	return xxx_messageInfo_Attribute.Size(m)
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	ErrorMessage         string     `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Success              bool       `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	DocumentID           string     `protobuf:"bytes,3,opt,name=documentID,proto3" json:"documentID,omitempty"`
	Publish              *Publish   `protobuf:"bytes,4,opt,name=publish,proto3" json:"publish,omitempty"`
	Publishes            []*Publish `protobuf:"bytes,5,rep,name=publishes,proto3" json:"publishes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_e133d7191f344f52, []int{4}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetDocumentID() string {
	if m != nil {
		return m.DocumentID
	}
	return ""
}

func (m *Response) GetPublish() *Publish {
	if m != nil {
		return m.Publish
	}
	return nil
}

func (m *Response) GetPublishes() []*Publish {
	if m != nil {
		return m.Publishes
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "publishen.publish.service.Request")
	proto.RegisterType((*Publish)(nil), "publishen.publish.service.Publish")
	proto.RegisterType((*Location)(nil), "publishen.publish.service.Location")
	proto.RegisterType((*Attribute)(nil), "publishen.publish.service.Attribute")
	proto.RegisterType((*Response)(nil), "publishen.publish.service.Response")
	proto.RegisterEnum("publishen.publish.service.RequestClass", RequestClass_name, RequestClass_value)
	proto.RegisterEnum("publishen.publish.service.RequestGroup", RequestGroup_name, RequestGroup_value)
	proto.RegisterEnum("publishen.publish.service.PublishPublishStatus", PublishPublishStatus_name, PublishPublishStatus_value)
	proto.RegisterEnum("publishen.publish.service.PublishShareState", PublishShareState_name, PublishShareState_value)
	proto.RegisterEnum("publishen.publish.service.PublishClass", PublishClass_name, PublishClass_value)
	proto.RegisterEnum("publishen.publish.service.PublishGroup", PublishGroup_name, PublishGroup_value)
}

func init() { proto.RegisterFile("publish.proto", fileDescriptor_publish_e133d7191f344f52) }

var fileDescriptor_publish_e133d7191f344f52 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x61, 0x6a, 0xdb, 0x4c,
	0x10, 0x8d, 0xec, 0xc8, 0xb6, 0xc6, 0xb1, 0x3f, 0xb1, 0x7c, 0x14, 0xd5, 0x94, 0xc6, 0xa8, 0x2d,
	0x18, 0x0a, 0x0e, 0x38, 0xbf, 0x0a, 0x85, 0x56, 0xb1, 0x4c, 0xaa, 0xd4, 0xad, 0xcd, 0xd6, 0x49,
	0x68, 0x28, 0x04, 0x59, 0x9e, 0x38, 0xa2, 0xb2, 0xa5, 0x6a, 0x57, 0x81, 0xf4, 0x40, 0xbd, 0x44,
	0x8f, 0xd2, 0x0b, 0xf4, 0x18, 0x65, 0x57, 0x5a, 0xc7, 0x29, 0xd4, 0xcd, 0x9f, 0xf4, 0xdf, 0xbc,
	0xd9, 0xf7, 0x86, 0xd9, 0x9d, 0x37, 0x0b, 0x8d, 0x24, 0x9b, 0x46, 0x21, 0xbb, 0xec, 0x26, 0x69,
	0xcc, 0x63, 0xf2, 0xb0, 0x80, 0xb8, 0xec, 0xaa, 0x03, 0x86, 0xe9, 0x55, 0x18, 0x60, 0x6b, 0x77,
	0x1e, 0xc7, 0xf3, 0x08, 0xf7, 0x24, 0x71, 0x9a, 0x5d, 0xec, 0xf1, 0x70, 0x81, 0x8c, 0xfb, 0x8b,
	0x24, 0xd7, 0xda, 0xdf, 0x35, 0xa8, 0x52, 0xfc, 0x92, 0x21, 0xe3, 0xe4, 0x01, 0x54, 0x32, 0x86,
	0xa9, 0xe7, 0x5a, 0x5a, 0x5b, 0xeb, 0x18, 0xb4, 0x40, 0x84, 0xc0, 0x76, 0xe2, 0xcf, 0xd1, 0x2a,
	0xb5, 0xb5, 0x8e, 0x4e, 0x65, 0x2c, 0x72, 0x2c, 0xfc, 0x8a, 0x56, 0x39, 0xcf, 0x89, 0xd8, 0xde,
	0x05, 0x3d, 0x88, 0x7c, 0xc6, 0x08, 0x40, 0xc5, 0x3b, 0x3f, 0x75, 0xde, 0x4f, 0xcc, 0xad, 0x3c,
	0x7e, 0xe3, 0x9c, 0x0c, 0x4c, 0xcd, 0x1e, 0x81, 0x3e, 0x4f, 0xe3, 0x2c, 0x21, 0x75, 0xa8, 0x3a,
	0x74, 0xe2, 0xf5, 0x87, 0x03, 0x73, 0x4b, 0x80, 0x31, 0x1d, 0xb9, 0xc7, 0xfd, 0x89, 0xa9, 0x09,
	0x70, 0xe8, 0x0c, 0x87, 0x03, 0xfa, 0xd1, 0x2c, 0x11, 0x03, 0xf4, 0x13, 0xcf, 0x1d, 0x8c, 0xcc,
	0x32, 0xa9, 0x42, 0xf9, 0x68, 0x74, 0x60, 0x6e, 0x17, 0xec, 0xa3, 0x41, 0x7f, 0x62, 0xea, 0xf6,
	0x37, 0x1d, 0xaa, 0xe3, 0xfc, 0xca, 0xa4, 0x09, 0xa5, 0x55, 0xe7, 0x25, 0xcf, 0x25, 0xff, 0x83,
	0xce, 0x43, 0x1e, 0xe5, 0x6d, 0x1b, 0x34, 0x07, 0xa4, 0x0d, 0xf5, 0x19, 0xb2, 0x20, 0x0d, 0x13,
	0x1e, 0xc6, 0x4b, 0xd9, 0xbe, 0x41, 0xd7, 0x53, 0xe2, 0x15, 0xa6, 0xfe, 0x72, 0x89, 0xa9, 0xb5,
	0x9d, 0xbf, 0x42, 0x8e, 0x44, 0x3e, 0x5c, 0xf8, 0x73, 0x64, 0x96, 0xde, 0x2e, 0x8b, 0x7c, 0x8e,
	0xc8, 0x2b, 0xa8, 0x45, 0x71, 0xe0, 0xcb, 0x72, 0x95, 0xb6, 0xd6, 0xa9, 0xf7, 0x9e, 0x74, 0xff,
	0x38, 0x90, 0xee, 0xb0, 0xa0, 0xd2, 0x95, 0x88, 0xb8, 0x00, 0x3e, 0xe7, 0x69, 0x38, 0xcd, 0x38,
	0x32, 0xab, 0xda, 0x2e, 0x77, 0xea, 0xbd, 0xa7, 0x1b, 0x4a, 0x38, 0x8a, 0x4c, 0xd7, 0x74, 0xc4,
	0x86, 0x1d, 0x76, 0xe9, 0xa7, 0x78, 0x28, 0x1e, 0xd8, 0x73, 0x2d, 0x43, 0x36, 0x7f, 0x2b, 0xb7,
	0x36, 0x60, 0xf8, 0x7d, 0xc0, 0xdc, 0x9f, 0x33, 0xab, 0x2e, 0x2f, 0x26, 0x63, 0x62, 0x42, 0xf9,
	0x2d, 0x5e, 0x5b, 0x3b, 0x92, 0x28, 0x42, 0xf2, 0x02, 0x20, 0x48, 0xd1, 0xe7, 0x38, 0x3b, 0xf7,
	0xb9, 0xd5, 0x90, 0x57, 0x6d, 0x75, 0x73, 0x83, 0x75, 0x95, 0xc1, 0xba, 0x13, 0x65, 0x30, 0x6a,
	0x14, 0x6c, 0x87, 0x0b, 0x69, 0x96, 0xcc, 0x94, 0xb4, 0xf9, 0x77, 0x69, 0xc1, 0x76, 0xb8, 0xfd,
	0x7c, 0xe5, 0xf6, 0x0f, 0xdc, 0xe7, 0x99, 0x34, 0x97, 0xd3, 0x9f, 0x78, 0x27, 0xc2, 0x3a, 0x4d,
	0x00, 0x77, 0x20, 0x91, 0x33, 0x11, 0x06, 0x7b, 0x06, 0x20, 0x2f, 0x2c, 0xa8, 0x28, 0x98, 0xe3,
	0xe3, 0x83, 0xa1, 0xd7, 0x57, 0x26, 0x53, 0xb4, 0x7f, 0x6f, 0x54, 0x17, 0x6a, 0x6a, 0xf2, 0xa4,
	0x05, 0xb5, 0xc8, 0xe7, 0x21, 0xcf, 0x66, 0x28, 0xed, 0x5a, 0xa2, 0x2b, 0x4c, 0x1e, 0x81, 0x11,
	0xc5, 0xcb, 0x79, 0x7e, 0x58, 0x92, 0x87, 0x37, 0x09, 0x7b, 0x1f, 0x8c, 0xd5, 0xf0, 0xc5, 0x80,
	0x3e, 0xe3, 0x75, 0x61, 0x78, 0x11, 0x0a, 0xc7, 0x5f, 0xf9, 0x51, 0xb6, 0x72, 0xbc, 0x04, 0xf6,
	0x4f, 0x0d, 0x6a, 0x14, 0x59, 0x12, 0x2f, 0x19, 0x0a, 0x97, 0x60, 0x9a, 0xc6, 0xe9, 0x3b, 0x64,
	0x4c, 0xac, 0x74, 0xae, 0xbe, 0x95, 0x23, 0x16, 0x54, 0x59, 0x16, 0x04, 0xc8, 0x98, 0x2c, 0x54,
	0xa3, 0x0a, 0x92, 0xc7, 0x00, 0xb3, 0x38, 0xc8, 0x16, 0xb8, 0xe4, 0x9e, 0x5b, 0xec, 0xce, 0x5a,
	0x86, 0xbc, 0x84, 0x6a, 0x31, 0x2b, 0xb9, 0x3b, 0xf5, 0x9e, 0xbd, 0xc1, 0xc6, 0xc5, 0xde, 0x52,
	0x25, 0x21, 0xaf, 0xc1, 0x50, 0xec, 0x7c, 0xc7, 0xee, 0xa6, 0xbf, 0x11, 0xf5, 0x7e, 0x94, 0xa1,
	0xa9, 0xcc, 0x92, 0xb3, 0xc8, 0x19, 0x34, 0xfa, 0xd2, 0x86, 0xea, 0x9b, 0xb8, 0x43, 0xc9, 0xd6,
	0xa6, 0x05, 0x56, 0x4f, 0x69, 0x6f, 0x91, 0x4f, 0xf0, 0xdf, 0x30, 0x64, 0x7c, 0x74, 0x31, 0x56,
	0x1d, 0x6c, 0xac, 0x5e, 0x7c, 0xb3, 0x77, 0xad, 0x7e, 0x0a, 0x70, 0x88, 0xfc, 0x1e, 0xda, 0x3e,
	0x83, 0xc6, 0xb1, 0x5c, 0xaf, 0xfb, 0xa9, 0xed, 0x62, 0x84, 0xf7, 0x51, 0x7b, 0x5a, 0x91, 0x1f,
	0xc5, 0xfe, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xce, 0x96, 0xb4, 0xfe, 0x06, 0x00, 0x00,
}
