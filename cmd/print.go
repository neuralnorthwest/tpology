package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/neuralnorthwest/tpology/resource"
	"github.com/neuralnorthwest/tpology/table"
	"gopkg.in/yaml.v3"
)

type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatYAML  Format = "yaml"
)

// printEntities prints entities in various formats.
func printEntities(ents []interface{}, format Format) error {
	switch format {
	case FormatTable:
		return printTable(ents)
	case FormatJSON:
		return printJSON(ents)
	case FormatYAML:
		return printYAML(ents)
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
}

// printTable prints entities as a table.
func printTable(ents []interface{}) error {
	t := table.New()
	switch ents[0].(type) {
	case string:
		t.InsertColumn("Name", table.AtEnd)
		for _, ent := range ents {
			t.InsertRow([]interface{}{ent}, table.AtEnd)
		}
	case *resource.Resource:
		t.InsertColumn("Kind", table.AtEnd)
		t.InsertColumn("Name", table.AtEnd)
		t.InsertColumn("Description", table.AtEnd)
		t.InsertColumn("Owner", table.AtEnd)
		for _, ent := range ents {
			r := ent.(*resource.Resource)
			t.InsertRow([]interface{}{r.Kind, r.Name, r.Description, r.Owner}, table.AtEnd)
		}
	}
	return t.Write(os.Stdout, table.MarkdownFormatter())
}

// printJSON prints entities as JSON.
func printJSON(ents []interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(ents)
}

// printYAML prints entities as YAML.
func printYAML(ents []interface{}) error {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	return enc.Encode(ents)
}
