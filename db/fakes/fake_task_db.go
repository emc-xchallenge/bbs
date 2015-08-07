// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type FakeTaskDB struct {
	TasksStub        func(logger lager.Logger, filter db.TaskFilter) (*models.Tasks, *models.Error)
	tasksMutex       sync.RWMutex
	tasksArgsForCall []struct {
		logger lager.Logger
		filter db.TaskFilter
	}
	tasksReturns struct {
		result1 *models.Tasks
		result2 *models.Error
	}
	TaskByGuidStub        func(logger lager.Logger, processGuid string) (*models.Task, *models.Error)
	taskByGuidMutex       sync.RWMutex
	taskByGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	taskByGuidReturns struct {
		result1 *models.Task
		result2 *models.Error
	}
	DesireTaskStub        func(logger lager.Logger, request *models.DesireTaskRequest) *models.Error
	desireTaskMutex       sync.RWMutex
	desireTaskArgsForCall []struct {
		logger  lager.Logger
		request *models.DesireTaskRequest
	}
	desireTaskReturns struct {
		result1 *models.Error
	}
	StartTaskStub        func(logger lager.Logger, request *models.StartTaskRequest) (bool, *models.Error)
	startTaskMutex       sync.RWMutex
	startTaskArgsForCall []struct {
		logger  lager.Logger
		request *models.StartTaskRequest
	}
	startTaskReturns struct {
		result1 bool
		result2 *models.Error
	}
}

func (fake *FakeTaskDB) Tasks(logger lager.Logger, filter db.TaskFilter) (*models.Tasks, *models.Error) {
	fake.tasksMutex.Lock()
	fake.tasksArgsForCall = append(fake.tasksArgsForCall, struct {
		logger lager.Logger
		filter db.TaskFilter
	}{logger, filter})
	fake.tasksMutex.Unlock()
	if fake.TasksStub != nil {
		return fake.TasksStub(logger, filter)
	} else {
		return fake.tasksReturns.result1, fake.tasksReturns.result2
	}
}

func (fake *FakeTaskDB) TasksCallCount() int {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return len(fake.tasksArgsForCall)
}

func (fake *FakeTaskDB) TasksArgsForCall(i int) (lager.Logger, db.TaskFilter) {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return fake.tasksArgsForCall[i].logger, fake.tasksArgsForCall[i].filter
}

func (fake *FakeTaskDB) TasksReturns(result1 *models.Tasks, result2 *models.Error) {
	fake.TasksStub = nil
	fake.tasksReturns = struct {
		result1 *models.Tasks
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeTaskDB) TaskByGuid(logger lager.Logger, processGuid string) (*models.Task, *models.Error) {
	fake.taskByGuidMutex.Lock()
	fake.taskByGuidArgsForCall = append(fake.taskByGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.taskByGuidMutex.Unlock()
	if fake.TaskByGuidStub != nil {
		return fake.TaskByGuidStub(logger, processGuid)
	} else {
		return fake.taskByGuidReturns.result1, fake.taskByGuidReturns.result2
	}
}

func (fake *FakeTaskDB) TaskByGuidCallCount() int {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return len(fake.taskByGuidArgsForCall)
}

func (fake *FakeTaskDB) TaskByGuidArgsForCall(i int) (lager.Logger, string) {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return fake.taskByGuidArgsForCall[i].logger, fake.taskByGuidArgsForCall[i].processGuid
}

func (fake *FakeTaskDB) TaskByGuidReturns(result1 *models.Task, result2 *models.Error) {
	fake.TaskByGuidStub = nil
	fake.taskByGuidReturns = struct {
		result1 *models.Task
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeTaskDB) DesireTask(logger lager.Logger, request *models.DesireTaskRequest) *models.Error {
	fake.desireTaskMutex.Lock()
	fake.desireTaskArgsForCall = append(fake.desireTaskArgsForCall, struct {
		logger  lager.Logger
		request *models.DesireTaskRequest
	}{logger, request})
	fake.desireTaskMutex.Unlock()
	if fake.DesireTaskStub != nil {
		return fake.DesireTaskStub(logger, request)
	} else {
		return fake.desireTaskReturns.result1
	}
}

func (fake *FakeTaskDB) DesireTaskCallCount() int {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return len(fake.desireTaskArgsForCall)
}

func (fake *FakeTaskDB) DesireTaskArgsForCall(i int) (lager.Logger, *models.DesireTaskRequest) {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return fake.desireTaskArgsForCall[i].logger, fake.desireTaskArgsForCall[i].request
}

func (fake *FakeTaskDB) DesireTaskReturns(result1 *models.Error) {
	fake.DesireTaskStub = nil
	fake.desireTaskReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeTaskDB) StartTask(logger lager.Logger, request *models.StartTaskRequest) (bool, *models.Error) {
	fake.startTaskMutex.Lock()
	fake.startTaskArgsForCall = append(fake.startTaskArgsForCall, struct {
		logger  lager.Logger
		request *models.StartTaskRequest
	}{logger, request})
	fake.startTaskMutex.Unlock()
	if fake.StartTaskStub != nil {
		return fake.StartTaskStub(logger, request)
	} else {
		return fake.startTaskReturns.result1, fake.startTaskReturns.result2
	}
}

func (fake *FakeTaskDB) StartTaskCallCount() int {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return len(fake.startTaskArgsForCall)
}

func (fake *FakeTaskDB) StartTaskArgsForCall(i int) (lager.Logger, *models.StartTaskRequest) {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return fake.startTaskArgsForCall[i].logger, fake.startTaskArgsForCall[i].request
}

func (fake *FakeTaskDB) StartTaskReturns(result1 bool, result2 *models.Error) {
	fake.StartTaskStub = nil
	fake.startTaskReturns = struct {
		result1 bool
		result2 *models.Error
	}{result1, result2}
}

var _ db.TaskDB = new(FakeTaskDB)
