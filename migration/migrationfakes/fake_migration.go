// This file was generated by counterfeiter
package migrationfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/db/etcd"
	"github.com/cloudfoundry-incubator/bbs/encryption"
	"github.com/cloudfoundry-incubator/bbs/migration"
	"github.com/pivotal-golang/lager"
)

type FakeMigration struct {
	VersionStub        func() int64
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 int64
	}
	UpStub        func(logger lager.Logger) error
	upMutex       sync.RWMutex
	upArgsForCall []struct {
		logger lager.Logger
	}
	upReturns struct {
		result1 error
	}
	DownStub        func(logger lager.Logger) error
	downMutex       sync.RWMutex
	downArgsForCall []struct {
		logger lager.Logger
	}
	downReturns struct {
		result1 error
	}
	SetStoreClientStub        func(storeClient etcd.StoreClient)
	setStoreClientMutex       sync.RWMutex
	setStoreClientArgsForCall []struct {
		storeClient etcd.StoreClient
	}
	SetCryptorStub        func(cryptor encryption.Cryptor)
	setCryptorMutex       sync.RWMutex
	setCryptorArgsForCall []struct {
		cryptor encryption.Cryptor
	}
}

func (fake *FakeMigration) Version() int64 {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeMigration) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeMigration) VersionReturns(result1 int64) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMigration) Up(logger lager.Logger) error {
	fake.upMutex.Lock()
	fake.upArgsForCall = append(fake.upArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.upMutex.Unlock()
	if fake.UpStub != nil {
		return fake.UpStub(logger)
	} else {
		return fake.upReturns.result1
	}
}

func (fake *FakeMigration) UpCallCount() int {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return len(fake.upArgsForCall)
}

func (fake *FakeMigration) UpArgsForCall(i int) lager.Logger {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return fake.upArgsForCall[i].logger
}

func (fake *FakeMigration) UpReturns(result1 error) {
	fake.UpStub = nil
	fake.upReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) Down(logger lager.Logger) error {
	fake.downMutex.Lock()
	fake.downArgsForCall = append(fake.downArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.downMutex.Unlock()
	if fake.DownStub != nil {
		return fake.DownStub(logger)
	} else {
		return fake.downReturns.result1
	}
}

func (fake *FakeMigration) DownCallCount() int {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return len(fake.downArgsForCall)
}

func (fake *FakeMigration) DownArgsForCall(i int) lager.Logger {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return fake.downArgsForCall[i].logger
}

func (fake *FakeMigration) DownReturns(result1 error) {
	fake.DownStub = nil
	fake.downReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) SetStoreClient(storeClient etcd.StoreClient) {
	fake.setStoreClientMutex.Lock()
	fake.setStoreClientArgsForCall = append(fake.setStoreClientArgsForCall, struct {
		storeClient etcd.StoreClient
	}{storeClient})
	fake.setStoreClientMutex.Unlock()
	if fake.SetStoreClientStub != nil {
		fake.SetStoreClientStub(storeClient)
	}
}

func (fake *FakeMigration) SetStoreClientCallCount() int {
	fake.setStoreClientMutex.RLock()
	defer fake.setStoreClientMutex.RUnlock()
	return len(fake.setStoreClientArgsForCall)
}

func (fake *FakeMigration) SetStoreClientArgsForCall(i int) etcd.StoreClient {
	fake.setStoreClientMutex.RLock()
	defer fake.setStoreClientMutex.RUnlock()
	return fake.setStoreClientArgsForCall[i].storeClient
}

func (fake *FakeMigration) SetCryptor(cryptor encryption.Cryptor) {
	fake.setCryptorMutex.Lock()
	fake.setCryptorArgsForCall = append(fake.setCryptorArgsForCall, struct {
		cryptor encryption.Cryptor
	}{cryptor})
	fake.setCryptorMutex.Unlock()
	if fake.SetCryptorStub != nil {
		fake.SetCryptorStub(cryptor)
	}
}

func (fake *FakeMigration) SetCryptorCallCount() int {
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	return len(fake.setCryptorArgsForCall)
}

func (fake *FakeMigration) SetCryptorArgsForCall(i int) encryption.Cryptor {
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	return fake.setCryptorArgsForCall[i].cryptor
}

var _ migration.Migration = new(FakeMigration)
