package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Table_InsertColumn tests the InsertColumn function.
func Test_Table_InsertColumn(t *testing.T) {
	t.Parallel()
	table := New()

	table.InsertColumn("foo", AtEnd)
	assert.Equal(t, []*Column{
		{Name: "foo", Width: 3},
	}, table.Columns)

	// Add a Row
	table.Rows = append(table.Rows, &Row{
		Cells: []Cell{
			{Value: "foo", Width: 3},
		},
	})

	table.InsertColumn("bar", AtBeginning)
	assert.Equal(t, []*Column{
		{Name: "bar", Width: 3},
		{Name: "foo", Width: 3},
	}, table.Columns)

	// Verify the Row
	assert.Equal(t, []Cell{
		{Value: nil, Width: 0},
		{Value: "foo", Width: 3},
	}, table.Rows[0].Cells)

	table.InsertColumn("baz", BeforeColumn("foo"))
	assert.Equal(t, []*Column{
		{Name: "bar", Width: 3},
		{Name: "baz", Width: 3},
		{Name: "foo", Width: 3},
	}, table.Columns)

	// Verify the Row
	assert.Equal(t, []Cell{
		{Value: nil, Width: 0},
		{Value: nil, Width: 0},
		{Value: "foo", Width: 3},
	}, table.Rows[0].Cells)

	table.InsertColumn("qux", AfterColumn("bar"))
	assert.Equal(t, []*Column{
		{Name: "bar", Width: 3},
		{Name: "qux", Width: 3},
		{Name: "baz", Width: 3},
		{Name: "foo", Width: 3},
	}, table.Columns)

	// Verify the Row
	assert.Equal(t, []Cell{
		{Value: nil, Width: 0},
		{Value: nil, Width: 0},
		{Value: nil, Width: 0},
		{Value: "foo", Width: 3},
	}, table.Rows[0].Cells)
}

// Test_Table_InsertColumn_BeforeUnknown tests the InsertColumn function with
// Before("unknown") insertion point.
func Test_Table_InsertColumn_BeforeUnknown(t *testing.T) {
	t.Parallel()
	table := New()

	table.InsertColumn("foo", AtEnd)
	table.InsertColumn("bar", BeforeColumn("unknown"))
	assert.Equal(t, []*Column{
		{Name: "foo", Width: 3},
		{Name: "bar", Width: 3},
	}, table.Columns)
}

// Test_Table_InsertColumn_AfterUnknown tests the InsertColumn function with
// After("unknown") insertion point.
func Test_Table_InsertColumn_AfterUnknown(t *testing.T) {
	t.Parallel()
	table := New()

	table.InsertColumn("foo", AtEnd)
	table.InsertColumn("bar", AfterColumn("unknown"))
	assert.Equal(t, []*Column{
		{Name: "foo", Width: 3},
		{Name: "bar", Width: 3},
	}, table.Columns)
}

// Test_Table_InsertRow tests the InsertRow function.
func Test_Table_InsertRow(t *testing.T) {
	t.Parallel()
	table := New()

	table.InsertRow(nil, AtEnd)
	assert.Equal(t, []*Row{
		{Cells: []Cell{}},
	}, table.Rows)

	// Add a Column
	table.InsertColumn("foo", AtEnd)

	table.InsertRow([]interface{}{"foo"}, AtBeginning)
	assert.Equal(t, []*Row{
		{Cells: []Cell{{Value: "foo", Width: 3}}},
		{Cells: []Cell{{Value: nil, Width: 0}}},
	}, table.Rows)

	table.InsertRow([]interface{}{"bar"}, BeforeRow(1))
	assert.Equal(t, []*Row{
		{Cells: []Cell{{Value: "foo", Width: 3}}},
		{Cells: []Cell{{Value: "bar", Width: 3}}},
		{Cells: []Cell{{Value: nil, Width: 0}}},
	}, table.Rows)

	table.InsertRow([]interface{}{"baz"}, AfterRow(0))
	assert.Equal(t, []*Row{
		{Cells: []Cell{{Value: "foo", Width: 3}}},
		{Cells: []Cell{{Value: "baz", Width: 3}}},
		{Cells: []Cell{{Value: "bar", Width: 3}}},
		{Cells: []Cell{{Value: nil, Width: 0}}},
	}, table.Rows)
}

// Test_InsertRow_WrongLength tests that the InsertRow function returns an
// error when the row length is different from the number of columns.
func Test_InsertRow_WrongLength(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("foo", AtEnd)
	assert.Error(t, table.InsertRow([]interface{}{"foo", "bar"}, AtBeginning))
}

// Test_InsertRow_BeforeColumn tests that the InsertRow function with
// BeforeColumn insertion point panics.
func Test_InsertRow_BeforeColumn(t *testing.T) {
	t.Parallel()
	table := New()
	assert.Panics(t, func() {
		table.InsertRow(nil, BeforeColumn("foo"))
	})
}

// Test_InsertRow_AfterColumn tests that the InsertRow function with
// AfterColumn insertion point panics.
func Test_InsertRow_AfterColumn(t *testing.T) {
	t.Parallel()
	table := New()
	assert.Panics(t, func() {
		table.InsertRow(nil, AfterColumn("foo"))
	})
}

// Test_Table_InsertColumn_BeforeRow tests that the InsertColumn function with
// BeforeRow insertion point panics.
func Test_Table_InsertColumn_BeforeRow(t *testing.T) {
	t.Parallel()
	table := New()
	assert.Panics(t, func() {
		table.InsertColumn("foo", BeforeRow(0))
	})
}

// Test_Table_InsertColumn_AfterRow tests that the InsertColumn function with
// AfterRow insertion point panics.
func Test_Table_InsertColumn_AfterRow(t *testing.T) {
	t.Parallel()
	table := New()
	assert.Panics(t, func() {
		table.InsertColumn("foo", AfterRow(0))
	})
}
