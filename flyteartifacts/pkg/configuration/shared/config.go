package shared

import (
	"fmt"
	"github.com/flyteorg/flyte/flytestdlib/config"
)

const sharedServer = "sharedServer"

type Metrics struct {
	MetricsScope    string      `json:"metricsScope" pflag:",MetricsScope"`
	Port            config.Port `json:"port" pflag:",Profile port to start listen for pprof and metric handlers on."`
	ProfilerEnabled bool        `json:"profilerEnabled" pflag:",Enable Profiler on server"`
}

type SslOptions struct {
	CertificateAuthorityFile string `json:"certificateAuthorityFile"`
	CertificateFile          string `json:"certificateFile"`
	KeyFile                  string `json:"keyFile"`
}

type ServerSecurityOptions struct {
	Secure               bool       `json:"secure"`
	Ssl                  SslOptions `json:"ssl"`
	UseAuth              bool       `json:"useAuth"`
	AllowLocalhostAccess bool       `json:"allowLocalhostAccess" pflag:",Whether to permit localhost unauthenticated access to the server"`
}

type ServerConfiguration struct {
	Metrics                    Metrics               `json:"metrics" pflag:",Metrics configuration"`
	Port                       config.Port           `json:"port" pflag:",On which grpc port to serve"`
	HttpPort                   config.Port           `json:"httpPort" pflag:",On which http port to serve"`
	GrpcMaxResponseStatusBytes int32                 `json:"grpcMaxResponseStatusBytes" pflag:", specifies the maximum (uncompressed) size of header list that the client is prepared to accept on grpc calls"`
	GrpcServerReflection       bool                  `json:"grpcServerReflection" pflag:",Enable GRPC Server Reflection"`
	Security                   ServerSecurityOptions `json:"security"`
	MaxConcurrentStreams       int                   `json:"maxConcurrentStreams" pflag:",Limit on the number of concurrent streams to each ServerTransport."`
}

var sharedServerConfiguration = ServerConfiguration{
	Metrics: Metrics{
		MetricsScope:    "service:",
		Port:            config.Port{Port: 10254},
		ProfilerEnabled: false,
	},
	Port:                       config.Port{Port: 8089},
	HttpPort:                   config.Port{Port: 8088},
	GrpcMaxResponseStatusBytes: 320000,
	GrpcServerReflection:       false,
	Security: ServerSecurityOptions{
		Secure:  false,
		UseAuth: false,
	},
	MaxConcurrentStreams: 100,
}

func (s ServerConfiguration) GetGrpcHostAddress() string {
	return fmt.Sprintf(":%s", s.Port.String())
}

func (s ServerConfiguration) GetHttpHostAddress() string {
	return fmt.Sprintf(":%s", s.HttpPort.String())
}

var SharedServerConfig = config.MustRegisterSection(sharedServer, &sharedServerConfiguration)