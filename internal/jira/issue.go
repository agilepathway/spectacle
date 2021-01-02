package jira

import (
	"fmt"
)

type issue struct {
	specs []spec
	key   string
}

func (i *issue) addSpec(spec spec) {
	i.specs = append(i.specs, spec)
}

func (i *issue) jiraFmtSpecs() string {
	var jiraFmtSpecs string
	for _, spec := range i.specs {
		jiraFmtSpecs += i.removeOpeningAndClosingQuotes(fmt.Sprintf("%#v", spec.jiraFmt()))
	}

	return jiraFmtSpecs
}

func (i *issue) removeOpeningAndClosingQuotes(spec string) string {
	return spec[1 : len(spec)-1]
}
