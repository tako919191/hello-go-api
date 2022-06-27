package driver

import (
	"database/sql"
	"fmt"
	"hello-go-api/adapter/controller"
	"hello-go-api/adapter/gateway"
	"hello-go-api/adapter/presenter"
	"hello-go-api/usecase/interactor"
	"log"
	"net/http"
	"os"
)

func Serve(addr string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}
	sample := controller.Sample{
		OutputFactory: presenter.NewSampleOutputPort,
		InputFactory:  interactor.NewSanpleInputPort,
		RepoFactory:   gateway.NewSampleRepositoy,
		Conn:          conn,
	}
	http.HandleFunc("/sample/", sample.GetSampleById)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("listen and serve failed. %+v", err)
	}
}
