package interactor

import (
	"context"
	"hello-go-api/usecase/port"
)

type Sample struct {
	OutputPort port.SampleOutputPort
	SampleRepo port.SampleRepository
}

func NewSanpleInputPort(outputPort port.SampleOutputPort, sampleRepository port.SampleRepository) port.SampleInputPort {
	return &Sample{
		OutputPort: outputPort,
		SampleRepo: sampleRepository,
	}
}

func (s *Sample) GetSampleById(ctx context.Context, sampleId int) {
	sample, err := s.SampleRepo.GetSampleById(ctx, sampleId)
	if err != nil {
		s.OutputPort.RenderError(err)
		return
	}
	s.OutputPort.Render(sample)
}
