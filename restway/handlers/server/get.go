package handlers

import (
	"net/http"

	repos "github.com/ummuys/courseproject/restway/repository"

	"github.com/gin-gonic/gin"
)

func Get(db repos.DataBase) gin.HandlerFunc {
	return func(g *gin.Context) {
		g.JSON(http.StatusOK, db.GetJSONFields())
	}
}
