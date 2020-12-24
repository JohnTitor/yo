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

// SnakeCase represents a row from 'snake_cases'.
type SnakeCase struct {
	ID        int64  `spanner:"id" json:"id"`                   // id
	StringID  string `spanner:"string_id" json:"string_id"`     // string_id
	FooBarBaz int64  `spanner:"foo_bar_baz" json:"foo_bar_baz"` // foo_bar_baz
}

func SnakeCasePrimaryKeys() []string {
	return []string{
		"id",
	}
}

func SnakeCaseColumns() []string {
	return []string{
		"id",
		"string_id",
		"foo_bar_baz",
	}
}

func (sc *SnakeCase) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "id":
			ret = append(ret, &sc.ID)
		case "string_id":
			ret = append(ret, &sc.StringID)
		case "foo_bar_baz":
			ret = append(ret, &sc.FooBarBaz)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (sc *SnakeCase) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "id":
			ret = append(ret, sc.ID)
		case "string_id":
			ret = append(ret, sc.StringID)
		case "foo_bar_baz":
			ret = append(ret, sc.FooBarBaz)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newSnakeCase_Decoder returns a decoder which reads a row from *spanner.Row
// into SnakeCase. The decoder is not goroutine-safe. Don't use it concurrently.
func newSnakeCase_Decoder(cols []string) func(*spanner.Row) (*SnakeCase, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*SnakeCase, error) {
		var sc SnakeCase
		ptrs, err := sc.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &sc, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (sc *SnakeCase) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("snake_cases", SnakeCaseColumns(), []interface{}{
		sc.ID, sc.StringID, sc.FooBarBaz,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (sc *SnakeCase) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("snake_cases", SnakeCaseColumns(), []interface{}{
		sc.ID, sc.StringID, sc.FooBarBaz,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (sc *SnakeCase) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("snake_cases", SnakeCaseColumns(), []interface{}{
		sc.ID, sc.StringID, sc.FooBarBaz,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (sc *SnakeCase) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, SnakeCasePrimaryKeys()...)

	values, err := sc.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "SnakeCase.UpdateColumns", "snake_cases", err)
	}

	return spanner.Update("snake_cases", colsWithPKeys, values), nil
}

// FindSnakeCase gets a SnakeCase by primary key
func FindSnakeCase(ctx context.Context, db YORODB, id int64) (*SnakeCase, error) {
	key := spanner.Key{id}
	row, err := db.ReadRow(ctx, "snake_cases", key, SnakeCaseColumns())
	if err != nil {
		return nil, newError("FindSnakeCase", "snake_cases", err)
	}

	decoder := newSnakeCase_Decoder(SnakeCaseColumns())
	sc, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindSnakeCase", "snake_cases", err)
	}

	return sc, nil
}

// ReadSnakeCase retrieves multiples rows from SnakeCase by KeySet as a slice.
func ReadSnakeCase(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*SnakeCase, error) {
	var res []*SnakeCase

	decoder := newSnakeCase_Decoder(SnakeCaseColumns())

	rows := db.Read(ctx, "snake_cases", keys, SnakeCaseColumns())
	err := rows.Do(func(row *spanner.Row) error {
		sc, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, sc)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadSnakeCase", "snake_cases", err)
	}

	return res, nil
}

// Delete deletes the SnakeCase from the database.
func (sc *SnakeCase) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := sc.columnsToValues(SnakeCasePrimaryKeys())
	return spanner.Delete("snake_cases", spanner.Key(values))
}

// FindSnakeCasesBySnakeCasesByStringID retrieves multiple rows from 'snake_cases' as a slice of SnakeCase.
//
// Generated from index 'snake_cases_by_string_id'.
func FindSnakeCasesBySnakeCasesByStringID(ctx context.Context, db YORODB, stringID string, fooBarBaz int64) ([]*SnakeCase, error) {
	const sqlstr = "SELECT " +
		"id, string_id, foo_bar_baz " +
		"FROM snake_cases@{FORCE_INDEX=snake_cases_by_string_id} " +
		"WHERE string_id = @param0 AND foo_bar_baz = @param1"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = stringID
	stmt.Params["param1"] = fooBarBaz

	decoder := newSnakeCase_Decoder(SnakeCaseColumns())

	// run query
	YOLog(ctx, sqlstr, stringID, fooBarBaz)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*SnakeCase{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindSnakeCasesBySnakeCasesByStringID", "snake_cases", err)
		}

		sc, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindSnakeCasesBySnakeCasesByStringID", "snake_cases", err)
		}

		res = append(res, sc)
	}

	return res, nil
}

// ReadSnakeCasesBySnakeCasesByStringID retrieves multiples rows from 'snake_cases' by KeySet as a slice.
//
// This does not retrives all columns of 'snake_cases' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'snake_cases_by_string_id'.
func ReadSnakeCasesBySnakeCasesByStringID(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*SnakeCase, error) {
	var res []*SnakeCase
	columns := []string{
		"id",
		"string_id",
		"foo_bar_baz",
	}

	decoder := newSnakeCase_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "snake_cases", "snake_cases_by_string_id", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		sc, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, sc)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadSnakeCasesBySnakeCasesByStringID", "snake_cases", err)
	}

	return res, nil
}
