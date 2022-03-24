package base

import (
	"elder-wand/errors"
	"elder-wand/models"
	"elder-wand/utils/tracer"
)

type ServicePlugin interface {
	Process(agreement *models.Agreement, service *models.AgreementDataService) *errors.Error
	InitTracerContext() *tracer.TracerContext
}
