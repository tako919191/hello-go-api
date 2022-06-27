package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hello-go-api/entity"
	"hello-go-api/usecase/port"
	"log"
)

type SampleRepository struct {
	conn *sql.DB
}

func NewSampleRepositoy(conn *sql.DB) port.SampleRepository {
	return &SampleRepository{
		conn: conn,
	}
}

func (s *SampleRepository) GetSampleById(ctx context.Context, sampleId int) (*entity.Sample, error) {
	conn := s.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM sample_table WHERE id=?", sampleId)
	sample := entity.Sample{}
	err := row.Scan(&sample.Id, &sample.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sample not found. SampleID = %d", sampleId)
		}
		log.Println(err)
		return nil, errors.New("internal server error. adapter/gateway/GetSampleById")
	}
	return &sample, nil
}

func (s *SampleRepository) GetDBConn() *sql.DB {
	return s.conn
}
