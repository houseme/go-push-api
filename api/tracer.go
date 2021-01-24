package api

import (
    `github.com/houseme/mipush/internal/builder`
    `github.com/houseme/mipush/internal/result`
)

// Tracer .
var Tracer = new(apiTracer)

type apiTracer struct {
}

// GetMessageStatus 通过RegId群推
func (a *apiTracer) GetMessageStatus(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    return nil, nil
}
