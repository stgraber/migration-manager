// Code generated by gowrap. DO NOT EDIT.
// template: ../../../logger/slog.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

package middleware

import (
	"context"
	"log/slog"

	_sourceMigration "github.com/FuturFusion/migration-manager/internal/migration"
)

// TargetRepoWithSlog implements _sourceMigration.TargetRepo that is instrumented with slog logger
type TargetRepoWithSlog struct {
	_log  *slog.Logger
	_base _sourceMigration.TargetRepo
}

// NewTargetRepoWithSlog instruments an implementation of the _sourceMigration.TargetRepo with simple logging
func NewTargetRepoWithSlog(base _sourceMigration.TargetRepo, log *slog.Logger) TargetRepoWithSlog {
	return TargetRepoWithSlog{
		_base: base,
		_log:  log,
	}
}

// Create implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) Create(ctx context.Context, target _sourceMigration.Target) (i1 int64, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("target", target),
	).Debug("TargetRepoWithSlog: calling Create")
	defer func() {
		log := _d._log.With(
			slog.Int64("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method Create returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method Create finished")
		}
	}()
	return _d._base.Create(ctx, target)
}

// DeleteByName implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) DeleteByName(ctx context.Context, name string) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("name", name),
	).Debug("TargetRepoWithSlog: calling DeleteByName")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method DeleteByName returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method DeleteByName finished")
		}
	}()
	return _d._base.DeleteByName(ctx, name)
}

// GetAll implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) GetAll(ctx context.Context) (t1 _sourceMigration.Targets, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("TargetRepoWithSlog: calling GetAll")
	defer func() {
		log := _d._log.With(
			slog.Any("t1", t1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method GetAll returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method GetAll finished")
		}
	}()
	return _d._base.GetAll(ctx)
}

// GetAllNames implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) GetAllNames(ctx context.Context) (sa1 []string, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("TargetRepoWithSlog: calling GetAllNames")
	defer func() {
		log := _d._log.With(
			slog.Any("sa1", sa1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method GetAllNames returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method GetAllNames finished")
		}
	}()
	return _d._base.GetAllNames(ctx)
}

// GetByName implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) GetByName(ctx context.Context, name string) (tp1 *_sourceMigration.Target, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("name", name),
	).Debug("TargetRepoWithSlog: calling GetByName")
	defer func() {
		log := _d._log.With(
			slog.Any("tp1", tp1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method GetByName returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method GetByName finished")
		}
	}()
	return _d._base.GetByName(ctx, name)
}

// Rename implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) Rename(ctx context.Context, oldName string, newName string) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("oldName", oldName),
		slog.String("newName", newName),
	).Debug("TargetRepoWithSlog: calling Rename")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method Rename returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method Rename finished")
		}
	}()
	return _d._base.Rename(ctx, oldName, newName)
}

// Update implements _sourceMigration.TargetRepo
func (_d TargetRepoWithSlog) Update(ctx context.Context, name string, target _sourceMigration.Target) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("name", name),
		slog.Any("target", target),
	).Debug("TargetRepoWithSlog: calling Update")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetRepoWithSlog: method Update returned an error")
		} else {
			log.Debug("TargetRepoWithSlog: method Update finished")
		}
	}()
	return _d._base.Update(ctx, name, target)
}
