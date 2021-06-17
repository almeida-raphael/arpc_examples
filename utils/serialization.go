package utils

import (
	"fmt"
	"time"

	"github.com/almeida-raphael/arpc/interfaces"
	"google.golang.org/protobuf/proto"
)

// TestSerialization get serialization and deserialization times
func TestSerialization(serializable interfaces.Serializable) (*time.Duration, *time.Duration, error) {
	serializableLen, err := serializable.MarshalLen()
	if err != nil {
		return nil, nil, err
	}

	buf := make([]byte, serializableLen)
	serializeStart := time.Now()
	serializable.MarshalTo(buf)
	serializationTime := time.Since(serializeStart)

	buffLen := len(buf)
	deserializeStart := time.Now()
	nConsumed, err := serializable.Unmarshal(buf)
	deserializationTime := time.Since(deserializeStart)
	if err != nil {
		return nil, nil, err
	}

	if buffLen != nConsumed {
		return nil, nil, fmt.Errorf("read byte amount is different from buffer size")
	}

	return &serializationTime, &deserializationTime, nil
}

// TestGRPCSerialization get serialization and deserialization times
func TestGRPCSerialization(serializable proto.Message) (*time.Duration, *time.Duration, error) {
	serializeStart := time.Now()
	serialized, err := proto.Marshal(serializable)
	if err != nil {
		return nil, nil, err
	}
	serializationTime := time.Since(serializeStart)

	deserializeStart := time.Now()
	if err := proto.Unmarshal(serialized, serializable); err != nil {
		return nil, nil, err
	}
	deserializationTime := time.Since(deserializeStart)

	return &serializationTime, &deserializationTime, nil
}
