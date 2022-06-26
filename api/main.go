package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "tako:password@tcp(localhost:3306)/sample")
	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/dbtest", dbtest)
	router.Run(":3000")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message1": "pong",
		"message2": "pang",
	})
}

func dbtest(c *gin.Context) {
	sqlInsert()
	c.JSON(200, gin.H{"message": "success"})
}

func sqlInsert() {
	db, err := sql.Open("mysql", "tako:password@tcp(db:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
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
}
