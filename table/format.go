package table

import (
	"fmt"
	"io"
	"strings"
)

// formatterBase is a base formatter interface.
type formatterBase interface {
	// FormatHeader formats the table header.
	FormatHeader(io.Writer, *Table)
	// FormatRow formats a table row.
	FormatRow(io.Writer, *Table, *Row)
	// FormatFooter formats the table footer.
	FormatFooter(io.Writer, *Table)
}

// Formatter is a formatter for a table.
type Formatter struct {
	formatterBase
}

// FormatOption is a format option.
type FormatOption func(*Formatter)

// NoHeader is a format option to disable the header.
var NoHeader = func(f *Formatter) {
	f.formatterBase = &noHeaderFormatter{f.formatterBase}
}

// NoFooter is a format option to disable the footer.
var NoFooter = func(f *Formatter) {
	f.formatterBase = &noFooterFormatter{f.formatterBase}
}

// noHeaderFormatter is a formatter without a header.
type noHeaderFormatter struct {
	formatterBase
}

// FormatHeader formats the table header.
func (f *noHeaderFormatter) FormatHeader(w io.Writer, t *Table) {}

// noFooterFormatter is a formatter without a footer.
type noFooterFormatter struct {
	formatterBase
}

// FormatFooter formats the table footer.
func (f *noFooterFormatter) FormatFooter(w io.Writer, t *Table) {}

// Format formats the table.
func (f *Formatter) Format(w io.Writer, t *Table) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("table: %v", r)
		}
	}()

	t.UpdateWidths()
	f.FormatHeader(w, t)
	for _, r := range t.Rows {
		f.FormatRow(w, t, r)
	}
	f.FormatFooter(w, t)
	return nil
}

// markdownFormatter is a Markdown formatter.
type markdownFormatter struct{}

// MarkdownFormatter returns a new Markdown formatter.
func MarkdownFormatter(opts ...FormatOption) *Formatter {
	f := &Formatter{formatterBase: &markdownFormatter{}}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

// FormatHeader formats the table header.
func (f *markdownFormatter) FormatHeader(w io.Writer, t *Table) {
	MustFprintf(w, "|")
	for _, c := range t.Columns {
		MustFprintf(w, " %-*s |", c.Width, c.Name)
	}
	MustFprintf(w, "\n")
	MustFprintf(w, "|")
	for _, c := range t.Columns {
		MustFprintf(w, "-%-*s-|", c.Width, strings.Repeat("-", c.Width))
	}
	MustFprintf(w, "\n")
}

// FormatRow formats a table row.
func (f *markdownFormatter) FormatRow(w io.Writer, t *Table, r *Row) {
	MustFprintf(w, "|")
	for i, c := range r.Cells {
		MustFprintf(w, " %-*s |", t.Columns[i].Width, c.Value)
	}
	MustFprintf(w, "\n")
}

// FormatFooter formats the table footer.
func (f *markdownFormatter) FormatFooter(w io.Writer, t *Table) {}

// csvFormatter is a CSV formatter.
type csvFormatter struct{}

// CSVFormatter returns a new CSV formatter.
func CSVFormatter(opts ...FormatOption) *Formatter {
	f := &Formatter{formatterBase: &csvFormatter{}}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

// FormatHeader formats the table header.
func (f *csvFormatter) FormatHeader(w io.Writer, t *Table) {
	for i, c := range t.Columns {
		if i > 0 {
			MustFprintf(w, ",")
		}
		MustFprintf(w, "%s", c.Name)
	}
	MustFprintf(w, "\n")
}

// FormatRow formats a table row.
func (f *csvFormatter) FormatRow(w io.Writer, t *Table, r *Row) {
	for i, c := range r.Cells {
		if i > 0 {
			MustFprintf(w, ",")
		}
		MustFprintf(w, "%s", c.Value)
	}
	MustFprintf(w, "\n")
}

// FormatFooter formats the table footer.
func (f *csvFormatter) FormatFooter(w io.Writer, t *Table) {}

// MustFprintf is a helper function to call MustFprintf and panic if an error occurs.
func MustFprintf(w io.Writer, format string, a ...interface{}) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		panic(err)
	}
}
