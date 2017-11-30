package models

import (
	// "encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type MyDate time.Time

func (self MyDate) MarshalJSON() ([]byte, error) {
	t := time.Time(self)
	if y := t.Year(); y < 0 || y >= 10000 {
		if y < 2000 {
			return []byte(`"2000-01-01 00:00:00"`), nil
		}
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(t.Format(`"2006-01-02 15:04:05"`)), nil
}

type VersionTag struct {
	Versionid        string `json:"versionid"`
	Agent            string `json:"agent"`
	Interact         string `json:"interact"`
	Interactbusiness string `json:"interactbusiness"`
	Jycenter         string `json:"jycenter"`
	Jyresource       string `json:"jyresource"`
	Middlecas        string `json:"middlecas"`
	Middlecenter     string `json:"middlecenter"`
	Middlecenterfile string `json:"middlecenterfile"`
	Middlecenterres  string `json:"middlecenterres"`
	Middleclient     string `json:"middleclient"`
	Middledriver     string `json:"middledriver"`
	Middleresource   string `json:"middleresource"`
	Middlewaremcu    string `json:"middleware-mcu`
	Middletherepair  string `json:"middletherepair"`
	Mysql            string `json:"mysql"`
	Nginx            string `json:"nginx"`
	Openfire         string `json:"openfire"`
	Redis            string `json:"redis"`
	Middledatabase   string `json:"middledatabase"`
	Filesrv          string `json:"filesrv"`
	Ftp              string `json:"ftp"`
	Mbs              string `json:"mbs"`
	Create_time      MyDate `json:"ctime" xorm:"created"`
	Update_time      MyDate `json:"ctime" xorm:"updated"`

	// Versionid        string    `xorm:"versionid varchar(36)"`
	// Agent            string    `xorm:"'agent' varchar(100)"`
	// Interact         string    `xorm:"'interact' varchar(100)"`
	// Interactbusiness string    `xorm:"'interactbusiness' varchar(100)"`
	// Jycenter         string    `xorm:"'jycenter' varchar(100)"`
	// Jyresource       string    `xorm:"'jyresource' varchar(100)"`
	// Middlecas        string    `xorm:"'middlecas' varchar(100)"`
	// Middlecenter     string    `xorm:"middlecenter varchar(36)"`
	// Middlecenterfile string    `xorm:"middlecenterfile varchar(36)"`
	// Middlecenterres  string    `xorm:"middlecenterres varchar(36)"`
	// Middleclient     string    `xorm:"middleclient varchar(36)"`
	// Middledriver     string    `xorm:"middledriver varchar(36)"`
	// Middleresource   string    `xorm:"middleresource varchar(36)"`
	// Middlewaremcu    string    `xorm:"middlewaremcu varchar(36)"`
	// Middletherepair  string    `xorm:"middletherepair varchar(36)"`
	// Mysql            string    `xorm:"mysql varchar(36)"`
	// Nginx            string    `xorm:"nginx varchar(36)"`
	// Openfire         string    `xorm:"openfire varchar(36)"`
	// Redis            string    `xorm:"redis varchar(36)"`
	// Middledatabase   string    `xorm:"middledatabase varchar(36)"`
	// Filesrv          string    `xorm:"filesrv varchar(36)"`
	// Ftp              string    `xorm:"ftp varchar(36)"`
	// Mbs              string    `xorm:"mbs varchar(36)"`
	// Create_time      time.Time `xorm:"create_time"`
	// Update_time      time.Time `xorm:"update_time"`
}

func (this *VersionTag) Add() (int64, error) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	num, err := engine.Insert(&(*this))
	return num, err
}

func GetJson() []map[string]string {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	// err := engine.Sync2(new(User))
	CheckErr(err)
	// ertr := engine.Sync2(new(User))
	// CheckErr(err)
	// engine.CreateTables(&case_info_t{})
	res, err := engine.QueryString("select * from case_info;")
	CheckErr(err)
	// fmt.Println(res)
	// jsonData, err := json.Marshal(res)
	// CheckErr(err)
	// fmt.Println(string(jsonData))
	// return string(jsonData)
	return res
}

func (this *VersionTag) GetAll() []VersionTag {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	engine.DatabaseTZ = time.UTC
	engine.TZLocation = time.UTC
	CheckErr(err)
	all := make([]VersionTag, 0)
	err = engine.Find(&all)
	CheckErr(err)
	if err != nil {
		return nil
	} else {
		return all
	}
}

func (this *VersionTag) DeleteOne(versionid string) bool {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	num, err := engine.Where("versionid=?", versionid).Delete(&(*this))
	if err != nil {
		return false
	} else {
		fmt.Println("delete num:", num)
		return true

	}
}

func GetVersionIds() []VersionTag {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/casedb?charset=utf8")
	this := make([]VersionTag, 0)
	err = engine.Cols("versionid").Find(&this)
	if err != nil {
		fmt.Println(err)
	}
	return this
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
