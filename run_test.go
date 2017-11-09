package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// router.StaticFS("/", http.Dir("html"))
	router.Static("/html", "html")
	router.GET("/v1/api/getinfo", getapi)
	router.Run(":8009")
}

// getapi ...
func getapi(c *gin.Context) {
	log.Println(c.Request.URL)
	d, err := json.Marshal(query())
	CheckErr(err)
	c.String(http.StatusOK, string(d))
}

func query() []map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from case_info")
	CheckErr(err)
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	resarr := []map[string]string{}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		resarr = append(resarr, record)
	}
	db.Close()
	return resarr
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
