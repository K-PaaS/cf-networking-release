// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"
)

type MetricsSender struct {
	IncrementCounterStub        func(string)
	incrementCounterMutex       sync.RWMutex
	incrementCounterArgsForCall []struct {
		arg1 string
	}
	SendDurationStub        func(string, time.Duration)
	sendDurationMutex       sync.RWMutex
	sendDurationArgsForCall []struct {
		arg1 string
		arg2 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetricsSender) IncrementCounter(arg1 string) {
	fake.incrementCounterMutex.Lock()
	fake.incrementCounterArgsForCall = append(fake.incrementCounterArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IncrementCounterStub
	fake.recordInvocation("IncrementCounter", []interface{}{arg1})
	fake.incrementCounterMutex.Unlock()
	if stub != nil {
		fake.IncrementCounterStub(arg1)
	}
}

func (fake *MetricsSender) IncrementCounterCallCount() int {
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	return len(fake.incrementCounterArgsForCall)
}

func (fake *MetricsSender) IncrementCounterCalls(stub func(string)) {
	fake.incrementCounterMutex.Lock()
	defer fake.incrementCounterMutex.Unlock()
	fake.IncrementCounterStub = stub
}

func (fake *MetricsSender) IncrementCounterArgsForCall(i int) string {
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	argsForCall := fake.incrementCounterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MetricsSender) SendDuration(arg1 string, arg2 time.Duration) {
	fake.sendDurationMutex.Lock()
	fake.sendDurationArgsForCall = append(fake.sendDurationArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	stub := fake.SendDurationStub
	fake.recordInvocation("SendDuration", []interface{}{arg1, arg2})
	fake.sendDurationMutex.Unlock()
	if stub != nil {
		fake.SendDurationStub(arg1, arg2)
	}
}

func (fake *MetricsSender) SendDurationCallCount() int {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return len(fake.sendDurationArgsForCall)
}

func (fake *MetricsSender) SendDurationCalls(stub func(string, time.Duration)) {
	fake.sendDurationMutex.Lock()
	defer fake.sendDurationMutex.Unlock()
	fake.SendDurationStub = stub
}

func (fake *MetricsSender) SendDurationArgsForCall(i int) (string, time.Duration) {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	argsForCall := fake.sendDurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *MetricsSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetricsSender) recordInvocation(key string, args []interface{}) {
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
