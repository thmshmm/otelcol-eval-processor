package evalprocessor

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type evalProcessor struct {
	logger *zap.Logger
	url    string
	client *http.Client
}

func newEvalProcessor(logger *zap.Logger, cfg *Config, client *http.Client) *evalProcessor {

	return &evalProcessor{
		logger: logger,
		url:    cfg.URL,
		client: client,
	}
}

func (p *evalProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	valid, err := p.evaluate(ld)
	if err != nil {
		p.logger.Warn("evaluation api returned an error", zapcore.Field{
			Key:       "error",
			Type:      zapcore.ErrorType,
			Interface: err,
		})

		ld.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().PutBool("evalError", true)
	} else {
		ld.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().PutBool("valid", valid)
	}

	return ld, nil
}

type response struct {
	Valid bool `json:"valid"`
}

func (p *evalProcessor) evaluate(ld plog.Logs) (bool, error) {
	resp, err := p.client.Get(p.url) // pass some actual data to the endpoint
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var responseObj response
	err = json.Unmarshal(body, &responseObj)
	if err != nil {
		return false, err
	}

	return responseObj.Valid, nil
}
