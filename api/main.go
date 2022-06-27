package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hello-go-api/driver"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Sample struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
	// router := gin.Default()
	// router.GET("/ping", ping)
	// router.GET("/insert", insert)
	// router.GET("/list", list)
	// router.Run(":3000")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message1": "pong",
		"message2": "pang",
	})
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "tako:password@tcp(db:3306)/sample")
	if err != nil {
		log.Fatal("Connect DB Error: ", err)
	} else {
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}
	log.Println("Connected DB.")
	return db
}

func insert(c *gin.Context) {
	db := connectDB()
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO sample_table (NAME) VALUES( ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	res, err := ins.Exec("api-sample")
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastInsertID)
	c.JSON(http.StatusOK, gin.H{"message": "success!"})
}

func list(c *gin.Context) {
	var samples []Sample
	db := connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * from sample_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		sample := Sample{}
		err = rows.Scan(&sample.Id, &sample.Name)
		if err != nil {
			panic(err)
		}
		samples = append(samples, sample)
	}

	res, err := json.Marshal(samples)
	if err != nil {
		panic(err)
	}
	log.Print("Samples: ", samples)
	log.Print("Res: ", res)
	c.JSON(http.StatusOK, &samples)
}
