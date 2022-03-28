package main

import (
	"kursach/CORS"
	"kursach/Methods"
	"kursach/Shifr"
	"kursach/port"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {

	r := gin.Default()
	CORS.ConnectDB()
	r.Use(cors.Default())
	r.SetTrustedProxies([]string{"192.168.1.2"})
	v1 := r.Group("/v1")
	{
		v1.POST("/aes", Shifr.Aes1)
		v1.POST("/des", Shifr.Desc1)
		v1.GET("/portUDP", port.Quest2)
		v1.GET("/portTCP", port.Quest3)

	}
	v3 := r.Group("/v3")
	{
		v3.GET("/GetAuto", Methods.GetAllAuto)
		v3.POST("/AddAuto", Methods.AddAuto)
		v3.GET("/DelAuto/:id", Methods.DelAuto)
		v3.POST("/UpdateAuto/:id", Methods.UpdateAuto)
	}

	//fmt.Println("Server is listening...")
	r.Run(":8080")
}
