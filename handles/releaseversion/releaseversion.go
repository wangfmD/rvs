package releaseversion

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/wangfmD/rvs/log"
	"github.com/wangfmD/rvs/setting"
	"log"
	"net/http"
)

// HandleUpdateTagvers update cfg tag version to mysql
func HandleUpdateTagvers(c *gin.Context) {
	err := SqlAddTagvers()
	if err == nil {
		log.Println("=====update tags end================")
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"result": "suc",
		})
	} else {
		log.Println("=====update tags end================")
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"result": err.Error(),
		})
	}
}

// HandleQueryDiffVers returns versions array by tag and ip address
func HandleQueryDiffVers(c *gin.Context) {
	tag := c.PostForm("tag")
	addr := c.PostForm("addr")
	releaseVers := queryTagVers(tag)
	actualVers := queryVersByIps(addr)
	contentArrs := make([]map[string]string, 0)
	contentArrs = append(contentArrs, releaseVers)
	contentArrs = append(contentArrs, actualVers)
	log.Println(tag, addr)
	log.Println(contentArrs)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": contentArrs,
	})
}

// HandleQueryTags 返回版本tags，example:[V4.021 V4.020]
func HandleQueryTags(c *gin.Context) {
	p := setting.GetTagVersions()
	tags := make([]string, 0, 0)
	for k, _ := range p {
		tags = append(tags, k)
	}

	log.Println(tags)
	c.JSON(http.StatusOK, gin.H{
		"tags": tags,
	})
}

// HandleQueryVersByIp returns versions by ip address
func HandleQueryVersByIp(c *gin.Context) {
	addr := c.Param("addr")
	vers := queryVersByIps(addr)
	c.JSON(http.StatusOK, gin.H{"status": "success",
		"result": vers,
	})
}

// HandleQueryTagVers returns versions by tags
func HandleQueryTagVers(c *gin.Context) {
	vertag := c.Param("vertag")
	vers := queryTagVers(vertag)
	c.JSON(http.StatusOK, gin.H{"status": "success",
		"result": vers,
	})
}

func queryVersByIps(vid string) map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	checkErr(err)
	sql := `SELECT
			  agent,
			  interact,
			  interactbusiness,
			  jycenter,
			  jyresource,
			  middlecas,
			  middlecenter,
			  middlecenterfile,
			  middlecenterres,
			  middleclient,
			  middledriver,
			  middleresource,
			  middlewaremcu,
			  middletherepair,
			  mysql,
			  nginx,
			  openfire,
			  redis,
			  middledatabase
		  FROM server_version
		  WHERE id="last" and serveraddr=?`
	rows, err := db.Query(sql, vid)

	checkErr(err)
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// resMap := make(map[string]map[string]string)

	record := make(map[string]string)

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			} else {
				record[columns[i]] = ""
			}

		}
		record["versiontag"] = vid
		// resMap[vid] = record
	}
	db.Close()
	// return resMap
	log.Println(record)
	return record
}

// queryTagvers ...
func queryTagVers(vid string) map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	checkErr(err)
	sql := `SELECT
			  agent,
			  interact,
			  interactbusiness,
			  jycenter,
			  jyresource,
			  middlecas,
			  middlecenter,
			  middlecenterfile,
			  middlecenterres,
			  middleclient,
			  middledriver,
			  middleresource,
			  middlewaremcu,
			  middletherepair,
			  mysql,
			  nginx,
			  openfire,
			  redis,
			  middledatabase
		  FROM version_tag
		  WHERE versionid=?`
	rows, err := db.Query(sql, vid)

	checkErr(err)
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// resMap := make(map[string]map[string]string)

	record := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		record["versiontag"] = vid
	}
	db.Close()
	// return resMap
	return record
}

// SqlAddversion TODO: NEEDS COMMENT INFO
func SqlAddTagvers() error {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	versionMaps := setting.GetTagVersions()

	for version, v := range versionMaps {
		var total int
		sqlQuery := "select count(*) from version_tag where versionid=?"
		err := db.QueryRow(sqlQuery, version).Scan(&total)
		checkErr(err)
		if total != 0 {
			//update
			sql := `
		  UPDATE version_tag
				SET agent=?,
				interact=?,
				interactbusiness=?,
				jycenter=?,
				jyresource=?,
				middlecas=?,
				middlecenter=?,
				middlecenterfile=?,
				middlecenterres=?,
				middleclient=?,
				middledriver=?,
				middleresource=?,
				middlewaremcu=?,
				middletherepair=?,
				mysql=?,
				nginx=?,
				openfire=?,
				redis=?,
				middledatabase=?,
				filesrv=?,
				ftp=?,
				mbs=?,
				update_time=now()
		  WHERE versionid=?`
			stmt, err := db.Prepare(sql)
			checkErr(err)
			res, err := stmt.Exec(
				v["agent"],
				v["interact"],
				v["interactbusiness"],
				v["jycenter"],
				v["jyresource"],
				v["middlecas"],
				v["middlecenter"],
				v["middlecenterfile"],
				v["middlecenterres"],
				v["middleclient"],
				v["middledriver"],
				v["middleresource"],
				v["middlewaremcu"],
				v["middletherepair"],
				v["mysql"],
				v["nginx"],
				v["openfire"],
				v["redis"],
				v["middledatabase"],
				v["filesrv"],
				v["ftp"],
				v["mbs"],
				version)
			checkErr(err)
			num, err := res.RowsAffected()
			checkErr(err)
			log.Println("update ", version, "succ,sums:", num)
		} else {
			sql := `
	INSERT INTO version_tag
		  (versionid,
		  agent,
		  interact,
		  interactbusiness,
		  jycenter,
		  jyresource,
		  middlecas,
		  middlecenter,
		  middlecenterfile,
		  middlecenterres,
		  middleclient,
		  middledriver,
		  middleresource,
		  middlewaremcu,
		  middletherepair,
		  mysql,
		  nginx,
		  openfire,
		  redis,
		  middledatabase,
		  filesrv,
		  ftp,
		  mbs,
		  create_time)
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,now());`
			rs, err := db.Exec(sql, version,
				v["agent"],
				v["interact"],
				v["interactbusiness"],
				v["jycenter"],
				v["jyresource"],
				v["middlecas"],
				v["middlecenter"],
				v["middlecenterfile"],
				v["middlecenterres"],
				v["middleclient"],
				v["middledriver"],
				v["middleresource"],
				v["middlewaremcu"],
				v["middletherepair"],
				v["mysql"],
				v["nginx"],
				v["openfire"],
				v["redis"],
				v["middledatabase"],
				v["filesrv"],
				v["ftp"],
				v["mbs"])
			if err != nil {
				log.Fatal(err)
			}
			i, err := rs.LastInsertId()
			log.Println("dd", i)
			checkErr(err)
		}
	}
	log.Println("update tags succ")
	return nil
}

// checkerr ...
func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
