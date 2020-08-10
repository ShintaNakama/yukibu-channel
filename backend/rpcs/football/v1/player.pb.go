// Code generated by protoc-gen-go. DO NOT EDIT.
// source: football/v1/player.proto

package football

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Player_Role int32

const (
	Player_PLAYER Player_Role = 0
	Player_COACH  Player_Role = 1
)

var Player_Role_name = map[int32]string{
	0: "PLAYER",
	1: "COACH",
}

var Player_Role_value = map[string]int32{
	"PLAYER": 0,
	"COACH":  1,
}

func (x Player_Role) String() string {
	return proto.EnumName(Player_Role_name, int32(x))
}

func (Player_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb33e31429a401ee, []int{0, 0}
}

type Player struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position             string               `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	DateOfBirth          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	CountryOfBirth       string               `protobuf:"bytes,5,opt,name=countryOfBirth,proto3" json:"countryOfBirth,omitempty"`
	Nationality          string               `protobuf:"bytes,6,opt,name=nationality,proto3" json:"nationality,omitempty"`
	ShirtNumber          int32                `protobuf:"varint,7,opt,name=shirtNumber,proto3" json:"shirtNumber,omitempty"`
	Role                 Player_Role          `protobuf:"varint,8,opt,name=role,proto3,enum=football.v1.Player_Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb33e31429a401ee, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Player) GetDateOfBirth() *timestamp.Timestamp {
	if m != nil {
		return m.DateOfBirth
	}
	return nil
}

func (m *Player) GetCountryOfBirth() string {
	if m != nil {
		return m.CountryOfBirth
	}
	return ""
}

func (m *Player) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *Player) GetShirtNumber() int32 {
	if m != nil {
		return m.ShirtNumber
	}
	return 0
}

func (m *Player) GetRole() Player_Role {
	if m != nil {
		return m.Role
	}
	return Player_PLAYER
}

func init() {
	proto.RegisterEnum("football.v1.Player_Role", Player_Role_name, Player_Role_value)
	proto.RegisterType((*Player)(nil), "football.v1.Player")
}

func init() { proto.RegisterFile("football/v1/player.proto", fileDescriptor_bb33e31429a401ee) }

var fileDescriptor_bb33e31429a401ee = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0xd7, 0x5e, 0xc7, 0x9b, 0xc8, 0x10, 0x82, 0x4e, 0x22, 0xb0, 0xd4, 0xe4, 0x50, 0x7c,
	0x68, 0x25, 0x92, 0x1e, 0xdb, 0x4b, 0x12, 0x0a, 0xa1, 0x94, 0x24, 0xb8, 0xbd, 0xb4, 0x37, 0xc9,
	0x56, 0x62, 0x11, 0x59, 0x32, 0xb2, 0x1c, 0xf0, 0xaf, 0xeb, 0x5f, 0x2b, 0x51, 0x70, 0x31, 0xbd,
	0xcd, 0xbc, 0xf9, 0x34, 0x7a, 0xbc, 0x01, 0xe8, 0xa0, 0xb5, 0x65, 0x54, 0x4a, 0x72, 0x9e, 0x93,
	0x4a, 0xd2, 0x96, 0x1b, 0x5c, 0x19, 0x6d, 0x35, 0x8c, 0xba, 0x09, 0x3e, 0xcf, 0xa7, 0x37, 0x47,
	0xad, 0x8f, 0x92, 0x13, 0x37, 0x62, 0xcd, 0x81, 0x58, 0x51, 0xf2, 0xda, 0xd2, 0xb2, 0xba, 0xd2,
	0xb3, 0x2f, 0x1f, 0x84, 0x7b, 0xf7, 0x1c, 0x8e, 0x81, 0x2f, 0x72, 0xe4, 0xc5, 0x5e, 0x32, 0x4a,
	0x7d, 0x91, 0x43, 0x08, 0x02, 0x45, 0x4b, 0x8e, 0x7c, 0xa7, 0xb8, 0x1a, 0x4e, 0xc1, 0xb0, 0xd2,
	0xb5, 0xb0, 0x42, 0x2b, 0xf4, 0xd7, 0xe9, 0x3f, 0x3d, 0x7c, 0x02, 0x51, 0x4e, 0x2d, 0xdf, 0x1d,
	0x56, 0xc2, 0xd8, 0x02, 0x05, 0xb1, 0x97, 0x44, 0x8b, 0x29, 0xbe, 0x3a, 0xc0, 0x9d, 0x03, 0xfc,
	0xde, 0x39, 0x48, 0xfb, 0x38, 0xbc, 0x05, 0xe3, 0x4c, 0x37, 0xca, 0x9a, 0xb6, 0x5b, 0x30, 0x70,
	0xfb, 0x7f, 0xa9, 0x30, 0x06, 0x91, 0xa2, 0x97, 0xff, 0xa8, 0x14, 0xb6, 0x45, 0xa1, 0x83, 0xfa,
	0xd2, 0x85, 0xa8, 0x0b, 0x61, 0xec, 0xb6, 0x29, 0x19, 0x37, 0xe8, 0x5f, 0xec, 0x25, 0x83, 0xb4,
	0x2f, 0xc1, 0x3b, 0x10, 0x18, 0x2d, 0x39, 0x1a, 0xc6, 0x5e, 0x32, 0x5e, 0x20, 0xdc, 0x4b, 0x0c,
	0x5f, 0xc3, 0xc0, 0xa9, 0x96, 0x3c, 0x75, 0xd4, 0xec, 0x3f, 0x08, 0x2e, 0x1d, 0x04, 0x20, 0xdc,
	0xbf, 0x2e, 0x3f, 0x9e, 0xd3, 0xc9, 0x1f, 0x38, 0x02, 0x83, 0xf5, 0x6e, 0xb9, 0xde, 0x4c, 0xbc,
	0xd5, 0xcb, 0xe7, 0xe6, 0x28, 0x6c, 0xd1, 0x30, 0x9c, 0xe9, 0x92, 0xbc, 0x15, 0x42, 0x59, 0xba,
	0xa5, 0x27, 0x5a, 0x52, 0xd2, 0x36, 0x27, 0xc1, 0x9a, 0xfb, 0xac, 0xa0, 0x4a, 0x71, 0x49, 0x18,
	0xcd, 0x4e, 0x5c, 0xe5, 0xc4, 0x54, 0x59, 0x4d, 0x7a, 0xf7, 0x7b, 0xec, 0x6a, 0x16, 0xba, 0x94,
	0x1e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x74, 0xf6, 0xa5, 0xde, 0x01, 0x00, 0x00,
}