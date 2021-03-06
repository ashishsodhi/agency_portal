// Package models contains the types for schema 'agency_portal'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Triage represents a row from 'agency_portal.TRIAGE'.
type Triage struct {
	UUID            []byte    `json:"UUID"`             // UUID
	Name            string    `json:"NAME"`             // NAME
	AcceptanceTimer JSON      `json:"ACCEPTANCE_TIMER"` // ACCEPTANCE_TIMER
	WhenCreated     time.Time `json:"WHEN_CREATED"`     // WHEN_CREATED
	Tlm             time.Time `json:"TLM"`              // TLM

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Triage exists in the database.
func (t *Triage) Exists() bool {
	return t._exists
}

// Deleted provides information if the Triage has been deleted from the database.
func (t *Triage) Deleted() bool {
	return t._deleted
}

// Insert inserts the Triage to the database.
func (t *Triage) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO agency_portal.TRIAGE (` +
		`UUID, NAME, ACCEPTANCE_TIMER, WHEN_CREATED, TLM` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.UUID, t.Name, t.AcceptanceTimer, t.WhenCreated, t.Tlm)
	_, err = db.Exec(sqlstr, t.UUID, t.Name, t.AcceptanceTimer, t.WhenCreated, t.Tlm)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Triage in the database.
func (t *Triage) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE agency_portal.TRIAGE SET ` +
		`NAME = ?, ACCEPTANCE_TIMER = ?, WHEN_CREATED = ?, TLM = ?` +
		` WHERE UUID = ?`

	// run query
	XOLog(sqlstr, t.Name, t.AcceptanceTimer, t.WhenCreated, t.Tlm, t.UUID)
	_, err = db.Exec(sqlstr, t.Name, t.AcceptanceTimer, t.WhenCreated, t.Tlm, t.UUID)
	return err
}

// Save saves the Triage to the database.
func (t *Triage) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Triage from the database.
func (t *Triage) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM agency_portal.TRIAGE WHERE UUID = ?`

	// run query
	XOLog(sqlstr, t.UUID)
	_, err = db.Exec(sqlstr, t.UUID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TriagesByTlm retrieves a row from 'agency_portal.TRIAGE' as a Triage.
//
// Generated from index 'IX_TLM'.
func TriagesByTlm(db XODB, tlm time.Time) ([]*Triage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`UUID, NAME, ACCEPTANCE_TIMER, WHEN_CREATED, TLM ` +
		`FROM agency_portal.TRIAGE ` +
		`WHERE TLM = ?`

	// run query
	XOLog(sqlstr, tlm)
	q, err := db.Query(sqlstr, tlm)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Triage{}
	for q.Next() {
		t := Triage{
			_exists: true,
		}

		// scan
		err = q.Scan(&t.UUID, &t.Name, &t.AcceptanceTimer, &t.WhenCreated, &t.Tlm)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

// TriageByUUID retrieves a row from 'agency_portal.TRIAGE' as a Triage.
//
// Generated from index 'TRIAGE_UUID_pkey'.
func TriageByUUID(db XODB, uuid []byte) (*Triage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`UUID, NAME, ACCEPTANCE_TIMER, WHEN_CREATED, TLM ` +
		`FROM agency_portal.TRIAGE ` +
		`WHERE UUID = ?`

	// run query
	XOLog(sqlstr, uuid)
	t := Triage{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, uuid).Scan(&t.UUID, &t.Name, &t.AcceptanceTimer, &t.WhenCreated, &t.Tlm)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// TriageByName retrieves a row from 'agency_portal.TRIAGE' as a Triage.
//
// Generated from index 'UNQ_NAME'.
func TriageByName(db XODB, name string) (*Triage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`UUID, NAME, ACCEPTANCE_TIMER, WHEN_CREATED, TLM ` +
		`FROM agency_portal.TRIAGE ` +
		`WHERE NAME = ?`

	// run query
	XOLog(sqlstr, name)
	t := Triage{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name).Scan(&t.UUID, &t.Name, &t.AcceptanceTimer, &t.WhenCreated, &t.Tlm)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
