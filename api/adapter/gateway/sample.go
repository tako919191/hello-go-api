package gateway

import (
	"database/sql"
	"errors"
	"fmt"
	"hello-go-api/entity"
	"hello-go-api/usecase/port"
	"log"

	"github.com/gin-gonic/gin"
)

type SampleRepository struct {
	conn *sql.DB
}

func NewSampleRepositoy(conn *sql.DB) port.SampleRepository {
	log.Println("gateway:NewSampleRepositoy: entry.")
	return &SampleRepository{
		conn: conn,
	}
}

func (s *SampleRepository) GetSampleById(c *gin.Context, sampleId int) (*entity.Sample, error) {
	log.Println("gateway:GetSampleById: entry.")
	conn := s.GetDBConn()
	row := conn.QueryRowContext(c, "SELECT * FROM sample_table WHERE id=?", sampleId)
	sample := entity.Sample{}
	err := row.Scan(&sample.Id, &sample.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("gateway:GetSampleById: sample not found. SampleID = %d", sampleId)
		}
		log.Println(err)
		return nil, errors.New("internal server error. adapter/gateway/GetSampleById")
	} else {
		log.Printf("gateway:GetSampleById: GetSampleById: sample found. Id: %d, Name: %s", sample.Id, sample.Name)
	}
	return &sample, nil
}

func (s *SampleRepository) GetDBConn() *sql.DB {
	log.Println("gateway:GetDBConn: entry.")
	return s.conn
}
