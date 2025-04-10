// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"sync"

	"github.com/FuturFusion/migration-manager/internal/migration"
)

// Ensure, that SourceRepoMock does implement migration.SourceRepo.
// If this is not the case, regenerate this file with moq.
var _ migration.SourceRepo = &SourceRepoMock{}

// SourceRepoMock is a mock implementation of migration.SourceRepo.
//
//	func TestSomethingThatUsesSourceRepo(t *testing.T) {
//
//		// make and configure a mocked migration.SourceRepo
//		mockedSourceRepo := &SourceRepoMock{
//			CreateFunc: func(ctx context.Context, source migration.Source) (int64, error) {
//				panic("mock out the Create method")
//			},
//			DeleteByNameFunc: func(ctx context.Context, name string) error {
//				panic("mock out the DeleteByName method")
//			},
//			GetAllFunc: func(ctx context.Context) (migration.Sources, error) {
//				panic("mock out the GetAll method")
//			},
//			GetAllNamesFunc: func(ctx context.Context) ([]string, error) {
//				panic("mock out the GetAllNames method")
//			},
//			GetByNameFunc: func(ctx context.Context, name string) (*migration.Source, error) {
//				panic("mock out the GetByName method")
//			},
//			RenameFunc: func(ctx context.Context, oldName string, newName string) error {
//				panic("mock out the Rename method")
//			},
//			UpdateFunc: func(ctx context.Context, name string, source migration.Source) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedSourceRepo in code that requires migration.SourceRepo
//		// and then make assertions.
//
//	}
type SourceRepoMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, source migration.Source) (int64, error)

	// DeleteByNameFunc mocks the DeleteByName method.
	DeleteByNameFunc func(ctx context.Context, name string) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) (migration.Sources, error)

	// GetAllNamesFunc mocks the GetAllNames method.
	GetAllNamesFunc func(ctx context.Context) ([]string, error)

	// GetByNameFunc mocks the GetByName method.
	GetByNameFunc func(ctx context.Context, name string) (*migration.Source, error)

	// RenameFunc mocks the Rename method.
	RenameFunc func(ctx context.Context, oldName string, newName string) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, name string, source migration.Source) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Source is the source argument value.
			Source migration.Source
		}
		// DeleteByName holds details about calls to the DeleteByName method.
		DeleteByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllNames holds details about calls to the GetAllNames method.
		GetAllNames []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByName holds details about calls to the GetByName method.
		GetByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// Rename holds details about calls to the Rename method.
		Rename []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OldName is the oldName argument value.
			OldName string
			// NewName is the newName argument value.
			NewName string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Source is the source argument value.
			Source migration.Source
		}
	}
	lockCreate       sync.RWMutex
	lockDeleteByName sync.RWMutex
	lockGetAll       sync.RWMutex
	lockGetAllNames  sync.RWMutex
	lockGetByName    sync.RWMutex
	lockRename       sync.RWMutex
	lockUpdate       sync.RWMutex
}

// Create calls CreateFunc.
func (mock *SourceRepoMock) Create(ctx context.Context, source migration.Source) (int64, error) {
	if mock.CreateFunc == nil {
		panic("SourceRepoMock.CreateFunc: method is nil but SourceRepo.Create was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Source migration.Source
	}{
		Ctx:    ctx,
		Source: source,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, source)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedSourceRepo.CreateCalls())
func (mock *SourceRepoMock) CreateCalls() []struct {
	Ctx    context.Context
	Source migration.Source
} {
	var calls []struct {
		Ctx    context.Context
		Source migration.Source
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DeleteByName calls DeleteByNameFunc.
func (mock *SourceRepoMock) DeleteByName(ctx context.Context, name string) error {
	if mock.DeleteByNameFunc == nil {
		panic("SourceRepoMock.DeleteByNameFunc: method is nil but SourceRepo.DeleteByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockDeleteByName.Lock()
	mock.calls.DeleteByName = append(mock.calls.DeleteByName, callInfo)
	mock.lockDeleteByName.Unlock()
	return mock.DeleteByNameFunc(ctx, name)
}

// DeleteByNameCalls gets all the calls that were made to DeleteByName.
// Check the length with:
//
//	len(mockedSourceRepo.DeleteByNameCalls())
func (mock *SourceRepoMock) DeleteByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockDeleteByName.RLock()
	calls = mock.calls.DeleteByName
	mock.lockDeleteByName.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *SourceRepoMock) GetAll(ctx context.Context) (migration.Sources, error) {
	if mock.GetAllFunc == nil {
		panic("SourceRepoMock.GetAllFunc: method is nil but SourceRepo.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedSourceRepo.GetAllCalls())
func (mock *SourceRepoMock) GetAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetAllNames calls GetAllNamesFunc.
func (mock *SourceRepoMock) GetAllNames(ctx context.Context) ([]string, error) {
	if mock.GetAllNamesFunc == nil {
		panic("SourceRepoMock.GetAllNamesFunc: method is nil but SourceRepo.GetAllNames was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllNames.Lock()
	mock.calls.GetAllNames = append(mock.calls.GetAllNames, callInfo)
	mock.lockGetAllNames.Unlock()
	return mock.GetAllNamesFunc(ctx)
}

// GetAllNamesCalls gets all the calls that were made to GetAllNames.
// Check the length with:
//
//	len(mockedSourceRepo.GetAllNamesCalls())
func (mock *SourceRepoMock) GetAllNamesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllNames.RLock()
	calls = mock.calls.GetAllNames
	mock.lockGetAllNames.RUnlock()
	return calls
}

// GetByName calls GetByNameFunc.
func (mock *SourceRepoMock) GetByName(ctx context.Context, name string) (*migration.Source, error) {
	if mock.GetByNameFunc == nil {
		panic("SourceRepoMock.GetByNameFunc: method is nil but SourceRepo.GetByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockGetByName.Lock()
	mock.calls.GetByName = append(mock.calls.GetByName, callInfo)
	mock.lockGetByName.Unlock()
	return mock.GetByNameFunc(ctx, name)
}

// GetByNameCalls gets all the calls that were made to GetByName.
// Check the length with:
//
//	len(mockedSourceRepo.GetByNameCalls())
func (mock *SourceRepoMock) GetByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockGetByName.RLock()
	calls = mock.calls.GetByName
	mock.lockGetByName.RUnlock()
	return calls
}

// Rename calls RenameFunc.
func (mock *SourceRepoMock) Rename(ctx context.Context, oldName string, newName string) error {
	if mock.RenameFunc == nil {
		panic("SourceRepoMock.RenameFunc: method is nil but SourceRepo.Rename was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		OldName string
		NewName string
	}{
		Ctx:     ctx,
		OldName: oldName,
		NewName: newName,
	}
	mock.lockRename.Lock()
	mock.calls.Rename = append(mock.calls.Rename, callInfo)
	mock.lockRename.Unlock()
	return mock.RenameFunc(ctx, oldName, newName)
}

// RenameCalls gets all the calls that were made to Rename.
// Check the length with:
//
//	len(mockedSourceRepo.RenameCalls())
func (mock *SourceRepoMock) RenameCalls() []struct {
	Ctx     context.Context
	OldName string
	NewName string
} {
	var calls []struct {
		Ctx     context.Context
		OldName string
		NewName string
	}
	mock.lockRename.RLock()
	calls = mock.calls.Rename
	mock.lockRename.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *SourceRepoMock) Update(ctx context.Context, name string, source migration.Source) error {
	if mock.UpdateFunc == nil {
		panic("SourceRepoMock.UpdateFunc: method is nil but SourceRepo.Update was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Name   string
		Source migration.Source
	}{
		Ctx:    ctx,
		Name:   name,
		Source: source,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, name, source)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedSourceRepo.UpdateCalls())
func (mock *SourceRepoMock) UpdateCalls() []struct {
	Ctx    context.Context
	Name   string
	Source migration.Source
} {
	var calls []struct {
		Ctx    context.Context
		Name   string
		Source migration.Source
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
