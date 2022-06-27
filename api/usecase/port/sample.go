package port

import (
	"context"
	"hello-go-api/entity"
)

type SampleInputPort interface {
	GetSampleById(ctx context.Context, sampleId int)
}

type SampleOutputPort interface {
	Render(*entity.Sample)
	RenderError(error)
}

type SampleRepository interface {
	GetSampleById(ctx context.Context, sampleId int) (*entity.Sample, error)
}
