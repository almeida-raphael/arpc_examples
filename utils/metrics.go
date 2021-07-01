package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc/interfaces"
)

type serializationTest struct {
	RequestSerializationTime   time.Duration `json:"request_serialization_time"`
	RequestDeserializationTime time.Duration `json:"request_deserialization_time"`

	ResponseSerializationTime   time.Duration `json:"response_serialization_time"`
	ResponseDeserializationTime time.Duration `json:"response_deserialization_time"`
}

type sample struct {
	SerializationTests    []serializationTest `json:"serialization_tests"`
	ARPCFullTestDurations []time.Duration     `json:"full_test_durations"`
}

type test struct {
	RequestSize  int      `json:"request_size"`
	ResponseSize int      `json:"response_size"`
	Samples      []sample `json:"samples"`
}

type SerializationComparison struct {
	SerializationTime   []time.Duration `json:"serialization_time"`
	DeserializationTime []time.Duration `json:"deserialization_time"`
	Size                []int           `json:"size"`
}

func saveJSON(data interface{}, jsonPath string) error {
	resultData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Dir(jsonPath), os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(jsonPath, resultData, 0664)
	if err != nil {
		return err
	}
	return nil
}

func ExtractSerializationMetrics(data interfaces.Serializable, trials int, jsonPath string) error {
	results := SerializationComparison{
		SerializationTime:   make([]time.Duration, trials),
		DeserializationTime: make([]time.Duration, trials),
		Size:                make([]int, trials),
	}
	for i := 0; i < trials; i++ {
		serializationTime, deserializationTime, err := TestSerialization(data)
		if err != nil {
			log.Fatal(err)
		}
		size, err := data.MarshalLen()
		if err != nil {
			log.Fatal(err)
		}
		results.SerializationTime[i] = *serializationTime
		results.DeserializationTime[i] = *deserializationTime
		results.Size[i] = size
	}

	return saveJSON(results, jsonPath)
}

func ExtractGRPCSerializationMetrics(data proto.Message, trials int, jsonPath string) error {
	results := SerializationComparison{
		SerializationTime:   make([]time.Duration, trials),
		DeserializationTime: make([]time.Duration, trials),
		Size:                make([]int, trials),
	}
	for i := 0; i < trials; i++ {
		serializationTime, deserializationTime, err := TestGRPCSerialization(data)
		if err != nil {
			log.Fatal(err)
		}
		results.SerializationTime[i] = *serializationTime
		results.DeserializationTime[i] = *deserializationTime
		results.Size[i] = proto.Size(data)
	}

	return saveJSON(results, jsonPath)
}

func runSerializationTest(
	runTrialsCount int, request, response interfaces.Serializable,
) ([]serializationTest, error) {
	serializationSamples := make([]serializationTest, runTrialsCount)
	for idxTrial := range serializationSamples {
		fmt.Printf("Running Serialization Trial %d\n", idxTrial+1)
		requestSerializationTime, requestDeserializationTime, err := TestSerialization(request)
		if err != nil {
			return nil, err
		}
		responseSerializationTime, responseDeserializationTime, err := TestSerialization(
			response,
		)
		if err != nil {
			return nil, err
		}
		serializationSamples[idxTrial] = serializationTest{
			RequestSerializationTime:    *requestSerializationTime,
			RequestDeserializationTime:  *requestDeserializationTime,
			ResponseSerializationTime:   *responseSerializationTime,
			ResponseDeserializationTime: *responseDeserializationTime,
		}
	}
	return serializationSamples, nil
}

func runGRPCSerializationTest(
	runTrialsCount int, request, response proto.Message,
) ([]serializationTest, error) {
	serializationSamples := make([]serializationTest, runTrialsCount)
	for idxTrial := range serializationSamples {
		fmt.Printf("Running Serialization Trial %d\n", idxTrial+1)
		requestSerializationTime, requestDeserializationTime, err := TestGRPCSerialization(request)
		if err != nil {
			return nil, err
		}
		responseSerializationTime, responseDeserializationTime, err := TestGRPCSerialization(
			response,
		)
		if err != nil {
			return nil, err
		}
		serializationSamples[idxTrial] = serializationTest{
			RequestSerializationTime:    *requestSerializationTime,
			RequestDeserializationTime:  *requestDeserializationTime,
			ResponseSerializationTime:   *responseSerializationTime,
			ResponseDeserializationTime: *responseDeserializationTime,
		}
	}
	return serializationSamples, nil
}

// RunClientRPCAndCollectMetrics Executes an RPC call for n samples and m trials for each sample then saves it's metrics
func RunClientRPCAndCollectMetrics(
	runSampleCount, runTrialsCount int, request interfaces.Serializable,
	rpcFunction func(interfaces.Serializable) (interfaces.Serializable, error),
	saveFilePath string,
) error {
	samples := make([]sample, runSampleCount)
	var response interfaces.Serializable

	wg := sync.WaitGroup{}
	wg.Add(len(samples))
	mu := sync.Mutex{}

	for idxSamples := range samples {
		_idxSamples := idxSamples
		go func() {
			fmt.Printf("Running Sample %d\n", _idxSamples+1)

			aRPCTestResults := make([]time.Duration, runTrialsCount)
			var err error
			for idxTrial := range aRPCTestResults {
				fmt.Printf("Running aRPC Call Trial %d\n", idxTrial+1)
				rpcStartTime := time.Now()
				response, err = rpcFunction(request)
				elapsedTime := time.Since(rpcStartTime)
				if HandleRemoteError(err) {
					log.Fatal(err)
				}
				aRPCTestResults[idxTrial] = elapsedTime
			}

			mu.Lock()
			serializationTestResults, err := runSerializationTest(
				runTrialsCount, request, response,
			)
			if err != nil {
				log.Fatal(err)
			}

			samples[_idxSamples] = sample{
				SerializationTests:    serializationTestResults,
				ARPCFullTestDurations: aRPCTestResults,
			}
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	requestSize, err := request.MarshalLen()
	if err != nil {
		return err
	}

	responseSize, err := response.MarshalLen()
	if err != nil {
		return err
	}

	testResults := test{
		RequestSize:  requestSize,
		ResponseSize: responseSize,
		Samples:      samples,
	}
	return saveJSON(testResults, saveFilePath)
}

// RunGRPCClientRPCAndCollectMetrics Executes an RPC call for n samples and m trials for each sample then saves it's metrics
func RunGRPCClientRPCAndCollectMetrics(
	runSampleCount, runTrialsCount int, request proto.Message,
	rpcFunction func(proto.Message) (proto.Message, error),
	saveFilePath string,
) error {
	samples := make([]sample, runSampleCount)
	var response proto.Message

	wg := sync.WaitGroup{}
	wg.Add(len(samples))
	mu := sync.Mutex{}

	for idxSamples := range samples {
		_idxSamples := idxSamples

		_request := proto.Clone(request)
		go func() {
			fmt.Printf("Running Sample %d\n", _idxSamples+1)

			aRPCTestResults := make([]time.Duration, runTrialsCount)
			var err error
			for idxTrial := range aRPCTestResults {
				fmt.Printf("Running gRPC Call Trial %d\n", idxTrial+1)
				rpcStartTime := time.Now()
				response, err = rpcFunction(_request)
				elapsedTime := time.Since(rpcStartTime)
				if HandleRemoteError(err) {
					log.Fatal(err)
				}
				aRPCTestResults[idxTrial] = elapsedTime
			}

			mu.Lock()
			serializationTestResults, err := runGRPCSerializationTest(
				runTrialsCount, request, response,
			)
			if err != nil {
				log.Fatal(err)
			}

			samples[_idxSamples] = sample{
				SerializationTests:    serializationTestResults,
				ARPCFullTestDurations: aRPCTestResults,
			}
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	testResults := test{
		RequestSize:  proto.Size(request),
		ResponseSize: proto.Size(response),
		Samples:      samples,
	}
	return saveJSON(testResults, saveFilePath)
}

func initServerMetricsDataCollector(runSampleCount, runTrialsCount int) [][]time.Duration {
	samples := make([][]time.Duration, runSampleCount)
	for idx := 0; idx < runSampleCount; idx++ {
		samples[idx] = make([]time.Duration, runTrialsCount)
	}
	return samples
}

var executionCounter = 0
var metricsRefMu sync.Mutex

func saveServerMetrics(
	sampleCount, trialsCount int, currentMetrics *[][]time.Duration, executionTime time.Duration, saveFilePath string,
) error {
	metricsRefMu.Lock()
	defer metricsRefMu.Unlock()

	if *currentMetrics == nil {
		*currentMetrics = initServerMetricsDataCollector(sampleCount, trialsCount)
	}
	(*currentMetrics)[executionCounter/trialsCount][executionCounter%trialsCount] = executionTime

	executionCounter += 1
	fmt.Printf("saving request %d\n", executionCounter)
	if executionCounter == sampleCount*trialsCount {
		executionCounter = 0
		err := saveJSON(currentMetrics, fmt.Sprintf(saveFilePath, time.Now().UnixNano()))
		*currentMetrics = nil
		return err
	}

	return nil
}

// CollectServerMetrics Wraps an server RPC function and collect it's execution time metrics
func CollectServerMetrics(
	runSampleCount, runTrialsCount int, rpcFunction func(interfaces.Serializable) (interfaces.Serializable, error),
	saveFilePath string,
) func(interfaces.Serializable) (interfaces.Serializable, error) {
	metricsRef := initServerMetricsDataCollector(runSampleCount, runTrialsCount)

	return func(request interfaces.Serializable) (interfaces.Serializable, error) {
		executionStartTime := time.Now()
		response, err := rpcFunction(request)
		if err != nil {
			return nil, err
		}
		executionElapsedTime := time.Since(executionStartTime)
		if err := saveServerMetrics(
			runSampleCount, runTrialsCount, &metricsRef, executionElapsedTime, saveFilePath,
		); err != nil {
			return nil, err
		}

		return response, nil
	}
}

// CollectGRPCServerMetrics Wraps an server RPC function and collect it's execution time metrics
func CollectGRPCServerMetrics(
	runSampleCount, runTrialsCount int, rpcFunction func(proto.Message) (proto.Message, error),
	saveFilePath string,
) func(proto.Message) (proto.Message, error) {
	metricsRef := initServerMetricsDataCollector(runSampleCount, runTrialsCount)

	return func(request proto.Message) (proto.Message, error) {
		executionStartTime := time.Now()
		response, err := rpcFunction(request)
		if err != nil {
			return nil, err
		}
		executionElapsedTime := time.Since(executionStartTime)
		if err := saveServerMetrics(
			runSampleCount, runTrialsCount, &metricsRef, executionElapsedTime, saveFilePath,
		); err != nil {
			return nil, err
		}

		return response, nil
	}
}
