/**
 * @Author : jinchunguang
 * @Date : 19-10-31 下午3:43
 * @Project : bingo-micro
 */
package main

import (
    "bingo-micro/pkg/setting"
    "bingo-micro/routers"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {

    router := routers.InitRouter()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()

}
