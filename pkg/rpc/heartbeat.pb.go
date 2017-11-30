// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/heartbeat.proto

/*
	Package rpc is a generated protocol buffer package.

	It is generated from these files:
		rpc/heartbeat.proto

	It has these top-level messages:
		RemoteOffset
		PingRequest
		PingResponse
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_cockroachdb_cockroach_pkg_util_uuid "github.com/cockroachdb/cockroach/pkg/util/uuid"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// RemoteOffset keeps track of this client's estimate of its offset from a
// remote server. Uncertainty is the maximum error in the reading of this
// offset, so that the real offset should be in the interval
// [Offset - Uncertainty, Offset + Uncertainty]. If the last heartbeat timed
// out, Offset = 0.
//
// Offset and Uncertainty are measured using the remote clock reading technique
// described in http://se.inf.tu-dresden.de/pubs/papers/SRDS1994.pdf, page 6.
type RemoteOffset struct {
	// The estimated offset from the remote server, in nanoseconds.
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// The maximum error of the measured offset, in nanoseconds.
	Uncertainty int64 `protobuf:"varint,2,opt,name=uncertainty" json:"uncertainty"`
	// Measurement time, in nanoseconds from unix epoch.
	MeasuredAt int64 `protobuf:"varint,3,opt,name=measured_at,json=measuredAt" json:"measured_at"`
}

func (m *RemoteOffset) Reset()                    { *m = RemoteOffset{} }
func (*RemoteOffset) ProtoMessage()               {}
func (*RemoteOffset) Descriptor() ([]byte, []int) { return fileDescriptorHeartbeat, []int{0} }

// A PingRequest specifies the string to echo in response.
// Fields are exported so that they will be serialized in the rpc call.
type PingRequest struct {
	// Echo this string with PingResponse.
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping"`
	// The last offset the client measured with the server.
	Offset RemoteOffset `protobuf:"bytes,2,opt,name=offset" json:"offset"`
	// The address of the client.
	Addr string `protobuf:"bytes,3,opt,name=addr" json:"addr"`
	// The configured maximum clock offset (in nanoseconds) on the server.
	MaxOffsetNanos int64 `protobuf:"varint,4,opt,name=max_offset_nanos,json=maxOffsetNanos" json:"max_offset_nanos"`
	// Cluster ID to prevent connections between nodes in different clusters.
	ClusterID *github_com_cockroachdb_cockroach_pkg_util_uuid.UUID `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,customtype=github.com/cockroachdb/cockroach/pkg/util/uuid.UUID" json:"cluster_id,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptorHeartbeat, []int{1} }

// A PingResponse contains the echoed ping request string.
type PingResponse struct {
	// An echo of value sent with PingRequest.
	Pong       string `protobuf:"bytes,1,opt,name=pong" json:"pong"`
	ServerTime int64  `protobuf:"varint,2,opt,name=server_time,json=serverTime" json:"server_time"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptorHeartbeat, []int{2} }

func init() {
	proto.RegisterType((*RemoteOffset)(nil), "cockroach.rpc.RemoteOffset")
	proto.RegisterType((*PingRequest)(nil), "cockroach.rpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "cockroach.rpc.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Heartbeat service

type HeartbeatClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/cockroach.rpc.Heartbeat/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Heartbeat service

type HeartbeatServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.rpc.Heartbeat/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.rpc.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Heartbeat_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/heartbeat.proto",
}

func (m *RemoteOffset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteOffset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.Offset))
	dAtA[i] = 0x10
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.Uncertainty))
	dAtA[i] = 0x18
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.MeasuredAt))
	return i, nil
}

func (m *PingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.Ping)))
	i += copy(dAtA[i:], m.Ping)
	dAtA[i] = 0x12
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.Offset.Size()))
	n1, err := m.Offset.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.Addr)))
	i += copy(dAtA[i:], m.Addr)
	dAtA[i] = 0x20
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.MaxOffsetNanos))
	if m.ClusterID != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHeartbeat(dAtA, i, uint64(m.ClusterID.Size()))
		n2, err := m.ClusterID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *PingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.Pong)))
	i += copy(dAtA[i:], m.Pong)
	dAtA[i] = 0x10
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.ServerTime))
	return i, nil
}

func encodeVarintHeartbeat(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RemoteOffset) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHeartbeat(uint64(m.Offset))
	n += 1 + sovHeartbeat(uint64(m.Uncertainty))
	n += 1 + sovHeartbeat(uint64(m.MeasuredAt))
	return n
}

func (m *PingRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ping)
	n += 1 + l + sovHeartbeat(uint64(l))
	l = m.Offset.Size()
	n += 1 + l + sovHeartbeat(uint64(l))
	l = len(m.Addr)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.MaxOffsetNanos))
	if m.ClusterID != nil {
		l = m.ClusterID.Size()
		n += 1 + l + sovHeartbeat(uint64(l))
	}
	return n
}

func (m *PingResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Pong)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.ServerTime))
	return n
}

func sovHeartbeat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeartbeat(x uint64) (n int) {
	return sovHeartbeat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteOffset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoteOffset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteOffset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uncertainty", wireType)
			}
			m.Uncertainty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uncertainty |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasuredAt", wireType)
			}
			m.MeasuredAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MeasuredAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func (m *PingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ping = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Offset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOffsetNanos", wireType)
			}
			m.MaxOffsetNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxOffsetNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cockroachdb_cockroach_pkg_util_uuid.UUID
			m.ClusterID = &v
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func (m *PingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pong = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			m.ServerTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func skipHeartbeat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeartbeat
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
					return 0, ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHeartbeat
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHeartbeat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeartbeat
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHeartbeat(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHeartbeat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeartbeat   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc/heartbeat.proto", fileDescriptorHeartbeat) }

var fileDescriptorHeartbeat = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x8d, 0x41, 0xca, 0x9b, 0x80, 0xd0, 0xc1, 0x60, 0xa5, 0xe0, 0x54, 0x91, 0x40,
	0x9d, 0x6c, 0xa9, 0x4c, 0xb0, 0x35, 0x64, 0x20, 0x4b, 0x8b, 0x2c, 0xba, 0xb0, 0x58, 0x97, 0xf3,
	0x5b, 0xe7, 0xd4, 0xfa, 0xce, 0xdc, 0x1f, 0x54, 0x46, 0xbe, 0x01, 0x6c, 0x8c, 0x7c, 0x9c, 0x8c,
	0x8c, 0x88, 0x21, 0x02, 0xf3, 0x45, 0x90, 0xed, 0x34, 0xb8, 0xa1, 0xdb, 0xab, 0xf7, 0x79, 0xfc,
	0xf8, 0x77, 0xcf, 0x1d, 0x3c, 0xd4, 0x25, 0x8f, 0x97, 0xc8, 0xb4, 0x5d, 0x20, 0xb3, 0x51, 0xa9,
	0x95, 0x55, 0xf4, 0x1e, 0x57, 0xfc, 0x42, 0x2b, 0xc6, 0x97, 0x91, 0x2e, 0xf9, 0xe8, 0x51, 0xae,
	0x72, 0xd5, 0x28, 0x71, 0x3d, 0xb5, 0xa6, 0xc9, 0x27, 0x02, 0xc3, 0x04, 0x0b, 0x65, 0xf1, 0xf4,
	0xfc, 0xdc, 0xa0, 0xa5, 0x8f, 0xe1, 0xae, 0x6a, 0xa6, 0x80, 0x1c, 0x90, 0xc3, 0xde, 0xd4, 0x5f,
	0xad, 0xc7, 0x5e, 0xb2, 0xd9, 0xd1, 0x67, 0x30, 0x70, 0x92, 0xa3, 0xb6, 0x4c, 0x48, 0xfb, 0x31,
	0xd8, 0xeb, 0x58, 0xba, 0x02, 0x7d, 0x0a, 0x83, 0x02, 0x99, 0x71, 0x1a, 0xb3, 0x94, 0xd9, 0xa0,
	0xd7, 0xf1, 0xc1, 0xb5, 0x70, 0x6c, 0x5f, 0xfa, 0x5f, 0xbf, 0x8d, 0xbd, 0xc9, 0x97, 0x3d, 0x18,
	0xbc, 0x11, 0x32, 0x4f, 0xf0, 0xbd, 0x43, 0x63, 0x69, 0x00, 0x7e, 0x29, 0x64, 0xde, 0x00, 0xf4,
	0x37, 0x5f, 0x35, 0x1b, 0xfa, 0x62, 0x0b, 0x57, 0xff, 0x79, 0x70, 0xb4, 0x1f, 0xdd, 0x38, 0x63,
	0xd4, 0x3d, 0xc9, 0x0e, 0x79, 0x00, 0x3e, 0xcb, 0x32, 0xdd, 0xa0, 0x6c, 0x43, 0xeb, 0x0d, 0x8d,
	0xe0, 0x41, 0xc1, 0xae, 0xd2, 0xd6, 0x97, 0x4a, 0x26, 0x95, 0x09, 0xfc, 0x0e, 0xf0, 0xfd, 0x82,
	0x5d, 0xb5, 0x91, 0x27, 0xb5, 0x46, 0x39, 0x00, 0xbf, 0x74, 0xc6, 0xa2, 0x4e, 0x45, 0x16, 0xdc,
	0x39, 0x20, 0x87, 0xc3, 0xe9, 0xec, 0xe7, 0x7a, 0xfc, 0x3c, 0x17, 0x76, 0xe9, 0x16, 0x11, 0x57,
	0x45, 0xbc, 0xc5, 0xca, 0x16, 0xff, 0xe6, 0xb8, 0xbc, 0xc8, 0x63, 0x67, 0xc5, 0x65, 0xec, 0x9c,
	0xc8, 0xa2, 0xb3, 0xb3, 0xf9, 0xac, 0x5a, 0x8f, 0xfb, 0xaf, 0xda, 0xb0, 0xf9, 0x2c, 0xe9, 0x6f,
	0x72, 0xe7, 0xd9, 0xe4, 0x14, 0x86, 0x6d, 0x25, 0xa6, 0x54, 0xd2, 0x60, 0xd3, 0x89, 0xfa, 0xaf,
	0x13, 0x25, 0xf3, 0xba, 0x6a, 0x83, 0xfa, 0x03, 0xea, 0xd4, 0x8a, 0x02, 0x6f, 0x5c, 0x09, 0xb4,
	0xc2, 0x5b, 0x51, 0xe0, 0xd1, 0x09, 0xf4, 0x5f, 0x5f, 0x3f, 0x10, 0x7a, 0x0c, 0x7e, 0x9d, 0x4e,
	0x47, 0x3b, 0xfd, 0x75, 0x6e, 0x61, 0xb4, 0x7f, 0xab, 0xd6, 0xe2, 0x4c, 0xbc, 0xe9, 0x93, 0xd5,
	0xef, 0xd0, 0x5b, 0x55, 0x21, 0xf9, 0x5e, 0x85, 0xe4, 0x47, 0x15, 0x92, 0x5f, 0x55, 0x48, 0x3e,
	0xff, 0x09, 0xbd, 0x77, 0x3d, 0x5d, 0xf2, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x1b, 0xfb,
	0xe9, 0x92, 0x02, 0x00, 0x00,
}
