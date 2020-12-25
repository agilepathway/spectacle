package export //nolint:stylecheck

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/getgauge/jira/util"
)

func Spec(specFile string, jiraIssue string) { //nolint:golint
	fmt.Printf("Exporting spec: %s", specFile)

	specBytes, err := ioutil.ReadFile(specFile) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", specFile), err)

	jsonSpec := jsonFmt(string(specBytes))

	base := os.Getenv("JIRA_BASE_URL")
	transport := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	jiraClient, err := jira.NewClient(transport.Client(), base)
	if err != nil {
		panic(err)
	}

	req, err := jiraClient.NewRawRequest("PUT", fmt.Sprintf("rest/api/2/issue/%s", jiraIssue), bytes.NewBufferString(`{"update":{"description":[{"set": "`+jsonSpec+`"}]}}`)) //nolint:lll
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-type", "application/json")

	_, err = jiraClient.Do(req, nil)
	if err != nil {
		panic(err)
	}
}

func jsonFmt(spec string) string {
	return removeOpeningAndClosingQuotes(fmt.Sprintf("%#v", spec))
}

func removeOpeningAndClosingQuotes(spec string) string {
	return spec[1 : len(spec)-1]
}
