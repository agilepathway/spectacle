package jira

import (
	"fmt"
	"io/ioutil"

	"github.com/getgauge/jira/util"
)

type spec struct {
	filename string
	markdown string
}

func newSpec(filename string) spec {
	return spec{filename, readMarkdown(filename)}
}

func (s *spec) issueKeys() []string {
	return parseIssueKeys(s.markdown)
}

func (s *spec) jiraFmt() string {
	return mdToJira(s.markdown)
}

func readMarkdown(filename string) string {
	specBytes, err := ioutil.ReadFile(filename) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", filename), err)

	return string(specBytes)
}