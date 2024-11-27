// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: protofiles/proto.proto

package protofiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ServicioRegionales_EnviarEstadoDigimon_FullMethodName        = "/protofiles.ServicioRegionales/EnviarEstadoDigimon"
	ServicioRegionales_EnviarDataNode_FullMethodName             = "/protofiles.ServicioRegionales/EnviarDataNode"
	ServicioRegionales_ConsultarEstadoDiaboromon_FullMethodName  = "/protofiles.ServicioRegionales/ConsultarEstadoDiaboromon"
	ServicioRegionales_ConsultarEstadoTerminacion_FullMethodName = "/protofiles.ServicioRegionales/ConsultarEstadoTerminacion"
)

// ServicioRegionalesClient is the client API for ServicioRegionales service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioRegionalesClient interface {
	EnviarEstadoDigimon(ctx context.Context, in *DigimonRequest, opts ...grpc.CallOption) (*DigimonResponse, error)
	EnviarDataNode(ctx context.Context, in *DataNodeRequest, opts ...grpc.CallOption) (*DataNodeResponse, error)
	ConsultarEstadoDiaboromon(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*EstadoResponse, error)
	ConsultarEstadoTerminacion(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*EstadoResponse, error)
}

type servicioRegionalesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioRegionalesClient(cc grpc.ClientConnInterface) ServicioRegionalesClient {
	return &servicioRegionalesClient{cc}
}

func (c *servicioRegionalesClient) EnviarEstadoDigimon(ctx context.Context, in *DigimonRequest, opts ...grpc.CallOption) (*DigimonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DigimonResponse)
	err := c.cc.Invoke(ctx, ServicioRegionales_EnviarEstadoDigimon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioRegionalesClient) EnviarDataNode(ctx context.Context, in *DataNodeRequest, opts ...grpc.CallOption) (*DataNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataNodeResponse)
	err := c.cc.Invoke(ctx, ServicioRegionales_EnviarDataNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioRegionalesClient) ConsultarEstadoDiaboromon(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*EstadoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EstadoResponse)
	err := c.cc.Invoke(ctx, ServicioRegionales_ConsultarEstadoDiaboromon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioRegionalesClient) ConsultarEstadoTerminacion(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*EstadoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EstadoResponse)
	err := c.cc.Invoke(ctx, ServicioRegionales_ConsultarEstadoTerminacion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioRegionalesServer is the server API for ServicioRegionales service.
// All implementations must embed UnimplementedServicioRegionalesServer
// for forward compatibility.
type ServicioRegionalesServer interface {
	EnviarEstadoDigimon(context.Context, *DigimonRequest) (*DigimonResponse, error)
	EnviarDataNode(context.Context, *DataNodeRequest) (*DataNodeResponse, error)
	ConsultarEstadoDiaboromon(context.Context, *EstadoRequest) (*EstadoResponse, error)
	ConsultarEstadoTerminacion(context.Context, *EstadoRequest) (*EstadoResponse, error)
	mustEmbedUnimplementedServicioRegionalesServer()
}

// UnimplementedServicioRegionalesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicioRegionalesServer struct{}

func (UnimplementedServicioRegionalesServer) EnviarEstadoDigimon(context.Context, *DigimonRequest) (*DigimonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEstadoDigimon not implemented")
}
func (UnimplementedServicioRegionalesServer) EnviarDataNode(context.Context, *DataNodeRequest) (*DataNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarDataNode not implemented")
}
func (UnimplementedServicioRegionalesServer) ConsultarEstadoDiaboromon(context.Context, *EstadoRequest) (*EstadoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarEstadoDiaboromon not implemented")
}
func (UnimplementedServicioRegionalesServer) ConsultarEstadoTerminacion(context.Context, *EstadoRequest) (*EstadoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarEstadoTerminacion not implemented")
}
func (UnimplementedServicioRegionalesServer) mustEmbedUnimplementedServicioRegionalesServer() {}
func (UnimplementedServicioRegionalesServer) testEmbeddedByValue()                            {}

// UnsafeServicioRegionalesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioRegionalesServer will
// result in compilation errors.
type UnsafeServicioRegionalesServer interface {
	mustEmbedUnimplementedServicioRegionalesServer()
}

func RegisterServicioRegionalesServer(s grpc.ServiceRegistrar, srv ServicioRegionalesServer) {
	// If the following call pancis, it indicates UnimplementedServicioRegionalesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServicioRegionales_ServiceDesc, srv)
}

func _ServicioRegionales_EnviarEstadoDigimon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigimonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioRegionalesServer).EnviarEstadoDigimon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioRegionales_EnviarEstadoDigimon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioRegionalesServer).EnviarEstadoDigimon(ctx, req.(*DigimonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioRegionales_EnviarDataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioRegionalesServer).EnviarDataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioRegionales_EnviarDataNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioRegionalesServer).EnviarDataNode(ctx, req.(*DataNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioRegionales_ConsultarEstadoDiaboromon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioRegionalesServer).ConsultarEstadoDiaboromon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioRegionales_ConsultarEstadoDiaboromon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioRegionalesServer).ConsultarEstadoDiaboromon(ctx, req.(*EstadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioRegionales_ConsultarEstadoTerminacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioRegionalesServer).ConsultarEstadoTerminacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioRegionales_ConsultarEstadoTerminacion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioRegionalesServer).ConsultarEstadoTerminacion(ctx, req.(*EstadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioRegionales_ServiceDesc is the grpc.ServiceDesc for ServicioRegionales service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioRegionales_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.ServicioRegionales",
	HandlerType: (*ServicioRegionalesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarEstadoDigimon",
			Handler:    _ServicioRegionales_EnviarEstadoDigimon_Handler,
		},
		{
			MethodName: "EnviarDataNode",
			Handler:    _ServicioRegionales_EnviarDataNode_Handler,
		},
		{
			MethodName: "ConsultarEstadoDiaboromon",
			Handler:    _ServicioRegionales_ConsultarEstadoDiaboromon_Handler,
		},
		{
			MethodName: "ConsultarEstadoTerminacion",
			Handler:    _ServicioRegionales_ConsultarEstadoTerminacion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/proto.proto",
}

const (
	ServicioDataNodes_GuardarData_FullMethodName          = "/protofiles.ServicioDataNodes/GuardarData"
	ServicioDataNodes_NotificarTerminacion_FullMethodName = "/protofiles.ServicioDataNodes/NotificarTerminacion"
	ServicioDataNodes_ConsultarAtributos_FullMethodName   = "/protofiles.ServicioDataNodes/ConsultarAtributos"
)

// ServicioDataNodesClient is the client API for ServicioDataNodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioDataNodesClient interface {
	GuardarData(ctx context.Context, in *DataNodeRequest, opts ...grpc.CallOption) (*DataNodeResponse, error)
	NotificarTerminacion(ctx context.Context, in *TerminarRequest, opts ...grpc.CallOption) (*TerminarResponse, error)
	ConsultarAtributos(ctx context.Context, in *DataNodeQuery, opts ...grpc.CallOption) (*DataNodeQueryResponse, error)
}

type servicioDataNodesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioDataNodesClient(cc grpc.ClientConnInterface) ServicioDataNodesClient {
	return &servicioDataNodesClient{cc}
}

func (c *servicioDataNodesClient) GuardarData(ctx context.Context, in *DataNodeRequest, opts ...grpc.CallOption) (*DataNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataNodeResponse)
	err := c.cc.Invoke(ctx, ServicioDataNodes_GuardarData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioDataNodesClient) NotificarTerminacion(ctx context.Context, in *TerminarRequest, opts ...grpc.CallOption) (*TerminarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminarResponse)
	err := c.cc.Invoke(ctx, ServicioDataNodes_NotificarTerminacion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioDataNodesClient) ConsultarAtributos(ctx context.Context, in *DataNodeQuery, opts ...grpc.CallOption) (*DataNodeQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataNodeQueryResponse)
	err := c.cc.Invoke(ctx, ServicioDataNodes_ConsultarAtributos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioDataNodesServer is the server API for ServicioDataNodes service.
// All implementations must embed UnimplementedServicioDataNodesServer
// for forward compatibility.
type ServicioDataNodesServer interface {
	GuardarData(context.Context, *DataNodeRequest) (*DataNodeResponse, error)
	NotificarTerminacion(context.Context, *TerminarRequest) (*TerminarResponse, error)
	ConsultarAtributos(context.Context, *DataNodeQuery) (*DataNodeQueryResponse, error)
	mustEmbedUnimplementedServicioDataNodesServer()
}

// UnimplementedServicioDataNodesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicioDataNodesServer struct{}

func (UnimplementedServicioDataNodesServer) GuardarData(context.Context, *DataNodeRequest) (*DataNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuardarData not implemented")
}
func (UnimplementedServicioDataNodesServer) NotificarTerminacion(context.Context, *TerminarRequest) (*TerminarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificarTerminacion not implemented")
}
func (UnimplementedServicioDataNodesServer) ConsultarAtributos(context.Context, *DataNodeQuery) (*DataNodeQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarAtributos not implemented")
}
func (UnimplementedServicioDataNodesServer) mustEmbedUnimplementedServicioDataNodesServer() {}
func (UnimplementedServicioDataNodesServer) testEmbeddedByValue()                           {}

// UnsafeServicioDataNodesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioDataNodesServer will
// result in compilation errors.
type UnsafeServicioDataNodesServer interface {
	mustEmbedUnimplementedServicioDataNodesServer()
}

func RegisterServicioDataNodesServer(s grpc.ServiceRegistrar, srv ServicioDataNodesServer) {
	// If the following call pancis, it indicates UnimplementedServicioDataNodesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServicioDataNodes_ServiceDesc, srv)
}

func _ServicioDataNodes_GuardarData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioDataNodesServer).GuardarData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioDataNodes_GuardarData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioDataNodesServer).GuardarData(ctx, req.(*DataNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioDataNodes_NotificarTerminacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioDataNodesServer).NotificarTerminacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioDataNodes_NotificarTerminacion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioDataNodesServer).NotificarTerminacion(ctx, req.(*TerminarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioDataNodes_ConsultarAtributos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataNodeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioDataNodesServer).ConsultarAtributos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioDataNodes_ConsultarAtributos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioDataNodesServer).ConsultarAtributos(ctx, req.(*DataNodeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioDataNodes_ServiceDesc is the grpc.ServiceDesc for ServicioDataNodes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioDataNodes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.ServicioDataNodes",
	HandlerType: (*ServicioDataNodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GuardarData",
			Handler:    _ServicioDataNodes_GuardarData_Handler,
		},
		{
			MethodName: "NotificarTerminacion",
			Handler:    _ServicioDataNodes_NotificarTerminacion_Handler,
		},
		{
			MethodName: "ConsultarAtributos",
			Handler:    _ServicioDataNodes_ConsultarAtributos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/proto.proto",
}

const (
	ServicioTai_SolicitarDatos_FullMethodName         = "/protofiles.ServicioTai/SolicitarDatos"
	ServicioTai_Atacar_FullMethodName                 = "/protofiles.ServicioTai/Atacar"
	ServicioTai_EnviarEstadoDiaboromon_FullMethodName = "/protofiles.ServicioTai/EnviarEstadoDiaboromon"
	ServicioTai_NotificarTerminacion_FullMethodName   = "/protofiles.ServicioTai/NotificarTerminacion"
	ServicioTai_NotificarEstado_FullMethodName        = "/protofiles.ServicioTai/NotificarEstado"
)

// ServicioTaiClient is the client API for ServicioTai service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioTaiClient interface {
	SolicitarDatos(ctx context.Context, in *TaiRequest, opts ...grpc.CallOption) (*TaiResponse, error)
	Atacar(ctx context.Context, in *TaiRequest, opts ...grpc.CallOption) (*TaiResponse, error)
	EnviarEstadoDiaboromon(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*EstadoResponse, error)
	NotificarTerminacion(ctx context.Context, in *TerminarRequest, opts ...grpc.CallOption) (*TerminarResponse, error)
	// Notificar estado de termino a diaboromon
	NotificarEstado(ctx context.Context, in *EstadoTaiRequest, opts ...grpc.CallOption) (*EstadoTaiResponse, error)
}

type servicioTaiClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioTaiClient(cc grpc.ClientConnInterface) ServicioTaiClient {
	return &servicioTaiClient{cc}
}

func (c *servicioTaiClient) SolicitarDatos(ctx context.Context, in *TaiRequest, opts ...grpc.CallOption) (*TaiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaiResponse)
	err := c.cc.Invoke(ctx, ServicioTai_SolicitarDatos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioTaiClient) Atacar(ctx context.Context, in *TaiRequest, opts ...grpc.CallOption) (*TaiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaiResponse)
	err := c.cc.Invoke(ctx, ServicioTai_Atacar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioTaiClient) EnviarEstadoDiaboromon(ctx context.Context, in *EstadoRequest, opts ...grpc.CallOption) (*EstadoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EstadoResponse)
	err := c.cc.Invoke(ctx, ServicioTai_EnviarEstadoDiaboromon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioTaiClient) NotificarTerminacion(ctx context.Context, in *TerminarRequest, opts ...grpc.CallOption) (*TerminarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminarResponse)
	err := c.cc.Invoke(ctx, ServicioTai_NotificarTerminacion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioTaiClient) NotificarEstado(ctx context.Context, in *EstadoTaiRequest, opts ...grpc.CallOption) (*EstadoTaiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EstadoTaiResponse)
	err := c.cc.Invoke(ctx, ServicioTai_NotificarEstado_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioTaiServer is the server API for ServicioTai service.
// All implementations must embed UnimplementedServicioTaiServer
// for forward compatibility.
type ServicioTaiServer interface {
	SolicitarDatos(context.Context, *TaiRequest) (*TaiResponse, error)
	Atacar(context.Context, *TaiRequest) (*TaiResponse, error)
	EnviarEstadoDiaboromon(context.Context, *EstadoRequest) (*EstadoResponse, error)
	NotificarTerminacion(context.Context, *TerminarRequest) (*TerminarResponse, error)
	// Notificar estado de termino a diaboromon
	NotificarEstado(context.Context, *EstadoTaiRequest) (*EstadoTaiResponse, error)
	mustEmbedUnimplementedServicioTaiServer()
}

// UnimplementedServicioTaiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicioTaiServer struct{}

func (UnimplementedServicioTaiServer) SolicitarDatos(context.Context, *TaiRequest) (*TaiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarDatos not implemented")
}
func (UnimplementedServicioTaiServer) Atacar(context.Context, *TaiRequest) (*TaiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Atacar not implemented")
}
func (UnimplementedServicioTaiServer) EnviarEstadoDiaboromon(context.Context, *EstadoRequest) (*EstadoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEstadoDiaboromon not implemented")
}
func (UnimplementedServicioTaiServer) NotificarTerminacion(context.Context, *TerminarRequest) (*TerminarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificarTerminacion not implemented")
}
func (UnimplementedServicioTaiServer) NotificarEstado(context.Context, *EstadoTaiRequest) (*EstadoTaiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificarEstado not implemented")
}
func (UnimplementedServicioTaiServer) mustEmbedUnimplementedServicioTaiServer() {}
func (UnimplementedServicioTaiServer) testEmbeddedByValue()                     {}

// UnsafeServicioTaiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioTaiServer will
// result in compilation errors.
type UnsafeServicioTaiServer interface {
	mustEmbedUnimplementedServicioTaiServer()
}

func RegisterServicioTaiServer(s grpc.ServiceRegistrar, srv ServicioTaiServer) {
	// If the following call pancis, it indicates UnimplementedServicioTaiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServicioTai_ServiceDesc, srv)
}

func _ServicioTai_SolicitarDatos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTaiServer).SolicitarDatos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioTai_SolicitarDatos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTaiServer).SolicitarDatos(ctx, req.(*TaiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioTai_Atacar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTaiServer).Atacar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioTai_Atacar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTaiServer).Atacar(ctx, req.(*TaiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioTai_EnviarEstadoDiaboromon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTaiServer).EnviarEstadoDiaboromon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioTai_EnviarEstadoDiaboromon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTaiServer).EnviarEstadoDiaboromon(ctx, req.(*EstadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioTai_NotificarTerminacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTaiServer).NotificarTerminacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioTai_NotificarTerminacion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTaiServer).NotificarTerminacion(ctx, req.(*TerminarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioTai_NotificarEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoTaiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTaiServer).NotificarEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioTai_NotificarEstado_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTaiServer).NotificarEstado(ctx, req.(*EstadoTaiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioTai_ServiceDesc is the grpc.ServiceDesc for ServicioTai service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioTai_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.ServicioTai",
	HandlerType: (*ServicioTaiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitarDatos",
			Handler:    _ServicioTai_SolicitarDatos_Handler,
		},
		{
			MethodName: "Atacar",
			Handler:    _ServicioTai_Atacar_Handler,
		},
		{
			MethodName: "EnviarEstadoDiaboromon",
			Handler:    _ServicioTai_EnviarEstadoDiaboromon_Handler,
		},
		{
			MethodName: "NotificarTerminacion",
			Handler:    _ServicioTai_NotificarTerminacion_Handler,
		},
		{
			MethodName: "NotificarEstado",
			Handler:    _ServicioTai_NotificarEstado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/proto.proto",
}

const (
	ServicioDiaboromon_RecibirNotificacion_FullMethodName = "/protofiles.ServicioDiaboromon/RecibirNotificacion"
)

// ServicioDiaboromonClient is the client API for ServicioDiaboromon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioDiaboromonClient interface {
	RecibirNotificacion(ctx context.Context, in *DiaboromonRequest, opts ...grpc.CallOption) (*DiaboromonResponse, error)
}

type servicioDiaboromonClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioDiaboromonClient(cc grpc.ClientConnInterface) ServicioDiaboromonClient {
	return &servicioDiaboromonClient{cc}
}

func (c *servicioDiaboromonClient) RecibirNotificacion(ctx context.Context, in *DiaboromonRequest, opts ...grpc.CallOption) (*DiaboromonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiaboromonResponse)
	err := c.cc.Invoke(ctx, ServicioDiaboromon_RecibirNotificacion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioDiaboromonServer is the server API for ServicioDiaboromon service.
// All implementations must embed UnimplementedServicioDiaboromonServer
// for forward compatibility.
type ServicioDiaboromonServer interface {
	RecibirNotificacion(context.Context, *DiaboromonRequest) (*DiaboromonResponse, error)
	mustEmbedUnimplementedServicioDiaboromonServer()
}

// UnimplementedServicioDiaboromonServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicioDiaboromonServer struct{}

func (UnimplementedServicioDiaboromonServer) RecibirNotificacion(context.Context, *DiaboromonRequest) (*DiaboromonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirNotificacion not implemented")
}
func (UnimplementedServicioDiaboromonServer) mustEmbedUnimplementedServicioDiaboromonServer() {}
func (UnimplementedServicioDiaboromonServer) testEmbeddedByValue()                            {}

// UnsafeServicioDiaboromonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioDiaboromonServer will
// result in compilation errors.
type UnsafeServicioDiaboromonServer interface {
	mustEmbedUnimplementedServicioDiaboromonServer()
}

func RegisterServicioDiaboromonServer(s grpc.ServiceRegistrar, srv ServicioDiaboromonServer) {
	// If the following call pancis, it indicates UnimplementedServicioDiaboromonServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServicioDiaboromon_ServiceDesc, srv)
}

func _ServicioDiaboromon_RecibirNotificacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiaboromonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioDiaboromonServer).RecibirNotificacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioDiaboromon_RecibirNotificacion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioDiaboromonServer).RecibirNotificacion(ctx, req.(*DiaboromonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioDiaboromon_ServiceDesc is the grpc.ServiceDesc for ServicioDiaboromon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioDiaboromon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.ServicioDiaboromon",
	HandlerType: (*ServicioDiaboromonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecibirNotificacion",
			Handler:    _ServicioDiaboromon_RecibirNotificacion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/proto.proto",
}
