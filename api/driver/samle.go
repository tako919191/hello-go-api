package driver

import (
	"database/sql"
	"fmt"
	"hello-go-api/adapter/controller"
	"hello-go-api/adapter/gateway"
	"hello-go-api/adapter/presenter"
	"hello-go-api/usecase/interactor"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve(addr string) {
	log.Println("driver:Serve: entry.")
	conn := connectDB()
	defer conn.Close()
	sample := controller.Sample{
		OutputFactory: presenter.NewSampleOutputPort,
		InputFactory:  interactor.NewSampleInputPort,
		RepoFactory:   gateway.NewSampleRepositoy,
		Conn:          conn,
	}
	router := gin.Default()
	router.GET("/sample/:id", sample.GetSampleById)
	router.Run(addr)
	// http.HandleFunc("/sample/", sample.GetSampleById)
	// err := http.ListenAndServe(addr, nil)
	// if err != nil {
	// 	log.Fatalf("listen and serve failed. %+v", err)
	// }
}

func connectDB() *sql.DB {
	log.Println("driver:connectDB: entry.")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("driver:connectDB: Connect DB Error: ", err)
	} else {
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}
	log.Println("driver:connectDB: Connected DB.")
	return db
}
