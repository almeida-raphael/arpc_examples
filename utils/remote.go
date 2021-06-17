package utils

import (
	"errors"
	"fmt"
	"log"

	arpcerrors "github.com/almeida-raphael/arpc/errors"
	"google.golang.org/grpc/status"
)

// HandleRemoteError remote error handler helper
func HandleRemoteError(err error) bool {
	if err != nil {
		if errors.Is(err, &arpcerrors.Remote{}) {
			fmt.Printf("Remote Error: %v", err)
		} else if errStatus, isStatus := status.FromError(err); isStatus {
			fmt.Printf(
				"gRPC Error: %v; Code: %s; Message: %s",
				err,
				errStatus.Code().String(),
				errStatus.Message(),
			)
		} else {
			log.Fatal(err)
		}
		return true
	}

	return false
}
