// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network_helper.proto

package network_helper

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type REQUEST_PING struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *REQUEST_PING) Reset()         { *m = REQUEST_PING{} }
func (m *REQUEST_PING) String() string { return proto.CompactTextString(m) }
func (*REQUEST_PING) ProtoMessage()    {}
func (*REQUEST_PING) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{0}
}
func (m *REQUEST_PING) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_REQUEST_PING.Unmarshal(m, b)
}
func (m *REQUEST_PING) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_REQUEST_PING.Marshal(b, m, deterministic)
}
func (dst *REQUEST_PING) XXX_Merge(src proto.Message) {
	xxx_messageInfo_REQUEST_PING.Merge(dst, src)
}
func (m *REQUEST_PING) XXX_Size() int {
	return xxx_messageInfo_REQUEST_PING.Size(m)
}
func (m *REQUEST_PING) XXX_DiscardUnknown() {
	xxx_messageInfo_REQUEST_PING.DiscardUnknown(m)
}

var xxx_messageInfo_REQUEST_PING proto.InternalMessageInfo

func (m *REQUEST_PING) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type REQUEST_CONTACT struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *REQUEST_CONTACT) Reset()         { *m = REQUEST_CONTACT{} }
func (m *REQUEST_CONTACT) String() string { return proto.CompactTextString(m) }
func (*REQUEST_CONTACT) ProtoMessage()    {}
func (*REQUEST_CONTACT) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{1}
}
func (m *REQUEST_CONTACT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_REQUEST_CONTACT.Unmarshal(m, b)
}
func (m *REQUEST_CONTACT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_REQUEST_CONTACT.Marshal(b, m, deterministic)
}
func (dst *REQUEST_CONTACT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_REQUEST_CONTACT.Merge(dst, src)
}
func (m *REQUEST_CONTACT) XXX_Size() int {
	return xxx_messageInfo_REQUEST_CONTACT.Size(m)
}
func (m *REQUEST_CONTACT) XXX_DiscardUnknown() {
	xxx_messageInfo_REQUEST_CONTACT.DiscardUnknown(m)
}

var xxx_messageInfo_REQUEST_CONTACT proto.InternalMessageInfo

func (m *REQUEST_CONTACT) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type REQUEST_DATA struct {
	KEY                  string   `protobuf:"bytes,1,opt,name=KEY,proto3" json:"KEY,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *REQUEST_DATA) Reset()         { *m = REQUEST_DATA{} }
func (m *REQUEST_DATA) String() string { return proto.CompactTextString(m) }
func (*REQUEST_DATA) ProtoMessage()    {}
func (*REQUEST_DATA) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{2}
}
func (m *REQUEST_DATA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_REQUEST_DATA.Unmarshal(m, b)
}
func (m *REQUEST_DATA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_REQUEST_DATA.Marshal(b, m, deterministic)
}
func (dst *REQUEST_DATA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_REQUEST_DATA.Merge(dst, src)
}
func (m *REQUEST_DATA) XXX_Size() int {
	return xxx_messageInfo_REQUEST_DATA.Size(m)
}
func (m *REQUEST_DATA) XXX_DiscardUnknown() {
	xxx_messageInfo_REQUEST_DATA.DiscardUnknown(m)
}

var xxx_messageInfo_REQUEST_DATA proto.InternalMessageInfo

func (m *REQUEST_DATA) GetKEY() string {
	if m != nil {
		return m.KEY
	}
	return ""
}

type REQUEST_STORE struct {
	KEY                  string   `protobuf:"bytes,1,opt,name=KEY,proto3" json:"KEY,omitempty"`
	VALUE                []byte   `protobuf:"bytes,2,opt,name=VALUE,proto3" json:"VALUE,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *REQUEST_STORE) Reset()         { *m = REQUEST_STORE{} }
func (m *REQUEST_STORE) String() string { return proto.CompactTextString(m) }
func (*REQUEST_STORE) ProtoMessage()    {}
func (*REQUEST_STORE) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{3}
}
func (m *REQUEST_STORE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_REQUEST_STORE.Unmarshal(m, b)
}
func (m *REQUEST_STORE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_REQUEST_STORE.Marshal(b, m, deterministic)
}
func (dst *REQUEST_STORE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_REQUEST_STORE.Merge(dst, src)
}
func (m *REQUEST_STORE) XXX_Size() int {
	return xxx_messageInfo_REQUEST_STORE.Size(m)
}
func (m *REQUEST_STORE) XXX_DiscardUnknown() {
	xxx_messageInfo_REQUEST_STORE.DiscardUnknown(m)
}

var xxx_messageInfo_REQUEST_STORE proto.InternalMessageInfo

func (m *REQUEST_STORE) GetKEY() string {
	if m != nil {
		return m.KEY
	}
	return ""
}

func (m *REQUEST_STORE) GetVALUE() []byte {
	if m != nil {
		return m.VALUE
	}
	return nil
}

type RETURN_PING struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RETURN_PING) Reset()         { *m = RETURN_PING{} }
func (m *RETURN_PING) String() string { return proto.CompactTextString(m) }
func (*RETURN_PING) ProtoMessage()    {}
func (*RETURN_PING) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{4}
}
func (m *RETURN_PING) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RETURN_PING.Unmarshal(m, b)
}
func (m *RETURN_PING) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RETURN_PING.Marshal(b, m, deterministic)
}
func (dst *RETURN_PING) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RETURN_PING.Merge(dst, src)
}
func (m *RETURN_PING) XXX_Size() int {
	return xxx_messageInfo_RETURN_PING.Size(m)
}
func (m *RETURN_PING) XXX_DiscardUnknown() {
	xxx_messageInfo_RETURN_PING.DiscardUnknown(m)
}

var xxx_messageInfo_RETURN_PING proto.InternalMessageInfo

func (m *RETURN_PING) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RETURN_CONTACTS struct {
	ContactInfo          []*RETURN_CONTACTS_CONTACT_INFO `protobuf:"bytes,1,rep,name=contact_info,json=contactInfo,proto3" json:"contact_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RETURN_CONTACTS) Reset()         { *m = RETURN_CONTACTS{} }
func (m *RETURN_CONTACTS) String() string { return proto.CompactTextString(m) }
func (*RETURN_CONTACTS) ProtoMessage()    {}
func (*RETURN_CONTACTS) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{5}
}
func (m *RETURN_CONTACTS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RETURN_CONTACTS.Unmarshal(m, b)
}
func (m *RETURN_CONTACTS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RETURN_CONTACTS.Marshal(b, m, deterministic)
}
func (dst *RETURN_CONTACTS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RETURN_CONTACTS.Merge(dst, src)
}
func (m *RETURN_CONTACTS) XXX_Size() int {
	return xxx_messageInfo_RETURN_CONTACTS.Size(m)
}
func (m *RETURN_CONTACTS) XXX_DiscardUnknown() {
	xxx_messageInfo_RETURN_CONTACTS.DiscardUnknown(m)
}

var xxx_messageInfo_RETURN_CONTACTS proto.InternalMessageInfo

func (m *RETURN_CONTACTS) GetContactInfo() []*RETURN_CONTACTS_CONTACT_INFO {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

type RETURN_CONTACTS_CONTACT_INFO struct {
	IP                   string   `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	PORT                 string   `protobuf:"bytes,2,opt,name=PORT,proto3" json:"PORT,omitempty"`
	ID                   string   `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RETURN_CONTACTS_CONTACT_INFO) Reset()         { *m = RETURN_CONTACTS_CONTACT_INFO{} }
func (m *RETURN_CONTACTS_CONTACT_INFO) String() string { return proto.CompactTextString(m) }
func (*RETURN_CONTACTS_CONTACT_INFO) ProtoMessage()    {}
func (*RETURN_CONTACTS_CONTACT_INFO) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{5, 0}
}
func (m *RETURN_CONTACTS_CONTACT_INFO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RETURN_CONTACTS_CONTACT_INFO.Unmarshal(m, b)
}
func (m *RETURN_CONTACTS_CONTACT_INFO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RETURN_CONTACTS_CONTACT_INFO.Marshal(b, m, deterministic)
}
func (dst *RETURN_CONTACTS_CONTACT_INFO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RETURN_CONTACTS_CONTACT_INFO.Merge(dst, src)
}
func (m *RETURN_CONTACTS_CONTACT_INFO) XXX_Size() int {
	return xxx_messageInfo_RETURN_CONTACTS_CONTACT_INFO.Size(m)
}
func (m *RETURN_CONTACTS_CONTACT_INFO) XXX_DiscardUnknown() {
	xxx_messageInfo_RETURN_CONTACTS_CONTACT_INFO.DiscardUnknown(m)
}

var xxx_messageInfo_RETURN_CONTACTS_CONTACT_INFO proto.InternalMessageInfo

func (m *RETURN_CONTACTS_CONTACT_INFO) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *RETURN_CONTACTS_CONTACT_INFO) GetPORT() string {
	if m != nil {
		return m.PORT
	}
	return ""
}

func (m *RETURN_CONTACTS_CONTACT_INFO) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RETURN_DATA struct {
	VALUE                []byte   `protobuf:"bytes,1,opt,name=VALUE,proto3" json:"VALUE,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RETURN_DATA) Reset()         { *m = RETURN_DATA{} }
func (m *RETURN_DATA) String() string { return proto.CompactTextString(m) }
func (*RETURN_DATA) ProtoMessage()    {}
func (*RETURN_DATA) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{6}
}
func (m *RETURN_DATA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RETURN_DATA.Unmarshal(m, b)
}
func (m *RETURN_DATA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RETURN_DATA.Marshal(b, m, deterministic)
}
func (dst *RETURN_DATA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RETURN_DATA.Merge(dst, src)
}
func (m *RETURN_DATA) XXX_Size() int {
	return xxx_messageInfo_RETURN_DATA.Size(m)
}
func (m *RETURN_DATA) XXX_DiscardUnknown() {
	xxx_messageInfo_RETURN_DATA.DiscardUnknown(m)
}

var xxx_messageInfo_RETURN_DATA proto.InternalMessageInfo

func (m *RETURN_DATA) GetVALUE() []byte {
	if m != nil {
		return m.VALUE
	}
	return nil
}

type RETURN_STORE struct {
	VALUE                string   `protobuf:"bytes,1,opt,name=VALUE,proto3" json:"VALUE,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RETURN_STORE) Reset()         { *m = RETURN_STORE{} }
func (m *RETURN_STORE) String() string { return proto.CompactTextString(m) }
func (*RETURN_STORE) ProtoMessage()    {}
func (*RETURN_STORE) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{7}
}
func (m *RETURN_STORE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RETURN_STORE.Unmarshal(m, b)
}
func (m *RETURN_STORE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RETURN_STORE.Marshal(b, m, deterministic)
}
func (dst *RETURN_STORE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RETURN_STORE.Merge(dst, src)
}
func (m *RETURN_STORE) XXX_Size() int {
	return xxx_messageInfo_RETURN_STORE.Size(m)
}
func (m *RETURN_STORE) XXX_DiscardUnknown() {
	xxx_messageInfo_RETURN_STORE.DiscardUnknown(m)
}

var xxx_messageInfo_RETURN_STORE proto.InternalMessageInfo

func (m *RETURN_STORE) GetVALUE() string {
	if m != nil {
		return m.VALUE
	}
	return ""
}

type Container struct {
	REQUEST_TYPE string `protobuf:"bytes,1,opt,name=REQUEST_TYPE,json=REQUESTTYPE,proto3" json:"REQUEST_TYPE,omitempty"`
	REQUEST_ID   string `protobuf:"bytes,2,opt,name=REQUEST_ID,json=REQUESTID,proto3" json:"REQUEST_ID,omitempty"`
	MSG_ID       string `protobuf:"bytes,3,opt,name=MSG_ID,json=MSGID,proto3" json:"MSG_ID,omitempty"`
	ID           string `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
	PORT         string `protobuf:"bytes,5,opt,name=PORT,proto3" json:"PORT,omitempty"`
	// Types that are valid to be assigned to Attachment:
	//	*Container_RequestPing
	//	*Container_RequestContact
	//	*Container_RequestData
	//	*Container_RequestStore
	//	*Container_ReturnPing
	//	*Container_ReturnContacts
	//	*Container_ReturnData
	//	*Container_ReturnStore
	Attachment           isContainer_Attachment `protobuf_oneof:"attachment"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_helper_f5fdaf30b5ca8af1, []int{8}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetREQUEST_TYPE() string {
	if m != nil {
		return m.REQUEST_TYPE
	}
	return ""
}

func (m *Container) GetREQUEST_ID() string {
	if m != nil {
		return m.REQUEST_ID
	}
	return ""
}

func (m *Container) GetMSG_ID() string {
	if m != nil {
		return m.MSG_ID
	}
	return ""
}

func (m *Container) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Container) GetPORT() string {
	if m != nil {
		return m.PORT
	}
	return ""
}

type isContainer_Attachment interface {
	isContainer_Attachment()
}

type Container_RequestPing struct {
	RequestPing *REQUEST_PING `protobuf:"bytes,6,opt,name=request_ping,json=requestPing,proto3,oneof"`
}

type Container_RequestContact struct {
	RequestContact *REQUEST_CONTACT `protobuf:"bytes,7,opt,name=request_contact,json=requestContact,proto3,oneof"`
}

type Container_RequestData struct {
	RequestData *REQUEST_DATA `protobuf:"bytes,8,opt,name=request_data,json=requestData,proto3,oneof"`
}

type Container_RequestStore struct {
	RequestStore *REQUEST_STORE `protobuf:"bytes,9,opt,name=request_store,json=requestStore,proto3,oneof"`
}

type Container_ReturnPing struct {
	ReturnPing *RETURN_PING `protobuf:"bytes,10,opt,name=return_ping,json=returnPing,proto3,oneof"`
}

type Container_ReturnContacts struct {
	ReturnContacts *RETURN_CONTACTS `protobuf:"bytes,11,opt,name=return_contacts,json=returnContacts,proto3,oneof"`
}

type Container_ReturnData struct {
	ReturnData *RETURN_DATA `protobuf:"bytes,12,opt,name=return_data,json=returnData,proto3,oneof"`
}

type Container_ReturnStore struct {
	ReturnStore *RETURN_STORE `protobuf:"bytes,13,opt,name=return_store,json=returnStore,proto3,oneof"`
}

func (*Container_RequestPing) isContainer_Attachment() {}

func (*Container_RequestContact) isContainer_Attachment() {}

func (*Container_RequestData) isContainer_Attachment() {}

func (*Container_RequestStore) isContainer_Attachment() {}

func (*Container_ReturnPing) isContainer_Attachment() {}

func (*Container_ReturnContacts) isContainer_Attachment() {}

func (*Container_ReturnData) isContainer_Attachment() {}

func (*Container_ReturnStore) isContainer_Attachment() {}

func (m *Container) GetAttachment() isContainer_Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *Container) GetRequestPing() *REQUEST_PING {
	if x, ok := m.GetAttachment().(*Container_RequestPing); ok {
		return x.RequestPing
	}
	return nil
}

func (m *Container) GetRequestContact() *REQUEST_CONTACT {
	if x, ok := m.GetAttachment().(*Container_RequestContact); ok {
		return x.RequestContact
	}
	return nil
}

func (m *Container) GetRequestData() *REQUEST_DATA {
	if x, ok := m.GetAttachment().(*Container_RequestData); ok {
		return x.RequestData
	}
	return nil
}

func (m *Container) GetRequestStore() *REQUEST_STORE {
	if x, ok := m.GetAttachment().(*Container_RequestStore); ok {
		return x.RequestStore
	}
	return nil
}

func (m *Container) GetReturnPing() *RETURN_PING {
	if x, ok := m.GetAttachment().(*Container_ReturnPing); ok {
		return x.ReturnPing
	}
	return nil
}

func (m *Container) GetReturnContacts() *RETURN_CONTACTS {
	if x, ok := m.GetAttachment().(*Container_ReturnContacts); ok {
		return x.ReturnContacts
	}
	return nil
}

func (m *Container) GetReturnData() *RETURN_DATA {
	if x, ok := m.GetAttachment().(*Container_ReturnData); ok {
		return x.ReturnData
	}
	return nil
}

func (m *Container) GetReturnStore() *RETURN_STORE {
	if x, ok := m.GetAttachment().(*Container_ReturnStore); ok {
		return x.ReturnStore
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Container) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Container_OneofMarshaler, _Container_OneofUnmarshaler, _Container_OneofSizer, []interface{}{
		(*Container_RequestPing)(nil),
		(*Container_RequestContact)(nil),
		(*Container_RequestData)(nil),
		(*Container_RequestStore)(nil),
		(*Container_ReturnPing)(nil),
		(*Container_ReturnContacts)(nil),
		(*Container_ReturnData)(nil),
		(*Container_ReturnStore)(nil),
	}
}

func _Container_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Container)
	// attachment
	switch x := m.Attachment.(type) {
	case *Container_RequestPing:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestPing); err != nil {
			return err
		}
	case *Container_RequestContact:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestContact); err != nil {
			return err
		}
	case *Container_RequestData:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestData); err != nil {
			return err
		}
	case *Container_RequestStore:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestStore); err != nil {
			return err
		}
	case *Container_ReturnPing:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnPing); err != nil {
			return err
		}
	case *Container_ReturnContacts:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnContacts); err != nil {
			return err
		}
	case *Container_ReturnData:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnData); err != nil {
			return err
		}
	case *Container_ReturnStore:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnStore); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Container.Attachment has unexpected type %T", x)
	}
	return nil
}

func _Container_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Container)
	switch tag {
	case 6: // attachment.request_ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(REQUEST_PING)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_RequestPing{msg}
		return true, err
	case 7: // attachment.request_contact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(REQUEST_CONTACT)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_RequestContact{msg}
		return true, err
	case 8: // attachment.request_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(REQUEST_DATA)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_RequestData{msg}
		return true, err
	case 9: // attachment.request_store
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(REQUEST_STORE)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_RequestStore{msg}
		return true, err
	case 10: // attachment.return_ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RETURN_PING)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_ReturnPing{msg}
		return true, err
	case 11: // attachment.return_contacts
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RETURN_CONTACTS)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_ReturnContacts{msg}
		return true, err
	case 12: // attachment.return_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RETURN_DATA)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_ReturnData{msg}
		return true, err
	case 13: // attachment.return_store
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RETURN_STORE)
		err := b.DecodeMessage(msg)
		m.Attachment = &Container_ReturnStore{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Container_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Container)
	// attachment
	switch x := m.Attachment.(type) {
	case *Container_RequestPing:
		s := proto.Size(x.RequestPing)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_RequestContact:
		s := proto.Size(x.RequestContact)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_RequestData:
		s := proto.Size(x.RequestData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_RequestStore:
		s := proto.Size(x.RequestStore)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_ReturnPing:
		s := proto.Size(x.ReturnPing)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_ReturnContacts:
		s := proto.Size(x.ReturnContacts)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_ReturnData:
		s := proto.Size(x.ReturnData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Container_ReturnStore:
		s := proto.Size(x.ReturnStore)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*REQUEST_PING)(nil), "network_helper.REQUEST_PING")
	proto.RegisterType((*REQUEST_CONTACT)(nil), "network_helper.REQUEST_CONTACT")
	proto.RegisterType((*REQUEST_DATA)(nil), "network_helper.REQUEST_DATA")
	proto.RegisterType((*REQUEST_STORE)(nil), "network_helper.REQUEST_STORE")
	proto.RegisterType((*RETURN_PING)(nil), "network_helper.RETURN_PING")
	proto.RegisterType((*RETURN_CONTACTS)(nil), "network_helper.RETURN_CONTACTS")
	proto.RegisterType((*RETURN_CONTACTS_CONTACT_INFO)(nil), "network_helper.RETURN_CONTACTS.CONTACT_INFO")
	proto.RegisterType((*RETURN_DATA)(nil), "network_helper.RETURN_DATA")
	proto.RegisterType((*RETURN_STORE)(nil), "network_helper.RETURN_STORE")
	proto.RegisterType((*Container)(nil), "network_helper.Container")
}

func init() {
	proto.RegisterFile("network_helper.proto", fileDescriptor_network_helper_f5fdaf30b5ca8af1)
}

var fileDescriptor_network_helper_f5fdaf30b5ca8af1 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x9b, 0x76, 0x2d, 0xe4, 0x9f, 0xb4, 0x43, 0xd6, 0x90, 0x22, 0xa0, 0xd0, 0x15, 0x0e,
	0x3d, 0xa0, 0x1e, 0xc6, 0x81, 0x1b, 0x52, 0xd6, 0x84, 0x36, 0xc0, 0x9a, 0xe0, 0xa4, 0x48, 0x3b,
	0x45, 0xa1, 0x64, 0x5b, 0x05, 0x38, 0x25, 0xf5, 0xc4, 0x37, 0xe2, 0x5b, 0xf0, 0xdd, 0x90, 0x9d,
	0xbf, 0x93, 0x74, 0x25, 0xbb, 0xb9, 0xf6, 0x7b, 0xaf, 0xef, 0xff, 0xb3, 0x15, 0x38, 0x61, 0x29,
	0xff, 0x9d, 0xe5, 0xdf, 0xe3, 0x9b, 0xf4, 0xc7, 0x36, 0xcd, 0xa7, 0xdb, 0x3c, 0xe3, 0x19, 0x19,
	0xec, 0xef, 0x8e, 0x9f, 0x83, 0x49, 0xdd, 0xcf, 0x2b, 0x37, 0x8c, 0xe2, 0xc0, 0x5b, 0xce, 0xc9,
	0x00, 0xda, 0x9e, 0x63, 0x69, 0x23, 0x6d, 0xa2, 0xd3, 0xb6, 0xe7, 0x8c, 0x4f, 0xe1, 0x58, 0x9d,
	0xcf, 0xfc, 0x65, 0x64, 0xcf, 0xa2, 0x03, 0xc9, 0xa8, 0x8a, 0x70, 0xec, 0xc8, 0x26, 0x8f, 0xa0,
	0xf3, 0xd1, 0xbd, 0x44, 0x81, 0x58, 0x8e, 0xdf, 0x42, 0x5f, 0x29, 0xc2, 0xc8, 0xa7, 0xee, 0xa1,
	0x84, 0x9c, 0x40, 0xf7, 0x8b, 0xfd, 0x69, 0xe5, 0x5a, 0xed, 0x91, 0x36, 0x31, 0x69, 0xf1, 0x63,
	0x3c, 0x04, 0x83, 0xba, 0xd1, 0x8a, 0x2e, 0xff, 0x5f, 0xee, 0x8f, 0x26, 0xda, 0xc9, 0x73, 0x2c,
	0x17, 0x12, 0x1f, 0xcc, 0x75, 0xc6, 0x78, 0xb2, 0xe6, 0xf1, 0x86, 0x5d, 0x65, 0x96, 0x36, 0xea,
	0x4c, 0x8c, 0xb3, 0xd7, 0xd3, 0x3b, 0x34, 0xee, 0xd8, 0xa6, 0xb8, 0x88, 0xbd, 0xe5, 0x7b, 0x9f,
	0x1a, 0x98, 0xe0, 0xb1, 0xab, 0xec, 0xc9, 0x39, 0x98, 0xf5, 0x43, 0x59, 0x22, 0x28, 0x4b, 0x04,
	0x84, 0xc0, 0x51, 0xe0, 0xd3, 0x48, 0x16, 0xd7, 0xa9, 0x5c, 0x63, 0xd1, 0x4e, 0x59, 0xf4, 0x65,
	0x39, 0x87, 0x24, 0x54, 0x0e, 0xab, 0xd5, 0x87, 0x7d, 0x25, 0x38, 0x4a, 0x51, 0x01, 0x69, 0x4f,
	0xa5, 0x2b, 0xd5, 0xdf, 0x2e, 0xe8, 0x33, 0x51, 0x6f, 0xc3, 0xd2, 0x9c, 0x9c, 0x56, 0xec, 0xa3,
	0xcb, 0x40, 0x49, 0x0d, 0xdc, 0x13, 0x5b, 0x64, 0x08, 0xa0, 0x24, 0x9e, 0x83, 0x2d, 0x75, 0xdc,
	0xf1, 0x1c, 0xf2, 0x18, 0x7a, 0x17, 0xe1, 0x3c, 0x2e, 0xeb, 0x76, 0x2f, 0xc2, 0xb9, 0xe7, 0xe0,
	0x04, 0x47, 0x6a, 0x82, 0x72, 0xca, 0x6e, 0x6d, 0x4a, 0x1b, 0xcc, 0x3c, 0xfd, 0x75, 0x9b, 0xee,
	0x78, 0xbc, 0xdd, 0xb0, 0x6b, 0xab, 0x37, 0xd2, 0x26, 0xc6, 0xd9, 0xb3, 0x43, 0xd4, 0xd5, 0xfb,
	0x5a, 0xb4, 0xa8, 0x81, 0x9e, 0x60, 0xc3, 0xae, 0xc9, 0x07, 0x38, 0x56, 0x11, 0xc8, 0xdc, 0x7a,
	0x20, 0x53, 0x5e, 0x34, 0xa5, 0xe0, 0x5d, 0x2c, 0x5a, 0x74, 0x80, 0xce, 0x59, 0x61, 0xac, 0xd7,
	0xf9, 0x96, 0xf0, 0xc4, 0x7a, 0x78, 0x7f, 0x1d, 0x71, 0x13, 0xb5, 0x3a, 0x4e, 0xc2, 0x13, 0xe2,
	0x40, 0x5f, 0x45, 0xec, 0x78, 0x96, 0xa7, 0x96, 0x2e, 0x33, 0x86, 0x4d, 0x19, 0xf2, 0xa2, 0x16,
	0x2d, 0xaa, 0xfe, 0x38, 0x14, 0x26, 0xf2, 0x0e, 0x8c, 0x3c, 0xe5, 0xb7, 0x39, 0x2b, 0xb0, 0x80,
	0xcc, 0x78, 0xda, 0xf0, 0x02, 0x91, 0x0a, 0x14, 0x8e, 0x0a, 0x8a, 0xf4, 0x23, 0x93, 0x9d, 0x65,
	0x34, 0x41, 0xd9, 0x7b, 0xc5, 0x05, 0x14, 0xe1, 0x44, 0x26, 0xbb, 0x5a, 0x17, 0xc9, 0xc4, 0xbc,
	0xb7, 0x0b, 0x22, 0xc1, 0x2e, 0x92, 0x88, 0x84, 0x2a, 0xfd, 0x05, 0x90, 0x7e, 0x13, 0xd4, 0xea,
	0xe1, 0x16, 0x50, 0x85, 0x47, 0xe2, 0x38, 0x37, 0x01, 0x12, 0xce, 0x93, 0xf5, 0xcd, 0xcf, 0x94,
	0xf1, 0xaf, 0x3d, 0xf9, 0x1d, 0x7a, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x56, 0x1d, 0x79, 0x54,
	0x9f, 0x04, 0x00, 0x00,
}
