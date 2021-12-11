package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TestPro struct {
	msgContent string
}

func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {
	g := gin.Default()
	g.GET("/index", func(c *gin.Context) {
		c.String(200, "ok")
	})

	g.Run(":8888")
}
