package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
	"text/tabwriter"
)

func FormatItemTable[T any](columns []string, item T, w *tabwriter.Writer) {
	val := reflect.ValueOf(item)
	values := []string{}
	for _, column := range columns {
		values = append(values, fmt.Sprintf("%v", val.FieldByName(column)))
	}
	fmt.Fprintf(w, "%s\n", strings.Join(values, "\t"))
}

func FormatItem[T any](columns []string, item T, outputFormat string, w io.Writer) {
	if outputFormat == "json" {
		projectJSON, err := json.Marshal(item)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}
		fmt.Fprintln(w, string(projectJSON))
		return
	}

	if outputFormat == "json-pretty" {
		projectJSON, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}
		fmt.Fprintln(w, string(projectJSON))
		return
	}

	if outputFormat == "table" {
		w_t := tabwriter.NewWriter(w, 10, 1, 5, ' ', 0)
		fmt.Fprintln(w_t, strings.Join(columns, "\t"))
		if reflect.TypeOf(item).Kind() == reflect.Slice {
			s := reflect.ValueOf(item)
			for i := 0; i < s.Len(); i++ {
				val := s.Index(i).Interface()
				FormatItemTable(columns, val, w_t)
			}

		} else {
			FormatItemTable(columns, item, w_t)
		}
		w_t.Flush()
		return
	}

	log.Fatalf("output format not handled: %s", outputFormat)
}
