package utils

import (
	"fmt"
	"github.com/almeida-raphael/arpc/interfaces"
	"time"
)

// TestSerialization get serialization and deserialization times
func TestSerialization(serializable interfaces.Serializable)(*time.Duration, *time.Duration, int, error){
	serializableLen, err := serializable.MarshalLen()
	if err != nil{
		return nil, nil, 0, err
	}

	buf := make([]byte, serializableLen)
	serializeStart := time.Now()
	serializable.MarshalTo(buf)
	serializationTime := time.Since(serializeStart)

	buffLen := len(buf)
	deserializeStart := time.Now()
	nConsumed, err := serializable.Unmarshal(buf)
	deserializationTime := time.Since(deserializeStart)
	if err != nil{
		return nil, nil, 0, err
	}

	if buffLen != nConsumed{
		return nil, nil, 0, fmt.Errorf("read byte amount is different from buffer size")
	}

	return &serializationTime, &deserializationTime, serializableLen, nil
}