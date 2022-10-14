package main

import (
	"context"
	"fmt"
)

func isServiceAbort(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}

func main() {

	ServiceAbortCtx, ServiceAbortFunc := context.WithCancel(context.Background())

	fmt.Printf("service is abort: %v\n", isServiceAbort(ServiceAbortCtx))
	ServiceAbortFunc()
	fmt.Printf("service is abort: %v\n", isServiceAbort(ServiceAbortCtx))
	fmt.Printf("service is abort: %v\n", isServiceAbort(ServiceAbortCtx))
}
