// Code generated by generate-database from the incus project - DO NOT EDIT.

package entities

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/mattn/go-sqlite3"
)

var sourceObjects = RegisterStmt(`
SELECT sources.id, sources.name, sources.source_type, sources.properties
  FROM sources
  ORDER BY sources.name
`)

var sourceObjectsByName = RegisterStmt(`
SELECT sources.id, sources.name, sources.source_type, sources.properties
  FROM sources
  WHERE ( sources.name = ? )
  ORDER BY sources.name
`)

var sourceNames = RegisterStmt(`
SELECT sources.name
  FROM sources
  ORDER BY sources.name
`)

var sourceID = RegisterStmt(`
SELECT sources.id FROM sources
  WHERE sources.name = ?
`)

var sourceCreate = RegisterStmt(`
INSERT INTO sources (name, source_type, properties)
  VALUES (?, ?, ?)
`)

var sourceUpdate = RegisterStmt(`
UPDATE sources
  SET name = ?, source_type = ?, properties = ?
 WHERE id = ?
`)

var sourceRename = RegisterStmt(`
UPDATE sources SET name = ? WHERE name = ?
`)

var sourceDeleteByName = RegisterStmt(`
DELETE FROM sources WHERE name = ?
`)

// GetSourceID return the ID of the source with the given key.
// generator: source ID
func GetSourceID(ctx context.Context, db tx, name string) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	stmt, err := Stmt(db, sourceID)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"sourceID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, ErrNotFound
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to get \"sources\" ID: %w", err)
	}

	return id, nil
}

// SourceExists checks if a source with the given key exists.
// generator: source Exists
func SourceExists(ctx context.Context, db dbtx, name string) (_ bool, _err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	stmt, err := Stmt(db, sourceID)
	if err != nil {
		return false, fmt.Errorf("Failed to get \"sourceID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("Failed to get \"sources\" ID: %w", err)
	}

	return true, nil
}

// GetSource returns the source with the given key.
// generator: source GetOne
func GetSource(ctx context.Context, db dbtx, name string) (_ *migration.Source, _err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	filter := SourceFilter{}
	filter.Name = &name

	objects, err := GetSources(ctx, db, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"sources\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, ErrNotFound
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"sources\" entry matches")
	}
}

// sourceColumns returns a string of column names to be used with a SELECT statement for the entity.
// Use this function when building statements to retrieve database entries matching the Source entity.
func sourceColumns() string {
	return "sources.id, sources.name, sources.source_type, sources.properties"
}

// getSources can be used to run handwritten sql.Stmts to return a slice of objects.
func getSources(ctx context.Context, stmt *sql.Stmt, args ...any) ([]migration.Source, error) {
	objects := make([]migration.Source, 0)

	dest := func(scan func(dest ...any) error) error {
		s := migration.Source{}
		err := scan(&s.ID, &s.Name, &s.SourceType, &s.Properties)
		if err != nil {
			return err
		}

		objects = append(objects, s)

		return nil
	}

	err := selectObjects(ctx, stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"sources\" table: %w", err)
	}

	return objects, nil
}

// getSourcesRaw can be used to run handwritten query strings to return a slice of objects.
func getSourcesRaw(ctx context.Context, db dbtx, sql string, args ...any) ([]migration.Source, error) {
	objects := make([]migration.Source, 0)

	dest := func(scan func(dest ...any) error) error {
		s := migration.Source{}
		err := scan(&s.ID, &s.Name, &s.SourceType, &s.Properties)
		if err != nil {
			return err
		}

		objects = append(objects, s)

		return nil
	}

	err := scan(ctx, db, sql, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"sources\" table: %w", err)
	}

	return objects, nil
}

// GetSources returns all available sources.
// generator: source GetMany
func GetSources(ctx context.Context, db dbtx, filters ...SourceFilter) (_ []migration.Source, _err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	var err error

	// Result slice.
	objects := make([]migration.Source, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	args := []any{}
	queryParts := [2]string{}

	if len(filters) == 0 {
		sqlStmt, err = Stmt(db, sourceObjects)
		if err != nil {
			return nil, fmt.Errorf("Failed to get \"sourceObjects\" prepared statement: %w", err)
		}
	}

	for i, filter := range filters {
		if filter.Name != nil {
			args = append(args, []any{filter.Name}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, sourceObjectsByName)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"sourceObjectsByName\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(sourceObjectsByName)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"sourceObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.Name == nil {
			return nil, fmt.Errorf("Cannot filter on empty SourceFilter")
		} else {
			return nil, fmt.Errorf("No statement exists for the given Filter")
		}
	}

	// Select.
	if sqlStmt != nil {
		objects, err = getSources(ctx, sqlStmt, args...)
	} else {
		queryStr := strings.Join(queryParts[:], "ORDER BY")
		objects, err = getSourcesRaw(ctx, db, queryStr, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"sources\" table: %w", err)
	}

	return objects, nil
}

// GetSourceNames returns the identifying field of source.
// generator: source GetNames
func GetSourceNames(ctx context.Context, db dbtx, filters ...SourceFilter) (_ []string, _err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	var err error

	// Result slice.
	names := make([]string, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	args := []any{}
	queryParts := [2]string{}

	if len(filters) == 0 {
		sqlStmt, err = Stmt(db, sourceNames)
		if err != nil {
			return nil, fmt.Errorf("Failed to get \"sourceNames\" prepared statement: %w", err)
		}
	}

	for _, filter := range filters {
		if filter.Name == nil {
			return nil, fmt.Errorf("Cannot filter on empty SourceFilter")
		} else {
			return nil, fmt.Errorf("No statement exists for the given Filter")
		}
	}

	// Select.
	var rows *sql.Rows
	if sqlStmt != nil {
		rows, err = sqlStmt.QueryContext(ctx, args...)
	} else {
		queryStr := strings.Join(queryParts[:], "ORDER BY")
		rows, err = db.QueryContext(ctx, queryStr, args...)
	}

	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var identifier string
		err := rows.Scan(&identifier)
		if err != nil {
			return nil, err
		}

		names = append(names, identifier)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"sources\" table: %w", err)
	}

	return names, nil
}

// CreateSource adds a new source to the database.
// generator: source Create
func CreateSource(ctx context.Context, db dbtx, object migration.Source) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	args := make([]any, 3)

	// Populate the statement arguments.
	args[0] = object.Name
	args[1] = object.SourceType
	args[2] = object.Properties

	// Prepared statement to use.
	stmt, err := Stmt(db, sourceCreate)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"sourceCreate\" prepared statement: %w", err)
	}

	// Execute the statement.
	result, err := stmt.Exec(args...)
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return -1, ErrConflict
		}
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to create \"sources\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"sources\" entry ID: %w", err)
	}

	return id, nil
}

// UpdateSource updates the source matching the given key parameters.
// generator: source Update
func UpdateSource(ctx context.Context, db tx, name string, object migration.Source) (_err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	id, err := GetSourceID(ctx, db, name)
	if err != nil {
		return err
	}

	stmt, err := Stmt(db, sourceUpdate)
	if err != nil {
		return fmt.Errorf("Failed to get \"sourceUpdate\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(object.Name, object.SourceType, object.Properties, id)
	if err != nil {
		return fmt.Errorf("Update \"sources\" entry failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}

// RenameSource renames the source matching the given key parameters.
// generator: source Rename
func RenameSource(ctx context.Context, db dbtx, name string, to string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	stmt, err := Stmt(db, sourceRename)
	if err != nil {
		return fmt.Errorf("Failed to get \"sourceRename\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(to, name)
	if err != nil {
		return fmt.Errorf("Rename Source failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows failed: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}

	return nil
}

// DeleteSource deletes the source matching the given key parameters.
// generator: source DeleteOne-by-Name
func DeleteSource(ctx context.Context, db dbtx, name string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Source")
	}()

	stmt, err := Stmt(db, sourceDeleteByName)
	if err != nil {
		return fmt.Errorf("Failed to get \"sourceDeleteByName\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("Delete \"sources\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n == 0 {
		return ErrNotFound
	} else if n > 1 {
		return fmt.Errorf("Query deleted %d Source rows instead of 1", n)
	}

	return nil
}
