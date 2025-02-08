package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Arg struct {
	Data string `form:"data" json:"data"`
}

var flagconf string

type ServerConfig struct {
	Port int `json:"port"`
}

var config struct {
	Server ServerConfig `json:"server"`
}

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func main() {
	if flagconf != "" {
		b, err := os.ReadFile(flagconf)
		if err != nil {
			log.Println("conf err:", err)
		} else {
			if err := yaml.Unmarshal(b, &config); err != nil {
				log.Println("yaml.Unmarshal err:", err)
			} else {
				log.Println("conf ", config.Server)
			}
		}
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		arg := &Arg{}
		if err := c.Bind(arg); err != nil {
			c.JSON(200, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "pong:" + arg.Data,
		})
	})
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
