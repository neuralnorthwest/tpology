// Copyright 2023 Scott M. Long
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package table

import (
	"io"
)

// Table is a table.
type Table struct {
	// Columns are the table columns.
	Columns []*Column
	// Rows are the table rows.
	Rows []*Row
}

// Column is a table column.
type Column struct {
	// Name is the column name.
	Name string
	// Width is the column width.
	Width int
}

// Row is a table row.
type Row struct {
	// Cells are the row cells.
	Cells []Cell
}

// Cell is a table cell.
type Cell struct {
	// Value is the cell value.
	Value interface{}
	// Width is the cell width.
	Width int
}

// New returns a new table.
func New() *Table {
	return &Table{}
}

// ColumnNames returns the column names.
func (t *Table) ColumnNames() []string {
	names := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		names[i] = c.Name
	}
	return names
}

// UpdateWidths updates the column widths.
func (t *Table) UpdateWidths() {
	for _, c := range t.Columns {
		c.Width = len(c.Name)
	}
	for _, r := range t.Rows {
		for i, c := range r.Cells {
			if c.Width > t.Columns[i].Width {
				t.Columns[i].Width = c.Width
			}
		}
	}
}

// Write writes the table to the writer.
func (t *Table) Write(w io.Writer, formatter *Formatter) error {
	return formatter.Format(w, t)
}
