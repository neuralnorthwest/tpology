package table

import "fmt"

// insertionPoint represents an insertion point.
type insertionPoint interface {
	// indexColumn returns the insertion point index for a column.
	indexColumn([]*Column) int
	// indexRow returns the insertion point index for a row.
	indexRow([]*Row) int
}

// atBeginning is an insertion point at the beginning.
type atBeginning struct{}

// AtBeginning is an insertion point at the beginning.
var AtBeginning = &atBeginning{}

// indexColumn returns the insertion point index.
func (a *atBeginning) indexColumn(columns []*Column) int {
	return 0
}

// indexRow returns the insertion point index.
func (a *atBeginning) indexRow(rows []*Row) int {
	return 0
}

// atEnd is an insertion point at the end.
type atEnd struct{}

// AtEnd is an insertion point at the end.
var AtEnd = &atEnd{}

// indexColumn returns the insertion point index.
func (a *atEnd) indexColumn(columns []*Column) int {
	return len(columns)
}

// indexRow returns the insertion point index.
func (a *atEnd) indexRow(rows []*Row) int {
	return len(rows)
}

// BeforeColumn is an insertion point before a name.
type BeforeColumn string

// indexColumn returns the insertion point index.
func (b BeforeColumn) indexColumn(columns []*Column) int {
	for i, c := range columns {
		if c.Name == string(b) {
			return i
		}
	}
	return len(columns)
}

// indexRow returns the insertion point index.
func (b BeforeColumn) indexRow(rows []*Row) int {
	panic("cannot insert column before row")
}

// AfterColumn is an insertion point after a name.
type AfterColumn string

// indexColumn returns the insertion point index.
func (a AfterColumn) indexColumn(columns []*Column) int {
	for i, c := range columns {
		if c.Name == string(a) {
			return i + 1
		}
	}
	return len(columns)
}

// indexRow returns the insertion point index.
func (a AfterColumn) indexRow(rows []*Row) int {
	panic("cannot insert column after row")
}

// BeforeRow is an insertion point before a row index.
type BeforeRow int

// indexColumn returns the insertion point index.
func (b BeforeRow) indexColumn(columns []*Column) int {
	panic("cannot insert row before column")
}

// indexRow returns the insertion point index.
func (b BeforeRow) indexRow(rows []*Row) int {
	return int(b)
}

// AfterRow is an insertion point after a row index.
type AfterRow int

// indexColumn returns the insertion point index.
func (a AfterRow) indexColumn(columns []*Column) int {
	panic("cannot insert row after column")
}

// index returns the insertion point index.
func (a AfterRow) indexRow(rows []*Row) int {
	return int(a) + 1
}

// InsertColumn inserts a column into the table.
//
//   - name is the column name.
//   - where is the insertion point
//
// Insertion points can be:
//   - AtBeginning
//   - AtEnd
//   - BeforeColumn("name")
//   - AfterColumn("name")
func (t *Table) InsertColumn(name string, where insertionPoint) {
	index := where.indexColumn(t.Columns)
	t.Columns = append(t.Columns, nil)
	copy(t.Columns[index+1:], t.Columns[index:])
	t.Columns[index] = &Column{Name: name, Width: len(name)}
	for _, r := range t.Rows {
		r.Cells = append(r.Cells, Cell{})
		copy(r.Cells[index+1:], r.Cells[index:])
		r.Cells[index] = Cell{Width: 0}
	}
}

// InsertRow inserts a row into the table.
//
//   - values are the row values.
//   - where is the insertion point
//
// Insertion points can be:
//   - AtBeginning
//   - AtEnd
//   - BeforeRow(n)
//   - AfterRow(n)
func (t *Table) InsertRow(values []interface{}, where insertionPoint) error {
	if len(values) != len(t.Columns) {
		return fmt.Errorf("invalid number of values")
	}
	index := where.indexRow(t.Rows)
	t.Rows = append(t.Rows, nil)
	copy(t.Rows[index+1:], t.Rows[index:])
	t.Rows[index] = &Row{Cells: make([]Cell, len(t.Columns))}
	for i, v := range values {
		t.Rows[index].Cells[i] = Cell{Value: v, Width: len(fmt.Sprint(v))}
	}
	return nil
}
