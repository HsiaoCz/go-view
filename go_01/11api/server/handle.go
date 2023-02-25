package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "please use number",
		})
	}
	user := s.store.GetUserByID(userID)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})
}
