package version

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/wangfmD/rvs/log"
	"github.com/wangfmD/rvs/setting"
	"github.com/wangfmD/rvs/sshv"
	"log"
	"net/http"
)

func QueryVersionById(c *gin.Context) {
	id := c.Param("versionid")
	pVersMap := sqlQueryVersionById(id)
	mVersMap := sqlQueryMediaVersionById(id)
	vers := make(map[string]map[string]string)
	vers["platform"] = pVersMap
	vers["media"] = mVersMap
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": vers,
	})

}

// QueryServerAddrs ...
func QueryServerAddrs(c *gin.Context) {
	p := setting.GetPlatformAddrs()
	addrs := make([]string, 0, 0)
	for _, v := range p {
		addrs = append(addrs, v.Addr)
	}
	log.Println(addrs)
	c.JSON(http.StatusOK, gin.H{
		"addrs": addrs,
	})
}

func AddVersion(c *gin.Context) {
	addr := c.Param("addr")
	// 校验 addr todo
	var m map[string]string
	var err error

	f := func() (*setting.ServerOs, error) {
		p := setting.GetPlatformAddrs()
		for _, v := range p {
			if v.Addr == addr {
				return v, nil
			}
		}
		return nil, errors.New("error address")
	}

	server, err := f()
	if err != nil {
		m = make(map[string]string)
		err = errors.New("查询服务器地址无效")
		m["msg"] = "查询服务器地址无效"

	} else {
		m, err = sshv.GetVersionMaps(server.Name, server.Pwd, addr)
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    m,
			"err":    err.Error(),
		})

	} else {
		err = sqlAddversion(addr, m)
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"msg":    m,
		})
	}
}

func QueryVersion(c *gin.Context) {
	res := sqlQueryVersion()
	c.JSON(http.StatusOK, gin.H{
		"versions": res,
	})
}

func sqlQueryVersion() []map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	sql := `select id,
      serveraddr,
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
	  mysql,
	  nginx,
	  openfire,
	  redis,
	  update_time from server_version where id='last' and type='mgr'`

	rows, err := db.Query(sql)
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

// 查询测试服务器时插入server_version表的版本信息
func sqlAddversion(addr string, m map[string]string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return err
	}
	// delete from server_version where id='last';
	sql_del := `delete from server_version where id='last' and serveraddr=?`
	_, err = db.Exec(sql_del, addr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	sql := `
   INSERT INTO server_version
	  (id,
      serveraddr,
	  type,
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
	  middledatabase,
	  mysql,
	  nginx,
	  openfire,
	  redis,
	  update_time)
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,now());`

	rs, err := db.Exec(sql, "last", addr, "mgr", m["agent"], m["interact"], m["interactbusiness"], m["jycenter"], m["jyresource"], m["middlecas"], m["middlecenter"], m["middlecenterfile"], m["middlecenterres"], m["middleclient"], m["middledriver"], m["middleresource"], m["middleware-mcu"], m["middledatabase"], m["mysql"], m["nginx"], m["openfire"], m["redis"])
	if err != nil {
		log.Fatal(err)
		return err
	}

	i, err := rs.LastInsertId()
	log.Println(i)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// 执行用例时插入server_version表的版本信息
func SqlAddVersionByExecCase(id string, addr string, m map[string]string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return err
	}

	sql := `
   INSERT INTO server_version
	  (id,
      serveraddr,
	  type,
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
	  middledatabase,
	  mysql,
	  nginx,
	  openfire,
	  redis,
	  update_time)
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,now());`

	rs, err := db.Exec(sql, id, addr, "mgr", m["agent"], m["interact"], m["interactbusiness"], m["jycenter"], m["jyresource"], m["middlecas"], m["middlecenter"], m["middlecenterfile"], m["middlecenterres"], m["middleclient"], m["middledriver"], m["middleresource"], m["middleware-mcu"], m["middledatabase"], m["mysql"], m["nginx"], m["openfire"], m["redis"])
	if err != nil {
		log.Fatal(err)
		return err
	}

	i, err := rs.LastInsertId()
	log.Println(i)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// uuid1 := uuid.NewV4()
// id = uuid1.String()

func sqlQueryVersionById(id string) map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	sql := `select id,
      serveraddr,
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
	  mysql,
	  nginx,
	  openfire,
	  redis,
	  update_time from server_version where id=? and type='mgr'`

	rows, err := db.Query(sql, id)
	CheckErr(err)
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
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
	}
	db.Close()
	return record
}

func sqlQueryMediaVersionById(id string) map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	sql := `select id,
      serveraddr,
	  filesrv,
	  ftp,
	  mbs,
	  nginx,
	  update_time from server_version where id =? and type='media'`

	rows, err := db.Query(sql, id)
	CheckErr(err)
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

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
	}
	db.Close()
	return record
}
