// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"sync"

	"github.com/FuturFusion/migration-manager/internal/migration"
)

// Ensure, that NetworkRepoMock does implement migration.NetworkRepo.
// If this is not the case, regenerate this file with moq.
var _ migration.NetworkRepo = &NetworkRepoMock{}

// NetworkRepoMock is a mock implementation of migration.NetworkRepo.
//
//	func TestSomethingThatUsesNetworkRepo(t *testing.T) {
//
//		// make and configure a mocked migration.NetworkRepo
//		mockedNetworkRepo := &NetworkRepoMock{
//			CreateFunc: func(ctx context.Context, network migration.Network) (int64, error) {
//				panic("mock out the Create method")
//			},
//			DeleteByNameFunc: func(ctx context.Context, name string) error {
//				panic("mock out the DeleteByName method")
//			},
//			GetAllFunc: func(ctx context.Context) (migration.Networks, error) {
//				panic("mock out the GetAll method")
//			},
//			GetAllNamesFunc: func(ctx context.Context) ([]string, error) {
//				panic("mock out the GetAllNames method")
//			},
//			GetByNameFunc: func(ctx context.Context, name string) (*migration.Network, error) {
//				panic("mock out the GetByName method")
//			},
//			RenameFunc: func(ctx context.Context, oldName string, newName string) error {
//				panic("mock out the Rename method")
//			},
//			UpdateFunc: func(ctx context.Context, network migration.Network) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedNetworkRepo in code that requires migration.NetworkRepo
//		// and then make assertions.
//
//	}
type NetworkRepoMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, network migration.Network) (int64, error)

	// DeleteByNameFunc mocks the DeleteByName method.
	DeleteByNameFunc func(ctx context.Context, name string) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) (migration.Networks, error)

	// GetAllNamesFunc mocks the GetAllNames method.
	GetAllNamesFunc func(ctx context.Context) ([]string, error)

	// GetByNameFunc mocks the GetByName method.
	GetByNameFunc func(ctx context.Context, name string) (*migration.Network, error)

	// RenameFunc mocks the Rename method.
	RenameFunc func(ctx context.Context, oldName string, newName string) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, network migration.Network) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Network is the network argument value.
			Network migration.Network
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
			// Network is the network argument value.
			Network migration.Network
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
func (mock *NetworkRepoMock) Create(ctx context.Context, network migration.Network) (int64, error) {
	if mock.CreateFunc == nil {
		panic("NetworkRepoMock.CreateFunc: method is nil but NetworkRepo.Create was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Network migration.Network
	}{
		Ctx:     ctx,
		Network: network,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, network)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedNetworkRepo.CreateCalls())
func (mock *NetworkRepoMock) CreateCalls() []struct {
	Ctx     context.Context
	Network migration.Network
} {
	var calls []struct {
		Ctx     context.Context
		Network migration.Network
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DeleteByName calls DeleteByNameFunc.
func (mock *NetworkRepoMock) DeleteByName(ctx context.Context, name string) error {
	if mock.DeleteByNameFunc == nil {
		panic("NetworkRepoMock.DeleteByNameFunc: method is nil but NetworkRepo.DeleteByName was just called")
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
//	len(mockedNetworkRepo.DeleteByNameCalls())
func (mock *NetworkRepoMock) DeleteByNameCalls() []struct {
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
func (mock *NetworkRepoMock) GetAll(ctx context.Context) (migration.Networks, error) {
	if mock.GetAllFunc == nil {
		panic("NetworkRepoMock.GetAllFunc: method is nil but NetworkRepo.GetAll was just called")
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
//	len(mockedNetworkRepo.GetAllCalls())
func (mock *NetworkRepoMock) GetAllCalls() []struct {
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
func (mock *NetworkRepoMock) GetAllNames(ctx context.Context) ([]string, error) {
	if mock.GetAllNamesFunc == nil {
		panic("NetworkRepoMock.GetAllNamesFunc: method is nil but NetworkRepo.GetAllNames was just called")
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
//	len(mockedNetworkRepo.GetAllNamesCalls())
func (mock *NetworkRepoMock) GetAllNamesCalls() []struct {
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
func (mock *NetworkRepoMock) GetByName(ctx context.Context, name string) (*migration.Network, error) {
	if mock.GetByNameFunc == nil {
		panic("NetworkRepoMock.GetByNameFunc: method is nil but NetworkRepo.GetByName was just called")
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
//	len(mockedNetworkRepo.GetByNameCalls())
func (mock *NetworkRepoMock) GetByNameCalls() []struct {
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
func (mock *NetworkRepoMock) Rename(ctx context.Context, oldName string, newName string) error {
	if mock.RenameFunc == nil {
		panic("NetworkRepoMock.RenameFunc: method is nil but NetworkRepo.Rename was just called")
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
//	len(mockedNetworkRepo.RenameCalls())
func (mock *NetworkRepoMock) RenameCalls() []struct {
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
func (mock *NetworkRepoMock) Update(ctx context.Context, network migration.Network) error {
	if mock.UpdateFunc == nil {
		panic("NetworkRepoMock.UpdateFunc: method is nil but NetworkRepo.Update was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Network migration.Network
	}{
		Ctx:     ctx,
		Network: network,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, network)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedNetworkRepo.UpdateCalls())
func (mock *NetworkRepoMock) UpdateCalls() []struct {
	Ctx     context.Context
	Network migration.Network
} {
	var calls []struct {
		Ctx     context.Context
		Network migration.Network
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
