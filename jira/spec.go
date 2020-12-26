package jira //nolint:stylecheck

import (
	"fmt"
	"io/ioutil"

	"github.com/getgauge/jira/util"
)

type Spec struct { //nolint:golint
	Filename string
}

func NewSpec(filename string) Spec { //nolint:golint
	return Spec{filename}
}

func (s *Spec) IssueKeys() []string { //nolint:golint
	// nolint:godox
	// TODO: implement this properly
	return []string{"JIRAGAUGE-1"}
}

func (s Spec) String() string {
	specBytes, err := ioutil.ReadFile(s.Filename) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", s.Filename), err)

	return string(specBytes)
}
