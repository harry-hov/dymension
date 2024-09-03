// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/app.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type App struct {
	// name is the unique App's name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// rollapp_id is the id of the Rollapp the App belongs to
	RollappId string `protobuf:"bytes,2,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	// description is the description of the App
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// image_url is the URL to the App's image
	ImageUrl string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// url is the URL to the App's website
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// order is the order of the App in the Rollapp
	Order int32 `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f01e4af248858b2, []int{0}
}
func (m *App) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_App.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return m.Size()
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *App) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *App) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *App) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *App) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func init() {
	proto.RegisterType((*App)(nil), "dymensionxyz.dymension.rollapp.App")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/app.proto", fileDescriptor_8f01e4af248858b2)
}

var fileDescriptor_8f01e4af248858b2 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xab, 0xa8, 0xac, 0xd2, 0x87, 0x73, 0xf4, 0x8b, 0xf2, 0x73, 0x72,
	0x12, 0x0b, 0x0a, 0xf4, 0x13, 0x0b, 0x0a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x90,
	0x55, 0xea, 0xc1, 0x39, 0x7a, 0x50, 0x95, 0x4a, 0xf3, 0x19, 0xb9, 0x98, 0x1d, 0x0b, 0x0a, 0x84,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0x21, 0x59, 0x2e, 0x2e, 0xa8, 0xb2, 0xf8, 0xcc, 0x14, 0x09, 0x26, 0xb0, 0x0c, 0x27, 0x54, 0xc4,
	0x33, 0x45, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0x33, 0x3f,
	0x4f, 0x82, 0x19, 0x2c, 0x8f, 0x2c, 0x24, 0x24, 0xcd, 0xc5, 0x99, 0x99, 0x9b, 0x98, 0x9e, 0x1a,
	0x5f, 0x5a, 0x94, 0x23, 0xc1, 0x02, 0x96, 0xe7, 0x00, 0x0b, 0x84, 0x16, 0xe5, 0x08, 0x09, 0x70,
	0x31, 0x83, 0x84, 0x59, 0xc1, 0xc2, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x7e, 0x51, 0x4a, 0x6a,
	0x91, 0x04, 0x9b, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xe4, 0x77, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c,
	0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0x38, 0x02, 0xa4, 0xcc, 0x58, 0xbf, 0x02, 0x1e, 0x2a, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0xe0, 0x80, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xfe, 0xed, 0x45, 0x44,
	0x01, 0x00, 0x00,
}

func (m *App) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *App) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *App) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order != 0 {
		i = encodeVarintApp(dAtA, i, uint64(m.Order))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintApp(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintApp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApp(dAtA []byte, offset int, v uint64) int {
	offset -= sovApp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *App) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	if m.Order != 0 {
		n += 1 + sovApp(uint64(m.Order))
	}
	return n
}

func sovApp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApp(x uint64) (n int) {
	return sovApp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *App) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: App: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: App: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			m.Order = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Order |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApp = fmt.Errorf("proto: unexpected end of group")
)
