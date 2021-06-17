package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"sync/atomic"
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
	for idxSamples := range samples {
		fmt.Printf("Running Sample %d\n", idxSamples+1)

		aRPCTestResults := make([]time.Duration, runTrialsCount)
		var err error
		for idxTrial := range aRPCTestResults {
			fmt.Printf("Running aRPC Call Trial %d\n", idxTrial+1)
			rpcStartTime := time.Now()
			response, err = rpcFunction(request)
			elapsedTime := time.Since(rpcStartTime)
			if HandleRemoteError(err) {
				return err
			}
			aRPCTestResults[idxTrial] = elapsedTime
		}

		serializationTestResults, err := runSerializationTest(
			runTrialsCount, request, response,
		)
		if err != nil {
			return err
		}

		samples[idxSamples] = sample{
			SerializationTests:    serializationTestResults,
			ARPCFullTestDurations: aRPCTestResults,
		}
	}

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
	for idxSamples := range samples {
		fmt.Printf("Running Sample %d\n", idxSamples+1)

		aRPCTestResults := make([]time.Duration, runTrialsCount)
		var err error
		for idxTrial := range aRPCTestResults {
			fmt.Printf("Running gRPC Call Trial %d\n", idxTrial+1)
			rpcStartTime := time.Now()
			response, err = rpcFunction(request)
			elapsedTime := time.Since(rpcStartTime)
			if HandleRemoteError(err) {
				return err
			}
			aRPCTestResults[idxTrial] = elapsedTime
		}

		serializationTestResults, err := runGRPCSerializationTest(
			runTrialsCount, request, response,
		)
		if err != nil {
			return err
		}

		samples[idxSamples] = sample{
			SerializationTests:    serializationTestResults,
			ARPCFullTestDurations: aRPCTestResults,
		}
	}

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

func saveServerMetrics(
	sampleCount, trialsCount int, currentMetrics *[][]time.Duration, currentMetricsMutex *sync.Mutex,
	currentExecution *int32, executionTime time.Duration, saveFilePath string,
) error {
	sample := int(*currentExecution) / trialsCount
	trial := int(*currentExecution) % trialsCount

	fmt.Printf("saving sample %d and trial %d\n", sample+1, trial+1)

	if *currentMetrics == nil {
		currentMetricsMutex.Lock()
		*currentMetrics = initServerMetricsDataCollector(sampleCount, trialsCount)
		currentMetricsMutex.Unlock()
	}

	currentMetricsMutex.Lock()
	(*currentMetrics)[sample][trial] = executionTime
	currentMetricsMutex.Unlock()

	if sample == sampleCount-1 && trial == trialsCount-1 {
		currentMetricsMutex.Lock()
		atomic.StoreInt32(currentExecution, -1) // This is -1 because CollectServerMetrics will sum 1 on defer
		err := saveJSON(currentMetrics, fmt.Sprintf(saveFilePath, time.Now().UnixNano()))
		*currentMetrics = initServerMetricsDataCollector(sampleCount, trialsCount)
		currentMetricsMutex.Unlock()
		return err
	}

	return nil
}

// CollectServerMetrics Wraps an server RPC function and collect it's execution time metrics
func CollectServerMetrics(
	runSampleCount, runTrialsCount int, rpcFunction func(interfaces.Serializable) (interfaces.Serializable, error),
	saveFilePath string,
) func(interfaces.Serializable) (interfaces.Serializable, error) {
	var executionCounter int32 = 0
	var metricsRefMu sync.Mutex
	metricsRef := initServerMetricsDataCollector(runSampleCount, runTrialsCount)

	return func(request interfaces.Serializable) (interfaces.Serializable, error) {
		defer atomic.AddInt32(&executionCounter, 1)

		executionStartTime := time.Now()
		response, err := rpcFunction(request)
		if err != nil {
			return nil, err
		}
		executionElapsedTime := time.Since(executionStartTime)
		if err := saveServerMetrics(
			runSampleCount, runTrialsCount, &metricsRef, &metricsRefMu, &executionCounter, executionElapsedTime,
			saveFilePath,
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
	var executionCounter int32 = 0
	var metricsRefMu sync.Mutex
	metricsRef := initServerMetricsDataCollector(runSampleCount, runTrialsCount)

	return func(request proto.Message) (proto.Message, error) {
		defer atomic.AddInt32(&executionCounter, 1)

		executionStartTime := time.Now()
		response, err := rpcFunction(request)
		if err != nil {
			return nil, err
		}
		executionElapsedTime := time.Since(executionStartTime)
		if err := saveServerMetrics(
			runSampleCount, runTrialsCount, &metricsRef, &metricsRefMu, &executionCounter, executionElapsedTime,
			saveFilePath,
		); err != nil {
			return nil, err
		}

		return response, nil
	}
}
