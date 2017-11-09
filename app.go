package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangfmD/rvs/handles/releaseversion"
	"github.com/wangfmD/rvs/handles/report"
	"github.com/wangfmD/rvs/handles/version"
)

func main() {
	router := gin.Default()

	router.Static("/html", "html")

	router.GET("/api/v1/getinfo", report.GetApi)
	router.POST("/api/v1/add", report.AddCaseInfo)
	router.POST("/api/v1/update", report.UpdateInfo)
	// router.GET("/api/v1/getversion", report.GetVersion)

	router.GET("/api/v1/get/:addr", version.AddVersion)
	router.GET("/api/v1/getvs", version.QueryVersion)
	router.GET("/api/v1/getselects", version.QueryServerAddrs)

	router.GET("/api/v1/getmedia/:addr", version.AddMediaVersion)
	router.GET("/api/v1/getmediaselects", version.QueryMediaServerAddrs)
	router.GET("/api/v1/getmediavs", version.QueryMediaVersion)

	// router.GET("/api/v1/querytagv/:vertag", releaseversion.HandleQueryTagVers)
	// router.GET("/api/v1/queryipv/:addr", releaseversion.HandleQueryVersByIp)
	router.POST("/api/v1/querydiffv", releaseversion.HandleQueryDiffVers)
	router.GET("/api/v1/gettags", releaseversion.HandleQueryTags)
	router.GET("/api/v1/updatetags", releaseversion.HandleUpdateTagvers)

	router.Run(":8009")
}
