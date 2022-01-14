package gvar

import (
	"github.com/go-kit/kit/log"
	"go.opencensus.io/trace"
)

var Logger log.Logger
var Tracer trace.Tracer
