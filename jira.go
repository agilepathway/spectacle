package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/getgauge/jira/export"
	"github.com/getgauge/jira/gauge_messages"
	"github.com/getgauge/jira/spec"
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
	jiraIssues := make(map[string][]spec.Spec)

	var files []string //nolint:prealloc
	for _, arg := range strings.Split(os.Getenv(gaugeSpecsDir), fileSeparator) {
		files = append(files, util.GetFiles(arg)...)
	}

	for _, file := range files {
		theSpec := spec.New(file)

		for _, issue := range theSpec.JiraIssues() {
			issues := jiraIssues[issue]
			jiraIssues[issue] = append(issues, theSpec)
		}
	}

	for jiraIssue, issueSpecs := range jiraIssues {
		// nolint:godox
		// TODO: export all the specs pertaining to the issue, not just the first one
		export.Spec(issueSpecs[0].Filename, jiraIssue)
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
