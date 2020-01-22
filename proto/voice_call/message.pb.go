// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voice_call/message.proto

package message

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CallingRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	UrlMedia             string   `protobuf:"bytes,2,opt,name=urlMedia,proto3" json:"urlMedia,omitempty"`
	UrlCallback          string   `protobuf:"bytes,3,opt,name=urlCallback,proto3" json:"urlCallback,omitempty"`
	NotificationID       string   `protobuf:"bytes,4,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	CallID               string   `protobuf:"bytes,5,opt,name=callID,proto3" json:"callID,omitempty"`
	CallbackType         string   `protobuf:"bytes,6,opt,name=callbackType,proto3" json:"callbackType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallingRequest) Reset()         { *m = CallingRequest{} }
func (m *CallingRequest) String() string { return proto.CompactTextString(m) }
func (*CallingRequest) ProtoMessage()    {}
func (*CallingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{0}
}

func (m *CallingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallingRequest.Unmarshal(m, b)
}
func (m *CallingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallingRequest.Marshal(b, m, deterministic)
}
func (m *CallingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallingRequest.Merge(m, src)
}
func (m *CallingRequest) XXX_Size() int {
	return xxx_messageInfo_CallingRequest.Size(m)
}
func (m *CallingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallingRequest proto.InternalMessageInfo

func (m *CallingRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CallingRequest) GetUrlMedia() string {
	if m != nil {
		return m.UrlMedia
	}
	return ""
}

func (m *CallingRequest) GetUrlCallback() string {
	if m != nil {
		return m.UrlCallback
	}
	return ""
}

func (m *CallingRequest) GetNotificationID() string {
	if m != nil {
		return m.NotificationID
	}
	return ""
}

func (m *CallingRequest) GetCallID() string {
	if m != nil {
		return m.CallID
	}
	return ""
}

func (m *CallingRequest) GetCallbackType() string {
	if m != nil {
		return m.CallbackType
	}
	return ""
}

type MessageRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	UrlCallback          string   `protobuf:"bytes,3,opt,name=urlCallback,proto3" json:"urlCallback,omitempty"`
	NotificationID       string   `protobuf:"bytes,4,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	MessageID            string   `protobuf:"bytes,5,opt,name=messageID,proto3" json:"messageID,omitempty"`
	CallbackType         string   `protobuf:"bytes,6,opt,name=callbackType,proto3" json:"callbackType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{1}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MessageRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MessageRequest) GetUrlCallback() string {
	if m != nil {
		return m.UrlCallback
	}
	return ""
}

func (m *MessageRequest) GetNotificationID() string {
	if m != nil {
		return m.NotificationID
	}
	return ""
}

func (m *MessageRequest) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *MessageRequest) GetCallbackType() string {
	if m != nil {
		return m.CallbackType
	}
	return ""
}

type GetCallRequest struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCallRequest) Reset()         { *m = GetCallRequest{} }
func (m *GetCallRequest) String() string { return proto.CompactTextString(m) }
func (*GetCallRequest) ProtoMessage()    {}
func (*GetCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{2}
}

func (m *GetCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCallRequest.Unmarshal(m, b)
}
func (m *GetCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCallRequest.Marshal(b, m, deterministic)
}
func (m *GetCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCallRequest.Merge(m, src)
}
func (m *GetCallRequest) XXX_Size() int {
	return xxx_messageInfo_GetCallRequest.Size(m)
}
func (m *GetCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCallRequest proto.InternalMessageInfo

func (m *GetCallRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

type GetCallResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	From                 string   `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	UrlMedia             string   `protobuf:"bytes,6,opt,name=urlMedia,proto3" json:"urlMedia,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,7,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	NotificationID       string   `protobuf:"bytes,8,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	CallID               string   `protobuf:"bytes,9,opt,name=callID,proto3" json:"callID,omitempty"`
	Sid                  string   `protobuf:"bytes,10,opt,name=sid,proto3" json:"sid,omitempty"`
	Error                string   `protobuf:"bytes,11,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCallResponse) Reset()         { *m = GetCallResponse{} }
func (m *GetCallResponse) String() string { return proto.CompactTextString(m) }
func (*GetCallResponse) ProtoMessage()    {}
func (*GetCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{3}
}

func (m *GetCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCallResponse.Unmarshal(m, b)
}
func (m *GetCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCallResponse.Marshal(b, m, deterministic)
}
func (m *GetCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCallResponse.Merge(m, src)
}
func (m *GetCallResponse) XXX_Size() int {
	return xxx_messageInfo_GetCallResponse.Size(m)
}
func (m *GetCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCallResponse proto.InternalMessageInfo

func (m *GetCallResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetCallResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetCallResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetCallResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GetCallResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *GetCallResponse) GetUrlMedia() string {
	if m != nil {
		return m.UrlMedia
	}
	return ""
}

func (m *GetCallResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *GetCallResponse) GetNotificationID() string {
	if m != nil {
		return m.NotificationID
	}
	return ""
}

func (m *GetCallResponse) GetCallID() string {
	if m != nil {
		return m.CallID
	}
	return ""
}

func (m *GetCallResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *GetCallResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetMessageResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	From                 string   `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	Body                 string   `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,7,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	NotificationID       string   `protobuf:"bytes,8,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	MessageID            string   `protobuf:"bytes,9,opt,name=messageID,proto3" json:"messageID,omitempty"`
	Sid                  string   `protobuf:"bytes,10,opt,name=sid,proto3" json:"sid,omitempty"`
	Error                string   `protobuf:"bytes,11,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageResponse) Reset()         { *m = GetMessageResponse{} }
func (m *GetMessageResponse) String() string { return proto.CompactTextString(m) }
func (*GetMessageResponse) ProtoMessage()    {}
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{4}
}

func (m *GetMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResponse.Unmarshal(m, b)
}
func (m *GetMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResponse.Marshal(b, m, deterministic)
}
func (m *GetMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResponse.Merge(m, src)
}
func (m *GetMessageResponse) XXX_Size() int {
	return xxx_messageInfo_GetMessageResponse.Size(m)
}
func (m *GetMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResponse proto.InternalMessageInfo

func (m *GetMessageResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetMessageResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetMessageResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetMessageResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GetMessageResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *GetMessageResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GetMessageResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *GetMessageResponse) GetNotificationID() string {
	if m != nil {
		return m.NotificationID
	}
	return ""
}

func (m *GetMessageResponse) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *GetMessageResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *GetMessageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CallingResponse struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	DateCreated          string   `protobuf:"bytes,5,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	DateUpdated          string   `protobuf:"bytes,6,opt,name=dateUpdated,proto3" json:"dateUpdated,omitempty"`
	StartTime            string   `protobuf:"bytes,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              string   `protobuf:"bytes,8,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Duration             string   `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	ForwardedFrom        string   `protobuf:"bytes,10,opt,name=forwardedFrom,proto3" json:"forwardedFrom,omitempty"`
	Price                float32  `protobuf:"fixed32,11,opt,name=price,proto3" json:"price,omitempty"`
	PriceUnit            string   `protobuf:"bytes,12,opt,name=priceUnit,proto3" json:"priceUnit,omitempty"`
	Type                 string   `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,14,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Error                string   `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
	NotificationID       string   `protobuf:"bytes,16,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	CallID               string   `protobuf:"bytes,17,opt,name=callID,proto3" json:"callID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallingResponse) Reset()         { *m = CallingResponse{} }
func (m *CallingResponse) String() string { return proto.CompactTextString(m) }
func (*CallingResponse) ProtoMessage()    {}
func (*CallingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{5}
}

func (m *CallingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallingResponse.Unmarshal(m, b)
}
func (m *CallingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallingResponse.Marshal(b, m, deterministic)
}
func (m *CallingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallingResponse.Merge(m, src)
}
func (m *CallingResponse) XXX_Size() int {
	return xxx_messageInfo_CallingResponse.Size(m)
}
func (m *CallingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallingResponse proto.InternalMessageInfo

func (m *CallingResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *CallingResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *CallingResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CallingResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CallingResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *CallingResponse) GetDateUpdated() string {
	if m != nil {
		return m.DateUpdated
	}
	return ""
}

func (m *CallingResponse) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *CallingResponse) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *CallingResponse) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *CallingResponse) GetForwardedFrom() string {
	if m != nil {
		return m.ForwardedFrom
	}
	return ""
}

func (m *CallingResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CallingResponse) GetPriceUnit() string {
	if m != nil {
		return m.PriceUnit
	}
	return ""
}

func (m *CallingResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CallingResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *CallingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CallingResponse) GetNotificationID() string {
	if m != nil {
		return m.NotificationID
	}
	return ""
}

func (m *CallingResponse) GetCallID() string {
	if m != nil {
		return m.CallID
	}
	return ""
}

type MessageResponse struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	DateCreated          string   `protobuf:"bytes,5,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	DateSent             string   `protobuf:"bytes,6,opt,name=dateSent,proto3" json:"dateSent,omitempty"`
	DateUpdated          string   `protobuf:"bytes,7,opt,name=dateUpdated,proto3" json:"dateUpdated,omitempty"`
	Segments             string   `protobuf:"bytes,8,opt,name=segments,proto3" json:"segments,omitempty"`
	Price                float32  `protobuf:"fixed32,9,opt,name=price,proto3" json:"price,omitempty"`
	PriceUnit            string   `protobuf:"bytes,10,opt,name=priceUnit,proto3" json:"priceUnit,omitempty"`
	Type                 string   `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,12,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Body                 string   `protobuf:"bytes,13,opt,name=body,proto3" json:"body,omitempty"`
	Error                string   `protobuf:"bytes,14,opt,name=error,proto3" json:"error,omitempty"`
	NotificationID       string   `protobuf:"bytes,15,opt,name=notificationID,proto3" json:"notificationID,omitempty"`
	MessageID            string   `protobuf:"bytes,16,opt,name=messageID,proto3" json:"messageID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bc5fe70d0b5992, []int{6}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *MessageResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MessageResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MessageResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MessageResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *MessageResponse) GetDateSent() string {
	if m != nil {
		return m.DateSent
	}
	return ""
}

func (m *MessageResponse) GetDateUpdated() string {
	if m != nil {
		return m.DateUpdated
	}
	return ""
}

func (m *MessageResponse) GetSegments() string {
	if m != nil {
		return m.Segments
	}
	return ""
}

func (m *MessageResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MessageResponse) GetPriceUnit() string {
	if m != nil {
		return m.PriceUnit
	}
	return ""
}

func (m *MessageResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MessageResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *MessageResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MessageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *MessageResponse) GetNotificationID() string {
	if m != nil {
		return m.NotificationID
	}
	return ""
}

func (m *MessageResponse) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func init() {
	proto.RegisterType((*CallingRequest)(nil), "CallingRequest")
	proto.RegisterType((*MessageRequest)(nil), "MessageRequest")
	proto.RegisterType((*GetCallRequest)(nil), "GetCallRequest")
	proto.RegisterType((*GetCallResponse)(nil), "GetCallResponse")
	proto.RegisterType((*GetMessageResponse)(nil), "GetMessageResponse")
	proto.RegisterType((*CallingResponse)(nil), "CallingResponse")
	proto.RegisterType((*MessageResponse)(nil), "MessageResponse")
}

func init() { proto.RegisterFile("voice_call/message.proto", fileDescriptor_78bc5fe70d0b5992) }

var fileDescriptor_78bc5fe70d0b5992 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x6e, 0x9c, 0x26, 0x4d, 0x26, 0x89, 0x9d, 0xdf, 0xfe, 0x10, 0xb2, 0x2a, 0x0e, 0x95, 0x85,
	0x10, 0x27, 0x57, 0x6a, 0x4f, 0x1c, 0xa1, 0x08, 0xc4, 0xa1, 0x97, 0xa4, 0x3d, 0x23, 0xc7, 0xde,
	0x44, 0x16, 0x89, 0x37, 0xac, 0x37, 0x45, 0x79, 0x24, 0x78, 0x01, 0xae, 0x88, 0x17, 0x41, 0xbc,
	0x09, 0xfb, 0x67, 0xbc, 0xb1, 0x9d, 0xa4, 0xe9, 0xa1, 0x82, 0xdb, 0xcc, 0xb7, 0xbb, 0xce, 0x7c,
	0xdf, 0x37, 0xb3, 0x1b, 0xf0, 0xef, 0x58, 0x1a, 0xd3, 0x8f, 0x71, 0x34, 0x9f, 0x9f, 0x2f, 0x68,
	0x9e, 0x47, 0x33, 0x1a, 0x2e, 0x39, 0x13, 0x2c, 0xf8, 0xd9, 0x00, 0xf7, 0x4a, 0xc2, 0x69, 0x36,
	0x1b, 0xd1, 0xcf, 0x2b, 0x9a, 0x0b, 0xe2, 0x82, 0x23, 0x98, 0xdf, 0x38, 0x6b, 0xbc, 0xec, 0x8e,
	0x64, 0x44, 0x4e, 0xa1, 0xb3, 0xe2, 0xf3, 0x6b, 0x9a, 0xa4, 0x91, 0xef, 0x68, 0xd4, 0xe6, 0xe4,
	0x0c, 0x7a, 0x32, 0x56, 0x1f, 0x98, 0x44, 0xf1, 0x27, 0xbf, 0xa9, 0x97, 0xcb, 0x10, 0x79, 0x01,
	0x6e, 0xc6, 0x44, 0x3a, 0x4d, 0xe3, 0x48, 0xa4, 0x2c, 0xfb, 0xf0, 0xd6, 0x3f, 0xd6, 0x9b, 0x6a,
	0x28, 0x79, 0x0a, 0x6d, 0x55, 0x9e, 0x5c, 0x6f, 0xe9, 0x75, 0xcc, 0x48, 0x00, 0xfd, 0x18, 0xbf,
	0x75, 0xb3, 0x5e, 0x52, 0xbf, 0xad, 0x57, 0x2b, 0x58, 0xf0, 0x43, 0x92, 0xb8, 0x36, 0xb4, 0xf6,
	0x91, 0x20, 0x70, 0x3c, 0x61, 0xc9, 0x1a, 0x09, 0xe8, 0xf8, 0x11, 0x8b, 0x7f, 0x06, 0x5d, 0x94,
	0xd5, 0xd6, 0xbf, 0x01, 0x1e, 0x44, 0x21, 0x00, 0xf7, 0x3d, 0x15, 0xea, 0x87, 0x0b, 0x06, 0x43,
	0x68, 0xe6, 0x69, 0x82, 0x14, 0x54, 0x18, 0x7c, 0x75, 0xc0, 0xb3, 0x9b, 0xf2, 0x25, 0xcb, 0x72,
	0xaa, 0x78, 0xe2, 0xa6, 0xd6, 0x48, 0x46, 0xaa, 0x92, 0x98, 0xd3, 0x48, 0xd0, 0xe4, 0xb5, 0x40,
	0xb2, 0x1b, 0x40, 0xad, 0xae, 0x96, 0x09, 0xae, 0x1a, 0xbe, 0x1b, 0x40, 0x69, 0x34, 0xe5, 0x6c,
	0x81, 0x1c, 0x75, 0x8c, 0x3a, 0xb6, 0x76, 0x36, 0x43, 0xbb, 0xd6, 0x0c, 0x92, 0x27, 0xe5, 0x9c,
	0x71, 0xb4, 0xc2, 0x3f, 0x31, 0x3c, 0xcb, 0xd8, 0x0e, 0x45, 0x3b, 0x07, 0xda, 0xa1, 0x5b, 0x69,
	0x07, 0x54, 0x05, 0xac, 0x2a, 0xe4, 0x09, 0xb4, 0xf4, 0x2f, 0xf8, 0x3d, 0x8d, 0x99, 0x24, 0xf8,
	0xe6, 0x00, 0x91, 0x5a, 0xd9, 0xae, 0xf8, 0x47, 0x72, 0x15, 0x6d, 0xd7, 0x2e, 0xb5, 0xdd, 0x63,
	0xca, 0x54, 0x69, 0xbc, 0x6e, 0xbd, 0xf1, 0x1e, 0x2a, 0xd6, 0xaf, 0x26, 0x78, 0xf6, 0x12, 0x40,
	0xa5, 0xb6, 0xda, 0xcf, 0xf2, 0x75, 0xb6, 0xf8, 0x36, 0x2d, 0x5f, 0x69, 0x5b, 0x2e, 0x22, 0xb1,
	0xca, 0x51, 0x15, 0xcc, 0xd4, 0xa8, 0x29, 0xd9, 0xae, 0x8c, 0xb4, 0x28, 0x50, 0x19, 0x2a, 0x76,
	0xdc, 0x1a, 0x79, 0x51, 0xb0, 0x32, 0xa4, 0xb8, 0xca, 0xaf, 0x71, 0x71, 0x93, 0x2e, 0x0a, 0xd1,
	0x36, 0x00, 0xf1, 0xe1, 0x84, 0x66, 0x89, 0x5e, 0x33, 0x52, 0x15, 0xa9, 0x6a, 0xd9, 0x64, 0xc5,
	0xb5, 0x62, 0x28, 0x91, 0xcd, 0xc9, 0x73, 0x18, 0x4c, 0x19, 0xff, 0x12, 0xf1, 0x84, 0x26, 0xef,
	0x14, 0x39, 0xa3, 0x55, 0x15, 0x54, 0xaa, 0x2d, 0xb9, 0xbc, 0x40, 0xb5, 0x6a, 0xce, 0xc8, 0x24,
	0xaa, 0x1e, 0x1d, 0xdc, 0x66, 0xa9, 0xf0, 0xfb, 0xa6, 0x1e, 0x0b, 0x28, 0xb5, 0x84, 0x1a, 0xf6,
	0x81, 0x51, 0x4b, 0xc5, 0x5b, 0xce, 0xbb, 0x3b, 0x9c, 0xb7, 0x0e, 0x79, 0x25, 0x87, 0x76, 0xf4,
	0xc3, 0xf0, 0xc0, 0xd8, 0xfc, 0x57, 0x1e, 0x9b, 0xe0, 0xbb, 0x74, 0xb8, 0x3e, 0x0b, 0x7f, 0xdb,
	0x61, 0xe5, 0x83, 0x0c, 0xc6, 0x34, 0x13, 0xc5, 0xd5, 0x51, 0xe4, 0x75, 0xf7, 0x4f, 0xb6, 0xdd,
	0x97, 0xa7, 0x73, 0x3a, 0x5b, 0xc8, 0xcd, 0x39, 0x1a, 0x6c, 0xf3, 0x8d, 0x3f, 0xdd, 0xbd, 0xfe,
	0xc0, 0x3e, 0x7f, 0x7a, 0xf7, 0xf8, 0xd3, 0xdf, 0xe1, 0x4f, 0x31, 0xd1, 0x83, 0xd2, 0x44, 0x5b,
	0xcf, 0xdc, 0xfb, 0x3d, 0xf3, 0x0e, 0xcf, 0xf0, 0xb0, 0x36, 0xc3, 0x17, 0xbf, 0x1b, 0x00, 0x38,
	0x9b, 0xe3, 0xbb, 0x98, 0x9c, 0x43, 0x47, 0x0a, 0x96, 0x28, 0x84, 0x78, 0x61, 0xf5, 0xe5, 0x3e,
	0x1d, 0x86, 0xb5, 0x29, 0x0e, 0x8e, 0xc8, 0x05, 0xf4, 0xd4, 0x81, 0xa2, 0x7c, 0x2f, 0xac, 0x3e,
	0x94, 0xf2, 0x4c, 0xad, 0x2f, 0xe4, 0x99, 0x4b, 0xe8, 0xe3, 0x3b, 0xf3, 0x66, 0x3d, 0x96, 0x7d,
	0xe1, 0x85, 0xd5, 0xb7, 0x49, 0x1e, 0xaa, 0xbd, 0x43, 0xf2, 0xd0, 0x2b, 0xfd, 0x38, 0xe1, 0xc7,
	0xf6, 0x9c, 0xfb, 0x3f, 0xdc, 0xbe, 0x93, 0x83, 0xa3, 0x49, 0x5b, 0xff, 0x17, 0xb9, 0xfc, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0xa9, 0xba, 0x2a, 0xa7, 0x08, 0x00, 0x00,
}