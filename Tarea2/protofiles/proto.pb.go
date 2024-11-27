// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: protofiles/proto.proto

package protofiles

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DigimonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MensajeCifrado string `protobuf:"bytes,1,opt,name=mensaje_cifrado,json=mensajeCifrado,proto3" json:"mensaje_cifrado,omitempty"`
}

func (x *DigimonRequest) Reset() {
	*x = DigimonRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DigimonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigimonRequest) ProtoMessage() {}

func (x *DigimonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigimonRequest.ProtoReflect.Descriptor instead.
func (*DigimonRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{0}
}

func (x *DigimonRequest) GetMensajeCifrado() string {
	if x != nil {
		return x.MensajeCifrado
	}
	return ""
}

type DigimonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *DigimonResponse) Reset() {
	*x = DigimonResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DigimonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigimonResponse) ProtoMessage() {}

func (x *DigimonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigimonResponse.ProtoReflect.Descriptor instead.
func (*DigimonResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{1}
}

func (x *DigimonResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type DataNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TipoDigimon string `protobuf:"bytes,3,opt,name=tipo_digimon,json=tipoDigimon,proto3" json:"tipo_digimon,omitempty"`
}

func (x *DataNodeRequest) Reset() {
	*x = DataNodeRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodeRequest) ProtoMessage() {}

func (x *DataNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodeRequest.ProtoReflect.Descriptor instead.
func (*DataNodeRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{2}
}

func (x *DataNodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataNodeRequest) GetTipoDigimon() string {
	if x != nil {
		return x.TipoDigimon
	}
	return ""
}

type DataNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *DataNodeResponse) Reset() {
	*x = DataNodeResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodeResponse) ProtoMessage() {}

func (x *DataNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodeResponse.ProtoReflect.Descriptor instead.
func (*DataNodeResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{3}
}

func (x *DataNodeResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type TaiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accion string `protobuf:"bytes,1,opt,name=accion,proto3" json:"accion,omitempty"`
}

func (x *TaiRequest) Reset() {
	*x = TaiRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaiRequest) ProtoMessage() {}

func (x *TaiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaiRequest.ProtoReflect.Descriptor instead.
func (*TaiRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{4}
}

func (x *TaiRequest) GetAccion() string {
	if x != nil {
		return x.Accion
	}
	return ""
}

type TaiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CantidadDatos float64 `protobuf:"fixed64,1,opt,name=cantidad_datos,json=cantidadDatos,proto3" json:"cantidad_datos,omitempty"`
	Mensaje       string  `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *TaiResponse) Reset() {
	*x = TaiResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaiResponse) ProtoMessage() {}

func (x *TaiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaiResponse.ProtoReflect.Descriptor instead.
func (*TaiResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{5}
}

func (x *TaiResponse) GetCantidadDatos() float64 {
	if x != nil {
		return x.CantidadDatos
	}
	return 0
}

func (x *TaiResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type EstadoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EstadoRequest) Reset() {
	*x = EstadoRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstadoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoRequest) ProtoMessage() {}

func (x *EstadoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoRequest.ProtoReflect.Descriptor instead.
func (*EstadoRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{6}
}

type EstadoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activo  bool   `protobuf:"varint,1,opt,name=activo,proto3" json:"activo,omitempty"`
	Mensaje string `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *EstadoResponse) Reset() {
	*x = EstadoResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstadoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoResponse) ProtoMessage() {}

func (x *EstadoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoResponse.ProtoReflect.Descriptor instead.
func (*EstadoResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{7}
}

func (x *EstadoResponse) GetActivo() bool {
	if x != nil {
		return x.Activo
	}
	return false
}

func (x *EstadoResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type TerminarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminarRequest) Reset() {
	*x = TerminarRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TerminarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminarRequest) ProtoMessage() {}

func (x *TerminarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminarRequest.ProtoReflect.Descriptor instead.
func (*TerminarRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{8}
}

type TerminarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *TerminarResponse) Reset() {
	*x = TerminarResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TerminarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminarResponse) ProtoMessage() {}

func (x *TerminarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminarResponse.ProtoReflect.Descriptor instead.
func (*TerminarResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{9}
}

func (x *TerminarResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

// Mensajes para notificar el estado de tai a diaboromon
type EstadoTaiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EstadoTaiRequest) Reset() {
	*x = EstadoTaiRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstadoTaiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoTaiRequest) ProtoMessage() {}

func (x *EstadoTaiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoTaiRequest.ProtoReflect.Descriptor instead.
func (*EstadoTaiRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{10}
}

type EstadoTaiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EstadoJuego string `protobuf:"bytes,1,opt,name=estado_juego,json=estadoJuego,proto3" json:"estado_juego,omitempty"` // "ganado" o "perdido"
	Mensaje     string `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *EstadoTaiResponse) Reset() {
	*x = EstadoTaiResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EstadoTaiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoTaiResponse) ProtoMessage() {}

func (x *EstadoTaiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoTaiResponse.ProtoReflect.Descriptor instead.
func (*EstadoTaiResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{11}
}

func (x *EstadoTaiResponse) GetEstadoJuego() string {
	if x != nil {
		return x.EstadoJuego
	}
	return ""
}

func (x *EstadoTaiResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type DataNodeQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"` // IDs de los Digimon sacrificados
}

func (x *DataNodeQuery) Reset() {
	*x = DataNodeQuery{}
	mi := &file_protofiles_proto_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataNodeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodeQuery) ProtoMessage() {}

func (x *DataNodeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodeQuery.ProtoReflect.Descriptor instead.
func (*DataNodeQuery) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{12}
}

func (x *DataNodeQuery) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DataNodeQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Atributos []string `protobuf:"bytes,1,rep,name=atributos,proto3" json:"atributos,omitempty"` // Las líneas crudas del archivo
}

func (x *DataNodeQueryResponse) Reset() {
	*x = DataNodeQueryResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataNodeQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodeQueryResponse) ProtoMessage() {}

func (x *DataNodeQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodeQueryResponse.ProtoReflect.Descriptor instead.
func (*DataNodeQueryResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{13}
}

func (x *DataNodeQueryResponse) GetAtributos() []string {
	if x != nil {
		return x.Atributos
	}
	return nil
}

type DiaboromonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VidaRestante int32  `protobuf:"varint,1,opt,name=vida_restante,json=vidaRestante,proto3" json:"vida_restante,omitempty"`
	Mensaje      string `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *DiaboromonRequest) Reset() {
	*x = DiaboromonRequest{}
	mi := &file_protofiles_proto_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiaboromonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiaboromonRequest) ProtoMessage() {}

func (x *DiaboromonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiaboromonRequest.ProtoReflect.Descriptor instead.
func (*DiaboromonRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{14}
}

func (x *DiaboromonRequest) GetVidaRestante() int32 {
	if x != nil {
		return x.VidaRestante
	}
	return 0
}

func (x *DiaboromonRequest) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

type DiaboromonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *DiaboromonResponse) Reset() {
	*x = DiaboromonResponse{}
	mi := &file_protofiles_proto_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiaboromonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiaboromonResponse) ProtoMessage() {}

func (x *DiaboromonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_proto_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiaboromonResponse.ProtoReflect.Descriptor instead.
func (*DiaboromonResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_proto_proto_rawDescGZIP(), []int{15}
}

func (x *DiaboromonResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_protofiles_proto_proto protoreflect.FileDescriptor

var file_protofiles_proto_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x5f, 0x63, 0x69, 0x66, 0x72, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x43, 0x69, 0x66, 0x72, 0x61, 0x64, 0x6f, 0x22,
	0x2b, 0x0a, 0x0f, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x44, 0x0a, 0x0f,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x70, 0x6f, 0x44, 0x69, 0x67, 0x69, 0x6d,
	0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x22, 0x24, 0x0a, 0x0a, 0x54, 0x61, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x0b, 0x54, 0x61, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63,
	0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c,
	0x0a, 0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x12, 0x0a, 0x10,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x54, 0x61, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x50, 0x0a, 0x11, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x54, 0x61, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x5f,
	0x6a, 0x75, 0x65, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x73, 0x22, 0x52, 0x0a, 0x11,
	0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x76, 0x69, 0x64, 0x61, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x22, 0x2e, 0x0a, 0x12, 0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x32, 0xda, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x13, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x67, 0x69,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61,
	0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x84, 0x02,
	0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x41, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x6f, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xff, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69,
	0x6f, 0x54, 0x61, 0x69, 0x12, 0x41, 0x0a, 0x0e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61,
	0x72, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x41, 0x74, 0x61, 0x63, 0x61,
	0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54,
	0x61, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x16, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x54, 0x61, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x54, 0x61, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6a, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x69, 0x6f, 0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x13,
	0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x44, 0x69, 0x61, 0x62, 0x6f, 0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x54, 0x61, 0x72, 0x65, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_proto_proto_rawDescOnce sync.Once
	file_protofiles_proto_proto_rawDescData = file_protofiles_proto_proto_rawDesc
)

func file_protofiles_proto_proto_rawDescGZIP() []byte {
	file_protofiles_proto_proto_rawDescOnce.Do(func() {
		file_protofiles_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_proto_proto_rawDescData)
	})
	return file_protofiles_proto_proto_rawDescData
}

var file_protofiles_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_protofiles_proto_proto_goTypes = []any{
	(*DigimonRequest)(nil),        // 0: protofiles.DigimonRequest
	(*DigimonResponse)(nil),       // 1: protofiles.DigimonResponse
	(*DataNodeRequest)(nil),       // 2: protofiles.DataNodeRequest
	(*DataNodeResponse)(nil),      // 3: protofiles.DataNodeResponse
	(*TaiRequest)(nil),            // 4: protofiles.TaiRequest
	(*TaiResponse)(nil),           // 5: protofiles.TaiResponse
	(*EstadoRequest)(nil),         // 6: protofiles.EstadoRequest
	(*EstadoResponse)(nil),        // 7: protofiles.EstadoResponse
	(*TerminarRequest)(nil),       // 8: protofiles.TerminarRequest
	(*TerminarResponse)(nil),      // 9: protofiles.TerminarResponse
	(*EstadoTaiRequest)(nil),      // 10: protofiles.EstadoTaiRequest
	(*EstadoTaiResponse)(nil),     // 11: protofiles.EstadoTaiResponse
	(*DataNodeQuery)(nil),         // 12: protofiles.DataNodeQuery
	(*DataNodeQueryResponse)(nil), // 13: protofiles.DataNodeQueryResponse
	(*DiaboromonRequest)(nil),     // 14: protofiles.DiaboromonRequest
	(*DiaboromonResponse)(nil),    // 15: protofiles.DiaboromonResponse
}
var file_protofiles_proto_proto_depIdxs = []int32{
	0,  // 0: protofiles.ServicioRegionales.EnviarEstadoDigimon:input_type -> protofiles.DigimonRequest
	2,  // 1: protofiles.ServicioRegionales.EnviarDataNode:input_type -> protofiles.DataNodeRequest
	6,  // 2: protofiles.ServicioRegionales.ConsultarEstadoDiaboromon:input_type -> protofiles.EstadoRequest
	6,  // 3: protofiles.ServicioRegionales.ConsultarEstadoTerminacion:input_type -> protofiles.EstadoRequest
	2,  // 4: protofiles.ServicioDataNodes.GuardarData:input_type -> protofiles.DataNodeRequest
	8,  // 5: protofiles.ServicioDataNodes.NotificarTerminacion:input_type -> protofiles.TerminarRequest
	12, // 6: protofiles.ServicioDataNodes.ConsultarAtributos:input_type -> protofiles.DataNodeQuery
	4,  // 7: protofiles.ServicioTai.SolicitarDatos:input_type -> protofiles.TaiRequest
	4,  // 8: protofiles.ServicioTai.Atacar:input_type -> protofiles.TaiRequest
	6,  // 9: protofiles.ServicioTai.EnviarEstadoDiaboromon:input_type -> protofiles.EstadoRequest
	8,  // 10: protofiles.ServicioTai.NotificarTerminacion:input_type -> protofiles.TerminarRequest
	10, // 11: protofiles.ServicioTai.NotificarEstado:input_type -> protofiles.EstadoTaiRequest
	14, // 12: protofiles.ServicioDiaboromon.RecibirNotificacion:input_type -> protofiles.DiaboromonRequest
	1,  // 13: protofiles.ServicioRegionales.EnviarEstadoDigimon:output_type -> protofiles.DigimonResponse
	3,  // 14: protofiles.ServicioRegionales.EnviarDataNode:output_type -> protofiles.DataNodeResponse
	7,  // 15: protofiles.ServicioRegionales.ConsultarEstadoDiaboromon:output_type -> protofiles.EstadoResponse
	7,  // 16: protofiles.ServicioRegionales.ConsultarEstadoTerminacion:output_type -> protofiles.EstadoResponse
	3,  // 17: protofiles.ServicioDataNodes.GuardarData:output_type -> protofiles.DataNodeResponse
	9,  // 18: protofiles.ServicioDataNodes.NotificarTerminacion:output_type -> protofiles.TerminarResponse
	13, // 19: protofiles.ServicioDataNodes.ConsultarAtributos:output_type -> protofiles.DataNodeQueryResponse
	5,  // 20: protofiles.ServicioTai.SolicitarDatos:output_type -> protofiles.TaiResponse
	5,  // 21: protofiles.ServicioTai.Atacar:output_type -> protofiles.TaiResponse
	7,  // 22: protofiles.ServicioTai.EnviarEstadoDiaboromon:output_type -> protofiles.EstadoResponse
	9,  // 23: protofiles.ServicioTai.NotificarTerminacion:output_type -> protofiles.TerminarResponse
	11, // 24: protofiles.ServicioTai.NotificarEstado:output_type -> protofiles.EstadoTaiResponse
	15, // 25: protofiles.ServicioDiaboromon.RecibirNotificacion:output_type -> protofiles.DiaboromonResponse
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_protofiles_proto_proto_init() }
func file_protofiles_proto_proto_init() {
	if File_protofiles_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protofiles_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_protofiles_proto_proto_goTypes,
		DependencyIndexes: file_protofiles_proto_proto_depIdxs,
		MessageInfos:      file_protofiles_proto_proto_msgTypes,
	}.Build()
	File_protofiles_proto_proto = out.File
	file_protofiles_proto_proto_rawDesc = nil
	file_protofiles_proto_proto_goTypes = nil
	file_protofiles_proto_proto_depIdxs = nil
}