// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/algorithm"
)

type FakeJobCombination struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	JobIDStub        func() int
	jobIDMutex       sync.RWMutex
	jobIDArgsForCall []struct{}
	jobIDReturns     struct {
		result1 int
	}
	jobIDReturnsOnCall map[int]struct {
		result1 int
	}
	ResourceSpaceIDStub        func() int
	resourceSpaceIDMutex       sync.RWMutex
	resourceSpaceIDArgsForCall []struct{}
	resourceSpaceIDReturns     struct {
		result1 int
	}
	resourceSpaceIDReturnsOnCall map[int]struct {
		result1 int
	}
	ResourceSpacesStub        func() map[string]string
	resourceSpacesMutex       sync.RWMutex
	resourceSpacesArgsForCall []struct{}
	resourceSpacesReturns     struct {
		result1 map[string]string
	}
	resourceSpacesReturnsOnCall map[int]struct {
		result1 map[string]string
	}
	CreateBuildStub        func() (db.Build, error)
	createBuildMutex       sync.RWMutex
	createBuildArgsForCall []struct{}
	createBuildReturns     struct {
		result1 db.Build
		result2 error
	}
	createBuildReturnsOnCall map[int]struct {
		result1 db.Build
		result2 error
	}
	EnsurePendingBuildExistsStub        func() error
	ensurePendingBuildExistsMutex       sync.RWMutex
	ensurePendingBuildExistsArgsForCall []struct{}
	ensurePendingBuildExistsReturns     struct {
		result1 error
	}
	ensurePendingBuildExistsReturnsOnCall map[int]struct {
		result1 error
	}
	GetNextBuildInputsStub        func() (algorithm.InputMapping, bool, error)
	getNextBuildInputsMutex       sync.RWMutex
	getNextBuildInputsArgsForCall []struct{}
	getNextBuildInputsReturns     struct {
		result1 algorithm.InputMapping
		result2 bool
		result3 error
	}
	getNextBuildInputsReturnsOnCall map[int]struct {
		result1 algorithm.InputMapping
		result2 bool
		result3 error
	}
	SaveNextInputMappingStub        func(inputMapping algorithm.InputMapping) error
	saveNextInputMappingMutex       sync.RWMutex
	saveNextInputMappingArgsForCall []struct {
		inputMapping algorithm.InputMapping
	}
	saveNextInputMappingReturns struct {
		result1 error
	}
	saveNextInputMappingReturnsOnCall map[int]struct {
		result1 error
	}
	SaveIndependentInputMappingStub        func(inputMapping algorithm.InputMapping) error
	saveIndependentInputMappingMutex       sync.RWMutex
	saveIndependentInputMappingArgsForCall []struct {
		inputMapping algorithm.InputMapping
	}
	saveIndependentInputMappingReturns struct {
		result1 error
	}
	saveIndependentInputMappingReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteNextInputMappingStub        func() error
	deleteNextInputMappingMutex       sync.RWMutex
	deleteNextInputMappingArgsForCall []struct{}
	deleteNextInputMappingReturns     struct {
		result1 error
	}
	deleteNextInputMappingReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJobCombination) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeJobCombination) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeJobCombination) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeJobCombination) IDReturnsOnCall(i int, result1 string) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeJobCombination) JobID() int {
	fake.jobIDMutex.Lock()
	ret, specificReturn := fake.jobIDReturnsOnCall[len(fake.jobIDArgsForCall)]
	fake.jobIDArgsForCall = append(fake.jobIDArgsForCall, struct{}{})
	fake.recordInvocation("JobID", []interface{}{})
	fake.jobIDMutex.Unlock()
	if fake.JobIDStub != nil {
		return fake.JobIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.jobIDReturns.result1
}

func (fake *FakeJobCombination) JobIDCallCount() int {
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	return len(fake.jobIDArgsForCall)
}

func (fake *FakeJobCombination) JobIDReturns(result1 int) {
	fake.JobIDStub = nil
	fake.jobIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeJobCombination) JobIDReturnsOnCall(i int, result1 int) {
	fake.JobIDStub = nil
	if fake.jobIDReturnsOnCall == nil {
		fake.jobIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.jobIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeJobCombination) ResourceSpaceID() int {
	fake.resourceSpaceIDMutex.Lock()
	ret, specificReturn := fake.resourceSpaceIDReturnsOnCall[len(fake.resourceSpaceIDArgsForCall)]
	fake.resourceSpaceIDArgsForCall = append(fake.resourceSpaceIDArgsForCall, struct{}{})
	fake.recordInvocation("ResourceSpaceID", []interface{}{})
	fake.resourceSpaceIDMutex.Unlock()
	if fake.ResourceSpaceIDStub != nil {
		return fake.ResourceSpaceIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceSpaceIDReturns.result1
}

func (fake *FakeJobCombination) ResourceSpaceIDCallCount() int {
	fake.resourceSpaceIDMutex.RLock()
	defer fake.resourceSpaceIDMutex.RUnlock()
	return len(fake.resourceSpaceIDArgsForCall)
}

func (fake *FakeJobCombination) ResourceSpaceIDReturns(result1 int) {
	fake.ResourceSpaceIDStub = nil
	fake.resourceSpaceIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeJobCombination) ResourceSpaceIDReturnsOnCall(i int, result1 int) {
	fake.ResourceSpaceIDStub = nil
	if fake.resourceSpaceIDReturnsOnCall == nil {
		fake.resourceSpaceIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.resourceSpaceIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeJobCombination) ResourceSpaces() map[string]string {
	fake.resourceSpacesMutex.Lock()
	ret, specificReturn := fake.resourceSpacesReturnsOnCall[len(fake.resourceSpacesArgsForCall)]
	fake.resourceSpacesArgsForCall = append(fake.resourceSpacesArgsForCall, struct{}{})
	fake.recordInvocation("ResourceSpaces", []interface{}{})
	fake.resourceSpacesMutex.Unlock()
	if fake.ResourceSpacesStub != nil {
		return fake.ResourceSpacesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceSpacesReturns.result1
}

func (fake *FakeJobCombination) ResourceSpacesCallCount() int {
	fake.resourceSpacesMutex.RLock()
	defer fake.resourceSpacesMutex.RUnlock()
	return len(fake.resourceSpacesArgsForCall)
}

func (fake *FakeJobCombination) ResourceSpacesReturns(result1 map[string]string) {
	fake.ResourceSpacesStub = nil
	fake.resourceSpacesReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeJobCombination) ResourceSpacesReturnsOnCall(i int, result1 map[string]string) {
	fake.ResourceSpacesStub = nil
	if fake.resourceSpacesReturnsOnCall == nil {
		fake.resourceSpacesReturnsOnCall = make(map[int]struct {
			result1 map[string]string
		})
	}
	fake.resourceSpacesReturnsOnCall[i] = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeJobCombination) CreateBuild() (db.Build, error) {
	fake.createBuildMutex.Lock()
	ret, specificReturn := fake.createBuildReturnsOnCall[len(fake.createBuildArgsForCall)]
	fake.createBuildArgsForCall = append(fake.createBuildArgsForCall, struct{}{})
	fake.recordInvocation("CreateBuild", []interface{}{})
	fake.createBuildMutex.Unlock()
	if fake.CreateBuildStub != nil {
		return fake.CreateBuildStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createBuildReturns.result1, fake.createBuildReturns.result2
}

func (fake *FakeJobCombination) CreateBuildCallCount() int {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return len(fake.createBuildArgsForCall)
}

func (fake *FakeJobCombination) CreateBuildReturns(result1 db.Build, result2 error) {
	fake.CreateBuildStub = nil
	fake.createBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeJobCombination) CreateBuildReturnsOnCall(i int, result1 db.Build, result2 error) {
	fake.CreateBuildStub = nil
	if fake.createBuildReturnsOnCall == nil {
		fake.createBuildReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 error
		})
	}
	fake.createBuildReturnsOnCall[i] = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeJobCombination) EnsurePendingBuildExists() error {
	fake.ensurePendingBuildExistsMutex.Lock()
	ret, specificReturn := fake.ensurePendingBuildExistsReturnsOnCall[len(fake.ensurePendingBuildExistsArgsForCall)]
	fake.ensurePendingBuildExistsArgsForCall = append(fake.ensurePendingBuildExistsArgsForCall, struct{}{})
	fake.recordInvocation("EnsurePendingBuildExists", []interface{}{})
	fake.ensurePendingBuildExistsMutex.Unlock()
	if fake.EnsurePendingBuildExistsStub != nil {
		return fake.EnsurePendingBuildExistsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.ensurePendingBuildExistsReturns.result1
}

func (fake *FakeJobCombination) EnsurePendingBuildExistsCallCount() int {
	fake.ensurePendingBuildExistsMutex.RLock()
	defer fake.ensurePendingBuildExistsMutex.RUnlock()
	return len(fake.ensurePendingBuildExistsArgsForCall)
}

func (fake *FakeJobCombination) EnsurePendingBuildExistsReturns(result1 error) {
	fake.EnsurePendingBuildExistsStub = nil
	fake.ensurePendingBuildExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) EnsurePendingBuildExistsReturnsOnCall(i int, result1 error) {
	fake.EnsurePendingBuildExistsStub = nil
	if fake.ensurePendingBuildExistsReturnsOnCall == nil {
		fake.ensurePendingBuildExistsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ensurePendingBuildExistsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) GetNextBuildInputs() (algorithm.InputMapping, bool, error) {
	fake.getNextBuildInputsMutex.Lock()
	ret, specificReturn := fake.getNextBuildInputsReturnsOnCall[len(fake.getNextBuildInputsArgsForCall)]
	fake.getNextBuildInputsArgsForCall = append(fake.getNextBuildInputsArgsForCall, struct{}{})
	fake.recordInvocation("GetNextBuildInputs", []interface{}{})
	fake.getNextBuildInputsMutex.Unlock()
	if fake.GetNextBuildInputsStub != nil {
		return fake.GetNextBuildInputsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getNextBuildInputsReturns.result1, fake.getNextBuildInputsReturns.result2, fake.getNextBuildInputsReturns.result3
}

func (fake *FakeJobCombination) GetNextBuildInputsCallCount() int {
	fake.getNextBuildInputsMutex.RLock()
	defer fake.getNextBuildInputsMutex.RUnlock()
	return len(fake.getNextBuildInputsArgsForCall)
}

func (fake *FakeJobCombination) GetNextBuildInputsReturns(result1 algorithm.InputMapping, result2 bool, result3 error) {
	fake.GetNextBuildInputsStub = nil
	fake.getNextBuildInputsReturns = struct {
		result1 algorithm.InputMapping
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeJobCombination) GetNextBuildInputsReturnsOnCall(i int, result1 algorithm.InputMapping, result2 bool, result3 error) {
	fake.GetNextBuildInputsStub = nil
	if fake.getNextBuildInputsReturnsOnCall == nil {
		fake.getNextBuildInputsReturnsOnCall = make(map[int]struct {
			result1 algorithm.InputMapping
			result2 bool
			result3 error
		})
	}
	fake.getNextBuildInputsReturnsOnCall[i] = struct {
		result1 algorithm.InputMapping
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeJobCombination) SaveNextInputMapping(inputMapping algorithm.InputMapping) error {
	fake.saveNextInputMappingMutex.Lock()
	ret, specificReturn := fake.saveNextInputMappingReturnsOnCall[len(fake.saveNextInputMappingArgsForCall)]
	fake.saveNextInputMappingArgsForCall = append(fake.saveNextInputMappingArgsForCall, struct {
		inputMapping algorithm.InputMapping
	}{inputMapping})
	fake.recordInvocation("SaveNextInputMapping", []interface{}{inputMapping})
	fake.saveNextInputMappingMutex.Unlock()
	if fake.SaveNextInputMappingStub != nil {
		return fake.SaveNextInputMappingStub(inputMapping)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveNextInputMappingReturns.result1
}

func (fake *FakeJobCombination) SaveNextInputMappingCallCount() int {
	fake.saveNextInputMappingMutex.RLock()
	defer fake.saveNextInputMappingMutex.RUnlock()
	return len(fake.saveNextInputMappingArgsForCall)
}

func (fake *FakeJobCombination) SaveNextInputMappingArgsForCall(i int) algorithm.InputMapping {
	fake.saveNextInputMappingMutex.RLock()
	defer fake.saveNextInputMappingMutex.RUnlock()
	return fake.saveNextInputMappingArgsForCall[i].inputMapping
}

func (fake *FakeJobCombination) SaveNextInputMappingReturns(result1 error) {
	fake.SaveNextInputMappingStub = nil
	fake.saveNextInputMappingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) SaveNextInputMappingReturnsOnCall(i int, result1 error) {
	fake.SaveNextInputMappingStub = nil
	if fake.saveNextInputMappingReturnsOnCall == nil {
		fake.saveNextInputMappingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveNextInputMappingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) SaveIndependentInputMapping(inputMapping algorithm.InputMapping) error {
	fake.saveIndependentInputMappingMutex.Lock()
	ret, specificReturn := fake.saveIndependentInputMappingReturnsOnCall[len(fake.saveIndependentInputMappingArgsForCall)]
	fake.saveIndependentInputMappingArgsForCall = append(fake.saveIndependentInputMappingArgsForCall, struct {
		inputMapping algorithm.InputMapping
	}{inputMapping})
	fake.recordInvocation("SaveIndependentInputMapping", []interface{}{inputMapping})
	fake.saveIndependentInputMappingMutex.Unlock()
	if fake.SaveIndependentInputMappingStub != nil {
		return fake.SaveIndependentInputMappingStub(inputMapping)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveIndependentInputMappingReturns.result1
}

func (fake *FakeJobCombination) SaveIndependentInputMappingCallCount() int {
	fake.saveIndependentInputMappingMutex.RLock()
	defer fake.saveIndependentInputMappingMutex.RUnlock()
	return len(fake.saveIndependentInputMappingArgsForCall)
}

func (fake *FakeJobCombination) SaveIndependentInputMappingArgsForCall(i int) algorithm.InputMapping {
	fake.saveIndependentInputMappingMutex.RLock()
	defer fake.saveIndependentInputMappingMutex.RUnlock()
	return fake.saveIndependentInputMappingArgsForCall[i].inputMapping
}

func (fake *FakeJobCombination) SaveIndependentInputMappingReturns(result1 error) {
	fake.SaveIndependentInputMappingStub = nil
	fake.saveIndependentInputMappingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) SaveIndependentInputMappingReturnsOnCall(i int, result1 error) {
	fake.SaveIndependentInputMappingStub = nil
	if fake.saveIndependentInputMappingReturnsOnCall == nil {
		fake.saveIndependentInputMappingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveIndependentInputMappingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) DeleteNextInputMapping() error {
	fake.deleteNextInputMappingMutex.Lock()
	ret, specificReturn := fake.deleteNextInputMappingReturnsOnCall[len(fake.deleteNextInputMappingArgsForCall)]
	fake.deleteNextInputMappingArgsForCall = append(fake.deleteNextInputMappingArgsForCall, struct{}{})
	fake.recordInvocation("DeleteNextInputMapping", []interface{}{})
	fake.deleteNextInputMappingMutex.Unlock()
	if fake.DeleteNextInputMappingStub != nil {
		return fake.DeleteNextInputMappingStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteNextInputMappingReturns.result1
}

func (fake *FakeJobCombination) DeleteNextInputMappingCallCount() int {
	fake.deleteNextInputMappingMutex.RLock()
	defer fake.deleteNextInputMappingMutex.RUnlock()
	return len(fake.deleteNextInputMappingArgsForCall)
}

func (fake *FakeJobCombination) DeleteNextInputMappingReturns(result1 error) {
	fake.DeleteNextInputMappingStub = nil
	fake.deleteNextInputMappingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) DeleteNextInputMappingReturnsOnCall(i int, result1 error) {
	fake.DeleteNextInputMappingStub = nil
	if fake.deleteNextInputMappingReturnsOnCall == nil {
		fake.deleteNextInputMappingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteNextInputMappingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobCombination) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	fake.resourceSpaceIDMutex.RLock()
	defer fake.resourceSpaceIDMutex.RUnlock()
	fake.resourceSpacesMutex.RLock()
	defer fake.resourceSpacesMutex.RUnlock()
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	fake.ensurePendingBuildExistsMutex.RLock()
	defer fake.ensurePendingBuildExistsMutex.RUnlock()
	fake.getNextBuildInputsMutex.RLock()
	defer fake.getNextBuildInputsMutex.RUnlock()
	fake.saveNextInputMappingMutex.RLock()
	defer fake.saveNextInputMappingMutex.RUnlock()
	fake.saveIndependentInputMappingMutex.RLock()
	defer fake.saveIndependentInputMappingMutex.RUnlock()
	fake.deleteNextInputMappingMutex.RLock()
	defer fake.deleteNextInputMappingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJobCombination) recordInvocation(key string, args []interface{}) {
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

var _ db.JobCombination = new(FakeJobCombination)