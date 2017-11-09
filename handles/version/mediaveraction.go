package version

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wangfmD/rvs/setting"
	"github.com/wangfmD/rvs/sshv"
	"log"
	"net/http"
)

func AddMediaVersion(c *gin.Context) {
	addr := c.Param("addr")
	// 校验 addr todo
	var m map[string]string
	var err error

	f := func() (*setting.ServerOs, error) {
		p := setting.GetMediaAddrs()
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
			"stauts": "failed",
			"msg":    m,
		})

	} else {
		// err = sqlAddversion(addr, m)
		err = sqlAddMediaversion(addr, m)
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"stauts": "success",
			"msg":    m,
		})
	}
}

// QueryServerAddrs ...
func QueryMediaServerAddrs(c *gin.Context) {
	p := setting.GetMediaAddrs()
	addrs := make([]string, 0, 0)
	for _, v := range p {

		addrs = append(addrs, v.Addr)
	}
	log.Println(addrs)
	c.JSON(http.StatusOK, gin.H{
		"addrs": addrs,
	})
}

func sqlAddMediaversion(addr string, m map[string]string) error {
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
	  filesrv,
	  ftp,
	  mbs,
	  nginx,
	  update_time)
	VALUES(?,?,?,?,?,?,?,now());`

	rs, err := db.Exec(sql, "last", addr, "media", m["filesrv"], m["ftp"], m["mbs"], m["nginx"])
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

func sqlQueryMediaVersion() []map[string]string {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	CheckErr(err)
	sql := `select id,
      serveraddr,
	  filesrv,
	  ftp,
	  mbs,
	  nginx,
	  update_time from server_version where id ='last' and type='media'`

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

func QueryMediaVersion(c *gin.Context) {
	res := sqlQueryMediaVersion()
	// log.Fatal(res)
	log.Println("query")
	c.JSON(http.StatusOK, gin.H{
		"versions": res,
	})
}
