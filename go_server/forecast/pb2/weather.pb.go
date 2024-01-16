// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: weather.proto

package pb2

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

type ForecastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Days int32  `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
}

func (x *ForecastRequest) Reset() {
	*x = ForecastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastRequest) ProtoMessage() {}

func (x *ForecastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastRequest.ProtoReflect.Descriptor instead.
func (*ForecastRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *ForecastRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ForecastRequest) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

type ForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date               string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Summary            string  `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	TemperatureCelsius float32 `protobuf:"fixed32,3,opt,name=temperatureCelsius,proto3" json:"temperatureCelsius,omitempty"`
}

func (x *ForecastResponse) Reset() {
	*x = ForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastResponse) ProtoMessage() {}

func (x *ForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastResponse.ProtoReflect.Descriptor instead.
func (*ForecastResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *ForecastResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ForecastResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ForecastResponse) GetTemperatureCelsius() float32 {
	if x != nil {
		return x.TemperatureCelsius
	}
	return 0
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64,
	0x61, 0x79, 0x73, 0x22, 0x70, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x43, 0x65, 0x6c, 0x73, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x65,
	0x6c, 0x73, 0x69, 0x75, 0x73, 0x32, 0x5a, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x62,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_weather_proto_goTypes = []interface{}{
	(*ForecastRequest)(nil),  // 0: weather.ForecastRequest
	(*ForecastResponse)(nil), // 1: weather.ForecastResponse
}
var file_weather_proto_depIdxs = []int32{
	0, // 0: weather.WeatherService.StreamForecasts:input_type -> weather.ForecastRequest
	1, // 1: weather.WeatherService.StreamForecasts:output_type -> weather.ForecastResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}