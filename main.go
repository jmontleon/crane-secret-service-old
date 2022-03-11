package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// "k8s.io/client-go/rest"
	// "k8s.io/client-go/tools/auth"
)

const (
	port = ":8080"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// log.Printf("auth header %v", c.Request.Header.Get("Authorization"))
		for name, values := range c.Request.Header {
			for _, value := range values {
				log.Println(name, value)
			}
		}

		// info := auth.Info{
		// 	User: "foo",
		// 	BearerToken: "1234",
		// 	Insecure: true,
		// }
		// clientConfig := rest.Config{}
		// clientConfig.Host = "example.com:4901"
		// clientConfig = info.MergeWithConfig()
		// client := client.New(clientConfig)
		// client.Pods(ns).List()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(port)
}
