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
	path := "serialization_results/Protobuffers/%s_.json"

	var data proto.Message = &typebinary.TypeBinary{Binary: []byte(utils.GenerateString(value))}
	err := utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "binary"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typebool.TypeBool{Bool: rand.Float32() >= 0.5}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "boolean"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typefloat32.TypeFloat32{Float32: rand.Float32()}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "float32"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typefloat64.TypeFloat64{Float64: rand.Float64()}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "float64"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeint32.TypeInt32{Int32: int32(rand.Uint32())}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "int32"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeint64.TypeInt64{Int64: int64(rand.Uint64())}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "int64"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typetext.TypeText{Text: utils.GenerateString(value)}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "string"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeuint32.TypeUInt32{Uint32: rand.Uint32()}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "int32"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeuint64.TypeUInt64{Uint64: rand.Uint64()}
	err = utils.ExtractGRPCSerializationMetrics(data, trials, fmt.Sprintf(path, "int64"))
	if err != nil {
		log.Fatal(err)
	}
}