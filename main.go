package main

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
//type Message struct {
//	From     string
//	To       string
//	Text     string
//	Validity int
//	Priority int
//}
//
//func sendMessage(c *gin.Context) {
//	var body Message
//
//	err := json.NewDecoder(c.Request.Body).Decode(&body)
//
//	if err != nil {
//		c.JSON(400, gin.H{
//			"message": "Something went wrong",
//		})
//	}
//
//	fmt.Println(body)
//}
//
//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//
//	r.POST("/sendMessage", sendMessage)
//	r.Run(":8080")
//}
