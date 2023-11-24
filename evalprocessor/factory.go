package evalprocessor

import (
	"context"
	"net/http"
	"time"

	"github.com/thmshmm/otelcol-eval-processor/internal/metadata"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

func NewFactory() processor.Factory {
	return processor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processor.WithLogs(createLogsProcessor, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		URL: "http://localhost:8080",
	}
}

func createLogsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	processorCfg := cfg.(*Config)

	client := &http.Client{
		Timeout: time.Duration(processorCfg.TimeoutSeconds) * time.Second,
	}

	evalProcessor := newEvalProcessor(set.Logger, processorCfg, client)

	return processorhelper.NewLogsProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		evalProcessor.processLogs,
		processorhelper.WithCapabilities(processorCapabilities),
	)
}
