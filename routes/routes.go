package main

import (
    // "log"
    // "net/http"

	"allergen_menu/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()     // ginのEngineインスタンスを生成
    r.GET("/", controller.SampleFunc) // ルーティング設定
    r.Run()                // サーバに接続。HTTPリクエストを受け付ける
}

// func sampleFunc(c *gin.Context) {
//     log.Println("I am sampleFunc")
//     c.JSON(http.StatusOK, gin.H{"message": "HELLO!"})
// }
