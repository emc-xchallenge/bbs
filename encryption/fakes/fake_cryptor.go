// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/encryption"
)

type FakeCryptor struct {
	EncryptStub        func(plaintext []byte) (encryption.Encrypted, error)
	encryptMutex       sync.RWMutex
	encryptArgsForCall []struct {
		plaintext []byte
	}
	encryptReturns struct {
		result1 encryption.Encrypted
		result2 error
	}
	DecryptStub        func(encrypted encryption.Encrypted) ([]byte, error)
	decryptMutex       sync.RWMutex
	decryptArgsForCall []struct {
		encrypted encryption.Encrypted
	}
	decryptReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeCryptor) Encrypt(plaintext []byte) (encryption.Encrypted, error) {
	fake.encryptMutex.Lock()
	fake.encryptArgsForCall = append(fake.encryptArgsForCall, struct {
		plaintext []byte
	}{plaintext})
	fake.encryptMutex.Unlock()
	if fake.EncryptStub != nil {
		return fake.EncryptStub(plaintext)
	} else {
		return fake.encryptReturns.result1, fake.encryptReturns.result2
	}
}

func (fake *FakeCryptor) EncryptCallCount() int {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return len(fake.encryptArgsForCall)
}

func (fake *FakeCryptor) EncryptArgsForCall(i int) []byte {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return fake.encryptArgsForCall[i].plaintext
}

func (fake *FakeCryptor) EncryptReturns(result1 encryption.Encrypted, result2 error) {
	fake.EncryptStub = nil
	fake.encryptReturns = struct {
		result1 encryption.Encrypted
		result2 error
	}{result1, result2}
}

func (fake *FakeCryptor) Decrypt(encrypted encryption.Encrypted) ([]byte, error) {
	fake.decryptMutex.Lock()
	fake.decryptArgsForCall = append(fake.decryptArgsForCall, struct {
		encrypted encryption.Encrypted
	}{encrypted})
	fake.decryptMutex.Unlock()
	if fake.DecryptStub != nil {
		return fake.DecryptStub(encrypted)
	} else {
		return fake.decryptReturns.result1, fake.decryptReturns.result2
	}
}

func (fake *FakeCryptor) DecryptCallCount() int {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return len(fake.decryptArgsForCall)
}

func (fake *FakeCryptor) DecryptArgsForCall(i int) encryption.Encrypted {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return fake.decryptArgsForCall[i].encrypted
}

func (fake *FakeCryptor) DecryptReturns(result1 []byte, result2 error) {
	fake.DecryptStub = nil
	fake.decryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ encryption.Cryptor = new(FakeCryptor)
