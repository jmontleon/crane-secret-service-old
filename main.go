package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// "k8s.io/client-go/rest"
	// "k8s.io/client-go/tools/auth"
)

const (
	bearerSchema string = "Bearer"
	apiServer    string = "https://kubernetes.default.svc"
	port         string = ":8443"
)

func main() {
	log.Printf("starting...")
	r := gin.Default()
	r.POST("/secrets", func(c *gin.Context) {

		log.Printf("auth header %v", c.Request.Header.Get("Authorization"))
		for name, values := range c.Request.Header {
			for _, value := range values {
				log.Println(name, value)
			}
		}


// 		namespace, _ := c.Params.Get("namespace")
// 		name, _ := c.Params.Get("name")

// 		authHeader := c.Request.Header.Get("Authorization")
// 		if authHeader == "" {
// 			log.Println("No auth header provided")
// 			return
// 		}

// 		token := authHeader[len(bearerSchema):]

		// log.Printf("auth header %v", c.Request.Header.Get("Authorization"))
		// for name, values := range c.Request.Header {
		// 	for _, value := range values {
		// 		log.Println(name, value)
		// 	}
		// }

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
