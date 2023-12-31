// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type SecurityGroupCLIAdapter struct {
	SecurityGroupStub        func(string) (string, error)
	securityGroupMutex       sync.RWMutex
	securityGroupArgsForCall []struct {
		arg1 string
	}
	securityGroupReturns struct {
		result1 string
		result2 error
	}
	securityGroupReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SecurityGroupCLIAdapter) SecurityGroup(arg1 string) (string, error) {
	fake.securityGroupMutex.Lock()
	ret, specificReturn := fake.securityGroupReturnsOnCall[len(fake.securityGroupArgsForCall)]
	fake.securityGroupArgsForCall = append(fake.securityGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SecurityGroupStub
	fakeReturns := fake.securityGroupReturns
	fake.recordInvocation("SecurityGroup", []interface{}{arg1})
	fake.securityGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SecurityGroupCLIAdapter) SecurityGroupCallCount() int {
	fake.securityGroupMutex.RLock()
	defer fake.securityGroupMutex.RUnlock()
	return len(fake.securityGroupArgsForCall)
}

func (fake *SecurityGroupCLIAdapter) SecurityGroupCalls(stub func(string) (string, error)) {
	fake.securityGroupMutex.Lock()
	defer fake.securityGroupMutex.Unlock()
	fake.SecurityGroupStub = stub
}

func (fake *SecurityGroupCLIAdapter) SecurityGroupArgsForCall(i int) string {
	fake.securityGroupMutex.RLock()
	defer fake.securityGroupMutex.RUnlock()
	argsForCall := fake.securityGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SecurityGroupCLIAdapter) SecurityGroupReturns(result1 string, result2 error) {
	fake.securityGroupMutex.Lock()
	defer fake.securityGroupMutex.Unlock()
	fake.SecurityGroupStub = nil
	fake.securityGroupReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *SecurityGroupCLIAdapter) SecurityGroupReturnsOnCall(i int, result1 string, result2 error) {
	fake.securityGroupMutex.Lock()
	defer fake.securityGroupMutex.Unlock()
	fake.SecurityGroupStub = nil
	if fake.securityGroupReturnsOnCall == nil {
		fake.securityGroupReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.securityGroupReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *SecurityGroupCLIAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.securityGroupMutex.RLock()
	defer fake.securityGroupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SecurityGroupCLIAdapter) recordInvocation(key string, args []interface{}) {
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
