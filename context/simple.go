package main

import (
	"context"
	"fmt"
)

var (
	ServiceAbortCtx, ServiceAbortFunc = context.WithCancel(context.Background())
)

func isServiceAbort(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}

func ResetServiceContext() {
	ServiceAbortCtx, ServiceAbortFunc = context.WithCancel(context.Background())
}

func main() {

	fmt.Printf("service is abort: %v\n", isServiceAbort(ServiceAbortCtx))
	ServiceAbortFunc()
	fmt.Printf("service is abort: %v\n", isServiceAbort(ServiceAbortCtx))
	ResetServiceContext()
	fmt.Printf("service is abort: %v\n", isServiceAbort(ServiceAbortCtx))
}
