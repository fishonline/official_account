// Code generated by protoc-gen-go.
// source: product.proto
// DO NOT EDIT!

package igrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ColorPatternRelationship struct {
	Id              int64  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Code            string `protobuf:"bytes,3,opt,name=Code" json:"Code,omitempty"`
	ColorName       string `protobuf:"bytes,4,opt,name=ColorName" json:"ColorName,omitempty"`
	ColorCode       string `protobuf:"bytes,5,opt,name=ColorCode" json:"ColorCode,omitempty"`
	CompanyId       int64  `protobuf:"varint,6,opt,name=CompanyId" json:"CompanyId,omitempty"`
	FabricPatternId int64  `protobuf:"varint,7,opt,name=FabricPatternId" json:"FabricPatternId,omitempty"`
	ColorId         int64  `protobuf:"varint,9,opt,name=ColorId" json:"ColorId,omitempty"`
	Amount          int64  `protobuf:"varint,10,opt,name=Amount" json:"Amount,omitempty"`
	Weight          int64  `protobuf:"varint,11,opt,name=Weight" json:"Weight,omitempty"`
	UnitType        int64  `protobuf:"varint,12,opt,name=UnitType" json:"UnitType,omitempty"`
}

func (m *ColorPatternRelationship) Reset()                    { *m = ColorPatternRelationship{} }
func (m *ColorPatternRelationship) String() string            { return proto.CompactTextString(m) }
func (*ColorPatternRelationship) ProtoMessage()               {}
func (*ColorPatternRelationship) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ColorPatternRelationship) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ColorPatternRelationship) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColorPatternRelationship) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ColorPatternRelationship) GetColorName() string {
	if m != nil {
		return m.ColorName
	}
	return ""
}

func (m *ColorPatternRelationship) GetColorCode() string {
	if m != nil {
		return m.ColorCode
	}
	return ""
}

func (m *ColorPatternRelationship) GetCompanyId() int64 {
	if m != nil {
		return m.CompanyId
	}
	return 0
}

func (m *ColorPatternRelationship) GetFabricPatternId() int64 {
	if m != nil {
		return m.FabricPatternId
	}
	return 0
}

func (m *ColorPatternRelationship) GetColorId() int64 {
	if m != nil {
		return m.ColorId
	}
	return 0
}

func (m *ColorPatternRelationship) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ColorPatternRelationship) GetWeight() int64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ColorPatternRelationship) GetUnitType() int64 {
	if m != nil {
		return m.UnitType
	}
	return 0
}

type ColorPatternRelationshipsResp struct {
	ColorPatternRelationships []*ColorPatternRelationship `protobuf:"bytes,1,rep,name=ColorPatternRelationships" json:"ColorPatternRelationships,omitempty"`
}

func (m *ColorPatternRelationshipsResp) Reset()                    { *m = ColorPatternRelationshipsResp{} }
func (m *ColorPatternRelationshipsResp) String() string            { return proto.CompactTextString(m) }
func (*ColorPatternRelationshipsResp) ProtoMessage()               {}
func (*ColorPatternRelationshipsResp) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *ColorPatternRelationshipsResp) GetColorPatternRelationships() []*ColorPatternRelationship {
	if m != nil {
		return m.ColorPatternRelationships
	}
	return nil
}

type FabricPattern struct {
	Id                      int64  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                    string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Code                    string `protobuf:"bytes,3,opt,name=Code" json:"Code,omitempty"`
	Breadth                 string `protobuf:"bytes,4,opt,name=Breadth" json:"Breadth,omitempty"`
	BreadthErrorRange       int64  `protobuf:"varint,5,opt,name=BreadthErrorRange" json:"BreadthErrorRange,omitempty"`
	Gsm                     string `protobuf:"bytes,6,opt,name=Gsm" json:"Gsm,omitempty"`
	GsmErrorRange           int64  `protobuf:"varint,7,opt,name=GsmErrorRange" json:"GsmErrorRange,omitempty"`
	WeavingComposition      string `protobuf:"bytes,8,opt,name=WeavingComposition" json:"WeavingComposition,omitempty"`
	LowPricePk              int64  `protobuf:"varint,9,opt,name=LowPricePk" json:"LowPricePk,omitempty"`
	RefPricePk              int64  `protobuf:"varint,10,opt,name=RefPricePk" json:"RefPricePk,omitempty"`
	SellingPricePk          int64  `protobuf:"varint,11,opt,name=SellingPricePk" json:"SellingPricePk,omitempty"`
	LowPricePm              int64  `protobuf:"varint,12,opt,name=LowPricePm" json:"LowPricePm,omitempty"`
	RefPricePm              int64  `protobuf:"varint,13,opt,name=RefPricePm" json:"RefPricePm,omitempty"`
	SellingPricePm          int64  `protobuf:"varint,14,opt,name=SellingPricePm" json:"SellingPricePm,omitempty"`
	Desc                    string `protobuf:"bytes,15,opt,name=Desc" json:"Desc,omitempty"`
	WeavingStructure        string `protobuf:"bytes,16,opt,name=WeavingStructure" json:"WeavingStructure,omitempty"`
	PaperTubeSpecifications string `protobuf:"bytes,17,opt,name=PaperTubeSpecifications" json:"PaperTubeSpecifications,omitempty"`
}

func (m *FabricPattern) Reset()                    { *m = FabricPattern{} }
func (m *FabricPattern) String() string            { return proto.CompactTextString(m) }
func (*FabricPattern) ProtoMessage()               {}
func (*FabricPattern) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *FabricPattern) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FabricPattern) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FabricPattern) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *FabricPattern) GetBreadth() string {
	if m != nil {
		return m.Breadth
	}
	return ""
}

func (m *FabricPattern) GetBreadthErrorRange() int64 {
	if m != nil {
		return m.BreadthErrorRange
	}
	return 0
}

func (m *FabricPattern) GetGsm() string {
	if m != nil {
		return m.Gsm
	}
	return ""
}

func (m *FabricPattern) GetGsmErrorRange() int64 {
	if m != nil {
		return m.GsmErrorRange
	}
	return 0
}

func (m *FabricPattern) GetWeavingComposition() string {
	if m != nil {
		return m.WeavingComposition
	}
	return ""
}

func (m *FabricPattern) GetLowPricePk() int64 {
	if m != nil {
		return m.LowPricePk
	}
	return 0
}

func (m *FabricPattern) GetRefPricePk() int64 {
	if m != nil {
		return m.RefPricePk
	}
	return 0
}

func (m *FabricPattern) GetSellingPricePk() int64 {
	if m != nil {
		return m.SellingPricePk
	}
	return 0
}

func (m *FabricPattern) GetLowPricePm() int64 {
	if m != nil {
		return m.LowPricePm
	}
	return 0
}

func (m *FabricPattern) GetRefPricePm() int64 {
	if m != nil {
		return m.RefPricePm
	}
	return 0
}

func (m *FabricPattern) GetSellingPricePm() int64 {
	if m != nil {
		return m.SellingPricePm
	}
	return 0
}

func (m *FabricPattern) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *FabricPattern) GetWeavingStructure() string {
	if m != nil {
		return m.WeavingStructure
	}
	return ""
}

func (m *FabricPattern) GetPaperTubeSpecifications() string {
	if m != nil {
		return m.PaperTubeSpecifications
	}
	return ""
}

type ColorInfo struct {
	Id        int64  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Code      string `protobuf:"bytes,2,opt,name=Code" json:"Code,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	CompanyId int64  `protobuf:"varint,4,opt,name=CompanyId" json:"CompanyId,omitempty"`
}

func (m *ColorInfo) Reset()                    { *m = ColorInfo{} }
func (m *ColorInfo) String() string            { return proto.CompactTextString(m) }
func (*ColorInfo) ProtoMessage()               {}
func (*ColorInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ColorInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ColorInfo) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ColorInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColorInfo) GetCompanyId() int64 {
	if m != nil {
		return m.CompanyId
	}
	return 0
}

type Colors struct {
	Colors []*ColorInfo `protobuf:"bytes,1,rep,name=Colors" json:"Colors,omitempty"`
}

func (m *Colors) Reset()                    { *m = Colors{} }
func (m *Colors) String() string            { return proto.CompactTextString(m) }
func (*Colors) ProtoMessage()               {}
func (*Colors) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *Colors) GetColors() []*ColorInfo {
	if m != nil {
		return m.Colors
	}
	return nil
}

type FabricPatternInfo struct {
	Id   int64  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=Code" json:"Code,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
}

func (m *FabricPatternInfo) Reset()                    { *m = FabricPatternInfo{} }
func (m *FabricPatternInfo) String() string            { return proto.CompactTextString(m) }
func (*FabricPatternInfo) ProtoMessage()               {}
func (*FabricPatternInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *FabricPatternInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FabricPatternInfo) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *FabricPatternInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ColorPatternRelationship)(nil), "igrpc.ColorPatternRelationship")
	proto.RegisterType((*ColorPatternRelationshipsResp)(nil), "igrpc.ColorPatternRelationshipsResp")
	proto.RegisterType((*FabricPattern)(nil), "igrpc.FabricPattern")
	proto.RegisterType((*ColorInfo)(nil), "igrpc.ColorInfo")
	proto.RegisterType((*Colors)(nil), "igrpc.Colors")
	proto.RegisterType((*FabricPatternInfo)(nil), "igrpc.FabricPatternInfo")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0x9b, 0x6e, 0xbb, 0x9d, 0xa5, 0xdd, 0x76, 0x0e, 0x60, 0x10, 0x1f, 0x55, 0x84, 0x50,
	0x85, 0x50, 0x0f, 0xcb, 0x85, 0x2b, 0x2c, 0xb0, 0x8a, 0x40, 0xa8, 0x72, 0x17, 0xed, 0x89, 0x83,
	0x9b, 0xb8, 0xad, 0x45, 0x1d, 0x47, 0x8e, 0x0b, 0xda, 0x0b, 0xbf, 0x8c, 0x0b, 0xff, 0x0c, 0xd9,
	0x71, 0xb6, 0x49, 0x3f, 0x2e, 0x70, 0x9b, 0x79, 0x6f, 0x3c, 0x63, 0xe7, 0xbd, 0x09, 0xf4, 0x32,
	0xad, 0x92, 0x4d, 0x6c, 0x26, 0x99, 0x56, 0x46, 0xe1, 0x89, 0x58, 0xea, 0x2c, 0x0e, 0x7f, 0x37,
	0x81, 0x5c, 0xaa, 0xb5, 0xd2, 0x53, 0x66, 0x0c, 0xd7, 0x29, 0xe5, 0x6b, 0x66, 0x84, 0x4a, 0xf3,
	0x95, 0xc8, 0xb0, 0x0f, 0xcd, 0x28, 0x21, 0x8d, 0x51, 0x63, 0x1c, 0xd0, 0x66, 0x94, 0x20, 0x42,
	0xeb, 0x0b, 0x93, 0x9c, 0x34, 0x47, 0x8d, 0x71, 0x97, 0xba, 0xd8, 0x62, 0x97, 0x2a, 0xe1, 0x24,
	0x28, 0x30, 0x1b, 0xe3, 0x63, 0xe8, 0xba, 0x9e, 0xae, 0xb8, 0xe5, 0x88, 0x2d, 0x70, 0xc7, 0xba,
	0x63, 0x27, 0x15, 0x76, 0x7b, 0x56, 0x66, 0x2c, 0xbd, 0x8d, 0x12, 0xd2, 0x76, 0xa3, 0xb7, 0x00,
	0x8e, 0xe1, 0xfc, 0x23, 0x9b, 0x6b, 0x11, 0xfb, 0xeb, 0x46, 0x09, 0xe9, 0xb8, 0x9a, 0x5d, 0x18,
	0x09, 0x74, 0x5c, 0xd3, 0x28, 0x21, 0x5d, 0x57, 0x51, 0xa6, 0x78, 0x1f, 0xda, 0x6f, 0xa5, 0xda,
	0xa4, 0x86, 0x80, 0x23, 0x7c, 0x66, 0xf1, 0x1b, 0x2e, 0x96, 0x2b, 0x43, 0xce, 0x0a, 0xbc, 0xc8,
	0xf0, 0x11, 0x9c, 0x7e, 0x4d, 0x85, 0xb9, 0xbe, 0xcd, 0x38, 0xb9, 0xe7, 0x98, 0xbb, 0x3c, 0xfc,
	0x05, 0x4f, 0x8e, 0x7d, 0xbd, 0x9c, 0xf2, 0x3c, 0xc3, 0x6f, 0xf0, 0xf0, 0x68, 0x01, 0x69, 0x8c,
	0x82, 0xf1, 0xd9, 0xc5, 0xb3, 0x89, 0x93, 0x62, 0x72, 0xac, 0x8e, 0x1e, 0xef, 0x10, 0xfe, 0x69,
	0x41, 0xaf, 0xf6, 0xf2, 0x7f, 0xd6, 0x8c, 0x40, 0xe7, 0x9d, 0xe6, 0x2c, 0x31, 0x2b, 0xaf, 0x58,
	0x99, 0xe2, 0x2b, 0x18, 0xfa, 0xf0, 0x83, 0xd6, 0x4a, 0x53, 0x96, 0x2e, 0x0b, 0xdd, 0x02, 0xba,
	0x4f, 0xe0, 0x00, 0x82, 0xab, 0x5c, 0x3a, 0xe5, 0xba, 0xd4, 0x86, 0xf8, 0x1c, 0x7a, 0x57, 0xb9,
	0xac, 0x9c, 0x2d, 0x14, 0xab, 0x83, 0x38, 0x01, 0xbc, 0xe1, 0xec, 0x87, 0x48, 0x97, 0x56, 0x6d,
	0x95, 0x0b, 0xfb, 0x48, 0x72, 0xea, 0xda, 0x1c, 0x60, 0xf0, 0x29, 0xc0, 0x67, 0xf5, 0x73, 0xaa,
	0x45, 0xcc, 0xa7, 0xdf, 0xbd, 0xc4, 0x15, 0xc4, 0xf2, 0x94, 0x2f, 0x4a, 0xbe, 0x50, 0xba, 0x82,
	0xe0, 0x0b, 0xe8, 0xcf, 0xf8, 0x7a, 0x2d, 0xd2, 0x65, 0x59, 0x53, 0xa8, 0xbe, 0x83, 0xd6, 0xe6,
	0x48, 0xaf, 0x7f, 0x05, 0xa9, 0xcd, 0x91, 0xa4, 0xb7, 0x33, 0x47, 0xee, 0xcd, 0x91, 0xa4, 0x7f,
	0x60, 0x8e, 0xb4, 0x9a, 0xbc, 0xe7, 0x79, 0x4c, 0xce, 0x0b, 0x4d, 0x6c, 0x8c, 0x2f, 0x61, 0xe0,
	0x5f, 0x3e, 0x33, 0x7a, 0x13, 0x9b, 0x8d, 0xe6, 0x64, 0xe0, 0xf8, 0x3d, 0x1c, 0xdf, 0xc0, 0x83,
	0x29, 0xcb, 0xb8, 0xbe, 0xde, 0xcc, 0xf9, 0x2c, 0xe3, 0xb1, 0x58, 0x88, 0xb8, 0x30, 0x0a, 0x19,
	0xba, 0x23, 0xc7, 0xe8, 0x90, 0xf9, 0x7d, 0x8c, 0xd2, 0x85, 0x3a, 0x64, 0x1f, 0x67, 0x95, 0x66,
	0xc5, 0x2a, 0xa5, 0xa5, 0x82, 0x8a, 0xa5, 0x6a, 0x6b, 0xdb, 0xda, 0x59, 0xdb, 0xf0, 0x02, 0xda,
	0x6e, 0x44, 0x8e, 0xe3, 0x32, 0xf2, 0xe6, 0x1f, 0x54, 0xcd, 0x6f, 0x6f, 0x40, 0x3d, 0x1f, 0x7e,
	0x82, 0x61, 0x7d, 0xa7, 0xff, 0xe3, 0x7a, 0xf3, 0xb6, 0xfb, 0xe9, 0xbd, 0xfe, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xce, 0xb8, 0x34, 0xbc, 0x05, 0x05, 0x00, 0x00,
}
