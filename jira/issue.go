package jira //nolint:stylecheck

import (
	"fmt"
)

type Issue struct { //nolint:golint
	specs []Spec
	Key   string
}

func (i *Issue) AddSpec(spec Spec) { //nolint:golint
	i.specs = append(i.specs, spec)
}

func (i *Issue) jiraFmtSpecs() string {
	// nolint: godox
	// TODO: do not just format the first spec
	// TODO: this method is currently doing two things: jira formatting and json formatting
	//   - so it should probably be split into separate methods
	return i.removeOpeningAndClosingQuotes(fmt.Sprintf("%#v", i.specs[0].jiraFmt()))
}

func (i *Issue) removeOpeningAndClosingQuotes(spec string) string {
	return spec[1 : len(spec)-1]
}
