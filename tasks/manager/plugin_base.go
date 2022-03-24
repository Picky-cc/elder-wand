package manager

import (
	"elder-wand/utils/tracer"
)

type BasePluginBase struct {
	Tracer *tracer.TracerContext
}

func (c *BasePluginBase) InitTracerContext() *tracer.TracerContext {
	c.Tracer = &tracer.TracerContext{}
	return c.Tracer
}
