package client

import (
	"fmt"
	"time"
)

// stateRefreshFunc is a function responsible for refreshing the item being watched for a state change.
type stateRefreshFunc[T any] func() (result *T, state string, err error)

type StateChangeOps[T any] struct {
	Pending []string
	Target  []string
	Refresh stateRefreshFunc[T]
	Timeout time.Duration
	Delay   time.Duration
}

// WaitForState polls a resource until it reaches the desired state.
// This is useful for resources whose creation is followed by an async process
// such as Data Pool connections or Job completions.
func WaitForState[T any](ops StateChangeOps[T]) (*T, error) {
	start := time.Now()

	for {
		elapsed := time.Since(start)
		if elapsed.Milliseconds() > ops.Timeout.Milliseconds() {
			return nil, fmt.Errorf("timeout waiting for state to change to %q", ops.Target)
		}

		res, currentState, err := ops.Refresh()
		if err != nil {
			return nil, err
		}

		for _, targetState := range ops.Target {
			if currentState == targetState {
				return res, nil
			}
		}

		pendingFound := false

		for _, pendingState := range ops.Pending {
			if currentState == pendingState {
				pendingFound = true
				break
			}
		}

		if !pendingFound {
			return nil, fmt.Errorf("received an unexpected state %s. Expected states are %q", currentState, ops.Target)
		}

		time.Sleep(ops.Delay)
	}
}
