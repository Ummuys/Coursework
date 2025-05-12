package server

import (
	hand "github.com/ummuys/courseproject/restway/handlers/server"
	"github.com/ummuys/courseproject/restway/repository"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	db := repository.NewMapDB()
	db.SetFields(1000)
	db.ToJSON()

	gin.SetMode(gin.ReleaseMode)     // Отключаем вывод логов в CMD
	g := gin.New()                   // Создаем новый веб-сервер
	g.GET("/students", hand.Get(db)) // Регестрируем маршрут
	g.Run()                          // Запускаем сервер
}
