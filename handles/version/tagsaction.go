package version

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/wangfmD/rvs/log"
	"github.com/wangfmD/rvs/models"
	"log"
	"net/http"
	"strings"
)

func QueryTagsHandle(c *gin.Context) {
	var tag = models.VersionTag{}
	tags := tag.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"result":  tags,
		"message": "query",
	})

}

func AddTagsHandle(c *gin.Context) {
	tags := c.PostForm("tags")
	versionid := c.PostForm("versionid")
	log.Println("tags is :", tags)
	log.Println("versionid is :", versionid)
	m := analysisString(tags)
	bData, err := json.Marshal(m)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	} else {
		log.Printf("tags is : %s\n", bData)
		tag := models.VersionTag{}

		err := json.Unmarshal(bData, &tag)
		tag.Versionid = versionid
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": err,
			})
		}
		num, err := tag.Add()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": err,
			})
		}
		log.Println("add tags total:", num)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": m,
		})
	}
}

func analysisString(s string) map[string]string {
	m := make(map[string]string)
	s1 := strings.TrimSpace(s)
	ss := strings.Split(s1, "\n")
	for _, value := range ss {
		log.Println("line split: ", value)
		tag := strings.TrimSpace(value)
		ss1 := strings.Split(tag, "=")
		log.Println("= split: ", ss1)
		m[strings.TrimSpace(ss1[0])] = strings.TrimSpace(ss1[1])
	}
	return m
}

func DeleteHandle(c *gin.Context) {
	m := c.PostForm("tagsId")
	log.Println("delete versionid:", m)
	s := make([]string, 5)
	err := json.Unmarshal([]byte(m), &s)
	if err != nil {
		log.Println(err)
	} else {
		for _, value := range s {
			tag := models.VersionTag{}
			tag.DeleteOne(value)
		}
	}
	c.JSON(http.StatusOK, m)
}

func GetTagsHandle(c *gin.Context) {
	tags := make([]string, 0)
	tagsM := models.GetVersionIds()
	for _, value := range tagsM {
		tags = append(tags, value.Versionid)
	}
	log.Println(tags)
	c.JSON(http.StatusOK, gin.H{
		"tags": tags,
	})
}
