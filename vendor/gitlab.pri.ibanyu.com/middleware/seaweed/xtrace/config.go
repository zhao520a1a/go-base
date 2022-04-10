package xtrace

import (
	"fmt"
	"log"
	"os"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const (
	defaultAgentHost = "127.0.0.1"
	defaultAgentPort = "6831"
)

// Environment variables for jaeger agent
const (
	EnvJaegerAgentHost = "JAEGER_AGENT_HOST"
	EnvJaegerAgentPort = "JAEGER_AGENT_PORT"
)

// TracerConfig keeps metadata for tracer
type TracerConfig struct {
	Payload interface{}
}

// TracerConfigManager defines the interface for an concrete implementation of config manager.
type TracerConfigManager interface {
	GetConfig(serviceName string, tracerType TracerType) TracerConfig
}

func newTracerConfigManager() TracerConfigManager {
	return newSimpleTracerConfigManager()
}

type simpleManager struct{}

func newSimpleTracerConfigManager() *simpleManager {
	return &simpleManager{}
}

func (s *simpleManager) GetConfig(serviceName string, tracerType TracerType) TracerConfig {
	if tracerType != TracerTypeJaeger {
		// TODO: use xlog later
		log.Panicf("unknown tracer type %s for simpleManager", tracerType)
	}

	agentHost, agentPort := defaultAgentHost, defaultAgentPort

	if h, ok := os.LookupEnv(EnvJaegerAgentHost); ok {
		agentHost = h
	}

	if p, ok := os.LookupEnv(EnvJaegerAgentPort); ok {
		agentPort = p
	}

	return TracerConfig{
		Payload: config.Configuration{
			ServiceName: serviceName,
			Disabled:    false,
			RPCMetrics:  false,
			Sampler: &config.SamplerConfig{
				Type:  jaeger.SamplerTypeRateLimiting,
				Param: 1,
			},
			Reporter: &config.ReporterConfig{
				LocalAgentHostPort: fmt.Sprintf("%s:%s", agentHost, agentPort),
			},
			Headers: &jaeger.HeadersConfig{
				JaegerDebugHeader:        TraceDebugHeader,
				JaegerBaggageHeader:      TraceBaggageHeader,
				TraceContextHeaderName:   TraceContextHeaderName,
				TraceBaggageHeaderPrefix: TraceBaggageHeaderPrefix,
			},
		},
	}
}
