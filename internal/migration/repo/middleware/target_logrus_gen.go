// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus
// gowrap: http://github.com/hexdigest/gowrap

package middleware

import (
	"context"

	_sourceMigration "github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/sirupsen/logrus"
)

// TargetRepoWithLogrus implements _sourceMigration.TargetRepo that is instrumented with logrus logger
type TargetRepoWithLogrus struct {
	_log  *logrus.Entry
	_base _sourceMigration.TargetRepo
}

// NewTargetRepoWithLogrus instruments an implementation of the _sourceMigration.TargetRepo with simple logging
func NewTargetRepoWithLogrus(base _sourceMigration.TargetRepo, log *logrus.Entry) TargetRepoWithLogrus {
	return TargetRepoWithLogrus{
		_base: base,
		_log:  log,
	}
}

// Create implements _sourceMigration.TargetRepo
func (_d TargetRepoWithLogrus) Create(ctx context.Context, target _sourceMigration.Target) (t1 _sourceMigration.Target, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":    ctx,
		"target": target})).Debug("TargetRepoWithLogrus: calling Create")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Error("TargetRepoWithLogrus: method Create returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Debug("TargetRepoWithLogrus: method Create finished")
		}
	}()
	return _d._base.Create(ctx, target)
}

// DeleteByName implements _sourceMigration.TargetRepo
func (_d TargetRepoWithLogrus) DeleteByName(ctx context.Context, name string) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":  ctx,
		"name": name})).Debug("TargetRepoWithLogrus: calling DeleteByName")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("TargetRepoWithLogrus: method DeleteByName returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("TargetRepoWithLogrus: method DeleteByName finished")
		}
	}()
	return _d._base.DeleteByName(ctx, name)
}

// GetAll implements _sourceMigration.TargetRepo
func (_d TargetRepoWithLogrus) GetAll(ctx context.Context) (t1 _sourceMigration.Targets, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx": ctx})).Debug("TargetRepoWithLogrus: calling GetAll")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Error("TargetRepoWithLogrus: method GetAll returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Debug("TargetRepoWithLogrus: method GetAll finished")
		}
	}()
	return _d._base.GetAll(ctx)
}

// GetByID implements _sourceMigration.TargetRepo
func (_d TargetRepoWithLogrus) GetByID(ctx context.Context, id int) (t1 _sourceMigration.Target, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx": ctx,
		"id":  id})).Debug("TargetRepoWithLogrus: calling GetByID")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Error("TargetRepoWithLogrus: method GetByID returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Debug("TargetRepoWithLogrus: method GetByID finished")
		}
	}()
	return _d._base.GetByID(ctx, id)
}

// GetByName implements _sourceMigration.TargetRepo
func (_d TargetRepoWithLogrus) GetByName(ctx context.Context, name string) (t1 _sourceMigration.Target, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":  ctx,
		"name": name})).Debug("TargetRepoWithLogrus: calling GetByName")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Error("TargetRepoWithLogrus: method GetByName returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Debug("TargetRepoWithLogrus: method GetByName finished")
		}
	}()
	return _d._base.GetByName(ctx, name)
}

// UpdateByName implements _sourceMigration.TargetRepo
func (_d TargetRepoWithLogrus) UpdateByName(ctx context.Context, target _sourceMigration.Target) (t1 _sourceMigration.Target, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":    ctx,
		"target": target})).Debug("TargetRepoWithLogrus: calling UpdateByName")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Error("TargetRepoWithLogrus: method UpdateByName returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"t1":  t1,
				"err": err})).Debug("TargetRepoWithLogrus: method UpdateByName finished")
		}
	}()
	return _d._base.UpdateByName(ctx, target)
}
