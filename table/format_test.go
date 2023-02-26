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
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Table_Markdown tests the Markdown format.
func Test_Table_Markdown(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("Column 1", AtEnd)
	table.InsertColumn("Column 2", AtEnd)
	table.InsertColumn("Column 3", AtEnd)
	assert.NoError(t, table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 2A ", "Row 2B  ", "Row 2C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C   "}, AtEnd))
	tableData := &bytes.Buffer{}
	formatter := MarkdownFormatter()
	assert.NoError(t, formatter.Format(tableData, table))
	assert.Equal(t, `| Column 1 | Column 2 | Column 3  |
|----------|----------|-----------|
| Row 1A   | Row 1B   | Row 1C    |
| Row 2A   | Row 2B   | Row 2C    |
| Row 3A   | Row 3B   | Row 3C    |
`, tableData.String())
}

// Test_Table_CSV tests the CSV format.
func Test_Table_CSV(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("Column 1", AtEnd)
	table.InsertColumn("Column 2", AtEnd)
	table.InsertColumn("Column 3", AtEnd)
	assert.NoError(t, table.InsertRow([]interface{}{"Row 1A  ", "Row 1B", "Row 1C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 2A", "Row 2B ", "Row 2C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd))
	tableData := &bytes.Buffer{}
	formatter := CSVFormatter()
	assert.NoError(t, formatter.Format(tableData, table))
	assert.Equal(t, `Column 1,Column 2,Column 3
Row 1A  ,Row 1B,Row 1C
Row 2A,Row 2B ,Row 2C
Row 3A,Row 3B,Row 3C
`, tableData.String())
}

// Test_Table_Markdown_NoHeader tests the NoHeader format.
func Test_Table_Markdown_NoHeader(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("Column 1", AtEnd)
	table.InsertColumn("Column 2", AtEnd)
	table.InsertColumn("Column 3", AtEnd)
	assert.NoError(t, table.InsertRow([]interface{}{"Row 1A  ", "Row 1B  ", "Row 1C  "}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 2A  ", "Row 2B  ", "Row 2C  "}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 3A  ", "Row 3B  ", "Row 3C  "}, AtEnd))
	tableData := &bytes.Buffer{}
	formatter := MarkdownFormatter(NoHeader)
	assert.NoError(t, formatter.Format(tableData, table))
	assert.Equal(t, `| Row 1A   | Row 1B   | Row 1C   |
| Row 2A   | Row 2B   | Row 2C   |
| Row 3A   | Row 3B   | Row 3C   |
`, tableData.String())
}

// Test_Table_CSV_NoHeader tests the NoHeader format.
func Test_Table_CSV_NoHeader(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("Column 1", AtEnd)
	table.InsertColumn("Column 2", AtEnd)
	table.InsertColumn("Column 3", AtEnd)
	assert.NoError(t, table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 2A", "Row 2B", "Row 2C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd))
	tableData := &bytes.Buffer{}
	formatter := CSVFormatter(NoHeader)
	assert.NoError(t, formatter.Format(tableData, table))
	assert.Equal(t, `Row 1A,Row 1B,Row 1C
Row 2A,Row 2B,Row 2C
Row 3A,Row 3B,Row 3C
`, tableData.String())
}

// Test_Table_Markdown_NoFooter tests the NoFooter format.
func Test_Table_Markdown_NoFooter(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("Column 1", AtEnd)
	table.InsertColumn("Column 2", AtEnd)
	table.InsertColumn("Column 3", AtEnd)
	assert.NoError(t, table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 2A", "Row 2B", "Row 2C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd))
	tableData := &bytes.Buffer{}
	formatter := MarkdownFormatter(NoFooter)
	assert.NoError(t, formatter.Format(tableData, table))
	assert.Equal(t, `| Column 1 | Column 2 | Column 3 |
|----------|----------|----------|
| Row 1A   | Row 1B   | Row 1C   |
| Row 2A   | Row 2B   | Row 2C   |
| Row 3A   | Row 3B   | Row 3C   |
`, tableData.String())
}

// errWriter is an io.Writer that always returns an error.
type errWriter struct{}

func (w *errWriter) Write(p []byte) (n int, err error) {
	return 0, assert.AnError
}

// Test_MustFprintf_Panic tests that MustFprintf panics when the underlying
// Fprintf returns an error.
func Test_MustFprintf_Panic(t *testing.T) {
	t.Parallel()
	assert.Panics(t, func() {
		MustFprintf(&errWriter{}, "Hello, %s!", "world")
	})
}

// Test_Format_Panic tests that Format returns an error when the underlying
// Fprintf returns an error.
func Test_Format_Panic(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("Column 1", AtEnd)
	table.InsertColumn("Column 2", AtEnd)
	table.InsertColumn("Column 3", AtEnd)
	assert.NoError(t, table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 2A", "Row 2B", "Row 2C"}, AtEnd))
	assert.NoError(t, table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd))
	formatter := MarkdownFormatter()
	err := formatter.Format(&errWriter{}, table)
	assert.Error(t, err)
}
