//go:build darwin && !cgo
// +build darwin,!cgo

package cpu

import "errors"

// CPU counters for darwin is unavailable without cgo.

func Get() (*Stats, error) {
	return nil, errors.New("cpu statistics not implemented for darwin nocgo")
}

// Stats represents cpu statistics
type Stats struct {
	User, System, Idle, Nice, Total uint64
}
