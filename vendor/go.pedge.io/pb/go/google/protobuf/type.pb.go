// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/type.proto

package google_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The syntax in which a protocol buffer element is defined.
type Syntax int32

const (
	// Syntax `proto2`.
	Syntax_SYNTAX_PROTO2 Syntax = 0
	// Syntax `proto3`.
	Syntax_SYNTAX_PROTO3 Syntax = 1
)

var Syntax_name = map[int32]string{
	0: "SYNTAX_PROTO2",
	1: "SYNTAX_PROTO3",
}
var Syntax_value = map[string]int32{
	"SYNTAX_PROTO2": 0,
	"SYNTAX_PROTO3": 1,
}

func (x Syntax) String() string {
	return proto.EnumName(Syntax_name, int32(x))
}
func (Syntax) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// Basic field types.
type Field_Kind int32

const (
	// Field type unknown.
	Field_TYPE_UNKNOWN Field_Kind = 0
	// Field type double.
	Field_TYPE_DOUBLE Field_Kind = 1
	// Field type float.
	Field_TYPE_FLOAT Field_Kind = 2
	// Field type int64.
	Field_TYPE_INT64 Field_Kind = 3
	// Field type uint64.
	Field_TYPE_UINT64 Field_Kind = 4
	// Field type int32.
	Field_TYPE_INT32 Field_Kind = 5
	// Field type fixed64.
	Field_TYPE_FIXED64 Field_Kind = 6
	// Field type fixed32.
	Field_TYPE_FIXED32 Field_Kind = 7
	// Field type bool.
	Field_TYPE_BOOL Field_Kind = 8
	// Field type string.
	Field_TYPE_STRING Field_Kind = 9
	// Field type group. Proto2 syntax only, and deprecated.
	Field_TYPE_GROUP Field_Kind = 10
	// Field type message.
	Field_TYPE_MESSAGE Field_Kind = 11
	// Field type bytes.
	Field_TYPE_BYTES Field_Kind = 12
	// Field type uint32.
	Field_TYPE_UINT32 Field_Kind = 13
	// Field type enum.
	Field_TYPE_ENUM Field_Kind = 14
	// Field type sfixed32.
	Field_TYPE_SFIXED32 Field_Kind = 15
	// Field type sfixed64.
	Field_TYPE_SFIXED64 Field_Kind = 16
	// Field type sint32.
	Field_TYPE_SINT32 Field_Kind = 17
	// Field type sint64.
	Field_TYPE_SINT64 Field_Kind = 18
)

var Field_Kind_name = map[int32]string{
	0:  "TYPE_UNKNOWN",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var Field_Kind_value = map[string]int32{
	"TYPE_UNKNOWN":  0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x Field_Kind) String() string {
	return proto.EnumName(Field_Kind_name, int32(x))
}
func (Field_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

// Whether a field is optional, required, or repeated.
type Field_Cardinality int32

const (
	// For fields with unknown cardinality.
	Field_CARDINALITY_UNKNOWN Field_Cardinality = 0
	// For optional fields.
	Field_CARDINALITY_OPTIONAL Field_Cardinality = 1
	// For required fields. Proto2 syntax only.
	Field_CARDINALITY_REQUIRED Field_Cardinality = 2
	// For repeated fields.
	Field_CARDINALITY_REPEATED Field_Cardinality = 3
)

var Field_Cardinality_name = map[int32]string{
	0: "CARDINALITY_UNKNOWN",
	1: "CARDINALITY_OPTIONAL",
	2: "CARDINALITY_REQUIRED",
	3: "CARDINALITY_REPEATED",
}
var Field_Cardinality_value = map[string]int32{
	"CARDINALITY_UNKNOWN":  0,
	"CARDINALITY_OPTIONAL": 1,
	"CARDINALITY_REQUIRED": 2,
	"CARDINALITY_REPEATED": 3,
}

func (x Field_Cardinality) String() string {
	return proto.EnumName(Field_Cardinality_name, int32(x))
}
func (Field_Cardinality) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 1} }

// A protocol buffer message type.
type Type struct {
	// The fully qualified message name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The list of fields.
	Fields []*Field `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	// The list of types appearing in `oneof` definitions in this type.
	Oneofs []string `protobuf:"bytes,3,rep,name=oneofs" json:"oneofs,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// The source syntax.
	Syntax Syntax `protobuf:"varint,6,opt,name=syntax,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetOneofs() []string {
	if m != nil {
		return m.Oneofs
	}
	return nil
}

func (m *Type) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Type) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Type) GetSyntax() Syntax {
	if m != nil {
		return m.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// A single field of a message type.
type Field struct {
	// The field type.
	Kind Field_Kind `protobuf:"varint,1,opt,name=kind,enum=google.protobuf.Field_Kind" json:"kind,omitempty"`
	// The field cardinality.
	Cardinality Field_Cardinality `protobuf:"varint,2,opt,name=cardinality,enum=google.protobuf.Field_Cardinality" json:"cardinality,omitempty"`
	// The field number.
	Number int32 `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	// The field name.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The field type URL, without the scheme, for message or enumeration
	// types. Example: `"type.googleapis.com/google.protobuf.Timestamp"`.
	TypeUrl string `protobuf:"bytes,6,opt,name=type_url,json=typeUrl" json:"type_url,omitempty"`
	// The index of the field type in `Type.oneofs`, for message or enumeration
	// types. The first type has index 1; zero means the type is not in the list.
	OneofIndex int32 `protobuf:"varint,7,opt,name=oneof_index,json=oneofIndex" json:"oneof_index,omitempty"`
	// Whether to use alternative packed wire representation.
	Packed bool `protobuf:"varint,8,opt,name=packed" json:"packed,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,9,rep,name=options" json:"options,omitempty"`
	// The field JSON name.
	JsonName string `protobuf:"bytes,10,opt,name=json_name,json=jsonName" json:"json_name,omitempty"`
	// The string value of the default value of this field. Proto2 syntax only.
	DefaultValue string `protobuf:"bytes,11,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Field) GetKind() Field_Kind {
	if m != nil {
		return m.Kind
	}
	return Field_TYPE_UNKNOWN
}

func (m *Field) GetCardinality() Field_Cardinality {
	if m != nil {
		return m.Cardinality
	}
	return Field_CARDINALITY_UNKNOWN
}

func (m *Field) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *Field) GetOneofIndex() int32 {
	if m != nil {
		return m.OneofIndex
	}
	return 0
}

func (m *Field) GetPacked() bool {
	if m != nil {
		return m.Packed
	}
	return false
}

func (m *Field) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Field) GetJsonName() string {
	if m != nil {
		return m.JsonName
	}
	return ""
}

func (m *Field) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

// Enum type definition.
type Enum struct {
	// Enum type name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Enum value definitions.
	Enumvalue []*EnumValue `protobuf:"bytes,2,rep,name=enumvalue" json:"enumvalue,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,4,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// The source syntax.
	Syntax Syntax `protobuf:"varint,5,opt,name=syntax,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Enum) Reset()                    { *m = Enum{} }
func (m *Enum) String() string            { return proto.CompactTextString(m) }
func (*Enum) ProtoMessage()               {}
func (*Enum) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Enum) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Enum) GetEnumvalue() []*EnumValue {
	if m != nil {
		return m.Enumvalue
	}
	return nil
}

func (m *Enum) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Enum) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Enum) GetSyntax() Syntax {
	if m != nil {
		return m.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// Enum value definition.
type EnumValue struct {
	// Enum value name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Enum value number.
	Number int32 `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
}

func (m *EnumValue) Reset()                    { *m = EnumValue{} }
func (m *EnumValue) String() string            { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()               {}
func (*EnumValue) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *EnumValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnumValue) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *EnumValue) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// A protocol buffer option, which can be attached to a message, field,
// enumeration, etc.
type Option struct {
	// The option's name. For protobuf built-in options (options defined in
	// descriptor.proto), this is the short name. For example, `"map_entry"`.
	// For custom options, it should be the fully-qualified name. For example,
	// `"google.api.http"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The option's value packed in an Any message. If the value is a primitive,
	// the corresponding wrapper type defined in google/protobuf/wrappers.proto
	// should be used. If the value is an enum, it should be stored as an int32
	// value using the google.protobuf.Int32Value type.
	Value *google_protobuf1.Any `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Option) Reset()                    { *m = Option{} }
func (m *Option) String() string            { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()               {}
func (*Option) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Option) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Option) GetValue() *google_protobuf1.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "google.protobuf.Type")
	proto.RegisterType((*Field)(nil), "google.protobuf.Field")
	proto.RegisterType((*Enum)(nil), "google.protobuf.Enum")
	proto.RegisterType((*EnumValue)(nil), "google.protobuf.EnumValue")
	proto.RegisterType((*Option)(nil), "google.protobuf.Option")
	proto.RegisterEnum("google.protobuf.Syntax", Syntax_name, Syntax_value)
	proto.RegisterEnum("google.protobuf.Field_Kind", Field_Kind_name, Field_Kind_value)
	proto.RegisterEnum("google.protobuf.Field_Cardinality", Field_Cardinality_name, Field_Cardinality_value)
}

func init() { proto.RegisterFile("google/protobuf/type.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x8e, 0xda, 0x56,
	0x14, 0x8e, 0x7f, 0xf0, 0xe0, 0xe3, 0x81, 0xb9, 0xb9, 0x89, 0x12, 0x67, 0x22, 0xa5, 0x88, 0x76,
	0x81, 0xb2, 0x60, 0x54, 0x18, 0x8d, 0xba, 0x35, 0x83, 0x87, 0x5a, 0x43, 0x6c, 0xf7, 0x62, 0x9a,
	0xb0, 0x42, 0x1e, 0x30, 0x11, 0x89, 0xb9, 0x46, 0xd8, 0x6e, 0x87, 0x87, 0xe8, 0x4b, 0x74, 0xd9,
	0x75, 0x1f, 0xa2, 0x8f, 0xd4, 0x5d, 0xab, 0x7b, 0x0d, 0xc6, 0xfc, 0x54, 0x4a, 0xdb, 0x1d, 0xe7,
	0x3b, 0xdf, 0xf9, 0xce, 0xcf, 0x3d, 0x3e, 0xc0, 0xe5, 0xc7, 0x28, 0xfa, 0x18, 0x06, 0x57, 0xcb,
	0x55, 0x94, 0x44, 0x0f, 0xe9, 0xec, 0x2a, 0x59, 0x2f, 0x83, 0x26, 0xb7, 0xf0, 0x45, 0xe6, 0x6b,
	0x6e, 0x7d, 0x97, 0xaf, 0x0e, 0xc9, 0x3e, 0x5d, 0x67, 0xde, 0xcb, 0x6f, 0x0e, 0x5d, 0x71, 0x94,
	0xae, 0x26, 0xc1, 0x78, 0x12, 0xd1, 0x24, 0x78, 0x4c, 0x32, 0x56, 0xfd, 0x17, 0x11, 0x64, 0x6f,
	0xbd, 0x0c, 0x30, 0x06, 0x99, 0xfa, 0x8b, 0x40, 0x17, 0x6a, 0x42, 0x43, 0x25, 0xfc, 0x37, 0x6e,
	0x82, 0x32, 0x9b, 0x07, 0xe1, 0x34, 0xd6, 0xc5, 0x9a, 0xd4, 0xd0, 0x5a, 0x2f, 0x9a, 0x07, 0xf9,
	0x9b, 0x77, 0xcc, 0x4d, 0x36, 0x2c, 0xfc, 0x02, 0x94, 0x88, 0x06, 0xd1, 0x2c, 0xd6, 0xa5, 0x9a,
	0xd4, 0x50, 0xc9, 0xc6, 0xc2, 0xdf, 0xc2, 0x59, 0xb4, 0x4c, 0xe6, 0x11, 0x8d, 0x75, 0x99, 0x0b,
	0xbd, 0x3c, 0x12, 0x72, 0xb8, 0x9f, 0x6c, 0x79, 0xd8, 0x84, 0xea, 0x7e, 0xbd, 0x7a, 0xa9, 0x26,
	0x34, 0xb4, 0xd6, 0x9b, 0xa3, 0xc8, 0x01, 0xa7, 0xdd, 0x66, 0x2c, 0x52, 0x89, 0x8b, 0x26, 0xbe,
	0x02, 0x25, 0x5e, 0xd3, 0xc4, 0x7f, 0xd4, 0x95, 0x9a, 0xd0, 0xa8, 0x9e, 0x48, 0x3c, 0xe0, 0x6e,
	0xb2, 0xa1, 0xd5, 0x7f, 0x57, 0xa0, 0xc4, 0x9b, 0xc2, 0x57, 0x20, 0x7f, 0x9e, 0xd3, 0x29, 0x1f,
	0x48, 0xb5, 0xf5, 0xfa, 0x74, 0xeb, 0xcd, 0xfb, 0x39, 0x9d, 0x12, 0x4e, 0xc4, 0x5d, 0xd0, 0x26,
	0xfe, 0x6a, 0x3a, 0xa7, 0x7e, 0x38, 0x4f, 0xd6, 0xba, 0xc8, 0xe3, 0xea, 0xff, 0x10, 0x77, 0xbb,
	0x63, 0x92, 0x62, 0x18, 0x9b, 0x21, 0x4d, 0x17, 0x0f, 0xc1, 0x4a, 0x97, 0x6a, 0x42, 0xa3, 0x44,
	0x36, 0x56, 0xfe, 0x3e, 0x72, 0xe1, 0x7d, 0x5e, 0x41, 0x99, 0x2d, 0xc7, 0x38, 0x5d, 0x85, 0xbc,
	0x3f, 0x95, 0x9c, 0x31, 0x7b, 0xb8, 0x0a, 0xf1, 0x57, 0xa0, 0xf1, 0xe1, 0x8f, 0xe7, 0x74, 0x1a,
	0x3c, 0xea, 0x67, 0x5c, 0x0b, 0x38, 0x64, 0x31, 0x84, 0xe5, 0x59, 0xfa, 0x93, 0xcf, 0xc1, 0x54,
	0x2f, 0xd7, 0x84, 0x46, 0x99, 0x6c, 0xac, 0xe2, 0x5b, 0xa9, 0x5f, 0xf8, 0x56, 0xaf, 0x41, 0xfd,
	0x14, 0x47, 0x74, 0xcc, 0xeb, 0x03, 0x5e, 0x47, 0x99, 0x01, 0x36, 0xab, 0xf1, 0x6b, 0xa8, 0x4c,
	0x83, 0x99, 0x9f, 0x86, 0xc9, 0xf8, 0x27, 0x3f, 0x4c, 0x03, 0x5d, 0xe3, 0x84, 0xf3, 0x0d, 0xf8,
	0x23, 0xc3, 0xea, 0x7f, 0x88, 0x20, 0xb3, 0x49, 0x62, 0x04, 0xe7, 0xde, 0xc8, 0x35, 0xc7, 0x43,
	0xfb, 0xde, 0x76, 0xde, 0xdb, 0xe8, 0x09, 0xbe, 0x00, 0x8d, 0x23, 0x5d, 0x67, 0xd8, 0xe9, 0x9b,
	0x48, 0xc0, 0x55, 0x00, 0x0e, 0xdc, 0xf5, 0x1d, 0xc3, 0x43, 0x62, 0x6e, 0x5b, 0xb6, 0x77, 0x73,
	0x8d, 0xa4, 0x3c, 0x60, 0x98, 0x01, 0x72, 0x91, 0xd0, 0x6e, 0xa1, 0x52, 0x9e, 0xe3, 0xce, 0xfa,
	0x60, 0x76, 0x6f, 0xae, 0x91, 0xb2, 0x8f, 0xb4, 0x5b, 0xe8, 0x0c, 0x57, 0x40, 0xe5, 0x48, 0xc7,
	0x71, 0xfa, 0xa8, 0x9c, 0x6b, 0x0e, 0x3c, 0x62, 0xd9, 0x3d, 0xa4, 0xe6, 0x9a, 0x3d, 0xe2, 0x0c,
	0x5d, 0x04, 0xb9, 0xc2, 0x3b, 0x73, 0x30, 0x30, 0x7a, 0x26, 0xd2, 0x72, 0x46, 0x67, 0xe4, 0x99,
	0x03, 0x74, 0xbe, 0x57, 0x56, 0xbb, 0x85, 0x2a, 0x79, 0x0a, 0xd3, 0x1e, 0xbe, 0x43, 0x55, 0xfc,
	0x14, 0x2a, 0x59, 0x8a, 0x6d, 0x11, 0x17, 0x07, 0xd0, 0xcd, 0x35, 0x42, 0xbb, 0x42, 0x32, 0x95,
	0xa7, 0x7b, 0xc0, 0xcd, 0x35, 0xc2, 0xf5, 0x04, 0xb4, 0xc2, 0x6e, 0xe1, 0x97, 0xf0, 0xec, 0xd6,
	0x20, 0x5d, 0xcb, 0x36, 0xfa, 0x96, 0x37, 0x2a, 0xcc, 0x55, 0x87, 0xe7, 0x45, 0x87, 0xe3, 0x7a,
	0x96, 0x63, 0x1b, 0x7d, 0x24, 0x1c, 0x7a, 0x88, 0xf9, 0xc3, 0xd0, 0x22, 0x66, 0x17, 0x89, 0xc7,
	0x1e, 0xd7, 0x34, 0x3c, 0xb3, 0x8b, 0xa4, 0xfa, 0x5f, 0x02, 0xc8, 0x26, 0x4d, 0x17, 0x27, 0xcf,
	0xc8, 0x77, 0xa0, 0x06, 0x34, 0x5d, 0x64, 0xcf, 0x9f, 0x5d, 0x92, 0xcb, 0xa3, 0xa5, 0x62, 0xd1,
	0x7c, 0x19, 0xc8, 0x8e, 0x5c, 0x5c, 0x46, 0xe9, 0x3f, 0x1f, 0x0e, 0xf9, 0xff, 0x1d, 0x8e, 0xd2,
	0x97, 0x1d, 0x8e, 0x4f, 0xa0, 0xe6, 0x2d, 0x9c, 0x9c, 0xc2, 0xee, 0xc3, 0x16, 0xf7, 0x3e, 0xec,
	0x7f, 0xdf, 0x63, 0xfd, 0x7b, 0x50, 0x32, 0xe8, 0x64, 0xa2, 0xb7, 0x50, 0xda, 0x8e, 0x9a, 0x35,
	0xfe, 0xfc, 0x48, 0xce, 0xa0, 0x6b, 0x92, 0x51, 0xde, 0x36, 0x41, 0xc9, 0xfa, 0x60, 0xcb, 0x36,
	0x18, 0xd9, 0x9e, 0xf1, 0x61, 0xec, 0x12, 0xc7, 0x73, 0x5a, 0xe8, 0xc9, 0x21, 0xd4, 0x46, 0x42,
	0xa7, 0x0f, 0xcf, 0x26, 0xd1, 0xe2, 0x50, 0xb1, 0xa3, 0xb2, 0xbf, 0x10, 0x97, 0x59, 0xae, 0xf0,
	0xa7, 0x20, 0xfc, 0x2a, 0x4a, 0x3d, 0xb7, 0xf3, 0x9b, 0xf8, 0xa6, 0x97, 0xf1, 0xdc, 0x6d, 0xe6,
	0xf7, 0x41, 0x18, 0xde, 0xd3, 0xe8, 0x67, 0xca, 0xf8, 0xf1, 0x83, 0xc2, 0x05, 0xda, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x1a, 0xdb, 0x3e, 0xf3, 0x06, 0x00, 0x00,
}
