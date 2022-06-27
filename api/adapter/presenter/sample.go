package presenter

import (
	"fmt"
	"hello-go-api/entity"
	"hello-go-api/usecase/port"
	"net/http"
)

type Sample struct {
	w http.ResponseWriter
}

func NewSampleOutputPort(w http.ResponseWriter) port.SampleOutputPort {
	return &Sample{
		w: w,
	}
}

func (s *Sample) Render(sample *entity.Sample) {
	s.w.WriteHeader(http.StatusOK)
	fmt.Fprint(s.w, sample.Name)
}

func (s *Sample) RenderError(err error) {
	s.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(s.w, err)
}
