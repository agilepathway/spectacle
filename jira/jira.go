package jira //nolint:stylecheck

func PublishSpecs(specFilenames []string) { //nolint:golint
	issues := NewIssues()
	issues.addSpecs(specFilenames)
	issues.publish()
}
