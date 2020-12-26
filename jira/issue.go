package jira //nolint:stylecheck

import (
	"bytes"
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

type Issue struct { //nolint:golint
	specs []Spec
	Key   string
}

func (i *Issue) AddSpec(spec Spec) { //nolint:golint
	i.specs = append(i.specs, spec)
}

func (i *Issue) Publish() { //nolint:golint
	base := os.Getenv("JIRA_BASE_URL")
	transport := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	jiraClient, err := jira.NewClient(transport.Client(), base)
	if err != nil {
		panic(err)
	}

	req, err := jiraClient.NewRawRequest("PUT", fmt.Sprintf("rest/api/2/issue/%s", i.Key), bytes.NewBufferString(`{"update":{"description":[{"set": "`+i.jiraFmtSpecs()+`"}]}}`)) //nolint:lll
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-type", "application/json")

	_, err = jiraClient.Do(req, nil)
	if err != nil {
		panic(err)
	}
}

func (i *Issue) jiraFmtSpecs() string {
	// nolint: godox
	// TODO: do not just format the first spec
	return i.removeOpeningAndClosingQuotes(fmt.Sprintf("%#v", i.specs[0].String()))
}

func (i *Issue) removeOpeningAndClosingQuotes(spec string) string {
	return spec[1 : len(spec)-1]
}
