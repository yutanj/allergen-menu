// package main

// import (
//     // "log"
//     // "net/http"

// 	"allergen_menu/controller"
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()     // ginのEngineインスタンスを生成
//     r.GET("/", controller.SampleFunc) // ルーティング設定
//     r.Run()                // サーバに接続。HTTPリクエストを受け付ける
// }

// // func sampleFunc(c *gin.Context) {
// //     log.Println("I am sampleFunc")
// //     c.JSON(http.StatusOK, gin.H{"message": "HELLO!"})
// // }


package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
    r.GET("/get", func(c *gin.Context) {
		s := c.Query("str")
		n := c.Query("num")
		b := c.Query("bool")
		l := c.DefaultQuery("limit", "10")
		message := fmt.Sprintf("s: %v, n: %v, b: %v, l: %v", s, n, b, l)
		c.String(http.StatusOK, message)
  })
	r.Run()
}


// package main
 
// import (
//     "fmt"
//     "html/template"
//     "net/http"
//     //"allergen_menu/templates"
// )
 
// // 入力フォーム画面
// func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
//     // テンプレートをパースする
//     tpl := template.Must(template.ParseFiles("user-form.gtpl"))
 
//     // テンプレートに出力する値をマップにセット
//     values := map[string]string{}
 
//     // マップを展開してテンプレートを出力する
//     if err := tpl.ExecuteTemplate(w, "user-form.gtpl", values); err != nil {
//         fmt.Println(err)
//     }
// }
 
// func main() {
 
//     // "user-form"へのリクエストを関数で処理する
//     http.HandleFunc("/user-form", HandlerUserForm)
 
//     // サーバーを起動
//     http.ListenAndServe(":8080", nil)
// }