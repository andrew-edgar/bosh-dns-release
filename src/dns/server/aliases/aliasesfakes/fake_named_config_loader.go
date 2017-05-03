// This file was generated by counterfeiter
package aliasesfakes

import (
	"sync"

	"github.com/cloudfoundry/dns-release/src/dns/server/aliases"
)

type FakeNamedConfigLoader struct {
	LoadStub        func(string) (aliases.Config, error)
	loadMutex       sync.RWMutex
	loadArgsForCall []struct {
		arg1 string
	}
	loadReturns struct {
		result1 aliases.Config
		result2 error
	}
	loadReturnsOnCall map[int]struct {
		result1 aliases.Config
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNamedConfigLoader) Load(arg1 string) (aliases.Config, error) {
	fake.loadMutex.Lock()
	ret, specificReturn := fake.loadReturnsOnCall[len(fake.loadArgsForCall)]
	fake.loadArgsForCall = append(fake.loadArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Load", []interface{}{arg1})
	fake.loadMutex.Unlock()
	if fake.LoadStub != nil {
		return fake.LoadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadReturns.result1, fake.loadReturns.result2
}

func (fake *FakeNamedConfigLoader) LoadCallCount() int {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return len(fake.loadArgsForCall)
}

func (fake *FakeNamedConfigLoader) LoadArgsForCall(i int) string {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return fake.loadArgsForCall[i].arg1
}

func (fake *FakeNamedConfigLoader) LoadReturns(result1 aliases.Config, result2 error) {
	fake.LoadStub = nil
	fake.loadReturns = struct {
		result1 aliases.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeNamedConfigLoader) LoadReturnsOnCall(i int, result1 aliases.Config, result2 error) {
	fake.LoadStub = nil
	if fake.loadReturnsOnCall == nil {
		fake.loadReturnsOnCall = make(map[int]struct {
			result1 aliases.Config
			result2 error
		})
	}
	fake.loadReturnsOnCall[i] = struct {
		result1 aliases.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeNamedConfigLoader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNamedConfigLoader) recordInvocation(key string, args []interface{}) {
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

var _ aliases.NamedConfigLoader = new(FakeNamedConfigLoader)
