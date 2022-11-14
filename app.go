package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name   string
	Age    int
	Number string
}

var dreamEmployess = map[string]Person{}

func PostPerson(c *gin.Context) {
	var bodyRequest Person

	body := c.Request.Body
	x, err := ioutil.ReadAll(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "problems with JSON body - 1",
		})
	}

	errJson := json.Unmarshal(x, &bodyRequest)
	if errJson != nil {
		c.JSON(400, gin.H{
			"message": "problems with JSON body - 2",
		})
	}

	if _, ok := dreamEmployess[bodyRequest.Name]; ok {
		c.JSON(409, gin.H{
			"message": "Person already exist",
		})
	}

	dreamEmployess[bodyRequest.Name] = bodyRequest
	c.JSON(200, gin.H{
		"message": "Person successfully created",
	})
}

func AllPersons(c *gin.Context) {
	json_data, err := json.Marshal(dreamEmployess)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "something went wrong",
		})
	}

	c.JSON(200, gin.H{
		"json": json_data,
	})

	fmt.Println(dreamEmployess)

}

type BodyRequest struct {
	Messages []struct {
		From     string
		To       string
		Text     string
		Validity int
		Priority int
	}
}

//var body BodyRequest
//
//err := json.NewDecoder(c.Request.Body).Decode(&body)
//
//if err != nil {
//	c.JSON(400, gin.H{
//		"message": "Something went wrong",
//	})
//}
//
//var buf bytes.Buffer
//errBuf := json.NewEncoder(&buf).Encode(body)
//if errBuf != nil {
//	c.JSON(400, gin.H{
//		"message": errBuf,
//	})
//}

func sendMessage(c *gin.Context) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.devino.online/sms/messages", c.Request.Body)
	req.Header.Add("Authorization", "Key e6876ea2-d760-43bd-8024-445d586ce5ae")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Connection troubles",
		})
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	defer resp.Body.Close()

	c.JSON(200, gin.H{
		"message": string(bytes),
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/addPerson", PostPerson)
	r.GET("/getAllPersons", AllPersons)
	r.POST("/sendMessage", sendMessage)
	r.Run(":8080")
}
