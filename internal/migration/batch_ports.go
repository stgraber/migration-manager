package migration

import (
	"context"

	"github.com/FuturFusion/migration-manager/shared/api"
)

//go:generate go run github.com/matryer/moq -fmt goimports -pkg migration_test -out batch_service_mock_gen_test.go -rm . BatchService

type BatchService interface {
	Create(ctx context.Context, batch Batch) (Batch, error)
	GetAll(ctx context.Context) (Batches, error)
	GetAllByState(ctx context.Context, status api.BatchStatusType) (Batches, error)
	GetAllNames(ctx context.Context) ([]string, error)
	GetAllNamesByState(ctx context.Context, status api.BatchStatusType) ([]string, error)
	GetByName(ctx context.Context, name string) (*Batch, error)
	Update(ctx context.Context, batch *Batch) error
	UpdateInstancesAssignedToBatch(ctx context.Context, batch Batch) error
	UpdateStatusByName(ctx context.Context, name string, status api.BatchStatusType, statusString string) (*Batch, error)
	Rename(ctx context.Context, oldName string, newName string) error
	DeleteByName(ctx context.Context, name string) error
	StartBatchByName(ctx context.Context, name string) error
	StopBatchByName(ctx context.Context, name string) error
}

//go:generate go run github.com/matryer/moq -fmt goimports -pkg mock -out repo/mock/batch_repo_mock_gen.go -rm . BatchRepo
//go:generate go run github.com/hexdigest/gowrap/cmd/gowrap gen -g -i BatchRepo -t ../logger/slog.gotmpl -o ./repo/middleware/batch_slog_gen.go
// disabled go:generate go run github.com/hexdigest/gowrap/cmd/gowrap gen -g -i BatchRepo -t prometheus -o ./repo/middleware/batch_prometheus_gen.go

type BatchRepo interface {
	Create(ctx context.Context, batch Batch) (int64, error)
	GetAll(ctx context.Context) (Batches, error)
	GetAllByState(ctx context.Context, status api.BatchStatusType) (Batches, error)
	GetAllNames(ctx context.Context) ([]string, error)
	GetAllNamesByState(ctx context.Context, status api.BatchStatusType) ([]string, error)
	GetByName(ctx context.Context, name string) (*Batch, error)
	Update(ctx context.Context, batch Batch) error
	Rename(ctx context.Context, oldName string, newName string) error
	DeleteByName(ctx context.Context, name string) error
}
