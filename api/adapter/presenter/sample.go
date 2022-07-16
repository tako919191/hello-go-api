package presenter

import (
	"hello-go-api/entity"
	"hello-go-api/usecase/port"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	c *gin.Context
}

func NewSampleOutputPort(c *gin.Context) port.SampleOutputPort {
	log.Println("presenter:NewSampleOutputPort: entry.")
	return &Sample{
		c: c,
	}
}

func (s *Sample) Render(sample *entity.Sample) {
	log.Println("presenter:Render: entry.")
	log.Printf("presenter:Render: render sample. Id: %d, Name: %s", sample.Id, sample.Name)
	s.c.JSON(http.StatusOK, gin.H{
		"name": sample.Name,
	})
	// s.c.JSON(http.StatusOK, gin.H{
	// 	"name": "test",
	// })
	// s.w.WriteHeader(http.StatusOK)
	// fmt.Fprint(s.w, sample.Name)
}

func (s *Sample) RenderError(err error) {
	log.Println("presenter:RenderError: entry.")
	s.c.JSON(http.StatusInternalServerError, err)
	// s.w.WriteHeader(http.StatusInternalServerError)
	// fmt.Fprint(s.w, err)
}
