package port

import (
	"hello-go-api/entity"

	"github.com/gin-gonic/gin"
)

type SampleInputPort interface {
	GetSampleById(c *gin.Context, sampleId int)
}

type SampleOutputPort interface {
	Render(*entity.Sample)
	RenderError(error)
}

type SampleRepository interface {
	GetSampleById(c *gin.Context, sampleId int) (*entity.Sample, error)
}
