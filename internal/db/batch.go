package db

import (
	"database/sql"
	"fmt"

	"github.com/FuturFusion/migration-manager/internal/batch"
)

func (n *Node) AddBatch(tx *sql.Tx, b batch.Batch) error {
	internalBatch, ok := b.(*batch.InternalBatch)
	if !ok {
		return fmt.Errorf("Wasn't given an InternalBatch?")
	}

	// Add batch to the database.
	q := `INSERT INTO batches (name,status,includeregex,excluderegex,migrationwindowstart,migrationwindowend) VALUES(?,?,?,?,?,?)`

	marshalledMigrationWindowStart, err := internalBatch.MigrationWindowStart.MarshalText()
	if err != nil {
		return err
	}
	marshalledMigrationWindowEnd, err := internalBatch.MigrationWindowEnd.MarshalText()
	if err != nil {
		return err
	}
	result, err := tx.Exec(q, internalBatch.Name, internalBatch.Status, internalBatch.IncludeRegex, internalBatch.ExcludeRegex, marshalledMigrationWindowStart, marshalledMigrationWindowEnd)
	if err != nil {
		return err
	}

	// Set the new ID assigned to the batch.
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	internalBatch.DatabaseID = int(lastInsertId)

	return nil
}

func (n *Node) GetBatch(tx *sql.Tx, name string) (batch.Batch, error) {
	ret, err := n.getBatchesHelper(tx, name)
	if err != nil {
		return nil, err
	}

	if len(ret) != 1 {
		return nil, fmt.Errorf("No batch exists with name '%s'", name)
	}

	return ret[0], nil
}

func (n *Node) GetAllBatches(tx *sql.Tx) ([]batch.Batch, error) {
	return n.getBatchesHelper(tx, "")
}

func (n *Node) DeleteBatch(tx *sql.Tx, name string) error {
	// Delete the batch from the database.
	q := `DELETE FROM batches WHERE name=?`
	result, err := tx.Exec(q, name)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("Batch with name '%s' doesn't exist, can't delete", name)
	}

	return nil
}

func (n *Node) UpdateBatch(tx *sql.Tx, b batch.Batch) error {
	// Update batch in the database.
	q := `UPDATE batches SET name=?,status=?,includeregex=?,excluderegex=?,migrationwindowstart=?,migrationwindowend=? WHERE id=?`

	internalBatch, ok := b.(*batch.InternalBatch)
	if !ok {
		return fmt.Errorf("Wasn't given an InternalBatch?")
	}

	marshalledMigrationWindowStart, err := internalBatch.MigrationWindowStart.MarshalText()
	if err != nil {
		return err
	}
	marshalledMigrationWindowEnd, err := internalBatch.MigrationWindowEnd.MarshalText()
	if err != nil {
		return err
	}
	result, err := tx.Exec(q, internalBatch.Name, internalBatch.Status, internalBatch.IncludeRegex, internalBatch.ExcludeRegex, marshalledMigrationWindowStart, marshalledMigrationWindowEnd, internalBatch.DatabaseID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("Batch with ID %d doesn't exist, can't update", internalBatch.DatabaseID)
	}

	return nil
}

func (n *Node) getBatchesHelper(tx *sql.Tx, name string) ([]batch.Batch, error) {
	ret := []batch.Batch{}

	// Get all batches in the database.
	q := `SELECT id,name,status,includeregex,excluderegex,migrationwindowstart,migrationwindowend FROM batches`
	var rows *sql.Rows
	var err error
	if name != "" {
		q += ` WHERE name=?`
		rows, err = tx.Query(q, name)
	} else {
		q += ` ORDER BY name`
		rows, err = tx.Query(q)
	}
	if err != nil {
		return ret, err
	}

	for rows.Next() {
		newBatch := &batch.InternalBatch{}
		marshalledMigrationWindowStart := ""
		marshalledMigrationWindowEnd := ""

		err := rows.Scan(&newBatch.DatabaseID, &newBatch.Name, &newBatch.Status, &newBatch.IncludeRegex, &newBatch.ExcludeRegex, &marshalledMigrationWindowStart, &marshalledMigrationWindowEnd)
		if err != nil {
			return nil, err
		}
		err = newBatch.MigrationWindowStart.UnmarshalText([]byte(marshalledMigrationWindowStart))
		if err != nil {
			return nil, err
		}
		err = newBatch.MigrationWindowEnd.UnmarshalText([]byte(marshalledMigrationWindowEnd))
		if err != nil {
			return nil, err
		}

		ret = append(ret, newBatch)
	}

	return ret, nil
}
