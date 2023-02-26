package table

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_New tests the New function.
func Test_New(t *testing.T) {
	t.Parallel()
	table := New()
	assert.NotNil(t, table)
}

// Test_Table_ColumnNames tests the ColumnNames function.
func Test_Table_ColumnNames(t *testing.T) {
	t.Parallel()
	table := New()
	assert.Equal(t, []string{}, table.ColumnNames())
	table.InsertColumn("foo", AtEnd)
	assert.Equal(t, []string{"foo"}, table.ColumnNames())
	table.InsertColumn("bar", AtBeginning)
	assert.Equal(t, []string{"bar", "foo"}, table.ColumnNames())
	table.InsertColumn("baz", BeforeColumn("foo"))
	assert.Equal(t, []string{"bar", "baz", "foo"}, table.ColumnNames())
	table.InsertColumn("qux", AfterColumn("bar"))
	assert.Equal(t, []string{"bar", "qux", "baz", "foo"}, table.ColumnNames())
}

// Test_Table_Write tests the Write function.
func Test_Table_Write(t *testing.T) {
	t.Parallel()
	table := New()
	table.InsertColumn("foo", AtEnd)
	table.InsertColumn("bar", AtEnd)
	table.InsertColumn("baz", AtEnd)
	table.InsertColumn("qux", AtEnd)
	table.InsertColumn("quux", AtEnd)
	table.InsertColumn("corge", AtEnd)
	table.InsertColumn("grault", AtEnd)
	table.InsertColumn("garply", AtEnd)
	table.InsertColumn("waldo", AtEnd)
	table.InsertColumn("fred", AtEnd)
	table.InsertColumn("plugh", AtEnd)
	table.InsertColumn("xyzzy", AtEnd)
	table.InsertColumn("thud", AtEnd)
	table.InsertColumn("foo", AtEnd)
	table.InsertColumn("bar", AtEnd)
	table.InsertColumn("baz", AtEnd)
	table.InsertColumn("qux", AtEnd)
	table.InsertColumn("quux", AtEnd)
	table.InsertColumn("corge", AtEnd)
	table.InsertColumn("grault", AtEnd)
	table.InsertColumn("garply", AtEnd)
	table.InsertColumn("waldo", AtEnd)
	table.InsertColumn("fred", AtEnd)
	table.InsertColumn("plugh", AtEnd)
	table.InsertColumn("xyzzy", AtEnd)
	table.InsertColumn("thud", AtEnd)

	table.InsertRow([]interface{}{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud", "foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud"}, AtEnd)
	table.InsertRow([]interface{}{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud", "foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud"}, AtEnd)
	table.InsertRow([]interface{}{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud", "foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud"}, AtEnd)
	table.InsertRow([]interface{}{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud", "foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy", "thud"}, AtEnd)

	buf := bytes.Buffer{}
	err := table.Write(&buf, MarkdownFormatter())
	assert.NoError(t, err)
	assert.Equal(t, `| foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud | foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud |
| --- | --- | --- | --- | ---- | ----- | ------ | ------ | ----- | ---- | ----- | ----- | ---- | --- | --- | --- | --- | ---- | ----- | ------ | ------ | ----- | ---- | ----- | ----- | ---- |
| foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud | foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud |
| foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud | foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud |
| foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud | foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud |
| foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud | foo | bar | baz | qux | quux | corge | grault | garply | waldo | fred | plugh | xyzzy | thud |
`, buf.String())
}
