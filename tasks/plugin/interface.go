package base

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/tracer"
)

type ServicePlugin interface {
	Process(task *models.Task, taskPlugin *models.TaskPlugin) *errors.Error
	InitTracerContext() *tracer.TracerContext
}
