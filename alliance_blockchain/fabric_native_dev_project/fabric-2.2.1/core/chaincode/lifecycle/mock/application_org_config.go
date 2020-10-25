// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/msp"
)

type ApplicationOrgConfig struct {
	AnchorPeersStub        func() []*peer.AnchorPeer
	anchorPeersMutex       sync.RWMutex
	anchorPeersArgsForCall []struct {
	}
	anchorPeersReturns struct {
		result1 []*peer.AnchorPeer
	}
	anchorPeersReturnsOnCall map[int]struct {
		result1 []*peer.AnchorPeer
	}
	MSPStub        func() msp.MSP
	mSPMutex       sync.RWMutex
	mSPArgsForCall []struct {
	}
	mSPReturns struct {
		result1 msp.MSP
	}
	mSPReturnsOnCall map[int]struct {
		result1 msp.MSP
	}
	MSPIDStub        func() string
	mSPIDMutex       sync.RWMutex
	mSPIDArgsForCall []struct {
	}
	mSPIDReturns struct {
		result1 string
	}
	mSPIDReturnsOnCall map[int]struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ApplicationOrgConfig) AnchorPeers() []*peer.AnchorPeer {
	fake.anchorPeersMutex.Lock()
	ret, specificReturn := fake.anchorPeersReturnsOnCall[len(fake.anchorPeersArgsForCall)]
	fake.anchorPeersArgsForCall = append(fake.anchorPeersArgsForCall, struct {
	}{})
	fake.recordInvocation("AnchorPeers", []interface{}{})
	fake.anchorPeersMutex.Unlock()
	if fake.AnchorPeersStub != nil {
		return fake.AnchorPeersStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.anchorPeersReturns
	return fakeReturns.result1
}

func (fake *ApplicationOrgConfig) AnchorPeersCallCount() int {
	fake.anchorPeersMutex.RLock()
	defer fake.anchorPeersMutex.RUnlock()
	return len(fake.anchorPeersArgsForCall)
}

func (fake *ApplicationOrgConfig) AnchorPeersCalls(stub func() []*peer.AnchorPeer) {
	fake.anchorPeersMutex.Lock()
	defer fake.anchorPeersMutex.Unlock()
	fake.AnchorPeersStub = stub
}

func (fake *ApplicationOrgConfig) AnchorPeersReturns(result1 []*peer.AnchorPeer) {
	fake.anchorPeersMutex.Lock()
	defer fake.anchorPeersMutex.Unlock()
	fake.AnchorPeersStub = nil
	fake.anchorPeersReturns = struct {
		result1 []*peer.AnchorPeer
	}{result1}
}

func (fake *ApplicationOrgConfig) AnchorPeersReturnsOnCall(i int, result1 []*peer.AnchorPeer) {
	fake.anchorPeersMutex.Lock()
	defer fake.anchorPeersMutex.Unlock()
	fake.AnchorPeersStub = nil
	if fake.anchorPeersReturnsOnCall == nil {
		fake.anchorPeersReturnsOnCall = make(map[int]struct {
			result1 []*peer.AnchorPeer
		})
	}
	fake.anchorPeersReturnsOnCall[i] = struct {
		result1 []*peer.AnchorPeer
	}{result1}
}

func (fake *ApplicationOrgConfig) MSP() msp.MSP {
	fake.mSPMutex.Lock()
	ret, specificReturn := fake.mSPReturnsOnCall[len(fake.mSPArgsForCall)]
	fake.mSPArgsForCall = append(fake.mSPArgsForCall, struct {
	}{})
	fake.recordInvocation("MSP", []interface{}{})
	fake.mSPMutex.Unlock()
	if fake.MSPStub != nil {
		return fake.MSPStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mSPReturns
	return fakeReturns.result1
}

func (fake *ApplicationOrgConfig) MSPCallCount() int {
	fake.mSPMutex.RLock()
	defer fake.mSPMutex.RUnlock()
	return len(fake.mSPArgsForCall)
}

func (fake *ApplicationOrgConfig) MSPCalls(stub func() msp.MSP) {
	fake.mSPMutex.Lock()
	defer fake.mSPMutex.Unlock()
	fake.MSPStub = stub
}

func (fake *ApplicationOrgConfig) MSPReturns(result1 msp.MSP) {
	fake.mSPMutex.Lock()
	defer fake.mSPMutex.Unlock()
	fake.MSPStub = nil
	fake.mSPReturns = struct {
		result1 msp.MSP
	}{result1}
}

func (fake *ApplicationOrgConfig) MSPReturnsOnCall(i int, result1 msp.MSP) {
	fake.mSPMutex.Lock()
	defer fake.mSPMutex.Unlock()
	fake.MSPStub = nil
	if fake.mSPReturnsOnCall == nil {
		fake.mSPReturnsOnCall = make(map[int]struct {
			result1 msp.MSP
		})
	}
	fake.mSPReturnsOnCall[i] = struct {
		result1 msp.MSP
	}{result1}
}

func (fake *ApplicationOrgConfig) MSPID() string {
	fake.mSPIDMutex.Lock()
	ret, specificReturn := fake.mSPIDReturnsOnCall[len(fake.mSPIDArgsForCall)]
	fake.mSPIDArgsForCall = append(fake.mSPIDArgsForCall, struct {
	}{})
	fake.recordInvocation("MSPID", []interface{}{})
	fake.mSPIDMutex.Unlock()
	if fake.MSPIDStub != nil {
		return fake.MSPIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mSPIDReturns
	return fakeReturns.result1
}

func (fake *ApplicationOrgConfig) MSPIDCallCount() int {
	fake.mSPIDMutex.RLock()
	defer fake.mSPIDMutex.RUnlock()
	return len(fake.mSPIDArgsForCall)
}

func (fake *ApplicationOrgConfig) MSPIDCalls(stub func() string) {
	fake.mSPIDMutex.Lock()
	defer fake.mSPIDMutex.Unlock()
	fake.MSPIDStub = stub
}

func (fake *ApplicationOrgConfig) MSPIDReturns(result1 string) {
	fake.mSPIDMutex.Lock()
	defer fake.mSPIDMutex.Unlock()
	fake.MSPIDStub = nil
	fake.mSPIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *ApplicationOrgConfig) MSPIDReturnsOnCall(i int, result1 string) {
	fake.mSPIDMutex.Lock()
	defer fake.mSPIDMutex.Unlock()
	fake.MSPIDStub = nil
	if fake.mSPIDReturnsOnCall == nil {
		fake.mSPIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.mSPIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ApplicationOrgConfig) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *ApplicationOrgConfig) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *ApplicationOrgConfig) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *ApplicationOrgConfig) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *ApplicationOrgConfig) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ApplicationOrgConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.anchorPeersMutex.RLock()
	defer fake.anchorPeersMutex.RUnlock()
	fake.mSPMutex.RLock()
	defer fake.mSPMutex.RUnlock()
	fake.mSPIDMutex.RLock()
	defer fake.mSPIDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ApplicationOrgConfig) recordInvocation(key string, args []interface{}) {
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