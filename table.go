package gosql

//CREATE [ [ GLOBAL | LOCAL ] { TEMPORARY | TEMP } | UNLOGGED ] TABLE [ IF NOT EXISTS ] имя_таблицы ( [
//  { имя_столбца тип_данных [ COLLATE правило_сортировки ] [ ограничение_столбца [ ... ] ]
//    | ограничение_таблицы
//    | LIKE исходная_таблица [ вариант_копирования ... ] }
//    [, ... ]
//] )
//[ INHERITS ( таблица_родитель [, ... ] ) ]
//[ WITH ( параметр_хранения [= значение] [, ... ] ) | WITH OIDS | WITHOUT OIDS ]
//[ ON COMMIT { PRESERVE ROWS | DELETE ROWS | DROP } ]
//[ TABLESPACE табл_пространство ]

// Table query builder
type Table struct {
	// table name
	name string
	// If not exists
	ifNotExists bool
	// is table temporary
	temp bool
	// unLogged table
	unLogged bool
	// inherits
	inherits expression
	// tablespace
	tablespace string
}

// Flags Set create flags
func (t *Table) Flags(ifNotExists, temp, unLogged bool) *Table {
	t.ifNotExists = ifNotExists
	t.unLogged = unLogged
	t.temp = temp
	return t
}

// SetName Set name
func (t *Table) SetName(name string) *Table {
	t.name = name
	return t
}

// GetName get name
func (t *Table) GetName() string {
	return t.name
}

// Inherits inherit form tables
func (t *Table) Inherits() *expression {
	return &t.inherits
}

// SetTableSpace set table space
func (t *Table) SetTableSpace(space string) *Table {
	t.tablespace = space
	return t
}

// GetTableSpace get table space
func (t *Table) GetTableSpace() string {
	return t.tablespace
}

// ResetTableSpace reset table space
func (t *Table) ResetTableSpace() *Table {
	t.tablespace = ""
	return t
}

//[ CONSTRAINT имя_ограничения ]
//{ CHECK ( выражение ) [ NO INHERIT ] |
//  UNIQUE ( имя_столбца [, ... ] ) параметры_индекса |
//  PRIMARY KEY ( имя_столбца [, ... ] ) параметры_индекса |
//  EXCLUDE [ USING индексный_метод ] ( элемент_исключения WITH оператор [, ... ] ) параметры_индекса [ WHERE ( предикат ) ] |
//  FOREIGN KEY ( имя_столбца [, ... ] ) REFERENCES целевая_таблица [ ( целевой_столбец [, ... ] ) ]
//    [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ] [ ON DELETE действие ] [ ON UPDATE действие ] }
//[ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]

// constraint
type constraintTable struct {
	// name
	name string
	// check expression
	check expression
	// unique index
	unique detailedExpression
	// primary key
	primary detailedExpression
	// references
	//references references
	// deferrable
	deferrable *bool
	// initially
	initially string
}
