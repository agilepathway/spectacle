package jira

import "github.com/getgauge/jira/internal/json"

type issue struct {
	specs []spec
	key   string
}

func (i *issue) addSpec(spec spec) {
	i.specs = append(i.specs, spec)
}

func (i *issue) publishSpecs() string {
	return json.Fmt(i.currentDescription() + i.jiraFmtSpecs())
}

func (i *issue) jiraFmtSpecs() string {
	var jiraFmtSpecs string
	for _, spec := range i.specs {
		jiraFmtSpecs += spec.jiraFmt()
	}

	return jiraFmtSpecs
}

func (i *issue) currentDescription() string {
	jiraClient := jiraClient()
	issue, _, _ := jiraClient.Issue.Get(i.key, nil)

	description := issue.Fields.Description
	if description == "" {
		return ""
	}

	return description + "\n"
}
