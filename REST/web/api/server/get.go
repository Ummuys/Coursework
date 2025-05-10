package handlers

import (
	repos "coursework/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(g *gin.Context) {
	g.JSON(http.StatusOK, repos.Db.ToJson())
}
