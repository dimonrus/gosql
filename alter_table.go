package gosql

import "strings"

// ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
//    action [, ... ]
//ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
//    RENAME [ COLUMN ] column_name TO new_column_name
//ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
//    RENAME CONSTRAINT constraint_name TO new_constraint_name
//ALTER TABLE [ IF EXISTS ] name
//    RENAME TO new_name
//ALTER TABLE [ IF EXISTS ] name
//    SET SCHEMA new_schema
//ALTER TABLE ALL IN TABLESPACE name [ OWNED BY role_name [, ... ] ]
//    SET TABLESPACE new_tablespace [ NOWAIT ]
//ALTER TABLE [ IF EXISTS ] name
//    ATTACH PARTITION partition_name { FOR VALUES partition_bound_spec | DEFAULT }
//ALTER TABLE [ IF EXISTS ] name
//    DETACH PARTITION partition_name [ CONCURRENTLY | FINALIZE ]
type alterTable struct {
	// if exists
	ifExists bool
	// only
	only bool
	// name of table
	name string
}

// IfExists set if exists
func (a *alterTable) IfExists() *alterTable {
	a.ifExists = true
	return a
}

// Only set if Only
func (a *alterTable) Only() *alterTable {
	a.only = true
	return a
}

// IsEmpty check if empty
func (a *alterTable) IsEmpty() bool {
	return a == nil
}

// String render alter table query
func (a *alterTable) String() string {
	if a.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	return b.String()
}

type alterTableAction struct {
	// FORCE ROW LEVEL SECURITY
	forceRowLevelSecurity bool
}

// ForceRowLevelSecurity set to tru
func (a *alterTableAction) ForceRowLevelSecurity() *alterTableAction {
	a.forceRowLevelSecurity = true
	return a
}

// IsEmpty check if empty
func (a *alterTableAction) IsEmpty() bool {
	return a == nil
}

// String render alter table query
func (a *alterTableAction) String() string {
	if a.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	return b.String()
}
