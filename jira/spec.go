package jira //nolint:stylecheck

import (
	"fmt"
	"io/ioutil"

	"github.com/getgauge/jira/util"
)

type Spec struct { //nolint:golint
	Filename string
	markdown string
}

func NewSpec(filename string) Spec { //nolint:golint
	return Spec{filename, readMarkdown(filename)}
}

func (s *Spec) IssueKeys() []string { //nolint:golint
	// nolint:godox
	// TODO: implement this properly
	return []string{"JIRAGAUGE-1"}
}

func (s *Spec) jiraFmt() string {
	return mdToJira(s.markdown)
}

func readMarkdown(filename string) string {
	specBytes, err := ioutil.ReadFile(filename) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", filename), err)

	return string(specBytes)
}
