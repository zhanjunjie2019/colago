// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserLoginCmd struct {
	Dto                  *DTO     `protobuf:"bytes,1,opt,name=dto,proto3" json:"dto,omitempty"`
	AccKey               string   `protobuf:"bytes,2,opt,name=accKey,proto3" json:"accKey,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginCmd) Reset()         { *m = UserLoginCmd{} }
func (m *UserLoginCmd) String() string { return proto.CompactTextString(m) }
func (*UserLoginCmd) ProtoMessage()    {}
func (*UserLoginCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}
func (m *UserLoginCmd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserLoginCmd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserLoginCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginCmd.Merge(m, src)
}
func (m *UserLoginCmd) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginCmd.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginCmd proto.InternalMessageInfo

func (m *UserLoginCmd) GetDto() *DTO {
	if m != nil {
		return m.Dto
	}
	return nil
}

func (m *UserLoginCmd) GetAccKey() string {
	if m != nil {
		return m.AccKey
	}
	return ""
}

func (m *UserLoginCmd) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserLoginResponse struct {
	Rsp                  *Response      `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	Data                 *UserLoginData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}
func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetRsp() *Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func (m *UserLoginResponse) GetData() *UserLoginData {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserLoginData struct {
	ClientToken          string   `protobuf:"bytes,1,opt,name=clientToken,proto3" json:"clientToken,omitempty"`
	Ext                  uint64   `protobuf:"varint,2,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginData) Reset()         { *m = UserLoginData{} }
func (m *UserLoginData) String() string { return proto.CompactTextString(m) }
func (*UserLoginData) ProtoMessage()    {}
func (*UserLoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}
func (m *UserLoginData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserLoginData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserLoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginData.Merge(m, src)
}
func (m *UserLoginData) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginData.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginData proto.InternalMessageInfo

func (m *UserLoginData) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

func (m *UserLoginData) GetExt() uint64 {
	if m != nil {
		return m.Ext
	}
	return 0
}

func init() {
	proto.RegisterType((*UserLoginCmd)(nil), "client.UserLoginCmd")
	proto.RegisterType((*UserLoginResponse)(nil), "client.UserLoginResponse")
	proto.RegisterType((*UserLoginData)(nil), "client.UserLoginData")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x27, 0xb6, 0x14, 0xfb, 0xaa, 0x50, 0x1f, 0x2a, 0xb5, 0x60, 0x29, 0x5d, 0x8d, 0x9b,
	0x2e, 0xea, 0x05, 0xd4, 0xce, 0x42, 0x50, 0x10, 0xc2, 0x78, 0x80, 0x4c, 0x13, 0xa4, 0xe8, 0x34,
	0x25, 0x89, 0xa8, 0x37, 0xf1, 0x48, 0x2e, 0x3d, 0x82, 0xd4, 0x8b, 0x48, 0xd3, 0x69, 0x71, 0x98,
	0x5d, 0xde, 0xf7, 0xfe, 0xfc, 0x5f, 0x08, 0xc0, 0xab, 0x16, 0x2a, 0x6f, 0x95, 0x34, 0x12, 0xbd,
	0xea, 0xa5, 0x16, 0x8d, 0x89, 0x7d, 0x6e, 0xe4, 0x80, 0x32, 0x06, 0x07, 0x8f, 0x5a, 0xa8, 0x7b,
	0xf9, 0x54, 0x37, 0xe5, 0x9a, 0xe3, 0x39, 0x38, 0xdc, 0xc8, 0x88, 0xa4, 0x64, 0x1e, 0x14, 0x41,
	0x3e, 0x5c, 0xc8, 0x17, 0xcb, 0x07, 0xda, 0x73, 0x3c, 0x05, 0x8f, 0x55, 0xd5, 0x9d, 0xf8, 0x88,
	0xf6, 0x52, 0x32, 0xf7, 0xe9, 0x66, 0xc2, 0x18, 0xf6, 0x5b, 0xa6, 0xf5, 0x9b, 0x54, 0x3c, 0x72,
	0xec, 0x66, 0x9a, 0xb3, 0x15, 0x1c, 0x4d, 0x0a, 0x2a, 0x74, 0x2b, 0x1b, 0x2d, 0x30, 0x03, 0x47,
	0xe9, 0x76, 0xe3, 0x09, 0x47, 0xcf, 0xb8, 0xa6, 0xfd, 0x12, 0x2f, 0xc0, 0xe5, 0xcc, 0x30, 0xab,
	0x0a, 0x8a, 0x93, 0x31, 0x34, 0x95, 0x2d, 0x98, 0x61, 0xd4, 0x46, 0xb2, 0x12, 0x0e, 0xb7, 0x30,
	0xa6, 0x10, 0x0c, 0xf1, 0xa5, 0x7c, 0x16, 0x8d, 0xf5, 0xf8, 0xf4, 0x3f, 0xc2, 0x10, 0x1c, 0xf1,
	0x6e, 0x6c, 0xb9, 0x4b, 0xfb, 0x63, 0x71, 0x0b, 0x6e, 0x5f, 0x82, 0x57, 0x10, 0xd8, 0xa2, 0xeb,
	0xca, 0xd4, 0xb2, 0xc1, 0xe3, 0x1d, 0x71, 0xb9, 0xe6, 0xf1, 0xd9, 0x0e, 0x1d, 0x1f, 0x9f, 0xcd,
	0x6e, 0xc2, 0xaf, 0x2e, 0x21, 0xdf, 0x5d, 0x42, 0x7e, 0xba, 0x84, 0x7c, 0xfe, 0x26, 0xb3, 0x95,
	0x67, 0xbf, 0xfb, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x8d, 0x30, 0xa1, 0x8f, 0x01, 0x00,
	0x00,
}

func (m *UserLoginCmd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginCmd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserLoginCmd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccKey) > 0 {
		i -= len(m.AccKey)
		copy(dAtA[i:], m.AccKey)
		i = encodeVarintUser(dAtA, i, uint64(len(m.AccKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Dto != nil {
		{
			size, err := m.Dto.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUser(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserLoginResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserLoginResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUser(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Rsp != nil {
		{
			size, err := m.Rsp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUser(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserLoginData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserLoginData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Ext != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Ext))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientToken) > 0 {
		i -= len(m.ClientToken)
		copy(dAtA[i:], m.ClientToken)
		i = encodeVarintUser(dAtA, i, uint64(len(m.ClientToken)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserLoginCmd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dto != nil {
		l = m.Dto.Size()
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.AccKey)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserLoginResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rsp != nil {
		l = m.Rsp.Size()
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserLoginData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientToken)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Ext != 0 {
		n += 1 + sovUser(uint64(m.Ext))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserLoginCmd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserLoginCmd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginCmd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dto", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dto == nil {
				m.Dto = &DTO{}
			}
			if err := m.Dto.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserLoginResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserLoginResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rsp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rsp == nil {
				m.Rsp = &Response{}
			}
			if err := m.Rsp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &UserLoginData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserLoginData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserLoginData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ext", wireType)
			}
			m.Ext = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ext |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
