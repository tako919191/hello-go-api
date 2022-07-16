package controller

import (
	"database/sql"
	"hello-go-api/usecase/port"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	OutputFactory func(c *gin.Context) port.SampleOutputPort
	InputFactory  func(o port.SampleOutputPort, s port.SampleRepository) port.SampleInputPort
	RepoFactory   func(c *sql.DB) port.SampleRepository
	Conn          *sql.DB
}

func (s *Sample) GetSampleById(c *gin.Context) {
	log.Printf("controller:GetSampleById: entry.[id:%s]", c.Param("id"))
	sampleId, _ := strconv.Atoi(c.Param("id"))
	outputPort := s.OutputFactory(c)
	repository := s.RepoFactory(s.Conn)
	inputPort := s.InputFactory(outputPort, repository)
	inputPort.GetSampleById(c, sampleId)
}
