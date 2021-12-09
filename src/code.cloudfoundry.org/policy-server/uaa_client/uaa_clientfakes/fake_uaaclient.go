// Code generated by counterfeiter. DO NOT EDIT.
package uaa_clientfakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/uaa_client"
)

type FakeUAAClient struct {
	CheckTokenStub        func(string) (uaa_client.CheckTokenResponse, error)
	checkTokenMutex       sync.RWMutex
	checkTokenArgsForCall []struct {
		arg1 string
	}
	checkTokenReturns struct {
		result1 uaa_client.CheckTokenResponse
		result2 error
	}
	checkTokenReturnsOnCall map[int]struct {
		result1 uaa_client.CheckTokenResponse
		result2 error
	}
	GetTokenStub        func() (string, error)
	getTokenMutex       sync.RWMutex
	getTokenArgsForCall []struct {
	}
	getTokenReturns struct {
		result1 string
		result2 error
	}
	getTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) CheckToken(arg1 string) (uaa_client.CheckTokenResponse, error) {
	fake.checkTokenMutex.Lock()
	ret, specificReturn := fake.checkTokenReturnsOnCall[len(fake.checkTokenArgsForCall)]
	fake.checkTokenArgsForCall = append(fake.checkTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CheckTokenStub
	fakeReturns := fake.checkTokenReturns
	fake.recordInvocation("CheckToken", []interface{}{arg1})
	fake.checkTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) CheckTokenCallCount() int {
	fake.checkTokenMutex.RLock()
	defer fake.checkTokenMutex.RUnlock()
	return len(fake.checkTokenArgsForCall)
}

func (fake *FakeUAAClient) CheckTokenCalls(stub func(string) (uaa_client.CheckTokenResponse, error)) {
	fake.checkTokenMutex.Lock()
	defer fake.checkTokenMutex.Unlock()
	fake.CheckTokenStub = stub
}

func (fake *FakeUAAClient) CheckTokenArgsForCall(i int) string {
	fake.checkTokenMutex.RLock()
	defer fake.checkTokenMutex.RUnlock()
	argsForCall := fake.checkTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) CheckTokenReturns(result1 uaa_client.CheckTokenResponse, result2 error) {
	fake.checkTokenMutex.Lock()
	defer fake.checkTokenMutex.Unlock()
	fake.CheckTokenStub = nil
	fake.checkTokenReturns = struct {
		result1 uaa_client.CheckTokenResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) CheckTokenReturnsOnCall(i int, result1 uaa_client.CheckTokenResponse, result2 error) {
	fake.checkTokenMutex.Lock()
	defer fake.checkTokenMutex.Unlock()
	fake.CheckTokenStub = nil
	if fake.checkTokenReturnsOnCall == nil {
		fake.checkTokenReturnsOnCall = make(map[int]struct {
			result1 uaa_client.CheckTokenResponse
			result2 error
		})
	}
	fake.checkTokenReturnsOnCall[i] = struct {
		result1 uaa_client.CheckTokenResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetToken() (string, error) {
	fake.getTokenMutex.Lock()
	ret, specificReturn := fake.getTokenReturnsOnCall[len(fake.getTokenArgsForCall)]
	fake.getTokenArgsForCall = append(fake.getTokenArgsForCall, struct {
	}{})
	stub := fake.GetTokenStub
	fakeReturns := fake.getTokenReturns
	fake.recordInvocation("GetToken", []interface{}{})
	fake.getTokenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) GetTokenCallCount() int {
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return len(fake.getTokenArgsForCall)
}

func (fake *FakeUAAClient) GetTokenCalls(stub func() (string, error)) {
	fake.getTokenMutex.Lock()
	defer fake.getTokenMutex.Unlock()
	fake.GetTokenStub = stub
}

func (fake *FakeUAAClient) GetTokenReturns(result1 string, result2 error) {
	fake.getTokenMutex.Lock()
	defer fake.getTokenMutex.Unlock()
	fake.GetTokenStub = nil
	fake.getTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.getTokenMutex.Lock()
	defer fake.getTokenMutex.Unlock()
	fake.GetTokenStub = nil
	if fake.getTokenReturnsOnCall == nil {
		fake.getTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkTokenMutex.RLock()
	defer fake.checkTokenMutex.RUnlock()
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
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

var _ uaa_client.UAAClient = new(FakeUAAClient)