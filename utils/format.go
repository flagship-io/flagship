package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

func FormatItemTable[T any](columns []string, item T, w *tabwriter.Writer) {
	val := reflect.ValueOf(item)
	values := []string{}
	for _, column := range columns {
		values = append(values, val.FieldByName(column).String())
	}
	fmt.Fprintf(w, "%s\n", strings.Join(values, "\t"))
}

func FormatItem[T any](columns []string, item T, outputFormat string) {
	if outputFormat == "json" {
		projectJSON, err := json.Marshal(item)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Println(string(projectJSON))
		return
	}

	if outputFormat == "json-pretty" {
		projectJSON, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Println(string(projectJSON))
		return
	}

	if outputFormat == "table" {
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
		fmt.Fprintln(w, strings.Join(columns, "\t"))
		if reflect.TypeOf(item).Kind() == reflect.Slice {
			s := reflect.ValueOf(item)
			for i := 0; i < s.Len(); i++ {
				val := s.Index(i).Interface()
				FormatItemTable(columns, val, w)
			}

		} else {
			FormatItemTable(columns, item, w)
		}
		w.Flush()
		return
	}

	log.Fatalf("output format not handled: %s", outputFormat)
}
