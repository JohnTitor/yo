// Code generated by yo. DO NOT EDIT.

// Package customtypes contains the types.
package customtypes

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

// CompositePrimaryKey represents a row from 'CompositePrimaryKeys'.
type CompositePrimaryKey struct {
	ID    uint64 `spanner:"Id" json:"Id"`       // Id
	PKey1 string `spanner:"PKey1" json:"PKey1"` // PKey1
	PKey2 uint32 `spanner:"PKey2" json:"PKey2"` // PKey2
	Error int8   `spanner:"Error" json:"Error"` // Error
	X     string `spanner:"X" json:"X"`         // X
	Y     string `spanner:"Y" json:"Y"`         // Y
	Z     string `spanner:"Z" json:"Z"`         // Z
}

func CompositePrimaryKeyPrimaryKeys() []string {
	return []string{
		"PKey1",
		"PKey2",
	}
}

func CompositePrimaryKeyColumns() []string {
	return []string{
		"Id",
		"PKey1",
		"PKey2",
		"Error",
		"X",
		"Y",
		"Z",
	}
}

func (cpk *CompositePrimaryKey) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "Id":
			ret = append(ret, &cpk.ID)
		case "PKey1":
			ret = append(ret, &cpk.PKey1)
		case "PKey2":
			ret = append(ret, &cpk.PKey2)
		case "Error":
			ret = append(ret, &cpk.Error)
		case "X":
			ret = append(ret, &cpk.X)
		case "Y":
			ret = append(ret, &cpk.Y)
		case "Z":
			ret = append(ret, &cpk.Z)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (cpk *CompositePrimaryKey) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "Id":
			ret = append(ret, int64(cpk.ID))
		case "PKey1":
			ret = append(ret, cpk.PKey1)
		case "PKey2":
			ret = append(ret, int64(cpk.PKey2))
		case "Error":
			ret = append(ret, int64(cpk.Error))
		case "X":
			ret = append(ret, cpk.X)
		case "Y":
			ret = append(ret, cpk.Y)
		case "Z":
			ret = append(ret, cpk.Z)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newCompositePrimaryKey_Decoder returns a decoder which reads a row from *spanner.Row
// into CompositePrimaryKey. The decoder is not goroutine-safe. Don't use it concurrently.
func newCompositePrimaryKey_Decoder(cols []string) func(*spanner.Row) (*CompositePrimaryKey, error) {
	var cID int64
	var cPKey2 int64
	var cError int64
	customPtrs := map[string]interface{}{
		"Id":    &cID,
		"PKey2": &cPKey2,
		"Error": &cError,
	}

	return func(row *spanner.Row) (*CompositePrimaryKey, error) {
		var cpk CompositePrimaryKey
		ptrs, err := cpk.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}
		cpk.ID = uint64(cID)
		cpk.PKey2 = uint32(cPKey2)
		cpk.Error = int8(cError)

		return &cpk, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (cpk *CompositePrimaryKey) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("CompositePrimaryKeys", CompositePrimaryKeyColumns(), []interface{}{
		int64(cpk.ID), cpk.PKey1, int64(cpk.PKey2), int64(cpk.Error), cpk.X, cpk.Y, cpk.Z,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (cpk *CompositePrimaryKey) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("CompositePrimaryKeys", CompositePrimaryKeyColumns(), []interface{}{
		int64(cpk.ID), cpk.PKey1, int64(cpk.PKey2), int64(cpk.Error), cpk.X, cpk.Y, cpk.Z,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (cpk *CompositePrimaryKey) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("CompositePrimaryKeys", CompositePrimaryKeyColumns(), []interface{}{
		int64(cpk.ID), cpk.PKey1, int64(cpk.PKey2), int64(cpk.Error), cpk.X, cpk.Y, cpk.Z,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (cpk *CompositePrimaryKey) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, CompositePrimaryKeyPrimaryKeys()...)

	values, err := cpk.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "CompositePrimaryKey.UpdateColumns", "CompositePrimaryKeys", err)
	}

	return spanner.Update("CompositePrimaryKeys", colsWithPKeys, values), nil
}

// FindCompositePrimaryKey gets a CompositePrimaryKey by primary key
func FindCompositePrimaryKey(ctx context.Context, db YORODB, pKey1 string, pKey2 uint32) (*CompositePrimaryKey, error) {
	key := spanner.Key{pKey1, int64(pKey2)}
	row, err := db.ReadRow(ctx, "CompositePrimaryKeys", key, CompositePrimaryKeyColumns())
	if err != nil {
		return nil, newError("FindCompositePrimaryKey", "CompositePrimaryKeys", err)
	}

	decoder := newCompositePrimaryKey_Decoder(CompositePrimaryKeyColumns())
	cpk, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindCompositePrimaryKey", "CompositePrimaryKeys", err)
	}

	return cpk, nil
}

// ReadCompositePrimaryKey retrieves multiples rows from CompositePrimaryKey by KeySet as a slice.
func ReadCompositePrimaryKey(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*CompositePrimaryKey, error) {
	var res []*CompositePrimaryKey

	decoder := newCompositePrimaryKey_Decoder(CompositePrimaryKeyColumns())

	rows := db.Read(ctx, "CompositePrimaryKeys", keys, CompositePrimaryKeyColumns())
	err := rows.Do(func(row *spanner.Row) error {
		cpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, cpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCompositePrimaryKey", "CompositePrimaryKeys", err)
	}

	return res, nil
}

// Delete deletes the CompositePrimaryKey from the database.
func (cpk *CompositePrimaryKey) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := cpk.columnsToValues(CompositePrimaryKeyPrimaryKeys())
	return spanner.Delete("CompositePrimaryKeys", spanner.Key(values))
}

// FindCompositePrimaryKeysByCompositePrimaryKeysByError retrieves multiple rows from 'CompositePrimaryKeys' as a slice of CompositePrimaryKey.
//
// Generated from index 'CompositePrimaryKeysByError'.
func FindCompositePrimaryKeysByCompositePrimaryKeysByError(ctx context.Context, db YORODB, e int8) ([]*CompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CompositePrimaryKeys@{FORCE_INDEX=CompositePrimaryKeysByError} " +
		"WHERE Error = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = int64(e)

	decoder := newCompositePrimaryKey_Decoder(CompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, e)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCompositePrimaryKeysByCompositePrimaryKeysByError", "CompositePrimaryKeys", err)
		}

		cpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCompositePrimaryKeysByCompositePrimaryKeysByError", "CompositePrimaryKeys", err)
		}

		res = append(res, cpk)
	}

	return res, nil
}

// ReadCompositePrimaryKeysByCompositePrimaryKeysByError retrieves multiples rows from 'CompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrives all columns of 'CompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CompositePrimaryKeysByError'.
func ReadCompositePrimaryKeysByCompositePrimaryKeysByError(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*CompositePrimaryKey, error) {
	var res []*CompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"Error",
	}

	decoder := newCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CompositePrimaryKeys", "CompositePrimaryKeysByError", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		cpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, cpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCompositePrimaryKeysByCompositePrimaryKeysByError", "CompositePrimaryKeys", err)
	}

	return res, nil
}

// FindCompositePrimaryKeysByCompositePrimaryKeysByError2 retrieves multiple rows from 'CompositePrimaryKeys' as a slice of CompositePrimaryKey.
//
// Generated from index 'CompositePrimaryKeysByError2'.
func FindCompositePrimaryKeysByCompositePrimaryKeysByError2(ctx context.Context, db YORODB, e int8) ([]*CompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CompositePrimaryKeys@{FORCE_INDEX=CompositePrimaryKeysByError2} " +
		"WHERE Error = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = int64(e)

	decoder := newCompositePrimaryKey_Decoder(CompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, e)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCompositePrimaryKeysByCompositePrimaryKeysByError2", "CompositePrimaryKeys", err)
		}

		cpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCompositePrimaryKeysByCompositePrimaryKeysByError2", "CompositePrimaryKeys", err)
		}

		res = append(res, cpk)
	}

	return res, nil
}

// ReadCompositePrimaryKeysByCompositePrimaryKeysByError2 retrieves multiples rows from 'CompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrives all columns of 'CompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CompositePrimaryKeysByError2'.
func ReadCompositePrimaryKeysByCompositePrimaryKeysByError2(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*CompositePrimaryKey, error) {
	var res []*CompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"Error",
		"Z",
	}

	decoder := newCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CompositePrimaryKeys", "CompositePrimaryKeysByError2", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		cpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, cpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCompositePrimaryKeysByCompositePrimaryKeysByError2", "CompositePrimaryKeys", err)
	}

	return res, nil
}

// FindCompositePrimaryKeysByCompositePrimaryKeysByError3 retrieves multiple rows from 'CompositePrimaryKeys' as a slice of CompositePrimaryKey.
//
// Generated from index 'CompositePrimaryKeysByError3'.
func FindCompositePrimaryKeysByCompositePrimaryKeysByError3(ctx context.Context, db YORODB, e int8) ([]*CompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CompositePrimaryKeys@{FORCE_INDEX=CompositePrimaryKeysByError3} " +
		"WHERE Error = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = int64(e)

	decoder := newCompositePrimaryKey_Decoder(CompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, e)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCompositePrimaryKeysByCompositePrimaryKeysByError3", "CompositePrimaryKeys", err)
		}

		cpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCompositePrimaryKeysByCompositePrimaryKeysByError3", "CompositePrimaryKeys", err)
		}

		res = append(res, cpk)
	}

	return res, nil
}

// ReadCompositePrimaryKeysByCompositePrimaryKeysByError3 retrieves multiples rows from 'CompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrives all columns of 'CompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CompositePrimaryKeysByError3'.
func ReadCompositePrimaryKeysByCompositePrimaryKeysByError3(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*CompositePrimaryKey, error) {
	var res []*CompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"Error",
		"Z",
		"Y",
	}

	decoder := newCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CompositePrimaryKeys", "CompositePrimaryKeysByError3", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		cpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, cpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCompositePrimaryKeysByCompositePrimaryKeysByError3", "CompositePrimaryKeys", err)
	}

	return res, nil
}

// FindCompositePrimaryKeysByCompositePrimaryKeysByXY retrieves multiple rows from 'CompositePrimaryKeys' as a slice of CompositePrimaryKey.
//
// Generated from index 'CompositePrimaryKeysByXY'.
func FindCompositePrimaryKeysByCompositePrimaryKeysByXY(ctx context.Context, db YORODB, x string, y string) ([]*CompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CompositePrimaryKeys@{FORCE_INDEX=CompositePrimaryKeysByXY} " +
		"WHERE X = @param0 AND Y = @param1"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = x
	stmt.Params["param1"] = y

	decoder := newCompositePrimaryKey_Decoder(CompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, x, y)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCompositePrimaryKeysByCompositePrimaryKeysByXY", "CompositePrimaryKeys", err)
		}

		cpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCompositePrimaryKeysByCompositePrimaryKeysByXY", "CompositePrimaryKeys", err)
		}

		res = append(res, cpk)
	}

	return res, nil
}

// ReadCompositePrimaryKeysByCompositePrimaryKeysByXY retrieves multiples rows from 'CompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrives all columns of 'CompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CompositePrimaryKeysByXY'.
func ReadCompositePrimaryKeysByCompositePrimaryKeysByXY(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*CompositePrimaryKey, error) {
	var res []*CompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"X",
		"Y",
	}

	decoder := newCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CompositePrimaryKeys", "CompositePrimaryKeysByXY", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		cpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, cpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCompositePrimaryKeysByCompositePrimaryKeysByXY", "CompositePrimaryKeys", err)
	}

	return res, nil
}
