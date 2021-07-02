package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc_examples/models/grpc/typebool"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typefloat32"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typefloat64"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typeint32"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typeint64"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typetext"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typeuint32"
	"github.com/almeida-raphael/arpc_examples/models/grpc/typeuint64"

	"github.com/almeida-raphael/arpc_examples/models/grpc/typebinary"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	trials := utils.Atoi(os.Getenv("TRIALS"))
	value := utils.Atoi(os.Getenv("VALUE"))
	path := "serialization_results/Protobuffers/%s.json"

	err := utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typebinary.TypeBinary{Binary: []byte(utils.GenerateString(value))} },
		trials, fmt.Sprintf(path, "binary"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typebool.TypeBool{Bool: rand.Float32() >= 0.5} },
		trials, fmt.Sprintf(path, "boolean"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typefloat32.TypeFloat32{Float32: rand.Float32()} },
		trials, fmt.Sprintf(path, "float32"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typefloat64.TypeFloat64{Float64: rand.Float64()} },
		trials, fmt.Sprintf(path, "float64"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typeint32.TypeInt32{Int32: int32(rand.Uint32())} },
		trials, fmt.Sprintf(path, "int32"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typeint64.TypeInt64{Int64: int64(rand.Uint64())} },
		trials, fmt.Sprintf(path, "int64"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typetext.TypeText{Text: utils.GenerateString(value)} },
		trials, fmt.Sprintf(path, "string"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typeuint32.TypeUInt32{Uint32: rand.Uint32()} },
		trials, fmt.Sprintf(path, "int32"))
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ExtractGRPCSerializationMetrics(
		func() proto.Message { return &typeuint64.TypeUInt64{Uint64: rand.Uint64()} },
		trials, fmt.Sprintf(path, "int64"))
	if err != nil {
		log.Fatal(err)
	}
}
