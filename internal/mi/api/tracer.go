package api

import (
	"github.com/houseme/go-push-api/internal/mi"
	"github.com/houseme/go-push-api/internal/mi/builder"
)

// Tracer .
var Tracer = apiTracer{}

type apiTracer struct {
}

// GetMessageStatus 通过RegId群推
func (a *apiTracer) GetMessageStatus(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return nil, nil
}
