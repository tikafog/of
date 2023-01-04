// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package eload

import (
	"fmt"
)

const (
	// RunnerStateInit is a RunnerState of type Init.
	RunnerStateInit RunnerState = iota
	// RunnerStateRun is a RunnerState of type Run.
	RunnerStateRun
	// RunnerStateStop is a RunnerState of type Stop.
	RunnerStateStop
	// RunnerStateMax is a RunnerState of type Max.
	RunnerStateMax
)

const _RunnerStateName = "initrunstopmax"

var _RunnerStateMap = map[RunnerState]string{
	RunnerStateInit: _RunnerStateName[0:4],
	RunnerStateRun:  _RunnerStateName[4:7],
	RunnerStateStop: _RunnerStateName[7:11],
	RunnerStateMax:  _RunnerStateName[11:14],
}

// String implements the Stringer interface.
func (x RunnerState) String() string {
	if str, ok := _RunnerStateMap[x]; ok {
		return str
	}
	return fmt.Sprintf("RunnerState(%d)", x)
}

var _RunnerStateValue = map[string]RunnerState{
	_RunnerStateName[0:4]:   RunnerStateInit,
	_RunnerStateName[4:7]:   RunnerStateRun,
	_RunnerStateName[7:11]:  RunnerStateStop,
	_RunnerStateName[11:14]: RunnerStateMax,
}

// ParseRunnerState attempts to convert a string to a RunnerState.
func ParseRunnerState(name string) (RunnerState, error) {
	if x, ok := _RunnerStateValue[name]; ok {
		return x, nil
	}
	return RunnerState(0), fmt.Errorf("%s is not a valid RunnerState", name)
}