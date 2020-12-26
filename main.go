package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/getgauge/jira/gauge_messages"
	"github.com/getgauge/jira/jira"
	"github.com/getgauge/jira/util"
	"google.golang.org/grpc"
)

const (
	gaugeSpecsDir = "GAUGE_SPEC_DIRS"
	fileSeparator = "||"
)

var projectRoot = util.GetProjectRoot() //nolint:gochecknoglobals

type handler struct {
	server *grpc.Server
}

func (h *handler) GenerateDocs(c context.Context, m *gauge_messages.SpecDetails) (*gauge_messages.Empty, error) {
	jiraIssues := make(map[string]jira.Issue)

	var files []string //nolint:prealloc
	for _, arg := range strings.Split(os.Getenv(gaugeSpecsDir), fileSeparator) {
		files = append(files, util.GetFiles(arg)...)
	}

	for _, file := range files {
		theSpec := jira.NewSpec(file)

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

	fmt.Println("Successfully exported specs to Jira")

	return &gauge_messages.Empty{}, nil
}

func (h *handler) Kill(c context.Context, m *gauge_messages.KillProcessRequest) (*gauge_messages.Empty, error) {
	defer h.stopServer()
	return &gauge_messages.Empty{}, nil
}

func (h *handler) stopServer() {
	h.server.Stop()
}

func main() {
	err := os.Chdir(projectRoot)
	if err != nil {
		util.Fatal("failed to start server.", err)
	}

	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		util.Fatal("failed to start server.", err)
	}

	l, err := net.ListenTCP("tcp", address)
	if err != nil {
		util.Fatal("failed to start server.", err)
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 10)) //nolint:gomnd
	h := &handler{server: server}
	gauge_messages.RegisterDocumenterServer(server, h)
	fmt.Printf("Listening on port:%d /n", l.Addr().(*net.TCPAddr).Port)
	server.Serve(l) //nolint:errcheck,gosec
}
