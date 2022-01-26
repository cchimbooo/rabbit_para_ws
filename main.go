package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	godotenv.Load()
	time.Sleep(20 * time.Second)
	r := gin.Default()

	melody := NovoWs()

	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		if err := melody.Ms.HandleRequest(c.Writer, c.Request); err != nil {
			log.Fatalln(err.Error())
		}
	})

	go ConsumerRabbit(Consumir(melody), Handle)

	r.Run(":5000")
}
