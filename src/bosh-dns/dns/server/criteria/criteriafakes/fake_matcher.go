// Code generated by counterfeiter. DO NOT EDIT.
package criteriafakes

import (
	"bosh-dns/dns/server/criteria"
	"bosh-dns/dns/server/record"
	"sync"
)

type FakeMatcher struct {
	MatchStub        func(r *record.Record) bool
	matchMutex       sync.RWMutex
	matchArgsForCall []struct {
		r *record.Record
	}
	matchReturns struct {
		result1 bool
	}
	matchReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMatcher) Match(r *record.Record) bool {
	fake.matchMutex.Lock()
	ret, specificReturn := fake.matchReturnsOnCall[len(fake.matchArgsForCall)]
	fake.matchArgsForCall = append(fake.matchArgsForCall, struct {
		r *record.Record
	}{r})
	fake.recordInvocation("Match", []interface{}{r})
	fake.matchMutex.Unlock()
	if fake.MatchStub != nil {
		return fake.MatchStub(r)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.matchReturns.result1
}

func (fake *FakeMatcher) MatchCallCount() int {
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	return len(fake.matchArgsForCall)
}

func (fake *FakeMatcher) MatchArgsForCall(i int) *record.Record {
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	return fake.matchArgsForCall[i].r
}

func (fake *FakeMatcher) MatchReturns(result1 bool) {
	fake.MatchStub = nil
	fake.matchReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMatcher) MatchReturnsOnCall(i int, result1 bool) {
	fake.MatchStub = nil
	if fake.matchReturnsOnCall == nil {
		fake.matchReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.matchReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMatcher) recordInvocation(key string, args []interface{}) {
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

var _ criteria.Matcher = new(FakeMatcher)
