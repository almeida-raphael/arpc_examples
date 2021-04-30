package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main(){
	assetPath := os.Getenv("ASSET_PATH")
	data, err := ioutil.ReadFile(assetPath)
	if err != nil {
		panic(err)
	}

	jsonEncodeStart := time.Now()
	text, err := json.Marshal(string(data))
	jsonEncodeTime := time.Since(jsonEncodeStart)

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(strings.NewReader(string(text)))

	var result = string(make([]byte, len(text)))

	jsonDecodeStart := time.Now()
	err = decoder.Decode(&result)
	jsonDecodeTime := time.Since(jsonDecodeStart)

	if err != nil {
		panic(err)
	}

	fmt.Printf(`
		Metrics:
		Json Serialization Time: %vs
		Json Deserialization Time: %vs
		Serialization size: %db
		Deserialization size: %db
		`,
		(jsonEncodeTime).Seconds(),
		(jsonDecodeTime).Seconds(),
		len(text),
		len(result),
	)
}

