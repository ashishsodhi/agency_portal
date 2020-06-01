// Package models contains the types for schema 'agency_portal'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// TriageHistory represents a row from 'agency_portal.TRIAGE_HISTORY'.
type TriageHistory struct {
	UUID         []byte    `json:"UUID"`           // UUID
	TriageUUID   []byte    `json:"TRIAGE_UUID"`    // TRIAGE_UUID
	Name         string    `json:"NAME"`           // NAME
	Zip          JSON      `json:"ZIP"`            // ZIP
	OpenJobTimer JSON      `json:"OPEN_JOB_TIMER"` // OPEN_JOB_TIMER
	TimeTable    JSON      `json:"TIME_TABLE"`     // TIME_TABLE
	CsrUUID      []byte    `json:"CSR_UUID"`       // CSR_UUID
	WhenCreated  time.Time `json:"WHEN_CREATED"`   // WHEN_CREATED
	Tlm          time.Time `json:"TLM"`            // TLM

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TriageHistory exists in the database.
func (th *TriageHistory) Exists() bool {
	return th._exists
}

// Deleted provides information if the TriageHistory has been deleted from the database.
func (th *TriageHistory) Deleted() bool {
	return th._deleted
}

// Insert inserts the TriageHistory to the database.
func (th *TriageHistory) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if th._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO agency_portal.TRIAGE_HISTORY (` +
		`UUID, TRIAGE_UUID, NAME, ZIP, OPEN_JOB_TIMER, TIME_TABLE, CSR_UUID, WHEN_CREATED, TLM` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, th.UUID, th.TriageUUID, th.Name, th.Zip, th.OpenJobTimer, th.TimeTable, th.CsrUUID, th.WhenCreated, th.Tlm)
	_, err = db.Exec(sqlstr, th.UUID, th.TriageUUID, th.Name, th.Zip, th.OpenJobTimer, th.TimeTable, th.CsrUUID, th.WhenCreated, th.Tlm)
	if err != nil {
		return err
	}

	// set existence
	th._exists = true

	return nil
}

// Update updates the TriageHistory in the database.
func (th *TriageHistory) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !th._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if th._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE agency_portal.TRIAGE_HISTORY SET ` +
		`TRIAGE_UUID = ?, NAME = ?, ZIP = ?, OPEN_JOB_TIMER = ?, TIME_TABLE = ?, CSR_UUID = ?, WHEN_CREATED = ?, TLM = ?` +
		` WHERE UUID = ?`

	// run query
	XOLog(sqlstr, th.TriageUUID, th.Name, th.Zip, th.OpenJobTimer, th.TimeTable, th.CsrUUID, th.WhenCreated, th.Tlm, th.UUID)
	_, err = db.Exec(sqlstr, th.TriageUUID, th.Name, th.Zip, th.OpenJobTimer, th.TimeTable, th.CsrUUID, th.WhenCreated, th.Tlm, th.UUID)
	return err
}

// Save saves the TriageHistory to the database.
func (th *TriageHistory) Save(db XODB) error {
	if th.Exists() {
		return th.Update(db)
	}

	return th.Insert(db)
}

// Delete deletes the TriageHistory from the database.
func (th *TriageHistory) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !th._exists {
		return nil
	}

	// if deleted, bail
	if th._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM agency_portal.TRIAGE_HISTORY WHERE UUID = ?`

	// run query
	XOLog(sqlstr, th.UUID)
	_, err = db.Exec(sqlstr, th.UUID)
	if err != nil {
		return err
	}

	// set deleted
	th._deleted = true

	return nil
}

// Triage returns the Triage associated with the TriageHistory's TriageUUID (TRIAGE_UUID).
//
// Generated from foreign key 'FK_TRIAGE_HISTORY_TRIAGE_UUID'.
func (th *TriageHistory) Triage(db XODB) (*Triage, error) {
	return TriageByUUID(db, th.TriageUUID)
}

// TriageHistoriesByTriageUUID retrieves a row from 'agency_portal.TRIAGE_HISTORY' as a TriageHistory.
//
// Generated from index 'FK_TRIAGE_HISTORY_TRIAGE_UUID'.
func TriageHistoriesByTriageUUID(db XODB, triageUUID []byte) ([]*TriageHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`UUID, TRIAGE_UUID, NAME, ZIP, OPEN_JOB_TIMER, TIME_TABLE, CSR_UUID, WHEN_CREATED, TLM ` +
		`FROM agency_portal.TRIAGE_HISTORY ` +
		`WHERE TRIAGE_UUID = ?`

	// run query
	XOLog(sqlstr, triageUUID)
	q, err := db.Query(sqlstr, triageUUID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TriageHistory{}
	for q.Next() {
		th := TriageHistory{
			_exists: true,
		}

		// scan
		err = q.Scan(&th.UUID, &th.TriageUUID, &th.Name, &th.Zip, &th.OpenJobTimer, &th.TimeTable, &th.CsrUUID, &th.WhenCreated, &th.Tlm)
		if err != nil {
			return nil, err
		}

		res = append(res, &th)
	}

	return res, nil
}

// TriageHistoriesByTlm retrieves a row from 'agency_portal.TRIAGE_HISTORY' as a TriageHistory.
//
// Generated from index 'IX_TLM'.
func TriageHistoriesByTlm(db XODB, tlm time.Time) ([]*TriageHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`UUID, TRIAGE_UUID, NAME, ZIP, OPEN_JOB_TIMER, TIME_TABLE, CSR_UUID, WHEN_CREATED, TLM ` +
		`FROM agency_portal.TRIAGE_HISTORY ` +
		`WHERE TLM = ?`

	// run query
	XOLog(sqlstr, tlm)
	q, err := db.Query(sqlstr, tlm)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TriageHistory{}
	for q.Next() {
		th := TriageHistory{
			_exists: true,
		}

		// scan
		err = q.Scan(&th.UUID, &th.TriageUUID, &th.Name, &th.Zip, &th.OpenJobTimer, &th.TimeTable, &th.CsrUUID, &th.WhenCreated, &th.Tlm)
		if err != nil {
			return nil, err
		}

		res = append(res, &th)
	}

	return res, nil
}

// TriageHistoryByUUID retrieves a row from 'agency_portal.TRIAGE_HISTORY' as a TriageHistory.
//
// Generated from index 'TRIAGE_HISTORY_UUID_pkey'.
func TriageHistoryByUUID(db XODB, uuid []byte) (*TriageHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`UUID, TRIAGE_UUID, NAME, ZIP, OPEN_JOB_TIMER, TIME_TABLE, CSR_UUID, WHEN_CREATED, TLM ` +
		`FROM agency_portal.TRIAGE_HISTORY ` +
		`WHERE UUID = ?`

	// run query
	XOLog(sqlstr, uuid)
	th := TriageHistory{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, uuid).Scan(&th.UUID, &th.TriageUUID, &th.Name, &th.Zip, &th.OpenJobTimer, &th.TimeTable, &th.CsrUUID, &th.WhenCreated, &th.Tlm)
	if err != nil {
		return nil, err
	}

	return &th, nil
}
