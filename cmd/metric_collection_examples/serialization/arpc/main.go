package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typebool"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typefloat32"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typefloat64"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typeint32"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typeint64"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typetext"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint32"
	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint64"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typebinary"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	trials := utils.Atoi(os.Getenv("TRIALS"))
	value := utils.Atoi(os.Getenv("VALUE"))
	path := "serialization_results/aRPC/%s_" + fmt.Sprintf("%d.json", value)

	var data interfaces.Serializable = &typebinary.Data{Value: []byte(utils.GenerateString(value))}
	err := utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "binary"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typebool.Data{Value: rand.Float32() >= 0.5}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "boolean"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typefloat32.Data{Value: rand.Float32()}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "float32"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typefloat64.Data{Value: rand.Float64()}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "float64"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeint32.Data{Value: int32(rand.Uint32())}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "int32"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeint64.Data{Value: int64(rand.Uint64())}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "int64"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typetext.Data{Value: utils.GenerateString(value)}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "string"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeuint32.Data{Value: rand.Uint32()}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "int32"))
	if err != nil {
		log.Fatal(err)
	}

	data = &typeuint64.Data{Value: rand.Uint64()}
	err = utils.ExtractSerializationMetrics(data, trials, fmt.Sprintf(path, "int64"))
	if err != nil {
		log.Fatal(err)
	}
}
