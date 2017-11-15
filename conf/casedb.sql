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
  `CASE_NAME` varchar(255) DEFAULT NULL,
  `REPORT_PATH` varchar(255) DEFAULT NULL,
  `exec_version` varchar(255) DEFAULT NULL,
  `TYPE` tinyint(4) DEFAULT '0',
  `STATUS` tinyint(4) DEFAULT '0',
  `START_TIME` datetime DEFAULT NULL,
  `STOP_TIME` datetime DEFAULT NULL,
  `DURATION` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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