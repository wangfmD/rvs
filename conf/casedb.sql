DROP TABLE IF EXISTS `server_version`;
CREATE TABLE `server_version` (
  `id` varchar(36) NOT NULL,
  `serveraddr` varchar(36) NOT NULL,
  `type` varchar(20) DEFAULT NULL,
  `agent` varchar(100) DEFAULT NULL,
  `interact` varchar(100) DEFAULT NULL ,
  `interactbusiness` varchar(100) DEFAULT NULL ,
  `jycenter` varchar(100) DEFAULT NULL,
  `jyresource` varchar(100) DEFAULT NULL,
  `middlecas` varchar(100) DEFAULT NULL,
  `middlecenter` varchar(100) DEFAULT NULL,
  `middlecenterfile` varchar(100) DEFAULT NULL,
  `middlecenterres` varchar(100) DEFAULT NULL,
  `middleclient` varchar(100) DEFAULT NULL,
  `middledriver` varchar(100) DEFAULT NULL,
  `middleresource` varchar(100) DEFAULT NULL,
  `middlewaremcu` varchar(100) DEFAULT NULL,
  `middletherepair` varchar(100) DEFAULT NULL,
  `mysql` varchar(100) DEFAULT NULL,
  `nginx` varchar(100) DEFAULT NULL,
  `openfire` varchar(100) DEFAULT NULL,
  `redis` varchar(100) DEFAULT NULL,
  `middledatabase` varchar(100) DEFAULT NULL,
  `filesrv` varchar(100) DEFAULT NULL,
  `ftp` varchar(100) DEFAULT NULL,
  `mbs` varchar(100) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='版本表';


DROP TABLE IF EXISTS `case_info`;
CREATE TABLE `case_info` (
  `ID` varchar(36) NOT NULL,
  `CASE_NAME` varchar(255) DEFAULT NULL COMMENT '执行名称',
  `REPORT_PATH` varchar(255) DEFAULT NULL COMMENT '说明',
  `TYPE` tinyint(4) DEFAULT '0' COMMENT '0：，1：，2：，3：',
  `STATUS` tinyint(4) DEFAULT '0' COMMENT '0：开始，1：执行中，2：结束，3：异常',
  `START_TIME` datetime DEFAULT NULL COMMENT '数据时间',
  `STOP_TIME` datetime DEFAULT NULL COMMENT '数据时间',
  `DURATION` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用例执行概要 表';

DROP TABLE IF EXISTS `version_tag`;
CREATE TABLE `version_tag` (
  `versionid` varchar(36) NOT NULL,
  `agent` varchar(100) DEFAULT NULL,
  `interact` varchar(100) DEFAULT NULL ,
  `interactbusiness` varchar(100) DEFAULT NULL ,
  `jycenter` varchar(100) DEFAULT NULL,
  `jyresource` varchar(100) DEFAULT NULL,
  `middlecas` varchar(100) DEFAULT NULL,
  `middlecenter` varchar(100) DEFAULT NULL,
  `middlecenterfile` varchar(100) DEFAULT NULL,
  `middlecenterres` varchar(100) DEFAULT NULL,
  `middleclient` varchar(100) DEFAULT NULL,
  `middledriver` varchar(100) DEFAULT NULL,
  `middleresource` varchar(100) DEFAULT NULL,
  `middlewaremcu` varchar(100) DEFAULT NULL,
  `middletherepair` varchar(100) DEFAULT NULL,
  `mysql` varchar(100) DEFAULT NULL,
  `nginx` varchar(100) DEFAULT NULL,
  `openfire` varchar(100) DEFAULT NULL,
  `redis` varchar(100) DEFAULT NULL,
  `middledatabase` varchar(100) DEFAULT NULL,
  `filesrv` varchar(100) DEFAULT NULL,
  `ftp` varchar(100) DEFAULT NULL,
  `mbs` varchar(100) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='发布版本表';