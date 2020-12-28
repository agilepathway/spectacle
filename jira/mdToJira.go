package jira //nolint:stylecheck

import (
	"github.com/kalafut/m2j"
)

func mdToJira(str string) string {
	return m2j.MDToJira(str)
}
