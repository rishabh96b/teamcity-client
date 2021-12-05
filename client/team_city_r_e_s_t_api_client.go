// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/client/agent"
	"github.com/rishabh96b/teamcity-client/client/agent_pool"
	"github.com/rishabh96b/teamcity-client/client/audit"
	"github.com/rishabh96b/teamcity-client/client/build"
	"github.com/rishabh96b/teamcity-client/client/build_queue"
	"github.com/rishabh96b/teamcity-client/client/build_type"
	"github.com/rishabh96b/teamcity-client/client/change"
	"github.com/rishabh96b/teamcity-client/client/cloud_instance"
	"github.com/rishabh96b/teamcity-client/client/group"
	"github.com/rishabh96b/teamcity-client/client/investigation"
	"github.com/rishabh96b/teamcity-client/client/mute"
	"github.com/rishabh96b/teamcity-client/client/problem"
	"github.com/rishabh96b/teamcity-client/client/problem_occurrence"
	"github.com/rishabh96b/teamcity-client/client/project"
	"github.com/rishabh96b/teamcity-client/client/root"
	serverops "github.com/rishabh96b/teamcity-client/client/server"
	"github.com/rishabh96b/teamcity-client/client/test"
	"github.com/rishabh96b/teamcity-client/client/test_occurrence"
	"github.com/rishabh96b/teamcity-client/client/user"
	"github.com/rishabh96b/teamcity-client/client/vcs_root"
	"github.com/rishabh96b/teamcity-client/client/vcs_root_instance"
)

// Default team city r e s t API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "teamcity.mesosphere.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new team city r e s t API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TeamCityRESTAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new team city r e s t API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *TeamCityRESTAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new team city r e s t API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TeamCityRESTAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(TeamCityRESTAPI)
	cli.Transport = transport
	cli.Agent = agent.New(transport, formats)
	cli.AgentPool = agent_pool.New(transport, formats)
	cli.Audit = audit.New(transport, formats)
	cli.Build = build.New(transport, formats)
	cli.BuildQueue = build_queue.New(transport, formats)
	cli.BuildType = build_type.New(transport, formats)
	cli.Change = change.New(transport, formats)
	cli.CloudInstance = cloud_instance.New(transport, formats)
	cli.Group = group.New(transport, formats)
	cli.Investigation = investigation.New(transport, formats)
	cli.Mute = mute.New(transport, formats)
	cli.Problem = problem.New(transport, formats)
	cli.ProblemOccurrence = problem_occurrence.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Root = root.New(transport, formats)
	cli.Server = serverops.New(transport, formats)
	cli.Test = test.New(transport, formats)
	cli.TestOccurrence = test_occurrence.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.VcsRoot = vcs_root.New(transport, formats)
	cli.VcsRootInstance = vcs_root_instance.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// TeamCityRESTAPI is a client for team city r e s t API
type TeamCityRESTAPI struct {
	Agent agent.ClientService

	AgentPool agent_pool.ClientService

	Audit audit.ClientService

	Build build.ClientService

	BuildQueue build_queue.ClientService

	BuildType build_type.ClientService

	Change change.ClientService

	CloudInstance cloud_instance.ClientService

	Group group.ClientService

	Investigation investigation.ClientService

	Mute mute.ClientService

	Problem problem.ClientService

	ProblemOccurrence problem_occurrence.ClientService

	Project project.ClientService

	Root root.ClientService

	Server serverops.ClientService

	Test test.ClientService

	TestOccurrence test_occurrence.ClientService

	User user.ClientService

	VcsRoot vcs_root.ClientService

	VcsRootInstance vcs_root_instance.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TeamCityRESTAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Agent.SetTransport(transport)
	c.AgentPool.SetTransport(transport)
	c.Audit.SetTransport(transport)
	c.Build.SetTransport(transport)
	c.BuildQueue.SetTransport(transport)
	c.BuildType.SetTransport(transport)
	c.Change.SetTransport(transport)
	c.CloudInstance.SetTransport(transport)
	c.Group.SetTransport(transport)
	c.Investigation.SetTransport(transport)
	c.Mute.SetTransport(transport)
	c.Problem.SetTransport(transport)
	c.ProblemOccurrence.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Root.SetTransport(transport)
	c.Server.SetTransport(transport)
	c.Test.SetTransport(transport)
	c.TestOccurrence.SetTransport(transport)
	c.User.SetTransport(transport)
	c.VcsRoot.SetTransport(transport)
	c.VcsRootInstance.SetTransport(transport)
}
