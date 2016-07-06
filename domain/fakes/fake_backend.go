// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/cloudfoundry-incubator/switchboard/domain"
)

type FakeBackend struct {
	HealthcheckUrlStub        func() string
	healthcheckUrlMutex       sync.RWMutex
	healthcheckUrlArgsForCall []struct{}
	healthcheckUrlReturns struct {
		result1 string
	}
	BridgeStub        func(clientConn net.Conn) error
	bridgeMutex       sync.RWMutex
	bridgeArgsForCall []struct {
		clientConn net.Conn
	}
	bridgeReturns struct {
		result1 error
	}
	SeverConnectionsStub        func()
	severConnectionsMutex       sync.RWMutex
	severConnectionsArgsForCall []struct{}
	AsJSONStub        func() domain.BackendJSON
	asJSONMutex       sync.RWMutex
	asJSONArgsForCall []struct{}
	asJSONReturns struct {
		result1 domain.BackendJSON
	}
	EnableTrafficStub        func()
	enableTrafficMutex       sync.RWMutex
	enableTrafficArgsForCall []struct{}
	DisableTrafficStub        func()
	disableTrafficMutex       sync.RWMutex
	disableTrafficArgsForCall []struct{}
}

func (fake *FakeBackend) HealthcheckUrl() string {
	fake.healthcheckUrlMutex.Lock()
	fake.healthcheckUrlArgsForCall = append(fake.healthcheckUrlArgsForCall, struct{}{})
	fake.healthcheckUrlMutex.Unlock()
	if fake.HealthcheckUrlStub != nil {
		return fake.HealthcheckUrlStub()
	} else {
		return fake.healthcheckUrlReturns.result1
	}
}

func (fake *FakeBackend) HealthcheckUrlCallCount() int {
	fake.healthcheckUrlMutex.RLock()
	defer fake.healthcheckUrlMutex.RUnlock()
	return len(fake.healthcheckUrlArgsForCall)
}

func (fake *FakeBackend) HealthcheckUrlReturns(result1 string) {
	fake.HealthcheckUrlStub = nil
	fake.healthcheckUrlReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBackend) Bridge(clientConn net.Conn) error {
	fake.bridgeMutex.Lock()
	fake.bridgeArgsForCall = append(fake.bridgeArgsForCall, struct {
		clientConn net.Conn
	}{clientConn})
	fake.bridgeMutex.Unlock()
	if fake.BridgeStub != nil {
		return fake.BridgeStub(clientConn)
	} else {
		return fake.bridgeReturns.result1
	}
}

func (fake *FakeBackend) BridgeCallCount() int {
	fake.bridgeMutex.RLock()
	defer fake.bridgeMutex.RUnlock()
	return len(fake.bridgeArgsForCall)
}

func (fake *FakeBackend) BridgeArgsForCall(i int) net.Conn {
	fake.bridgeMutex.RLock()
	defer fake.bridgeMutex.RUnlock()
	return fake.bridgeArgsForCall[i].clientConn
}

func (fake *FakeBackend) BridgeReturns(result1 error) {
	fake.BridgeStub = nil
	fake.bridgeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) SeverConnections() {
	fake.severConnectionsMutex.Lock()
	fake.severConnectionsArgsForCall = append(fake.severConnectionsArgsForCall, struct{}{})
	fake.severConnectionsMutex.Unlock()
	if fake.SeverConnectionsStub != nil {
		fake.SeverConnectionsStub()
	}
}

func (fake *FakeBackend) SeverConnectionsCallCount() int {
	fake.severConnectionsMutex.RLock()
	defer fake.severConnectionsMutex.RUnlock()
	return len(fake.severConnectionsArgsForCall)
}

func (fake *FakeBackend) AsJSON() domain.BackendJSON {
	fake.asJSONMutex.Lock()
	fake.asJSONArgsForCall = append(fake.asJSONArgsForCall, struct{}{})
	fake.asJSONMutex.Unlock()
	if fake.AsJSONStub != nil {
		return fake.AsJSONStub()
	} else {
		return fake.asJSONReturns.result1
	}
}

func (fake *FakeBackend) AsJSONCallCount() int {
	fake.asJSONMutex.RLock()
	defer fake.asJSONMutex.RUnlock()
	return len(fake.asJSONArgsForCall)
}

func (fake *FakeBackend) AsJSONReturns(result1 domain.BackendJSON) {
	fake.AsJSONStub = nil
	fake.asJSONReturns = struct {
		result1 domain.BackendJSON
	}{result1}
}

func (fake *FakeBackend) EnableTraffic() {
	fake.enableTrafficMutex.Lock()
	fake.enableTrafficArgsForCall = append(fake.enableTrafficArgsForCall, struct{}{})
	fake.enableTrafficMutex.Unlock()
	if fake.EnableTrafficStub != nil {
		fake.EnableTrafficStub()
	}
}

func (fake *FakeBackend) EnableTrafficCallCount() int {
	fake.enableTrafficMutex.RLock()
	defer fake.enableTrafficMutex.RUnlock()
	return len(fake.enableTrafficArgsForCall)
}

func (fake *FakeBackend) DisableTraffic() {
	fake.disableTrafficMutex.Lock()
	fake.disableTrafficArgsForCall = append(fake.disableTrafficArgsForCall, struct{}{})
	fake.disableTrafficMutex.Unlock()
	if fake.DisableTrafficStub != nil {
		fake.DisableTrafficStub()
	}
}

func (fake *FakeBackend) DisableTrafficCallCount() int {
	fake.disableTrafficMutex.RLock()
	defer fake.disableTrafficMutex.RUnlock()
	return len(fake.disableTrafficArgsForCall)
}

var _ domain.Backend = new(FakeBackend)
