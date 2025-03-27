// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tkd/printing/v1/printing.proto

package printingv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PrinterState int32

const (
	PrinterState_PRINTERSTATE_UNSPECIFIED PrinterState = 0
	PrinterState_PRINTERSTATE_IDLE        PrinterState = 1
	PrinterState_PRINTERSTATE_PRINTING    PrinterState = 2
	PrinterState_PRINTERSTATE_STOPPED     PrinterState = 3
)

// Enum value maps for PrinterState.
var (
	PrinterState_name = map[int32]string{
		0: "PRINTERSTATE_UNSPECIFIED",
		1: "PRINTERSTATE_IDLE",
		2: "PRINTERSTATE_PRINTING",
		3: "PRINTERSTATE_STOPPED",
	}
	PrinterState_value = map[string]int32{
		"PRINTERSTATE_UNSPECIFIED": 0,
		"PRINTERSTATE_IDLE":        1,
		"PRINTERSTATE_PRINTING":    2,
		"PRINTERSTATE_STOPPED":     3,
	}
)

func (x PrinterState) Enum() *PrinterState {
	p := new(PrinterState)
	*p = x
	return p
}

func (x PrinterState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrinterState) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_printing_v1_printing_proto_enumTypes[0].Descriptor()
}

func (PrinterState) Type() protoreflect.EnumType {
	return &file_tkd_printing_v1_printing_proto_enumTypes[0]
}

func (x PrinterState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrinterState.Descriptor instead.
func (PrinterState) EnumDescriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{0}
}

type Orientation int32

const (
	Orientation_ORIENTATION_PORTRAIT  Orientation = 0
	Orientation_ORIENTATION_LANDSCAPE Orientation = 1
)

// Enum value maps for Orientation.
var (
	Orientation_name = map[int32]string{
		0: "ORIENTATION_PORTRAIT",
		1: "ORIENTATION_LANDSCAPE",
	}
	Orientation_value = map[string]int32{
		"ORIENTATION_PORTRAIT":  0,
		"ORIENTATION_LANDSCAPE": 1,
	}
)

func (x Orientation) Enum() *Orientation {
	p := new(Orientation)
	*p = x
	return p
}

func (x Orientation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Orientation) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_printing_v1_printing_proto_enumTypes[1].Descriptor()
}

func (Orientation) Type() protoreflect.EnumType {
	return &file_tkd_printing_v1_printing_proto_enumTypes[1]
}

func (x Orientation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Orientation.Descriptor instead.
func (Orientation) EnumDescriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{1}
}

type ColorMode int32

const (
	ColorMode_COLORMODE_AUTO      ColorMode = 0
	ColorMode_COLORMODE_COLOR     ColorMode = 1
	ColorMode_COLORMODE_GRAYSCALE ColorMode = 2
)

// Enum value maps for ColorMode.
var (
	ColorMode_name = map[int32]string{
		0: "COLORMODE_AUTO",
		1: "COLORMODE_COLOR",
		2: "COLORMODE_GRAYSCALE",
	}
	ColorMode_value = map[string]int32{
		"COLORMODE_AUTO":      0,
		"COLORMODE_COLOR":     1,
		"COLORMODE_GRAYSCALE": 2,
	}
)

func (x ColorMode) Enum() *ColorMode {
	p := new(ColorMode)
	*p = x
	return p
}

func (x ColorMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColorMode) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_printing_v1_printing_proto_enumTypes[2].Descriptor()
}

func (ColorMode) Type() protoreflect.EnumType {
	return &file_tkd_printing_v1_printing_proto_enumTypes[2]
}

func (x ColorMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColorMode.Descriptor instead.
func (ColorMode) EnumDescriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{2}
}

type PrintState int32

const (
	PrintState_PRINTSTATE_UNSPECIFIED PrintState = 0
	PrintState_PRINTSTATE_PENDING     PrintState = 1
	PrintState_PRINTSTATE_PRINTING    PrintState = 2
	PrintState_PRINTSTATE_COMPLETED   PrintState = 3
	PrintState_PRINTSTATE_CANCELLED   PrintState = 4
	PrintState_PRINTSTATE_HELD        PrintState = 5
	PrintState_PRINTSTATE_STOPPED     PrintState = 6
)

// Enum value maps for PrintState.
var (
	PrintState_name = map[int32]string{
		0: "PRINTSTATE_UNSPECIFIED",
		1: "PRINTSTATE_PENDING",
		2: "PRINTSTATE_PRINTING",
		3: "PRINTSTATE_COMPLETED",
		4: "PRINTSTATE_CANCELLED",
		5: "PRINTSTATE_HELD",
		6: "PRINTSTATE_STOPPED",
	}
	PrintState_value = map[string]int32{
		"PRINTSTATE_UNSPECIFIED": 0,
		"PRINTSTATE_PENDING":     1,
		"PRINTSTATE_PRINTING":    2,
		"PRINTSTATE_COMPLETED":   3,
		"PRINTSTATE_CANCELLED":   4,
		"PRINTSTATE_HELD":        5,
		"PRINTSTATE_STOPPED":     6,
	}
)

func (x PrintState) Enum() *PrintState {
	p := new(PrintState)
	*p = x
	return p
}

func (x PrintState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrintState) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_printing_v1_printing_proto_enumTypes[3].Descriptor()
}

func (PrintState) Type() protoreflect.EnumType {
	return &file_tkd_printing_v1_printing_proto_enumTypes[3]
}

func (x PrintState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrintState.Descriptor instead.
func (PrintState) EnumDescriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{3}
}

type Printer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name holds the unique name of the printer.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Manufactorer holds the manufactorer of the printer.
	Manufactorer string `protobuf:"bytes,2,opt,name=manufactorer,proto3" json:"manufactorer,omitempty"`
	// Model holds the printer model.
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Location holds an optional location for the printer.
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	// Description might hold an optional description of the printer.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Parameters might hold printing parameters.
	Parameters    map[string]string `protobuf:"bytes,6,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Printer) Reset() {
	*x = Printer{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Printer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Printer) ProtoMessage() {}

func (x *Printer) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Printer.ProtoReflect.Descriptor instead.
func (*Printer) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{0}
}

func (x *Printer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Printer) GetManufactorer() string {
	if x != nil {
		return x.Manufactorer
	}
	return ""
}

func (x *Printer) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Printer) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Printer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Printer) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type Document struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the name of the document to print.
	// It's used to for the long-running operation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ContentType holds the MIME type of the document to be printed.
	// If set to * or */* the document's MIME type will be autodetected by
	// checking the first few bytes.
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Source defines the source of the document content.
	//
	// Types that are valid to be assigned to Source:
	//
	//	*Document_Data
	//	*Document_Url
	//	*Document_FilePath
	Source      isDocument_Source `protobuf_oneof:"source"`
	Orientation Orientation       `protobuf:"varint,6,opt,name=Orientation,proto3,enum=tkd.printing.v1.Orientation" json:"Orientation,omitempty"`
	ColorMode   ColorMode         `protobuf:"varint,7,opt,name=color_mode,json=colorMode,proto3,enum=tkd.printing.v1.ColorMode" json:"color_mode,omitempty"`
	// Printer might be used to specify the printer where the document should be printed.
	// If empty, the services default printer will be used.
	Printer       string `protobuf:"bytes,10,opt,name=printer,proto3" json:"printer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Document) Reset() {
	*x = Document{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{1}
}

func (x *Document) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Document) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Document) GetSource() isDocument_Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Document) GetData() []byte {
	if x != nil {
		if x, ok := x.Source.(*Document_Data); ok {
			return x.Data
		}
	}
	return nil
}

func (x *Document) GetUrl() string {
	if x != nil {
		if x, ok := x.Source.(*Document_Url); ok {
			return x.Url
		}
	}
	return ""
}

func (x *Document) GetFilePath() string {
	if x != nil {
		if x, ok := x.Source.(*Document_FilePath); ok {
			return x.FilePath
		}
	}
	return ""
}

func (x *Document) GetOrientation() Orientation {
	if x != nil {
		return x.Orientation
	}
	return Orientation_ORIENTATION_PORTRAIT
}

func (x *Document) GetColorMode() ColorMode {
	if x != nil {
		return x.ColorMode
	}
	return ColorMode_COLORMODE_AUTO
}

func (x *Document) GetPrinter() string {
	if x != nil {
		return x.Printer
	}
	return ""
}

type isDocument_Source interface {
	isDocument_Source()
}

type Document_Data struct {
	// Data includes the whole document data in this request.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3,oneof"`
}

type Document_Url struct {
	// URL holds a web URL where the document should be fetched.
	Url string `protobuf:"bytes,4,opt,name=url,proto3,oneof"`
}

type Document_FilePath struct {
	// FilePath holds the path to the document on the file-storage.
	// Not all service implementations will support this source type.
	FilePath string `protobuf:"bytes,5,opt,name=file_path,json=filePath,proto3,oneof"`
}

func (*Document_Data) isDocument_Source() {}

func (*Document_Url) isDocument_Source() {}

func (*Document_FilePath) isDocument_Source() {}

type Job struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State         PrintState             `protobuf:"varint,3,opt,name=state,proto3,enum=tkd.printing.v1.PrintState" json:"state,omitempty"`
	Progress      int32                  `protobuf:"varint,4,opt,name=progress,proto3" json:"progress,omitempty"`
	Printer       string                 `protobuf:"bytes,5,opt,name=printer,proto3" json:"printer,omitempty"`
	OperationId   string                 `protobuf:"bytes,6,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{2}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetState() PrintState {
	if x != nil {
		return x.State
	}
	return PrintState_PRINTSTATE_UNSPECIFIED
}

func (x *Job) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Job) GetPrinter() string {
	if x != nil {
		return x.Printer
	}
	return ""
}

func (x *Job) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

type ListJobsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Printers      []string               `protobuf:"bytes,1,rep,name=printers,proto3" json:"printers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{3}
}

func (x *ListJobsRequest) GetPrinters() []string {
	if x != nil {
		return x.Printers
	}
	return nil
}

type ListJobsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jobs          []*Job                 `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{4}
}

func (x *ListJobsResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type PrintDocumentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Message:
	//
	//	*PrintDocumentRequest_Document
	//	*PrintDocumentRequest_Data
	Message       isPrintDocumentRequest_Message `protobuf_oneof:"message"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintDocumentRequest) Reset() {
	*x = PrintDocumentRequest{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintDocumentRequest) ProtoMessage() {}

func (x *PrintDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintDocumentRequest.ProtoReflect.Descriptor instead.
func (*PrintDocumentRequest) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{5}
}

func (x *PrintDocumentRequest) GetMessage() isPrintDocumentRequest_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PrintDocumentRequest) GetDocument() *Document {
	if x != nil {
		if x, ok := x.Message.(*PrintDocumentRequest_Document); ok {
			return x.Document
		}
	}
	return nil
}

func (x *PrintDocumentRequest) GetData() []byte {
	if x != nil {
		if x, ok := x.Message.(*PrintDocumentRequest_Data); ok {
			return x.Data
		}
	}
	return nil
}

type isPrintDocumentRequest_Message interface {
	isPrintDocumentRequest_Message()
}

type PrintDocumentRequest_Document struct {
	Document *Document `protobuf:"bytes,1,opt,name=document,proto3,oneof"`
}

type PrintDocumentRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*PrintDocumentRequest_Document) isPrintDocumentRequest_Message() {}

func (*PrintDocumentRequest_Data) isPrintDocumentRequest_Message() {}

type ListPrintersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PrinterState  PrinterState           `protobuf:"varint,1,opt,name=printer_state,json=printerState,proto3,enum=tkd.printing.v1.PrinterState" json:"printer_state,omitempty"`
	Location      string                 `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPrintersRequest) Reset() {
	*x = ListPrintersRequest{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPrintersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPrintersRequest) ProtoMessage() {}

func (x *ListPrintersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPrintersRequest.ProtoReflect.Descriptor instead.
func (*ListPrintersRequest) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{6}
}

func (x *ListPrintersRequest) GetPrinterState() PrinterState {
	if x != nil {
		return x.PrinterState
	}
	return PrinterState_PRINTERSTATE_UNSPECIFIED
}

func (x *ListPrintersRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type ListPrintersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Printers      []*Printer             `protobuf:"bytes,1,rep,name=printers,proto3" json:"printers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPrintersResponse) Reset() {
	*x = ListPrintersResponse{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPrintersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPrintersResponse) ProtoMessage() {}

func (x *ListPrintersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPrintersResponse.ProtoReflect.Descriptor instead.
func (*ListPrintersResponse) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{7}
}

func (x *ListPrintersResponse) GetPrinters() []*Printer {
	if x != nil {
		return x.Printers
	}
	return nil
}

type PrintOperationState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	State PrintState             `protobuf:"varint,1,opt,name=state,proto3,enum=tkd.printing.v1.PrintState" json:"state,omitempty"`
	// Document holds the document that is being printed.
	// If the document content has been specified using the source.data field,
	// it is omitted to avoid large blobs of data.
	Document      *Document `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintOperationState) Reset() {
	*x = PrintOperationState{}
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintOperationState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintOperationState) ProtoMessage() {}

func (x *PrintOperationState) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_printing_v1_printing_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintOperationState.ProtoReflect.Descriptor instead.
func (*PrintOperationState) Descriptor() ([]byte, []int) {
	return file_tkd_printing_v1_printing_proto_rawDescGZIP(), []int{8}
}

func (x *PrintOperationState) GetState() PrintState {
	if x != nil {
		return x.State
	}
	return PrintState_PRINTSTATE_UNSPECIFIED
}

func (x *PrintOperationState) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

var File_tkd_printing_v1_printing_proto protoreflect.FileDescriptor

const file_tkd_printing_v1_printing_proto_rawDesc = "" +
	"\n" +
	"\x1etkd/printing/v1/printing.proto\x12\x0ftkd.printing.v1\x1a\x1bbuf/validate/validate.proto\x1a\"tkd/longrunning/v1/operation.proto\"\x9e\x02\n" +
	"\aPrinter\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\"\n" +
	"\fmanufactorer\x18\x02 \x01(\tR\fmanufactorer\x12\x14\n" +
	"\x05model\x18\x03 \x01(\tR\x05model\x12\x1a\n" +
	"\blocation\x18\x04 \x01(\tR\blocation\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12H\n" +
	"\n" +
	"parameters\x18\x06 \x03(\v2(.tkd.printing.v1.Printer.ParametersEntryR\n" +
	"parameters\x1a=\n" +
	"\x0fParametersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xbb\x02\n" +
	"\bDocument\x12\x1b\n" +
	"\x04name\x18\x01 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\x04name\x12*\n" +
	"\fcontent_type\x18\x02 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\vcontentType\x12\x14\n" +
	"\x04data\x18\x03 \x01(\fH\x00R\x04data\x12\x12\n" +
	"\x03url\x18\x04 \x01(\tH\x00R\x03url\x12\x1d\n" +
	"\tfile_path\x18\x05 \x01(\tH\x00R\bfilePath\x12>\n" +
	"\vOrientation\x18\x06 \x01(\x0e2\x1c.tkd.printing.v1.OrientationR\vOrientation\x129\n" +
	"\n" +
	"color_mode\x18\a \x01(\x0e2\x1a.tkd.printing.v1.ColorModeR\tcolorMode\x12\x18\n" +
	"\aprinter\x18\n" +
	" \x01(\tR\aprinterB\b\n" +
	"\x06source\"\xb5\x01\n" +
	"\x03Job\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x121\n" +
	"\x05state\x18\x03 \x01(\x0e2\x1b.tkd.printing.v1.PrintStateR\x05state\x12\x1a\n" +
	"\bprogress\x18\x04 \x01(\x05R\bprogress\x12\x18\n" +
	"\aprinter\x18\x05 \x01(\tR\aprinter\x12!\n" +
	"\foperation_id\x18\x06 \x01(\tR\voperationId\"-\n" +
	"\x0fListJobsRequest\x12\x1a\n" +
	"\bprinters\x18\x01 \x03(\tR\bprinters\"<\n" +
	"\x10ListJobsResponse\x12(\n" +
	"\x04jobs\x18\x01 \x03(\v2\x14.tkd.printing.v1.JobR\x04jobs\"x\n" +
	"\x14PrintDocumentRequest\x127\n" +
	"\bdocument\x18\x01 \x01(\v2\x19.tkd.printing.v1.DocumentH\x00R\bdocument\x12\x14\n" +
	"\x04data\x18\x02 \x01(\fH\x00R\x04dataB\x11\n" +
	"\amessage\x12\x06\xfa\xf7\x18\x02\b\x01\"u\n" +
	"\x13ListPrintersRequest\x12B\n" +
	"\rprinter_state\x18\x01 \x01(\x0e2\x1d.tkd.printing.v1.PrinterStateR\fprinterState\x12\x1a\n" +
	"\blocation\x18\x02 \x01(\tR\blocation\"L\n" +
	"\x14ListPrintersResponse\x124\n" +
	"\bprinters\x18\x01 \x03(\v2\x18.tkd.printing.v1.PrinterR\bprinters\"\x7f\n" +
	"\x13PrintOperationState\x121\n" +
	"\x05state\x18\x01 \x01(\x0e2\x1b.tkd.printing.v1.PrintStateR\x05state\x125\n" +
	"\bdocument\x18\x02 \x01(\v2\x19.tkd.printing.v1.DocumentR\bdocument*x\n" +
	"\fPrinterState\x12\x1c\n" +
	"\x18PRINTERSTATE_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11PRINTERSTATE_IDLE\x10\x01\x12\x19\n" +
	"\x15PRINTERSTATE_PRINTING\x10\x02\x12\x18\n" +
	"\x14PRINTERSTATE_STOPPED\x10\x03*B\n" +
	"\vOrientation\x12\x18\n" +
	"\x14ORIENTATION_PORTRAIT\x10\x00\x12\x19\n" +
	"\x15ORIENTATION_LANDSCAPE\x10\x01*M\n" +
	"\tColorMode\x12\x12\n" +
	"\x0eCOLORMODE_AUTO\x10\x00\x12\x13\n" +
	"\x0fCOLORMODE_COLOR\x10\x01\x12\x17\n" +
	"\x13COLORMODE_GRAYSCALE\x10\x02*\xba\x01\n" +
	"\n" +
	"PrintState\x12\x1a\n" +
	"\x16PRINTSTATE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12PRINTSTATE_PENDING\x10\x01\x12\x17\n" +
	"\x13PRINTSTATE_PRINTING\x10\x02\x12\x18\n" +
	"\x14PRINTSTATE_COMPLETED\x10\x03\x12\x18\n" +
	"\x14PRINTSTATE_CANCELLED\x10\x04\x12\x13\n" +
	"\x0fPRINTSTATE_HELD\x10\x05\x12\x16\n" +
	"\x12PRINTSTATE_STOPPED\x10\x062\xe6\x02\n" +
	"\fPrintService\x12[\n" +
	"\fListPrinters\x12$.tkd.printing.v1.ListPrintersRequest\x1a%.tkd.printing.v1.ListPrintersResponse\x12I\n" +
	"\rPrintDocument\x12\x19.tkd.printing.v1.Document\x1a\x1d.tkd.longrunning.v1.Operation\x12]\n" +
	"\x13PrintDocumentStream\x12%.tkd.printing.v1.PrintDocumentRequest\x1a\x1d.tkd.longrunning.v1.Operation(\x01\x12O\n" +
	"\bListJobs\x12 .tkd.printing.v1.ListJobsRequest\x1a!.tkd.printing.v1.ListJobsResponseB\xcb\x01\n" +
	"\x13com.tkd.printing.v1B\rPrintingProtoP\x01ZGgithub.com/tierklinik-dobersberg/apis/gen/go/tkd/printing/v1;printingv1\xa2\x02\x03TPX\xaa\x02\x0fTkd.Printing.V1\xca\x02\x0fTkd\\Printing\\V1\xe2\x02\x1bTkd\\Printing\\V1\\GPBMetadata\xea\x02\x11Tkd::Printing::V1b\x06proto3"

var (
	file_tkd_printing_v1_printing_proto_rawDescOnce sync.Once
	file_tkd_printing_v1_printing_proto_rawDescData []byte
)

func file_tkd_printing_v1_printing_proto_rawDescGZIP() []byte {
	file_tkd_printing_v1_printing_proto_rawDescOnce.Do(func() {
		file_tkd_printing_v1_printing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_printing_v1_printing_proto_rawDesc), len(file_tkd_printing_v1_printing_proto_rawDesc)))
	})
	return file_tkd_printing_v1_printing_proto_rawDescData
}

var file_tkd_printing_v1_printing_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_tkd_printing_v1_printing_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tkd_printing_v1_printing_proto_goTypes = []any{
	(PrinterState)(0),            // 0: tkd.printing.v1.PrinterState
	(Orientation)(0),             // 1: tkd.printing.v1.Orientation
	(ColorMode)(0),               // 2: tkd.printing.v1.ColorMode
	(PrintState)(0),              // 3: tkd.printing.v1.PrintState
	(*Printer)(nil),              // 4: tkd.printing.v1.Printer
	(*Document)(nil),             // 5: tkd.printing.v1.Document
	(*Job)(nil),                  // 6: tkd.printing.v1.Job
	(*ListJobsRequest)(nil),      // 7: tkd.printing.v1.ListJobsRequest
	(*ListJobsResponse)(nil),     // 8: tkd.printing.v1.ListJobsResponse
	(*PrintDocumentRequest)(nil), // 9: tkd.printing.v1.PrintDocumentRequest
	(*ListPrintersRequest)(nil),  // 10: tkd.printing.v1.ListPrintersRequest
	(*ListPrintersResponse)(nil), // 11: tkd.printing.v1.ListPrintersResponse
	(*PrintOperationState)(nil),  // 12: tkd.printing.v1.PrintOperationState
	nil,                          // 13: tkd.printing.v1.Printer.ParametersEntry
	(*v1.Operation)(nil),         // 14: tkd.longrunning.v1.Operation
}
var file_tkd_printing_v1_printing_proto_depIdxs = []int32{
	13, // 0: tkd.printing.v1.Printer.parameters:type_name -> tkd.printing.v1.Printer.ParametersEntry
	1,  // 1: tkd.printing.v1.Document.Orientation:type_name -> tkd.printing.v1.Orientation
	2,  // 2: tkd.printing.v1.Document.color_mode:type_name -> tkd.printing.v1.ColorMode
	3,  // 3: tkd.printing.v1.Job.state:type_name -> tkd.printing.v1.PrintState
	6,  // 4: tkd.printing.v1.ListJobsResponse.jobs:type_name -> tkd.printing.v1.Job
	5,  // 5: tkd.printing.v1.PrintDocumentRequest.document:type_name -> tkd.printing.v1.Document
	0,  // 6: tkd.printing.v1.ListPrintersRequest.printer_state:type_name -> tkd.printing.v1.PrinterState
	4,  // 7: tkd.printing.v1.ListPrintersResponse.printers:type_name -> tkd.printing.v1.Printer
	3,  // 8: tkd.printing.v1.PrintOperationState.state:type_name -> tkd.printing.v1.PrintState
	5,  // 9: tkd.printing.v1.PrintOperationState.document:type_name -> tkd.printing.v1.Document
	10, // 10: tkd.printing.v1.PrintService.ListPrinters:input_type -> tkd.printing.v1.ListPrintersRequest
	5,  // 11: tkd.printing.v1.PrintService.PrintDocument:input_type -> tkd.printing.v1.Document
	9,  // 12: tkd.printing.v1.PrintService.PrintDocumentStream:input_type -> tkd.printing.v1.PrintDocumentRequest
	7,  // 13: tkd.printing.v1.PrintService.ListJobs:input_type -> tkd.printing.v1.ListJobsRequest
	11, // 14: tkd.printing.v1.PrintService.ListPrinters:output_type -> tkd.printing.v1.ListPrintersResponse
	14, // 15: tkd.printing.v1.PrintService.PrintDocument:output_type -> tkd.longrunning.v1.Operation
	14, // 16: tkd.printing.v1.PrintService.PrintDocumentStream:output_type -> tkd.longrunning.v1.Operation
	8,  // 17: tkd.printing.v1.PrintService.ListJobs:output_type -> tkd.printing.v1.ListJobsResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_tkd_printing_v1_printing_proto_init() }
func file_tkd_printing_v1_printing_proto_init() {
	if File_tkd_printing_v1_printing_proto != nil {
		return
	}
	file_tkd_printing_v1_printing_proto_msgTypes[1].OneofWrappers = []any{
		(*Document_Data)(nil),
		(*Document_Url)(nil),
		(*Document_FilePath)(nil),
	}
	file_tkd_printing_v1_printing_proto_msgTypes[5].OneofWrappers = []any{
		(*PrintDocumentRequest_Document)(nil),
		(*PrintDocumentRequest_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_printing_v1_printing_proto_rawDesc), len(file_tkd_printing_v1_printing_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_printing_v1_printing_proto_goTypes,
		DependencyIndexes: file_tkd_printing_v1_printing_proto_depIdxs,
		EnumInfos:         file_tkd_printing_v1_printing_proto_enumTypes,
		MessageInfos:      file_tkd_printing_v1_printing_proto_msgTypes,
	}.Build()
	File_tkd_printing_v1_printing_proto = out.File
	file_tkd_printing_v1_printing_proto_goTypes = nil
	file_tkd_printing_v1_printing_proto_depIdxs = nil
}
