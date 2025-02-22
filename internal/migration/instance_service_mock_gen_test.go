// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package migration_test

import (
	"context"
	"sync"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/shared/api"
	"github.com/google/uuid"
)

// Ensure, that InstanceServiceMock does implement migration.InstanceService.
// If this is not the case, regenerate this file with moq.
var _ migration.InstanceService = &InstanceServiceMock{}

// InstanceServiceMock is a mock implementation of migration.InstanceService.
//
//	func TestSomethingThatUsesInstanceService(t *testing.T) {
//
//		// make and configure a mocked migration.InstanceService
//		mockedInstanceService := &InstanceServiceMock{
//			CreateFunc: func(ctx context.Context, instance migration.Instance) (migration.Instance, error) {
//				panic("mock out the Create method")
//			},
//			CreateOverridesFunc: func(ctx context.Context, overrides migration.Overrides) (migration.Overrides, error) {
//				panic("mock out the CreateOverrides method")
//			},
//			DeleteByIDFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the DeleteByID method")
//			},
//			DeleteOverridesByIDFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the DeleteOverridesByID method")
//			},
//			GetAllFunc: func(ctx context.Context) (migration.Instances, error) {
//				panic("mock out the GetAll method")
//			},
//			GetAllByBatchIDFunc: func(ctx context.Context, batchID int) (migration.Instances, error) {
//				panic("mock out the GetAllByBatchID method")
//			},
//			GetAllByStateFunc: func(ctx context.Context, status api.MigrationStatusType) (migration.Instances, error) {
//				panic("mock out the GetAllByState method")
//			},
//			GetAllUUIDsFunc: func(ctx context.Context) ([]uuid.UUID, error) {
//				panic("mock out the GetAllUUIDs method")
//			},
//			GetAllUnassignedFunc: func(ctx context.Context) (migration.Instances, error) {
//				panic("mock out the GetAllUnassigned method")
//			},
//			GetByIDFunc: func(ctx context.Context, id uuid.UUID) (migration.Instance, error) {
//				panic("mock out the GetByID method")
//			},
//			GetByIDWithDetailsFunc: func(ctx context.Context, id uuid.UUID) (migration.InstanceWithDetails, error) {
//				panic("mock out the GetByIDWithDetails method")
//			},
//			GetOverridesByIDFunc: func(ctx context.Context, id uuid.UUID) (migration.Overrides, error) {
//				panic("mock out the GetOverridesByID method")
//			},
//			ProcessWorkerUpdateFunc: func(ctx context.Context, id uuid.UUID, workerResponseTypeArg api.WorkerResponseType, statusString string) (migration.Instance, error) {
//				panic("mock out the ProcessWorkerUpdate method")
//			},
//			UnassignFromBatchFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the UnassignFromBatch method")
//			},
//			UpdateByIDFunc: func(ctx context.Context, instance migration.Instance) (migration.Instance, error) {
//				panic("mock out the UpdateByID method")
//			},
//			UpdateOverridesByIDFunc: func(ctx context.Context, overrides migration.Overrides) (migration.Overrides, error) {
//				panic("mock out the UpdateOverridesByID method")
//			},
//			UpdateStatusByUUIDFunc: func(ctx context.Context, id uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (migration.Instance, error) {
//				panic("mock out the UpdateStatusByUUID method")
//			},
//		}
//
//		// use mockedInstanceService in code that requires migration.InstanceService
//		// and then make assertions.
//
//	}
type InstanceServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, instance migration.Instance) (migration.Instance, error)

	// CreateOverridesFunc mocks the CreateOverrides method.
	CreateOverridesFunc func(ctx context.Context, overrides migration.Overrides) (migration.Overrides, error)

	// DeleteByIDFunc mocks the DeleteByID method.
	DeleteByIDFunc func(ctx context.Context, id uuid.UUID) error

	// DeleteOverridesByIDFunc mocks the DeleteOverridesByID method.
	DeleteOverridesByIDFunc func(ctx context.Context, id uuid.UUID) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) (migration.Instances, error)

	// GetAllByBatchIDFunc mocks the GetAllByBatchID method.
	GetAllByBatchIDFunc func(ctx context.Context, batchID int) (migration.Instances, error)

	// GetAllByStateFunc mocks the GetAllByState method.
	GetAllByStateFunc func(ctx context.Context, status api.MigrationStatusType) (migration.Instances, error)

	// GetAllUUIDsFunc mocks the GetAllUUIDs method.
	GetAllUUIDsFunc func(ctx context.Context) ([]uuid.UUID, error)

	// GetAllUnassignedFunc mocks the GetAllUnassigned method.
	GetAllUnassignedFunc func(ctx context.Context) (migration.Instances, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, id uuid.UUID) (migration.Instance, error)

	// GetByIDWithDetailsFunc mocks the GetByIDWithDetails method.
	GetByIDWithDetailsFunc func(ctx context.Context, id uuid.UUID) (migration.InstanceWithDetails, error)

	// GetOverridesByIDFunc mocks the GetOverridesByID method.
	GetOverridesByIDFunc func(ctx context.Context, id uuid.UUID) (migration.Overrides, error)

	// ProcessWorkerUpdateFunc mocks the ProcessWorkerUpdate method.
	ProcessWorkerUpdateFunc func(ctx context.Context, id uuid.UUID, workerResponseTypeArg api.WorkerResponseType, statusString string) (migration.Instance, error)

	// UnassignFromBatchFunc mocks the UnassignFromBatch method.
	UnassignFromBatchFunc func(ctx context.Context, id uuid.UUID) error

	// UpdateByIDFunc mocks the UpdateByID method.
	UpdateByIDFunc func(ctx context.Context, instance migration.Instance) (migration.Instance, error)

	// UpdateOverridesByIDFunc mocks the UpdateOverridesByID method.
	UpdateOverridesByIDFunc func(ctx context.Context, overrides migration.Overrides) (migration.Overrides, error)

	// UpdateStatusByUUIDFunc mocks the UpdateStatusByUUID method.
	UpdateStatusByUUIDFunc func(ctx context.Context, id uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (migration.Instance, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Instance is the instance argument value.
			Instance migration.Instance
		}
		// CreateOverrides holds details about calls to the CreateOverrides method.
		CreateOverrides []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Overrides is the overrides argument value.
			Overrides migration.Overrides
		}
		// DeleteByID holds details about calls to the DeleteByID method.
		DeleteByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// DeleteOverridesByID holds details about calls to the DeleteOverridesByID method.
		DeleteOverridesByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllByBatchID holds details about calls to the GetAllByBatchID method.
		GetAllByBatchID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BatchID is the batchID argument value.
			BatchID int
		}
		// GetAllByState holds details about calls to the GetAllByState method.
		GetAllByState []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Status is the status argument value.
			Status api.MigrationStatusType
		}
		// GetAllUUIDs holds details about calls to the GetAllUUIDs method.
		GetAllUUIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllUnassigned holds details about calls to the GetAllUnassigned method.
		GetAllUnassigned []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetByIDWithDetails holds details about calls to the GetByIDWithDetails method.
		GetByIDWithDetails []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetOverridesByID holds details about calls to the GetOverridesByID method.
		GetOverridesByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// ProcessWorkerUpdate holds details about calls to the ProcessWorkerUpdate method.
		ProcessWorkerUpdate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
			// WorkerResponseTypeArg is the workerResponseTypeArg argument value.
			WorkerResponseTypeArg api.WorkerResponseType
			// StatusString is the statusString argument value.
			StatusString string
		}
		// UnassignFromBatch holds details about calls to the UnassignFromBatch method.
		UnassignFromBatch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// UpdateByID holds details about calls to the UpdateByID method.
		UpdateByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Instance is the instance argument value.
			Instance migration.Instance
		}
		// UpdateOverridesByID holds details about calls to the UpdateOverridesByID method.
		UpdateOverridesByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Overrides is the overrides argument value.
			Overrides migration.Overrides
		}
		// UpdateStatusByUUID holds details about calls to the UpdateStatusByUUID method.
		UpdateStatusByUUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
			// Status is the status argument value.
			Status api.MigrationStatusType
			// StatusString is the statusString argument value.
			StatusString string
			// NeedsDiskImport is the needsDiskImport argument value.
			NeedsDiskImport bool
		}
	}
	lockCreate              sync.RWMutex
	lockCreateOverrides     sync.RWMutex
	lockDeleteByID          sync.RWMutex
	lockDeleteOverridesByID sync.RWMutex
	lockGetAll              sync.RWMutex
	lockGetAllByBatchID     sync.RWMutex
	lockGetAllByState       sync.RWMutex
	lockGetAllUUIDs         sync.RWMutex
	lockGetAllUnassigned    sync.RWMutex
	lockGetByID             sync.RWMutex
	lockGetByIDWithDetails  sync.RWMutex
	lockGetOverridesByID    sync.RWMutex
	lockProcessWorkerUpdate sync.RWMutex
	lockUnassignFromBatch   sync.RWMutex
	lockUpdateByID          sync.RWMutex
	lockUpdateOverridesByID sync.RWMutex
	lockUpdateStatusByUUID  sync.RWMutex
}

// Create calls CreateFunc.
func (mock *InstanceServiceMock) Create(ctx context.Context, instance migration.Instance) (migration.Instance, error) {
	if mock.CreateFunc == nil {
		panic("InstanceServiceMock.CreateFunc: method is nil but InstanceService.Create was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Instance migration.Instance
	}{
		Ctx:      ctx,
		Instance: instance,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, instance)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedInstanceService.CreateCalls())
func (mock *InstanceServiceMock) CreateCalls() []struct {
	Ctx      context.Context
	Instance migration.Instance
} {
	var calls []struct {
		Ctx      context.Context
		Instance migration.Instance
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// CreateOverrides calls CreateOverridesFunc.
func (mock *InstanceServiceMock) CreateOverrides(ctx context.Context, overrides migration.Overrides) (migration.Overrides, error) {
	if mock.CreateOverridesFunc == nil {
		panic("InstanceServiceMock.CreateOverridesFunc: method is nil but InstanceService.CreateOverrides was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Overrides migration.Overrides
	}{
		Ctx:       ctx,
		Overrides: overrides,
	}
	mock.lockCreateOverrides.Lock()
	mock.calls.CreateOverrides = append(mock.calls.CreateOverrides, callInfo)
	mock.lockCreateOverrides.Unlock()
	return mock.CreateOverridesFunc(ctx, overrides)
}

// CreateOverridesCalls gets all the calls that were made to CreateOverrides.
// Check the length with:
//
//	len(mockedInstanceService.CreateOverridesCalls())
func (mock *InstanceServiceMock) CreateOverridesCalls() []struct {
	Ctx       context.Context
	Overrides migration.Overrides
} {
	var calls []struct {
		Ctx       context.Context
		Overrides migration.Overrides
	}
	mock.lockCreateOverrides.RLock()
	calls = mock.calls.CreateOverrides
	mock.lockCreateOverrides.RUnlock()
	return calls
}

// DeleteByID calls DeleteByIDFunc.
func (mock *InstanceServiceMock) DeleteByID(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteByIDFunc == nil {
		panic("InstanceServiceMock.DeleteByIDFunc: method is nil but InstanceService.DeleteByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteByID.Lock()
	mock.calls.DeleteByID = append(mock.calls.DeleteByID, callInfo)
	mock.lockDeleteByID.Unlock()
	return mock.DeleteByIDFunc(ctx, id)
}

// DeleteByIDCalls gets all the calls that were made to DeleteByID.
// Check the length with:
//
//	len(mockedInstanceService.DeleteByIDCalls())
func (mock *InstanceServiceMock) DeleteByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDeleteByID.RLock()
	calls = mock.calls.DeleteByID
	mock.lockDeleteByID.RUnlock()
	return calls
}

// DeleteOverridesByID calls DeleteOverridesByIDFunc.
func (mock *InstanceServiceMock) DeleteOverridesByID(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteOverridesByIDFunc == nil {
		panic("InstanceServiceMock.DeleteOverridesByIDFunc: method is nil but InstanceService.DeleteOverridesByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteOverridesByID.Lock()
	mock.calls.DeleteOverridesByID = append(mock.calls.DeleteOverridesByID, callInfo)
	mock.lockDeleteOverridesByID.Unlock()
	return mock.DeleteOverridesByIDFunc(ctx, id)
}

// DeleteOverridesByIDCalls gets all the calls that were made to DeleteOverridesByID.
// Check the length with:
//
//	len(mockedInstanceService.DeleteOverridesByIDCalls())
func (mock *InstanceServiceMock) DeleteOverridesByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDeleteOverridesByID.RLock()
	calls = mock.calls.DeleteOverridesByID
	mock.lockDeleteOverridesByID.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *InstanceServiceMock) GetAll(ctx context.Context) (migration.Instances, error) {
	if mock.GetAllFunc == nil {
		panic("InstanceServiceMock.GetAllFunc: method is nil but InstanceService.GetAll was just called")
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
//	len(mockedInstanceService.GetAllCalls())
func (mock *InstanceServiceMock) GetAllCalls() []struct {
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

// GetAllByBatchID calls GetAllByBatchIDFunc.
func (mock *InstanceServiceMock) GetAllByBatchID(ctx context.Context, batchID int) (migration.Instances, error) {
	if mock.GetAllByBatchIDFunc == nil {
		panic("InstanceServiceMock.GetAllByBatchIDFunc: method is nil but InstanceService.GetAllByBatchID was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		BatchID int
	}{
		Ctx:     ctx,
		BatchID: batchID,
	}
	mock.lockGetAllByBatchID.Lock()
	mock.calls.GetAllByBatchID = append(mock.calls.GetAllByBatchID, callInfo)
	mock.lockGetAllByBatchID.Unlock()
	return mock.GetAllByBatchIDFunc(ctx, batchID)
}

// GetAllByBatchIDCalls gets all the calls that were made to GetAllByBatchID.
// Check the length with:
//
//	len(mockedInstanceService.GetAllByBatchIDCalls())
func (mock *InstanceServiceMock) GetAllByBatchIDCalls() []struct {
	Ctx     context.Context
	BatchID int
} {
	var calls []struct {
		Ctx     context.Context
		BatchID int
	}
	mock.lockGetAllByBatchID.RLock()
	calls = mock.calls.GetAllByBatchID
	mock.lockGetAllByBatchID.RUnlock()
	return calls
}

// GetAllByState calls GetAllByStateFunc.
func (mock *InstanceServiceMock) GetAllByState(ctx context.Context, status api.MigrationStatusType) (migration.Instances, error) {
	if mock.GetAllByStateFunc == nil {
		panic("InstanceServiceMock.GetAllByStateFunc: method is nil but InstanceService.GetAllByState was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Status api.MigrationStatusType
	}{
		Ctx:    ctx,
		Status: status,
	}
	mock.lockGetAllByState.Lock()
	mock.calls.GetAllByState = append(mock.calls.GetAllByState, callInfo)
	mock.lockGetAllByState.Unlock()
	return mock.GetAllByStateFunc(ctx, status)
}

// GetAllByStateCalls gets all the calls that were made to GetAllByState.
// Check the length with:
//
//	len(mockedInstanceService.GetAllByStateCalls())
func (mock *InstanceServiceMock) GetAllByStateCalls() []struct {
	Ctx    context.Context
	Status api.MigrationStatusType
} {
	var calls []struct {
		Ctx    context.Context
		Status api.MigrationStatusType
	}
	mock.lockGetAllByState.RLock()
	calls = mock.calls.GetAllByState
	mock.lockGetAllByState.RUnlock()
	return calls
}

// GetAllUUIDs calls GetAllUUIDsFunc.
func (mock *InstanceServiceMock) GetAllUUIDs(ctx context.Context) ([]uuid.UUID, error) {
	if mock.GetAllUUIDsFunc == nil {
		panic("InstanceServiceMock.GetAllUUIDsFunc: method is nil but InstanceService.GetAllUUIDs was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllUUIDs.Lock()
	mock.calls.GetAllUUIDs = append(mock.calls.GetAllUUIDs, callInfo)
	mock.lockGetAllUUIDs.Unlock()
	return mock.GetAllUUIDsFunc(ctx)
}

// GetAllUUIDsCalls gets all the calls that were made to GetAllUUIDs.
// Check the length with:
//
//	len(mockedInstanceService.GetAllUUIDsCalls())
func (mock *InstanceServiceMock) GetAllUUIDsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllUUIDs.RLock()
	calls = mock.calls.GetAllUUIDs
	mock.lockGetAllUUIDs.RUnlock()
	return calls
}

// GetAllUnassigned calls GetAllUnassignedFunc.
func (mock *InstanceServiceMock) GetAllUnassigned(ctx context.Context) (migration.Instances, error) {
	if mock.GetAllUnassignedFunc == nil {
		panic("InstanceServiceMock.GetAllUnassignedFunc: method is nil but InstanceService.GetAllUnassigned was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllUnassigned.Lock()
	mock.calls.GetAllUnassigned = append(mock.calls.GetAllUnassigned, callInfo)
	mock.lockGetAllUnassigned.Unlock()
	return mock.GetAllUnassignedFunc(ctx)
}

// GetAllUnassignedCalls gets all the calls that were made to GetAllUnassigned.
// Check the length with:
//
//	len(mockedInstanceService.GetAllUnassignedCalls())
func (mock *InstanceServiceMock) GetAllUnassignedCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllUnassigned.RLock()
	calls = mock.calls.GetAllUnassigned
	mock.lockGetAllUnassigned.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *InstanceServiceMock) GetByID(ctx context.Context, id uuid.UUID) (migration.Instance, error) {
	if mock.GetByIDFunc == nil {
		panic("InstanceServiceMock.GetByIDFunc: method is nil but InstanceService.GetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedInstanceService.GetByIDCalls())
func (mock *InstanceServiceMock) GetByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// GetByIDWithDetails calls GetByIDWithDetailsFunc.
func (mock *InstanceServiceMock) GetByIDWithDetails(ctx context.Context, id uuid.UUID) (migration.InstanceWithDetails, error) {
	if mock.GetByIDWithDetailsFunc == nil {
		panic("InstanceServiceMock.GetByIDWithDetailsFunc: method is nil but InstanceService.GetByIDWithDetails was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetByIDWithDetails.Lock()
	mock.calls.GetByIDWithDetails = append(mock.calls.GetByIDWithDetails, callInfo)
	mock.lockGetByIDWithDetails.Unlock()
	return mock.GetByIDWithDetailsFunc(ctx, id)
}

// GetByIDWithDetailsCalls gets all the calls that were made to GetByIDWithDetails.
// Check the length with:
//
//	len(mockedInstanceService.GetByIDWithDetailsCalls())
func (mock *InstanceServiceMock) GetByIDWithDetailsCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetByIDWithDetails.RLock()
	calls = mock.calls.GetByIDWithDetails
	mock.lockGetByIDWithDetails.RUnlock()
	return calls
}

// GetOverridesByID calls GetOverridesByIDFunc.
func (mock *InstanceServiceMock) GetOverridesByID(ctx context.Context, id uuid.UUID) (migration.Overrides, error) {
	if mock.GetOverridesByIDFunc == nil {
		panic("InstanceServiceMock.GetOverridesByIDFunc: method is nil but InstanceService.GetOverridesByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetOverridesByID.Lock()
	mock.calls.GetOverridesByID = append(mock.calls.GetOverridesByID, callInfo)
	mock.lockGetOverridesByID.Unlock()
	return mock.GetOverridesByIDFunc(ctx, id)
}

// GetOverridesByIDCalls gets all the calls that were made to GetOverridesByID.
// Check the length with:
//
//	len(mockedInstanceService.GetOverridesByIDCalls())
func (mock *InstanceServiceMock) GetOverridesByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetOverridesByID.RLock()
	calls = mock.calls.GetOverridesByID
	mock.lockGetOverridesByID.RUnlock()
	return calls
}

// ProcessWorkerUpdate calls ProcessWorkerUpdateFunc.
func (mock *InstanceServiceMock) ProcessWorkerUpdate(ctx context.Context, id uuid.UUID, workerResponseTypeArg api.WorkerResponseType, statusString string) (migration.Instance, error) {
	if mock.ProcessWorkerUpdateFunc == nil {
		panic("InstanceServiceMock.ProcessWorkerUpdateFunc: method is nil but InstanceService.ProcessWorkerUpdate was just called")
	}
	callInfo := struct {
		Ctx                   context.Context
		ID                    uuid.UUID
		WorkerResponseTypeArg api.WorkerResponseType
		StatusString          string
	}{
		Ctx:                   ctx,
		ID:                    id,
		WorkerResponseTypeArg: workerResponseTypeArg,
		StatusString:          statusString,
	}
	mock.lockProcessWorkerUpdate.Lock()
	mock.calls.ProcessWorkerUpdate = append(mock.calls.ProcessWorkerUpdate, callInfo)
	mock.lockProcessWorkerUpdate.Unlock()
	return mock.ProcessWorkerUpdateFunc(ctx, id, workerResponseTypeArg, statusString)
}

// ProcessWorkerUpdateCalls gets all the calls that were made to ProcessWorkerUpdate.
// Check the length with:
//
//	len(mockedInstanceService.ProcessWorkerUpdateCalls())
func (mock *InstanceServiceMock) ProcessWorkerUpdateCalls() []struct {
	Ctx                   context.Context
	ID                    uuid.UUID
	WorkerResponseTypeArg api.WorkerResponseType
	StatusString          string
} {
	var calls []struct {
		Ctx                   context.Context
		ID                    uuid.UUID
		WorkerResponseTypeArg api.WorkerResponseType
		StatusString          string
	}
	mock.lockProcessWorkerUpdate.RLock()
	calls = mock.calls.ProcessWorkerUpdate
	mock.lockProcessWorkerUpdate.RUnlock()
	return calls
}

// UnassignFromBatch calls UnassignFromBatchFunc.
func (mock *InstanceServiceMock) UnassignFromBatch(ctx context.Context, id uuid.UUID) error {
	if mock.UnassignFromBatchFunc == nil {
		panic("InstanceServiceMock.UnassignFromBatchFunc: method is nil but InstanceService.UnassignFromBatch was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockUnassignFromBatch.Lock()
	mock.calls.UnassignFromBatch = append(mock.calls.UnassignFromBatch, callInfo)
	mock.lockUnassignFromBatch.Unlock()
	return mock.UnassignFromBatchFunc(ctx, id)
}

// UnassignFromBatchCalls gets all the calls that were made to UnassignFromBatch.
// Check the length with:
//
//	len(mockedInstanceService.UnassignFromBatchCalls())
func (mock *InstanceServiceMock) UnassignFromBatchCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockUnassignFromBatch.RLock()
	calls = mock.calls.UnassignFromBatch
	mock.lockUnassignFromBatch.RUnlock()
	return calls
}

// UpdateByID calls UpdateByIDFunc.
func (mock *InstanceServiceMock) UpdateByID(ctx context.Context, instance migration.Instance) (migration.Instance, error) {
	if mock.UpdateByIDFunc == nil {
		panic("InstanceServiceMock.UpdateByIDFunc: method is nil but InstanceService.UpdateByID was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Instance migration.Instance
	}{
		Ctx:      ctx,
		Instance: instance,
	}
	mock.lockUpdateByID.Lock()
	mock.calls.UpdateByID = append(mock.calls.UpdateByID, callInfo)
	mock.lockUpdateByID.Unlock()
	return mock.UpdateByIDFunc(ctx, instance)
}

// UpdateByIDCalls gets all the calls that were made to UpdateByID.
// Check the length with:
//
//	len(mockedInstanceService.UpdateByIDCalls())
func (mock *InstanceServiceMock) UpdateByIDCalls() []struct {
	Ctx      context.Context
	Instance migration.Instance
} {
	var calls []struct {
		Ctx      context.Context
		Instance migration.Instance
	}
	mock.lockUpdateByID.RLock()
	calls = mock.calls.UpdateByID
	mock.lockUpdateByID.RUnlock()
	return calls
}

// UpdateOverridesByID calls UpdateOverridesByIDFunc.
func (mock *InstanceServiceMock) UpdateOverridesByID(ctx context.Context, overrides migration.Overrides) (migration.Overrides, error) {
	if mock.UpdateOverridesByIDFunc == nil {
		panic("InstanceServiceMock.UpdateOverridesByIDFunc: method is nil but InstanceService.UpdateOverridesByID was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Overrides migration.Overrides
	}{
		Ctx:       ctx,
		Overrides: overrides,
	}
	mock.lockUpdateOverridesByID.Lock()
	mock.calls.UpdateOverridesByID = append(mock.calls.UpdateOverridesByID, callInfo)
	mock.lockUpdateOverridesByID.Unlock()
	return mock.UpdateOverridesByIDFunc(ctx, overrides)
}

// UpdateOverridesByIDCalls gets all the calls that were made to UpdateOverridesByID.
// Check the length with:
//
//	len(mockedInstanceService.UpdateOverridesByIDCalls())
func (mock *InstanceServiceMock) UpdateOverridesByIDCalls() []struct {
	Ctx       context.Context
	Overrides migration.Overrides
} {
	var calls []struct {
		Ctx       context.Context
		Overrides migration.Overrides
	}
	mock.lockUpdateOverridesByID.RLock()
	calls = mock.calls.UpdateOverridesByID
	mock.lockUpdateOverridesByID.RUnlock()
	return calls
}

// UpdateStatusByUUID calls UpdateStatusByUUIDFunc.
func (mock *InstanceServiceMock) UpdateStatusByUUID(ctx context.Context, id uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (migration.Instance, error) {
	if mock.UpdateStatusByUUIDFunc == nil {
		panic("InstanceServiceMock.UpdateStatusByUUIDFunc: method is nil but InstanceService.UpdateStatusByUUID was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		ID              uuid.UUID
		Status          api.MigrationStatusType
		StatusString    string
		NeedsDiskImport bool
	}{
		Ctx:             ctx,
		ID:              id,
		Status:          status,
		StatusString:    statusString,
		NeedsDiskImport: needsDiskImport,
	}
	mock.lockUpdateStatusByUUID.Lock()
	mock.calls.UpdateStatusByUUID = append(mock.calls.UpdateStatusByUUID, callInfo)
	mock.lockUpdateStatusByUUID.Unlock()
	return mock.UpdateStatusByUUIDFunc(ctx, id, status, statusString, needsDiskImport)
}

// UpdateStatusByUUIDCalls gets all the calls that were made to UpdateStatusByUUID.
// Check the length with:
//
//	len(mockedInstanceService.UpdateStatusByUUIDCalls())
func (mock *InstanceServiceMock) UpdateStatusByUUIDCalls() []struct {
	Ctx             context.Context
	ID              uuid.UUID
	Status          api.MigrationStatusType
	StatusString    string
	NeedsDiskImport bool
} {
	var calls []struct {
		Ctx             context.Context
		ID              uuid.UUID
		Status          api.MigrationStatusType
		StatusString    string
		NeedsDiskImport bool
	}
	mock.lockUpdateStatusByUUID.RLock()
	calls = mock.calls.UpdateStatusByUUID
	mock.lockUpdateStatusByUUID.RUnlock()
	return calls
}
