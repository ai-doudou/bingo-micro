/**
 * @Author : jinchunguang
 * @Date : 19-10-31 下午3:52
 * @Project : bingo-micro
 */
package util

import (
    "github.com/gin-gonic/gin"
    "github.com/unknwon/com"
    "bingo-micro/pkg/setting"
)

func GetPage(c *gin.Context) int {
    result := 0
    page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        result = (page - 1) * setting.PageSize
    }

    return result
}