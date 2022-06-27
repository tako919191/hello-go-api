package controller

import (
	"database/sql"
	"hello-go-api/usecase/port"
	"net/http"
	"strconv"
	"strings"
)

type Sample struct {
	OutputFactory func(w http.ResponseWriter) port.SampleOutputPort
	InputFactory  func(o port.SampleOutputPort, s port.SampleRepository) port.SampleInputPort
	RepoFactory   func(c *sql.DB) port.SampleRepository
	Conn          *sql.DB
}

func (s *Sample) GetSampleById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sampleId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/sample/"))
	outputPort := s.OutputFactory(w)
	repository := s.RepoFactory(s.Conn)
	inputPort := s.InputFactory(outputPort, repository)
	inputPort.GetSampleById(ctx, sampleId)
}
