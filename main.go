package main

import (
	"io"
	"log"
	"time"

	"wxn/global"
	"wxn/weixin"

	"github.com/gin-gonic/gin"
)

func UpdateAccessTokenTimer() {
	timer := time.NewTimer(0)
	for {
		<-timer.C
		accessToken, expiresIn, err := weixin.GetAccessToken(global.Appid, global.AppSecret)
		if err == nil {
			global.AccessToken = accessToken
			timer.Reset(time.Duration(expiresIn-300) * time.Second)
			continue
		}

		log.Print("更新AccessToken失败，将于2s后重试")
		timer.Reset(2 * time.Second)
	}
}

func main() {
	go UpdateAccessTokenTimer()

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/invoke", func(c *gin.Context) {
		f := c.Request.URL.Query().Get("func")
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "Invalid JSON"})
			return
		}

		res, err := weixin.InvokeCloudFunc(global.AccessToken, global.EnvId, f, bodyBytes)
		if err != nil {
			c.JSON(200, gin.H{"code": 222, "msg": "参数异常"})
		} else {
			// var r map[string]interface{}
			// json.Unmarshal(res, &r)
			c.JSON(200, string(res))
		}
	})

	r.Run(":1126")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
