package server

import (
	hand "coursework/web/api/server"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.GET("/students", hand.Get)
	g.Run()
}
