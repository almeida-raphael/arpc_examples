package utils

import (
	"io/ioutil"
	"os"
)

// LoadAsset loads an asset using env path
func LoadAsset()[]byte{
	assetPath := os.Getenv("ASSET_PATH")
	data, err := ioutil.ReadFile(assetPath)
	if err != nil {
		panic(err)
	}
	return data
}
