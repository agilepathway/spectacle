package jira //nolint:stylecheck

func PublishSpecs(specFilenames []string) { //nolint:golint
	jiraIssues := make(map[string]Issue)

	for _, filename := range specFilenames {
		theSpec := NewSpec(filename)

		for _, issueKey := range theSpec.IssueKeys() {
			issue := jiraIssues[issueKey]
			// nolint:godox
			// TODO: not elegant to always set the key here
			issue.Key = issueKey
			issue.AddSpec(theSpec)
			jiraIssues[issueKey] = issue
		}
	}

	for _, issue := range jiraIssues {
		issue.Publish()
	}
}
