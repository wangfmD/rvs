package report

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"github.com/wangfmD/rvs/handles/version"
	"github.com/wangfmD/rvs/setting"
	"github.com/wangfmD/rvs/sshv"
	"log"
	"net/http"
	"time"
)

type case_info_t struct {
	Id         string    `json:"ID"`
	Name       string    `json:"CASE_NAME"`
	ReportPath string    `json:"REPORT_PATH"`
	Type       int       `json:"TYPE"`
	Status     int       `json:"STATUS"`
	StartTime  time.Time `json:"START_TIME"`
	StopTime   time.Time `json:"STOP_TIME"`
	Duration   string    `json:"DURATION"`
	Mid        string    `json:"maddr"`
	Pid        string    `json:"paddr"`
}

func GetApi(c *gin.Context) {
	log.Println(c.Request.URL)
	data := make(map[string]interface{})
	data["caseInfots"] = query()
	d, err := json.Marshal(data)
	CheckErr(err)
	c.String(http.StatusOK, string(d))
}

func UpdateInfo(c *gin.Context) {
	var id string
	var status int
	var reportPath string
	var ra int64
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		var caseInfo case_info_t
		err := c.BindJSON(&caseInfo)
		id = caseInfo.Id
		status = caseInfo.Status
		reportPath = caseInfo.ReportPath
		db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
		CheckErr(err)
		sql := "update case_info set STOP_TIME=now(), STATUS=?, REPORT_PATH=?  where ID=?;"
		stmt, err := db.Prepare(sql)
		defer stmt.Close()
		CheckErr(err)
		rs, err := stmt.Exec(status, reportPath, id)
		CheckErr(err)
		ra, err = rs.RowsAffected()
		CheckErr(err)
		db.Close()
	}
	msg := fmt.Sprintf("Update person %d successful %d", id, ra)
	// msg := "Update person" + id + "successful" + ra

	c.JSON(http.StatusOK, gin.H{
		"status": "suc",
		"result": map[string]string{
			"id":  id,
			"msg": msg,
		},
	})

}

//根据ip匹配cfg的配置的用户名，密码
func getServerPl(addr string) (*setting.ServerOs, error) {
	p := setting.GetPlatformAddrs()
	for _, v := range p {
		if v.Addr == addr {
			return v, nil
		}
	}
	return nil, errors.New("error address")
}

//根据ip匹配cfg的配置的用户名，密码
func getServerMe(addr string) (*setting.ServerOs, error) {
	p := setting.GetMediaAddrs()
	for _, v := range p {
		if v.Addr == addr {
			return v, nil
		}
	}
	return nil, errors.New("error address")
}

//执行用例时ssh查询版本并入库
func sSHSqlAddPlatformvers(id, addr string) {
	server, err := getServerPl(addr)

	if err != nil {
		log.Println(err)
	} else {
		versMap, err := sshv.GetVersionMaps(server.Name, server.Pwd, addr)
		if err != nil {
			log.Println(err)
		} else {
			version.SqlAddVersionByExecCase(id, addr, versMap)
		}
	}
	log.Println("exec case: insert platform version to db")

}

//执行用例时ssh查询版本并入库
func sSHSqlAddMediavers(id, addr string) {
	server, err := getServerMe(addr)
	if err != nil {
		log.Println(err)
	} else {
		versMap, err := sshv.GetVersionMaps(server.Name, server.Pwd, addr)
		if err != nil {
			log.Println(err)
		} else {
			version.SqlAddMediaVersionByExecCase(id, addr, versMap)
		}
	}
	log.Println("exec case: insert media version to db")
}

//执行用例时插入执行信息
func AddCaseInfo(c *gin.Context) {
	var id, reportId, mid, pid string
	var caseName string
	// var reportPath string
	var caseType int

	uuid1 := uuid.NewV4()
	id = uuid1.String()

	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		var caseinfo case_info_t
		err := c.BindJSON(&caseinfo)
		// err := c.MustBindWith(&ci, binding.JSON)
		// err := binding.JSON.Bind(c.Request, &ci)
		// id = caseinfo.Id
		// reportPath = caseinfo.ReportPath
		caseType = caseinfo.Type
		caseName = caseinfo.Name
		log.Println("caseType:", caseType)
		log.Println("casename:", caseName)
		pid = caseinfo.Pid
		mid = caseinfo.Mid
		log.Println(id)
		log.Println(caseName)
		// log.Println(reportPath)
		log.Println(caseType)
		CheckErr(err)
	}

	uuid2 := uuid.NewV4()
	reportId = uuid2.String()
	log.Println("pid: ", pid)
	if pid != "" {
		go sSHSqlAddPlatformvers(reportId, pid)
	}

	log.Println("mid: ", mid)
	if mid != "" {
		go sSHSqlAddMediavers(reportId, mid)
	}

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	sql := "INSERT INTO `case_info` (`ID`, `CASE_NAME`, `REPORT_PATH`,`EXEC_VERSION`, `TYPE`, `STATUS`, `START_TIME`, `STOP_TIME`) VALUES(?,?,'/opt',?,?,0,now(),now())"
	rs, err := db.Exec(sql, id, caseName, reportId, caseType)
	CheckErr(err)
	i, err := rs.LastInsertId()
	CheckErr(err)
	log.Println(i)
	err = db.Close()
	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"status": "suc",
		"msg":    map[string]string{"id": id},
	})
}

func GetCaseInfo(c *gin.Context) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from case_info order by `START_TIME` DESC")
	CheckErr(err)
	caseInfots := make([]case_info_t, 0)
	for rows.Next() {
		var caseInfot case_info_t
		rows.Scan(&caseInfot.Id, &caseInfot.Name, &caseInfot.ReportPath, &caseInfot.Type, &caseInfot.Status, &caseInfot.StartTime, &caseInfot.StopTime, &caseInfot.Duration)
		caseInfots = append(caseInfots, caseInfot)
	}
	db.Close()
	log.Println(caseInfots)
	c.JSON(http.StatusOK, gin.H{"caseInfots": caseInfots})

}

func query() []map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from case_info order by `START_TIME` DESC")
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
		log.Println(err)
	}
}
