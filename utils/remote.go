package utils

import (
	"errors"
	"fmt"
	arpcerrors "github.com/almeida-raphael/arpc/errors"
	"log"
)

// HandleRemoteError remote error handler helper
func HandleRemoteError(err error)bool{
	if err != nil {
		if errors.Is(err, &arpcerrors.Remote{}) {
			fmt.Printf("Remote Error: %v", err)
		}else{
			log.Fatal(err)
		}
		return true
	}

	return false
}