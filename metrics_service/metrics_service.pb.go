// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics_service.proto

/*
Package metrics_service is a generated protocol buffer package.

It is generated from these files:
	metrics_service.proto

It has these top-level messages:
	Metric
	MetricsBatch
	MetricsBatchRequest
	MetricsBatchResponse
*/
package metrics_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metric represents one pieces of time series data
type Metric struct {
	Name      string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Timestamp int64   `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Value     float32 `protobuf:"fixed32,3,opt,name=value" json:"value,omitempty"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metric) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// MetricsBatch is a collection of metrics sent at one point in time
type MetricsBatch struct {
	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	Count   int32     `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *MetricsBatch) Reset()                    { *m = MetricsBatch{} }
func (m *MetricsBatch) String() string            { return proto.CompactTextString(m) }
func (*MetricsBatch) ProtoMessage()               {}
func (*MetricsBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MetricsBatch) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *MetricsBatch) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type MetricsBatchRequest struct {
	MetricsData *MetricsBatch `protobuf:"bytes,1,opt,name=metricsData" json:"metricsData,omitempty"`
	Uuid        string        `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *MetricsBatchRequest) Reset()                    { *m = MetricsBatchRequest{} }
func (m *MetricsBatchRequest) String() string            { return proto.CompactTextString(m) }
func (*MetricsBatchRequest) ProtoMessage()               {}
func (*MetricsBatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetricsBatchRequest) GetMetricsData() *MetricsBatch {
	if m != nil {
		return m.MetricsData
	}
	return nil
}

func (m *MetricsBatchRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type MetricsBatchResponse struct {
	Ok   bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *MetricsBatchResponse) Reset()                    { *m = MetricsBatchResponse{} }
func (m *MetricsBatchResponse) String() string            { return proto.CompactTextString(m) }
func (*MetricsBatchResponse) ProtoMessage()               {}
func (*MetricsBatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetricsBatchResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *MetricsBatchResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Metric)(nil), "Metric")
	proto.RegisterType((*MetricsBatch)(nil), "MetricsBatch")
	proto.RegisterType((*MetricsBatchRequest)(nil), "MetricsBatchRequest")
	proto.RegisterType((*MetricsBatchResponse)(nil), "MetricsBatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MetricsBridge service

type MetricsBridgeClient interface {
	ProcessMetrics(ctx context.Context, in *MetricsBatchRequest, opts ...grpc.CallOption) (*MetricsBatchResponse, error)
}

type metricsBridgeClient struct {
	cc *grpc.ClientConn
}

func NewMetricsBridgeClient(cc *grpc.ClientConn) MetricsBridgeClient {
	return &metricsBridgeClient{cc}
}

func (c *metricsBridgeClient) ProcessMetrics(ctx context.Context, in *MetricsBatchRequest, opts ...grpc.CallOption) (*MetricsBatchResponse, error) {
	out := new(MetricsBatchResponse)
	err := grpc.Invoke(ctx, "/MetricsBridge/ProcessMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetricsBridge service

type MetricsBridgeServer interface {
	ProcessMetrics(context.Context, *MetricsBatchRequest) (*MetricsBatchResponse, error)
}

func RegisterMetricsBridgeServer(s *grpc.Server, srv MetricsBridgeServer) {
	s.RegisterService(&_MetricsBridge_serviceDesc, srv)
}

func _MetricsBridge_ProcessMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsBridgeServer).ProcessMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetricsBridge/ProcessMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsBridgeServer).ProcessMetrics(ctx, req.(*MetricsBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsBridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MetricsBridge",
	HandlerType: (*MetricsBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessMetrics",
			Handler:    _MetricsBridge_ProcessMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics_service.proto",
}

func init() { proto.RegisterFile("metrics_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0xc6, 0xdd, 0xd4, 0xb6, 0x66, 0x62, 0x7b, 0x58, 0x5b, 0x08, 0xe2, 0x21, 0xee, 0x29, 0xa7,
	0x15, 0xe2, 0xcd, 0x8b, 0x20, 0x82, 0xa7, 0x42, 0xd8, 0xa3, 0x17, 0x59, 0xb7, 0x83, 0x2e, 0x35,
	0xd9, 0xb8, 0x8f, 0xfe, 0xfd, 0xd2, 0x4d, 0x83, 0x2d, 0xe4, 0x36, 0xaf, 0xef, 0xf7, 0xcd, 0x30,
	0xb0, 0x6e, 0xd0, 0x5b, 0xad, 0xdc, 0x87, 0x43, 0xbb, 0xd7, 0x0a, 0x79, 0x67, 0x8d, 0x37, 0xac,
	0x86, 0xd9, 0x26, 0x36, 0x28, 0x85, 0xcb, 0x56, 0x36, 0x98, 0x93, 0x82, 0x94, 0xa9, 0x88, 0x31,
	0xbd, 0x83, 0xd4, 0xeb, 0x06, 0x9d, 0x97, 0x4d, 0x97, 0x27, 0x05, 0x29, 0x27, 0xe2, 0xbf, 0x40,
	0x57, 0x30, 0xdd, 0xcb, 0x9f, 0x80, 0xf9, 0xa4, 0x20, 0x65, 0x22, 0xfa, 0x84, 0xbd, 0xc1, 0x75,
	0x4f, 0x74, 0x2f, 0xd2, 0xab, 0x6f, 0x7a, 0x0f, 0xf3, 0xa3, 0x75, 0x4e, 0x8a, 0x49, 0x99, 0x55,
	0x73, 0xde, 0xf7, 0xc5, 0x50, 0x3f, 0x80, 0x94, 0x09, 0xad, 0x8f, 0x16, 0x53, 0xd1, 0x27, 0xec,
	0x1d, 0x6e, 0x4e, 0x41, 0x02, 0x7f, 0x03, 0x3a, 0x4f, 0x1f, 0x20, 0x3b, 0xea, 0x5e, 0xa5, 0x97,
	0x71, 0xdd, 0xac, 0x5a, 0xf0, 0xb3, 0xd1, 0xd3, 0x89, 0xc3, 0x61, 0x21, 0xe8, 0x6d, 0x84, 0xa7,
	0x22, 0xc6, 0xec, 0x09, 0x56, 0xe7, 0x6c, 0xd7, 0x99, 0xd6, 0x21, 0x5d, 0x42, 0x62, 0x76, 0x91,
	0x79, 0x25, 0x12, 0xb3, 0x1b, 0xd3, 0x56, 0x35, 0x2c, 0x06, 0xad, 0xd5, 0xdb, 0x2f, 0xa4, 0xcf,
	0xb0, 0xac, 0xad, 0x51, 0xe8, 0xdc, 0x66, 0x38, 0x88, 0x8f, 0x6c, 0x7e, 0xbb, 0xe6, 0x63, 0x9e,
	0xec, 0xe2, 0x73, 0x16, 0x7f, 0xf1, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x5e, 0xb8, 0x95,
	0xa4, 0x01, 0x00, 0x00,
}
