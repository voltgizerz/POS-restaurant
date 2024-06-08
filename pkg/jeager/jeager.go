package jeager

import (
	"io"

	jeagercfg "github.com/uber/jaeger-client-go/config"
)

func NewJeager(serviceName string) (io.Closer, error) {
	cfg := &jeagercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jeagercfg.SamplerConfig{
			Type:  "const",
			Param: 10,
		},
		Reporter: &jeagercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	closer, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return nil, err
	}

	return closer, nil
}
