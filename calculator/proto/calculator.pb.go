// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: calculator/proto/calculator.proto

package grpc_course

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_calculator_proto_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_calculator_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x1a,
	0x1a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x76, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x71, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x90, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x61, 0x76, 0x67, 0x2e, 0x41, 0x76,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x76, 0x67, 0x2e, 0x41,
	0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x2e, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x78, 0x2e, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x78, 0x2e, 0x4d, 0x61,
	0x78, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x0a,
	0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x28, 0x0a, 0x03, 0x53, 0x75, 0x6d,
	0x12, 0x0f, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x53, 0x71, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x71,
	0x72, 0x74, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x73, 0x71, 0x72, 0x74, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x7a, 0x61, 0x61, 0x6b, 0x64, 0x61, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_calculator_proto_calculator_proto_goTypes = []interface{}{
	(*AvgRequest)(nil),    // 0: avg.AvgRequest
	(*MaxRequest)(nil),    // 1: max.MaxRequest
	(*PrimeRequest)(nil),  // 2: primes.PrimeRequest
	(*SumRequest)(nil),    // 3: sum.SumRequest
	(*SqrtRequest)(nil),   // 4: sqrt.SqrtRequest
	(*AvgResponse)(nil),   // 5: avg.AvgResponse
	(*MaxReponse)(nil),    // 6: max.MaxReponse
	(*PrimeResponse)(nil), // 7: primes.PrimeResponse
	(*SumResponse)(nil),   // 8: sum.SumResponse
	(*SqrtResponse)(nil),  // 9: sqrt.SqrtResponse
}
var file_calculator_proto_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.RequestAverage:input_type -> avg.AvgRequest
	1, // 1: calculator.CalculatorService.GetMax:input_type -> max.MaxRequest
	2, // 2: calculator.CalculatorService.ManyPrimes:input_type -> primes.PrimeRequest
	3, // 3: calculator.CalculatorService.Sum:input_type -> sum.SumRequest
	4, // 4: calculator.CalculatorService.Sqrt:input_type -> sqrt.SqrtRequest
	5, // 5: calculator.CalculatorService.RequestAverage:output_type -> avg.AvgResponse
	6, // 6: calculator.CalculatorService.GetMax:output_type -> max.MaxReponse
	7, // 7: calculator.CalculatorService.ManyPrimes:output_type -> primes.PrimeResponse
	8, // 8: calculator.CalculatorService.Sum:output_type -> sum.SumResponse
	9, // 9: calculator.CalculatorService.Sqrt:output_type -> sqrt.SqrtResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_calculator_proto_init() }
func file_calculator_proto_calculator_proto_init() {
	if File_calculator_proto_calculator_proto != nil {
		return
	}
	file_calculator_proto_avg_proto_init()
	file_calculator_proto_max_proto_init()
	file_calculator_proto_primes_proto_init()
	file_calculator_proto_sum_proto_init()
	file_calculator_proto_sqrt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_proto_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_calculator_proto_depIdxs,
	}.Build()
	File_calculator_proto_calculator_proto = out.File
	file_calculator_proto_calculator_proto_rawDesc = nil
	file_calculator_proto_calculator_proto_goTypes = nil
	file_calculator_proto_calculator_proto_depIdxs = nil
}
