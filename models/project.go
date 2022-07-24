package models

import (
	"fmt"
	"io"
)

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	Items       []Project `json:"items"`
	CurrentPage int       `json:"current_page"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}

type ProjectToggleRequest struct {
	State string `json:"state"`
}

type ProjectItems []Project

type FormatableItem interface {
	FormatTable(w io.Writer)
}

func (projects ProjectItems) FormatTable(w io.Writer) {
	fmt.Fprintln(w, "Id\tName")
	for _, project := range projects {
		fmt.Fprintf(w, "%s\t%s\n", project.Id, project.Name)
	}
}
