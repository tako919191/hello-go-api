package interactor

import (
	"hello-go-api/usecase/port"
	"log"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	OutputPort port.SampleOutputPort
	SampleRepo port.SampleRepository
}

func NewSampleInputPort(o port.SampleOutputPort, r port.SampleRepository) port.SampleInputPort {
	log.Println("interactor:NewSampleInputPort: entry.")
	return &Sample{
		OutputPort: o,
		SampleRepo: r,
	}
}

func (s *Sample) GetSampleById(c *gin.Context, sampleId int) {
	log.Println("interactor:GetSampleById: entry.")
	sample, err := s.SampleRepo.GetSampleById(c, sampleId)
	if err != nil {
		s.OutputPort.RenderError(err)
		return
	} else {
		log.Printf("interactor:GetSampleById: sample found. Id: %d, Name: %s", sample.Id, sample.Name)
	}
	s.OutputPort.Render(sample)
}
