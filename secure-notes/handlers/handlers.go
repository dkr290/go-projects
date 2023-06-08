package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", "index.html")

}
