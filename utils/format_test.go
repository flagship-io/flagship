package utils

import (
	"bytes"
	"testing"
	"text/tabwriter"

	"github.com/stretchr/testify/assert"
)

type TestModel struct {
	Id   string
	Name string
	Type string
}

func TestFormatItemTable(t *testing.T) {
	item := TestModel{"1", "Test1", "string"}

	var b bytes.Buffer

	w := tabwriter.NewWriter(&b, 10, 1, 5, ' ', 0)
	columns := []string{"Id", "Name", "Type"}

	FormatItemTable(columns, item, w)
	w.Flush()

	assert.Contains(t, b.String(), "Test1")
	assert.Equal(t, "1         Test1     string\n", b.String())
}

func TestFormatItem(t *testing.T) {
	item := TestModel{"1", "Test1", "string"}
	items := []TestModel{{"1", "Test1", "string"}, {"2", "Test2", "string"}, {"1", "Test3", "string"}}

	var b_json bytes.Buffer
	var b_json_pretty bytes.Buffer
	var b_table_items bytes.Buffer
	var b_table_item bytes.Buffer

	columns := []string{"Id", "Name", "Type"}

	FormatItem(columns, items, "json", &b_json)

	assert.Contains(t, b_json.String(), "Test1")
	assert.Equal(t, "[{\"Id\":\"1\",\"Name\":\"Test1\",\"Type\":\"string\"},{\"Id\":\"2\",\"Name\":\"Test2\",\"Type\":\"string\"},{\"Id\":\"1\",\"Name\":\"Test3\",\"Type\":\"string\"}]\n", b_json.String())

	FormatItem(columns, items, "json-pretty", &b_json_pretty)

	assert.Contains(t, b_json_pretty.String(), "Test1")
	assert.Equal(t, "[\n  {\n    \"Id\": \"1\",\n    \"Name\": \"Test1\",\n    \"Type\": \"string\"\n  },\n  {\n    \"Id\": \"2\",\n    \"Name\": \"Test2\",\n    \"Type\": \"string\"\n  },\n  {\n    \"Id\": \"1\",\n    \"Name\": \"Test3\",\n    \"Type\": \"string\"\n  }\n]\n", b_json_pretty.String())

	FormatItem(columns, items, "table", &b_table_items)

	assert.Contains(t, b_table_items.String(), "Test1")
	assert.Equal(t, "Id        Name      Type\n1         Test1     string\n2         Test2     string\n1         Test3     string\n", b_table_items.String())

	FormatItem(columns, item, "table", &b_table_item)
	assert.Contains(t, b_table_item.String(), "Test1")
	assert.Equal(t, "Id        Name      Type\n1         Test1     string\n", b_table_item.String())

}
