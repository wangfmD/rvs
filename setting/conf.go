package setting

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	"strings"
)

type ServerOs struct {
	Addr string
	Name string
	Pwd  string
}

// GetPlatformAddrs ...
func GetPlatformAddrs() []*ServerOs {
	cfg, err := goconfig.LoadConfigFile("/opt/gopath/src/github.com/wangfmD/rvs/conf/app.ini")
	// cfg, err := goconfig.LoadConfigFile("conf/app.ini")

	if err != nil {
		fmt.Println("Fail to load configuration file: " + err.Error())
		os.Exit(2)
	}

	os := cfg.GetSectionList()
	servers := make([]*ServerOs, 0, 0)

	for _, v := range os {
		var server ServerOs = ServerOs{}
		if strings.HasPrefix(v, "mgr_") {
			server.Addr = cfg.MustValue(v, "platform", "0.0.0.0")
			server.Name = cfg.MustValue(v, "name", "0.0.0.0")
			server.Pwd = cfg.MustValue(v, "pwd", "0.0.0.0")
			servers = append(servers, &server)
		}
		// fmt.Println(server)
	}

	return servers
}

func GetMediaAddrs() []*ServerOs {
	cfg, err := goconfig.LoadConfigFile("/opt/gopath/src/github.com/wangfmD/rvs/conf/app.ini")
	// cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		fmt.Println("Fail to load configuration file: " + err.Error())
		os.Exit(2)
	}

	os := cfg.GetSectionList()
	servers := make([]*ServerOs, 0, 0)

	for _, v := range os {
		var server ServerOs = ServerOs{}
		if strings.HasPrefix(v, "media_") {
			server.Addr = cfg.MustValue(v, "platform", "0.0.0.0")
			server.Name = cfg.MustValue(v, "name", "0.0.0.0")
			server.Pwd = cfg.MustValue(v, "pwd", "0.0.0.0")
			servers = append(servers, &server)
		}
		// fmt.Println(server)
	}
	return servers
}

func GetTagVersions() map[string]map[string]string {
	path := "/opt/gopath/src/github.com/wangfmD/rvs/conf/releaseversion/version.ini"
	// path:="conf/releaseversion/version.ini"
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		fmt.Println("Fail to load configuration file: " + err.Error())
		os.Exit(2)
	}
	ses := cfg.GetSectionList()
	versions := make(map[string]map[string]string)
	for _, se := range ses {
		versions[se] = fromCfgGetModelver(se, cfg)
	}
	fmt.Println(versions)
	return versions
}

func fromCfgGetModelver(se string, c *goconfig.ConfigFile) map[string]string {
	ver := make(map[string]string)
	ver["mysql"] = c.MustValue(se, "mysql", "")
	ver["nginx"] = c.MustValue(se, "nginx", "")
	ver["middledatabase"] = c.MustValue(se, "middledatabase", "")
	ver["middlecas"] = c.MustValue(se, "middlecas", "")
	ver["middlecenter"] = c.MustValue(se, "middlecenter", "")
	ver["middlecenterfile"] = c.MustValue(se, "middlecenterfile", "")
	ver["jycenter"] = c.MustValue(se, "jycenter", "")
	ver["middlecenterres"] = c.MustValue(se, "middlecenterres", "")
	ver["redis"] = c.MustValue(se, "redis", "")
	ver["interact"] = c.MustValue(se, "interact", "")
	ver["middleclient"] = c.MustValue(se, "middleclient", "")
	ver["middletherepair"] = c.MustValue(se, "middletherepair", "")
	ver["middledriver"] = c.MustValue(se, "middledriver", "")
	ver["middlewaremcu"] = c.MustValue(se, "middleware-mcu", "")
	ver["interactbusiness"] = c.MustValue(se, "interactbusiness", "")
	ver["openfire"] = c.MustValue(se, "openfire", "")
	ver["middleresource"] = c.MustValue(se, "middleresource", "")
	ver["jyresource"] = c.MustValue(se, "jyresource", "")
	ver["ftp"] = c.MustValue(se, "ftp", "")
	ver["filesrv"] = c.MustValue(se, "filesrv", "")
	ver["mbs"] = c.MustValue(se, "mbs", "")
	ver["agent"] = c.MustValue(se, "agent", "")
	return ver
}
