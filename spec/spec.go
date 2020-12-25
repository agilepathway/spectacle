package spec //nolint:stylecheck

type Spec struct { //nolint:golint
	Filename string
}

func New(filename string) Spec { //nolint:golint
	return Spec{filename}
}

func (s *Spec) JiraIssues() []string { //nolint:golint
	return []string{"JIRAGAUGE-1"}
}
