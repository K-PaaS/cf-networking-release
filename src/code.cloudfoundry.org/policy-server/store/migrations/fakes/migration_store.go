// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type MigrationStore struct {
	HasV1MigrationOccurredStub        func() (bool, error)
	hasV1MigrationOccurredMutex       sync.RWMutex
	hasV1MigrationOccurredArgsForCall []struct {
	}
	hasV1MigrationOccurredReturns struct {
		result1 bool
		result2 error
	}
	hasV1MigrationOccurredReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	HasV2MigrationOccurredStub        func() (bool, error)
	hasV2MigrationOccurredMutex       sync.RWMutex
	hasV2MigrationOccurredArgsForCall []struct {
	}
	hasV2MigrationOccurredReturns struct {
		result1 bool
		result2 error
	}
	hasV2MigrationOccurredReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	HasV3MigrationOccurredStub        func() (bool, error)
	hasV3MigrationOccurredMutex       sync.RWMutex
	hasV3MigrationOccurredArgsForCall []struct {
	}
	hasV3MigrationOccurredReturns struct {
		result1 bool
		result2 error
	}
	hasV3MigrationOccurredReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MigrationStore) HasV1MigrationOccurred() (bool, error) {
	fake.hasV1MigrationOccurredMutex.Lock()
	ret, specificReturn := fake.hasV1MigrationOccurredReturnsOnCall[len(fake.hasV1MigrationOccurredArgsForCall)]
	fake.hasV1MigrationOccurredArgsForCall = append(fake.hasV1MigrationOccurredArgsForCall, struct {
	}{})
	stub := fake.HasV1MigrationOccurredStub
	fakeReturns := fake.hasV1MigrationOccurredReturns
	fake.recordInvocation("HasV1MigrationOccurred", []interface{}{})
	fake.hasV1MigrationOccurredMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MigrationStore) HasV1MigrationOccurredCallCount() int {
	fake.hasV1MigrationOccurredMutex.RLock()
	defer fake.hasV1MigrationOccurredMutex.RUnlock()
	return len(fake.hasV1MigrationOccurredArgsForCall)
}

func (fake *MigrationStore) HasV1MigrationOccurredCalls(stub func() (bool, error)) {
	fake.hasV1MigrationOccurredMutex.Lock()
	defer fake.hasV1MigrationOccurredMutex.Unlock()
	fake.HasV1MigrationOccurredStub = stub
}

func (fake *MigrationStore) HasV1MigrationOccurredReturns(result1 bool, result2 error) {
	fake.hasV1MigrationOccurredMutex.Lock()
	defer fake.hasV1MigrationOccurredMutex.Unlock()
	fake.HasV1MigrationOccurredStub = nil
	fake.hasV1MigrationOccurredReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MigrationStore) HasV1MigrationOccurredReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasV1MigrationOccurredMutex.Lock()
	defer fake.hasV1MigrationOccurredMutex.Unlock()
	fake.HasV1MigrationOccurredStub = nil
	if fake.hasV1MigrationOccurredReturnsOnCall == nil {
		fake.hasV1MigrationOccurredReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasV1MigrationOccurredReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MigrationStore) HasV2MigrationOccurred() (bool, error) {
	fake.hasV2MigrationOccurredMutex.Lock()
	ret, specificReturn := fake.hasV2MigrationOccurredReturnsOnCall[len(fake.hasV2MigrationOccurredArgsForCall)]
	fake.hasV2MigrationOccurredArgsForCall = append(fake.hasV2MigrationOccurredArgsForCall, struct {
	}{})
	stub := fake.HasV2MigrationOccurredStub
	fakeReturns := fake.hasV2MigrationOccurredReturns
	fake.recordInvocation("HasV2MigrationOccurred", []interface{}{})
	fake.hasV2MigrationOccurredMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MigrationStore) HasV2MigrationOccurredCallCount() int {
	fake.hasV2MigrationOccurredMutex.RLock()
	defer fake.hasV2MigrationOccurredMutex.RUnlock()
	return len(fake.hasV2MigrationOccurredArgsForCall)
}

func (fake *MigrationStore) HasV2MigrationOccurredCalls(stub func() (bool, error)) {
	fake.hasV2MigrationOccurredMutex.Lock()
	defer fake.hasV2MigrationOccurredMutex.Unlock()
	fake.HasV2MigrationOccurredStub = stub
}

func (fake *MigrationStore) HasV2MigrationOccurredReturns(result1 bool, result2 error) {
	fake.hasV2MigrationOccurredMutex.Lock()
	defer fake.hasV2MigrationOccurredMutex.Unlock()
	fake.HasV2MigrationOccurredStub = nil
	fake.hasV2MigrationOccurredReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MigrationStore) HasV2MigrationOccurredReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasV2MigrationOccurredMutex.Lock()
	defer fake.hasV2MigrationOccurredMutex.Unlock()
	fake.HasV2MigrationOccurredStub = nil
	if fake.hasV2MigrationOccurredReturnsOnCall == nil {
		fake.hasV2MigrationOccurredReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasV2MigrationOccurredReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MigrationStore) HasV3MigrationOccurred() (bool, error) {
	fake.hasV3MigrationOccurredMutex.Lock()
	ret, specificReturn := fake.hasV3MigrationOccurredReturnsOnCall[len(fake.hasV3MigrationOccurredArgsForCall)]
	fake.hasV3MigrationOccurredArgsForCall = append(fake.hasV3MigrationOccurredArgsForCall, struct {
	}{})
	stub := fake.HasV3MigrationOccurredStub
	fakeReturns := fake.hasV3MigrationOccurredReturns
	fake.recordInvocation("HasV3MigrationOccurred", []interface{}{})
	fake.hasV3MigrationOccurredMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MigrationStore) HasV3MigrationOccurredCallCount() int {
	fake.hasV3MigrationOccurredMutex.RLock()
	defer fake.hasV3MigrationOccurredMutex.RUnlock()
	return len(fake.hasV3MigrationOccurredArgsForCall)
}

func (fake *MigrationStore) HasV3MigrationOccurredCalls(stub func() (bool, error)) {
	fake.hasV3MigrationOccurredMutex.Lock()
	defer fake.hasV3MigrationOccurredMutex.Unlock()
	fake.HasV3MigrationOccurredStub = stub
}

func (fake *MigrationStore) HasV3MigrationOccurredReturns(result1 bool, result2 error) {
	fake.hasV3MigrationOccurredMutex.Lock()
	defer fake.hasV3MigrationOccurredMutex.Unlock()
	fake.HasV3MigrationOccurredStub = nil
	fake.hasV3MigrationOccurredReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MigrationStore) HasV3MigrationOccurredReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasV3MigrationOccurredMutex.Lock()
	defer fake.hasV3MigrationOccurredMutex.Unlock()
	fake.HasV3MigrationOccurredStub = nil
	if fake.hasV3MigrationOccurredReturnsOnCall == nil {
		fake.hasV3MigrationOccurredReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasV3MigrationOccurredReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MigrationStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.hasV1MigrationOccurredMutex.RLock()
	defer fake.hasV1MigrationOccurredMutex.RUnlock()
	fake.hasV2MigrationOccurredMutex.RLock()
	defer fake.hasV2MigrationOccurredMutex.RUnlock()
	fake.hasV3MigrationOccurredMutex.RLock()
	defer fake.hasV3MigrationOccurredMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MigrationStore) recordInvocation(key string, args []interface{}) {
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
