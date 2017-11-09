package models

import (
	"time"
)

type Versions struct {
	Id               string    `json:"id"`
	ServerAddr       string    `json:"serveraddr"`
	Agent            string    `json:"agent"`
	Interact         string    `json:"interact"`
	InteractBusiness string    `json:"interactbusiness"`
	JyCenter         string    `json:"jycenter"`
	JyResource       string    `json:"jyresource"`
	MiddleCas        string    `json:"middlecas"`
	MiddleCenter     string    `json:"middlecenter"`
	MiddleCenterFile string    `json:"middlecenterfile"`
	MiddleCenterres  string    `json:"middlecenterres"`
	MiddleClient     string    `json:"middleclient"`
	MiddleDriver     string    `json:"middledriver"`
	MiddleResource   string    `json:"middleresource"`
	MiddleWareMcu    string    `json:"middleware_mcu"`
	Mysql            string    `json:"mysql"`
	Nginx            string    `json:"nginx"`
	OpenFire         string    `json:"openfire"`
	Redis            string    `json:"redis"`
	FileSrv          string    `json:"redis"`
	Ftp              string    `json:"redis"`
	FileSrv          string    `json:"redis"`
	Ftp              string    `json:"redis"`
	MiddleDatabase   string    `json:"redis"`
	MbsStorage       string    `json:"mbs-storage"`
	MbsManager       string    `json:"mbs-manager"`
	MbsTransfer      string    `json:"mbs-transfer"`
	MbsJobmgrDoc     string    `json:"mbs-jobmgr-doc"`
	MbsHttp          string    `json:"mbs-http"`
	MbsRecording     string    `json:"mbs-recording"`
	MbsStreaming     string    `json:"mbs-streaming"`
	MbsJobmgr        string    `json:"mbs-jobmgr"`
	UpdateTime       time.Time `json:"middleware_mcu"`
}
