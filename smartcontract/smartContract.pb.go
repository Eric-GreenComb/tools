// Code generated by protoc-gen-go.
// source: smartContract.proto
// DO NOT EDIT!

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SmartContract model
type SmartContract struct {
	Addr                 string       `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Id                   string       `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Detail               string       `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
	Goal                 uint64       `protobuf:"varint,5,opt,name=goal" json:"goal,omitempty"`
	PartyA               string       `protobuf:"bytes,6,opt,name=partyA" json:"partyA,omitempty"`
	PartyB               string       `protobuf:"bytes,7,opt,name=partyB" json:"partyB,omitempty"`
	Type                 uint64       `protobuf:"varint,8,opt,name=type" json:"type,omitempty"`
	FundManangerFee      uint64       `protobuf:"varint,9,opt,name=fundManangerFee" json:"fundManangerFee,omitempty"`
	ChannelFee           uint64       `protobuf:"varint,10,opt,name=channelFee" json:"channelFee,omitempty"`
	CreateTimestamp      int64        `protobuf:"varint,11,opt,name=createTimestamp" json:"createTimestamp,omitempty"`
	UtilTimestamp        int64        `protobuf:"varint,12,opt,name=utilTimestamp" json:"utilTimestamp,omitempty"`
	TerminationTimestamp int64        `protobuf:"varint,13,opt,name=terminationTimestamp" json:"terminationTimestamp,omitempty"`
	Foundation           string       `protobuf:"bytes,14,opt,name=foundation" json:"foundation,omitempty"`
	Attachhash           string       `protobuf:"bytes,15,opt,name=attachhash" json:"attachhash,omitempty"`
	TContract            []*TContract `protobuf:"bytes,16,rep,name=tContract" json:"tContract,omitempty"`
	Status               uint64       `protobuf:"varint,17,opt,name=status" json:"status,omitempty"`
	Remark               string       `protobuf:"bytes,18,opt,name=remark" json:"remark,omitempty"`
}

func (m *SmartContract) Reset()                    { *m = SmartContract{} }
func (m *SmartContract) String() string            { return proto.CompactTextString(m) }
func (*SmartContract) ProtoMessage()               {}
func (*SmartContract) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SmartContract) GetTContract() []*TContract {
	if m != nil {
		return m.TContract
	}
	return nil
}

type TContract struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PartyA          string `protobuf:"bytes,2,opt,name=partyA" json:"partyA,omitempty"`
	PartyB          string `protobuf:"bytes,3,opt,name=partyB" json:"partyB,omitempty"`
	SignedTimestamp int64  `protobuf:"varint,4,opt,name=signedTimestamp" json:"signedTimestamp,omitempty"`
	Hash            string `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
}

func (m *TContract) Reset()                    { *m = TContract{} }
func (m *TContract) String() string            { return proto.CompactTextString(m) }
func (*TContract) ProtoMessage()               {}
func (*TContract) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type SmartContractTrack struct {
	Addr  string                  `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Id    string                  `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Total uint64                  `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Trans []*SmartContractHistory `protobuf:"bytes,4,rep,name=trans" json:"trans,omitempty"`
}

func (m *SmartContractTrack) Reset()                    { *m = SmartContractTrack{} }
func (m *SmartContractTrack) String() string            { return proto.CompactTextString(m) }
func (*SmartContractTrack) ProtoMessage()               {}
func (*SmartContractTrack) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *SmartContractTrack) GetTrans() []*SmartContractHistory {
	if m != nil {
		return m.Trans
	}
	return nil
}

type SmartContractHistory struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type            string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Amount          string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	SignedTimestamp int64  `protobuf:"varint,4,opt,name=signedTimestamp" json:"signedTimestamp,omitempty"`
}

func (m *SmartContractHistory) Reset()                    { *m = SmartContractHistory{} }
func (m *SmartContractHistory) String() string            { return proto.CompactTextString(m) }
func (*SmartContractHistory) ProtoMessage()               {}
func (*SmartContractHistory) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func init() {
	proto.RegisterType((*SmartContract)(nil), "protos.SmartContract")
	proto.RegisterType((*TContract)(nil), "protos.TContract")
	proto.RegisterType((*SmartContractTrack)(nil), "protos.SmartContractTrack")
	proto.RegisterType((*SmartContractHistory)(nil), "protos.SmartContractHistory")
}

func init() { proto.RegisterFile("smartContract.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xda, 0xb4, 0xd0, 0x59, 0xba, 0x65, 0x4d, 0xb5, 0xf2, 0x01, 0xa1, 0xaa, 0xe2, 0x90,
	0xd3, 0x22, 0x95, 0x2f, 0x60, 0x91, 0x10, 0x17, 0x2e, 0xa1, 0x3f, 0x30, 0x24, 0xde, 0xd6, 0xda,
	0xc4, 0xae, 0xec, 0xc9, 0xa1, 0x37, 0xc4, 0x9d, 0xdf, 0xe0, 0x3b, 0x91, 0x3d, 0xcd, 0x26, 0xa9,
	0xba, 0xd2, 0x9e, 0x32, 0xf3, 0xde, 0x93, 0x3d, 0x7e, 0xf3, 0x02, 0xef, 0x7c, 0x8d, 0x8e, 0xbe,
	0x5a, 0x43, 0x0e, 0x0b, 0xba, 0x3b, 0x38, 0x4b, 0x56, 0x4c, 0xe3, 0xc7, 0xaf, 0xff, 0xa5, 0x30,
	0xff, 0xd9, 0xe7, 0x85, 0x80, 0x14, 0xcb, 0xd2, 0xc9, 0x64, 0x95, 0x64, 0xb3, 0x3c, 0xd6, 0xe2,
	0x1a, 0x46, 0xba, 0x94, 0xa3, 0x88, 0x8c, 0x74, 0x19, 0x34, 0x06, 0x6b, 0x25, 0xc7, 0xac, 0x09,
	0xb5, 0xb8, 0x85, 0x69, 0xa9, 0x08, 0x75, 0x25, 0xd3, 0x88, 0x9e, 0xba, 0xa0, 0xdd, 0x59, 0xac,
	0xe4, 0x64, 0x95, 0x64, 0x69, 0x1e, 0xeb, 0xa0, 0x3d, 0xa0, 0xa3, 0xe3, 0x17, 0x39, 0x65, 0x2d,
	0x77, 0x4f, 0xf8, 0xbd, 0x7c, 0xd5, 0xc3, 0xef, 0xc3, 0x19, 0x74, 0x3c, 0x28, 0xf9, 0x9a, 0xcf,
	0x08, 0xb5, 0xc8, 0x60, 0xf1, 0xd0, 0x98, 0xf2, 0x07, 0x1a, 0x34, 0x3b, 0xe5, 0xbe, 0x29, 0x25,
	0x67, 0x91, 0x3e, 0x87, 0xc5, 0x07, 0x80, 0x62, 0x8f, 0xc6, 0xa8, 0x2a, 0x88, 0x20, 0x8a, 0x7a,
	0x48, 0x38, 0xa9, 0x70, 0x0a, 0x49, 0x6d, 0x75, 0xad, 0x3c, 0x61, 0x7d, 0x90, 0x57, 0xab, 0x24,
	0x1b, 0xe7, 0xe7, 0xb0, 0xf8, 0x08, 0xf3, 0x86, 0x74, 0xd5, 0xe9, 0xde, 0x44, 0xdd, 0x10, 0x14,
	0x1b, 0x58, 0x92, 0x72, 0xb5, 0x36, 0x48, 0xda, 0x9a, 0x4e, 0x3c, 0x8f, 0xe2, 0x8b, 0x5c, 0x98,
	0xf1, 0xc1, 0x36, 0xa6, 0x8c, 0xb0, 0xbc, 0x8e, 0xaf, 0xef, 0x21, 0x81, 0x47, 0x22, 0x2c, 0xf6,
	0x7b, 0xf4, 0x7b, 0xb9, 0x60, 0xbe, 0x43, 0xc4, 0x27, 0x98, 0x3d, 0xad, 0x50, 0xbe, 0x5d, 0x8d,
	0xb3, 0xab, 0xcd, 0x0d, 0xaf, 0xda, 0xdf, 0x6d, 0x5b, 0x22, 0xef, 0x34, 0xc1, 0x6a, 0x4f, 0x48,
	0x8d, 0x97, 0x37, 0xd1, 0x90, 0x53, 0x17, 0x70, 0xa7, 0x6a, 0x74, 0x8f, 0x52, 0xf0, 0x0a, 0xb8,
	0x5b, 0xff, 0x4d, 0x60, 0xb6, 0xed, 0x87, 0x24, 0x06, 0x20, 0x19, 0x06, 0xe0, 0xb4, 0xd4, 0xd1,
	0x33, 0x4b, 0x1d, 0x0f, 0x96, 0x9a, 0xc1, 0xc2, 0xeb, 0x9d, 0x51, 0x65, 0xe7, 0x50, 0xca, 0xb6,
	0x9f, 0xc1, 0xe1, 0xb6, 0xf8, 0xec, 0x09, 0xdf, 0x16, 0xea, 0xf5, 0x9f, 0x04, 0xc4, 0x20, 0xb8,
	0x5b, 0x87, 0xc5, 0xe3, 0x8b, 0xd2, 0xbb, 0x84, 0x09, 0x59, 0xc2, 0x2a, 0xce, 0x93, 0xe6, 0xdc,
	0x88, 0x0d, 0x4c, 0xc8, 0xa1, 0xf1, 0x32, 0x8d, 0xee, 0xbd, 0x6f, 0xdd, 0x1b, 0x5c, 0xf2, 0x5d,
	0x7b, 0xb2, 0xee, 0x98, 0xb3, 0x74, 0xfd, 0x3b, 0x81, 0xe5, 0x25, 0xfe, 0xa2, 0x3f, 0x6d, 0x88,
	0x79, 0x10, 0x0e, 0xf1, 0x2d, 0x4c, 0xb1, 0xb6, 0x8d, 0xa1, 0xd6, 0x1b, 0xee, 0x5e, 0xee, 0xcd,
	0x2f, 0xfe, 0x91, 0x3f, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x63, 0x93, 0x9a, 0xe6, 0x03,
	0x00, 0x00,
}
