// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"
	"vxlan-policy-agent/agent_metrics"
)

type TimeMetricsEmitter struct {
	EmitDurationStub        func(string, time.Duration)
	emitDurationMutex       sync.RWMutex
	emitDurationArgsForCall []struct {
		arg1 string
		arg2 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TimeMetricsEmitter) EmitDuration(arg1 string, arg2 time.Duration) {
	fake.emitDurationMutex.Lock()
	fake.emitDurationArgsForCall = append(fake.emitDurationArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("EmitDuration", []interface{}{arg1, arg2})
	fake.emitDurationMutex.Unlock()
	if fake.EmitDurationStub != nil {
		fake.EmitDurationStub(arg1, arg2)
	}
}

func (fake *TimeMetricsEmitter) EmitDurationCallCount() int {
	fake.emitDurationMutex.RLock()
	defer fake.emitDurationMutex.RUnlock()
	return len(fake.emitDurationArgsForCall)
}

func (fake *TimeMetricsEmitter) EmitDurationArgsForCall(i int) (string, time.Duration) {
	fake.emitDurationMutex.RLock()
	defer fake.emitDurationMutex.RUnlock()
	return fake.emitDurationArgsForCall[i].arg1, fake.emitDurationArgsForCall[i].arg2
}

func (fake *TimeMetricsEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitDurationMutex.RLock()
	defer fake.emitDurationMutex.RUnlock()
	return fake.invocations
}

func (fake *TimeMetricsEmitter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ agent_metrics.TimeMetricsEmitter = new(TimeMetricsEmitter)
