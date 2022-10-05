package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"simple-kv/cmd"
	"simple-kv/kv"

	"github.com/gin-gonic/gin"
)

var DEBUG string = "1"

func main() {
	cmd.Execute()
	if DEBUG == "1" {
		fmt.Printf("addr: %s, port %s\n", cmd.Addr, cmd.Port)
		gin.SetMode(gin.DebugMode)
	} else {
		fmt.Printf("Listening on %s:%s\n", cmd.Addr, cmd.Port)
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(tokenChecker())
	r.GET("/:token/:key", handleGet)
	r.POST("/:token/:key", handleSet)
	r.PUT("/:token/:key", handleSet)

	r.Run(cmd.Addr + ":" + cmd.Port)
}

func handleGet(c *gin.Context) {
	token := c.Param("token")
	key := c.Param("key")
	vaule := kv.GetValue(token, key)
	c.String(http.StatusOK, vaule)
}

func handleSet(c *gin.Context) {
	token := c.Param("token")
	key := c.Param("key")
	b, _ := io.ReadAll(c.Request.Body)
	value := string(b)
	if err := kv.SetValue(token, key, value); err != nil {
		c.String(http.StatusBadGateway, err.Error())
	} else {
		c.String(http.StatusOK, string(b))
	}

}

func tokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		matched, _ := regexp.MatchString("[^a-zA-Z0-9]", token)
		if len(token) < 32 || matched {
			c.AbortWithError(http.StatusBadRequest, errors.New("token is invalid"))
		}
		c.Next()
	}
}
