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
	table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 2A ", "Row 2B  ", "Row 2C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C   "}, AtEnd)
	tableData := &bytes.Buffer{}
	formatter := MarkdownFormatter()
	formatter.Format(tableData, table)
	assert.Equal(t, `| Column 1 | Column 2 | Column 3  |
| -------- | -------- | --------- |
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
	table.InsertRow([]interface{}{"Row 1A  ", "Row 1B", "Row 1C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 2A", "Row 2B ", "Row 2C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd)
	tableData := &bytes.Buffer{}
	formatter := CSVFormatter()
	formatter.Format(tableData, table)
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
	table.InsertRow([]interface{}{"Row 1A  ", "Row 1B  ", "Row 1C  "}, AtEnd)
	table.InsertRow([]interface{}{"Row 2A  ", "Row 2B  ", "Row 2C  "}, AtEnd)
	table.InsertRow([]interface{}{"Row 3A  ", "Row 3B  ", "Row 3C  "}, AtEnd)
	tableData := &bytes.Buffer{}
	formatter := MarkdownFormatter(NoHeader)
	formatter.Format(tableData, table)
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
	table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 2A", "Row 2B", "Row 2C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd)
	tableData := &bytes.Buffer{}
	formatter := CSVFormatter(NoHeader)
	formatter.Format(tableData, table)
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
	table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 2A", "Row 2B", "Row 2C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd)
	tableData := &bytes.Buffer{}
	formatter := MarkdownFormatter(NoFooter)
	formatter.Format(tableData, table)
	assert.Equal(t, `| Column 1 | Column 2 | Column 3 |
| -------- | -------- | -------- |
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
	table.InsertRow([]interface{}{"Row 1A", "Row 1B", "Row 1C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 2A", "Row 2B", "Row 2C"}, AtEnd)
	table.InsertRow([]interface{}{"Row 3A", "Row 3B", "Row 3C"}, AtEnd)
	formatter := MarkdownFormatter()
	err := formatter.Format(&errWriter{}, table)
	assert.Error(t, err)
}
