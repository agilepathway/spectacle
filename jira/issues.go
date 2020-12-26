package jira //nolint:stylecheck

import (
	"bytes"
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/getgauge/jira/util"
)

type Issues map[string]Issue //nolint:golint

func NewIssues() Issues { //nolint:golint
	return make(map[string]Issue)
}

func (i Issues) addSpecs(specFilenames []string) { //nolint:golint
	for _, filename := range specFilenames {
		i.addSpecToAllItsLinkedIssues(NewSpec(filename))
	}
}

func (i Issues) addSpecToAllItsLinkedIssues(spec Spec) {
	for _, issueKey := range spec.IssueKeys() {
		i.addSpecToIssue(spec, issueKey)
	}
}

func (i Issues) addSpecToIssue(spec Spec, issueKey string) {
	issue := i[issueKey]
	if issue.Key == "" {
		issue.Key = issueKey
	}

	issue.AddSpec(spec)
	i[issueKey] = issue
}

func (i Issues) publish() {
	jiraClient := jiraClient()

	for _, issue := range i {
		i.publishIssue(issue, jiraClient)
	}
}

func (i Issues) publishIssue(issue Issue, jiraClient *jira.Client) {
	req, err := jiraClient.NewRawRequest("PUT", fmt.Sprintf("rest/api/2/issue/%s", issue.Key), bytes.NewBufferString(`{"update":{"description":[{"set": "`+issue.jiraFmtSpecs()+`"}]}}`)) //nolint:lll
	util.Fatal("Error while creating Jira request %v", err)

	req.Header.Set("Content-type", "application/json")

	_, err = jiraClient.Do(req, nil)
	util.Fatal(fmt.Sprintf("Error while executing Jira request: %v", req), err)
}

func jiraClient() *jira.Client {
	base := os.Getenv("JIRA_BASE_URL")
	transport := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	jiraClient, err := jira.NewClient(transport.Client(), base)
	util.Fatal("Error while creating Jira Client", err)

	return jiraClient
}
